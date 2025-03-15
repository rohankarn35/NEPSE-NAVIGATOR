package cmd

import (
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func BotMessageReply(bot *tgbotapi.BotAPI, db *gorm.DB) {
	// Register commands
	commands := []tgbotapi.BotCommand{
		{Command: "ipo", Description: "See the IPO Details"},
		{Command: "fpo", Description: "See the FPO details"},
		{Command: "help", Description: "Show help information"},
	}

	if _, err := bot.Request(tgbotapi.NewSetMyCommands(commands...)); err != nil {
		log.Printf("Error setting commands: %v", err)
	}

	// Create single updates channel
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	// Process all updates in one loop
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "ipo":
				var wg sync.WaitGroup
				wg.Add(1)
				go func(upd tgbotapi.Update) {
					defer wg.Done()
					SendCommandMessage(bot, upd.Message.Chat.ID, updates, db, "IPO")
				}(update)
				wg.Wait()
			case "fpo":
				var wg sync.WaitGroup
				wg.Add(1)
				go func(upd tgbotapi.Update) {
					defer wg.Done()
					SendCommandMessage(bot, upd.Message.Chat.ID, updates, db, "FPO")
				}(update)
				wg.Wait()
			case "help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here is the help information...")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use valid command")
				bot.Send(msg)
			}
		}
	}
}
