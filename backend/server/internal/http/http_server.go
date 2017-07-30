package http

import (
	"net"
	"net/http"

	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/api"
)

type HttpServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func NewHttpServer(app *app.App, address string) (*HttpServer, error) {
	// Create a router
	router := api.NewRouter(app)

	// Create the logged handler
	loggedHandler := logHandler(router)

	// Listen on the given address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	// Create an http server with the logged handler
	httpServer := &http.Server{
		Handler: loggedHandler,
	}

	server := HttpServer{
		httpServer: httpServer,
		listener:   listener,
	}

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
