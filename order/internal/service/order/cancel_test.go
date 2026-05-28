package order

import (
	"github.com/ianagovitsyn/project/order/internal/model"
)

func (s *OrderServiceSuite) TestCancelOrder_Success() {
	orderUUID := "123"
	order := model.Order{
		OrderUUID: orderUUID,
		Status:    "PENDING_PAYMENT",
	}

	s.repo.EXPECT().
		Get(s.ctx, orderUUID).
		Return(order, nil)

	cancelledOrder := order
	cancelledOrder.Status = "CANCELLED"
	s.repo.EXPECT().
		Update(s.ctx, cancelledOrder).
		Return(nil)

	err := s.service.CancelOrder(s.ctx, orderUUID)
	s.NoError(err)
}

func (s *OrderServiceSuite) TestCancelOrder_NotFound() {
	s.repo.EXPECT().
		Get(s.ctx, "456").
		Return(model.Order{}, model.ErrOrderNotFound)

	err := s.service.CancelOrder(s.ctx, "456")
	s.ErrorIs(err, model.ErrOrderNotFound)
}

func (s *OrderServiceSuite) TestCancelOrder_WrongStatus() {
	s.repo.EXPECT().
		Get(s.ctx, "789").
		Return(model.Order{
			OrderUUID: "789",
			Status:    "PAID",
		}, nil)

	err := s.service.CancelOrder(s.ctx, "789")
	s.ErrorIs(err, model.ErrConflict)
}
