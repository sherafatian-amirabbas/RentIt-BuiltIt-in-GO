package domain

import "time"

// PlantHireRequest will hold the information about the plant hiring request send by site engineer
type PlantHireRequest struct {
	ID                  int64
	PlantName           string
	SiteName            string
	SupplierName        string
	RequesterName       string
	StartDate           time.Time
	EndDate             time.Time
	TotalHiringCost     float64
	Regulator           string // work engineer name who approved the request
	WorkEngineerComment string
	Status              int64 // 0: pending, 1: accepted, 2: rejected, 3: canceled
}

func GetRequestStatusDescription(status int64) string {
	var statusDesc string
	switch status {
	case 0:
		statusDesc = "Pending"
	case 1:
		statusDesc = "Accepted"
	case 2:
		statusDesc = "Rejected"
	case 3:
		statusDesc = "Canceled"
	}
	return statusDesc
}
