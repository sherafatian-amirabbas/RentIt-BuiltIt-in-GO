package domain

// Plant will hold the information about the equipment
type Plant struct {
	Id          int64   `bson:"_id,omitempty", json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	PricePerDay float64 `json:"PricePerDay"`
}
