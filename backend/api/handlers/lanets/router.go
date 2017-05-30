package lanets

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.Path("/").Handler(indexHandler()).Methods(http.MethodGet)
}
