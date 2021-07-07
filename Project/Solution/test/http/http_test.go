package httptest

import (
	"strconv"
	"strings"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
	"github.com/cs-ut-ee/project-group-3/pkg/transport/dto"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func ApiUrl() string {
	url, success := os.LookupEnv("apiUrl")
	if !success {
		panic("Environment variable 'httpUrl' is not defined")
		// url = "http://localhost:8081"
	}

	return url
}

// CC1
func TestNewPlantHireRequest(t *testing.T) {
	makeNewPlantHireRequest(t)
}

// CC2
func TestModifyRequestBySiteEngineersRejectionTest(t *testing.T) {
	// make a new request
	plantHireRequest := makeNewPlantHireRequest(t)

	// accept a request, to show site engineers cannot modify the requests after being accepted
	confirmPlantHireRequest(t, plantHireRequest, true)

	plantHireRequest.PlantName = "PlantName - modified by site engineers"
	plantHireRequest.SiteName = "SiteName - modified by site engineers"
	plantHireRequest.SupplierName = "SupplierName - modified by site engineers"
	plantHireRequest.RequesterName = "RequesterName - modified by site engineers"
	plantHireRequest.StartDate = plantHireRequest.StartDate.AddDate(0, 1, 0)
	plantHireRequest.EndDate = plantHireRequest.EndDate.AddDate(0, 2, 0)
	plantHireRequest.TotalHiringCost *= 3

	// modify request
	modifyPlantHireRequest(t, dto.GetPlantHireRequestFromDTO(plantHireRequest), true)

	// get request
	modifiedRequestObject := getPlantHireRequestById(t, plantHireRequest.ID)

	if plantHireRequest.PlantName == modifiedRequestObject.PlantName ||
		plantHireRequest.SiteName == modifiedRequestObject.SiteName ||
		plantHireRequest.SupplierName == modifiedRequestObject.SupplierName ||
		plantHireRequest.RequesterName == modifiedRequestObject.RequesterName ||
		plantHireRequest.StartDate.Equal(modifiedRequestObject.StartDate) ||
		plantHireRequest.EndDate.Equal(modifiedRequestObject.EndDate) ||
		plantHireRequest.TotalHiringCost == modifiedRequestObject.TotalHiringCost {
		t.Error("TestModifyRequestBySiteEngineersRejectionTest: the PlantHireRequestID was modified by site engineers after being accepted!")
		return
	}
}

// CC2
func TestModifyRequestBySiteEngineersAcceptanceTest(t *testing.T) {
	// make a request
	plantHireRequest := makeNewPlantHireRequest(t)

	plantHireRequest.PlantName = "PlantName - modified by work engineers"
	plantHireRequest.SiteName = "SiteName - modified by work engineers"
	plantHireRequest.SupplierName = "SupplierName - modified by work engineers"
	plantHireRequest.RequesterName = "RequesterName - modified by work engineers"
	plantHireRequest.StartDate = plantHireRequest.StartDate.AddDate(0, 0, 1)
	plantHireRequest.EndDate = plantHireRequest.EndDate.AddDate(0, 0, 2)
	plantHireRequest.TotalHiringCost *= 2

	modifyPlantHireRequest(t, dto.GetPlantHireRequestFromDTO(plantHireRequest), true)

	// get
	modifiedRequestObject := getPlantHireRequestById(t, plantHireRequest.ID)

	if plantHireRequest.PlantName != modifiedRequestObject.PlantName ||
		plantHireRequest.SiteName != modifiedRequestObject.SiteName ||
		plantHireRequest.SupplierName != modifiedRequestObject.SupplierName ||
		plantHireRequest.RequesterName != modifiedRequestObject.RequesterName ||
		!plantHireRequest.StartDate.Equal(modifiedRequestObject.StartDate) ||
		!plantHireRequest.EndDate.Equal(modifiedRequestObject.EndDate) ||
		plantHireRequest.TotalHiringCost != modifiedRequestObject.TotalHiringCost {
		t.Error("TestModifyRequestBySiteEngineersAcceptanceTest: the PlantHireRequest was not modified properly!")
		return
	}
}

// CC4
func TestGetPlantHireRequestById(t *testing.T) {
	plantHireRequest := makeNewPlantHireRequest(t)
	requestObj := getPlantHireRequestById(t, plantHireRequest.ID)
	if requestObj.ID != plantHireRequest.ID {
		t.Error("TestGetPlantHireRequestById: Couldn't get PlantHireRequest")
	}
}

// CC5
func TestModifyRequestByWorkEngineers(t *testing.T) {
	// make a request
	plantHireRequest := makeNewPlantHireRequest(t)

	// accept a request, to show work engineers can modify the requests without limitation
	confirmPlantHireRequest(t, plantHireRequest, true)

	plantHireRequest.PlantName = "PlantName - modified by work engineers"
	plantHireRequest.SiteName = "SiteName - modified by work engineers"
	plantHireRequest.SupplierName = "SupplierName - modified by work engineers"
	plantHireRequest.RequesterName = "RequesterName - modified by work engineers"
	plantHireRequest.StartDate = plantHireRequest.StartDate.AddDate(0, 0, 1)
	plantHireRequest.EndDate = plantHireRequest.EndDate.AddDate(0, 0, 2)
	plantHireRequest.TotalHiringCost *= 2

	modifyPlantHireRequest(t, dto.GetPlantHireRequestFromDTO(plantHireRequest), false)

	// get
	modifiedRequestObject := getPlantHireRequestById(t, plantHireRequest.ID)

	if plantHireRequest.PlantName != modifiedRequestObject.PlantName ||
		plantHireRequest.SiteName != modifiedRequestObject.SiteName ||
		plantHireRequest.SupplierName != modifiedRequestObject.SupplierName ||
		plantHireRequest.RequesterName != modifiedRequestObject.RequesterName ||
		!plantHireRequest.StartDate.Equal(modifiedRequestObject.StartDate) ||
		!plantHireRequest.EndDate.Equal(modifiedRequestObject.EndDate) ||
		plantHireRequest.TotalHiringCost != modifiedRequestObject.TotalHiringCost {
		t.Error("TestModifyRequestByWorkEngineers: the PlantHireRequest was not modified properly!")
		return
	}
}

// CC5
func TestAcceptRequest(t *testing.T) {
	plantHireRequest := makeNewPlantHireRequest(t)
	confirmPlantHireRequest(t, plantHireRequest, true)
}

// CC5
func TestRejectRequest(t *testing.T) {
	plantHireRequest := makeNewPlantHireRequest(t)
	confirmPlantHireRequest(t, plantHireRequest, false)
}

// CC7
func TestGetCompleteOrders(t *testing.T) {
	// make a new request
	plantHireRequest := makeNewPlantHireRequest(t)

	// accept a request
	confirmPlantHireRequest(t, plantHireRequest, true)

	// create the order
	plantOrder, err := MakePurchaseOrder(plantHireRequest.ID)
	if err != nil {
		t.Error("TestGetCompleteOrders: couldn't create plant order. " + err.Error())
		return
	}

	orders := getCompleteOrders(t)

	var isCurrentRequestAndOrderFound bool = false
	for _, order := range orders {
		if order.RequestID == plantHireRequest.ID &&
			order.OrderID == plantOrder.ID {
			isCurrentRequestAndOrderFound = true
			break
		}
	}

	if isCurrentRequestAndOrderFound == false {
		t.Error("TestGetCompleteOrders: problem getting order")
		return
	}
}

//CC3
func TestCancelPlantHireRequest(t *testing.T) {

	var workEngineerComment = "Please cancel the request"

	from, _ := time.Parse("2006-01-02", "2021-08-25")
	to, _ := time.Parse("2006-01-02", "2021-09-17")
	plantRequest := domain.PlantHireRequest{
		PlantName:       "new-Plant1",
		SiteName:        "new-Site1",
		SupplierName:    "new-RentIt",
		RequesterName:   "new-SiteEngineer1",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          0,
	}

	respRequest, err := CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	id := respRequest.ID
	requester_name := respRequest.RequesterName

	//It should be successfully cancelled
	resp, err := http.Post(RequestUrl(id, requester_name, workEngineerComment), "", nil)

	if err != nil {
		t.Error("Plant Hire Request Cancellation Unsuccessful. Error: " + err.Error())
	}

	if resp != nil {
		if resp.Status != "200 OK" {
			t.Error("Plant Hire Request Cancellation Unsuccessful.")
		}
	}

	//Start date is already passed
	from, _ = time.Parse("2006-01-02", "2021-05-10")
	to, _ = time.Parse("2006-01-02", "2021-06-17")
	plantRequest = domain.PlantHireRequest{
		PlantName:       "new-Plant2",
		SiteName:        "new-Site2",
		SupplierName:    "new-RentIt",
		RequesterName:   "new-SiteEngineer2",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          0,
	}

	respRequest, err = CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	id = respRequest.ID
	requester_name = respRequest.RequesterName

	//It should be rejected cancellation can be done before one day of start date
	resp, err = http.Post(RequestUrl(id, requester_name, workEngineerComment), "", nil)

	if err != nil {
		t.Error("Posting cancellation request failed. Error: " + err.Error())
	}

	if resp != nil {
		if resp.Status == "200 OK" {
			t.Error("Cancellation should not be accepted!")
		}
	}

	//Cancellation is sent to supplier as its accepted
	from, _ = time.Parse("2006-01-02", "2021-08-25")
	to, _ = time.Parse("2006-01-02", "2021-09-17")
	plantRequest = domain.PlantHireRequest{
		PlantName:       "new-Plant3",
		SiteName:        "new-Site3",
		SupplierName:    "new-RentIt",
		RequesterName:   "new-SiteEngineer23",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1,
	}

	respRequest, err = CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	id = respRequest.ID
	requester_name = respRequest.RequesterName

	//if request status is accepted and startdate is after current date, cancellation request should be sent to supplier. Supplier can accept or reject it
	resp, err = http.Post(RequestUrl(id, requester_name, workEngineerComment), "", nil)

	if err != nil {
		t.Error("Posting cancellation request failed. Error: " + err.Error())
	}

	if resp.StatusCode != 200 && resp.StatusCode != 400 {
		t.Error("Cancellation request is not processed by supplier.")
	}

	if resp.StatusCode == 400 {
		respText, _ := ioutil.ReadAll(resp.Body)
		if !strings.Contains(string(respText), "request is rejected by supplier") {
			t.Error("Cancellation request is not processed by supplier.")
		}
	}
}

func RequestUrl(id int64, requester_name string, workEngineerComment string) string {
	return fmt.Sprintf("%s/requests/cancel/%d/%s/%s", ApiUrl(), id, requester_name, workEngineerComment)
}

//CC6
//Produce a PO for every approved plant hire request
func TestGeneratePurchaseOrder(t *testing.T) {
	from, _ := time.Parse("2006-01-02", "2021-08-25")
	to, _ := time.Parse("2006-01-02", "2021-09-17")
	plantRequest := domain.PlantHireRequest{
		PlantName:       "new-Plant1",
		SiteName:        "new-Site1",
		SupplierName:    "new-RentIt",
		RequesterName:   "Lauri Leiten",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1,
	}

	respRequest, err := CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	//For this request PO should be generated and PO status should be updated due to supplier feedback
	purchaseOrder, err := MakePurchaseOrder(respRequest.ID)

	if err != nil {
		t.Error("Error :" + err.Error() + " ID: " + strconv.FormatInt(respRequest.ID, 10))
	}

	if purchaseOrder == nil {
		t.Error("Purchase order is not successfully created")
	} else if purchaseOrder.Status == 0 {
		t.Error("Purchase order is not processed by supplier")
	}

	//Not approved hire request
	from, _ = time.Parse("2006-01-02", "2021-08-25")
	to, _ = time.Parse("2006-01-02", "2021-09-17")
	plantRequest2 := domain.PlantHireRequest{
		PlantName:       "new-Plant23",
		SiteName:        "new-Site13",
		SupplierName:    "new-RentIt",
		RequesterName:   "Einar Linde",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          0, //Pending
	}

	respRequest2, err := CreatePlantHireRequest(&plantRequest2)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	//This request should raise error as purchase request status is not approved
	_, err = MakePurchaseOrder(respRequest2.ID)

	if err == nil {
		t.Error("Purchase request should raise error")
	}

}

//CC8
//Request Extension
func TestRequestExtension(t *testing.T) {
	//Send Accepted Senario
	from, _ := time.Parse("2006-01-02", "2021-08-25")
	to, _ := time.Parse("2006-01-02", "2021-09-17")
	plantRequest := domain.PlantHireRequest{
		PlantName:       "new-Plant23",
		SiteName:        "new-Site13",
		SupplierName:    "new-RentIt",
		RequesterName:   "Einar Linde",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1, //accepted
	}

	respRequest, err := CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	po, err := RequestExtension(respRequest.ID, "2021-09-27")

	if err != nil {
		t.Error("Should be successfully extended. Error: " + err.Error())
	}

	if po == nil {
		t.Error("Purchase order is not successfully created")
	} else if po.Status == 0 {
		t.Error("Purchase order is not processed by supplier")
	}

	//Should be rejected as end date is already passed
	from, _ = time.Parse("2006-01-02", "2021-03-25")
	to, _ = time.Parse("2006-01-02", "2021-04-17")
	plantRequest = domain.PlantHireRequest{
		PlantName:       "new-Plant23",
		SiteName:        "new-Site13",
		SupplierName:    "new-RentIt",
		RequesterName:   "Einar Linde",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1, //accepted
	}

	respRequest, err = CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	_, err = RequestExtension(respRequest.ID, "2021-09-27")

	if err == nil {
		t.Error("Extension should be rejected as end date is already passed")
	}

	//Should be rejected as it is not accepted
	from, _ = time.Parse("2006-01-02", "2021-06-25")
	to, _ = time.Parse("2006-01-02", "2021-06-17")
	plantRequest = domain.PlantHireRequest{
		PlantName:       "new-Plant23",
		SiteName:        "new-Site13",
		SupplierName:    "new-RentIt",
		RequesterName:   "Einar Linde",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          0, //pending
	}

	respRequest, err = CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}

	_, err = RequestExtension(respRequest.ID, "2021-09-27")

	if err == nil {
		t.Error("Extension should be rejected as hire request is not approved")
	}

}

