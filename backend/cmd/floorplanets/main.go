package main

import (
	"os"

	"github.com/lanets/floorplanets/backend/cmd/floorplanets/cli"
)

func main() {
	// Retrieve args and Shift binary name off argument list.
	args := os.Args[1:]
	cli.Run(args, os.Stdout)
}
