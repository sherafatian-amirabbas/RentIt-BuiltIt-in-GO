package main

import (
	grpc2 "github.com/cs-ut-ee/hw2-group-3/pkg/transport/gRPC"

	"context"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func ApiUrl() string {
	url, success := os.LookupEnv("grpcUrl")
	if !success {
		panic("Environment variable 'grpcUrl' is not defined")
	}

	return url
}

func TestGetAllPlants(t *testing.T) {

	conn, err := grpc.Dial(ApiUrl(), grpc.WithInsecure())
	if err != nil {
		t.Error("TestGetAllPlants: Problem establishing the connection via gRPC")
		return
	}

	defer conn.Close()
	client := grpc2.NewGRPCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	plantsRequest := &grpc2.PlantsRequest{}
	plantsResponse, err := client.GetAllPlants(ctx, plantsRequest)
	if err != nil {
		t.Error("TestGetAllPlants: Problem receiving the reponse via gRPC")
		return
	}

	if len(plantsResponse.Items) < 2 { // since we have 2 records initially in postgress and it's not going to be chnaged
		t.Error("TestGetAllPlants: there should 2 plants available!")
		return
	}
}

func TestGetPlantPrice(t *testing.T) {

	conn, err := grpc.Dial(ApiUrl(), grpc.WithInsecure())
	if err != nil {
		t.Error("TestGetPlantPrice: Problem establishing the connection via gRPC")
		return
	}

	defer conn.Close()
	client := grpc2.NewGRPCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	plantPriceRequest := &grpc2.PlantPriceRequest{
		Id:        1,
		StartDate: "2020-03-10",
		EndDate:   "2020-03-12",
	}
	responsePlantPrice, err := client.GetPlantPrice(ctx, plantPriceRequest)
	if err != nil {
		t.Error("TestGetPlantPrice: Problem receiving the reponse via gRPC")
		return
	}

	if responsePlantPrice.Item.PricePerDuration != 21 { // for 2 days this is the price, this should be in postgres
		t.Error("TestGetPlantPrice: there is something wrong with calculating the price!")
		return
	}
}

func TestIsPlantAvailable(t *testing.T) {

	conn, err := grpc.Dial(ApiUrl(), grpc.WithInsecure())
	if err != nil {
		t.Error("TestIsPlantAvailable: Problem establishing the connection via gRPC")
		return
	}

	defer conn.Close()
	client := grpc2.NewGRPCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	plantAvailabilityRequest := &grpc2.PlantAvailabilityRequest{
		Id:        1,
		StartDate: "2020-03-21",
		EndDate:   "2020-03-22",
	}
	plantAvailabilityResponse, err := client.IsPlantAvailable(ctx, plantAvailabilityRequest)
	if err != nil {
		t.Error("TestIsPlantAvailable: Problem receiving the reponse via gRPC")
		return
	}

	if plantAvailabilityResponse.IsAvailable { // this item in the postgres has order in this time period (it's created when db is initialized)
		t.Error("TestIsPlantAvailable: the plant should not be available!")
		return
	}
}
