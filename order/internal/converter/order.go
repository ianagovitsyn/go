package converter

import (
	"github.com/google/uuid"

	"github.com/ianagovitsyn/project/order/internal/model"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
)

func OrderRequestToModel(r *orderV1.CreateOrderRequest) model.CreateOrderParams {
	uuids := make([]string, 0, len(r.PartUuids))
	for _, v := range r.GetPartUuids() {
		uuids = append(uuids, v.String())
	}

	return model.CreateOrderParams{
		UserUUID:  r.UserUUID.String(),
		PartUUIDs: uuids,
	}
}

func ModelToNewOrderResponse(order model.Order) *orderV1.CreateOrderResponse {
	return &orderV1.CreateOrderResponse{
		UUID:       uuid.MustParse(order.OrderUUID),
		TotalPrice: order.TotalPrice,
	}
}

func ModelToGetOrderResponse(order model.Order) *orderV1.OrderDto {
	partUUIDs := make([]uuid.UUID, 0, len(order.PartUuids))
	for _, v := range order.PartUuids {
		partUUIDs = append(partUUIDs, uuid.MustParse(v))
	}

	dto := &orderV1.OrderDto{
		OrderUUID:  uuid.MustParse(order.OrderUUID),
		UserUUID:   uuid.MustParse(order.UserUUID),
		PartUuids:  partUUIDs,
		TotalPrice: order.TotalPrice,
		Status:     orderV1.OrderStatus(order.Status),
	}

	if order.TransactionUUID != nil {
		dto.TransactionUUID = orderV1.OptUUID{
			Value: uuid.MustParse(*order.TransactionUUID),
			Set:   true,
		}
	}

	if order.PaymentMethod != nil {
		dto.PaymentMethod = orderV1.OptPaymentMethod{
			Value: orderV1.PaymentMethod(*order.PaymentMethod),
			Set:   true,
		}
	}

	return dto
}