//CC10
//Check PO number with exisiting upaid invoice
func TestCheckInvoice(t *testing.T) {
	//Create PlantHireRequest
	from, _ := time.Parse("2006-01-02", "2021-08-25")
	to, _ := time.Parse("2006-01-02", "2021-09-17")
	plantRequest := domain.PlantHireRequest{
		PlantName:       "new-Plant23",
		SiteName:        "new-Site13",
		SupplierName:    "new-RentIt",
		RequesterName:   "Einar Linde",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1, //accepted
	}

	respRequest, err := CreatePlantHireRequest(&plantRequest)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}
	//Make PurchaseOrder
	purchaseOrder, err := MakePurchaseOrder(respRequest.ID)

	if err != nil {
		t.Error("Error :" + err.Error() + " ID: " + strconv.FormatInt(respRequest.ID, 10))
	}

	if purchaseOrder == nil {
		t.Error("Purchase order is not successfully created")
	} else if purchaseOrder.Status == 0 {
		t.Error("Purchase order is not processed by supplier")
	}
	//Check Invoice for that purchase Order
	invoice := domain.Invoice{
		ID:              10,
		PurchaseOrderID: purchaseOrder.ID,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          0,
	}

	url := ApiUrl() + "/invoices/checkInvoice"

	jsonInv, _ := json.Marshal(invoice)
	byteJsonPO := []byte(jsonInv)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(byteJsonPO))

	if err != nil {
		t.Error("Could not post to check invoice" + err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Invoice should be accepted but its rejected")
	}

	invoice = domain.Invoice{
		ID:              11,
		PurchaseOrderID: purchaseOrder.ID,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          3,
	}

	jsonInv, _ = json.Marshal(invoice)
	byteJsonPO = []byte(jsonInv)

	resp, err = http.Post(url, "application/json", bytes.NewBuffer(byteJsonPO))

	if err != nil {
		t.Error("Could not post to check invoice" + err.Error())
	}

	if resp.StatusCode == 200 {
		t.Error("Invoice should be rejected but its accepted")
	}
}

