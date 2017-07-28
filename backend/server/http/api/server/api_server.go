package server

import (
	"github.com/jinzhu/gorm"
)

//APIServer is everything that the API package needs to work
type APIServer interface {
	Database() *gorm.DB
}
