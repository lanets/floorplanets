// Package api provides a router that calls app actions
package api

import (
	"github.com/gorilla/mux"

	floorplans_handlers "github.com/lanets/floorplanets/backend/api/internal/handlers/floorplans"
	lanets_handlers "github.com/lanets/floorplanets/backend/api/internal/handlers/lanets"
	"github.com/lanets/floorplanets/backend/app"
)

//NewRouter creates the main router
func NewRouter(app *app.App) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Register routes for: /lanets
	lanets_handlers.RegisterRoutes(
		app,
		router.PathPrefix("/lanets").Subrouter(),
	)

	//Register routes for: /floorplan
	floorplans_handlers.RegisterRoutes(
		app,
		router.PathPrefix("/floorplans").Subrouter(),
	)

	return router
}
