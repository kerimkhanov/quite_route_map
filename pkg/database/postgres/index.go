package postgres

import (
	"database/sql"
	"github.com/kerimkhanov/quick_route_map/pkg/config"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL для работы с базой данных
)

func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PGUrl)
	if err != nil {
		return nil, err
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
