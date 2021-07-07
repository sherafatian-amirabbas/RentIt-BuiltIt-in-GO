package domain

import "time"

// CompleteOrder will hold the complete information about the order
type CompleteOrder struct {
	OrderID             int64
	RequestID           int64
	PlantName           string
	SiteName            string
	SupplierName        string
	RequesterName       string
	StartDate           time.Time
	EndDate             time.Time
	TotalHiringCost     float64
	Regulator           string
	WorkEngineerComment string
	RequestStatus       int64
	RequestStatusDesc   string
	OrderStatus         int64 // 0: Sent, 1: RejectedBySupplier, 2: Delivered, 3: RejectedByTheSite, 4: AcceptedBySupplier, 5: Paid
	OrderStatusDesc     string
	InvoiceStatus       int64 // 0: Sent, 1: Approved, 2: Rejected, 3: Paid
	InvoiceStatusDesc   string
}
