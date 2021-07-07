package service

import (
	"time"

	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
)

type IRepository interface {
	GetAllPlants() ([]*domain.Plant, error)
	GetPlantPrice(plantID int64, startDate time.Time, endDate time.Time) (float64, error)
	GetAllCustomers() ([]*domain.Customer, error)
	GetPlantOrdersByStartDate(startDate time.Time) ([]*domain.PlantOrder, error)
	GetPagedPlantOrdersByStartDate(startDate time.Time, pageNumber int, pageSize int) ([]*domain.PlantOrder, error)
	NewPlantOrder(plantID int64, customerID int64, startDate time.Time, endDate time.Time) error
	IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error)
	UpdatePlantOrder(plantOrderID int64, plantID int64, startDate time.Time, endDate time.Time) error
	CancelPlantOrder(plantOrderID int64) (bool, error)
	RejectPlantByCustomer(plantOrderID int64) (bool, error)
	RentalPeriodExpired(plantOrderID int64) (bool, error)
	GetUnPaidInvoicesFor(customerID int64) ([]*domain.Invoice, error)
	CreateRemittance(plantOrderID int64, referenceNumber string) (bool, error)
	AcceptRemittance(remittanceId int64) (bool, error)
}

type Service struct {
	Repository IRepository
}

func NewService(repo IRepository) Service {
	return Service{
		Repository: repo,
	}
}

func (service Service) GetAllCustomers() ([]*domain.Customer, error) {

	return service.Repository.GetAllCustomers()
}

func (service Service) GetAllPlants() ([]*domain.Plant, error) {
	return service.Repository.GetAllPlants()
}

func (service Service) GetPlantPrice(Id int64, startDate time.Time, endDate time.Time) (float64, error) {
	return service.Repository.GetPlantPrice(Id, startDate, endDate)
}

func (service Service) GetPlantOrdersByStartDate(startDate time.Time) ([]*domain.PlantOrder, error) {

	return service.Repository.GetPlantOrdersByStartDate(startDate)
}

func (service Service) GetPagedPlantOrdersByStartDate(startDate time.Time, pageNumber int, pageSize int) ([]*domain.PlantOrder, error) {

	return service.Repository.GetPagedPlantOrdersByStartDate(startDate, pageNumber, pageSize)
}

func (service Service) NewPlantOrder(plantID int64, customerID int64,
	startDate time.Time, endDate time.Time) error {

	return service.Repository.NewPlantOrder(plantID, customerID, startDate, endDate)
}

func (service Service) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {

	return service.Repository.IsPlantAvailable(Id, startDate, endDate)
}

func (service Service) UpdatePlantOrder(plantOrderID int64, plantID int64, startDate time.Time, endDate time.Time) error {

	return service.Repository.UpdatePlantOrder(plantOrderID, plantID, startDate, endDate)
}

//PS7
func (service Service) CancelPlantOrder(plantOrderID int64) (bool, error) {
	return service.Repository.CancelPlantOrder(plantOrderID)
}

//PS8
func (service Service) RejectPlantByCustomer(plantOrderID int64) (bool, error) {
	return service.Repository.RejectPlantByCustomer(plantOrderID)
}

//PS9
func (service Service) RentalPeriodExpired(plantOrderID int64) (bool, error) {
	return service.Repository.RentalPeriodExpired(plantOrderID)
}

//PS12
func (service Service) CreateRemittance(plantOrderID int64, referenceNumber string) (bool, error) {
	return service.Repository.CreateRemittance(plantOrderID, referenceNumber)
}

func (service Service) AcceptRemittance(remittanceId int64) (bool, error) {
	return service.Repository.AcceptRemittance(remittanceId)
}
