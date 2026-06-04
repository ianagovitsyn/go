package order

import (
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func NewRepository(DB *pgx.Conn) *Repository {
	return &Repository{
		DB: DB,
	}
}
