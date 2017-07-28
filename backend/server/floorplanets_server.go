// Package server encapsulates operations on the http server
package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lanets/floorplanets/backend/server/internal/models"
	"github.com/lanets/floorplanets/backend/server/http"
)

type FloorplanetsServer struct {
	httpServer *http.HttpServer
	database   *gorm.DB
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {
	s := FloorplanetsServer{}

	database, err := setupDatabase()
	if err != nil {
		return nil, err
	}
	s.database = database

	httpServer, err := setupHttpServer(&s, address)
	if err != nil {
		return nil, err
	}
	s.httpServer = httpServer

	return &s, nil
}

func setupHttpServer(server *FloorplanetsServer, address string) (*http.HttpServer, error) {
	httpServer, err := http.NewHttpServer(server, address)
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

func (s *FloorplanetsServer) Database() *gorm.DB {
	return s.database
}

func (s *FloorplanetsServer) Close() error {
	return s.httpServer.Close()
}
