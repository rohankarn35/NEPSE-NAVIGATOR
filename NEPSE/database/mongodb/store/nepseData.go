package store

import (
	"context"

	applog "nepseserver/log"
	"nepseserver/server"

	models "nepseserver/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StoreNepseData(collection *mongo.Collection) {
	nepseData, err := server.FetchNepseData()
	if err != nil {
		applog.Log(applog.ERROR, "Failed to fetch nepse data")
		return
	}

	data := models.NepseIndex{
		MarketIndex:          nepseData.IndexName,
		CurrentValue:         nepseData.IndexValue,
		PreviousClose:        nepseData.PreviousValue,
		OpeningValue:         nepseData.OpeningValue,
		PercentageChange:     nepseData.PercentChange,
		PointChange:          nepseData.Difference,
		TotalTurnover:        nepseData.Turnover,
		TradedVolume:         int32(nepseData.Volume),
		MarketCapitalization: nepseData.MarketCap,
		DailyHigh:            nepseData.DayHigh,
		DailyLow:             nepseData.DayLow,
		YearlyHigh:           nepseData.YearHigh,
		YearlyLow:            nepseData.YearLow,
		Date:                 nepseData.AsOfDate,
	}
	filter := bson.M{"index_name": data.MarketIndex}
	update := bson.M{"$set": data}
	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to upsert nepse data: %v", err)
		return
	}
	applog.Log(applog.INFO, "Nepse Data updated")
}
