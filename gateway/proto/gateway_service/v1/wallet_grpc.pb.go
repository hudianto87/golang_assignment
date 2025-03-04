// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/gateway_service/v1/wallet.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WalletService_GetWallets_FullMethodName    = "/proto.wallet_service.v1.WalletService/GetWallets"
	WalletService_GetWalletByID_FullMethodName = "/proto.wallet_service.v1.WalletService/GetWalletByID"
	WalletService_CreateWallet_FullMethodName  = "/proto.wallet_service.v1.WalletService/CreateWallet"
	WalletService_UpdateWallet_FullMethodName  = "/proto.wallet_service.v1.WalletService/UpdateWallet"
	WalletService_DeleteWallet_FullMethodName  = "/proto.wallet_service.v1.WalletService/DeleteWallet"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetWallets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWalletResponse, error)
	GetWalletByID(ctx context.Context, in *GetWalletByIDRequest, opts ...grpc.CallOption) (*GetWalletByIDResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetWallets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletResponse)
	err := c.cc.Invoke(ctx, WalletService_GetWallets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetWalletByID(ctx context.Context, in *GetWalletByIDRequest, opts ...grpc.CallOption) (*GetWalletByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletByIDResponse)
	err := c.cc.Invoke(ctx, WalletService_GetWalletByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Wallet)
	err := c.cc.Invoke(ctx, WalletService_CreateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Wallet)
	err := c.cc.Invoke(ctx, WalletService_UpdateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Wallet)
	err := c.cc.Invoke(ctx, WalletService_DeleteWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility.
type WalletServiceServer interface {
	GetWallets(context.Context, *emptypb.Empty) (*GetWalletResponse, error)
	GetWalletByID(context.Context, *GetWalletByIDRequest) (*GetWalletByIDResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error)
	UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error)
	DeleteWallet(context.Context, *DeleteWalletRequest) (*Wallet, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) GetWallets(context.Context, *emptypb.Empty) (*GetWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallets not implemented")
}
func (UnimplementedWalletServiceServer) GetWalletByID(context.Context, *GetWalletByIDRequest) (*GetWalletByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletByID not implemented")
}
func (UnimplementedWalletServiceServer) CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServiceServer) UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWallet not implemented")
}
func (UnimplementedWalletServiceServer) DeleteWallet(context.Context, *DeleteWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWallet not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}
func (UnimplementedWalletServiceServer) testEmbeddedByValue()                       {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWallets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWallets(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWalletByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWalletByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWalletByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWalletByID(ctx, req.(*GetWalletByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_UpdateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UpdateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_UpdateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UpdateWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_DeleteWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).DeleteWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_DeleteWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).DeleteWallet(ctx, req.(*DeleteWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.wallet_service.v1.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWallets",
			Handler:    _WalletService_GetWallets_Handler,
		},
		{
			MethodName: "GetWalletByID",
			Handler:    _WalletService_GetWalletByID_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _WalletService_CreateWallet_Handler,
		},
		{
			MethodName: "UpdateWallet",
			Handler:    _WalletService_UpdateWallet_Handler,
		},
		{
			MethodName: "DeleteWallet",
			Handler:    _WalletService_DeleteWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gateway_service/v1/wallet.proto",
}
