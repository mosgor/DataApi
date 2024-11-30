package repositories

import (
	"ModelOrchestrator/pkg/internal/structs"
	"context"
)

type ModelRepository interface {
	Create(ctx context.Context, model *structs.Resp) error
	ReadAll(ctx context.Context) ([]structs.Resp, error)
	ReadOne(ctx context.Context, id int) (structs.Resp, error)
	Update(ctx context.Context, model *structs.Resp) error
	Delete(ctx context.Context, id int) error
}
