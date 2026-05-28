package repository

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/model"
)

type InventoryRepository interface {
	GetByUUID(ctx context.Context, uuid string) (model.Part, error)
	GetListByFilter(ctx context.Context, filter *model.Filter) ([]model.Part, error)
}
