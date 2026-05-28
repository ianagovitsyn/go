package inventory

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/model"
)

func (s *Service) Get(ctx context.Context, uuid string) (model.Part, error) {
	part, err := s.inventoryRepository.GetByUUID(ctx, uuid)
	if err != nil {
		return model.Part{}, err
	}

	return part, nil
}
