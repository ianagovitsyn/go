//go:build integration

package integration

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"github.com/ianagovitsyn/project/order/internal/migrator"
	orderRepo "github.com/ianagovitsyn/project/order/internal/repository/order"
	"github.com/ianagovitsyn/project/platform/pkg/logger"
	"github.com/ianagovitsyn/project/platform/pkg/testcontainers/postgres"
)

const (
	migrationsDir = "../../migrations"
)

type TestEnvironment struct {
	Postgres *postgres.Container
	Repo     *orderRepo.Repository
}

func setupTestEnvironment(ctx context.Context) *TestEnvironment {
	logger.Info(ctx, "🚀 Подготовка тестового окружения...")

	pg, err := postgres.NewContainer(ctx,
		postgres.WithDatabase("orders_test"),
		postgres.WithLogger(logger.Logger()),
	)
	if err != nil {
		logger.Fatal(ctx, "не удалось запустить контейнер PostgreSQL", zap.Error(err))
	}
	logger.Info(ctx, "✅ Контейнер PostgreSQL успешно запущен")

	db, err := sql.Open("pgx", pg.URI())
	if err != nil {
		logger.Fatal(ctx, "не удалось открыть подключение для миграций", zap.Error(err))
	}
	defer db.Close()

	m := migrator.NewMigrator(db, migrationsDir)
	if err = m.Up(); err != nil {
		logger.Fatal(ctx, "не удалось выполнить миграции", zap.Error(err))
	}
	logger.Info(ctx, "✅ Миграции успешно применены")

	repo := orderRepo.NewRepository(pg.Conn())

	logger.Info(ctx, "🎉 Тестовое окружение готово")
	return &TestEnvironment{
		Postgres: pg,
		Repo:     repo,
	}
}
