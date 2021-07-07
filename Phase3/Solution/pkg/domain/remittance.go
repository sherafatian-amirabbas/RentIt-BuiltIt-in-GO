package domain

// Plant will hold the information about the equipment
type Remittance struct {
	ID          int64     `gorm:"column:id"`
	InvoiceId        int64     `gorm:"column:invoiceid"`
	ReferenceNumber string     `gorm:"column:referencenumber"`
	Status int64     `gorm:"column:status"` // 0 pending | 1 received
}
