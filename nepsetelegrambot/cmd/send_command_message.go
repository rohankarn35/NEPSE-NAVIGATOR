package cmd

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rohankarn35/nepsemarketbot/services"
	"gorm.io/gorm"
)

func SendCommandMessage(bot *tgbotapi.BotAPI, chatID int64, updates tgbotapi.UpdatesChannel, db *gorm.DB, stockType string) {
	// Send initial message with buttons
	msg := tgbotapi.NewMessage(chatID, "Please choose an option:")
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Upcoming", "upcoming"),
			tgbotapi.NewInlineKeyboardButtonData("Open", "open"),
		),
	)
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
		return
	}

	// Handle button clicks using the existing updates channel
	for update := range updates {
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "open":
				services.SendCommandMessageService(bot, db, chatID, "Open", stockType)
			case "upcoming":
				services.SendCommandMessageService(bot, db, chatID, "Upcoming", stockType)

			default:
				continue
			}

			// Acknowledge callback
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			if _, err := bot.Request(callback); err != nil {
				log.Printf("Error sending callback: %v", err)
			}

			break // Exit after handling one callback
		}
	}
}
