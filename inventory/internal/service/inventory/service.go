package inventory

import "github.com/ianagovitsyn/project/inventory/internal/repository"

type Service struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(inventoryRepository repository.InventoryRepository) *Service {
	return &Service{
		inventoryRepository: inventoryRepository,
	}
}
