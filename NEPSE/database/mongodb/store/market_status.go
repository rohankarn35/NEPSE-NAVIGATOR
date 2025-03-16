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

func MarketStatus(collection *mongo.Collection) {

	marketStatus, err := server.GetMarketStatus()
	if err != nil {
		applog.Log(applog.ERROR, "Error getting market status: %v", err)
		return
	}

	marketData := dbmodels.MarketStatus{
		IsOpen: marketStatus.IsOpen,
	}

	filter := bson.M{}

	_, err = collection.UpdateOne(
		context.TODO(),
		filter,
		bson.M{"$set": marketData},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		applog.Log(applog.ERROR, "Error updating market status in MongoDB: %v", err)
	}
}
