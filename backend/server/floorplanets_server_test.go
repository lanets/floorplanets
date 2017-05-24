package server_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lanets/floorplanets/backend/server"
)

func getFreePort() int {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func TestStartStopServer(t *testing.T) {
	port := getFreePort()

	s, _ := server.NewFloorplanetsServer(fmt.Sprintf(":%d", port))

	// Something should be listening on this port
	_, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		t.Error("Server is not listening")
	}

	assert.Equal(t, s.Port(), port)

	// Stop the server
	s.Close()

	_, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		t.Error("Server did not close")
	}
}
