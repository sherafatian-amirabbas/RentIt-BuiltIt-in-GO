package dto

import (
	"github.com/cs-ut-ee/project-group-3/pkg/domain"
)

// customer will hold the information about the customer
type PlantOrderDTO struct {
	Id         int64  `json:"id"`
	RequestID  int64  `json:"requestID"`
	StatusCode int64  `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}

func GetPlantDTOList(plantOrders []*domain.PlantOrder) []*PlantOrderDTO {

	var dtoList = []*PlantOrderDTO{}
	for _, order := range plantOrders {
		dto := GetPlantOrderDTO(order)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetPlantOrderDTO(order *domain.PlantOrder) *PlantOrderDTO {

	statusDesc := domain.GetOrderStatusDescription(order.Status)
	dto := &PlantOrderDTO{
		Id:         order.ID,
		RequestID:  order.RequestID,
		StatusCode: order.Status,
		StatusDesc: statusDesc,
	}

	return dto
}
