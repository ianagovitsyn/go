package v1

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/converter"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

func (a *Api) CreateNewOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateNewOrderRes, error) {
	order, err := a.orderService.CreateNewOrder(ctx, converter.OrderRequestToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.ModelToNewOrderResponse(order), nil
}
