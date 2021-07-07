package dto

import (
	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
)

// customer will hold the information about the customer
type Customer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetCustomerDTOList(customers []*domain.Customer) []*Customer {

	var dtoList = []*Customer{}
	for _, customer := range customers {
		dto := GetCustomerDTO(customer)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetCustomerDTO(customer *domain.Customer) *Customer {

	dto := &Customer{
		Id:   customer.ID,
		Name: customer.Name,
	}

	return dto
}
