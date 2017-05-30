package lanets

import (
	"fmt"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
	"net/http"
)

func indexHandler() http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, lanetslogo.Logo)
	}
	return http.HandlerFunc(handler)
}
