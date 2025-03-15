package server

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/machinebox/graphql"
	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/robfig/cron/v3"
	"github.com/rohankarn35/htmlcapture"
	dbgraphql "github.com/rohankarn35/nepsemarketbot/graphql"
	"github.com/rohankarn35/nepsemarketbot/services"
)

func ScheduleMarketSummary(bot *tgbotapi.BotAPI, c *cron.Cron, chatID int64, client *graphql.Client) {

	_, err := c.AddFunc("0 15 * * *", func() {
		SendMarketSummaryMessage(bot, chatID, client)
	})
	if err != nil {
		log.Printf("Error scheduling market summary: %v", err)
		return
	}

	log.Print("Market Summary Scheduled")

}

func SendMarketSummaryMessage(bot *tgbotapi.BotAPI, chatID int64, client *graphql.Client) {

	marketSummary, err := dbgraphql.MarketSummary(client)
	if err != nil {
		log.Printf("Error fetching market summary: %v", err)
		return
	}
	if marketSummary.MarketStatus.IsMarketOpen {
		datetimeStr := "2025-03-12T15:00:00"

		datetime, err := time.Parse("2006-01-02T15:04:05", datetimeStr)
		if err != nil {
			log.Fatalf("Error converting datetime string to time.Time: %v", err)
		}

		er, err := nepalitime.FromEnglishTime(datetime)
		if err != nil {
			log.Fatalf("Error parsing time: %v", err)
		}
		nep := er.String()[:10]
		opts := htmlcapture.CaptureOptions{
			Input: "templates/marketSummary.html",
			Variables: map[string]string{

				"Date":               services.BSDateConvert(nep),
				"IndexPoint":         fmt.Sprintf("%.2f", marketSummary.NepseIndex.IndexValue),
				"PointChange":        fmt.Sprintf("%.2f", marketSummary.NepseIndex.Difference),
				"PercentageChange":   fmt.Sprintf("%.2f%%", marketSummary.NepseIndex.PercentChange),
				"Turnover":           fmt.Sprintf("%.2f", marketSummary.NepseIndex.Turnover),
				"ShareTraded":        fmt.Sprintf("%d", marketSummary.NepseIndex.Volume),
				"Sector1":            marketSummary.Indices[0].IndexName,
				"PointSector1":       fmt.Sprintf("%.2f", marketSummary.Indices[0].Difference),
				"PonitChangeSector1": fmt.Sprintf("%.2f%%", marketSummary.Indices[0].PercentChange),
				"Sector2":            marketSummary.Indices[1].IndexName,
				"PointSector2":       fmt.Sprintf("%.2f", marketSummary.Indices[1].Difference),
				"PonitChangeSector2": fmt.Sprintf("%.2f%%", marketSummary.Indices[1].PercentChange),
				"Sector3":            marketSummary.Indices[2].IndexName,
				"PointSector3":       fmt.Sprintf("%.2f", marketSummary.Indices[2].Difference),
				"PonitChangeSector3": fmt.Sprintf("%.2f%%", marketSummary.Indices[2].PercentChange),
				"GainerName1":        marketSummary.MarketMovers.Gainers[0].StockSymbol,
				"GainerPoint1":       fmt.Sprintf("%.2f", marketSummary.MarketMovers.Gainers[0].Amount),
				"GainerPointChange1": fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Gainers[0].PercentChange),
				"GainerName2":        marketSummary.MarketMovers.Gainers[1].StockSymbol,
				"GainerPoint2":       fmt.Sprintf("%.2f", marketSummary.MarketMovers.Gainers[1].Amount),
				"GainerPointChange2": fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Gainers[1].PercentChange),
				"GainerName3":        marketSummary.MarketMovers.Gainers[2].StockSymbol,
				"GainerPoint3":       fmt.Sprintf("%.2f", marketSummary.MarketMovers.Gainers[2].Amount),
				"GainerPointChange3": fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Gainers[2].PercentChange),
				"GainerName4":        marketSummary.MarketMovers.Gainers[3].StockSymbol,
				"GainerPoint4":       fmt.Sprintf("%.2f", marketSummary.MarketMovers.Gainers[3].Amount),
				"GainerPointChange4": fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Gainers[3].PercentChange),
				"GainerName5":        marketSummary.MarketMovers.Gainers[4].StockSymbol,
				"GainerPoint5":       fmt.Sprintf("%.2f", marketSummary.MarketMovers.Gainers[4].Amount),
				"GainerPointChange5": fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Gainers[4].PercentChange),
				"LoserName1":         marketSummary.MarketMovers.Losers[0].StockSymbol,
				"LoserPoint1":        fmt.Sprintf("%.2f", marketSummary.MarketMovers.Losers[0].Amount),
				"LoserPointChange1":  fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Losers[0].PercentChange),
				"LoserName2":         marketSummary.MarketMovers.Losers[1].StockSymbol,
				"LoserPoint2":        fmt.Sprintf("%.2f", marketSummary.MarketMovers.Losers[1].Amount),
				"LoserPointChange2":  fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Losers[1].PercentChange),
				"LoserName3":         marketSummary.MarketMovers.Losers[2].StockSymbol,
				"LoserPoint3":        fmt.Sprintf("%.2f", marketSummary.MarketMovers.Losers[2].Amount),
				"LoserPointChange3":  fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Losers[2].PercentChange),
				"LoserName4":         marketSummary.MarketMovers.Losers[3].StockSymbol,
				"LoserPoint4":        fmt.Sprintf("%.2f", marketSummary.MarketMovers.Losers[3].Amount),
				"LoserPointChange4":  fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Losers[3].PercentChange),
				"LoserName5":         marketSummary.MarketMovers.Losers[4].StockSymbol,
				"LoserPoint5":        fmt.Sprintf("%.2f", marketSummary.MarketMovers.Losers[4].Amount),
				"LoserPointChange5":  fmt.Sprintf("%.2f%%", marketSummary.MarketMovers.Losers[4].PercentChange),
			},
			Selector:  ".container",
			ViewportW: 900,
			ViewportH: 1200,
		}

		img, err := htmlcapture.Capture(opts)
		if err != nil {
			log.Fatalf("Error capturing screenshot: %v", err)
		}

		responseText := `📊 NEPSE Daily Market Summary - ` + services.BSDateConvert(nep) + `

📈 NEPSE Index: ` + fmt.Sprintf("%.2f", marketSummary.NepseIndex.IndexValue) + ` (` + fmt.Sprintf("%.2f", marketSummary.NepseIndex.Difference) + ` | ` + fmt.Sprintf("%.2f%%", marketSummary.NepseIndex.PercentChange) + `)
💰 Total Turnover: Rs ` + fmt.Sprintf("%.2f", marketSummary.NepseIndex.Turnover) + `
📉 Total Traded Shares: ` + fmt.Sprintf("%d", marketSummary.NepseIndex.Volume)

		photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "market_summary.png", Bytes: img})
		photo.Caption = responseText
		photo.ParseMode = "Markdown"

		if _, err := bot.Send(photo); err != nil {
			log.Printf("Error sending market summary image: %v", err)
		}
	} else {
		log.Print("Market Closed")
	}
}
