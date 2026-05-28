package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	orderAPI "github.com/ianagovitsyn/project/order/internal/api/order/v1"
	inventoryClient "github.com/ianagovitsyn/project/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/ianagovitsyn/project/order/internal/client/grpc/payment/v1"
	orderRepository "github.com/ianagovitsyn/project/order/internal/repository/order"
	orderService "github.com/ianagovitsyn/project/order/internal/service/order"
	orderV1 "github.com/ianagovitsyn/project/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/ianagovitsyn/project/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = "8080"
	// Таймауты для HTTP-сервера
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second

	serverAddr    = "localhost"
	inventoryPort = "50051"
	paymentPort   = "50052"
)

func main() {
	//inventoryConn := createConn(inventoryPort)
	//paymentConn := createConn(paymentPort)
	//
	//service := &OrderService{
	//	orders:          make(map[string]*orderV1.OrderDto),
	//	inventoryClient: inventoryV1.NewInventoryServiceClient(inventoryConn),
	//	paymentClient:   paymentV1.NewPaymentServiceClient(paymentConn),
	//}

	paymentConn := createConn(paymentPort)
	inventoryConn := createConn(inventoryPort)

	paymentGenerated := paymentV1.NewPaymentServiceClient(paymentConn)
	inventoryGenerated := inventoryV1.NewInventoryServiceClient(inventoryConn)

	pClient := paymentClient.NewClient(paymentGenerated)
	iClient := inventoryClient.NewClient(inventoryGenerated)
	repository := orderRepository.NewRepository()

	service := orderService.NewService(repository, pClient, iClient)

	api := orderAPI.NewAPI(service)

	// Создаем OpenAPI сервер
	orderServer, err := orderV1.NewServer(api)
	if err != nil {
		log.Fatalf("ошибка создания сервера OpenAPI: %v", err)
	}

	// Инициализируем роутер Chi
	r := chi.NewRouter()

	r.Mount("/", orderServer)

	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout, // Защита от Slowloris атак - тип DDoS-атаки, при которой
		// атакующий умышленно медленно отправляет HTTP-заголовки, удерживая соединения открытыми и истощая
		// пул доступных соединений на сервере. ReadHeaderTimeout принудительно закрывает соединение,
		// если клиент не успел отправить все заголовки за отведенное время.
	}

	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}

func createConn(port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(
		serverAddr+":"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return nil
	}

	return conn
}
