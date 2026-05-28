package v1

import (
	"context"
	"errors"

	"github.com/ianagovitsyn/project/order/internal/model"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

func (a *Api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	err := a.orderService.CancelOrder(ctx, params.OrderUUID.String())
	if errors.Is(err, model.ErrOrderNotFound) {
		return &orderV1.NotFoundError{Code: 404, Message: "order not found"}, nil
	}
	if errors.Is(err, model.ErrConflict) {
		return &orderV1.ConflictError{}, nil
	}

	return &orderV1.CancelOrderNoContent{}, nil
}
