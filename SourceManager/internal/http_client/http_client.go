package http_client

import (
	"SourceManager/internal/config"
	"SourceManager/internal/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateData(cfg config.Config) [][]int32 {

	log := logger.SetupLogger(cfg.Env)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.Timeout)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://mongo:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Error("FROM http_client/handlers/get", "unable to connect to mongo", err.Error())
		return [][]int32{}
	}

	coll := client.Database("DataApi").Collection("mappings")

	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", "$source_id"}}}},
	}
	cursor, err := coll.Aggregate(ctx, pipeline)

	indexes := make([][]int32, 0)

	type OneIndex struct {
		SourceId []int32 `bson:"_id"`
	}

	var temp OneIndex

	if err != nil {
		log.Error(err.Error())
		return [][]int32{}
	}
	for cursor.Next(ctx) {
		cursor.Decode(&temp)
		indexes = append(indexes, temp.SourceId)
	}
	log.Info("Sources read successfully")
	return indexes

}
