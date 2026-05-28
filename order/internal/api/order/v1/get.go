package v1

import (
	"context"
	"errors"

	"github.com/ianagovitsyn/project/order/internal/converter"
	"github.com/ianagovitsyn/project/order/internal/model"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

func (a *Api) GetOrderByUuid(ctx context.Context, params orderV1.GetOrderByUuidParams) (orderV1.GetOrderByUuidRes, error) {
	order, err := a.orderService.GetByUUID(ctx, params.OrderUUID.String())
	if errors.Is(err, model.ErrOrderNotFound) {
		return &orderV1.NotFoundError{Code: 404, Message: "order not found"}, nil
	}

	return converter.ModelToGetOrderResponse(order), nil
}
