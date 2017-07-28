// Package server encapsulates operations on the http server
package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lanets/floorplanets/backend/models"
)

type FloorplanetsServer struct {
	httpServer *HttpServer
	database   *gorm.DB
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {
	database, err := setupDatabase()
	if err != nil {
		return nil, err
	}

	httpServer, err := setupHttpServer(address, database)
	if err != nil {
		return nil, err
	}

	s := FloorplanetsServer{
		httpServer: httpServer,
		database:   database,
	}

	return &s, nil
}

func setupHttpServer(address string, database *gorm.DB) (*HttpServer, error) {
	httpServer, err := NewHttpServer(address)
	if err != nil {
		return nil, err
	}

	return httpServer, nil
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

func (s *FloorplanetsServer) Close() error {
	return s.httpServer.Close()
}
