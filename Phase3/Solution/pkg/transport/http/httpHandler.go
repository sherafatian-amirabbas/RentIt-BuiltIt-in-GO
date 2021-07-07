package http

import (
	"github.com/cs-ut-ee/hw3-group-3/pkg/service"
	"github.com/cs-ut-ee/hw3-group-3/pkg/transport/dto"

	log "github.com/sirupsen/logrus"

	"encoding/json"
	"net/http"

	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	service service.Service
	router  *mux.Router
}

func NewHTTPHandler(service service.Service, router *mux.Router) httpHandler {
	return httpHandler{
		service: service,
		router:  router,
	}
}

func (handler httpHandler) RegisterRoutes() {
	handler.router.HandleFunc("/plants", handler.GetAllPlants).Methods(http.MethodGet)
	handler.router.HandleFunc("/plants/GetPrice/{id}/{from}/{to}", handler.GetPlantPrice).Methods(http.MethodGet)
	handler.router.HandleFunc("/plants/IsPlantAvailable/{id}/{from}/{to}", handler.IsPlantAvailable).Methods(http.MethodGet)

	handler.router.HandleFunc("/customers", handler.GetAllCustomers).Methods(http.MethodGet)
	handler.router.HandleFunc("/orders/{startDate}", handler.GetPlantOrdersByStartDate).Methods(http.MethodGet)
	handler.router.HandleFunc("/orders/{startDate}/{pageNumber}/{pageSize}", handler.GetPagedPlantOrdersByStartDate).Methods(http.MethodGet)
	handler.router.HandleFunc("/orders/new/{plantID}/{customerID}/{startDate}/{endDate}", handler.NewPlantOrder).Methods(http.MethodPost)
	handler.router.HandleFunc("/orders/update/{orderPlantID}/{plantID}/{startDate}/{endDate}", handler.updatePlantOrder).Methods(http.MethodPost)

	handler.router.HandleFunc("/customers/{customerID}/invoices/sendReminder", handler.sendInvoiceRemiderToCustomer).Methods(http.MethodPost)
	handler.router.HandleFunc("/invoices/sendReminder", handler.sendInvoiceRemiders).Methods(http.MethodPost)

	handler.router.HandleFunc("/orders/cancel/{orderPlantID}", handler.cancelPlantOrder).Methods(http.MethodPost)
	handler.router.HandleFunc("/orders/reject/{orderPlantID}", handler.rejectPlantByCustomer).Methods(http.MethodPost)
	handler.router.HandleFunc("/orders/rental/{orderPlantID}", handler.rentalPeriodExpired).Methods(http.MethodPost)

	handler.router.HandleFunc("/remittance/create/{orderPlantID}/{referenceNumber}", handler.createRemittance).Methods(http.MethodPost)
	handler.router.HandleFunc("/remittance/accept/{remittanceId}", handler.acceptRemittance).Methods(http.MethodPost)
}

func (handler httpHandler) GetAllPlants(w http.ResponseWriter, _ *http.Request) {

	plants, err := handler.service.GetAllPlants()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plantsDTO := dto.GetPlantDTOList(plants)

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plantsDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) GetPlantPrice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	from, err := time.Parse("2006-01-02", params["from"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	to, err := time.Parse("2006-01-02", params["to"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	price, err := handler.service.GetPlantPrice(id, from, to)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&price)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) IsPlantAvailable(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	from, err := time.Parse("2006-01-02", params["from"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	to, err := time.Parse("2006-01-02", params["to"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isAvailable, err := handler.service.IsPlantAvailable(id, from, to)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&isAvailable)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) GetAllCustomers(w http.ResponseWriter, _ *http.Request) {

	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customersDTO := dto.GetCustomerDTOList(customers)

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&customersDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) GetPlantOrdersByStartDate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	startDate, err := time.Parse("2006-01-02", params["startDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plantOrders, err := handler.service.GetPlantOrdersByStartDate(startDate)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ordersDTO := dto.GetPlantOrderDTOList(plantOrders)

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&ordersDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) GetPagedPlantOrdersByStartDate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pageNumber, _ := strconv.Atoi(params["pageNumber"])
	pageSize, _ := strconv.Atoi(params["pageSize"])

	startDate, err := time.Parse("2006-01-02", params["startDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plantOrders, err := handler.service.GetPagedPlantOrdersByStartDate(startDate, pageNumber, pageSize)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ordersDTO := dto.GetPlantOrderDTOList(plantOrders)

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&ordersDTO)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler httpHandler) NewPlantOrder(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	plantID, _ := strconv.ParseInt(params["plantID"], 10, 64)
	customerID, _ := strconv.ParseInt(params["customerID"], 10, 64)

	startDate, err := time.Parse("2006-01-02", params["startDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	endDate, err := time.Parse("2006-01-02", params["endDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.service.NewPlantOrder(plantID, customerID, startDate, endDate)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

func (handler httpHandler) updatePlantOrder(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	orderPlantID, _ := strconv.ParseInt(params["orderPlantID"], 10, 64)
	plantID, _ := strconv.ParseInt(params["plantID"], 10, 64)

	startDate, err := time.Parse("2006-01-02", params["startDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	endDate, err := time.Parse("2006-01-02", params["endDate"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.service.UpdatePlantOrder(orderPlantID, plantID, startDate, endDate)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

func (handler httpHandler) sendInvoiceRemiderToCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID, err := strconv.ParseInt(params["customerID"], 10, 64)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.service.SendReminderFor(customerID)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

func (handler httpHandler) sendInvoiceRemiders(w http.ResponseWriter, r *http.Request) {
	err := handler.service.SendReminders()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
}

//PS7
func (handler httpHandler) cancelPlantOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plantOrderId, err := strconv.ParseInt(params["orderPlantID"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := handler.service.CancelPlantOrder(plantOrderId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res == true {
		// write success response
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

//PS8
func (handler httpHandler) rejectPlantByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plantOrderId, err := strconv.ParseInt(params["orderPlantID"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := handler.service.RejectPlantByCustomer(plantOrderId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res == true {
		// write success response
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

//PS9
func (handler httpHandler) rentalPeriodExpired(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plantOrderId, err := strconv.ParseInt(params["orderPlantID"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := handler.service.RentalPeriodExpired(plantOrderId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res == true {
		// write success response
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

//PS12
func (handler httpHandler) createRemittance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plantOrderId, err := strconv.ParseInt(params["orderPlantID"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	referenceNumber := string(params["referenceNumber"])
	res, err := handler.service.CreateRemittance(plantOrderId, referenceNumber)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res == true {
		// write success response
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (handler httpHandler) acceptRemittance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	remittanceId, err := strconv.ParseInt(params["remittanceId"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := handler.service.AcceptRemittance(remittanceId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res == true {
		// write success response
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
