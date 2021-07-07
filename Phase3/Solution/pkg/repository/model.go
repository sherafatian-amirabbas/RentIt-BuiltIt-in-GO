package repository

import (
	"gorm.io/gorm"
)

// ------------------------------------------------------------- DB Model

type modelStruct struct {
	db *gorm.DB
}

// Will provide access to the gorm model
func Model(db *gorm.DB) *modelStruct {
	return &modelStruct{
		db: db,
	}
}

// Will return the specific table defined in gorm model
func (model *modelStruct) Table(tableName string) *gorm.DB {

	return model.db.Table(tableName)
}

// ------------------------------------------------------------- DB Schema

type tablesStruct struct {
	customers            string
	plants               string
	plantOrders          string
	cancellationRequests string
	invoices             string
	remittances          string
}

var Tables = &tablesStruct{
	customers:            "customers",
	plants:               "plants",
	plantOrders:          "plantorders",
	cancellationRequests: "cancellationrequests",
	invoices:             "invoices",
	remittances:          "remittances",
}

// ------------------------------------------------------------- Table Schema

type plantOrdersColumnsStruct struct {
	ID         string
	PlantID    string
	CustomerID string
	StartDate  string
	EndDate    string
	Status     string
	InvoiceID  string
}

var PlantOrdersColumns = &plantOrdersColumnsStruct{
	ID:         "id",
	PlantID:    "plantid",
	CustomerID: "customerid",
	StartDate:  "startdate",
	EndDate:    "enddate",
	Status:     "status",
	InvoiceID:  "invoiceid",
}
