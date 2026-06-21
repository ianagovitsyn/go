package converter

import (
	"github.com/ianagovitsyn/project/inventory/internal/model"
	inventoryV1 "github.com/ianagovitsyn/project/shared/pkg/proto/inventory/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartToProto(part model.Part) *inventoryV1.Part {
	partProto := inventoryV1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: int64(part.StockQuantity),
		Category:      inventoryV1.Category(part.Category),
		Dimensions: &inventoryV1.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: &inventoryV1.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		Metadata:  metadataToProto(part.Metadata),
		CreatedAt: timestamppb.New(part.CreatedAt),
		UpdatedAt: timestamppb.New(part.UpdatedAt),
	}

	return &partProto
}

func PartsFilterToModel(filter *inventoryV1.PartsFilter) *model.Filter {
	if filter == nil {
		return nil
	}
	categories := make([]model.Category, len(filter.Categories))
	for i, c := range filter.Categories {
		categories[i] = model.Category(c)
	}

	return &model.Filter{
		UUIDs:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            categories,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func PartsToProto(serviceParts []model.Part) []*inventoryV1.Part {
	parts := make([]*inventoryV1.Part, 0, len(serviceParts))

	for _, v := range serviceParts {
		parts = append(parts, PartToProto(v))
	}

	return parts
}

func metadataToProto(m map[string]any) map[string]*inventoryV1.Value {
	if m == nil {
		return nil
	}
	result := make(map[string]*inventoryV1.Value, len(m))
	for k, v := range m {
		result[k] = anyToValue(v)
	}
	return result
}

func anyToValue(v any) *inventoryV1.Value {
	switch val := v.(type) {
	case string:
		return &inventoryV1.Value{Kind: &inventoryV1.Value_StringValue{StringValue: val}}
	case int64:
		return &inventoryV1.Value{Kind: &inventoryV1.Value_Int64Value{Int64Value: val}}
	case int:
		return &inventoryV1.Value{Kind: &inventoryV1.Value_Int64Value{Int64Value: int64(val)}}
	case float64:
		return &inventoryV1.Value{Kind: &inventoryV1.Value_DoubleValue{DoubleValue: val}}
	case bool:
		return &inventoryV1.Value{Kind: &inventoryV1.Value_BoolValue{BoolValue: val}}
	default:
		return nil
	}
}
