// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shortener

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

// ShortenerServiceClient is the client API for ShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerServiceClient interface {
	CreateShortener(ctx context.Context, in *CreateShortenRequest, opts ...grpc.CallOption) (*CreateShortenResponse, error)
}

type shortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerServiceClient(cc grpc.ClientConnInterface) ShortenerServiceClient {
	return &shortenerServiceClient{cc}
}

func (c *shortenerServiceClient) CreateShortener(ctx context.Context, in *CreateShortenRequest, opts ...grpc.CallOption) (*CreateShortenResponse, error) {
	out := new(CreateShortenResponse)
	err := c.cc.Invoke(ctx, "/shortener.v1.ShortenerService/CreateShortener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServiceServer is the server API for ShortenerService service.
// All implementations should embed UnimplementedShortenerServiceServer
// for forward compatibility
type ShortenerServiceServer interface {
	CreateShortener(context.Context, *CreateShortenRequest) (*CreateShortenResponse, error)
}

// UnimplementedShortenerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShortenerServiceServer struct {
}

func (UnimplementedShortenerServiceServer) CreateShortener(context.Context, *CreateShortenRequest) (*CreateShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortener not implemented")
}

// UnsafeShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServiceServer will
// result in compilation errors.
type UnsafeShortenerServiceServer interface {
	mustEmbedUnimplementedShortenerServiceServer()
}

func RegisterShortenerServiceServer(s grpc.ServiceRegistrar, srv ShortenerServiceServer) {
	s.RegisterService(&ShortenerService_ServiceDesc, srv)
}

func _ShortenerService_CreateShortener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServiceServer).CreateShortener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.v1.ShortenerService/CreateShortener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServiceServer).CreateShortener(ctx, req.(*CreateShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenerService_ServiceDesc is the grpc.ServiceDesc for ShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.v1.ShortenerService",
	HandlerType: (*ShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortener",
			Handler:    _ShortenerService_CreateShortener_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrastructure/grpc/protofile/shortener/v1/shortener.proto",
}
