package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryV1API "github.com/ianagovitsyn/project/inventory/internal/api/inventory/v1"
	inventoryRepository "github.com/ianagovitsyn/project/inventory/internal/repository/inventory"
	inventoryService "github.com/ianagovitsyn/project/inventory/internal/service/inventory"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	// Создаем слушатель на порт
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	// Проверяем на ошибку
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	// defer на закрытие слушателя
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	s := grpc.NewServer()

	repository := inventoryRepository.NewRepository()
	service := inventoryService.NewService(repository)
	api := inventoryV1API.NewAPI(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
