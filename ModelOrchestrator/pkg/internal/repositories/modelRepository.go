package repositories

import (
	"ModelOrchestrator/pkg/internal/structs"
	"context"
)

type ModelRepository interface {
	Create(ctx context.Context, model *structs.Model) error
	ReadAll(ctx context.Context) ([]structs.Model, error)
	ReadOne(ctx context.Context, id int) (structs.Model, error)
	Update(ctx context.Context, model *structs.Model) error
	Delete(ctx context.Context, id int) error
}
