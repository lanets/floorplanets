package floorplans

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/server/http/api/server"
)

func RegisterRoutes(server server.APIServer, r *mux.Router) {
	r.Path("/").Handler(floorplansGetHandler(server)).Methods(http.MethodGet)
	r.Path("/").Handler(floorplansPostHandler(server)).Methods(http.MethodPost)
	r.Path("/{id:[0-9]+}").Handler(floorplanGetHandler(server)).Methods(http.MethodGet)
}
