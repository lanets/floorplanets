package floorplans_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	api_test "github.com/lanets/floorplanets/backend/tests/api"
)

func TestFloorplansGetHandlerEmpty(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "[]", response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestFloorplansGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	apitest.App().CreateFloorplan("floorplan1")

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `[{"id":1,"name":"floorplan1"}]`, response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestFloorplanGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	apitest.App().CreateFloorplan("floorplan1")

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `{"id":1,"name":"floorplan1"}`, response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestFloorplanGetHandlerNotFound(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "", response.Body.String(), "the response body should be empty")
	assert.Equal(t, http.StatusNotFound, response.Result().StatusCode)
}

func TestFloorplansPostHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(
		http.MethodPost,
		"/floorplans/",
		strings.NewReader(`{"name":"createdFloorplan"}`),
	)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `{"id":1,"name":"createdFloorplan"}`, response.Body.String())
	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
}

func TestFloorplansPostHandlerBadInput(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(
		http.MethodPost,
		"/floorplans/",
		strings.NewReader(`Hello :) {::::}`),
	)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "", response.Body.String())
	assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
}
