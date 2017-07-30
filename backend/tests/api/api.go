package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api"
)

type ApiTest struct {
	*mux.Router
}

func NewApiTest() *ApiTest {
	return &ApiTest{api.NewRouter(nil)}
}

func (apitest *ApiTest) ServeHTTP(*http.Request) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/lanets/", nil)
	res := httptest.NewRecorder()
	api.NewRouter(nil).ServeHTTP(res, req)
	return res
}
