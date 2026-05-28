package inventory

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/model"
)

func (s *Service) List(ctx context.Context, filter *model.Filter) ([]model.Part, error) {
	parts, err := s.inventoryRepository.GetListByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return parts, nil
}
