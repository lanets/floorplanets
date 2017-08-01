package app_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/lanets/floorplanets/backend/tests/app"
)

func TestListGetAllFloorplansNameIdEmpty(t *testing.T){
    testApp := app.NewTestApp(t)
    defer testApp.Close()

    floorplans, err := testApp.App.GetAllFloorplansNameId()
    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, 0, len(floorplans), "floorpans should be empty")
}


func TestCreateFloorplan(t *testing.T){
    testApp := app.NewTestApp(t)
    defer testApp.Close()

    floorplan, err := testApp.App.GetFloorplan(1)
    assert.Nil(t, err, "getting an unexisting floorplan should raise no error")
    assert.Nil(t, floorplan, "the floorplan should not exist yet")

    floorplan, err = testApp.App.CreateFloorplan("testfloorplan")
    assert.Nil(t, err, "creating a floorplan should not raise an error")
    assert.Equal(t, "testfloorplan", floorplan.Name)
    assert.Equal(t, uint(1), floorplan.ID)

    floorplan, err = testApp.App.GetFloorplan(1)
    assert.Equal(t, "testfloorplan", floorplan.Name)
    assert.Equal(t, uint(1), floorplan.ID)
}
