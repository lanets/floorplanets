package api

import (
	"github.com/gorilla/mux"
	floorplans_handlers "github.com/lanets/floorplanets/backend/api/handlers/floorplans"
	lanets_handlers "github.com/lanets/floorplanets/backend/api/handlers/lanets"
)

//NewRouter creates the main router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Register routes for: /lanets
	lanets_handlers.RegisterRoutes(router.PathPrefix("/lanets").Subrouter())

	//Register routes for: /floorplan
	floorplans_handlers.RegisterRoutes(router.PathPrefix("/floorplans").Subrouter())

	return router
}
