// Package server encapsulates operations on the http server
package server

import (
	"net"
	"net/http"

	"github.com/lanets/floorplanets/backend/api"
)

type FloorplanetsServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {
	// Create a router
	router := api.NewRouter()

	// Create an http server with the router
	httpServer := &http.Server{Handler: router}

	// Listen on the given address
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	s := FloorplanetsServer{httpServer: httpServer, listener: listener}

	go httpServer.Serve(listener)

	return &s, nil
}

func (s *FloorplanetsServer) Close() error {
	s.listener.Close()
	s.httpServer.Close()
	return nil
}

func (s *FloorplanetsServer) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}
