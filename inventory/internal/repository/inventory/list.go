package inventory

import (
	"context"
	"slices"

	"github.com/ianagovitsyn/project/inventory/internal/model"
	"github.com/ianagovitsyn/project/inventory/internal/repository/converter"
	modelRepo "github.com/ianagovitsyn/project/inventory/internal/repository/model"
)

func (r *Repository) GetListByFilter(_ context.Context, filter *model.Filter) ([]model.Part, error) {
	parts := make([]modelRepo.Part, 0, len(r.storage))

	for _, part := range r.storage {
		parts = append(parts, part)
	}

	if filter == nil {
		return converter.PartsToModel(parts), nil
	}

	if len(filter.UUIDs) > 0 {
		var filtered []modelRepo.Part
		for _, part := range parts {
			if slices.Contains(filter.UUIDs, part.UUID) {
				filtered = append(filtered, part)
			}
		}
		parts = filtered
	}

	if len(filter.Names) > 0 {
		var filtered []modelRepo.Part
		for _, part := range parts {
			if slices.Contains(filter.Names, part.Name) {
				filtered = append(filtered, part)
			}
		}
		parts = filtered
	}

	if len(filter.Categories) > 0 {
		var filtered []modelRepo.Part
		for _, part := range parts {
			if slices.Contains(filter.Categories, model.Category(part.Category)) {
				filtered = append(filtered, part)
			}
		}
		parts = filtered
	}

	if len(filter.ManufacturerCountries) > 0 {
		var filtered []modelRepo.Part
		for _, part := range parts {
			if slices.Contains(filter.ManufacturerCountries, part.Manufacturer.Country) {
				filtered = append(filtered, part)
			}
		}
		parts = filtered
	}

	if len(filter.Tags) > 0 {
		var filtered []modelRepo.Part
		for _, part := range parts {
			found := slices.ContainsFunc(filter.Tags, func(filterTag string) bool {
				return slices.Contains(part.Tags, filterTag)
			})

			if found {
				filtered = append(filtered, part)
			}
		}
		parts = filtered
	}
	return converter.PartsToModel(parts), nil
}
