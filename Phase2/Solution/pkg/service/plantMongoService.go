package service

import (
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"
	"github.com/cs-ut-ee/hw2-group-3/pkg/repository"

	"time"
)

type PlantMongoService struct {
	PlantMongoRepo *repository.PlantMongoRepo
}

func NewPlantMongoService(plantMongoRepo *repository.PlantMongoRepo) *PlantMongoService {
	return &PlantMongoService{
		PlantMongoRepo: plantMongoRepo,
	}
}

func (service *PlantMongoService) GetAllPlants() ([]*domain.Plant, error) {
	plants, err := service.PlantMongoRepo.GetAllPlants()
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service *PlantMongoService) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error) {
	plantPrice, err := service.PlantMongoRepo.GetPlantPrice(Id, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return plantPrice, nil
}

func (service *PlantMongoService) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {
	availability, err := service.PlantMongoRepo.IsPlantAvailable(Id, startDate, endDate)
	if err != nil {
		return false, err
	}
	return availability, nil
}
