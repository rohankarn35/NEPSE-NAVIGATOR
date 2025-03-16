package server

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/machinebox/graphql"
	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/robfig/cron/v3"
	"github.com/rohankarn35/htmlcapture"
	"github.com/rohankarn35/nepsemarketbot/applog"
	dbgraphql "github.com/rohankarn35/nepsemarketbot/graphql"
	"github.com/rohankarn35/nepsemarketbot/services"
	"github.com/rohankarn35/nepsemarketbot/utils"
)

func ScheduleMarketSummary(bot *tgbotapi.BotAPI, c *cron.Cron, chatID int64, client *graphql.Client) {

	var isMarketOpen bool

	_, err := c.AddFunc("58 14 * * 0-4", func() {
		marketSummary, err := dbgraphql.MarketSummary(client)
		if err != nil {
			applog.Log(applog.ERROR, "Error fetching market summary: %v", err)
			return
		}
		isMarketOpen = marketSummary.MarketStatus.IsMarketOpen
		applog.Log(applog.INFO, "Market status fetched at 2:58 PM: %v", isMarketOpen)
	})
	if err != nil {
		applog.Log(applog.ERROR, "Error scheduling market status fetch: %v", err)
		return
	}

	_, err = c.AddFunc("17 15 * * 0-4", func() {
		SendMarketSummaryMessage(bot, chatID, client, isMarketOpen)
	})
	if err != nil {
		applog.Log(applog.ERROR, "Error scheduling market summary: %v", err)
		return
	}

	applog.Log(applog.INFO, "Market Summary Scheduled")
}

func SendMarketSummaryMessage(bot *tgbotapi.BotAPI, chatID int64, client *graphql.Client, ismarketopen bool) {

	marketSummary, err := dbgraphql.MarketSummary(client)
	if err != nil {
		applog.Log(applog.ERROR, "Error fetching market summary: %v", err)
		return
	}
	if ismarketopen {
		datetimeStr := marketSummary.NepseIndex.AsOfDate

		datetime, err := time.Parse("2006-01-02T15:04:05", datetimeStr)
		if err != nil {
			applog.Log(applog.ERROR, "Error converting datetime string to time.Time: %v", err)
		}

		er, err := nepalitime.FromEnglishTime(datetime)
		if err != nil {
			applog.Log(applog.ERROR, "Error parsing time: %v", err)
		}
		nep := er.String()[:10]
		opts := htmlcapture.CaptureOptions{
			Input: "templates/marketSummary.html",
			Variables: map[string]string{

				"Date":               services.BSDateConvert(nep),
				"IndexPoint":         fmt.Sprintf("%.2f", marketSummary.NepseIndex.IndexValue),
				"PointChange":        fmt.Sprintf("%.2f", marketSummary.NepseIndex.Difference),
				"PercentageChange":   fmt.Sprintf("%.2f%%", marketSummary.NepseIndex.PercentChange),
				"Turnover":           utils.NumberToCroreArab(marketSummary.NepseIndex.Turnover),
				"ShareTraded":        utils.NumberToCroreArab(float64(marketSummary.NepseIndex.Volume)),
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
			applog.Log(applog.ERROR, "Error capturing screenshot: %v", err)
		}

		responseText := `📊 NEPSE Daily Market Summary - ` + services.BSDateConvert(nep) + `

📈 NEPSE Index: ` + fmt.Sprintf("%.2f", marketSummary.NepseIndex.IndexValue) + ` (` + fmt.Sprintf("%.2f", marketSummary.NepseIndex.Difference) + ` | ` + fmt.Sprintf("%.2f%%", marketSummary.NepseIndex.PercentChange) + `)
💰 Total Turnover: Rs ` + utils.NumberToCroreArabFull(marketSummary.NepseIndex.Turnover) + `
📉 Total Traded Shares: ` + utils.NumberToCroreArabFull(float64(marketSummary.NepseIndex.Volume))

		photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "market_summary.png", Bytes: img})
		photo.Caption = responseText
		photo.ParseMode = "Markdown"

		if _, err := bot.Send(photo); err != nil {
			applog.Log(applog.ERROR, "Error sending market summary image: %v", err)
		}
	} else {
		applog.Log(applog.INFO, "Market Closed")
	}
}
