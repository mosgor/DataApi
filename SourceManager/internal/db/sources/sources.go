package sources

import (
	"SourceManager/internal/db"
	"SourceManager/internal/repositories"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	postgresClient *pgxpool.Pool
	mongoClient    *mongo.Collection
	log            *slog.Logger
}

func (r *repository) Create(ctx context.Context, source *db.Source) error {
	q := `
    INSERT INTO data_sources (name, connection_string)
    VALUES ($1, $2)
    RETURNING id, created_at`

	if err := r.postgresClient.QueryRow(ctx, q, source.Name, source.URL).Scan(&source.SourceID, &source.CreatedAt); err != nil {
		var pgEr *pgconn.PgError
		if errors.As(err, &pgEr) {
			r.log.Error(pgEr.Error())
			return err
		}
		return err
	}

	r.log.Info("Successfully connected to PostgreSQL")
	r.log.Info("Created new source in PostgreSQL")

	_, err := r.mongoClient.InsertOne(ctx, source.MongoData)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	r.log.Info("Successfully connected to MongoDB")
	r.log.Info("Created new source in MongoDB")

	return nil
}

func (r *repository) ReadOne(ctx context.Context, ID int) (db.Source, error) {
	q := `SELECT * FROM data_sources WHERE id = $1`
	source := db.Source{}
	err := r.postgresClient.QueryRow(ctx, q, ID).Scan(&source.PostgresID, &source.Name, &source.URL, &source.CreatedAt)
	if err != nil {
		return db.Source{}, err
	}

	filter := bson.M{"source_id": source.PostgresID}
	err = r.mongoClient.FindOne(ctx, filter).Decode(&source.MongoData)
	if err != nil {
		return source, err
	}

	r.log.Info(fmt.Sprintf("Data from data source with id %d proceeded successfully", ID))
	return source, nil
}

func (r *repository) ReadAll(ctx context.Context) ([]db.Source, error) {
	q := `SELECT * FROM data_sources`

	rows, err := r.postgresClient.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sources := []db.Source{}

	for rows.Next() {
		source := db.Source{}
		err := rows.Scan(&source.PostgresID, &source.Name, &source.URL, &source.CreatedAt)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"source_id": source.PostgresID}
		err = r.mongoClient.FindOne(ctx, filter).Decode(&source.MongoData)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)

	}
	r.log.Info("Data from 'source_schemas read successfully")

	return sources, nil

}

func NewRepository(pool *pgxpool.Pool, mongo *mongo.Client, log *slog.Logger) repositories.SourceRepository {
	return &repository{
		pool, mongo.Database("DataApi").Collection("source_schemas"), log,
	}
}
