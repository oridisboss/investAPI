// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package investapi

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

// OperationsServiceClient is the client API for OperationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsServiceClient interface {
	//Метод получения списка операций по счёту.
	GetOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
	//Метод получения портфеля по счёту.
	GetPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error)
	//Метод получения списка позиций по счёту.
	GetPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error)
	//Метод получения доступного остатка для вывода средств.
	GetWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error)
}

type operationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsServiceClient(cc grpc.ClientConnInterface) OperationsServiceClient {
	return &operationsServiceClient{cc}
}

func (c *operationsServiceClient) GetOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error) {
	out := new(PortfolioResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error) {
	out := new(PositionsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error) {
	out := new(WithdrawLimitsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServiceServer is the server API for OperationsService service.
// All implementations must embed UnimplementedOperationsServiceServer
// for forward compatibility
type OperationsServiceServer interface {
	//Метод получения списка операций по счёту.
	GetOperations(context.Context, *OperationsRequest) (*OperationsResponse, error)
	//Метод получения портфеля по счёту.
	GetPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error)
	//Метод получения списка позиций по счёту.
	GetPositions(context.Context, *PositionsRequest) (*PositionsResponse, error)
	//Метод получения доступного остатка для вывода средств.
	GetWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error)
	mustEmbedUnimplementedOperationsServiceServer()
}

// UnimplementedOperationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsServiceServer struct {
}

func (UnimplementedOperationsServiceServer) GetOperations(context.Context, *OperationsRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedOperationsServiceServer) GetPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortfolio not implemented")
}
func (UnimplementedOperationsServiceServer) GetPositions(context.Context, *PositionsRequest) (*PositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedOperationsServiceServer) GetWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawLimits not implemented")
}
func (UnimplementedOperationsServiceServer) mustEmbedUnimplementedOperationsServiceServer() {}

// UnsafeOperationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsServiceServer will
// result in compilation errors.
type UnsafeOperationsServiceServer interface {
	mustEmbedUnimplementedOperationsServiceServer()
}

func RegisterOperationsServiceServer(s grpc.ServiceRegistrar, srv OperationsServiceServer) {
	s.RegisterService(&OperationsService_ServiceDesc, srv)
}

func _OperationsService_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetOperations(ctx, req.(*OperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetPortfolio(ctx, req.(*PortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetPositions(ctx, req.(*PositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetWithdrawLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetWithdrawLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetWithdrawLimits(ctx, req.(*WithdrawLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationsService_ServiceDesc is the grpc.ServiceDesc for OperationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.OperationsService",
	HandlerType: (*OperationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperations",
			Handler:    _OperationsService_GetOperations_Handler,
		},
		{
			MethodName: "GetPortfolio",
			Handler:    _OperationsService_GetPortfolio_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _OperationsService_GetPositions_Handler,
		},
		{
			MethodName: "GetWithdrawLimits",
			Handler:    _OperationsService_GetWithdrawLimits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations.proto",
}
