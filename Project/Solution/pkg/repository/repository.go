package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
	"github.com/cs-ut-ee/project-group-3/pkg/model"

	"time"
)

// postgresRepo Makes possible to access DB
type Repository struct {
	DBModel          *model.Model
	SupplierEndPoint string
}

// NewPostgresRepo creates the repository by receiving the a DB
func NewPostgresRepo(dbModel *model.Model, supplierEndPoint string) *Repository {
	return &Repository{
		DBModel:          dbModel,
		SupplierEndPoint: supplierEndPoint,
	}
}

// CC1
// NewPlantHireRequest creates a new plant hire request
func (repo *Repository) NewPlantHireRequest(plantName string, siteName string, supplierName string,
	requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64, status int64) (*domain.PlantHireRequest, error) {

	request := &domain.PlantHireRequest{
		PlantName:           plantName,
		SiteName:            siteName,
		SupplierName:        supplierName,
		RequesterName:       requesterName,
		StartDate:           startDate,
		EndDate:             endDate,
		TotalHiringCost:     totalHiringCost,
		Regulator:           "",
		WorkEngineerComment: "",
		Status:              status,
	}

	result := repo.DBModel.Database.Create(request)
	if result.Error != nil {
		return nil, fmt.Errorf("NewPlantHireRequest: error creating the request, err: %v", result.Error)
	}

	return request, nil
}

// CC4
// GetPlantHireRequestById returns the PlantHireReequest object containing status
func (repo *Repository) GetPlantHireRequestById(requestId int64) (*domain.PlantHireRequest, error) {
	var request *domain.PlantHireRequest
	whereClause := fmt.Sprintf("%s = ?", model.PlantHireRequestsColumns.ID)
	result := repo.DBModel.Database.Where(whereClause, requestId).Take(&request)
	if result.Error != nil {
		return nil, fmt.Errorf("GetPlantHireRequestById: error getting PlantHireRequest, err: %v", result.Error)
	}

	return request, nil
}

func (repo *Repository) GetPurchaseOrderById(requestId int64) (*domain.PlantOrder, error) {
	var request *domain.PlantOrder
	whereClause := fmt.Sprintf("%s = ?", model.PlantOrdersColumns.ID)
	result := repo.DBModel.Database.Where(whereClause, requestId).Take(&request)
	if result.Error != nil {
		return nil, fmt.Errorf("GetPurchaseOrderById: error getting PlantOrder err: %v", result.Error)
	}

	return request, nil
}

func (repo *Repository) GetInvoiceById(invoiceId int64) (*domain.Invoice, error) {
	var invoice *domain.Invoice
	whereClause := fmt.Sprintf("%s = ?", model.InvoiceColumns.ID)
	result := repo.DBModel.Database.Where(whereClause, invoiceId).Take(&invoice)
	if result.Error != nil {
		return nil, fmt.Errorf("GetInvoiceById: error fetching invoice, err: %v", result.Error)
	}

	return invoice, nil
}

