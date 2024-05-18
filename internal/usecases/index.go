package usecases

import (
	"github.com/kerimkhanov/quick_route_map/internal/services"
	"github.com/rs/zerolog"
)

type UseCase struct {
	cr  *services.Core
	log zerolog.Logger
}

func New(cr *services.Core, log zerolog.Logger) *UseCase {
	return &UseCase{
		cr:  cr,
		log: log,
	}
}
