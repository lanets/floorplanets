package floorplans

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api/internal/handlers/decorators"
	"github.com/lanets/floorplanets/backend/app"
)

func floorplansGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Bogus response for now :)
		fmt.Fprint(w, "[]")
	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func floorplansPostHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusCreated)
	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func floorplanGetHandler(app *app.App) http.Handler {

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
