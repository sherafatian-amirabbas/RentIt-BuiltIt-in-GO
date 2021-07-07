package domain

type PlantPrice struct {
	PlantId          int64   `bson:"_id,omitempty", json:"PlantId"`
	StartDate        string  `json:"StartDate"`
	EndDate          string  `json:"EndDate"`
	PricePerDuration float64 `json:"PricePerDuration"`
}
