package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (s *Service) GetByUUID(ctx context.Context, orderUUID string) (model.Order, error) {
	order, err := s.OrderRepository.Get(ctx, orderUUID)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