func (repo *Repository) GetCompletePurchaseOrderByInvoiceId(invoiceId int64) (*domain.CompleteOrder, error) {

	query := `
		SELECT plant_orders.id as order_id, plant_hire_requests.ID as request_id, plant_hire_requests.plant_name, 
			plant_hire_requests.site_name, plant_hire_requests.supplier_name, plant_hire_requests.requester_name,
			plant_hire_requests.start_date, plant_hire_requests.end_date, plant_hire_requests.total_hiring_cost, 
			plant_hire_requests.regulator, plant_hire_requests.work_engineer_comment, 
			plant_hire_requests.status as request_status, plant_orders.status as order_status,
			invoices.status as invoice_status
		FROM invoices
		LEFT JOIN plant_orders ON invoices.purchase_order_id = plant_orders.id
		LEFT JOIN plant_hire_requests ON plant_orders.request_id = plant_hire_requests.id
		WHERE invoices.id = ?
	`

	var completeOrders = []*domain.CompleteOrder{}
	result := repo.DBModel.Database.Raw(query, invoiceId).Scan(&completeOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(completeOrders) <= 0 {
		return nil, fmt.Errorf("GetCompletePurchaseOrderByInvoiceId: no valid purchase order was found!")
	}

	var firstOrder = completeOrders[0]
	firstOrder.RequestStatusDesc = domain.GetRequestStatusDescription(firstOrder.RequestStatus)
	firstOrder.OrderStatusDesc = domain.GetOrderStatusDescription(firstOrder.OrderStatus)
	firstOrder.InvoiceStatusDesc = domain.GetInvoiceStatusDescription(firstOrder.InvoiceStatus)

	return firstOrder, nil
}

// CC2
// ModifyRequestBySiteEngineers modifies the PlantHireRequest object by site engineers, prior to its approval by
func (repo *Repository) ModifyRequestBySiteEngineers(plantRequestId int64, plantName string, siteName string,
	supplierName string, requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error {

	return repo.ModifyRequest(plantRequestId, plantName, siteName, supplierName, requesterName, startDate,
		endDate, totalHiringCost, true)
}

// CC5
// ModifyRequestByWorkEngineers modifies PlantHireRequest without limitation
func (repo *Repository) ModifyRequestByWorkEngineers(plantRequestId int64, plantName string, siteName string,
	supplierName string, requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64) error {

	return repo.ModifyRequest(plantRequestId, plantName, siteName, supplierName, requesterName, startDate,
		endDate, totalHiringCost, false)
}

// CC5
// AcceptRequest changes the status of the PlantHireRequest object to accepted
func (repo *Repository) AcceptRequest(plantRequestId int64, regulator string, comment string) error {
	request, err := repo.GetPlantHireRequestById(plantRequestId)
	if err != nil {
		return err
	}

	if request.Status != 0 { // if it's not Pending
		return fmt.Errorf("AcceptRequestByWorkEngineers: only when the request is Pending, it can be approved!")
	}

	return repo.changePlantHireRequestStatus(request, regulator, comment, 1)
}

// CC5
// RejectRequest changes the status of the PlantHireRequest object to rejected
func (repo *Repository) RejectRequest(plantRequestId int64, regulator string, comment string) error {
	request, err := repo.GetPlantHireRequestById(plantRequestId)
	if err != nil {
		return err
	}

	if request.Status != 0 { // if it's not Pending
		return fmt.Errorf("RejectRequestByWorkEngineers: only when the request is Pending, it can be rejected!")
	}

	return repo.changePlantHireRequestStatus(request, regulator, comment, 2)
}

// CC7
// GetCompleteOrders will return the complete list of PlantHireRequests and Orders
func (repo *Repository) GetCompleteOrders() ([]*domain.CompleteOrder, error) {

	query := `
	SELECT plant_orders.id as order_id, plant_hire_requests.ID as request_id, plant_hire_requests.plant_name,
		plant_hire_requests.site_name, plant_hire_requests.supplier_name, plant_hire_requests.requester_name,
		plant_hire_requests.start_date, plant_hire_requests.end_date, plant_hire_requests.total_hiring_cost,
		plant_hire_requests.regulator, plant_hire_requests.work_engineer_comment,
		plant_hire_requests.status as request_status, plant_orders.status as order_status,
		invoices.status as invoice_status
	FROM plant_orders
	LEFT JOIN plant_hire_requests ON plant_orders.request_id = plant_hire_requests.id
	LEFT JOIN invoices ON invoices.purchase_order_id = plant_orders.id
	`
	var completeOrders = []*domain.CompleteOrder{}
	result := repo.DBModel.Database.Raw(query).Scan(&completeOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, order := range completeOrders {
		order.RequestStatusDesc = domain.GetRequestStatusDescription(order.RequestStatus)
		order.OrderStatusDesc = domain.GetOrderStatusDescription(order.OrderStatus)
		order.InvoiceStatusDesc = domain.GetInvoiceStatusDescription(order.InvoiceStatus)
	}

	return completeOrders, nil
}

//it will match request id and requesterName from database and if match found then change the status to cancel else throw error
//its done by site engineer
//if cancellation request is done after Purchase Order has been sent to supplier then it should be sent to supplier - how?
//May be - if request is accepted that means PO is generated and forwarded to supplier
//So check if request is accepted, if yes send request to mock supplier end point
//if response is true then update database
func (repo *Repository) CancelPlantHireRequest(plantHireRequestId int64, requesterName string, comment string) error {
	request, err := repo.GetPlantHireRequestById(plantHireRequestId)
	if err != nil || request.RequesterName != requesterName {
		return fmt.Errorf("requester name is not matching")
	}
	if request.Status != 2 && request.Status != 1 && request.StartDate.After(time.Now().Add(24*time.Hour)) {
		return repo.modifyPlantHireRequest(request, request.PlantName, request.SiteName, request.SupplierName, request.RequesterName, request.StartDate, request.EndDate,
			request.TotalHiringCost, request.Regulator, comment, 3)
	} else if request.Status == 1 && request.StartDate.After(time.Now().Add(24*time.Hour)) {

		//Send request to dummy supplier endpoint
		//if reponse is OK' then update repo else error
		resp, err := http.Post(repo.SupplierEndPoint + "/supplier/request/cancel/" + strconv.Itoa(int(plantHireRequestId)), "", nil)

		if err != nil {
			return err
		}

		if resp.StatusCode == 200 {
			return repo.modifyPlantHireRequest(request, request.PlantName, request.SiteName, request.SupplierName, requesterName, request.StartDate, request.EndDate,
				request.TotalHiringCost, request.Regulator, comment, 3)
		} else {
			return fmt.Errorf("request is rejected by supplier")
		}
	} else if request.Status == 3 {
		return fmt.Errorf("request is already cancelled")
	} else {
		return fmt.Errorf("cancel request is allowed before one day of start date. (start date is after current date) :  " + strconv.FormatBool(request.StartDate.After(time.Now().Add(24*time.Hour))))
	}
}

//CC6
func (repo *Repository) GeneratePurchaseOrder(PlantHireRequestID int64) (*domain.PlantOrder, error) {
	//get the plant hire request and check it accepted or not
	request, err := repo.GetPlantHireRequestById(PlantHireRequestID)
	if err != nil {
		return nil, err
	}

	if request.Status != 1 {
		return nil, fmt.Errorf("purchase order can be generated from approved request only")
	}

	PlantOrder := &domain.PlantOrder{
		RequestID: request.ID,
		Status:    0,
	}

	//if accepted then create CompleteOrder to sent to supplier
	PurchaseOrder := &domain.CompleteOrder{
		RequestID:           request.ID,
		PlantName:           request.PlantName,
		SupplierName:        request.SupplierName,
		SiteName:            request.SiteName,
		RequesterName:       request.RequesterName,
		StartDate:           request.StartDate,
		EndDate:             request.EndDate,
		TotalHiringCost:     request.TotalHiringCost,
		Regulator:           request.Regulator,
		WorkEngineerComment: request.WorkEngineerComment,
		RequestStatus:       request.Status,
		RequestStatusDesc:   domain.GetRequestStatusDescription(request.Status),
		OrderStatus:         PlantOrder.Status,
		OrderStatusDesc:     domain.GetOrderStatusDescription(PlantOrder.Status),
		InvoiceStatus:       0,
		InvoiceStatusDesc:   domain.GetInvoiceStatusDescription(0),
	}

	result := repo.DBModel.Database.Create(PlantOrder)

	if result.Error != nil {
		return nil, fmt.Errorf("Error saving the PO to database")
	}

	//send PO to supplier
	jsonPo, err := json.Marshal(PurchaseOrder)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	byteJsonPO := []byte(jsonPo)

	req, err := http.NewRequest("POST", repo.SupplierEndPoint+"/supplier/request/purchase_order", bytes.NewBuffer(byteJsonPO))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	//if supplier reject then change PO status to rejectedBySupplier
	if resp.StatusCode != 200 {
		PlantOrder.Status = 1
		result = repo.DBModel.Database.Save(PlantOrder)
		if result.Error != nil {
			return nil, fmt.Errorf("Error saving the PO to database")
		}
	} else {
		PlantOrder.Status = 4
		result = repo.DBModel.Database.Save(PlantOrder)
		if result.Error != nil {
			return nil, fmt.Errorf("Error saving the PO to database")
		}
	}

	return PlantOrder, nil
}

//CC8
//Get the original plant hire request
//Check its accepted and its enddate is not passed yet
//Modify it. End date should be higher than the previous end date
//Sent to supplier
func (repo *Repository) RequestExtension(plantHireRequestID int64, endDate time.Time) (*domain.PlantOrder, error) {

	//Get order by RequestID
	request, err := repo.GetPlantHireRequestById(plantHireRequestID)

	if err != nil {
		return nil, err
	}

	if request.Status != 1 || request.EndDate.Before(time.Now()) {
		return nil, fmt.Errorf("Wrong Request. Please generate new PO instead of extension!")
	}

	PlantOrder := &domain.PlantOrder{
		RequestID: request.ID,
		Status:    0,
	}

	//if accepted then create CompleteOrder to sent to supplier
	ModifiedPO := &domain.CompleteOrder{
		RequestID:           request.ID,
		PlantName:           request.PlantName,
		SupplierName:        request.SupplierName,
		SiteName:            request.SiteName,
		RequesterName:       request.RequesterName,
		StartDate:           request.StartDate,
		EndDate:             endDate,
		TotalHiringCost:     request.TotalHiringCost,
		Regulator:           request.Regulator,
		WorkEngineerComment: request.WorkEngineerComment,
		RequestStatus:       request.Status,
		RequestStatusDesc:   domain.GetRequestStatusDescription(request.Status),
		OrderStatus:         PlantOrder.Status,
		OrderStatusDesc:     domain.GetOrderStatusDescription(PlantOrder.Status),
		InvoiceStatus:       0,
		InvoiceStatusDesc:   domain.GetInvoiceStatusDescription(0),
	}

	//nolint:staticcheck
	result := repo.DBModel.Database.Create(PlantOrder)
	if result.Error != nil {
		return nil, fmt.Errorf("Error creating PO on Database")
	}

	//send PO to supplier
	jsonPo, err := json.Marshal(ModifiedPO)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	byteJsonPO := []byte(jsonPo)

	req, err := http.NewRequest("POST", repo.SupplierEndPoint+"/supplier/request/purchase_order", bytes.NewBuffer(byteJsonPO))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	//if supplier reject then change PO status to rejectedBySupplier
	if resp.StatusCode != 200 {
		PlantOrder.Status = 1
		result = repo.DBModel.Database.Save(PlantOrder)
		if result.Error != nil {
			return nil, fmt.Errorf("Error saving the PO to database")
		}
	} else {
		PlantOrder.Status = 4
		result = repo.DBModel.Database.Save(PlantOrder)
		request.EndDate = endDate
		if result.Error != nil {
			return nil, fmt.Errorf("Error saving the PO to database")
		}
		result = repo.DBModel.Database.Save(request) //request updated with new enddate
		if result.Error != nil {
			return nil, fmt.Errorf("Error saving the PO to database")
		}
	}

	return PlantOrder, nil
}

//CC10
//If no error returned then its OK to proceed further with the invoice
func (repo *Repository) CheckInvoice(invoice *domain.Invoice) error {
	//get order from invoice
	purchaseOrder, err := repo.GetPurchaseOrderById(invoice.PurchaseOrderID)

	if err != nil {
		return err
	}

	//Check the invoice status
	if invoice.Status == 3 || purchaseOrder.Status == 5 { //didn't count invoice rejected as rejected invoice can be modified and resubmitted
		return fmt.Errorf("Order is already paid") //No need to reject the invoice here as it is not mentioned on requirements
	}

	return nil
}

func (repo *Repository) NewInvoice(invoice *domain.Invoice) (*domain.Invoice, error) {
	order, err := repo.GetPurchaseOrderById(invoice.PurchaseOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed getting purchase order: %v", err.Error())
	}

	if order.Status == 5 {
		return nil, fmt.Errorf("purchase order is already paid")
	}

	result := repo.DBModel.Database.Create(invoice)
	if result.Error != nil {
		return nil, fmt.Errorf("NewInvoice: error creating the request, err: %v", result.Error)
	}

	return invoice, nil
}

// CC11
// Approves the invoice if possible and returns the new invoice object
// If the approval was successful, calls the RentIT aPI remittance creation endpoint
func (repo *Repository) AcceptInvoice(invoiceId int64, regulator string, comment string) (*domain.Invoice, error) {
	invoice, err := repo.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, err
	}

	if invoice.Status != 0 { // if it's not Pending
		return nil, fmt.Errorf("AcceptInvoiceBySiteEngineers: only when the request is Pending, it can be approved!")
	}

	err = repo.changeInvoiceAcceptanceStatus(invoice, regulator, comment, 1)
	if err != nil {
		return nil, fmt.Errorf("AcceptInvoiceBySiteEngineers: approving didn't work as expected!")
	}

	//CC12
	//send remittance advice to supplier
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/supplier/remittance/create/%d/%d", repo.SupplierEndPoint, invoice.PurchaseOrderID, invoiceId), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// revert approval
		_ = repo.changeInvoiceAcceptanceStatus(invoice, regulator, comment, 0)
		return nil, fmt.Errorf("AcceptInvoiceBySiteEngineers: remittance was not accepted (invoice wasn't accepted)")
	}

	finalInvoice, err := repo.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, err
	}

	return finalInvoice, nil
}

