package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

type Floorplan struct {
	gorm.Model
	Name string
}

// floorplanJson is used to serialize and deserialize Floorplan
type floorplanJson struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (floorplan *Floorplan) ToFloorplanJson() floorplanJson {
	floorplanJson := floorplanJson{
		ID:   floorplan.ID,
		Name: floorplan.Name,
	}
	return floorplanJson
}

func (floorplan *Floorplan) ToJson() string {
	b, err := json.Marshal(floorplan.ToFloorplanJson())
	if err != nil {
		return ""
	}

	return string(b)
}

func FloorplanFromJson(data io.Reader) *floorplanJson {
	decoder := json.NewDecoder(data)
	var f floorplanJson
	err := decoder.Decode(&f)
	if err == nil {
		return &f
	}
	return nil
}

func FloorplanListToJson(floorplans []Floorplan) string {
	var floorplanJsons []floorplanJson

	for _, floorplan := range floorplans {
		floorplanJsons = append(floorplanJsons, floorplan.ToFloorplanJson())
	}

	b, err := json.Marshal(floorplanJsons)
	if err != nil {
		return ""
	}

	return string(b)
}
