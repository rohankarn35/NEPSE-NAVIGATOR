package cmd

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/services"
	"gorm.io/gorm"
)

func BotMessageReply(bot *tgbotapi.BotAPI, db *gorm.DB, nepsenavigatorChatID int64) {
	// Initialize the logger
	err := applog.InitLogger("bot.log", applog.DEBUG)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to initialize logger: %v", err)
		return
	}
	defer applog.CloseLogger()

	// Register commands
	commands := []tgbotapi.BotCommand{
		{Command: "ipo", Description: "See the IPO Details"},
		{Command: "fpo", Description: "See the FPO details"},
		{Command: "help", Description: "Show help information"},
	}
	if _, err := bot.Request(tgbotapi.NewSetMyCommands(commands...)); err != nil {
		applog.Log(applog.ERROR, "Error setting commands: %v", err)
	}

	// Create updates channel
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	// Use a WaitGroup to keep the main goroutine alive
	var wg sync.WaitGroup

	// Process all updates concurrently
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil { // Ignore non-message, non-callback updates
			continue
		}

		wg.Add(1)
		go func(upd tgbotapi.Update) {
			defer wg.Done()

			// Handle commands
			if upd.Message != nil && upd.Message.IsCommand() {
				switch upd.Message.Command() {
				case "ipo":
					handleCommand(bot, db, nepsenavigatorChatID, upd, "IPO")

				case "fpo":
					handleCommand(bot, db, nepsenavigatorChatID, upd, "FPO")

				case "help":
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Here are the Commands:\n/ipo - See the IPO Details\n"+
						"/fpo - See the FPO details\n"+
						"/help - Show help information")
					if _, err := bot.Send(msg); err != nil {
						applog.Log(applog.ERROR, "Error sending help message: %v", err)
					}

				case "start":
					welcomeText := "Welcome to NEPSE Navigator! Click the button below to join our channel."
					joinButton := tgbotapi.NewInlineKeyboardButtonURL("Join NEPSE Navigator", "https://t.me/nepsenavigator")
					keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(joinButton))
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, welcomeText)
					msg.ReplyMarkup = keyboard
					if _, err := bot.Send(msg); err != nil {
						applog.Log(applog.ERROR, "Error sending start message: %v", err)
					}

				default:
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Use a valid command")
					if _, err := bot.Send(msg); err != nil {
						applog.Log(applog.ERROR, "Error sending invalid command message: %v", err)
					}
				}
			}

			// Handle callback queries (button presses)
			if upd.CallbackQuery != nil {
				chatID := upd.CallbackQuery.Message.Chat.ID
				switch upd.CallbackQuery.Data {
				case "open":
					services.SendCommandMessageService(bot, db, chatID, "Open", getStockTypeFromContext(upd))
				case "upcoming":
					services.SendCommandMessageService(bot, db, chatID, "Upcoming", getStockTypeFromContext(upd))
				default:
					applog.Log(applog.DEBUG, "Unknown callback data: %s", upd.CallbackQuery.Data)
				}

				// Acknowledge callback
				callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, "")
				if _, err := bot.Request(callback); err != nil {
					applog.Log(applog.ERROR, "Error sending callback acknowledgment: %v", err)
				}
			}
		}(update)
	}

	// Wait for all goroutines (only if updates channel closes)
	wg.Wait()
}

// handleCommand processes /ipo and /fpo commands
func handleCommand(bot *tgbotapi.BotAPI, db *gorm.DB, nepsenavigatorChatID int64, update tgbotapi.Update, stockType string) {
	// Check if the user has joined the NEPSE Navigator channel
	chatMember, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: nepsenavigatorChatID,
			UserID: update.Message.From.ID,
		},
	})
	if err != nil {
		applog.Log(applog.ERROR, "Error checking chat member status: %v", err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "An error occurred while checking your membership status.")
		bot.Send(msg)
		return
	}

	if chatMember.Status == "left" || chatMember.Status == "kicked" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please join the NEPSE Navigator channel to use this command. After joining, try /"+stockType+" again.")
		joinButton := tgbotapi.NewInlineKeyboardButtonURL("Join NEPSE Navigator", "https://t.me/nepsenavigator")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(joinButton))
		msg.ReplyMarkup = keyboard
		if _, err := bot.Send(msg); err != nil {
			applog.Log(applog.ERROR, "Error sending join message: %v", err)
		}
	} else {
		SendCommandMessage(bot, update.Message.Chat.ID, db, stockType)
	}
}