// CC11
func TestApproveInvoice(t *testing.T) {
	var id int64 = 1
	var regulator string = "approve_regulator"
	var comment string = "ok"

	resp, err := http.Post(fmt.Sprintf("%s/invoices/%d/accept/%s/%s", ApiUrl(), id, regulator, comment), "", nil)
	if err != nil {
		t.Error("TestApproveInvoice: couldn't approve the invoice")
		return
	}

	savedInvoiceJSON, _ := ioutil.ReadAll(resp.Body)
	savedInvoice := dto.InvoiceDTO{}
	err = json.Unmarshal(savedInvoiceJSON, &savedInvoice)
	if err != nil {
		t.Error("TestApproveInvoice: error on Unmarshaling the response")
		return
	}

	if savedInvoice.Id != id {
		t.Error("TestApproveInvoice: invoice.ID is not the same as requested!")
		return
	}

	if savedInvoice.Regulator != regulator {
		t.Error("TestApproveInvoice: the regulator should be the same as sent!")
		return
	}

	if savedInvoice.Comment != comment {
		t.Error("TestApproveInvoice: the regulator should be the same as sent!")
		return
	}

	if savedInvoice.StatusCode != 1 {
		t.Error("TestApproveInvoice: the status should be approved!")
		return
	}

	if savedInvoice.StatusDesc != "Approved" {
		t.Error("TestApproveInvoice: the status should be approved!")
		return
	}
}