// CC11
// Rejects the invoice if possible and returns the new invoice object
func (repo *Repository) RejectInvoice(invoiceId int64, regulator string, comment string) (*domain.Invoice, error) {
	invoice, err := repo.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, err
	}

	if invoice.Status != 0 { // if it's not Pending
		return nil, fmt.Errorf("RejectInvoiceBySiteEngineers: only when the request is Pending, it can be approved!")
	}

	err = repo.changeInvoiceAcceptanceStatus(invoice, regulator, comment, 2)
	if err != nil {
		return nil, err
	}

	finalInvoice, err := repo.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, err
	}

	return finalInvoice, nil
}

// CC11
// Retrieves the purchase order associated with the given invoice ID
func (repo *Repository) GetPurchaseOrderByInvoiceId(invoiceId int64) (*domain.CompleteOrder, error) {
	invoice, err := repo.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, err
	}

	if invoice.PurchaseOrderID == 0 { // if it's not found
		return nil, fmt.Errorf("GetPurchaseOrderByInvoiceId: there is no purchase order attached to the invoice!")
	}

	purchaseOrder, err := repo.GetCompletePurchaseOrderByInvoiceId(invoice.ID)
	if err != nil {
		return nil, err
	}

	return purchaseOrder, nil
}

// CC13
func (repo *Repository) DeleteRequestsBySupplierName(supplierName string) ([]*domain.PlantHireRequest, error) {
	plantHire := []*domain.PlantHireRequest{}

	repo.DBModel.Database.Model(domain.PlantHireRequest{}).Where("supplier_name", supplierName).Find(&plantHire).Delete(&plantHire)

	return plantHire, nil
}

