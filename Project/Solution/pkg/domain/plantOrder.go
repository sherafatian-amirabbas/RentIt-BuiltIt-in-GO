package domain

// PlantOrder will hold the information about the order
//Its actually purchase order
type PlantOrder struct {
	ID        int64
	RequestID int64
	Status    int64 // 0: Sent, 1: RejectedBySupplier, 2: Delivered, 3: RejectedByTheSite, 4: AcceptedBySupplier, 5: Paid
}

func GetOrderStatusDescription(status int64) string {
	var statusDesc string
	switch status {
	case 0:
		statusDesc = "Sent"
	case 1:
		statusDesc = "RejectedBySupplier"
	case 2:
		statusDesc = "Delivered"
	case 3:
		statusDesc = "RejectedByTheSite"
	case 4:
		statusDesc = "AcceptedBySupplier"
	case 5:
		statusDesc = "Paid"
	}
	return statusDesc
}