// CC11
func TestRejectInvoice(t *testing.T) {
	var id int64 = 2
	var regulator string = "reject_regulator"
	var comment string = "definitely not ok"

	resp, err := http.Post(fmt.Sprintf("%s/invoices/%d/reject/%s/%s", ApiUrl(), id, regulator, comment), "", nil)
	if err != nil {
		t.Error("TestRejectInvoice: couldn't reject the invoice")
		return
	}

	savedInvoiceJSON, _ := ioutil.ReadAll(resp.Body)
	savedInvoice := dto.InvoiceDTO{}
	err = json.Unmarshal(savedInvoiceJSON, &savedInvoice)
	if err != nil {
		t.Error("TestRejectInvoice: error on Unmarshaling the response")
		return
	}

	if savedInvoice.Id != id {
		t.Error("TestRejectInvoice: invoice.ID is not the same as requested!")
		return
	}

	if savedInvoice.Regulator != regulator {
		t.Error("TestRejectInvoice: the regulator should be the same as sent!")
		return
	}

	if savedInvoice.Comment != comment {
		t.Error("TestRejectInvoice: the regulator should be the same as sent!")
		return
	}

	if savedInvoice.StatusCode != 2 {
		t.Error("TestRejectInvoice: the status should be rejected!")
		return
	}

	if savedInvoice.StatusDesc != "Rejected" {
		t.Error("TestRejectInvoice: the status should be rejected!")
		return
	}
}

