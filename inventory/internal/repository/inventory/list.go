package inventory

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"

	"github.com/ianagovitsyn/project/inventory/internal/model"
	"github.com/ianagovitsyn/project/inventory/internal/repository/converter"
	modelRepo "github.com/ianagovitsyn/project/inventory/internal/repository/model"
	"github.com/ianagovitsyn/project/platform/pkg/logger"
)

func (r *Repository) GetListByFilter(ctx context.Context, filter *model.Filter) ([]model.Part, error) {
	query := bson.M{}

	if filter != nil {
		if len(filter.UUIDs) > 0 {
			query["_id"] = bson.M{"$in": filter.UUIDs}
		}
		if len(filter.Names) > 0 {
			query["name"] = bson.M{"$in": filter.Names}
		}
		if len(filter.Categories) > 0 {
			query["category"] = bson.M{"$in": filter.Categories}
		}
		if len(filter.ManufacturerCountries) > 0 {
			query["manufacturer.country"] = bson.M{"$in": filter.ManufacturerCountries}
		}
		if len(filter.Tags) > 0 {
			query["tags"] = bson.M{"$in": filter.Tags}
		}
	}

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			logger.Warn(ctx, "failed to close cursor", zap.Error(err))
		}
	}()

	var parts []modelRepo.Part

	if err := cursor.All(ctx, &parts); err != nil {
		return nil, err
	}

	return converter.PartsToModel(parts), nil
}
