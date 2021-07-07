package dto

import (
	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
)

// customer will hold the information about the customer
type Plant struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PricePerDay float64 `json:"priceperday"`
}

func GetPlantDTOList(plants []*domain.Plant) []*Plant {

	var dtoList = []*Plant{}
	for _, plant := range plants {
		dto := GetPlantDTO(plant)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetPlantDTO(plant *domain.Plant) *Plant {

	dto := &Plant{
		Id:   plant.ID,
		Name: plant.Name,
		Description: plant.Description,
		PricePerDay: plant.PricePerDay,
	}

	return dto
}
