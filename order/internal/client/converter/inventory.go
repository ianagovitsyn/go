package converter

import (
	"github.com/ianagovitsyn/project/order/internal/model"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
)

func PartUUIDsToGrpc(uuids []string) *inventoryV1.ListPartsRequest {
	return &inventoryV1.ListPartsRequest{
		Filter: &inventoryV1.PartsFilter{
			Uuids: uuids,
		},
	}
}

func PartsToModel(parts []*inventoryV1.Part) []model.Part {
	result := make([]model.Part, 0, len(parts))
	for _, p := range parts {
		result = append(result, PartToModel(p))
	}
	return result
}

func PartToModel(p *inventoryV1.Part) model.Part {
	part := model.Part{
		UUID:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: int(p.StockQuantity),
		Category:      model.Category(p.Category),
		Tags:          p.Tags,
	}

	if p.Dimensions != nil {
		part.Dimensions = model.Dimensions{
			Length: p.Dimensions.Length,
			Width:  p.Dimensions.Width,
			Height: p.Dimensions.Height,
			Weight: p.Dimensions.Weight,
		}
	}

	if p.Manufacturer != nil {
		part.Manufacturer = model.Manufacturer{
			Name:    p.Manufacturer.Name,
			Country: p.Manufacturer.Country,
			Website: p.Manufacturer.Website,
		}
	}

	if p.Metadata != nil {
		part.Metadata = make(map[string]any, len(p.Metadata))
		for k, v := range p.Metadata {
			part.Metadata[k] = valueToAny(v)
		}
	}

	if p.CreatedAt != nil {
		part.CreatedAt = p.CreatedAt.AsTime()
	}

	if p.UpdatedAt != nil {
		part.UpdatedAt = p.UpdatedAt.AsTime()
	}

	return part
}

func valueToAny(v *inventoryV1.Value) any {
	if v == nil {
		return nil
	}
	switch k := v.Kind.(type) {
	case *inventoryV1.Value_StringValue:
		return k.StringValue
	case *inventoryV1.Value_Int64Value:
		return k.Int64Value
	case *inventoryV1.Value_DoubleValue:
		return k.DoubleValue
	case *inventoryV1.Value_BoolValue:
		return k.BoolValue
	default:
		return nil
	}
}
