package floorplans

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api/handlers/decorators"
)

func floorplansGetHandler() http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Bogus response for now :)
		fmt.Fprint(w, "[]")
	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func floorplanGetHandler() http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// For now, hardcode only one floorplan
		if id != 1 {
			http.NotFound(w, r)
			return
		}

		// :)
		fmt.Fprint(w, `{ "0": { name: "A-1", x: 957, y: 171, label: "A-1" } }`)
	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}
