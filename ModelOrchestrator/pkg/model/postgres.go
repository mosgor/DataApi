package model

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/storageClients"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"log/slog"
)

type repository struct {
	client storageClients.PostgresClient
	log    *slog.Logger
}

func (r *repository) Create(ctx context.Context, model *structs.Model) error {
	return nil
}

func (r *repository) ReadAll(ctx context.Context) ([]structs.Model, error) {
	return nil, nil
}

func (r *repository) ReadOne(ctx context.Context, id int) (structs.Model, error) {
	return structs.Model{}, nil
}

func (r *repository) Update(ctx context.Context, model *structs.Model) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}

func NewRepository(client storageClients.PostgresClient, log *slog.Logger) repositories.ModelRepository {
	return &repository{
		client: client,
		log:    log,
	}
}
