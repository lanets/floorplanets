package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"

	"github.com/lanets/floorplanets/backend/api"
)

type ApiTest struct {
	router *mux.Router
}

func NewApiTest() *ApiTest {
	return &ApiTest{
		router: api.NewRouter(nil),
	}
}

func (apitest *ApiTest) ServeHTTP(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	apitest.router.ServeHTTP(res, req)
	return res
}
