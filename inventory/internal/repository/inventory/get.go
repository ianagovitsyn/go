package inventory

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/model"
	"github.com/ianagovitsyn/project/inventory/internal/repository/converter"
)

func (r *Repository) GetByUUID(_ context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.storage[uuid]
	if !ok {
		return model.Part{}, model.ErrOrderNotFound
	}

	return converter.PartToModel(part), nil
}
