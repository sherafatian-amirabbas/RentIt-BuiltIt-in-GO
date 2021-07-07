package domain

// Invoice will contain the necessary information about an invoice
type Invoice struct {
	ID              int64     `gorm:"column:id"`
	PlantOrderId    int64     `gorm:"column:plantorderid"`
	Price 			float64   `gorm:"column:price"`
	Status     		int64     `gorm:"column:status"` // 0: unpaid, 1: paid
	RemittanceID  	int64     `gorm:"column:remittanceid"`
}