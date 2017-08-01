package lanets_test

import (
	"net/http"
	"testing"

	api_test "github.com/lanets/floorplanets/backend/tests/api"
	"github.com/lanets/floorplanets/backend/util/lanetslogo"
)

func TestIndexHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	req, _ := http.NewRequest(http.MethodGet, "/lanets/", nil)
	res := apitest.ServeHTTP(req)

	if res.Body.String() != lanetslogo.Logo {
		t.Error("Expected Lan ETS logo ", res.Body.String())
	}
}

func TextIndexHandlerPost(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	req, _ := http.NewRequest(http.MethodPost, "/lanets/", nil)
	res := apitest.ServeHTTP(req)

	if res.Result().StatusCode != http.StatusNotFound {
		t.Error("IndexHandler should not respond to a POST", res.Body.String())
	}
}
