package payment

import def "github.com/ianagovitsyn/project/payment/internal/service"

var _ def.PaymentService = (*Service)(nil)

type Service struct{}

func NewService() *Service {
	return &Service{}
}
