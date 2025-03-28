package store

import (
	"context"
	dbmodels "nepseserver/database/models"
	applog "nepseserver/log"
	"nepseserver/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MarketMovers(collection *mongo.Collection) {
	loserData, err := server.GetMarketMovers("losers")
	if err != nil {
		applog.Log(applog.ERROR, "Failed to get IPO data: %v", err)
		return
	}

	gainerData, err := server.GetMarketMovers("gainers")
	if err != nil {
		applog.Log(applog.ERROR, "Failed to get FPO data: %v", err)
		return
	}
	var loserdata []dbmodels.MarketMover
	for _, ipo := range loserData {
		loserdata = append(loserdata, dbmodels.MarketMover{
			StockCode:         ipo.StockSymbol,
			Company:           ipo.CompanyName,
			TransactionsCount: int32(ipo.NoOfTransactions),
			HighestPrice:      ipo.MaxPrice,
			LowestPrice:       ipo.MinPrice,
			OpeningPrice:      ipo.OpeningPrice,
			ClosingPrice:      ipo.ClosingPrice,
			Turnover:          ipo.Amount,
			PreviousClose:     ipo.PreviousClosing,
			PriceChange:       ipo.DifferenceRs,
			PercentageChange:  ipo.PercentChange,
			TradedVolume:      int32(ipo.Volume),
			TradeDate:         ipo.TradeDate,
		})
	}

	var gainerdata []dbmodels.MarketMover
	for _, gainer := range gainerData {
		gainerdata = append(gainerdata,

			dbmodels.MarketMover{
				StockCode:         gainer.StockSymbol,
				Company:           gainer.CompanyName,
				TransactionsCount: int32(gainer.NoOfTransactions),
				HighestPrice:      gainer.MaxPrice,
				LowestPrice:       gainer.MinPrice,
				OpeningPrice:      gainer.OpeningPrice,
				ClosingPrice:      gainer.ClosingPrice,
				Turnover:          gainer.Amount,
				PreviousClose:     gainer.PreviousClosing,
				PriceChange:       gainer.DifferenceRs,
				PercentageChange:  gainer.PercentChange,
				TradedVolume:      int32(gainer.Volume),
				TradeDate:         gainer.TradeDate,
			})

	}

	totalData := dbmodels.MarketMovers{
		Gainers: gainerdata,
		Loser:   loserdata,
	}

	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		applog.Log(applog.ERROR, "Failed to count documents: %v", err)
		return
	}

	if count == 0 {
		_, err := collection.InsertOne(context.TODO(), totalData)
		if err != nil {
			applog.Log(applog.ERROR, "Failed to insert data: %v", err)
		}
	} else {
		_, err := collection.ReplaceOne(context.TODO(), bson.M{}, totalData)
		if err != nil {
			applog.Log(applog.ERROR, "Failed to update data: %v", err)
		}
	}
	applog.Log(applog.INFO, "Market movers successfully updated")
}
