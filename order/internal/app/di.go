package app

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderV1API "github.com/ianagovitsyn/project/order/internal/api/order/v1"
	grpcClient "github.com/ianagovitsyn/project/order/internal/client/grpc"
	inventoryClient "github.com/ianagovitsyn/project/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/ianagovitsyn/project/order/internal/client/grpc/payment/v1"
	"github.com/ianagovitsyn/project/order/internal/config"
	"github.com/ianagovitsyn/project/order/internal/repository"
	orderRepository "github.com/ianagovitsyn/project/order/internal/repository/order"
	"github.com/ianagovitsyn/project/order/internal/service"
	orderService "github.com/ianagovitsyn/project/order/internal/service/order"
	"github.com/ianagovitsyn/project/platform/pkg/closer"
	"github.com/ianagovitsyn/project/platform/pkg/logger"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
)

type diContainer struct {
	orderV1API      orderV1.Handler
	orderService    service.OrderService
	orderRepository repository.OrderRepository
	inventoryClient grpcClient.InventoryClient
	paymentClient   grpcClient.PaymentClient
	pgConn          *pgx.Conn
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1API(ctx context.Context) orderV1.Handler {
	if d.orderV1API == nil {
		d.orderV1API = orderV1API.NewAPI(d.OrderService(ctx))
	}

	return d.orderV1API
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewService(d.OrderRepository(ctx), d.PaymentClient(ctx), d.InventoryClient(ctx))
	}

	return d.orderService
}

func (d *diContainer) InventoryClient(ctx context.Context) grpcClient.InventoryClient {
	if d.inventoryClient == nil {
		conn, err := grpc.NewClient(
			config.AppConfig().InventoryGRPC.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			logger.Error(ctx, "failed to connect", zap.Error(err))
			return nil
		}
		closer.AddNamed("inventory gRPC conn", func(ctx context.Context) error {
			return conn.Close()
		})

		generatedClient := inventoryV1.NewInventoryServiceClient(conn)

		d.inventoryClient = inventoryClient.NewClient(generatedClient)
	}

	return d.inventoryClient
}

func (d *diContainer) PaymentClient(ctx context.Context) grpcClient.PaymentClient {
	if d.paymentClient == nil {
		conn, err := grpc.NewClient(
			config.AppConfig().PaymentGRPC.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			logger.Error(ctx, "failed to connect", zap.Error(err))
			return nil
		}
		closer.AddNamed("payment gRPC conn", func(ctx context.Context) error {
			return conn.Close()
		})

		generatedClient := paymentV1.NewPaymentServiceClient(conn)

		d.paymentClient = paymentClient.NewClient(generatedClient)
	}

	return d.paymentClient
}

func (d *diContainer) OrderRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		d.orderRepository = orderRepository.NewRepository(d.PgConn(ctx))
	}

	return d.orderRepository
}

func (d *diContainer) PgConn(ctx context.Context) *pgx.Conn {
	if d.pgConn == nil {
		conn, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			logger.Error(ctx, "failed to connect", zap.Error(err))
			return nil
		}
		closer.AddNamed("postgres conn", func(ctx context.Context) error {
			return conn.Close(ctx)
		})

		err = conn.Ping(ctx)
		if err != nil {
			logger.Error(ctx, "failed to ping", zap.Error(err))
			return nil
		}

		d.pgConn = conn
	}

	return d.pgConn
}
