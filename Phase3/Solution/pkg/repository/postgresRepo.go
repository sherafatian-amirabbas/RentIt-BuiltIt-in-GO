package repository

import (
	"time"

	"fmt"

	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
	"gorm.io/gorm"
)

// postgresRepo Makes possible to access DB
type PostgresRepo struct {
	Database *gorm.DB
}

// NewPostgresRepo creates the repository by receiving the a DB
func NewPostgresRepo(db *gorm.DB) PostgresRepo {
	return PostgresRepo{
		Database: db,
	}
}

// GetAllCustomers returns all the customers
func (repo *PostgresRepo) GetAllCustomers() ([]*domain.Customer, error) {

	var customers []*domain.Customer
	result := Model(repo.Database).Table(Tables.customers).Find(&customers)
	if result.Error != nil {
		return nil, fmt.Errorf("GetAllCustomers: error querying Customers, err: %v", result.Error)
	}

	return customers, nil
}

// GetAllPlants returns all the customers
func (repo *PostgresRepo) GetAllPlants() ([]*domain.Plant, error) {
	var plants []*domain.Plant
	result := Model(repo.Database).Table(Tables.plants).Find(&plants)

	if result.Error != nil {
		return nil, fmt.Errorf("GetAllPlants: error querying Plants, err: %v", result.Error)
	}

	return plants, nil
}

func (repo *PostgresRepo) GetPlantPrice(plantID int64, startDate time.Time,
	endDate time.Time) (float64, error) {

	var plant *domain.Plant

	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plants).Where(whereClause, plantID).Take(&plant)
	if result.Error != nil {
		return 0.0, fmt.Errorf("GetPlantPrice: error querying Plants, err: %v", result.Error)
	}

	var daysPassed = endDate.Sub(startDate).Hours() / 24.0
	return plant.PricePerDay * daysPassed, nil
}

// UpdatePlantOrder updates a record from the table “PlantOrders” if its status is not Returned or Canceled
func (repo *PostgresRepo) UpdatePlantOrder(plantOrderID int64, plantID int64, startDate time.Time,
	endDate time.Time) error {

	// this always apply changes to the record then checks for an overlap and then change the
	// Status again to Accepted or Rejected.

	var order *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantOrderID).Take(&order)
	if result.Error != nil {

		return fmt.Errorf("UpdatePlantOrder: error updating order, err: %v", result.Error)
	}

	if order.Status == 2 || order.Status == 3 { // if it's canceled or returned, shouldn't be allowed to update
		return fmt.Errorf("UpdatePlantOrder: orders with status Canceled or Returned cannot be updated")
	}

	order.PlantID = plantID
	order.StartDate = startDate
	order.EndDate = endDate
	order.Status = 1 // rejected

	// If it has an overlap should be Rejected, if not it’s Accepted

	currentOrders, err := repo.getPlantOrdersByPlantID(plantID)
	if err != nil {
		return err
	}

	var orders = []*domain.PlantOrder{}
	for _, order := range currentOrders {
		if order.ID != plantOrderID {
			orders = append(orders, order)
		}
	}

	isAvailable := repo.isThereAnyOverlapAmongOrders(startDate, endDate, orders)
	if isAvailable {
		order.Status = 0 // accepted
	}

	result = Model(repo.Database).Table(Tables.plantOrders).Save(&order)
	if result.Error != nil {

		return fmt.Errorf("UpdatePlantOrder: error saving the changes, err: %v", result.Error)
	}

	return nil
}

// NewPlantOrder creates a new order
func (repo *PostgresRepo) NewPlantOrder(plantID int64, customerID int64, startDate time.Time,
	endDate time.Time) error {

	id := repo.getMaxPlantOrderID() + 1
	var status int64 = 1 // rejected

	// If it has an overlap should be Rejected, if not it’s Accepted
	isAvailable, err := repo.IsPlantAvailable(plantID, startDate, endDate)
	if err != nil {
		return err
	}

	if isAvailable {
		status = 0 // accepted
	}

	order := &domain.PlantOrder{
		ID:         id,
		PlantID:    plantID,
		CustomerID: customerID,
		StartDate:  startDate,
		EndDate:    endDate,
		Status:     status,
	}

	result := Model(repo.Database).Table(Tables.plantOrders).Create(order)
	if result.Error != nil {

		return fmt.Errorf("NewPlantOrder: error creating order, err: %v", result.Error)
	}

	return nil
}

// IsPlantAvailable returns true if the plant can be rented at that specific duration
func (repo *PostgresRepo) IsPlantAvailable(Id int64, startDate time.Time, endDate time.Time) (bool, error) {

	currentOrders, err := repo.getPlantOrdersByPlantID(Id)
	if err != nil {
		return false, err
	}

	isAvailable := repo.isThereAnyOverlapAmongOrders(startDate, endDate, currentOrders)

	return isAvailable, nil
}

