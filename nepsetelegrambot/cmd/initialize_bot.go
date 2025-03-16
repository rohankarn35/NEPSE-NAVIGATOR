package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rohankarn35/nepsemarketbot/applog"
)

func InitializeDataBase(tokenKey string) *tgbotapi.BotAPI {
	// Initialize the logger
	err := applog.InitLogger("bot.log", applog.DEBUG)
	if err != nil {
		applog.Log(applog.ERROR, "Error initializing logger: %v", err)
		return nil
	}
	defer applog.CloseLogger()

	bot, err := tgbotapi.NewBotAPI(tokenKey)
	if err != nil {
		applog.Log(applog.ERROR, "Error creating bot: %v", err)
		return nil
	}
	bot.Debug = true
	applog.Log(applog.INFO, "Authorized on account %s", bot.Self.UserName)

	// Set up updates channel
	// updateConfig := tgbotapi.NewUpdate(0)
	// updateConfig.Timeout = 60
	// updates := bot.GetUpdatesChan(updateConfig)

	return bot
}
