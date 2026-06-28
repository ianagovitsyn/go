package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ianagovitsyn/project/inventory/internal/converter"
	"github.com/ianagovitsyn/project/inventory/internal/model"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
)

func (a *Api) GetPart(ctx context.Context, r *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.Get(ctx, r.Uuid)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, "part not found")
		}
		return nil, err
	}

	return &inventoryV1.GetPartResponse{
		Part: converter.PartToProto(part),
	}, nil
}
