package seats

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/app"
)

func RegisterRoutes(app *app.App, r *mux.Router) {
	r.Path("").Handler(seatsGetHandler(app)).Methods(http.MethodGet)
	r.Path("").Handler(seatsPostHandler(app)).Methods(http.MethodPost)
	r.Path("/{id:[0-9]+}").Handler(seatGetHandler(app)).Methods(http.MethodGet)
}
