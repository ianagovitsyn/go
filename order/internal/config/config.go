package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/ianagovitsyn/project/order/internal/config/env"
)

var appConfig *config

type config struct {
	Logger        LoggerConfig
	OrderHTTP     OrderHTTPConfig
	InventoryGRPC InventoryGRPCConfig
	PaymentGRPC   PaymentGRPCConfig
	Postgres      PostgresConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	OrderHTTPCfg, err := env.NewOrderHTTPConfig()
	if err != nil {
		return err
	}

	InventoryGRPCCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		return err
	}

	PaymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:        loggerCfg,
		OrderHTTP:     OrderHTTPCfg,
		InventoryGRPC: InventoryGRPCCfg,
		PaymentGRPC:   PaymentGRPCCfg,
		Postgres:      postgresCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
