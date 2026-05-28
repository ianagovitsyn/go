package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (r *Repository) Insert(_ context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[order.OrderUUID] = order

	return nil
}
