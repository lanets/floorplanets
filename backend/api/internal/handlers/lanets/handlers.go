package lanets

import (
	"fmt"
	"net/http"

	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func indexHandler(app *app.App) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, lanetslogo.Logo)
	}
	return http.HandlerFunc(handler)
}
