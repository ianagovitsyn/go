package repository

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

type OrderRepository interface {
	Get(ctx context.Context, uuid string) (model.Order, error)
	Insert(ctx context.Context, order model.Order) error
	UpdatePayment(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, order model.Order) error
}
