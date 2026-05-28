package order

import (
	"sync"

	"github.com/ianagovitsyn/project/order/internal/model"
)

type Repository struct {
	mu sync.RWMutex

	storage map[string]model.Order
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string]model.Order),
	}
}
