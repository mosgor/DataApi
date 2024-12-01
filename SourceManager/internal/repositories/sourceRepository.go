package repositories

import (
	"SourceManager/internal/db"
	"context"
)

type SourceRepository interface {
	Create(ctx context.Context, source *db.Source) error
	ReadOne(ctx context.Context, ID int) (db.Source, error)
	ReadAll(ctx context.Context) ([]db.Source, error)
}