func TestGetPurchaseOrderById(t *testing.T) {
	var id int64 = 1

	resp, err := http.Get(fmt.Sprintf("%s/invoices/%d/purchaseOrder", ApiUrl(), id))
	if err != nil {
		t.Error("TestGetPurchaseOrderById: couldn't fetch the purchase order")
		return
	}

	orderJSON, _ := ioutil.ReadAll(resp.Body)
	order := dto.CompleteOrderDTO{}
	err = json.Unmarshal(orderJSON, &order)
	if err != nil {
		fmt.Println(err)
		t.Error("TestGetPurchaseOrderById: error on Unmarshaling the response")
		return
	}
}

// CC13
func TestPlantRequestDeletion(t *testing.T) {
	var supplierName string = "RentItNewer"
	var requesterName string = "Lauri Leiten"

	from, _ := time.Parse("2006-01-02", "2021-08-25")
	to, _ := time.Parse("2006-01-02", "2021-09-17")

	plantRequestCreation := domain.PlantHireRequest{
		PlantName:       "new-Plant432",
		SiteName:        "new-Site1432",
		SupplierName:    supplierName,
		RequesterName:   requesterName,
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
		Status:          1, //accepted
	}

	_, err := CreatePlantHireRequest(&plantRequestCreation)

	if err != nil {
		t.Error("Plant Hire Request Creation Unsuccessful. Error: " + err.Error())
	}
	
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/requests/delete/%s", ApiUrl(), supplierName), nil)
	if err != nil {
		t.Error("TestPlantRequestDeletion: couldn't create request")
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("TestPlantRequestDeletion: couldn't delete the necessary data")
		return
	}
	defer resp.Body.Close()
	
	plantRequestJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestPlantRequestDeletion: couldn't read the response")
		return
	}
	plantRequest := []*dto.PlanRequestDTO{}
	err = json.Unmarshal(plantRequestJSON, &plantRequest)
	if err != nil {
		fmt.Println(err)
		t.Error("TestPlantRequestDeletion: error on Unmarshaling the response")
		return
	}

	for _, element := range plantRequest {
		if element.SupplierName != supplierName {
			t.Error("TestPlantRequestDeletion: an element with the wrong supplier name was returned")
			return
		}
		if element.RequesterName != requesterName {
			t.Error("TestPlantRequestDeletion: an element with the wrong requester name was returned")
			return
		}
	}
}


