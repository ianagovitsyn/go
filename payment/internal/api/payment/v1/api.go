package v1

import (
	"github.com/ianagovitsyn/project/payment/internal/service"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

type Api struct {
	paymentV1.UnimplementedPaymentServiceServer

	paymentService service.PaymentService
}

func NewAPI(paymentService service.PaymentService) *Api {
	return &Api{
		paymentService: paymentService,
	}
}
