package models

import (
	"github.com/jinzhu/gorm"
)

type Floorplan struct {
	gorm.Model
	Name string
}
