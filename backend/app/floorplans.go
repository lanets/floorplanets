package app

import (
	"github.com/lanets/floorplanets/backend/models"
)

func (app *App) CreateFloorplan(name string) error {
	return app.Database.Create(&models.Floorplan{Name: name}).Error
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
