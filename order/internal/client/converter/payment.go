package converter

import (
	"github.com/ianagovitsyn/project/order/internal/model"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

func PayToGrpc(order model.Order, paymentMethod model.PaymentMethod) *paymentV1.PayOrderRequest {
	return &paymentV1.PayOrderRequest{
		OrderUuid:     order.OrderUUID,
		UserUuid:      order.UserUUID,
		PaymentMethod: paymentMethodToProto(paymentMethod),
	}
}

func paymentMethodToProto(m model.PaymentMethod) paymentV1.PaymentMethod {
	switch m {
	case model.PaymentMethodCard:
		return paymentV1.PaymentMethod_CARD
	case model.PaymentMethodSBP:
		return paymentV1.PaymentMethod_SBP
	case model.PaymentMethodCreditCard:
		return paymentV1.PaymentMethod_CREDIT_CARD
	case model.PaymentMethodInvestorMoney:
		return paymentV1.PaymentMethod_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_UNKNOWN
	}
}
