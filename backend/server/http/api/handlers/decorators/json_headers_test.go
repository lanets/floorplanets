package decorators_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/lanets/floorplanets/backend/server/http/api/handlers/decorators"
)

func TestJsonHeaders(t *testing.T) {
	// Create an empty handler
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
	}
	handler := http.HandlerFunc(handlerFunc)

	// Test the empty handler
	request, _ := http.NewRequest("GET", "/", nil)
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	contentType := responseRecorder.Result().Header.Get("Content-Type")
	if contentType != "" {
		t.Error("Content-Type should not be set", contentType)
	}

	// Decoreate the handler
	handlerFunc = decorators.JsonHeaders(handlerFunc)
	handler = http.HandlerFunc(handlerFunc)

	// Test the decorated handler
	request, _ = http.NewRequest("GET", "/", nil)
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	contentType = responseRecorder.Result().Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Error("Content-Type should be application/json: ", contentType)
	}
}
