package app

import (
	"github.com/lanets/floorplanets/backend/models"
)

func (app *App) CreateFloorplan(name string) (*models.Floorplan, error) {
	floorplan := &models.Floorplan{
		Name: name,
	}

	err := app.Database.Create(floorplan).Error
	if err != nil {
		return nil, err
	}

	return floorplan, nil
}

func (app *App) GetFloorplan(id int) (*models.Floorplan, error) {
	var floorplan models.Floorplan
	if app.Database.First(&floorplan, id).RecordNotFound() {
		return nil, app.Database.Error
	}
	return &floorplan, app.Database.Error
}

func (app *App) GetAllFloorplansNameId() ([]models.Floorplan, error) {
	var floorplans []models.Floorplan

	app.Database.Model(&models.Floorplan{}).Select("name, id").Find(&floorplans)

	err := app.Database.Error
	if err != nil {
		return nil, err
	}

	return floorplans, nil
}
