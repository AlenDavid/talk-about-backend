package postgres

import (
	"github.com/alendavid/talk-about-backend/pkg/plants"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&plants.Plant{})
}
