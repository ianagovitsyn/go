//go:build e2e

package e2e

import (
	"context"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (env *TestEnvironment) ClearInventoryCollection(ctx context.Context) error {
	dbName := os.Getenv("MONGO_DATABASE")

	_, err := env.Mongo.Client().Database(dbName).Collection(inventoryCollectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func (env *TestEnvironment) InsertTestPart(ctx context.Context) (string, error) {
	partUUID := gofakeit.UUID()
	part := bson.M{
		"_id":            partUUID,
		"name":           gofakeit.ProductName(),
		"description":    gofakeit.LoremIpsumSentence(10),
		"price":          gofakeit.Price(10, 1000),
		"stock_quantity": gofakeit.Number(1, 100),
		"category":       1,
		"dimensions": bson.M{
			"length": gofakeit.Float64Range(1, 100),
			"width":  gofakeit.Float64Range(1, 100),
			"height": gofakeit.Float64Range(1, 100),
			"weight": gofakeit.Float64Range(0.1, 50),
		},
		"manufacturer": bson.M{
			"name":    gofakeit.Company(),
			"country": gofakeit.Country(),
			"website": gofakeit.URL(),
		},
		"tags":       []string{gofakeit.Word(), gofakeit.Word()},
		"metadata":   bson.M{},
		"created_at": primitive.NewDateTimeFromTime(time.Now()),
		"updated_at": primitive.NewDateTimeFromTime(time.Now()),
	}

	databaseName := os.Getenv("MONGO_DATABASE")

	_, err := env.Mongo.Client().Database(databaseName).Collection(inventoryCollectionName).InsertOne(ctx, part)
	if err != nil {
		return "", err
	}

	return partUUID, nil
}
