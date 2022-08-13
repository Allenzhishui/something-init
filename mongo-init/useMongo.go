package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func UserMongo(mongoClient *mongo.Client) {
	_ = mongoClient.Database("my").Collection("test")
}
