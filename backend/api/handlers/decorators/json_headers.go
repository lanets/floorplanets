package decorators

import (
	"net/http"
)

func JsonHeaders(h http.HandlerFunc) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	}
	return handler
}
