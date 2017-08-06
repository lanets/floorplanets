package seats_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	api_test "github.com/lanets/floorplanets/backend/tests/api"
)

func TestSeatsGetHandlerEmpty(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	_, err := apitest.App().CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1/seats", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "[]", response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestSeatsGetHandlerBadFloorplan(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1/seats", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "", response.Body.String())
	assert.Equal(t, http.StatusNotFound, response.Result().StatusCode)
}

func TestSeatsGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	floorplan, err := apitest.App().CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	_, err = apitest.App().CreateSeat(floorplan.ID, "A-1", 2, 3)
	assert.Nil(t, err)

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1/seats", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `[{"id":1,"label":"A-1","x":2,"y":3}]`, response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestSeatGetHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	floorplan, err := apitest.App().CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	_, err = apitest.App().CreateSeat(floorplan.ID, "A-1", 2, 3)
	assert.Nil(t, err)

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/1/seats/1", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `{"id":1,"label":"A-1","x":2,"y":3}`, response.Body.String())
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestSeatGetHandlerFloorplanNotFound(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	floorplan, err := apitest.App().CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	_, err = apitest.App().CreateSeat(floorplan.ID, "A-1", 2, 3)
	assert.Nil(t, err)

	request, _ := http.NewRequest(http.MethodGet, "/floorplans/2/seats/1", nil)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, "", response.Body.String())
	assert.Equal(t, http.StatusNotFound, response.Result().StatusCode)
}

func TestSeatsPostHandler(t *testing.T) {
	apitest := api_test.NewApiTest(t)
	defer apitest.Close()

	_, err := apitest.App().CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	request, _ := http.NewRequest(
		http.MethodPost,
		"/floorplans/1/seats",
		strings.NewReader(`{"label":"A-1","x":2,"y":3}`),
	)
	response := apitest.ServeHTTP(request)

	assert.Equal(t, `{"id":1,"label":"A-1","x":2,"y":3}`, response.Body.String())
	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
}
