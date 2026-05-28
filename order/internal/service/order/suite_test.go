package order

import (
	"context"
	"testing"

	grpcMocks "github.com/ianagovitsyn/project/order/internal/client/grpc/mocks"
	repoMocks "github.com/ianagovitsyn/project/order/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type OrderServiceSuite struct {
	suite.Suite
	ctx             context.Context
	repo            *repoMocks.OrderRepository
	paymentClient   *grpcMocks.PaymentClient
	inventoryClient *grpcMocks.InventoryClient
	service         *Service
}

func (s *OrderServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = repoMocks.NewOrderRepository(s.T())
	s.paymentClient = grpcMocks.NewPaymentClient(s.T())
	s.inventoryClient = grpcMocks.NewInventoryClient(s.T())
	s.service = NewService(s.repo, s.paymentClient, s.inventoryClient)
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(OrderServiceSuite))
}
