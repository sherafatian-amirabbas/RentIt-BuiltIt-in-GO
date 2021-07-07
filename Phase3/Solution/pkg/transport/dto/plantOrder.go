package dto

import (
	"time"

	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
)

// PlantOrder will hold the information about the order
type PlantOrder struct {
	Id         int64     `json:"id"`
	PlantId    int64     `json:"plantId"`
	CustomerId int64     `json:"customerId"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     int64     `json:"status"` // 0: accepted, 1: rejected, 2: canceled, 3: returned
	InvoiceId  int64     `json:"invoiceId"`
}

func GetPlantOrderDTOList(orders []*domain.PlantOrder) []*PlantOrder {

	var dtoList = []*PlantOrder{}
	for _, order := range orders {
		dto := GetPlantOrderDTO(order)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetPlantOrderDTO(order *domain.PlantOrder) *PlantOrder {

	dto := &PlantOrder{
		Id:         order.ID,
		PlantId:    order.PlantID,
		CustomerId: order.CustomerID,
		StartDate:  order.StartDate,
		EndDate:    order.EndDate,
		Status:     order.Status,
		InvoiceId:  order.InvoiceID,
	}

	return dto
}
