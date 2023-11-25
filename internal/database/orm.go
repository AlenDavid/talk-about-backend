package database

import (
	"github.com/alendavid/talk-about-backend/internal/persistence/plants"
	"gorm.io/gorm"
)

func setupQueries(db *gorm.DB) {
	plants.SetDefault(db)
}
