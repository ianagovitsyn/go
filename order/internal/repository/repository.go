package repository

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

type OrderRepository interface {
	Get(ctx context.Context, uuid string) (model.Order, error)
	Update(ctx context.Context, order model.Order) error
	Insert(ctx context.Context, order model.Order) error
}
