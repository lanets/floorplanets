package lanets

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/server/http/api/server"
)

func RegisterRoutes(server server.APIServer, r *mux.Router) {
	r.Path("/").Handler(indexHandler(server)).Methods(http.MethodGet)
}
