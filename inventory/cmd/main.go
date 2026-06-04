package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Не удалось загрузить файл .env: %v\n", err)
		return
	}

	dbURI := os.Getenv("MONGO_URI")
	if dbURI == "" {
		log.Println("Ошибка: переменная окружения MONGO_URI не установлена")
		return
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("Ошибка подключения к MongoDB: %v\n", err)
		return
	}

	defer func() {
		if cerr := client.Disconnect(ctx); cerr != nil {
			log.Printf("Ошибка при отключении от MongoDB: %v\n", cerr)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("MongoDB недоступна, ошибка ping: %v\n", err)
		return
	}
	log.Println("Успешное подключение к MongoDB")

	collection := client.Database("inventory").Collection("inventory")

	repository := inventoryRepository.NewRepository(collection)
	service := inventoryService.NewService(repository)
	api := inventoryV1API.NewAPI(service)

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
