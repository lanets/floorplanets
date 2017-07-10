package floorplans

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.Path("/").Handler(floorplansGetHandler()).Methods(http.MethodGet)
	r.Path("/").Handler(floorplansPostHandler()).Methods(http.MethodPost)
	r.Path("/{id:[0-9]+}").Handler(floorplanGetHandler()).Methods(http.MethodGet)
}
