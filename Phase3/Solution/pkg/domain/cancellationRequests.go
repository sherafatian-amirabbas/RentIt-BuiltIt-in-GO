package domain

import "time"

type CancellationRequest struct {
	ID             int64     `gorm:"column:id"`
	PlantOrderId   int64     `gorm:"column:plantorderid"`
	SubmissionDate time.Time `gorm:"column:submissiondate"`
	Status         int64     `gorm:"column:status"` //0: rejected, 1: Accepted, 2: Pending
}
