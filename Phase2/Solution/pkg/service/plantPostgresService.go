package service

import (
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"
	"github.com/cs-ut-ee/hw2-group-3/pkg/repository"

	"time"
)

type PlantPostgresService struct {
	PlantPostgresRepo *repository.PlantPostgresRepo
}

func NewPlantPostgresService(plantPostgresRepo *repository.PlantPostgresRepo) *PlantPostgresService {
	return &PlantPostgresService{
		PlantPostgresRepo: plantPostgresRepo,
	}
}

func (service *PlantPostgresService) GetAllPlants() ([]*domain.Plant, error) {
	plants, err := service.PlantPostgresRepo.GetAllPlants()
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service *PlantPostgresService) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error) {
	plantPrice, err := service.PlantPostgresRepo.GetPlantPrice(Id, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return plantPrice, nil
}

func (service *PlantPostgresService) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {
	availability, err := service.PlantPostgresRepo.IsPlantAvailable(Id, startDate, endDate)
	if err != nil {
		return false, err
	}
	return availability, nil
}
