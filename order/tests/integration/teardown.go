//go:build integration

package integration

import (
	"context"

	"github.com/ianagovitsyn/project/platform/pkg/logger"
	"go.uber.org/zap"
)

func teardownTestEnvironment(ctx context.Context, env *TestEnvironment) {
	logger.Info(ctx, "🧹 Очистка тестового окружения...")

	if env.Postgres != nil {
		if err := env.Postgres.Terminate(ctx); err != nil {
			logger.Error(ctx, "не удалось остановить контейнер PostgreSQL", zap.Error(err))
		} else {
			logger.Info(ctx, "🛑 Контейнер PostgreSQL остановлен")
		}
	}

	logger.Info(ctx, "Тестовое окружение успешно очищено")
}
