package store

import (
	"context"
	"log"
	dbmodels "nepseserver/database/models"
	"nepseserver/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StoreNepseData(collection *mongo.Collection) {
	nepseData, err := server.FetchNepseData()
	if err != nil {
		log.Print("Failed to fetch nepse data")
	}

	data := dbmodels.NepseIndex{
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
		log.Printf("Failed to upsert nepse data: %v", err)
	}
	log.Print("Nepse Data updated")

}
