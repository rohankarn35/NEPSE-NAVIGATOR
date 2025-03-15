package cmd

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitializeDataBase(tokenKey string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(tokenKey)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up updates channel
	// updateConfig := tgbotapi.NewUpdate(0)
	// updateConfig.Timeout = 60
	// updates := bot.GetUpdatesChan(updateConfig)

	return bot

}
