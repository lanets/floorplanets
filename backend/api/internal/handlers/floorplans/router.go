package floorplans

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/app"
)

func RegisterRoutes(app *app.App, r *mux.Router) {
	r.Path("/").Handler(floorplansGetHandler(app)).Methods(http.MethodGet)
	r.Path("/").Handler(floorplansPostHandler(app)).Methods(http.MethodPost)
	r.Path("/{floorplan:[0-9]+}").Handler(floorplanGetHandler(app)).Methods(http.MethodGet)
}
