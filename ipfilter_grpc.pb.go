// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ipfilter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// IpFilterClient is the client API for IpFilter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IpFilterClient interface {
	Rewrite(ctx context.Context, in *IpReq, opts ...grpc.CallOption) (*IpRep, error)
}

type ipFilterClient struct {
	cc grpc.ClientConnInterface
}

func NewIpFilterClient(cc grpc.ClientConnInterface) IpFilterClient {
	return &ipFilterClient{cc}
}

func (c *ipFilterClient) Rewrite(ctx context.Context, in *IpReq, opts ...grpc.CallOption) (*IpRep, error) {
	out := new(IpRep)
	err := c.cc.Invoke(ctx, "/ipfilter.IpFilter/Rewrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IpFilterServer is the server API for IpFilter service.
// All implementations must embed UnimplementedIpFilterServer
// for forward compatibility
type IpFilterServer interface {
	Rewrite(context.Context, *IpReq) (*IpRep, error)
	mustEmbedUnimplementedIpFilterServer()
}

// UnimplementedIpFilterServer must be embedded to have forward compatible implementations.
type UnimplementedIpFilterServer struct {
}

func (UnimplementedIpFilterServer) Rewrite(context.Context, *IpReq) (*IpRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rewrite not implemented")
}
func (UnimplementedIpFilterServer) mustEmbedUnimplementedIpFilterServer() {}

// UnsafeIpFilterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IpFilterServer will
// result in compilation errors.
type UnsafeIpFilterServer interface {
	mustEmbedUnimplementedIpFilterServer()
}

func RegisterIpFilterServer(s grpc.ServiceRegistrar, srv IpFilterServer) {
	s.RegisterService(&_IpFilter_serviceDesc, srv)
}

func _IpFilter_Rewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpFilterServer).Rewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfilter.IpFilter/Rewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpFilterServer).Rewrite(ctx, req.(*IpReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _IpFilter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ipfilter.IpFilter",
	HandlerType: (*IpFilterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rewrite",
			Handler:    _IpFilter_Rewrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipfilter.proto",
}