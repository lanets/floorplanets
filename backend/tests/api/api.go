package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api"
	"github.com/lanets/floorplanets/backend/app"
	test_app "github.com/lanets/floorplanets/backend/tests/app"
)

type ApiTest struct {
	router  *mux.Router
	testApp *test_app.TestApp
	t       *testing.T
}

func NewApiTest(t *testing.T) *ApiTest {
	testApp := test_app.NewTestApp(t)

	apiTest := ApiTest{
		router:  api.NewRouter(testApp.App),
		testApp: testApp,
		t:       t,
	}

	return &apiTest
}

func (apiTest *ApiTest) ServeHTTP(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	apiTest.router.ServeHTTP(res, req)
	return res
}

func (apiTest *ApiTest) Close() {
	apiTest.testApp.Close()
}

func (apiTest *ApiTest) App() *app.App {
	return apiTest.testApp.App
}
