package decorators_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanets/floorplanets/backend/api/internal/handlers/decorators"
)

func TestJsonHeaders(t *testing.T) {
	// Create an empty handler
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"hello": "world"}`)
	}
	handler := http.HandlerFunc(handlerFunc)

	// Test the empty handler
	request, _ := http.NewRequest("GET", "/", nil)
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	contentType := responseRecorder.Result().Header.Get("Content-Type")
	if contentType == "application/json" {
		t.Error("Content-Type should not be application/json yet")
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
