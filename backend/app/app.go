// Package app is reponsible for executing the application's actions.
// It dosen't know anything about http, but it knows about the database.
package app

import (
	"github.com/jinzhu/gorm"
)

type App struct {
	Database *gorm.DB
}

func NewApp(database *gorm.DB) *App {
	app := App{
		Database: database,
	}
	return &app
}