// GetPlantOrders returns a list of “PlantOrders” by receiving a date to be compared with “StartDate”.
func (repo *PostgresRepo) GetPlantOrdersByStartDate(startDate time.Time) ([]*domain.PlantOrder, error) {

	var orders []*domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.StartDate)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, startDate.Format("2006-01-02")).Find(&orders)
	if result.Error != nil {
		return nil, fmt.Errorf("GetPlantOrdersByStartDate: error querying plantOrders, err: %v", result.Error)
	}

	return orders, nil
}

// GetPagedPlantOrdersByStartDate returns a list of “PlantOrders” by receiving a date to be compared with “StartDate”. This method is supporting pagination
func (repo *PostgresRepo) GetPagedPlantOrdersByStartDate(startDate time.Time, pageNumber int, pageSize int) ([]*domain.PlantOrder, error) {

	var orders []*domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.StartDate)
	list := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, startDate.Format("2006-01-02"))
	result := list.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&orders)

	if result.Error != nil {
		return nil, fmt.Errorf("GetPlantOrdersByStartDate: error querying plantOrders, err: %v", result.Error)
	}

	return orders, nil
}

//Cancel PlantOrder
func (repo *PostgresRepo) CancelPlantOrder(plantOrderID int64) (bool, error) {
	//Get the plantOrder
	var plantOrder *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantOrderID).Find(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("CancelPlantOrder: error querying plantOrder, err: %v", result.Error)
	}

	startDate := plantOrder.StartDate
	curTime := time.Now()
	var cancelRequest *domain.CancellationRequest

	if curTime.Before(startDate) == true && plantOrder.Status == 0 {
		//Update the plantorder status with 'canlcelled'
		plantOrder.Status = 2
		result = Model(repo.Database).Table(Tables.plantOrders).Save(&plantOrder)

		if result.Error != nil {
			return false, fmt.Errorf("CancelPlantOrder: error saving the changes, err: %v", result.Error)
		}

		cancelRequest.PlantOrderId = plantOrder.ID
		cancelRequest.SubmissionDate = curTime
		cancelRequest.Status = 2

		result = Model(repo.Database).Table(Tables.cancellationRequests).Save(&cancelRequest)

		return true, nil
	}

	return false, nil
}

//PS8
func (repo *PostgresRepo) RejectPlantByCustomer(plantOrderID int64) (bool, error) {
	//Get the plantOrder
	var plantOrder *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantOrderID).Find(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("RejectPlantByCustomer: error querying plantOrder, err: %v", result.Error)
	}

	//Update the plantorder status with 'rejected'
	plantOrder.Status = 1
	result = Model(repo.Database).Table(Tables.plantOrders).Save(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("RejectPlantByCustomer: error saving the changes, err: %v", result.Error)
	}

	return true, nil
}

//PS9
func (repo *PostgresRepo) RentalPeriodExpired(plantOrderID int64) (bool, error) {
	//Get the plantOrder
	var plantOrder *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantOrderID).Find(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("RentalPeriodExpired: error querying plantOrder, err: %v", result.Error)
	}

	//Update the plantorder status with 'returned'
	plantOrder.Status = 3
	result = Model(repo.Database).Table(Tables.plantOrders).Save(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("RentalPeriodExpired: error saving the changes, err: %v", result.Error)
	}

	repo.generateReturnedPlantInvoice(plantOrder)

	return true, nil
}

// PS 12
func (repo *PostgresRepo) CreateRemittance(plantOrderID int64, referenceNumber string) (bool, error) {
	//Get the plantOrder
	var plantOrder *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantOrderID).Find(&plantOrder)

	if result.Error != nil {
		return false, fmt.Errorf("CreateRemittance: error querying plantOrder, err: %v", result.Error)
	}

	// Get the invoice
	var invoice *domain.Invoice
	whereClause = fmt.Sprintf("%s = ?", "id")
	result = Model(repo.Database).Table(Tables.invoices).Where(whereClause, plantOrder.InvoiceID).Find(&invoice)

	if result.Error != nil {
		return false, fmt.Errorf("CreateRemittance: error querying invoice, err: %v", result.Error)
	}

	var remittanceId = repo.getMaxRemittanceID() + 1
	remittance := &domain.Remittance{
		ID:              remittanceId,
		InvoiceId:       invoice.ID,
		ReferenceNumber: referenceNumber,
		Status:          0,
	}

	result = Model(repo.Database).Table(Tables.remittances).Create(remittance)

	if result.Error != nil {
		return false, fmt.Errorf("CreateRemittance: error creating remittance, err: %v", result.Error)
	}

	invoice.RemittanceID = remittanceId
	result = Model(repo.Database).Table(Tables.invoices).Save(&invoice)

	return true, nil
}