//.........Helper functions..........................................
//ReQuestExtension
func RequestExtension(plantHireRequestId int64, endDate string) (*domain.PlantOrder, error) {
	url := fmt.Sprintf("%s/requests/extension/%d/%s", ApiUrl(), plantHireRequestId, endDate)
	resp, err := http.Post(url, "", nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf(string(respBody))
	}

	purchaseOrderJSON, _ := ioutil.ReadAll(resp.Body)
	purchaseOrder := domain.PlantOrder{}
	err = json.Unmarshal(purchaseOrderJSON, &purchaseOrder)
	if err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

//MakePurchaseOrder
func MakePurchaseOrder(plantHireRequestId int64) (*domain.PlantOrder, error) {
	url := fmt.Sprintf("%s/requests/generatePO/%d", ApiUrl(), plantHireRequestId)
	resp, err := http.Post(url, "", nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf(string(respBody))
	}

	purchaseOrderJSON, _ := ioutil.ReadAll(resp.Body)
	purchaseOrder := domain.PlantOrder{}
	err = json.Unmarshal(purchaseOrderJSON, &purchaseOrder)
	if err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

//Created this method to get id of Plant Request
//And it might be needed for testing
func CreatePlantHireRequest(hireRequest *domain.PlantHireRequest) (*domain.PlantHireRequest, error) {
	plantRequestJSON, _ := json.Marshal(hireRequest)
	resp, err := http.Post(ApiUrl()+"/requests/new", "", bytes.NewBuffer(plantRequestJSON))
	if err != nil {
		return nil, err

	}

	savedPlantRequestJSON, _ := ioutil.ReadAll(resp.Body)
	savedPlantRequest := domain.PlantHireRequest{}
	err = json.Unmarshal(savedPlantRequestJSON, &savedPlantRequest)
	if err != nil {
		return nil, err
	}

	return &savedPlantRequest, nil
}

// --------------- private methods

func makeNewPlantHireRequest(t *testing.T) *dto.PlanRequestDTO {
	from, _ := time.Parse("2006-01-02", "2021-05-11")
	to, _ := time.Parse("2006-01-02", "2021-05-17")
	plantRequest := domain.PlantHireRequest{
		PlantName:       "new-Plant1",
		SiteName:        "new-Site1",
		SupplierName:    "new-RentIt",
		RequesterName:   "new-SiteEngineer1",
		StartDate:       from,
		EndDate:         to,
		TotalHiringCost: 123.50,
	}

	plantRequestJSON, _ := json.Marshal(plantRequest)
	resp, err := http.Post(ApiUrl()+"/requests/new", "", bytes.NewBuffer(plantRequestJSON))
	if err != nil {
		t.Error("makeNewPlantHireRequest: couldn't create the plants")
		return nil
	}

	savedPlantRequestJSON, _ := ioutil.ReadAll(resp.Body)
	savedPlantRequest := dto.PlanRequestDTO{}
	err = json.Unmarshal(savedPlantRequestJSON, &savedPlantRequest)
	if err != nil {
		t.Error("makeNewPlantHireRequest: error on Unmarshaling the response")
		return nil
	}

	if savedPlantRequest.ID == 0 {
		t.Error("makeNewPlantHireRequest: savedPlantRequest.ID is not valid")
		return nil
	}

	if savedPlantRequest.StatusCode != 0 {
		t.Error("makeNewPlantHireRequest: when the plantHiringRequest is saved, the status should be pending!")
		return nil
	}

	return &savedPlantRequest
}

func getPlantHireRequestById(t *testing.T, plantHireRequestId int64) *dto.PlanRequestDTO {
	endpoint := fmt.Sprintf("%s/requests/get/%d", ApiUrl(), plantHireRequestId)
	resp, err := http.Get(endpoint)
	if err != nil {
		t.Error("getPlantHireRequestById: Problem reading request via REST.")
		return nil
	}

	requestedObjectJSON, _ := ioutil.ReadAll(resp.Body)
	requestedObject := dto.PlanRequestDTO{}
	_ = json.Unmarshal(requestedObjectJSON, &requestedObject)
	if requestedObject.ID != plantHireRequestId {
		t.Error("getPlantHireRequestById: Couldn't find or parse PlantHireRequest")
		return nil
	}

	return &requestedObject
}

func confirmPlantHireRequest(t *testing.T, plantHireRequest *dto.PlanRequestDTO, isAccepted bool) {
	confirmation := dto.WorkEngineerConfirmationDTO{
		RequestId: plantHireRequest.ID,
		Regulator: "work-engineer-1",
		Comment:   "it's approved",
	}

	// approve the request
	confirmationJSON, _ := json.Marshal(confirmation)
	var endpoint string
	if isAccepted {
		endpoint = ApiUrl() + "/requests/accept"
	} else {
		endpoint = ApiUrl() + "/requests/reject"
	}
	_, err := http.Post(endpoint, "", bytes.NewBuffer(confirmationJSON))
	if err != nil {
		t.Error("confirmPlantHireRequest: couldn't accept/reject the plantHireRequest")
		return
	}

	// get the request
	modifiedPlantHireRequest := getPlantHireRequestById(t, plantHireRequest.ID)

	// check the status
	if isAccepted {
		if modifiedPlantHireRequest.StatusCode != 1 {
			t.Error("confirmPlantHireRequest: the status of the plantHireRequest wasn't accepted!")
			return
		}
	} else {
		if modifiedPlantHireRequest.StatusCode != 2 {
			t.Error("confirmPlantHireRequest: the status of the plantHireRequest wasn't rejected!")
			return
		}
	}
}

func modifyPlantHireRequest(t *testing.T, plantHireRequest *domain.PlantHireRequest, isBySiteEngineers bool) {
	plantHireRequestJSON, _ := json.Marshal(plantHireRequest)

	var endpoint string
	if isBySiteEngineers {
		endpoint = ApiUrl() + "/requests/modifyBySiteEngineers"
	} else {
		endpoint = ApiUrl() + "/requests/modifyByWorkEngineers"
	}

	_, err := http.Post(endpoint, "", bytes.NewBuffer(plantHireRequestJSON))
	if err != nil {
		t.Error("modifyPlantHireRequest: couldn't update the request")
		return
	}
}

func getCompleteOrders(t *testing.T) []dto.CompleteOrderDTO {
	endpoint := fmt.Sprintf("%s/orders/getAll", ApiUrl())
	resp, err := http.Get(endpoint)
	if err != nil {
		t.Error("getCompleteOrders: Problem reading orders via REST.")
		return nil
	}

	ordersJSON, _ := ioutil.ReadAll(resp.Body)

	var orders []dto.CompleteOrderDTO
	_ = json.Unmarshal([]byte(ordersJSON), &orders)

	return orders
}
