package v1

import (
	"context"

	"go.uber.org/zap"

	"github.com/ianagovitsyn/project/payment/internal/converter"
	"github.com/ianagovitsyn/project/platform/pkg/logger"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

func (a *Api) PayOrder(ctx context.Context, r *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	transaction, err := a.paymentService.Pay(ctx, converter.PaymentInfoToModel(r))
	if err != nil {
		logger.Error(ctx, "failed", zap.Error(err))
		return nil, err
	}

	return converter.TransactionToProto(transaction), nil
}
