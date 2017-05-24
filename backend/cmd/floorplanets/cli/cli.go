package cli

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/lanets/floorplanets/backend/api"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func Run(args []string, writerOutput io.Writer) {

	fs := flag.NewFlagSet("default", flag.ExitOnError)

	var h, help bool
	fs.BoolVar(&help, "help", false, "")
	fs.BoolVar(&h, "h", false, "")

	var port int
	fs.IntVar(&port, "port", 8080, "port to listen to")

	fs.Parse(args)

	if h || help {
		fmt.Fprintln(writerOutput, `Usage: floorplanets -port <port number>`)
		return
	}

	router := api.NewRouter()
	fmt.Fprintln(writerOutput, lanetslogo.Logo)
	fmt.Fprintf(writerOutput, "Starting floorplanets API on port %d.\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
