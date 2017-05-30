package lanets_test

import (
	"net/http"
	"testing"

	api_test "github.com/lanets/floorplanets/backend/tests/api"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func TestIndexHandler(t *testing.T) {
	apitest := api_test.NewApiTest()
	req, _ := http.NewRequest("GET", "/lanets/", nil)
	res := apitest.ServeHTTP(req)

	if res.Body.String() != lanetslogo.Logo {
		t.Error("Expected Lan ETS logo ", res.Body.String())
	}
}

func TextIndexHandlerPost(t *testing.T) {
	apitest := api_test.NewApiTest()
	req, _ := http.NewRequest("POST", "/lanets/", nil)
	res := apitest.ServeHTTP(req)

	if res.Result().StatusCode != 404 {
		t.Error("IndexHandler should not respond to a POST", res.Body.String())
	}
}
