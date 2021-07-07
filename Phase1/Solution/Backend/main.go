package main

import (
	controller "ESI-Homework1/Backend/Controllers"
	"log"
	"net/http"

	// Type go get -u github.com/gorilla/mux to install
	"github.com/gorilla/mux"
)

func main() {
	registerRoutes()
}

func registerRoutes() {

	var homeController = controller.HomeController{}
	var itemController = controller.ItemController{}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeController.GetHomePage).Methods("GET")

	router.HandleFunc("/items", itemController.GetItems).Methods("GET")
	router.HandleFunc("/items/{key}", itemController.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}/complete", itemController.CompleteItem).Methods("GET")
	router.HandleFunc("/items/create", itemController.CreateItem).Methods("POST")
	router.HandleFunc("/items/delete/{id}", itemController.DeleteItem).Methods("DELETE")
	router.HandleFunc("/items/update/{id}", itemController.UpdateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
