package lanets

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.Path("/").HandlerFunc(IndexHandler)
}
