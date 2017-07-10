// Package server encapsulates operations on the http server
package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type FloorplanetsServer struct {
	httpServer *HttpServer
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {

	// Connect to the database
	_, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	// Create the HTTP server
	httpServer, err := NewHttpServer(address)

	if err != nil {
		return nil, err
	}

	s := FloorplanetsServer{httpServer: httpServer}

	return &s, nil
}

func (s *FloorplanetsServer) Close() error {
	return s.httpServer.Close()
}
