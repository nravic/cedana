// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: daemon.proto

package daemon

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
	Daemon_Dump_FullMethodName    = "/cedana.daemon.Daemon/Dump"
	Daemon_Restore_FullMethodName = "/cedana.daemon.Daemon/Restore"
	Daemon_Start_FullMethodName   = "/cedana.daemon.Daemon/Start"
	Daemon_Manage_FullMethodName  = "/cedana.daemon.Daemon/Manage"
	Daemon_List_FullMethodName    = "/cedana.daemon.Daemon/List"
)

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	// C/R
	Dump(ctx context.Context, in *DumpReq, opts ...grpc.CallOption) (*DumpResp, error)
	Restore(ctx context.Context, in *RestoreReq, opts ...grpc.CallOption) (*RestoreResp, error)
	// Job management
	Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*StartResp, error)
	Manage(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*ManageResp, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Dump(ctx context.Context, in *DumpReq, opts ...grpc.CallOption) (*DumpResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DumpResp)
	err := c.cc.Invoke(ctx, Daemon_Dump_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Restore(ctx context.Context, in *RestoreReq, opts ...grpc.CallOption) (*RestoreResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestoreResp)
	err := c.cc.Invoke(ctx, Daemon_Restore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*StartResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartResp)
	err := c.cc.Invoke(ctx, Daemon_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Manage(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*ManageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManageResp)
	err := c.cc.Invoke(ctx, Daemon_Manage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResp)
	err := c.cc.Invoke(ctx, Daemon_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility.
type DaemonServer interface {
	// C/R
	Dump(context.Context, *DumpReq) (*DumpResp, error)
	Restore(context.Context, *RestoreReq) (*RestoreResp, error)
	// Job management
	Start(context.Context, *StartReq) (*StartResp, error)
	Manage(context.Context, *ManageReq) (*ManageResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDaemonServer struct{}

func (UnimplementedDaemonServer) Dump(context.Context, *DumpReq) (*DumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (UnimplementedDaemonServer) Restore(context.Context, *RestoreReq) (*RestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedDaemonServer) Start(context.Context, *StartReq) (*StartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedDaemonServer) Manage(context.Context, *ManageReq) (*ManageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Manage not implemented")
}
func (UnimplementedDaemonServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}
func (UnimplementedDaemonServer) testEmbeddedByValue()                {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	// If the following call pancis, it indicates UnimplementedDaemonServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Dump_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Dump(ctx, req.(*DumpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Restore(ctx, req.(*RestoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Start(ctx, req.(*StartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Manage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Manage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Manage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Manage(ctx, req.(*ManageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cedana.daemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dump",
			Handler:    _Daemon_Dump_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Daemon_Restore_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Daemon_Start_Handler,
		},
		{
			MethodName: "Manage",
			Handler:    _Daemon_Manage_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Daemon_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon.proto",
}