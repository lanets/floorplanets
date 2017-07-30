// Package app is reponsible for executing the application's actions.
// It dosen't know anything about http.
package app

import (
	"github.com/jinzhu/gorm"
)

type App struct {
	database *gorm.DB
}

func NewApp(database *gorm.DB) *App {
	app := App{
		database: database,
	}
	return &app
}
