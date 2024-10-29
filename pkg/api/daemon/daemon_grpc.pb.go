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
)

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	Dump(ctx context.Context, in *DumpReq, opts ...grpc.CallOption) (*DumpResp, error)
	Restore(ctx context.Context, in *RestoreReq, opts ...grpc.CallOption) (*RestoreResp, error)
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

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility.
type DaemonServer interface {
	Dump(context.Context, *DumpReq) (*DumpResp, error)
	Restore(context.Context, *RestoreReq) (*RestoreResp, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon.proto",
}
