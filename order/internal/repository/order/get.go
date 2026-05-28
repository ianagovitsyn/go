package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (r *Repository) Get(_ context.Context, orderUUID string) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.storage[orderUUID]
	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}

	return order, nil
}
