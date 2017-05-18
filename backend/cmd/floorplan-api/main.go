package main

import (
	"fmt"
	"net/http"

	"github.com/lanets/floorplanets/backend/api"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func main() {
	// TODO: Parse command line arguments (for example, to configure port)
	router := api.NewRouter()
	fmt.Println(lanetslogo.Logo)
	fmt.Println("Starting floorplanets API on port 8080.")
	http.ListenAndServe(":8080", router)
}
