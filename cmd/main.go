package main

import (
	"github.com/kerimkhanov/quick_route_map/internal/handler"
	"github.com/kerimkhanov/quick_route_map/internal/repository"
	"github.com/kerimkhanov/quick_route_map/internal/services"
	"github.com/kerimkhanov/quick_route_map/internal/usecases"
	"github.com/kerimkhanov/quick_route_map/pkg/config"
	"github.com/kerimkhanov/quick_route_map/pkg/database/postgres"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	cfg, err := config.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("get config")
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("db connect")
	}
	defer db.Close()

	repo := repository.New(db)
	cr := services.New(repo)
	uc := usecases.New(cr, log)
	delivery := handler.New(uc, log)

	r := delivery.InitRoutes()

	log.Info().Msg("starting on port: %v" + cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Err(err).Msg("run server error")
		return
	}

}
