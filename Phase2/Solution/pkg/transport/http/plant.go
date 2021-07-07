package http

import (
	"time"

	"github.com/cs-ut-ee/hw2-group-3/pkg/service"

	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type PlantHandler struct {
	service service.IService
	router  *mux.Router
}

func NewPlantHandler(service service.IService, router *mux.Router) *PlantHandler {
	return &PlantHandler{
		service: service,
		router:  router,
	}
}

func (handler *PlantHandler) RegisterRoutes() {
	handler.router.HandleFunc("/plants", handler.GetAllPlants).Methods(http.MethodGet)
	handler.router.HandleFunc("/plants/GetPrice/{id}/{from}/{to}", handler.GetPlantPrice).Methods(http.MethodGet)
	handler.router.HandleFunc("/plants/IsPlantAvailable/{id}/{from}/{to}", handler.IsPlantAvailable).Methods(http.MethodGet)
}

func (handler *PlantHandler) GetAllPlants(w http.ResponseWriter, _ *http.Request) {

	plants, err := handler.service.GetAllPlants()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler *PlantHandler) GetPlantPrice(w http.ResponseWriter, r *http.Request) {

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

	plants, err := handler.service.GetPlantPrice(id, from, to)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (handler *PlantHandler) IsPlantAvailable(w http.ResponseWriter, r *http.Request) {

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
