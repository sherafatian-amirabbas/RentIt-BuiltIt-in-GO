package cachetest

import (
	"os"
	"testing"

	"github.com/cs-ut-ee/hw2-group-3/pkg/cache"
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"
)

func CreatePlantCacheClient(db int) cache.PostPlantCache {
	host, success := os.LookupEnv("redisUrl")
	if !success {
		panic("Environment variable 'redisUrl' is not defined")
	}

	return cache.NewRedisPlantCache(host, db, 30)
}

func CreatePlantPriceCacheClient(db int) cache.PostPlantPriceCache {
	host, success := os.LookupEnv("redisUrl")
	if !success {
		panic("Environment variable 'redisUrl' is not defined")
	}
	return cache.NewRedisPlantPriceCache(host, db, 30)
}

func Plants() []*domain.Plant {
	plant1 := domain.Plant{Id: 1, Name: "Plant1", Description: "Test Plant1", PricePerDay: 100}
	plant2 := domain.Plant{Id: 2, Name: "Plant2", Description: "Test Plant2", PricePerDay: 200}
	plant3 := domain.Plant{Id: 3, Name: "Plant3", Description: "Test Plant3", PricePerDay: 300}

	return []*domain.Plant{&plant1, &plant2, &plant3}
}

func Plant() *domain.Plant {
	plant1 := domain.Plant{Id: 4, Name: "Plant4", Description: "Test Plant4", PricePerDay: 400}
	return &plant1
}

func TestGetAllPlants(t *testing.T) {
	client := CreatePlantCacheClient(2)
	plants := Plants()
	plantsByGet := client.GetListAvailablePlant("list")

	if plantsByGet != nil {
		t.Fatalf("Plants should not be available in cache!")
	}

	client.SetListAvailablePlant("list", plants)

	plantsByGet = client.GetListAvailablePlant("list")

	if plantsByGet == nil {
		t.Fatalf("Plants should be retrieved from cache!")
	}
}

func TestIsPlantAvailable(t *testing.T) {
	client := CreatePlantCacheClient(1)
	_, err := client.GetIsAvailable("plant1")

	if err == nil {
		t.Fatalf("Plant should not be available in cache!")
	}
	client.SetIsAvailable("plant1", true)
	_, err = client.GetIsAvailable("plant1")
	if err != nil {
		t.Error("Plant should be available in cache!", err)
	}
}

func TestGetPlant(t *testing.T) {
	client := CreatePlantCacheClient(3)
	res := client.GetPlant("plant1")

	if res != nil {
		t.Fatalf("Plant should not be available in cache!")
	}

	plant := Plant()
	client.SetPlant("plant1", plant)

	res = client.GetPlant("plant1")

	if res == nil {
		t.Fatalf("Plant should be available in cache!")
	}
}

func PlantPrice() *domain.PlantPrice {
	plant1 := domain.PlantPrice{PlantId: 1, StartDate: "20/4/2021", EndDate: "25/4/2021", PricePerDuration: 100}
	return &plant1
}

func TestGetPlantPrice(t *testing.T) {
	client := CreatePlantPriceCacheClient(4)
	res := client.GetPlantPrice("price1")

	if res != nil {
		t.Fatalf("Plant Price should not be available in cache!")
	}

	client.SetPlantPrice("price1", PlantPrice())

	if res != nil {
		t.Fatalf("Plant Price should be available in cache!")
	}
}
