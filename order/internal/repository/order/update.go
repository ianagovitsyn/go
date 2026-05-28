package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (r *Repository) Update(_ context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.storage[order.OrderUUID]
	if !ok {
		return model.ErrOrderNotFound
	}

	r.storage[order.OrderUUID] = order

	return nil
}
