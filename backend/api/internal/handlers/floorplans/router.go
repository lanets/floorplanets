package floorplans

import (
	"net/http"

	"github.com/gorilla/mux"

	seats_handlers "github.com/lanets/floorplanets/backend/api/internal/handlers/floorplans/seats"
	"github.com/lanets/floorplanets/backend/app"
)

func RegisterRoutes(app *app.App, r *mux.Router) {
	r.Path("/").Handler(floorplansGetHandler(app)).Methods(http.MethodGet)
	r.Path("/").Handler(floorplansPostHandler(app)).Methods(http.MethodPost)

	// Floorplan router
	floorplanRouter := r.PathPrefix("/{floorplan:[0-9]}").Subrouter()
	floorplanRouter.Path("").Handler(floorplanGetHandler(app)).Methods(http.MethodGet)

	//Register routes for: /seats
	seats_handlers.RegisterRoutes(app, floorplanRouter.PathPrefix("/seats").Subrouter())
}
