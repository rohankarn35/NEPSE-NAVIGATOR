package store

import (
	"context"
	dbmodels "nepseserver/database/models"
	applog "nepseserver/log"
	"nepseserver/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StoreIndicesData(collection *mongo.Collection) {

	indices, err := server.GetIndices()
	if err != nil {
		applog.Log(applog.ERROR, "Failed to get indices: %v", err)
		return
	}

	for _, index := range indices {
		index := dbmodels.Indices{
			IndexName:        index.IndexName,
			IndexValue:       index.IndexValue,
			PreviousValue:    index.PreviousValue,
			OpeningValue:     index.OpeningValue,
			PercentChange:    index.PercentChange,
			Difference:       index.Difference,
			Turnover:         index.Turnover,
			Volume:           int32(index.Volume),
			TotalCompanies:   int32(index.NoOfListedCompanies),
			TradedCompanies:  int32(index.NoOfTradedCompanies),
			Transactions:     int32(index.NoOfTransactions),
			ListedShares:     int64(index.NoOfListedShares),
			MarketCap:        index.MarketCap,
			DailyHigh:        index.DayHigh,
			DailyLow:         index.DayLow,
			YearlyHigh:       index.YearHigh,
			YearlyLow:        index.YearLow,
			ReportDate:       index.AsOfDate,
			ReportDateString: index.AsOfDateString,
			GainingCompanies: int32(index.NoOfGainers),
			LosingCompanies:  int32(index.NoOfLosers),
			Unchanged:        int32(index.NoOfUnchanged),
		}
		filter := bson.M{"index_name": index.IndexName}
		update := bson.M{
			"$set": index,
		}
		opts := options.Update().SetUpsert(true)
		_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			applog.Log(applog.ERROR, "Failed to update index: %v", err)
			return
		}
		applog.Log(applog.INFO, "Indices data updated successfully")
	}

}
