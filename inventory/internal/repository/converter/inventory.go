package converter

import (
	"github.com/ianagovitsyn/project/inventory/internal/model"
	repoModel "github.com/ianagovitsyn/project/inventory/internal/repository/model"
)

func PartToModel(part repoModel.Part) model.Part {
	return model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category),
		Dimensions: model.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: model.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		Metadata:  part.Metadata,
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
	}
}

func PartsToModel(parts []repoModel.Part) []model.Part {
	serviceParts := make([]model.Part, 0, len(parts))

	for _, v := range parts {
		serviceParts = append(serviceParts, PartToModel(v))
	}

	return serviceParts
}
