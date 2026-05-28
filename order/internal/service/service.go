package service

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, params model.CreateOrderParams) (model.Order, error)
	OrderPay(ctx context.Context, paymentMethod model.PaymentMethod, OrderUUID string) (string, error)
	GetByUUID(ctx context.Context, OrderUUID string) (model.Order, error)
	CancelOrder(ctx context.Context, orderUUID string) error
}
