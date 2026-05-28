package converter

import (
	"github.com/ianagovitsyn/project/payment/internal/model"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

func PaymentInfoToModel(info *paymentV1.PayOrderRequest) model.PaymentInfo {
	return model.PaymentInfo{
		OrderUUID:     info.OrderUuid,
		UserUUID:      info.UserUuid,
		PaymentMethod: model.PaymentMethod(info.PaymentMethod),
	}
}

func TransactionToProto(transaction model.Transaction) *paymentV1.PayOrderResponse {
	return &paymentV1.PayOrderResponse{
		TransactionUuid: transaction.TransactionUUID,
	}
}
