package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

type Seat struct {
	gorm.Model
	Label string
	X     int
	Y     int
}

type seatJson struct {
	ID    uint   `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	X     int    `json:"label,omitempty"`
	Y     int    `json:"label,omitempty"`
}

func (seat *Seat) ToSeatJson() seatJson {
	seatJson := seatJson{
		ID:    seat.ID,
		Label: seat.Label,
		X:     seat.X,
		Y:     seat.Y,
	}
	return seatJson
}

func (seat *Seat) ToJson() string {
	b, err := json.Marshal(seat.ToSeatJson())
	if err != nil {
		return ""
	}
	return string(b)
}

func SeatFromJson(data io.Reader) (*seatJson, error) {
	decoder := json.NewDecoder(data)
	var s seatJson
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s, err
}
