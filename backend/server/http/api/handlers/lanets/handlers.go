package lanets

import (
	"fmt"
	"net/http"

	"github.com/lanets/floorplanets/backend/util/lanetslogo"
	"github.com/lanets/floorplanets/backend/server/http/api/server"
)

func indexHandler(server server.APIServer) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, lanetslogo.Logo)
	}
	return http.HandlerFunc(handler)
}
