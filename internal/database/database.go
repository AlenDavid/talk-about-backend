package database

import (
	"context"

	"github.com/alendavid/talk-about-backend/internal/database/postgres"
	"github.com/alendavid/talk-about-backend/internal/persistence/plants"
	"gorm.io/gorm"
)

type Database struct {
	DB gorm.DB
}

// TODO: use context here
func (d Database) Plants() plants.IPlantDo {
	return plants.Q.Plant.WithContext(context.TODO())
}

func Run(host string, port int, user, password, name string) (Database, error) {
	db, err := postgres.Run(host, port, user, password, name)
	if err != nil {
		return Database{}, err
	}

	// TODO: instead of using global queries, it would be best to broadcast this responsibility up.
	setupQueries(db)

	return Database{*db}, nil
}
