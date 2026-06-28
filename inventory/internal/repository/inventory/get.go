package inventory

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ianagovitsyn/project/inventory/internal/model"
	"github.com/ianagovitsyn/project/inventory/internal/repository/converter"
	repoModel "github.com/ianagovitsyn/project/inventory/internal/repository/model"
)

func (r *Repository) GetByUUID(ctx context.Context, uuid string) (model.Part, error) {
	bsonUUID := bson.M{"_id": uuid}
	var part repoModel.Part
	err := r.collection.FindOne(ctx, bsonUUID).Decode(&part)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Part{}, repoModel.ErrNotFound
	}

	return converter.PartToModel(part), nil
}
