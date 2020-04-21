// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heldamount.proto

package pbgrpc

import (
	context "context"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"

	. "storj.io/common/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeldAmountClient is the client API for HeldAmount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeldAmountClient interface {
	GetPayStub(ctx context.Context, in *GetHeldAmountRequest, opts ...grpc.CallOption) (*GetHeldAmountResponse, error)
	GetAllPaystubs(ctx context.Context, in *GetAllPaystubsRequest, opts ...grpc.CallOption) (*GetAllPaystubsResponse, error)
}

type heldAmountClient struct {
	cc *grpc.ClientConn
}

func NewHeldAmountClient(cc *grpc.ClientConn) HeldAmountClient {
	return &heldAmountClient{cc}
}

func (c *heldAmountClient) GetPayStub(ctx context.Context, in *GetHeldAmountRequest, opts ...grpc.CallOption) (*GetHeldAmountResponse, error) {
	out := new(GetHeldAmountResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetPayStub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heldAmountClient) GetAllPaystubs(ctx context.Context, in *GetAllPaystubsRequest, opts ...grpc.CallOption) (*GetAllPaystubsResponse, error) {
	out := new(GetAllPaystubsResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetAllPaystubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeldAmountServer is the server API for HeldAmount service.
type HeldAmountServer interface {
	GetPayStub(context.Context, *GetHeldAmountRequest) (*GetHeldAmountResponse, error)
	GetAllPaystubs(context.Context, *GetAllPaystubsRequest) (*GetAllPaystubsResponse, error)
}

func RegisterHeldAmountServer(s *grpc.Server, srv HeldAmountServer) {
	s.RegisterService(&_HeldAmount_serviceDesc, srv)
}

func _HeldAmount_GetPayStub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeldAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeldAmountServer).GetPayStub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heldamount.HeldAmount/GetPayStub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeldAmountServer).GetPayStub(ctx, req.(*GetHeldAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeldAmount_GetAllPaystubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPaystubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeldAmountServer).GetAllPaystubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heldamount.HeldAmount/GetAllPaystubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeldAmountServer).GetAllPaystubs(ctx, req.(*GetAllPaystubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeldAmount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heldamount.HeldAmount",
	HandlerType: (*HeldAmountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPayStub",
			Handler:    _HeldAmount_GetPayStub_Handler,
		},
		{
			MethodName: "GetAllPaystubs",
			Handler:    _HeldAmount_GetAllPaystubs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heldamount.proto",
}
