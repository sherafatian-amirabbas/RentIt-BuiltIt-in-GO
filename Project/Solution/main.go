package main

import (
	"fmt"
	"net/http"

	"github.com/cs-ut-ee/project-group-3/pkg/model"
	"github.com/cs-ut-ee/project-group-3/pkg/repository"
	"github.com/cs-ut-ee/project-group-3/pkg/service"
	http2 "github.com/cs-ut-ee/project-group-3/pkg/transport/http"

	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const httpServicePort = 8081

var connectionString string

func main() {

	log.Println("Server Started!")

	rentItEndpoint := GetRentItApiUrl()

	connectionString = getPostgresConnection()
	dbModel := updateAndGetModel()
	repo := repository.NewPostgresRepo(dbModel, rentItEndpoint)
	service := service.NewService(repo)

	// this should be run once, when it's needed to populate the DB with some primitive data
	err := dbModel.InitialDatabase()
	if err != nil {
		log.Errorf("Failed to initialize database")
	}

	// Start goroutine to receive invcoices from kafka
	go service.ReceiveInvoiceJob()

	runHTTPServer(service, dbModel)

	log.Infof("Server Stopped!")
}

func GetRentItApiUrl() string {
	url, success := os.LookupEnv("apiUrl")
	if !success {
		panic("Environment variable 'httpUrl' is not defined")
		// url = "http://localhost:8081"
	}

	return url
}

func getPostgresConnection() string {
	postgresConnection, success := os.LookupEnv("postgresConnectionString")
	if !success {
		postgresConnection = "dbname=buildit host=localhost password=postgres user=postgres sslmode=disable port=5432"
	}
	return postgresConnection
}

func updateAndGetModel() *model.Model {
	mdl := model.NewModel(connectionString)
	mdl.InitialMigration()
	return mdl
}

func runHTTPServer(service *service.Service, dbModel *model.Model) {

	httpRouter := mux.NewRouter().StrictSlash(true)

	plantHTTPHandler := http2.NewHTTPHandler(service, httpRouter)
	plantHTTPHandler.RegisterRoutes()

	// setup http server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}

	err := httpSrv.ListenAndServe()
	if err != nil {
		dbModel.CloseConnection()
		log.Fatalf("Could not start server")
	}
}