// ------------------------------------------------------------------------------- private members

func (repo *Repository) ModifyRequest(plantRequestId int64, plantName string, siteName string,
	supplierName string, requesterName string, startDate time.Time, endDate time.Time, totalHiringCost float64,
	isSiteEngineer bool) error {

	request, err := repo.GetPlantHireRequestById(plantRequestId)
	if err != nil {
		return err
	}

	if isSiteEngineer {
		if request.Status != 0 { // if it's not Pending
			return fmt.Errorf("ModifyRequest: Site engineers are allowed to modify the request if the request is still Pending!")
		}
	}

	return repo.modifyPlantHireRequest(request, plantName, siteName, supplierName, requesterName, startDate,
		endDate, totalHiringCost, request.Regulator, request.WorkEngineerComment, request.Status)
}

func (repo *Repository) changePlantHireRequestStatus(plantRequest *domain.PlantHireRequest, regulator string,
	comment string, status int64) error {

	return repo.modifyPlantHireRequest(plantRequest, plantRequest.PlantName, plantRequest.SiteName,
		plantRequest.SupplierName, plantRequest.RequesterName, plantRequest.StartDate, plantRequest.EndDate,
		plantRequest.TotalHiringCost, regulator, comment, status)
}

func (repo *Repository) changeInvoiceAcceptanceStatus(invoice *domain.Invoice, regulator string,
	comment string, status int64) error {

	return repo.modifyInvoice(invoice, invoice.Amount, regulator, comment, status)
}

