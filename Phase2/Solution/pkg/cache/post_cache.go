package cache

import "github.com/cs-ut-ee/hw2-group-3/pkg/domain"

type PostPlantCache interface {
	SetPlant(key string, value *domain.Plant)
	GetPlant(key string) *domain.Plant
	SetListAvailablePlant(key string, value []*domain.Plant)
	GetListAvailablePlant(key string) []*domain.Plant
	SetIsAvailable(key string, value bool)
	GetIsAvailable(key string) (bool, error)
}

type PostPlantPriceCache interface {
	SetPlantPrice(key string, value *domain.PlantPrice)
	GetPlantPrice(key string) *domain.PlantPrice
}
