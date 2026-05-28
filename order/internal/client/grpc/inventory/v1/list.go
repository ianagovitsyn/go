package v1

import (
	"context"

	"github.com/ianagovitsyn/project/order/internal/client/converter"
	"github.com/ianagovitsyn/project/order/internal/model"
)

func (c *Client) ListParts(ctx context.Context, partUUIDs []string) ([]model.Part, error) {
	res, err := c.generatedClient.ListParts(ctx, converter.PartUUIDsToGrpc(partUUIDs))
	if err != nil {
		return nil, err
	}

	return converter.PartsToModel(res.Parts), nil
}
