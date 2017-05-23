package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/lanets/floorplanets/backend/cmd/floorplanets/cli"
)

func main() {
	// Retrieve args and Shift binary name off argument list.
	args := os.Args[1:]

	// Run the CLI, this may return a FloorplanetsServer instance
	srv := cli.Run(args, os.Stdout)

	if srv != nil {

		// Handle SIGINT
		go func() {
			sigchan := make(chan os.Signal, 10)
			signal.Notify(sigchan, os.Interrupt)
			<-sigchan
			fmt.Println("\nStopping server...")
			srv.Close()
			os.Exit(0)
		}()

		// Wait for SIGINT
		select {}

	}

}
