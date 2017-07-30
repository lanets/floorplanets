// Package server encapsulates operations on the http server
package server

type FloorplanetsServer struct {
	httpServer *HttpServer
}

func NewFloorplanetsServer(address string) (*FloorplanetsServer, error) {

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
