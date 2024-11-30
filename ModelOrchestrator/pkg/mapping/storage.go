package mapping

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type repository struct {
	mongoClient *mongo.Collection
	log         *slog.Logger
}

func (r *repository) Create(ctx context.Context, mapping *structs.MappingModel) error {
	res, err := r.mongoClient.InsertOne(ctx, mapping)
	mapping.Id = res.InsertedID.(primitive.ObjectID).Hex()
	if err != nil {
		return fmt.Errorf("%w error while inserting in Mongo", err)
	}
	r.log.Info("Inserted new mapping in Mongo")
	return nil
}

func (r *repository) ReadAll(ctx context.Context) ([]structs.MappingModel, error) {
	var resp []structs.MappingModel
	curs, err := r.mongoClient.Find(ctx, bson.D{})
	if err != nil {
		r.log.Error("Error while reading mappings from Mongo ", err)
		return nil, err
	}
	for curs.Next(ctx) {
		var temp structs.MappingModel
		err := curs.Decode(&temp)
		if err != nil {
			r.log.Error("Error while decoding from mongo", err)
			return nil, err
		}
		resp = append(resp, temp)
	}
	r.log.Info("Read all mappings from Mongo ")
	return resp, nil
}

func (r *repository) ReadOne(ctx context.Context, id string) (structs.MappingModel, error) {
	var resp structs.MappingModel
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.log.Error("Error while converting objectId to primitive.ObjectID", err)
		return structs.MappingModel{}, err
	}
	err = r.mongoClient.FindOne(ctx, bson.M{"_id": objId}).Decode(&resp)
	if err != nil {
		r.log.Error("Error while decoding from mongo", err)
		return structs.MappingModel{}, err
	}
	r.log.Info("Read mapping from Mongo")
	return resp, nil
}

// Update TODO: implement
func (r *repository) Update(ctx context.Context, mapping *structs.MappingModel) error {
	return nil
}

// Delete TODO: implement
func (r *repository) Delete(ctx context.Context, id string) error {
	return nil
}

func NewRepository(mClient *mongo.Client, log *slog.Logger) repositories.MappingRepository {
	return &repository{
		mongoClient: mClient.Database("DataApi").Collection("mappings"),
		log:         log,
	}
}
