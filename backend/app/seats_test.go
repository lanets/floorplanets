package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lanets/floorplanets/backend/tests/app"
)

func TestGetSeatNil(t *testing.T) {
	testApp := app.NewTestApp(t)
	defer testApp.Close()

	seat, err := testApp.App.GetSeat(1)

	assert.Nil(t, err, "getting an unexisting seat should raise no error")
	assert.Nil(t, seat, "the seat should not exist yet")
}

func TestCreateSeatBadFloorplanID(t *testing.T) {
	testApp := app.NewTestApp(t)
	defer testApp.Close()

	seat, err := testApp.App.CreateSeat(1, "A-1", 0, 0)

	assert.NotNil(t, err, "creating a seat with a bad floorplan ID shoul raise an error")
	assert.Nil(t, seat, "no seat should be returned")
	assert.Equal(t, "Could not find floorplan with id '1'", err.Error())
}

func TestCreateSeat(t *testing.T) {
	testApp := app.NewTestApp(t)
	defer testApp.Close()

	floorplan, err := testApp.App.CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	seat, err := testApp.App.CreateSeat(floorplan.ID, "A-1", 2, 3)
	assert.NotNil(t, seat, "CreateSeat should return a seat")
	assert.Nil(t, err)
	assert.Equal(t, "A-1", seat.Label)
	assert.Equal(t, uint(1), seat.ID)
	assert.Equal(t, 2, seat.X)
	assert.Equal(t, 3, seat.Y)
}

func TestLoadSeats(t *testing.T) {
	testApp := app.NewTestApp(t)
	defer testApp.Close()

	floorplan, err := testApp.App.CreateFloorplan("floorplan1")
	assert.Nil(t, err)

	err = testApp.App.LoadFloorplanSeats(floorplan)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(floorplan.Seats))

	_, err = testApp.App.CreateSeat(floorplan.ID, "A-1", 2, 3)
	assert.Nil(t, err)

	err = testApp.App.LoadFloorplanSeats(floorplan)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(floorplan.Seats))
}
