package seats

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api/internal/handlers/decorators"
	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
)

func seatsGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		floorplan := decorators.FloorplanFromContext(r.Context())

		err := app.LoadFloorplanSeats(floorplan)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, models.SeatListtoJson(floorplan.Seats))

	}

	handler = decorators.FloorplanContext(app, handler)
	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func seatsPostHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {

		postedFloorplan, err := models.FloorplanFromJson(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		floorplan, err := app.CreateFloorplan(postedFloorplan.Name)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, floorplan.ToJson())

	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func seatGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["seat"])

		seat, err := app.GetSeat(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if seat == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Fprint(w, seat.ToJson())

	}

	handler = decorators.FloorplanContext(app, handler)
	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}
