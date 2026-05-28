package v1

import (
	"context"

	"github.com/ianagovitsyn/project/inventory/internal/converter"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
)

func (a *Api) ListParts(ctx context.Context, r *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	parts, err := a.inventoryService.List(ctx, converter.PartsFilterToModel(r.Filter))
	if err != nil {
		return nil, err
	}

	return &inventoryV1.ListPartsResponse{
		Parts: converter.PartsToProto(parts),
	}, nil
}
