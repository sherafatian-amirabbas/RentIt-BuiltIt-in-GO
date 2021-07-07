package domain

import "time"

// PlantOrder will hold the information about the order
type PlantOrder struct {
	ID         int64     `gorm:"column:id"`
	PlantID    int64     `gorm:"column:plantid"`
	CustomerID int64     `gorm:"column:customerid"`
	StartDate  time.Time `gorm:"column:startdate"`
	EndDate    time.Time `gorm:"column:enddate"`
	Status     int64     `gorm:"column:status"` // 0: accepted, 1: rejected, 2: canceled, 3: returned
	InvoiceID  int64     `gorm:"column:invoiceid"`
}
