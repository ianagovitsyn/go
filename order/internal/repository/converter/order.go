package converter

import (
	"github.com/ianagovitsyn/project/order/internal/model"
	repoModel "github.com/ianagovitsyn/project/order/internal/repository/model"
)

func OrderToModel(order repoModel.Order) model.Order {
	var pm *model.PaymentMethod
	if order.PaymentMethod != nil {
		v := model.PaymentMethod(*order.PaymentMethod)
		pm = &v
	}

	return model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   pm,
		Status:          model.Status(order.Status),
	}
}

func OrderToRepo(order model.Order) repoModel.Order {
	var pm *string
	if order.PaymentMethod != nil {
		v := string(*order.PaymentMethod)
		pm = &v
	}
	return repoModel.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   pm,
		Status:          string(order.Status),
	}
}
