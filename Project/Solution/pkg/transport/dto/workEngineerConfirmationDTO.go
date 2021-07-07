package dto

// workEngineerConfirmationDTO will hold the information about the Approval/Rejection issued by work engineers
type WorkEngineerConfirmationDTO struct {
	RequestId int64  `json:"requestId"`
	Regulator string `json:"regulator"`
	Comment   string `json:"comment"`
}
