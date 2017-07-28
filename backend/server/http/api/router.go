package api

import (
	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/server/http/api/server"
	floorplans_handlers "github.com/lanets/floorplanets/backend/server/http/api/handlers/floorplans"
	lanets_handlers "github.com/lanets/floorplanets/backend/server/http/api/handlers/lanets"
)

//NewRouter creates the main router
func NewRouter(server server.APIServer) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Register routes for: /lanets
	lanets_handlers.RegisterRoutes(
		server,
		router.PathPrefix("/lanets").Subrouter(),
	)

	//Register routes for: /floorplan
	floorplans_handlers.RegisterRoutes(
		server,
		router.PathPrefix("/floorplans").Subrouter(),
	)

	return router
}
