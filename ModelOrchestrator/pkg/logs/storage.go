package logs

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/storageClient"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

type repository struct {
	client storageClient.PostgresClient
	log    *slog.Logger
}

func (r *repository) Create(ctx context.Context, model *structs.LogsModel) error {
	q := `
	INSERT INTO logs (source_id, model_id, time, time_with_response, status) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at; 
	`
	if err := r.client.QueryRow(
		ctx, q, model.SourceId,
		model.ModelId, model.Time,
		model.TimeWithResponse, model.Status,
	).Scan(&model.Id, &model.CreationDate); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			r.log.Error("Postgres error", *pgErr)
			return err
		}
		return err
	}
	r.log.Info("Created new log in Postgres")
	return nil
}

func (r *repository) ReadAll(ctx context.Context) ([]structs.LogsModel, error) {
	q := `
		SELECT * FROM logs;
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var res []structs.LogsModel
	for rows.Next() {
		var logModel structs.LogsModel
		err = rows.Scan(
			&logModel.Id, &logModel.SourceId,
			&logModel.ModelId, &logModel.Time, &logModel.TimeWithResponse,
			&logModel.Status, &logModel.CreationDate,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, logModel)
	}
	return res, nil
}

func NewRepository(p storageClient.PostgresClient, log *slog.Logger) repositories.LogsRepository {
	return &repository{
		client: p,
		log:    log,
	}
}
