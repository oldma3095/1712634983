// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: common/api.proto

package api

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

const (
	ApiService_ClientInfo_FullMethodName = "/common.ApiService/ClientInfo"
	ApiService_Result_FullMethodName     = "/common.ApiService/Result"
)

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	ClientInfo(ctx context.Context, opts ...grpc.CallOption) (ApiService_ClientInfoClient, error)
	Result(ctx context.Context, in *ResultReq, opts ...grpc.CallOption) (*ResultRes, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) ClientInfo(ctx context.Context, opts ...grpc.CallOption) (ApiService_ClientInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiService_ServiceDesc.Streams[0], ApiService_ClientInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &apiServiceClientInfoClient{stream}
	return x, nil
}

type ApiService_ClientInfoClient interface {
	Send(*ClientInfoReq) error
	CloseAndRecv() (*ClientInfoRes, error)
	grpc.ClientStream
}

type apiServiceClientInfoClient struct {
	grpc.ClientStream
}

func (x *apiServiceClientInfoClient) Send(m *ClientInfoReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *apiServiceClientInfoClient) CloseAndRecv() (*ClientInfoRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientInfoRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiServiceClient) Result(ctx context.Context, in *ResultReq, opts ...grpc.CallOption) (*ResultRes, error) {
	out := new(ResultRes)
	err := c.cc.Invoke(ctx, ApiService_Result_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	ClientInfo(ApiService_ClientInfoServer) error
	Result(context.Context, *ResultReq) (*ResultRes, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) ClientInfo(ApiService_ClientInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientInfo not implemented")
}
func (UnimplementedApiServiceServer) Result(context.Context, *ResultReq) (*ResultRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_ClientInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApiServiceServer).ClientInfo(&apiServiceClientInfoServer{stream})
}

type ApiService_ClientInfoServer interface {
	SendAndClose(*ClientInfoRes) error
	Recv() (*ClientInfoReq, error)
	grpc.ServerStream
}

type apiServiceClientInfoServer struct {
	grpc.ServerStream
}

func (x *apiServiceClientInfoServer) SendAndClose(m *ClientInfoRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *apiServiceClientInfoServer) Recv() (*ClientInfoReq, error) {
	m := new(ClientInfoReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApiService_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_Result_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Result(ctx, req.(*ResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Result",
			Handler:    _ApiService_Result_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientInfo",
			Handler:       _ApiService_ClientInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "common/api.proto",
}
