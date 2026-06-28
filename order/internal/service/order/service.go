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

func NewService(orderRepository repository.OrderRepository, paymentClient grpc.PaymentClient, inventoryClient grpc.InventoryClient) *Service {
	return &Service{
		OrderRepository: orderRepository,
		PaymentClient:   paymentClient,
		InventoryClient: inventoryClient,
	}
}
