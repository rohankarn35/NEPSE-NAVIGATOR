package cmd

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/machinebox/graphql"
	"github.com/robfig/cron/v3"
	"github.com/rohankarn35/htmlcapture"
	dbgraphql "github.com/rohankarn35/nepsemarketbot/graphql"
	"github.com/rohankarn35/nepsemarketbot/models"
	"github.com/rohankarn35/nepsemarketbot/services"
	"gorm.io/gorm"

	"github.com/rohankarn35/nepsemarketbot/applog"
	ipodb "github.com/rohankarn35/nepsemarketbot/db"
	gorm_models "github.com/rohankarn35/nepsemarketbot/db/models"
	"github.com/rohankarn35/nepsemarketbot/server"
)

func SendMessages(db *gorm.DB, c *cron.Cron, bot *tgbotapi.BotAPI, chatID int64, client *graphql.Client) error {

	// Fetch latest data
	ipoData, fpoData, err := dbgraphql.GetIPOFPODetails(client)
	if err != nil {
		return err
	}
	applog.Log(applog.INFO, "Fetched IPO and FPO data successfully")

	var IPOdata []models.IPOAlertModel
	for _, ipo := range ipoData {
		IPOdata = append(IPOdata, models.IPOAlertModel{
			UniqueSymbol: ipo.UniqueSymbol,

			CompanyName:            ipo.CompanyName,
			StockSymbol:            ipo.StockSymbol,
			ShareType:              ipo.ShareType,
			Status:                 ipo.Status,
			PricePerUnit:           ipo.PricePerUnit,
			MinUnits:               ipo.MinUnits,
			MaxUnits:               ipo.MaxUnits,
			OpeningDateAD:          ipo.OpeningDateAD,
			OpeningDateBS:          ipo.OpeningDateBS,
			ClosingDateAD:          ipo.ClosingDateAD,
			ClosingDateBS:          ipo.ClosingDateBS,
			ClosingDateClosingTime: ipo.ClosingDateClosingTime,
			ShareRegistrar:         ipo.ShareRegistrar,
			Rating:                 ipo.Rating,
			SectorName:             ipo.SectorName,
			Type:                   "IPO",
		})
	}

	for _, fpo := range fpoData {
		IPOdata = append(IPOdata, models.IPOAlertModel{
			UniqueSymbol:           fpo.UniqueSymbol,
			CompanyName:            fpo.CompanyName,
			StockSymbol:            fpo.StockSymbol,
			ShareType:              fpo.ShareType,
			Status:                 fpo.Status,
			PricePerUnit:           fpo.PricePerUnit,
			MinUnits:               fpo.MinUnits,
			MaxUnits:               fpo.MaxUnits,
			OpeningDateAD:          fpo.OpeningDateAD,
			OpeningDateBS:          fpo.OpeningDateBS,
			ClosingDateAD:          fpo.ClosingDateAD,
			ClosingDateBS:          fpo.ClosingDateBS,
			ClosingDateClosingTime: fpo.ClosingDateClosingTime,
			ShareRegistrar:         fpo.ShareRegistrar,
			Rating:                 fpo.Rating,
			SectorName:             fpo.SectorName,
			Type:                   "FPO",
		})
	}

	// Send IPO updates
	for _, ipo := range IPOdata {
		fmt.Printf("the status of %s is %s", ipo.StockSymbol, ipo.Status)
		if strings.EqualFold(ipo.Status, "nearing") || strings.EqualFold(ipo.Status, "open") || strings.EqualFold(ipo.Status, "upcoming") {
			isAvailable := ipodb.CheckAndUpdateIPOStatus(db, ipo.UniqueSymbol, ipo.Status)
			if isAvailable {
				createUpdateIpo := gorm_models.NepseData{
					UniqueSymbol:           ipo.UniqueSymbol,
					CompanyName:            ipo.CompanyName,
					StockSymbol:            ipo.StockSymbol,
					ShareRegistrar:         ipo.ShareRegistrar,
					SectorName:             ipo.SectorName,
					ShareType:              ipo.ShareType,
					PricePerUnit:           ipo.PricePerUnit,
					Rating:                 ipo.Rating,
					MinUnits:               ipo.MinUnits,
					MaxUnits:               ipo.MaxUnits,
					OpeningDateAD:          ipo.OpeningDateAD,
					OpeningDateBS:          ipo.OpeningDateBS,
					ClosingDateAD:          ipo.ClosingDateAD,
					ClosingDateBS:          ipo.ClosingDateBS,
					ClosingDateClosingTime: ipo.ClosingDateClosingTime,
					Status:                 ipo.Status,
					Type:                   ipo.Type,
				}

				ipodb.CreateOrUpdateDB(db, createUpdateIpo)
				server.Scheduler(ipo.ClosingDateAD, ipo.ClosingDateClosingTime, ipo, bot, c, chatID)
				var ipoCronData gorm_models.CronJob
				if strings.ToLower(ipo.Status) == "open" {
					ipoCronData = gorm_models.CronJob{
						UniqueSymbol:           ipo.UniqueSymbol,
						StockSymbol:            ipo.StockSymbol,
						OpeningDateAD:          ipo.OpeningDateAD,
						OpeningDateBS:          ipo.OpeningDateBS,
						ClosingDateAD:          ipo.ClosingDateAD,
						ClosingDateBS:          ipo.ClosingDateBS,
						ClosingDateClosingTime: ipo.ClosingDateClosingTime,
						Status:                 ipo.Status,
						NepseData: gorm_models.NepseData{
							UniqueSymbol:           ipo.UniqueSymbol,
							CompanyName:            ipo.CompanyName,
							StockSymbol:            ipo.StockSymbol,
							ShareRegistrar:         ipo.ShareRegistrar,
							SectorName:             ipo.SectorName,
							ShareType:              ipo.ShareType,
							PricePerUnit:           ipo.PricePerUnit,
							Rating:                 ipo.Rating,
							MinUnits:               ipo.MinUnits,
							MaxUnits:               ipo.MaxUnits,
							OpeningDateAD:          ipo.OpeningDateAD,
							OpeningDateBS:          ipo.OpeningDateBS,
							ClosingDateAD:          ipo.ClosingDateAD,
							ClosingDateBS:          ipo.ClosingDateBS,
							ClosingDateClosingTime: ipo.ClosingDateClosingTime,
							Status:                 ipo.Status,
							Type:                   ipo.Type,
						},
					}
					ipodb.StoreCron(db, ipoCronData)
				}
				ipoType := ipo.ShareType
				if ipo.ShareType == "ordinary" {
					ipoType = "General Public"
				}
				status := "Upcoming"
				if ipo.Status != "Nearing" {
					status = ipo.Status
				}
				openingDate := services.ConvertDate(ipo.OpeningDateAD, ipo.OpeningDateBS)
				closingDate := services.ConvertDate(ipo.ClosingDateAD, ipo.ClosingDateBS)
				opts := htmlcapture.CaptureOptions{
					Input: "templates/ipoAlert.html",
					Variables: map[string]string{
						"CompanyName": ipo.CompanyName,
						"Title":       status + " " + ipo.Type + " Alert",
						"Subtitle":    "(" + "For " + ipoType + ")",

						"IssueDate":   openingDate,
						"ClosingDate": closingDate,
						"IssuePrice":  "Rs. " + ipo.PricePerUnit,
						"Sector":      ipo.SectorName,
					},
					Selector:  ".container",
					ViewportW: 700,
					ViewportH: 600,
				}
				img, err := htmlcapture.Capture(opts)
				if err != nil {
					applog.Log(applog.ERROR, "Error capturing screenshot: %v", err)
				}

				// Prepare the message text
				responseText := services.FormatIPOMessage(ipo)

				// Send the photo
				photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "ipoimage", Bytes: img})
				photo.Caption = responseText
				photo.ParseMode = "Markdown"

				// If IPO is open, add a button
				if strings.ToLower(ipo.Status) == "open" {
					button1 := tgbotapi.NewInlineKeyboardButtonURL("APPLY HERE", "https://meroshare.cdsc.com.np/")
					inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(button1),
					)
					photo.ReplyMarkup = inlineKeyboard
				}

				// Send the photo with caption and button
				if _, err := bot.Send(photo); err != nil {
					applog.Log(applog.ERROR, "Error sending IPO image: %v", err)
					continue
				}

				applog.Log(applog.INFO, "Successfully sent IPO image to chat ID: %d", chatID)
			} else {
				applog.Log(applog.INFO, "Message Already Sent: %d", chatID)
			}
		} else {
			isAvailable := ipodb.CheckAndUpdateIPOStatus(db, ipo.UniqueSymbol, ipo.Status)
			if isAvailable {
				createUpdateIpo := gorm_models.NepseData{
					UniqueSymbol:           ipo.UniqueSymbol,
					CompanyName:            ipo.CompanyName,
					StockSymbol:            ipo.StockSymbol,
					ShareRegistrar:         ipo.ShareRegistrar,
					SectorName:             ipo.SectorName,
					ShareType:              ipo.ShareType,
					PricePerUnit:           ipo.PricePerUnit,
					Rating:                 ipo.Rating,
					MinUnits:               ipo.MinUnits,
					MaxUnits:               ipo.MaxUnits,
					OpeningDateAD:          ipo.OpeningDateAD,
					OpeningDateBS:          ipo.OpeningDateBS,
					ClosingDateAD:          ipo.ClosingDateAD,
					ClosingDateBS:          ipo.ClosingDateBS,
					ClosingDateClosingTime: ipo.ClosingDateClosingTime,
					Status:                 ipo.Status,
					Type:                   ipo.Type,
				}

				ipodb.CreateOrUpdateDB(db, createUpdateIpo)
			}

		}
	}
	return nil
}

func ScheduleSendMessage(db *gorm.DB, c *cron.Cron, bot *tgbotapi.BotAPI, chatID int64, client *graphql.Client) {

	if err := SendMessages(db, c, bot, chatID, client); err != nil {
		applog.Log(applog.ERROR, "Error sending messages: %v", err)
		if err := SendMessages(db, c, bot, chatID, client); err != nil {
			applog.Log(applog.ERROR, "Error sending messages on retry: %v", err)
		}
	}

	_, err := c.AddFunc("0 */2 9-17 * *", func() {
		if err := SendMessages(db, c, bot, chatID, client); err != nil {
			applog.Log(applog.ERROR, "Error sending messages: %v", err)
			if err := SendMessages(db, c, bot, chatID, client); err != nil {
				applog.Log(applog.ERROR, "Error sending messages on retry: %v", err)
			}
		}
	})
	if err != nil {
		applog.Log(applog.ERROR, "Error scheduling send messages: %v", err)
	}
}
