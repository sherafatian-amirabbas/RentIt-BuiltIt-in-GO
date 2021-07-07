package dto

import (
	"time"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
)

// customer will hold the information about the customer
type PlanRequestDTO struct {
	ID                  int64     `json:"id"`
	PlantName           string    `json:"plantName"`
	SiteName            string    `json:"siteName"`
	SupplierName        string    `json:"supplierName"`
	RequesterName       string    `json:"requesterName"`
	StartDate           time.Time `json:"startDate"`
	EndDate             time.Time `json:"endDate"`
	TotalHiringCost     float64   `json:"totalHiringCost"`
	Regulator           string    `json:"regulator"`
	WorkEngineerComment string    `json:"workEngineerComment"`
	StatusCode          int64     `json:"statusCode"`
	StatusDesc          string    `json:"statusDesc"`
}

func GetPlantHireRequestDTOList(plantHireRequests []*domain.PlantHireRequest) []*PlanRequestDTO {

	var dtoList = []*PlanRequestDTO{}
	for _, request := range plantHireRequests {
		dto := GetPlantHireRequestDTO(request)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetPlantHireRequestDTO(plantHireRequest *domain.PlantHireRequest) *PlanRequestDTO {

	statusDesc := domain.GetRequestStatusDescription(plantHireRequest.Status)
	dto := &PlanRequestDTO{
		ID:                  plantHireRequest.ID,
		PlantName:           plantHireRequest.PlantName,
		SiteName:            plantHireRequest.SiteName,
		SupplierName:        plantHireRequest.SupplierName,
		RequesterName:       plantHireRequest.RequesterName,
		StartDate:           plantHireRequest.StartDate,
		EndDate:             plantHireRequest.EndDate,
		TotalHiringCost:     plantHireRequest.TotalHiringCost,
		Regulator:           plantHireRequest.Regulator,
		WorkEngineerComment: plantHireRequest.WorkEngineerComment,
		StatusCode:          plantHireRequest.Status,
		StatusDesc:          statusDesc,
	}

	return dto
}

func GetPlantHireRequestFromDTO(dto *PlanRequestDTO) *domain.PlantHireRequest {
	PlantRequest := &domain.PlantHireRequest{
		ID:                  dto.ID,
		PlantName:           dto.PlantName,
		SiteName:            dto.SiteName,
		SupplierName:        dto.SupplierName,
		RequesterName:       dto.RequesterName,
		StartDate:           dto.StartDate,
		EndDate:             dto.EndDate,
		TotalHiringCost:     dto.TotalHiringCost,
		Regulator:           dto.Regulator,
		WorkEngineerComment: dto.WorkEngineerComment,
		Status:              dto.StatusCode,
	}

	return PlantRequest
}
