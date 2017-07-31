package floorplans_test

import (
	"net/http"
	"strings"
	"testing"

	api_test "github.com/lanets/floorplanets/backend/tests/api"
)

func TestFloorplansGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/", nil)
	response := apitest.ServeHTTP(request)

	expected := `[]`
	result := response.Body.String()

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Errorf("the status code should be 200, got %d", statusCode)
	}
}

func TestFloorplansGetHandlerEmpty(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	apitest.App().CreateFloorplan("floorplan1")

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/", nil)
	response := apitest.ServeHTTP(request)

	expected := `[{"id":1,"name":"floorplan1"}]`
	result := response.Body.String()

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Errorf("the status code should be 200, got %d", statusCode)
	}
}


func TestFloorplanGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	apitest.App().CreateFloorplan("floorplan1")

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1", nil)
	response := apitest.ServeHTTP(request)

	expected := `{"id":1,"name":"floorplan1"}`
	result := response.Body.String()

	if result != expected {
		t.Errorf("Got `%s`, expected `%s`", result, expected)
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Errorf("the status code should be 200, got %d", statusCode)
	}
}

func TestFloorplanGetHandlerNotFound(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1", nil)
	response := apitest.ServeHTTP(request)

	if response.Body.String() != "" {
		t.Error("the response body should be empty")
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusNotFound {
		t.Errorf("the status code should be 404, got `%d`", statusCode)
	}
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

	expected := `{"id":1,"name":"createdFloorplan"}`
	result := response.Body.String()

	if result != expected {
		t.Errorf("Got `%s`, expected `%s`", result, expected)
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusCreated {
		t.Errorf("the status code should be 201, got `%d`", statusCode)
	}

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

	result := response.Body.String()
	if result != "" {
		t.Errorf("Got `%s`, expected empty body", result)
	}

	statusCode := response.Result().StatusCode
	if statusCode != http.StatusBadRequest {
		t.Errorf("the status code should be 400, got `%d`", statusCode)
	}
}
