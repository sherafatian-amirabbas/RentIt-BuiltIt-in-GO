package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cs-ut-ee/hw3-group-3/pkg/repository"
	"github.com/cs-ut-ee/hw3-group-3/pkg/service"
	http2 "github.com/cs-ut-ee/hw3-group-3/pkg/transport/http"

	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const httpServicePort = 8081

func main() {

	log.Println("Server Started!")

	gormDB := getGormDB()
	repo := repository.NewPostgresRepo(gormDB)
	service := service.NewService(&repo)

	runHTTPServer(service)

	log.Infof("Server Stopped!")
}

func getPostgresConnection() string {
	postgresConnection, success := os.LookupEnv("postgresConnectionString")
	if !success {
		postgresConnection = "dbname=postgres host=localhost password=postgres user=postgres sslmode=disable port=55554"
	}

	return postgresConnection
}

func getGormDB() *gorm.DB {
	postgresConnection := getPostgresConnection()

	sqlDB, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}

func runHTTPServer(service service.Service) {

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
		sqlDB, err := (service.Repository.(*repository.PostgresRepo)).Database.DB()
		if err != nil {
			log.Fatalf("Could not get DB to close")
		}

		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Could not disconnect from Postgres DB")
		}
		log.Fatalf("Could not start server")
	}
}
