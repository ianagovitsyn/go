package grpc

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

type PaymentClient interface {
	Pay(ctx context.Context, order model.Order, paymentMethod model.PaymentMethod) (string, error)
}

type InventoryClient interface {
	ListParts(ctx context.Context, partUUIDs []string) ([]model.Part, error)
}
