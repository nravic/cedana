// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: agent-task.proto

package agent_task

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Health
	DetailedHealthCheck(ctx context.Context, in *DetailedHealthCheckRequest, opts ...grpc.CallOption) (*DetailedHealthCheckResponse, error)
	// Kata
	KataDump(ctx context.Context, in *DumpArgs, opts ...grpc.CallOption) (*DumpResp, error)
	KataRestore(ctx context.Context, in *RestoreArgs, opts ...grpc.CallOption) (*RestoreResp, error)
	// Config
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) DetailedHealthCheck(ctx context.Context, in *DetailedHealthCheckRequest, opts ...grpc.CallOption) (*DetailedHealthCheckResponse, error) {
	out := new(DetailedHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/cedana.services.agent_task.TaskService/DetailedHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) KataDump(ctx context.Context, in *DumpArgs, opts ...grpc.CallOption) (*DumpResp, error) {
	out := new(DumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.agent_task.TaskService/KataDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) KataRestore(ctx context.Context, in *RestoreArgs, opts ...grpc.CallOption) (*RestoreResp, error) {
	out := new(RestoreResp)
	err := c.cc.Invoke(ctx, "/cedana.services.agent_task.TaskService/KataRestore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/cedana.services.agent_task.TaskService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// Health
	DetailedHealthCheck(context.Context, *DetailedHealthCheckRequest) (*DetailedHealthCheckResponse, error)
	// Kata
	KataDump(context.Context, *DumpArgs) (*DumpResp, error)
	KataRestore(context.Context, *RestoreArgs) (*RestoreResp, error)
	// Config
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) DetailedHealthCheck(context.Context, *DetailedHealthCheckRequest) (*DetailedHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailedHealthCheck not implemented")
}
func (UnimplementedTaskServiceServer) KataDump(context.Context, *DumpArgs) (*DumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KataDump not implemented")
}
func (UnimplementedTaskServiceServer) KataRestore(context.Context, *RestoreArgs) (*RestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KataRestore not implemented")
}
func (UnimplementedTaskServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_DetailedHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailedHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DetailedHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.agent_task.TaskService/DetailedHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DetailedHealthCheck(ctx, req.(*DetailedHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_KataDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).KataDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.agent_task.TaskService/KataDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).KataDump(ctx, req.(*DumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_KataRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).KataRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.agent_task.TaskService/KataRestore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).KataRestore(ctx, req.(*RestoreArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.agent_task.TaskService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cedana.services.agent_task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetailedHealthCheck",
			Handler:    _TaskService_DetailedHealthCheck_Handler,
		},
		{
			MethodName: "KataDump",
			Handler:    _TaskService_KataDump_Handler,
		},
		{
			MethodName: "KataRestore",
			Handler:    _TaskService_KataRestore_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _TaskService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent-task.proto",
}