func (repo *Repository) modifyPlantHireRequest(plantRequest *domain.PlantHireRequest, plantName string,
	siteName string, supplierName string, requesterName string, startDate time.Time, endDate time.Time,
	totalHiringCost float64, regulator string, workEngineerComment string, status int64) error {

	plantRequest.PlantName = plantName
	plantRequest.SiteName = siteName
	plantRequest.SupplierName = supplierName
	plantRequest.RequesterName = requesterName
	plantRequest.StartDate = startDate
	plantRequest.EndDate = endDate
	plantRequest.TotalHiringCost = totalHiringCost
	plantRequest.Regulator = regulator
	plantRequest.WorkEngineerComment = workEngineerComment
	plantRequest.Status = status

	result := repo.DBModel.Database.Save(plantRequest)
	if result.Error != nil {
		return fmt.Errorf("modifyPlantHireRequest: error saving the changes, err: %v", result.Error)
	}

	return nil
}

func (repo *Repository) modifyInvoice(invoice *domain.Invoice, amount float64, regulator string, comment string, status int64) error {

	invoice.Amount = amount
	invoice.Regulator = regulator
	invoice.Comment = comment
	invoice.Status = status

	result := repo.DBModel.Database.Save(invoice)
	if result.Error != nil {
		return fmt.Errorf("modifyInvoice: error saving the changes, err: %v", result.Error)
	}

	return nil
}
