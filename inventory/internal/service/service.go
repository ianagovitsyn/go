package service

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/model"
)

type InventoryService interface {
	Get(ctx context.Context, uuid string) (model.Part, error)
	List(ctx context.Context, filter *model.Filter) ([]model.Part, error)
}
