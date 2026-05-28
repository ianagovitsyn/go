package v1

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ianagovitsyn/project/order/internal/model"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

func (a *Api) OrderPay(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.OrderPayParams) (orderV1.OrderPayRes, error) {
	transactionUUID, err := a.orderService.OrderPay(ctx, model.PaymentMethod(req.PaymentMethod), params.OrderUUID.String())
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderV1.NotFoundError{
				Code:    404,
				Message: "order not found",
			}, nil
		}

		return nil, err
	}

	return &orderV1.PayOrderResponse{
		TransactionUUID: uuid.MustParse(transactionUUID),
	}, nil
}
