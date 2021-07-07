package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"
	"github.com/go-redis/redis/v8"
)

//credit : Pragmatic Reviews on Youtube (www.youtube.com/watch?v=x5GGLrTuQCA&t)

//for creating redis client struct
type redis_cache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisPlantCache(host string, db int, exp time.Duration) PostPlantCache {
	return &redis_cache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func NewRedisPlantPriceCache(host string, db int, exp time.Duration) PostPlantPriceCache {
	return &redis_cache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redis_cache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

//Set Plant info to redis db
func (cache *redis_cache) SetPlant(key string, value *domain.Plant) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

//Get Plant from Redis db
func (cache *redis_cache) GetPlant(key string) *domain.Plant {
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()

	if err != nil {
		return nil
	}

	plant := domain.Plant{}
	err = json.Unmarshal([]byte(val), &plant)
	if err != nil {
		panic(err)
	}

	return &plant
}

//Set list of available plant
func (cache *redis_cache) SetListAvailablePlant(key string, value []*domain.Plant) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

//Get list of available plant
func (cache *redis_cache) GetListAvailablePlant(key string) []*domain.Plant {
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()

	if err != nil {
		return nil
	}

	plants := []*domain.Plant{}
	err = json.Unmarshal([]byte(val), &plants)
	if err != nil {
		panic(err)
	}

	return plants
}

//Set availability of a plant based on key
func (cache *redis_cache) SetIsAvailable(key string, value bool) {
	client := cache.getClient()
	client.Set(context.Background(), key, value, cache.expires*time.Second)
}

//Get plant is available or not
func (cache *redis_cache) GetIsAvailable(key string) (bool, error) {
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()

	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(val)

	if err != nil {
		return false, err
	}

	return b, nil
}

//Set Plant price info to redis db
func (cache *redis_cache) SetPlantPrice(key string, value *domain.PlantPrice) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

//Get Plant Price from Redis db
func (cache *redis_cache) GetPlantPrice(key string) *domain.PlantPrice {
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()

	if err != nil {
		return nil
	}

	plant := domain.PlantPrice{}
	err = json.Unmarshal([]byte(val), &plant)
	if err != nil {
		panic(err)
	}

	return &plant
}
