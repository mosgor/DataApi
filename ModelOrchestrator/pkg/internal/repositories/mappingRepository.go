package repositories

import (
	"ModelOrchestrator/pkg/internal/structs"
	"context"
)

type MappingRepository interface {
	Create(ctx context.Context, mapping *structs.MappingModel) error
	ReadAll(ctx context.Context) ([]structs.MappingModel, error)
	ReadOne(ctx context.Context, id string) (structs.MappingModel, error)
	Update(ctx context.Context, mapping *structs.MappingModel) error
	Delete(ctx context.Context, id string) error
}
