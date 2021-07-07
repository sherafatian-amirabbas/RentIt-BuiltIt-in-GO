package http

import (
	"github.com/cs-ut-ee/project-group-3/pkg/domain"
	"github.com/cs-ut-ee/project-group-3/pkg/service"
	"github.com/cs-ut-ee/project-group-3/pkg/transport/dto"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type httpHandler struct {
	service *service.Service
	router  *mux.Router
}

func NewHTTPHandler(service *service.Service, router *mux.Router) *httpHandler {
	return &httpHandler{
		service: service,
		router:  router,
	}
}

func (handler httpHandler) RegisterRoutes() {
	handler.router.HandleFunc("/requests/new", handler.NewPlantHireRequest).Methods(http.MethodPost)

	handler.router.HandleFunc("/requests/modifyBySiteEngineers",
		handler.ModifyRequestBySiteEngineers).Methods(http.MethodPost)

	handler.router.HandleFunc("/requests/get/{Id}", handler.GetPlantHireRequestById).Methods(http.MethodGet)

	handler.router.HandleFunc("/requests/modifyByWorkEngineers",
		handler.ModifyRequestByWorkEngineers).Methods(http.MethodPost)

	handler.router.HandleFunc("/requests/accept", handler.AcceptRequest).Methods(http.MethodPost)

	handler.router.HandleFunc("/requests/reject", handler.RejectRequest).Methods(http.MethodPost)

	handler.router.HandleFunc("/requests/cancel/{requestId}/{requesterName}/{comment}",
		handler.CancelPlantHireRequest).Methods(http.MethodPost)
	handler.router.HandleFunc("/requests/generatePO/{plantHireRequestID}",
		handler.GeneratePurchaseOrder).Methods(http.MethodPost)
	handler.router.HandleFunc("/requests/extension/{plantHireRequestID}/{endDate}",
		handler.RequestExtension).Methods(http.MethodPost)
	handler.router.HandleFunc("/requests/delete/{supplierName}", handler.DeleteRequestsBySupplierName).Methods(http.MethodDelete)

	handler.router.HandleFunc("/orders/getAll", handler.GetCompleteOrders).Methods(http.MethodGet)

	handler.router.HandleFunc("/invoices/{invoiceId}/accept/{regulator}/{comment}", handler.AcceptInvoice).Methods(http.MethodPost)
	handler.router.HandleFunc("/invoices/{invoiceId}/reject/{regulator}/{comment}", handler.RejectInvoice).Methods(http.MethodPost)
	handler.router.HandleFunc("/invoices/{invoiceId}/purchaseOrder", handler.GetInvoicePurchaseOrder).Methods(http.MethodGet)
	handler.router.HandleFunc("/invoices/{invoiceId}", handler.GetInvoice).Methods(http.MethodGet)
	handler.router.HandleFunc("/invoices/checkInvoice", handler.CheckInvoice).Methods(http.MethodPost) //accept invoice as json

	handler.router.HandleFunc("/supplier/request/cancel/{orderId}", handler.CancelRequest).Methods(http.MethodPost)           //dummy supplier endpoint
	handler.router.HandleFunc("/supplier/request/purchase_order", handler.PurchaseOrderAcceptance).Methods(http.MethodPost) //dummy to process PO
	handler.router.HandleFunc("/supplier/remittance/create/{purchaseOrderId}/{referenceNumber}", handler.CreateRemittance).Methods(http.MethodPost)
} // dummy supplier remittance

// CC1
func (handler httpHandler) NewPlantHireRequest(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	requestObject := &domain.PlantHireRequest{}
	err := json.Unmarshal(reqBody, requestObject)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := handler.service.NewPlantHireRequest(requestObject.PlantName, requestObject.SiteName,
		requestObject.SupplierName, requestObject.RequesterName, requestObject.StartDate, requestObject.EndDate,
		requestObject.TotalHiringCost, requestObject.Status)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestDTO := dto.GetPlantHireRequestDTO(result)

	err = json.NewEncoder(w).Encode(requestDTO)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CC4
func (handler httpHandler) GetPlantHireRequestById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	requestId, _ := strconv.ParseInt(params["Id"], 10, 64)

	plantRequest, err := handler.service.GetPlantHireRequestById(requestId)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestDTO := dto.GetPlantHireRequestDTO(plantRequest)

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(requestDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// CC2
func (handler httpHandler) ModifyRequestBySiteEngineers(w http.ResponseWriter, r *http.Request) {
	handler.modifyRequest(w, r, true)
}

// CC5
func (handler httpHandler) ModifyRequestByWorkEngineers(w http.ResponseWriter, r *http.Request) {
	handler.modifyRequest(w, r, false)
}

// CC5
func (handler httpHandler) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	handler.confirmRequest(w, r, true)
}

// CC5
func (handler httpHandler) RejectRequest(w http.ResponseWriter, r *http.Request) {
	handler.confirmRequest(w, r, false)
}

func (handler httpHandler) CancelPlantHireRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	requestId, _ := strconv.ParseInt(params["requestId"], 10, 64)
	requesterName := params["requesterName"]
	comment := params["comment"]

	err := handler.service.CancelPlantHireRequest(requestId, requesterName, comment)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

func (handler httpHandler) GeneratePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestId, _ := strconv.ParseInt(params["plantHireRequestID"], 10, 64)

	po, err := handler.service.GeneratePurchaseOrder(requestId)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(domain.PlantOrder(*po))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) RequestExtension(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestId, _ := strconv.ParseInt(params["plantHireRequestID"], 10, 64)
	endDate, err := time.Parse("2006-01-02", params["endDate"])

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	po, err := handler.service.RequestExtension(requestId, endDate)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(domain.PlantOrder(*po))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// CC7
func (handler httpHandler) GetCompleteOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := handler.service.GetCompleteOrders()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.GetCompleteOrderDTOList(orders))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// CC11
func (handler httpHandler) AcceptInvoice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	invoiceId, _ := strconv.ParseInt(params["invoiceId"], 10, 64)
	regulator := params["regulator"]
	comment := params["comment"]

	invoice, err := handler.service.AcceptInvoice(invoiceId, regulator, comment)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.GetInvoiceDTO(invoice))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// CC11
