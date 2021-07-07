package service

import (
	"strconv"

	"github.com/cs-ut-ee/hw2-group-3/pkg/cache"
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"

	"time"
)

type IService interface {
	GetAllPlants() ([]*domain.Plant, error)
	GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error)
	IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error)
}

type Service struct {
	PostgresService *PlantPostgresService
	MongoService    *PlantMongoService
}

func NewService(postgresService *PlantPostgresService, mongoService *PlantMongoService) *Service {
	return &Service{
		PostgresService: postgresService,
		MongoService:    mongoService,
	}
}

func (service *Service) GetAllPlants() ([]*domain.Plant, error) {

	redis_client := cache.NewRedisPlantCache("localhost:6379", 1, 30) //using database 1 and timeout 30seconds for plant list
	cache_list := redis_client.GetListAvailablePlant("available")

	if cache_list != nil {
		return cache_list, nil
	}

	plants_postgres, err := service.PostgresService.GetAllPlants()
	if err != nil {
		return nil, err
	}

	plants_mongo, err := service.MongoService.GetAllPlants()
	if err != nil {
		return nil, err
	}

	redis_client.SetListAvailablePlant("available", append(plants_postgres, plants_mongo...)) //for list of plants "available" is the key
	return append(plants_postgres, plants_mongo...), nil
}

func (service *Service) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error) {
	redis_client := cache.NewRedisPlantPriceCache("localhost:6379", 2, 30)
	key := strconv.FormatInt(Id, 16) + startDate.String() + endDate.String()
	res := redis_client.GetPlantPrice(key)

	if res != nil {
		return res, nil
	}

	plants_postgres, err := service.PostgresService.GetPlantPrice(Id, startDate, endDate)
	if err != nil {
		return nil, err
	}

	if plants_postgres != nil {
		redis_client.SetPlantPrice(key, plants_postgres)
		return plants_postgres, nil
	}

	plants_mongo, err := service.MongoService.GetPlantPrice(Id, startDate, endDate)
	if err != nil {
		return nil, err
	}

	redis_client.SetPlantPrice(key, plants_mongo)

	return plants_mongo, nil
}

func (service *Service) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {
	redis_client := cache.NewRedisPlantCache("localhost:6379", 2, 30)
	key := strconv.FormatInt(Id, 16) + startDate.String() + endDate.String()
	res, err := redis_client.GetIsAvailable(key)

	if err == nil {
		return res, nil
	}

	available, err := service.PostgresService.IsPlantAvailable(Id, startDate, endDate)
	if err != nil {
		return false, err
	}

	if !available {
		return false, nil // to prevent unnecessary queries to MongoDB since the plant was already found
	}

	available, err = service.MongoService.IsPlantAvailable(Id, startDate, endDate)
	if err != nil {
		return false, err
	}

	redis_client.SetIsAvailable(key, available)
	return available, nil
}
