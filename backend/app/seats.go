package app

import (
	"fmt"

	"github.com/lanets/floorplanets/backend/models"
)

func (app *App) CreateSeat(floorplanID uint, label string, x, y int) (*models.Seat, error) {
	// Get the corresponding floorplan
	floorplan, err := app.GetFloorplan(int(floorplanID))
	if err != nil {
		return nil, err
	}
	if floorplan == nil {
		return nil, fmt.Errorf("Could not find floorplan with id '%d'", floorplanID)
	}

	// Create the seat
	seat := &models.Seat{
		FloorplanID: floorplan.ID,
		Label:       label,
		X:           x,
		Y:           y,
	}

	err = app.Database.Create(seat).Error
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (app *App) GetSeat(id int) (*models.Seat, error) {
	var seat models.Seat
	if app.Database.First(&seat, id).RecordNotFound() {
		return nil, app.Database.Error
	}
	return &seat, app.Database.Error
}
