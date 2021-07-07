package dto

import (
	"github.com/cs-ut-ee/project-group-3/pkg/domain"
)

// invoice will hold the information about the invoice
type InvoiceDTO struct {
	Id         		 int64  `json:"id"`
	PurchaseOrderID  int64  `json:"purchaseOrderID"`
	Amount  		 float64  `json:"amount"`
	Regulator  		 string  `json:"regulator"`
	Comment  		 string  `json:"comment"`
	StatusCode 		 int64  `json:"statusCode"`
	StatusDesc 		 string `json:"statusDesc"`
}

func GetInvoiceDTOList(invoices []*domain.Invoice) []*InvoiceDTO {

	var dtoList = []*InvoiceDTO{}
	for _, order := range invoices {
		dto := GetInvoiceDTO(order)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetInvoiceDTO(invoice *domain.Invoice) *InvoiceDTO {

	statusDesc := domain.GetInvoiceStatusDescription(invoice.Status)
	dto := &InvoiceDTO{
		Id:         invoice.ID,
		PurchaseOrderID:  invoice.PurchaseOrderID,
		Amount:  invoice.Amount,
		Regulator:  invoice.Regulator,
		Comment:  invoice.Comment,
		StatusCode: invoice.Status,
		StatusDesc: statusDesc,
	}

	return dto
}