func (handler httpHandler) RejectInvoice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	invoiceId, _ := strconv.ParseInt(params["invoiceId"], 10, 64)
	regulator := params["regulator"]
	comment := params["comment"]

	invoice, err := handler.service.RejectInvoice(invoiceId, regulator, comment)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.GetInvoiceDTO(invoice))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// CC11
func (handler httpHandler) GetInvoicePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	invoiceId, _ := strconv.ParseInt(params["invoiceId"], 10, 64)

	purchaseOrder, err := handler.service.GetPurchaseOrderByInvoiceId(invoiceId)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseOrderDTO := dto.GetCompleteOrderDTO(purchaseOrder)
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(purchaseOrderDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	invoiceId, _ := strconv.ParseInt(params["invoiceId"], 10, 64)

	invoice, err := handler.service.GetInvoice(invoiceId)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoiceDTO := dto.GetInvoiceDTO(invoice)
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(invoiceDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) CheckInvoice(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var invoice domain.Invoice
	err := decoder.Decode(&invoice)

	if err != nil {
		log.Errorf("Could not decode json, err %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.CheckInvoice(&invoice)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

// CC13
// Our "unique" feature was the possibility to delete all plant hire request with a specific supplier
// For example if we exchange supplier and the previous one doesn't want us to store their information anymore then we should delete the data connected to them
func (handler httpHandler) DeleteRequestsBySupplierName(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	supplierName := params["supplierName"]

	deleted, err := handler.service.DeleteRequestsBySupplierName(supplierName)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.GetPlantHireRequestDTOList(deleted))
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

//..............Dummy cancel EndPoint------------
//If the response header is accepted then request accepted else rejected
func (handler httpHandler) CancelRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId, _ := strconv.ParseInt(params["orderId"], 10, 64)

	if orderId%2 == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

//dummy for processing Purchase Order
//It accept PO as JSON
func (handler httpHandler) PurchaseOrderAcceptance(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var po domain.CompleteOrder
	err := decoder.Decode(&po)

	if err != nil {
		log.Errorf("Could not decode json, err %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if po.RequestID%2 == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (handler httpHandler) CreateRemittance(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ---------------------------------------------- private methods

func (handler httpHandler) modifyRequest(w http.ResponseWriter, r *http.Request, isSiteEngineer bool) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	requestObject := &domain.PlantHireRequest{}
	err := json.Unmarshal(reqBody, requestObject)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSiteEngineer {
		err = handler.service.ModifyRequestBySiteEngineers(requestObject.ID, requestObject.PlantName,
			requestObject.SiteName, requestObject.SupplierName, requestObject.RequesterName,
			requestObject.StartDate, requestObject.EndDate, requestObject.TotalHiringCost)
	} else {
		err = handler.service.ModifyRequestByWorkEngineers(requestObject.ID, requestObject.PlantName,
			requestObject.SiteName, requestObject.SupplierName, requestObject.RequesterName,
			requestObject.StartDate, requestObject.EndDate, requestObject.TotalHiringCost)
	}
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

func (handler httpHandler) confirmRequest(w http.ResponseWriter, r *http.Request, isAccepted bool) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	confirmationObject := &dto.WorkEngineerConfirmationDTO{}
	err := json.Unmarshal(reqBody, confirmationObject)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isAccepted {
		err = handler.service.AcceptRequest(confirmationObject.RequestId, confirmationObject.Regulator,
			confirmationObject.Comment)
	} else {
		err = handler.service.RejectRequest(confirmationObject.RequestId, confirmationObject.Regulator,
			confirmationObject.Comment)
	}

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}
