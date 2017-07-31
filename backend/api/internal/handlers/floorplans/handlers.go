package floorplans

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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
		name := r.FormValue("name")

		err := app.CreateFloorplan(name)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}

func floorplanGetHandler(app *app.App) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		floorplan, err := app.GetFloorplan(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if floorplan == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Fprint(w, floorplan.ToJson())

	}

	handler = decorators.JsonHeaders(handler)

	return http.HandlerFunc(handler)
}
