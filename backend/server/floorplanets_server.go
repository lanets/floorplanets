// Package server contains the core of the flooplanets server
package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
	"github.com/lanets/floorplanets/backend/server/internal/http"
)

type FloorplanetsServer struct {
	httpServer *http.HttpServer
	database   *gorm.DB
	app        *app.App
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {

	database, err := setupDatabase()
	if err != nil {
		return nil, err
	}

	application := app.NewApp(database)

	httpServer, err := http.NewHttpServer(application, address)
	if err != nil {
		return nil, err
	}

	floorplanetsServer := FloorplanetsServer{
		httpServer: httpServer,
		database:   database,
		app:        application,
	}

	return &floorplanetsServer, nil
}

func setupDatabase() (*gorm.DB, error) {
	database, err := gorm.Open("sqlite3", "database.sqlite")
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	database.AutoMigrate(&models.Floorplan{})

	return database, nil
}

func (s *FloorplanetsServer) Database() *gorm.DB {
	return s.database
}

func (s *FloorplanetsServer) Close() error {
	return s.httpServer.Close()
}
