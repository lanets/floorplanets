package models

import (
	"github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&Floorplan{})
	db.AutoMigrate(&Seat{})
}
