package lanets

import (
	"fmt"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, lanetslogo.Logo)
}
