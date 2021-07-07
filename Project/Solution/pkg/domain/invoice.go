package domain

// Invoice will hold the information about an invoice
type Invoice struct {
	ID              int64
	PurchaseOrderID int64
	Amount          float64
	Regulator       string // site engineer name who approved the request
	Comment         string
	Status          int64 // 0: Sent, 1: Approved, 2: Rejected, 3: Paid
}

func GetInvoiceStatusDescription(status int64) string {
	var statusDesc string
	switch status {
	case 0:
		statusDesc = "Sent"
	case 1:
		statusDesc = "Approved"
	case 2:
		statusDesc = "Rejected"
	case 3:
		statusDesc = "Paid"
	}
	return statusDesc
}
