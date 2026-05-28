package v1

import (
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
)

type Client struct {
	generatedClient inventoryV1.InventoryServiceClient
}

func NewClient(generatedClient inventoryV1.InventoryServiceClient) *Client {
	return &Client{
		generatedClient: generatedClient,
	}
}
