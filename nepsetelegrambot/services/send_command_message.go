package services

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rohankarn35/htmlcapture"
	"github.com/rohankarn35/nepsemarketbot/applog" // Import the applog package
	ipodb "github.com/rohankarn35/nepsemarketbot/db"
	"github.com/rohankarn35/nepsemarketbot/utils"
	"gorm.io/gorm"
)

func SendCommandMessageService(bot *tgbotapi.BotAPI, db *gorm.DB, chatId int64, status string, stockType string) {

	datas := ipodb.ReadDB(db, status, stockType)

	if len(datas) == 0 {
		message := tgbotapi.NewMessage(chatId, "No "+" "+status+" "+stockType+" available for now.")
		if _, err := bot.Send(message); err != nil {
			applog.Log(applog.ERROR, "Error sending no data message: %v", err)
		}
		return
	}

	for _, data := range datas {
		ipoType := data.ShareType
		if data.ShareType == "ordinary" {
			ipoType = "General Public"
		}
		subs := GetIPOOverscribeData(data.StockSymbol)
		openingDate := ConvertDate(data.OpeningDateAD, data.OpeningDateBS)
		closingDate := ConvertDate(data.ClosingDateAD, data.ClosingDateBS)
		opts := htmlcapture.CaptureOptions{
			Input: "templates/ipoCommand.html",
			Variables: map[string]string{
				"CompanyName": data.CompanyName,
				"Title":       status + " " + data.Type + " Alert",
				"Subtitle":    "(" + "For " + utils.CapitalizeFirstLetter(ipoType) + ")",

				"IssueDate":    openingDate,
				"ClosingDate":  closingDate,
				"IssuePrice":   "Rs. " + data.PricePerUnit,
				"Sector":       data.SectorName,
				"Subscription": subs,
			},
			Selector:  ".container",
			ViewportW: 600,
			ViewportH: 600,
		}
		img, err := htmlcapture.Capture(opts)
		if err != nil {
			applog.Log(applog.ERROR, "Error capturing screenshot: %v", err)
			return
		}

		responseText := "Here is the latest " + ipoType + " information for " + data.CompanyName + "."
		photo := tgbotapi.NewPhoto(chatId, tgbotapi.FileBytes{Name: "ipoimage", Bytes: img})
		photo.Caption = responseText
		photo.ParseMode = "Markdown"

		if strings.ToLower(data.Status) == "open" {
			button1 := tgbotapi.NewInlineKeyboardButtonURL("APPLY HERE", "https://meroshare.cdsc.com.np/")
			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(button1),
			)
			photo.ReplyMarkup = inlineKeyboard
		}
		if _, err := bot.Send(photo); err != nil {
			applog.Log(applog.ERROR, "Error sending market summary image: %v", err)
		}

	}

}
