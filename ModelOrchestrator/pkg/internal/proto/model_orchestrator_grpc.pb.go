// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: model_orchestrator.proto

package proto

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
	ModelOrchestrator_SendData_FullMethodName = "/ModelOrchestrator.ModelOrchestrator/SendData"
)

// ModelOrchestratorClient is the client API for ModelOrchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelOrchestratorClient interface {
	SendData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ProcessedData, Status], error)
}

type modelOrchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewModelOrchestratorClient(cc grpc.ClientConnInterface) ModelOrchestratorClient {
	return &modelOrchestratorClient{cc}
}

func (c *modelOrchestratorClient) SendData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ProcessedData, Status], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ModelOrchestrator_ServiceDesc.Streams[0], ModelOrchestrator_SendData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ProcessedData, Status]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelOrchestrator_SendDataClient = grpc.ClientStreamingClient[ProcessedData, Status]

// ModelOrchestratorServer is the server API for ModelOrchestrator service.
// All implementations must embed UnimplementedModelOrchestratorServer
// for forward compatibility.
type ModelOrchestratorServer interface {
	SendData(grpc.ClientStreamingServer[ProcessedData, Status]) error
	mustEmbedUnimplementedModelOrchestratorServer()
}

// UnimplementedModelOrchestratorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModelOrchestratorServer struct{}

func (UnimplementedModelOrchestratorServer) SendData(grpc.ClientStreamingServer[ProcessedData, Status]) error {
	return status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedModelOrchestratorServer) mustEmbedUnimplementedModelOrchestratorServer() {}
func (UnimplementedModelOrchestratorServer) testEmbeddedByValue()                           {}

// UnsafeModelOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelOrchestratorServer will
// result in compilation errors.
type UnsafeModelOrchestratorServer interface {
	mustEmbedUnimplementedModelOrchestratorServer()
}

func RegisterModelOrchestratorServer(s grpc.ServiceRegistrar, srv ModelOrchestratorServer) {
	// If the following call pancis, it indicates UnimplementedModelOrchestratorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModelOrchestrator_ServiceDesc, srv)
}

func _ModelOrchestrator_SendData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelOrchestratorServer).SendData(&grpc.GenericServerStream[ProcessedData, Status]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelOrchestrator_SendDataServer = grpc.ClientStreamingServer[ProcessedData, Status]

// ModelOrchestrator_ServiceDesc is the grpc.ServiceDesc for ModelOrchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelOrchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ModelOrchestrator.ModelOrchestrator",
	HandlerType: (*ModelOrchestratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendData",
			Handler:       _ModelOrchestrator_SendData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "model_orchestrator.proto",
}
