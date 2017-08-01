package lanets

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/app"
)

func RegisterRoutes(app *app.App, r *mux.Router) {
	r.Path("/").Handler(indexHandler(app)).Methods(http.MethodGet)
}
