package v1

import (
	"github.com/ianagovitsyn/project/order/internal/service"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

type Api struct {
	orderV1.UnimplementedHandler

	orderService service.OrderService
}

func NewAPI(orderService service.OrderService) *Api {
	return &Api{
		orderService: orderService,
	}
}
