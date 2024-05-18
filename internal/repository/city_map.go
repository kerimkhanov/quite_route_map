package repository

import "database/sql"

type cityMap struct {
	db *sql.DB
}

func newCityMap(db *sql.DB) *cityMap {
	return &cityMap{
		db: db,
	}
}

func (c *cityMap) GetMap() error {
	//c.db.QueryRow()
	return nil
}
