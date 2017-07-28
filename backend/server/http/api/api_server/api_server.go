package api_server

import (
	"github.com/jinzhu/gorm"
)

type APIServer interface {
	Database() *gorm.DB
}
