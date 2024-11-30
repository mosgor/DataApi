package model

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/storageClient"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
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
	).Scan(&model.SourceID, &model.CreationDate); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			r.log.Error("Postgres error", *pgErr)
			return err
		}
		return err
	}
	model.DateCreated = model.CreationDate
	_, err := r.mongoClient.InsertOne(ctx, model.MongoModel)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ReadAll(ctx context.Context) ([]structs.Resp, error) {
	return nil, nil
}

func (r *repository) ReadOne(ctx context.Context, id int) (structs.Resp, error) {
	return structs.Resp{}, nil
}

func (r *repository) Update(ctx context.Context, model *structs.Resp) error {
	return nil
}

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
