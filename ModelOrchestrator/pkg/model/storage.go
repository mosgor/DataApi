package model

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/storageClient"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type repository struct {
	postgresClient storageClient.PostgresClient // *pgxpool.Pool
	mongoClient    *mongo.Collection
	log            *slog.Logger
}

func (r *repository) Create(ctx context.Context, model *structs.Resp) error {
	q := `
		INSERT INTO ml_models (name, connection_string)
		VALUES ($1, $2)
		RETURNING id, created_at`
	if err := r.postgresClient.QueryRow(
		ctx, q, model.Name,
		model.Url,
	).Scan(&model.ModelID, &model.DateCreated); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			r.log.Error("Postgres error", *pgErr)
			return err
		}
		return err
	}
	r.log.Info("Created new model in Postgres")
	_, err := r.mongoClient.InsertOne(ctx, model.MongoModel)
	if err != nil {
		return fmt.Errorf("%w error while inserting in Mongo", err)
	}
	r.log.Info("Inserted new model in Mongo")
	return nil
}

func (r *repository) ReadAll(ctx context.Context) ([]structs.Resp, error) {
	q := `
		SELECT * FROM ml_models;
	`
	rows, err := r.postgresClient.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var res []structs.Resp
	for rows.Next() {
		var respModel structs.Resp
		err = rows.Scan(
			&respModel.Id, &respModel.Name,
			&respModel.Url, &respModel.DateCreated,
		)
		if err != nil {
			return nil, err
		}
		err = r.mongoClient.FindOne(ctx, bson.M{"model_id": respModel.Id}).Decode(&respModel.MongoModel)
		if err != nil {
			return nil, err
		}
		res = append(res, respModel)
	}
	return res, nil
}

func (r *repository) ReadOne(ctx context.Context, id int) (structs.Resp, error) {
	q := `
	SELECT * FROM ml_models WHERE id = $1;
	`
	var res structs.Resp
	rw := r.postgresClient.QueryRow(ctx, q, id)
	if err := rw.Scan(
		&res.Id, &res.Name,
		&res.Url, &res.DateCreated,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			r.log.Error("Data base error", *pgErr)
			return structs.Resp{}, err
		}
		return structs.Resp{}, err
	}
	err := r.mongoClient.FindOne(ctx, bson.M{"model_id": res.Id}).Decode(&res.MongoModel)
	if err != nil {
		return structs.Resp{}, err
	}
	return res, nil
}

// Update TODO: implement update
func (r *repository) Update(ctx context.Context, model *structs.Resp) error {
	return nil
}

// Delete TODO: implement delete
func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}

func NewRepository(pClient storageClient.PostgresClient, mClient *mongo.Client, log *slog.Logger) repositories.ModelRepository {
	return &repository{
		postgresClient: pClient,
		mongoClient:    mClient.Database("DataApi").Collection("model_schemas"),
		log:            log,
	}
}
