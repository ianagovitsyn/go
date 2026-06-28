package order

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/model"
)

func (s *Service) OrderPay(ctx context.Context, paymentMethod model.PaymentMethod, orderUUID string) (string, error) {
	order, err := s.OrderRepository.Get(ctx, orderUUID)
	if err != nil {
		return "", err
	}

	transactionUUID, err := s.PaymentClient.Pay(ctx, order, paymentMethod)
	if err != nil {
		return "", err
	}

	order.TransactionUUID = &transactionUUID
	order.Status = model.StatusPaid
	order.PaymentMethod = &paymentMethod

	err = s.OrderRepository.UpdatePayment(ctx, order)
	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}
