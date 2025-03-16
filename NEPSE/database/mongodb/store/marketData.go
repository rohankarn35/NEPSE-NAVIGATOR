package store

import (
	"context"
	"fmt"
	dbmodels "nepseserver/database/models"
	applog "nepseserver/log"
	"nepseserver/server"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StoreOrUpdateMarketData(collection *mongo.Collection) error {
	// Fetch stock market data
	stocks, err := server.MarketData()
	if err != nil {
		applog.Log(applog.ERROR, "error fetching market data: %v", err)
		return fmt.Errorf("error fetching market data: %v", err)
	}

	// Create a slice of bulk write operations
	var bulkOps []mongo.WriteModel

	for _, stock := range stocks {
		stock := dbmodels.Market{
			Symbol:           stock.StockSymbol,
			Company:          stock.CompanyName,
			TradeVolume:      int32(stock.NoOfTransactions),
			High:             stock.MaxPrice,
			Low:              stock.MinPrice,
			Open:             stock.OpeningPrice,
			Close:            stock.ClosingPrice,
			TotalTradedValue: stock.Amount,
			PrevClose:        stock.PreviousClosing,
			PriceChange:      stock.DifferenceRs,
			PercentChange:    stock.PercentChange,
			ShareVolume:      int32(stock.Volume),
			LastUpdated:      strings.Replace(stock.AsOfDateString, "As of ", "", 1),
		}
		filter := bson.M{"symbol": stock.Symbol} // Find by stock symbol
		update := bson.M{"$set": stock}          // Update existing fields
		upsert := true                           // Insert if not found

		// Define the update operation
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(upsert)

		// Add to bulk operations
		bulkOps = append(bulkOps, model)
	}

	// Execute bulk write operation
	if len(bulkOps) > 0 {
		_, err := collection.BulkWrite(context.TODO(), bulkOps)
		if err != nil {
			applog.Log(applog.ERROR, "error performing bulk write: %v", err)
			return fmt.Errorf("error performing bulk write: %v", err)
		}
		applog.Log(applog.INFO, "Market data updated successfully.")
	} else {
		applog.Log(applog.INFO, "No stock data to update.")
	}

	return nil
}
