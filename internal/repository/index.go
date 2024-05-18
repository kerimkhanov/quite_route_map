package repository

import (
	"database/sql"
	citymapservice "github.com/kerimkhanov/quick_route_map/internal/services/citymap"
)

type Repository struct {
	CityMap citymapservice.Repository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		CityMap: newCityMap(db),
	}
}
