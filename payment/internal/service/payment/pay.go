package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/ianagovitsyn/project/payment/internal/model"
)

func (s *Service) Pay(_ context.Context, _ model.PaymentInfo) (model.Transaction, error) {
	orderUUID := uuid.NewString()
	idempotencyKey := uuid.NewString()

	return model.Transaction{
		TransactionUUID: orderUUID,
		IdempotencyKey:  idempotencyKey,
	}, nil
}
