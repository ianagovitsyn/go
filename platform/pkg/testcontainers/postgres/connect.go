package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func connectPostgresClient(ctx context.Context, uri string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, errors.Errorf("failed to connect to postgres: %v", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, errors.Errorf("failed to ping postgres: %v", err)
	}

	return conn, nil
}
