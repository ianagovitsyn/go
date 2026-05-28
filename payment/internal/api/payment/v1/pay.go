package v1

import (
	"context"

	"github.com/ianagovitsyn/project/payment/internal/converter"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

func (a *Api) PayOrder(ctx context.Context, r *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	transaction, _ := a.paymentService.Pay(ctx, converter.PaymentInfoToModel(r))

	return converter.TransactionToProto(transaction), nil
}
