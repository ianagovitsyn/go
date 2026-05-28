package inventory

import (
	"sync"

	"github.com/ianagovitsyn/project/inventory/internal/repository/model"
)

type Repository struct {
	mu      sync.RWMutex
	storage map[string]model.Part
}

func NewRepository() *Repository {
	return &Repository{
		storage: map[string]model.Part{
			"550e8400-e29b-41d4-a716-446655440001": {
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
				Tags:     []string{"engine", "ion", "interplanetary"},
				Metadata: map[string]any{"thrust": 0.5, "fuel_type": "xenon"},
			},
			"550e8400-e29b-41d4-a716-446655440002": {
				UUID:          "550e8400-e29b-41d4-a716-446655440002",
				Name:          "Titanium Wing",
				Description:   "Титановое крыло повышенной прочности",
				Price:         42000.00,
				StockQuantity: 8,
				Category:      model.CategoryWing,
				Dimensions: model.Dimensions{
					Length: 1200,
					Width:  400,
					Height: 50,
					Weight: 800,
				},
				Manufacturer: model.Manufacturer{
					Name:    "AeroTitan",
					Country: "Japan",
					Website: "https://aerotitan.example.com",
				},
				Tags:     []string{"wing", "titanium", "durable"},
				Metadata: map[string]any{"material": "titanium", "max_load": 5000.0},
			},
			"550e8400-e29b-41d4-a716-446655440003": {
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
				Tags:     []string{"fuel", "plasma", "energy"},
				Metadata: map[string]any{"capacity": 500, "output": "plasma"},
			},
		},
	}
}
