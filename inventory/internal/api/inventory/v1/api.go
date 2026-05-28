package v1

import (
	"github.com/ianagovitsyn/project/inventory/internal/service"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
)

type Api struct {
	inventoryV1.UnimplementedInventoryServiceServer

	inventoryService service.InventoryService
}

func NewAPI(inventoryService service.InventoryService) *Api {
	return &Api{
		inventoryService: inventoryService,
	}
}
