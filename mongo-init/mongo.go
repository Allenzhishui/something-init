package mongo

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func InitMongo(ctx context.Context) (*mongo.Client, error) {
	mongoCtx, cancel2 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel2()
	mongoClient, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatal().Str("mongodb", viper.GetString("mongo.uri")).Err(err).Msg("connect conn fail.")
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal().Str("mongodb", viper.GetString("mongo.uri")).Err(err).Msg("ping conn fail.")
		return nil, err
	}
	return mongoClient, nil

}
