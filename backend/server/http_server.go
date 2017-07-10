package server

import (
	"net"
	"net/http"

	"github.com/lanets/floorplanets/backend/api"
)

type HttpServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func NewHttpServer(address string) (*HttpServer, error) {
	// Create a router
	router := api.NewRouter()

	// Create the logged handler
	loggedHandler := logHandler(router)

	// Create an http server with the logged handler
	httpServer := &http.Server{Handler: loggedHandler}

	// Listen on the given address
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	server := HttpServer{httpServer: httpServer, listener: listener}

	go httpServer.Serve(listener)

	return &server, nil
}

func (s *HttpServer) Close() error {
	s.listener.Close()
	s.httpServer.Close()
	return nil
}

func (s *HttpServer) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}
