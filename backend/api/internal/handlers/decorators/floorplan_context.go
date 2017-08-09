package decorators

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
)

func FloorplanContext(app *app.App, h http.HandlerFunc) http.HandlerFunc {

	handler := func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["floorplan"])

		floorplan, err := app.GetFloorplan(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if floorplan == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "floorplan", floorplan)

		r = r.WithContext(ctx)

		h(w, r)
	}

	return handler
}

func FloorplanFromContext(ctx context.Context) *models.Floorplan {
	floorplan, ok := ctx.Value("floorplan").(*models.Floorplan)
	if !ok {
		panic("floorplan was not found in context")
	}
	return floorplan
}
