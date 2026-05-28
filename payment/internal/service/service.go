package service

import (
	"context"

	"github.com/ianagovitsyn/project/payment/internal/model"
)

type PaymentService interface {
	Pay(ctx context.Context, info model.PaymentInfo) (model.Transaction, error)
}
