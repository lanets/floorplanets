package lanets

import (
	"fmt"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, lanetslogo.Logo)
}
