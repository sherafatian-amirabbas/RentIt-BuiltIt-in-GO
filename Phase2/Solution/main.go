package main

import (
	"github.com/cs-ut-ee/hw2-group-3/pkg/repository"
	"github.com/cs-ut-ee/hw2-group-3/pkg/service"
	http2 "github.com/cs-ut-ee/hw2-group-3/pkg/transport/http"
	"github.com/cs-ut-ee/hw2-group-3/pkg/transport/websocket"

	"database/sql"

	"fmt"
	"net"
	"net/http"
	"os"

	grpc2 "github.com/cs-ut-ee/hw2-group-3/pkg/transport/gRPC"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	httpServicePort = 8080
	GRPCServicePort = 50051
)

func main() {

	log.Println("Server Started!")

	service := createAndGetTheService()

	go func() {
		runGRPCServer(service)
	}()

	runHTTPServer(service)

	log.Infof("Server Stopped!")
}

func createAndGetTheService() *service.Service {
	postgresConnection, success := os.LookupEnv("postgresConnectionString")
	if !success {
		postgresConnection = "dbname=postgres host=localhost password=password user=postgres sslmode=disable port=5432"
	}

	mongoConnection, success := os.LookupEnv("mongoConnectionString")
	if !success {
		mongoConnection = "mongodb://mongo:mongo@localhost:27017"
	}

	// open Postgres connection
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	plantPostgresRepo := repository.NewPlantPostgresRepo(dbConn)
	plantPostgresService := service.NewPlantPostgresService(plantPostgresRepo)

	dbConnMongo, err := mongo.NewClient(options.Client().ApplyURI(mongoConnection))
	if err != nil {
		log.Fatal(err)
	}
	err = dbConnMongo.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	plantMongoRepo := repository.NewPlantMongoRepo(dbConnMongo)
	plantMongoService := service.NewPlantMongoService(plantMongoRepo)

	return service.NewService(plantPostgresService, plantMongoService)
}

func runGRPCServer(service *service.Service) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", GRPCServicePort))
	if err != nil {
		log.Fatalf("GRPCService: failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := grpc2.NewGRPCService(service)
	grpc2.RegisterGRPCServiceServer(grpcServer, grpcService)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer: Failed to start server: %v", err)
	}
}

func runHTTPServer(service *service.Service) {

	httpRouter := mux.NewRouter().StrictSlash(true)
	plantHTTPHandler := http2.NewPlantHandler(service, httpRouter)
	plantHTTPHandler.RegisterRoutes()

	websocketHandler := websocket.NewWebsocketHandler(service, httpRouter)
	websocketHandler.RegisterRoutes()

	// setup http server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}

	err := httpSrv.ListenAndServe()
	if err != nil {
		err = service.PostgresService.PlantPostgresRepo.DB.Close()
		if err != nil {
			log.Fatalf("Could not disconnect from Postgres DB")
		}
		log.Fatalf("Could not start server")
	}
}
