package inventory

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"

	"github.com/ianagovitsyn/project/inventory/internal/repository/model"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	repo := &Repository{
		collection: collection,
	}

	parts := []any{
		model.Part{
			UUID:          "550e8400-e29b-41d4-a716-446655440001",
			Name:          "Ion Engine",
			Description:   "Ионный двигатель для межпланетных перелётов",
			Price:         75000.00,
			StockQuantity: 12,
			Category:      model.CategoryEngine,
			Dimensions: model.Dimensions{
				Length: 320,
				Width:  180,
				Height: 200,
				Weight: 1500,
			},
			Manufacturer: model.Manufacturer{
				Name:    "SpaceDrive Corp",
				Country: "Germany",
				Website: "https://spacedrive.example.com",
			},
			Tags:      []string{"engine", "ion", "interplanetary"},
			Metadata:  map[string]any{"thrust": 0.5, "fuel_type": "xenon"},
			CreatedAt: time.Now(),
		},
		model.Part{
			UUID:          "550e8400-e29b-41d4-a716-446655440003",
			Name:          "Plasma Fuel Cell",
			Description:   "Плазменный топливный элемент",
			Price:         18500.50,
			StockQuantity: 25,
			Category:      model.CategoryFuel,
			Dimensions: model.Dimensions{
				Length: 60,
				Width:  60,
				Height: 120,
				Weight: 200,
			},
			Manufacturer: model.Manufacturer{
				Name:    "FuelTech",
				Country: "Germany",
				Website: "https://fueltech.example.com",
			},
			Tags:      []string{"fuel", "plasma", "energy"},
			Metadata:  map[string]any{"capacity": 500, "output": "plasma"},
			CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
	}

	ctx := context.Background()
	res, err := repo.collection.InsertMany(ctx, parts)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)

	return repo
}
