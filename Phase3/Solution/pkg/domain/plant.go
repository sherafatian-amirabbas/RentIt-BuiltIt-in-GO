package domain

// Plant will hold the information about the equipment
type Plant struct {
	ID          int64     `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	PricePerDay float64     `gorm:"column:priceperday"`
}
