package floorplans

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.Path("/").Handler(floorplansGetHandler()).Methods("GET")
	r.Path("/{id:[0-9]+}").Handler(floorplanGetHandler()).Methods("GET")
}
