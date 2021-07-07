package grpc

import (
	"context"
	"time"

	"github.com/cs-ut-ee/hw2-group-3/pkg/service"
)

type gRPCServer struct {
	Service service.IService
}

func NewGRPCService(service service.IService) *gRPCServer {
	return &gRPCServer{
		Service: service,
	}
}

func (server *gRPCServer) GetAllPlants(_ context.Context, r *PlantsRequest) (*PlantsResponse, error) {
	plants, err := server.Service.GetAllPlants()
	if err != nil {
		return nil, err
	}

	var items []*Plant
	for _, plant := range plants {

		item := &Plant{
			Id:          plant.Id,
			Name:        plant.Name,
			Description: plant.Description,
			PricePerDay: plant.PricePerDay,
		}

		items = append(items, item)
	}

	res := &PlantsResponse{
		Items: items,
	}
	return res, nil
}

func (server *gRPCServer) GetPlantPrice(_ context.Context, r *PlantPriceRequest) (*PlantPriceResponse, error) {

	from, err := time.Parse("2006-01-02", r.StartDate)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse("2006-01-02", r.EndDate)
	if err != nil {
		return nil, err
	}

	plantPrice, err := server.Service.GetPlantPrice(r.Id, from, to)
	if err != nil {
		return nil, err
	}

	res := &PlantPriceResponse{
		Item: &PlantPrice{
			PlantId:          plantPrice.PlantId,
			StartDate:        r.StartDate,
			EndDate:          r.EndDate,
			PricePerDuration: plantPrice.PricePerDuration,
		},
	}
	return res, nil
}

func (server *gRPCServer) IsPlantAvailable(_ context.Context, r *PlantAvailabilityRequest) (*PlantAvailabilityResponse, error) {
	from, err := time.Parse("2006-01-02", r.StartDate)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse("2006-01-02", r.EndDate)
	if err != nil {
		return nil, err
	}

	availability, err := server.Service.IsPlantAvailable(r.Id, from, to)
	if err != nil {
		return nil, err
	}

	res := &PlantAvailabilityResponse{
		IsAvailable: availability,
	}
	return res, nil
}
