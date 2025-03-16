package mongodb

import (
	"context"
	"log"
	"nepseserver/constants"
	applog "nepseserver/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Client {
	url := constants.DB_URL

	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	applog.Log(applog.INFO, "Connected to Database/n")

	return client

}
