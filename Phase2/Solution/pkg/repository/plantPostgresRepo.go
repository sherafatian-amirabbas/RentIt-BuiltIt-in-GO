package repository

import (
	"github.com/cs-ut-ee/hw2-group-3/pkg/domain"

	"context"
	"database/sql"
	"fmt"
	"math"
	"time"
)

// PlantPostgresRepo Makes possible to access DB through methods
type PlantPostgresRepo struct {
	DB *sql.DB
}

// ------------------------------------------------------------- DB Schema

type tables struct {
	Plants      string
	PlantOrders string
}

var tableNames = &tables{
	Plants:      "plants",
	PlantOrders: "plantorders",
}

type plantsColumns struct {
	Id          string
	Name        string
	Description string
	PricePerDay string
}

var plantsColumnsName = &plantsColumns{
	Id:          "Id",
	Name:        "Name",
	Description: "Description",
	PricePerDay: "PricePerDay",
}

type plantOrdersColumns struct {
	Id        string
	PlantId   string
	StartDate string
	EndDate   string
}

var plantOrdersColumnsName = &plantOrdersColumns{
	Id:        "Id",
	PlantId:   "PlantId",
	StartDate: "StartDate",
	EndDate:   "EndDate",
}

// -------------------------------------------------------------

// NewPlantPostgresRepo creates the repository by receiving the a Postgres DB
func NewPlantPostgresRepo(db *sql.DB) *PlantPostgresRepo {
	return &PlantPostgresRepo{
		DB: db,
	}
}

// GetAll returns all the plants
func (repo *PlantPostgresRepo) GetAllPlants() ([]*domain.Plant, error) {

	query := fmt.Sprintf("SELECT * FROM %s", tableNames.Plants)
	rows, err := repo.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("GetAll: error querying plants, err: %v", err)
	}

	plants := []*domain.Plant{}
	for rows.Next() {
		plant := &domain.Plant{}
		err := rows.Scan(&plant.Id, &plant.Name, &plant.Description, &plant.PricePerDay)
		if err != nil {
			return nil, fmt.Errorf("GetAll: error scaning query, err: %v", err)
		}
		plants = append(plants, plant)
	}

	// close rows to avoid memory leak
	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("GetAll: could not close rows, err %v", err)
	}

	return plants, nil
}

func (repo *PlantPostgresRepo) GetPlantById(Id int64) (*domain.Plant, error) {
	query := fmt.Sprintf("SELECT * FROM %s where %s=$1", tableNames.Plants, plantsColumnsName.Id)
	row := repo.DB.QueryRow(query, Id)

	plant := domain.Plant{}
	err := row.Scan(&plant.Id, &plant.Name, &plant.Description, &plant.PricePerDay)
	if err != nil {
		return nil, fmt.Errorf("GetPlantById: the plant was not found, err: %v", err)
	}

	return &plant, nil
}

func (repo *PlantPostgresRepo) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (*domain.PlantPrice, error) {

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
}

// GetPlantOrdersById returns all the plants
func (repo *PlantPostgresRepo) GetPlantOrdersByPlantId(Id int64) ([]*domain.PlantOrder, error) {

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", tableNames.PlantOrders, plantOrdersColumnsName.PlantId)
	rows, err := repo.DB.QueryContext(context.Background(), query, Id)
	if err != nil {
		return nil, fmt.Errorf("GetPlantOrdersByPlantId: error querying plantOrders, err: %v", err)
	}

	plantOrders := []*domain.PlantOrder{}
	for rows.Next() {
		plantOrder := &domain.PlantOrder{}
		err := rows.Scan(&plantOrder.Id, &plantOrder.PlantId, &plantOrder.StartDate, &plantOrder.EndDate)
		if err != nil {
			return nil, fmt.Errorf("GetPlantOrdersByPlantId: error scaning query, err: %v", err)
		}
		plantOrders = append(plantOrders, plantOrder)
	}

	// close rows to avoid memory leak
	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("GetPlantOrdersByPlantId: could not close rows, err %v", err)
	}

	return plantOrders, nil
}

func (repo *PlantPostgresRepo) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {

	currentOrders, err := repo.GetPlantOrdersByPlantId(Id)
	if err != nil {
		return false, err
	}

	isAvailable := true
	for _, order := range currentOrders {

		isStartDateInTheRange := (startDate.Equal(order.StartDate) || startDate.After(order.StartDate)) && startDate.Before(order.EndDate)
		isEndDateInTheRange := endDate.After(order.StartDate) && endDate.Before(order.EndDate)
		if isStartDateInTheRange || isEndDateInTheRange {
			isAvailable = false
			break
		}
	}

	return isAvailable, nil
}
