package graph

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
	MongoClient *mongo.Database
}
