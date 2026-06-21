package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"go.uber.org/zap"
)

const (
	postgresPort           = "5432"
	postgresStartupTimeout = 1 * time.Minute
)

type Container struct {
	container testcontainers.Container
	conn      *pgx.Conn
	cfg       *Config
}

func NewContainer(ctx context.Context, opts ...Option) (*Container, error) {
	cfg := buildConfig(opts...)

	container, err := startPostgresContainer(ctx, cfg)
	if err != nil {
		return nil, err
	}

	success := false
	defer func() {
		if !success {
			if err = container.Terminate(ctx); err != nil {
				cfg.Logger.Error(ctx, "failed to terminate postgres container", zap.Error(err))
			}
		}
	}()

	cfg.Host, cfg.Port, err = getContainerHostPort(ctx, container)
	if err != nil {
		return nil, err
	}

	uri := buildPostgresURI(cfg)

	conn, err := connectPostgresClient(ctx, uri)
	if err != nil {
		return nil, err
	}

	cfg.Logger.Info(ctx, "Postgres container started", zap.String("uri", uri))
	success = true

	return &Container{
		container: container,
		conn:      conn,
		cfg:       cfg,
	}, nil
}

func (c *Container) Conn() *pgx.Conn {
	return c.conn
}

func (c *Container) Config() *Config {
	return c.cfg
}

func (c *Container) URI() string {
	return buildPostgresURI(c.cfg)
}

func (c *Container) Terminate(ctx context.Context) error {
	if err := c.conn.Close(ctx); err != nil {
		c.cfg.Logger.Error(ctx, "failed to close postgres connection", zap.Error(err))
	}

	if err := c.container.Terminate(ctx); err != nil {
		c.cfg.Logger.Error(ctx, "failed to terminate postgres container", zap.Error(err))
	}

	c.cfg.Logger.Info(ctx, "Postgres container terminated")

	return nil
}
