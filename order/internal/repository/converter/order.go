package converter

import (
	"github.com/ianagovitsyn/project/order/internal/model"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"
)

func OrderToModel(order repoModel.Order) model.Order {
	return model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   model.PaymentMethod(order.PaymentMethod),
		Status:          model.Status(order.Status),
	}
}

func OrderToRepo(order model.Order) repoModel.Order {
	return repoModel.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   string(order.PaymentMethod),
		Status:          string(order.Status),
	}
}
