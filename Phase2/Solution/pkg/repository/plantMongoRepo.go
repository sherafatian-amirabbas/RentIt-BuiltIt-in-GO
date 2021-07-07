package repository

import (
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"fmt"
	"math"
)

// PlantMongoRepo Makes possible to access DB through methods
type PlantMongoRepo struct {
	db *mongo.Client
}

// ------------------------------------------------------------- private members

// Tables holds the table names of the DB
type Collections struct {
	Plants string
	PlantOrders string
}

var collections = &Collections{
	Plants: "plants",
	PlantOrders: "plant_orders",
}

// ------------------------------------------------------------- private members

// NewPlantMongoRepo creates the repository by receiving the a Postgres DB
func NewPlantMongoRepo(db *mongo.Client) *PlantMongoRepo {
	return &PlantMongoRepo{
		db: db,
	}
}

// GetAll returns all the plants
func (repo *PlantMongoRepo) GetAllPlants() ([]*domain.Plant, error) {
	collection := repo.db.Database("local").Collection(collections.Plants)

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting plants, err: %v", err)
	}
	plants := []*domain.Plant{}
	for cursor.Next(context.Background()) {
		b := &domain.Plant{}
		err := cursor.Decode(&b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		plants = append(plants, b)
	}

	err = cursor.Close(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting plants, err: %v", err)
	}

	return plants, nil
}

func (repo *PlantMongoRepo) GetPlantById(Id int64) (*domain.Plant, error) {
	collection := repo.db.Database("local").Collection(collections.Plants)

	var plant *domain.Plant
	err := collection.FindOne(context.Background(), bson.M{"_id": Id}).Decode(&plant)
	
	if err != nil {
		return nil, fmt.Errorf("given plant not found")
	}

	return plant, nil
}

func (repo *PlantMongoRepo) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error) {

	if endDate.After(startDate) {
		plant, err := repo.GetPlantById(Id)

		if err != nil {
			return nil, err
		}

		days := math.Ceil(endDate.Sub(startDate).Hours() / 24)
		plantPrice := &domain.PlantPrice{
			PlantId:          plant.Id,
			StartDate:        startDate.Format("2006-01-02"),
			EndDate:          endDate.Format("2006-01-02"),
			PricePerDuration: (days * plant.PricePerDay),
		}

		return plantPrice, nil
	} else {

		return nil, fmt.Errorf("GetPlantPrice: end date should be greate than start date")
	}

	return nil, nil
}

func (repo *PlantMongoRepo) GetPlantOrdersByPlantId(Id int64) ([]*domain.PlantOrder, error) {
	collection := repo.db.Database("local").Collection(collections.PlantOrders)
	cursor, err := collection.Find(context.Background(), bson.M{"PlantId": Id})
	if err != nil {
		return nil, fmt.Errorf("error getting plants, err: %v", err)
	}

	plantOrders := []*domain.PlantOrder{}
	for cursor.Next(context.Background()) {
		b := &domain.PlantOrder{}
		err := cursor.Decode(&b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		plantOrders = append(plantOrders, b)
	}

	err = cursor.Close(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting plants, err: %v", err)
	}

	return plantOrders, nil
}

func (repo *PlantMongoRepo) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {
	currentOrders, err := repo.GetPlantOrdersByPlantId(Id)

	if err != nil {
		return false, err
	}

	isAvailable := true
	for _, order := range currentOrders {

		isStartDateInTheRange := (startDate.Equal(order.StartDate) || startDate.After(order.StartDate)) && startDate.Before(order.EndDate)
		isEndDateInTheRange := (endDate.Equal(order.StartDate) || endDate.After(order.StartDate)) && endDate.Before(order.EndDate)
		if isStartDateInTheRange || isEndDateInTheRange {
			isAvailable = false
			break
		}
	}

	return isAvailable, nil
}
