package dto

import (
	"time"
	"github.com/cs-ut-ee/project-group-3/pkg/domain"
)

// customer will hold the information about the customer
type CompleteOrderDTO struct {
	OrderID             int64  `json:"orderId"`
	RequestID           int64  `json:"requestId"`
	PlantName           string  `json:"plantName"`
	SiteName            string  `json:"siteName"`
	SupplierName        string  `json:"supplierName"`
	RequesterName       string  `json:"requesterName"`
	StartDate           time.Time  `json:"startDate"`
	EndDate             time.Time  `json:"endDate"`
	TotalHiringCost     float64  `json:"totalHiringCost"`
	Regulator           string  `json:"regulator"`
	WorkEngineerComment string  `json:"workEngineerComment"`
	RequestStatus       int64  `json:"requestStatus"`
	RequestStatusDesc   string  `json:"requestStatusDesc"`
	OrderStatus         int64  `json:"orderStatus"`
	OrderStatusDesc     string  `json:"orderStatusDesc"`
	InvoiceStatus         int64  `json:"invoiceStatus"`
	InvoiceStatusDesc     string  `json:"invoiceStatusDesc"`
}

func GetCompleteOrderDTOList(plantOrders []*domain.CompleteOrder) []*CompleteOrderDTO {

	var dtoList = []*CompleteOrderDTO{}
	for _, order := range plantOrders {
		dto := GetCompleteOrderDTO(order)
		dtoList = append(dtoList, dto)
	}

	return dtoList
}

func GetCompleteOrderDTO(order *domain.CompleteOrder) *CompleteOrderDTO {

	order.RequestStatusDesc = domain.GetRequestStatusDescription(order.RequestStatus)
	order.OrderStatusDesc = domain.GetOrderStatusDescription(order.OrderStatus)
	order.InvoiceStatusDesc = domain.GetInvoiceStatusDescription(order.InvoiceStatus)

	dto := &CompleteOrderDTO{
		OrderID:         order.OrderID,
		RequestID:  order.RequestID,
		PlantName: order.PlantName,
		SiteName: order.SiteName,
		SupplierName: order.SupplierName,
		RequesterName: order.RequesterName,
		StartDate: order.StartDate,
		EndDate: order.EndDate,
		TotalHiringCost: order.TotalHiringCost,
		Regulator: order.Regulator,
		WorkEngineerComment: order.WorkEngineerComment,
		RequestStatus: order.RequestStatus,
		RequestStatusDesc: order.RequestStatusDesc,
		OrderStatus: order.OrderStatus,
		OrderStatusDesc: order.OrderStatusDesc,
		InvoiceStatus: order.InvoiceStatus,
		InvoiceStatusDesc: order.InvoiceStatusDesc,
	}

	return dto
}
