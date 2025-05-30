// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: net.proto

package netsend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MyService_TaskStream_FullMethodName = "/netsend.MyService/TaskStream"
)

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	// 双向流：客户端和服务端都可以持续发送消息
	TaskStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskMessage, TaskNotification], error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) TaskStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskMessage, TaskNotification], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], MyService_TaskStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TaskMessage, TaskNotification]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_TaskStreamClient = grpc.BidiStreamingClient[TaskMessage, TaskNotification]

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility.
type MyServiceServer interface {
	// 双向流：客户端和服务端都可以持续发送消息
	TaskStream(grpc.BidiStreamingServer[TaskMessage, TaskNotification]) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMyServiceServer struct{}

func (UnimplementedMyServiceServer) TaskStream(grpc.BidiStreamingServer[TaskMessage, TaskNotification]) error {
	return status.Errorf(codes.Unimplemented, "method TaskStream not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}
func (UnimplementedMyServiceServer) testEmbeddedByValue()                   {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	// If the following call pancis, it indicates UnimplementedMyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_TaskStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).TaskStream(&grpc.GenericServerStream[TaskMessage, TaskNotification]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_TaskStreamServer = grpc.BidiStreamingServer[TaskMessage, TaskNotification]

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netsend.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskStream",
			Handler:       _MyService_TaskStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "net.proto",
}
