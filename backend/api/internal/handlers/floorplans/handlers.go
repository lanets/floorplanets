package floorplans

import (
	"fmt"
	"net/http"

	"github.com/lanets/floorplanets/backend/api/internal/handlers/decorators"
	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
)

func floorplansGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {

		floorplans, err := app.GetAllFloorplansNameId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, models.FloorplanListToJson(floorplans))

	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func floorplansPostHandler(app *app.App) http.Handler {

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

func floorplanGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		floorplan := decorators.FloorplanFromContext(r.Context())
		fmt.Fprint(w, floorplan.ToJson())

	}

	handler = decorators.FloorplanContext(app, handler)
	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}
