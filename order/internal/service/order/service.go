package order

import (
	"github.com/ianagovitsyn/project/order/internal/client/grpc"
	"github.com/ianagovitsyn/project/order/internal/repository"
	def "github.com/ianagovitsyn/project/order/internal/service"
)

var _ def.OrderService = (*Service)(nil)

type Service struct {
	OrderRepository repository.OrderRepository
	PaymentClient   grpc.PaymentClient
	InventoryClient grpc.InventoryClient
}

func NewService(OrderRepository repository.OrderRepository, PaymentClient grpc.PaymentClient, InventoryClient grpc.InventoryClient) *Service {
	return &Service{
		OrderRepository: OrderRepository,
		PaymentClient:   PaymentClient,
		InventoryClient: InventoryClient,
	}
}
