package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (s *Service) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.OrderRepository.Get(ctx, orderUUID)
	if err != nil {
		return err
	}

	if order.Status == "PENDING_PAYMENT" {
		order.Status = "CANCELLED"
		err := s.OrderRepository.UpdateStatus(ctx, order)
		if err != nil {
			return err
		}
		return nil
	}

	return model.ErrConflict
}
