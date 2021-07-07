// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GRPCServiceClient is the client API for GRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCServiceClient interface {
	GetAllPlants(ctx context.Context, in *PlantsRequest, opts ...grpc.CallOption) (*PlantsResponse, error)
	GetPlantPrice(ctx context.Context, in *PlantPriceRequest, opts ...grpc.CallOption) (*PlantPriceResponse, error)
	IsPlantAvailable(ctx context.Context, in *PlantAvailabilityRequest, opts ...grpc.CallOption) (*PlantAvailabilityResponse, error)
}

type gRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCServiceClient(cc grpc.ClientConnInterface) GRPCServiceClient {
	return &gRPCServiceClient{cc}
}

func (c *gRPCServiceClient) GetAllPlants(ctx context.Context, in *PlantsRequest, opts ...grpc.CallOption) (*PlantsResponse, error) {
	out := new(PlantsResponse)
	err := c.cc.Invoke(ctx, "/grpc.gRPCService/GetAllPlants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCServiceClient) GetPlantPrice(ctx context.Context, in *PlantPriceRequest, opts ...grpc.CallOption) (*PlantPriceResponse, error) {
	out := new(PlantPriceResponse)
	err := c.cc.Invoke(ctx, "/grpc.gRPCService/GetPlantPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCServiceClient) IsPlantAvailable(ctx context.Context, in *PlantAvailabilityRequest, opts ...grpc.CallOption) (*PlantAvailabilityResponse, error) {
	out := new(PlantAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/grpc.gRPCService/IsPlantAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServiceServer is the server API for GRPCService service.
// All implementations should embed UnimplementedGRPCServiceServer
// for forward compatibility
type GRPCServiceServer interface {
	GetAllPlants(context.Context, *PlantsRequest) (*PlantsResponse, error)
	GetPlantPrice(context.Context, *PlantPriceRequest) (*PlantPriceResponse, error)
	IsPlantAvailable(context.Context, *PlantAvailabilityRequest) (*PlantAvailabilityResponse, error)
}

// UnimplementedGRPCServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGRPCServiceServer struct {
}

func (UnimplementedGRPCServiceServer) GetAllPlants(context.Context, *PlantsRequest) (*PlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlants not implemented")
}
func (UnimplementedGRPCServiceServer) GetPlantPrice(context.Context, *PlantPriceRequest) (*PlantPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlantPrice not implemented")
}
func (UnimplementedGRPCServiceServer) IsPlantAvailable(context.Context, *PlantAvailabilityRequest) (*PlantAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPlantAvailable not implemented")
}

// UnsafeGRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCServiceServer will
// result in compilation errors.
type UnsafeGRPCServiceServer interface {
	mustEmbedUnimplementedGRPCServiceServer()
}

func RegisterGRPCServiceServer(s grpc.ServiceRegistrar, srv GRPCServiceServer) {
	s.RegisterService(&GRPCService_ServiceDesc, srv)
}

func _GRPCService_GetAllPlants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).GetAllPlants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gRPCService/GetAllPlants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).GetAllPlants(ctx, req.(*PlantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCService_GetPlantPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).GetPlantPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gRPCService/GetPlantPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).GetPlantPrice(ctx, req.(*PlantPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCService_IsPlantAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).IsPlantAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gRPCService/IsPlantAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).IsPlantAvailable(ctx, req.(*PlantAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCService_ServiceDesc is the grpc.ServiceDesc for GRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gRPCService",
	HandlerType: (*GRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPlants",
			Handler:    _GRPCService_GetAllPlants_Handler,
		},
		{
			MethodName: "GetPlantPrice",
			Handler:    _GRPCService_GetPlantPrice_Handler,
		},
		{
			MethodName: "IsPlantAvailable",
			Handler:    _GRPCService_IsPlantAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/transport/gRPC/gRPCService.proto",
}
