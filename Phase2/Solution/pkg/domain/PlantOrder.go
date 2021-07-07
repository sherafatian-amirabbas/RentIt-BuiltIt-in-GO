package domain

import "time"

// PlantOrder will hold the information about the ordered dates of the equipment
type PlantOrder struct {
	Id        int64 `bson:"_id,omitempty", json:"Id"`
	PlantId   int64
	StartDate time.Time
	EndDate   time.Time
}