func (repo *PostgresRepo) AcceptRemittance(remittanceId int64) (bool, error) {
	//Get the plantOrder
	var remittance *domain.Remittance
	whereClause := fmt.Sprintf("%s = ?", "id")
	result := Model(repo.Database).Table(Tables.remittances).Where(whereClause, remittanceId).Find(&remittance)

	if result.Error != nil {
		return false, fmt.Errorf("AcceptRemittance: error querying remittance, err: %v", result.Error)
	}

	// Get the invoice
	var invoice *domain.Invoice
	whereClause = fmt.Sprintf("%s = ?", "id")
	result = Model(repo.Database).Table(Tables.invoices).Where(whereClause, remittance.InvoiceId).Find(&invoice)

	if result.Error != nil {
		return false, fmt.Errorf("AcceptRemittance: error querying invoice, err: %v", result.Error)
	}

	remittance.Status = 1
	result = Model(repo.Database).Table(Tables.remittances).Save(&remittance)

	if result.Error != nil {
		return false, fmt.Errorf("AcceptRemittance: error saving remittance, err: %v", result.Error)
	}

	invoice.Status = 1
	result = Model(repo.Database).Table(Tables.invoices).Save(&invoice)

	if result.Error != nil {
		return false, fmt.Errorf("AcceptRemittance: error saving invoice, err: %v", result.Error)
	}

	return true, nil
}

// --------------------------------------------------------- private members

func (repo *PostgresRepo) isThereAnyOverlapAmongOrders(newStartDate time.Time, newEndDate time.Time, orders []*domain.PlantOrder) bool {

	isAvailable := true
	for _, order := range orders {

		if order.Status == 0 { // if its accepted, then we need to check for an overlap

			isStartDateInTheRange := (newStartDate.Equal(order.StartDate) || newStartDate.After(order.StartDate)) && newStartDate.Before(order.EndDate)
			if isStartDateInTheRange || (newStartDate.Before(order.StartDate) && newEndDate.After(order.StartDate)) {
				isAvailable = false
				break
			}
		}
	}

	return isAvailable
}

// PS10
func (repo *PostgresRepo) generateReturnedPlantInvoice(plantOrder *domain.PlantOrder) {
	var plant *domain.Plant
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.ID)
	result := Model(repo.Database).Table(Tables.plants).Where(whereClause, plantOrder.PlantID).Take(&plant)

	if result.Error != nil {
		return
	}

	var invoiceId = repo.getMaxInvoiceID() + 1
	var daysPassed = plantOrder.EndDate.Sub(plantOrder.StartDate).Hours() / 24
	invoice := &domain.Invoice{
		ID:           invoiceId,
		PlantOrderId: plantOrder.ID,
		Price:        plant.PricePerDay * daysPassed,
		Status:       0,
	}

	result = Model(repo.Database).Table(Tables.invoices).Create(invoice)

	if result.Error != nil {
		return
	}

	plantOrder.InvoiceID = invoiceId
	result = Model(repo.Database).Table(Tables.plantOrders).Save(&plantOrder)
}

// getMaxPlantOrderID returns the MAX ID
func (repo *PostgresRepo) getMaxPlantOrderID() int64 {

	var id int64

	maxClause := fmt.Sprintf("max(%s)", PlantOrdersColumns.ID)
	row := Model(repo.Database).Table(Tables.plantOrders).Select(maxClause).Row()
	row.Scan(&id)

	return id
}

// getMaxInvoiceID returns the MAX INVOICE ID
func (repo *PostgresRepo) getMaxInvoiceID() int64 {

	var id int64

	row := Model(repo.Database).Table(Tables.invoices).Select("max(id)").Row()
	row.Scan(&id)

	return id
}

// getMaxRemittanceID returns the MAX REMITTANCE ID
func (repo *PostgresRepo) getMaxRemittanceID() int64 {

	var id int64

	row := Model(repo.Database).Table(Tables.remittances).Select("max(id)").Row()
	row.Scan(&id)

	return id
}

// getPlantOrdersByPlantID returns all the plants
func (repo *PostgresRepo) getPlantOrdersByPlantID(plantID int64) ([]*domain.PlantOrder, error) {

	var plantOrders []*domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", PlantOrdersColumns.PlantID)
	result := Model(repo.Database).Table(Tables.plantOrders).Where(whereClause, plantID).Find(&plantOrders)
	if result.Error != nil {
		return nil, fmt.Errorf("GetPlantOrdersByPlantID: error querying plantOrders, err: %v", result.Error)
	}

	return plantOrders, nil
}

func (repo *PostgresRepo) GetUnPaidInvoicesFor(customerID int64) ([]*domain.Invoice, error) {

	var invoices []*domain.Invoice
	whereClause := "invoices.status = ? and plantorders.customerid = ?"
	joinClause := "join plantorders on invoices.plantorderid = plantorders.id"
	result := Model(repo.Database).Table(Tables.invoices).Joins(joinClause).Where(whereClause, 0, customerID).Find(&invoices)
	if result.Error != nil {
		return nil, fmt.Errorf("GetunPaidInvoices: error querying invoices, err: %v", result.Error)
	}

	return invoices, nil
}
