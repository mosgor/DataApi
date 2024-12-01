package repositories

import (
	"ModelOrchestrator/pkg/internal/structs"
	"context"
)

type LogsRepository interface {
	Create(ctx context.Context, model *structs.LogsModel) error
	ReadAll(ctx context.Context) ([]structs.LogsModel, error)
}
