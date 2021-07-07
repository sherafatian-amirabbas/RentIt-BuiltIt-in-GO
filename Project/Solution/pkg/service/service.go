package service

import (
	"time"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
)

type IRepository interface {
	NewPlantHireRequest(plantName string, siteName string, supplierName string, requesterName string,
		startDate time.Time, endDate time.Time, totalHiringCost float64, status int64) (*domain.PlantHireRequest, error)

	GetPlantHireRequestById(plantRequestId int64) (*domain.PlantHireRequest, error)

	ModifyRequestBySiteEngineers(plantRequestId int64, plantName string, siteName string, supplierName string,
		requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error

	ModifyRequestByWorkEngineers(plantRequestId int64, plantName string, siteName string, supplierName string,
		requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error

	AcceptRequest(plantRequestId int64, regulator string, comment string) error
	RejectRequest(plantRequestId int64, regulator string, comment string) error
	GetCompleteOrders() ([]*domain.CompleteOrder, error)
	CancelPlantHireRequest(plantHireRequestId int64, requesterName string, comment string) error
	RequestExtension(plantHireRequestID int64, endTime time.Time) (*domain.PlantOrder, error)

	AcceptInvoice(invoiceId int64, requesterName string, comment string) (*domain.Invoice, error)
	RejectInvoice(invoiceId int64, requesterName string, comment string) (*domain.Invoice, error)
	GetPurchaseOrderByInvoiceId(invoiceId int64) (*domain.CompleteOrder, error)
	DeleteRequestsBySupplierName(supplierName string) ([]*domain.PlantHireRequest, error)

	GeneratePurchaseOrder(plantHireRequestId int64) (*domain.PlantOrder, error)
	CheckInvoice(invoice *domain.Invoice) error
	NewInvoice(invoice *domain.Invoice) (*domain.Invoice, error)
	GetInvoiceById(invoiceId int64) (*domain.Invoice, error)
}

type Service struct {
	Repository IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (service Service) NewPlantHireRequest(plantName string, siteName string, supplierName string,
	requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64, status int64) (*domain.PlantHireRequest, error) {

	return service.Repository.NewPlantHireRequest(plantName, siteName, supplierName, requesterName,
		startDate, endDate, totalHiringCost, status)
}

func (service Service) GetPlantHireRequestById(plantRequestId int64) (*domain.PlantHireRequest, error) {
	return service.Repository.GetPlantHireRequestById(plantRequestId)
}

func (service Service) ModifyRequestBySiteEngineers(plantRequestId int64, plantName string, siteName string,
	supplierName string, requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error {

	return service.Repository.ModifyRequestBySiteEngineers(plantRequestId, plantName, siteName, supplierName,
		requesterName, startDate, endDate, totalHiringCost)
}

func (service Service) ModifyRequestByWorkEngineers(plantRequestId int64, plantName string, siteName string,
	supplierName string, requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error {

	return service.Repository.ModifyRequestByWorkEngineers(plantRequestId, plantName, siteName, supplierName,
		requesterName, startDate, endDate, totalHiringCost)
}

func (service Service) AcceptRequest(plantRequestId int64, regulator string, comment string) error {
	return service.Repository.AcceptRequest(plantRequestId, regulator, comment)
}

func (service Service) RejectRequest(plantRequestId int64, regulator string, comment string) error {
	return service.Repository.RejectRequest(plantRequestId, regulator, comment)
}

func (service Service) GetCompleteOrders() ([]*domain.CompleteOrder, error) {
	return service.Repository.GetCompleteOrders()
}

func (service Service) CancelPlantHireRequest(plantHireRequestId int64, requesterName string, comment string) error {
	return service.Repository.CancelPlantHireRequest(plantHireRequestId, requesterName, comment)
}

func (service Service) GeneratePurchaseOrder(plantHireRequestId int64) (*domain.PlantOrder, error) {
	return service.Repository.GeneratePurchaseOrder(plantHireRequestId)
}

func (service Service) AcceptInvoice(invoiceId int64, requesterName string, comment string) (*domain.Invoice, error) {
	return service.Repository.AcceptInvoice(invoiceId, requesterName, comment)
}

func (service Service) RejectInvoice(invoiceId int64, requesterName string, comment string) (*domain.Invoice, error) {
	return service.Repository.RejectInvoice(invoiceId, requesterName, comment)
}

func (service Service) GetPurchaseOrderByInvoiceId(invoiceId int64) (*domain.CompleteOrder, error) {
	return service.Repository.GetPurchaseOrderByInvoiceId(invoiceId)
}

func (service Service) DeleteRequestsBySupplierName(supplierName string) ([]*domain.PlantHireRequest, error) {
	return service.Repository.DeleteRequestsBySupplierName(supplierName)
}

func (service Service) RequestExtension(plantHireRequestID int64, endTime time.Time) (*domain.PlantOrder, error) {
	return service.Repository.RequestExtension(plantHireRequestID, endTime)
}

func (service Service) CheckInvoice(invoice *domain.Invoice) error {
	return service.Repository.CheckInvoice(invoice)
}

func (service Service) GetInvoice(invoiceId int64) (*domain.Invoice, error) {
	return service.Repository.GetInvoiceById(invoiceId)
}
