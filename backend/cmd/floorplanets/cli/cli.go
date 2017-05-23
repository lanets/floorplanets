// Package cli is responsible for parsing command line arguments and creating
// a server instance.
package cli

import (
	"flag"
	"fmt"
	"io"

	"github.com/lanets/floorplanets/backend/server"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func Run(args []string, writerOutput io.Writer) *server.FloorplanetsServer {

	fs := flag.NewFlagSet("default", flag.ExitOnError)

	var h, help bool
	fs.BoolVar(&help, "help", false, "")
	fs.BoolVar(&h, "h", false, "")

	var address string
	fs.StringVar(&address, "address", "0.0.0.0", "address to listen to")

	var port int
	fs.IntVar(&port, "port", 8080, "port to listen to")

	fs.Parse(args)

	if h || help {
		fmt.Fprintln(writerOutput, `Usage: floorplanets -port <port number>`)
		return nil
	}

	fmt.Fprintln(writerOutput, lanetslogo.Logo)

	listenAddress := fmt.Sprintf("%s:%d", address, port)

	fmt.Fprintf(writerOutput, "Starting floorplanets API on %s.\n", listenAddress)

	srv, _ := server.NewFloorplanetsServer(listenAddress)

	return srv
}
