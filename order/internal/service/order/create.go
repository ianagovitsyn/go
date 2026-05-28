package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/ianagovitsyn/project/order/internal/model"
)

func (s *Service) CreateNewOrder(ctx context.Context, params model.CreateOrderParams) (model.Order, error) {
	parts, err := s.InventoryClient.ListParts(ctx, params.PartUUIDs)
	if err != nil {
		return model.Order{}, err
	}

	if len(params.PartUUIDs) != len(parts) {
		return model.Order{}, model.ErrPartsNotFound
	}

	var totalPrice float64
	for _, v := range parts {
		totalPrice += v.Price
	}

	orderUuid := uuid.NewString()

	order := model.Order{
		OrderUUID:       orderUuid,
		UserUUID:        params.UserUUID,
		PartUuids:       params.PartUUIDs,
		TotalPrice:      totalPrice,
		TransactionUUID: "",
		PaymentMethod:   "",
		Status:          model.StatusPendingPayment,
	}

	err = s.OrderRepository.Insert(ctx, order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
