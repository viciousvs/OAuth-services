// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tokenService.proto

package __

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	GenerateTokens(ctx context.Context, in *GenerateTokensRequest, opts ...grpc.CallOption) (*Tokens, error)
	ValidateAccess(ctx context.Context, in *ValidateAccessRequest, opts ...grpc.CallOption) (*ValidateAccessResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Tokens, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) GenerateTokens(ctx context.Context, in *GenerateTokensRequest, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/tokenService.TokenService/GenerateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ValidateAccess(ctx context.Context, in *ValidateAccessRequest, opts ...grpc.CallOption) (*ValidateAccessResponse, error) {
	out := new(ValidateAccessResponse)
	err := c.cc.Invoke(ctx, "/tokenService.TokenService/ValidateAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/tokenService.TokenService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	GenerateTokens(context.Context, *GenerateTokensRequest) (*Tokens, error)
	ValidateAccess(context.Context, *ValidateAccessRequest) (*ValidateAccessResponse, error)
	Refresh(context.Context, *RefreshRequest) (*Tokens, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) GenerateTokens(context.Context, *GenerateTokensRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTokens not implemented")
}
func (UnimplementedTokenServiceServer) ValidateAccess(context.Context, *ValidateAccessRequest) (*ValidateAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAccess not implemented")
}
func (UnimplementedTokenServiceServer) Refresh(context.Context, *RefreshRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_GenerateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GenerateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenService.TokenService/GenerateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GenerateTokens(ctx, req.(*GenerateTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ValidateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ValidateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenService.TokenService/ValidateAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ValidateAccess(ctx, req.(*ValidateAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenService.TokenService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenService.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTokens",
			Handler:    _TokenService_GenerateTokens_Handler,
		},
		{
			MethodName: "ValidateAccess",
			Handler:    _TokenService_ValidateAccess_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _TokenService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenService.proto",
}
