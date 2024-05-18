package handler

import (
	"github.com/kerimkhanov/quick_route_map/internal/usecases"
	"github.com/rs/zerolog"
)

type Handler struct {
	uc  *usecases.UseCase
	log zerolog.Logger
}

func New(uc *usecases.UseCase, log zerolog.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}
