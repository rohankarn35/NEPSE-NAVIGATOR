package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rohankarn35/nepsemarketbot/applog"
	"gorm.io/gorm"
)

// SendCommandMessage sends the initial message with buttons
func SendCommandMessage(bot *tgbotapi.BotAPI, chatID int64, db *gorm.DB, stockType string) {
	// Initialize logger (optional, could reuse bot.log if configured globally)
	err := applog.InitLogger("app.log", applog.DEBUG)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to initialize logger: %v", err)
		return
	}
	defer applog.CloseLogger()

	// Send initial message with buttons
	msg := tgbotapi.NewMessage(chatID, "Please choose an option for "+stockType+":")
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Upcoming", "upcoming"),
			tgbotapi.NewInlineKeyboardButtonData("Open", "open"),
		),
	)
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		applog.Log(applog.ERROR, "Error sending message: %v", err)
	}
}

// getStockTypeFromContext infers stock type from the original command (simplified approach)
func getStockTypeFromContext(update tgbotapi.Update) string {
	// Since callbacks don't carry the original command, we assume context from the message text
	// This is a limitation; ideally, store state or pass it via callback data
	if update.CallbackQuery != nil && update.CallbackQuery.Message != nil {
		if update.CallbackQuery.Message.ReplyToMessage != nil {
			switch update.CallbackQuery.Message.ReplyToMessage.Text {
			case "Please choose an option for IPO:":
				return "IPO"
			case "Please choose an option for FPO:":
				return "FPO"
			}
		}
	}
	return "IPO" // Default fallback (could be improved with better state management)
}
