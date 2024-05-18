package services

import (
	"github.com/kerimkhanov/quick_route_map/internal/repository"
	citymapservice "github.com/kerimkhanov/quick_route_map/internal/services/citymap"
)

type Core struct {
	CityMap *citymapservice.CityMap
}

func New(repository *repository.Repository) *Core {
	return &Core{
		CityMap: citymapservice.New(repository.CityMap),
	}
}
