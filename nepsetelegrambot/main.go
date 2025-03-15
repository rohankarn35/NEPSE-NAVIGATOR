package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"github.com/robfig/cron/v3"
	"github.com/rohankarn35/nepsemarketbot/cmd"
	ipodb "github.com/rohankarn35/nepsemarketbot/db"
	"github.com/rohankarn35/nepsemarketbot/server"
)

func main() {

	// Load environment variables from .env file
	attempts := 3
	for attempts > 0 {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
			attempts--
			continue

		}
		var (
			botToken  string
			chatIDStr string
		)
		isTest := os.Getenv("DEV_TYPE")
		api_url := os.Getenv("GRAPHQL_API")
		dsn := os.Getenv("DATABASE_URL")

		if isTest == "test" {
			botToken = os.Getenv("TEST_BOT_TOKEN")
			chatIDStr = os.Getenv("TEST_CHAT_ID")
		} else {
			botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
			chatIDStr = os.Getenv("CHAT_ID")
		}
		chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
		if err != nil {
			log.Printf("Error converting TEST_CHAT_ID to int64: %v", err)
			attempts--
			continue
		}

		// Connect to PostgreSQL
		db := cmd.InitializeDb(dsn)
		if db == nil {
			log.Printf("Error initializing database")
			attempts--
			continue
		}
		// Connect to local GraphQL server
		client := graphql.NewClient(api_url)
		if client == nil {
			log.Printf("Error initializing GraphQL client")
			attempts--
			continue
		}
		fmt.Print(client)
		ipodb.ReadCron(db)
		c := cron.New(cron.WithLocation(time.FixedZone("NPT", 5*3600+45*60)))

		//initializebot
		bot := cmd.InitializeDataBase(botToken)
		if bot == nil {
			log.Printf("Error initializing bot")
			attempts--
			continue
		}

		server.InitializeScheduleronRestart(bot, c, db, chatID)
		// cmd.BotMessageReply(bot)

		server.ScheduleMarketSummary(bot, c, chatID, client)
		server.SendMarketSummaryMessage(bot, chatID, client)

		// Add initial message to show bot is running
		log.Println("Bot started and waiting for messages...")
		cmd.ScheduleSendMessage(db, c, bot, chatID, client)

		c.Start()

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmd.BotMessageReply(bot, db)
		}()
		wg.Wait()

		// Keep the program running
		select {}
	}
	log.Fatalf("Failed to start the bot after 3 attempts")
}
