package model

import (
	"time"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"

	log "github.com/sirupsen/logrus"
)

// postgresRepo Makes possible to access DB
type Model struct {
	connectionString string
	Database         *gorm.DB
}

// NewModel creates the Model by receiving the ConnectionString
func NewModel(connection string) *Model {
	return &Model{
		connectionString: connection,
	}
}

// InitialMigration keeps the database up to date
func (dbModel *Model) InitialMigration() {
	dbModel.OpenConnection()
	err := dbModel.Database.AutoMigrate(&domain.PlantHireRequest{}, &domain.PlantOrder{}, &domain.Invoice{})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// OpenConnection opens the connection
func (dbModel *Model) OpenConnection() {
	db, err := gorm.Open(postgres.Open(dbModel.connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	dbModel.Database = db
}

// CloseConnection closes the connection
func (dbModel *Model) CloseConnection() {
	sqlDB, err := dbModel.Database.DB()
	if err != nil {
		log.Fatalf("Could not get DB to close")
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Could not disconnect from Postgres DB")
	}
}

func (dbModel *Model) InitialDatabase() error {

	// ------------------------- PlantHireRequest sample data

	from, _ := time.Parse("2006-01-02", "2020-05-20")
	to, _ := time.Parse("2006-01-02", "2020-06-26")
	plantRequest1 := &domain.PlantHireRequest{
		PlantName:       "Plant1",
		SiteName:        "Site1",
		SupplierName:    "RentIt",
		RequesterName:   "SiteEngineer1",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 100.50,
		Status:          0, // pending
	}
	result := dbModel.Database.Create(plantRequest1)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	from, _ = time.Parse("2006-01-02", "2021-05-20")
	to, _ = time.Parse("2006-01-02", "2021-06-20")
	plantRequest2 := &domain.PlantHireRequest{
		PlantName:           "Plant1",
		SiteName:            "Site1",
		SupplierName:        "RentIt",
		RequesterName:       "SiteEngineer2",
		StartDate:           from,
		EndDate:             to,
		TotalHiringCost:     100.50,
		Regulator:           "WorkEngineer1",
		WorkEngineerComment: "Duplicate Request",
		Status:              1, // accepted
	}
	result = dbModel.Database.Create(plantRequest2)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	from, _ = time.Parse("2006-01-02", "2020-05-20")
	to, _ = time.Parse("2006-01-02", "2020-05-25")
	plantRequest3 := &domain.PlantHireRequest{
		PlantName:           "Plant2",
		SiteName:            "Site1",
		SupplierName:        "RentIt",
		RequesterName:       "SiteEngineer2",
		StartDate:           from,
		EndDate:             to,
		TotalHiringCost:     150.50,
		Regulator:           "WorkEngineer1",
		WorkEngineerComment: "Request Approved",
		Status:              2, // rejected
	}
	result = dbModel.Database.Create(plantRequest3)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	from, _ = time.Parse("2006-01-02", "2020-05-20")
	to, _ = time.Parse("2006-01-02", "2020-06-25")
	plantRequest4 := &domain.PlantHireRequest{
		PlantName:           "Plant4",
		SiteName:            "Site4",
		SupplierName:        "RentIt",
		RequesterName:       "SiteEngineer2",
		StartDate:           from,
		EndDate:             to,
		TotalHiringCost:     250.50,
		Regulator:           "WorkEngineer4",
		WorkEngineerComment: "",
		Status:              3, // Cancelled
	}
	result = dbModel.Database.Create(plantRequest4)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	from, _ = time.Parse("2006-01-02", "2020-05-10") //current date passed, could not be cancelled
	to, _ = time.Parse("2006-01-02", "2020-06-25")
	plantRequest5 := &domain.PlantHireRequest{
		PlantName:           "Plant5",
		SiteName:            "Site5",
		SupplierName:        "RentIt",
		RequesterName:       "SiteEngineer2",
		StartDate:           from,
		EndDate:             to,
		TotalHiringCost:     250.50,
		Regulator:           "WorkEngineer4",
		WorkEngineerComment: "",
		Status:              0, // accepted
	}
	result = dbModel.Database.Create(plantRequest5)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	// ------------------------- PlantOrder sample data

	order1 := &domain.PlantOrder{
		RequestID: plantRequest3.ID,
		Status:    0, // Sent
	}
	result = dbModel.Database.Create(order1)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	order2 := &domain.PlantOrder{
		RequestID: plantRequest2.ID,
		Status:    5, // Paid
	}
	result = dbModel.Database.Create(order2)
	if result.Error != nil {
		return fmt.Errorf("database initialization failed, err: %v", result.Error)
	}

	// ------------------------- Invoice sample data

	invoice1 := &domain.Invoice{
		ID:              1,
		PurchaseOrderID: order1.ID,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          0,
	}
	if dbModel.Database.Model(&invoice1).Select("purchaseOrderId", "amount", "regulator", "comment", "status").Updates(invoice1).RowsAffected == 0 {
		dbModel.Database.Create(invoice1)
	}

	invoice2 := &domain.Invoice{
		ID:              2,
		PurchaseOrderID: order1.ID,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          0,
	}
	if dbModel.Database.Model(&invoice2).Select("purchaseOrderId", "amount", "regulator", "comment", "status").Updates(invoice2).RowsAffected == 0 {
		dbModel.Database.Create(invoice2)
	}

	invoice3 := &domain.Invoice{
		ID:              3,
		PurchaseOrderID: order2.ID,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          0,
	}
	if dbModel.Database.Model(&invoice2).Select("purchaseOrderId", "amount", "regulator", "comment", "status").Updates(invoice2).RowsAffected == 0 {
		dbModel.Database.Create(invoice3)
	}

	return nil
}

// ------------------------------------------------------------- DB Schema

type tablesStruct struct {
	PlantHireRequest string
	PlantOrders      string
	Invoices         string
}

var Tables = &tablesStruct{
	PlantHireRequest: "plant_hire_requests",
	PlantOrders:      "plant_orders",
	Invoices:         "invoices",
}

// ------------------------------------------------------------- Table Schema - PlantHireRequests

type PlantHireRequestsColumnsStruct struct {
	ID                  string
	PlantName           string
	SiteName            string
	SupplierName        string
	RequesterName       string
	StartDate           string
	EndDate             string
	TotalHiringCost     string
	Regulator           string
	WorkEngineerComment string
	Status              string
}

var PlantHireRequestsColumns = &PlantHireRequestsColumnsStruct{
	ID:                  "id",
	PlantName:           "plant_name",
	SiteName:            "site_name",
	SupplierName:        "supplier_name",
	RequesterName:       "requester_name",
	StartDate:           "start_date",
	EndDate:             "end_date",
	TotalHiringCost:     "total_hiring_cost",
	Regulator:           "regulator",
	WorkEngineerComment: "work_engineer_comment",
	Status:              "status",
}

// ------------------------------------------------------------- Table Schema - PlantOrder

type PlantOrdersColumnsStruct struct {
	ID        string
	RequestID string
	Status    string
}

var PlantOrdersColumns = &PlantOrdersColumnsStruct{
	ID:        "id",
	RequestID: "request_id",
	Status:    "status",
}

// ------------------------------------------------------------- Table Schema - Invoice
type InvoiceColumnsStruct struct {
	ID              string
	PurchaseOrderID string
	Amount          string
	Regulator       string
	Comment         string
	Status          string
}

var InvoiceColumns = &InvoiceColumnsStruct{
	ID:              "id",
	PurchaseOrderID: "purchase_order_id",
	Amount:          "amount",
	Regulator:       "regulator",
	Comment:         "comment",
	Status:          "status",
}
