// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: auth/v1/user.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_List_FullMethodName             = "/auth.UserService/List"
	UserService_Get_FullMethodName              = "/auth.UserService/Get"
	UserService_GetSelf_FullMethodName          = "/auth.UserService/GetSelf"
	UserService_All_FullMethodName              = "/auth.UserService/All"
	UserService_UpdateUserGroups_FullMethodName = "/auth.UserService/UpdateUserGroups"
	UserService_UpdateUserRoles_FullMethodName  = "/auth.UserService/UpdateUserRoles"
	UserService_LoginByLDAP_FullMethodName      = "/auth.UserService/LoginByLDAP"
	UserService_LoginByIdaas_FullMethodName     = "/auth.UserService/LoginByIdaas"
	UserService_Sync_FullMethodName             = "/auth.UserService/Sync"
	UserService_Ban_FullMethodName              = "/auth.UserService/Ban"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetSelf(ctx context.Context, in *UserGetSelfRequest, opts ...grpc.CallOption) (*UserResponse, error)
	All(ctx context.Context, in *UserAllRequest, opts ...grpc.CallOption) (*UserAllResponse, error)
	UpdateUserGroups(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserRoles(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LoginByLDAP(ctx context.Context, in *LoginByLDAPRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LoginByIdaas(ctx context.Context, in *LoginByIdaasRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Sync(ctx context.Context, in *SyncUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Ban(ctx context.Context, in *UserBanRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, UserService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSelf(ctx context.Context, in *UserGetSelfRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetSelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) All(ctx context.Context, in *UserAllRequest, opts ...grpc.CallOption) (*UserAllResponse, error) {
	out := new(UserAllResponse)
	err := c.cc.Invoke(ctx, UserService_All_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserGroups(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateUserGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserRoles(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateUserRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginByLDAP(ctx context.Context, in *LoginByLDAPRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_LoginByLDAP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginByIdaas(ctx context.Context, in *LoginByIdaasRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_LoginByIdaas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Sync(ctx context.Context, in *SyncUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Ban(ctx context.Context, in *UserBanRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_Ban_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	List(context.Context, *UserListRequest) (*UserListResponse, error)
	Get(context.Context, *UserGetRequest) (*UserResponse, error)
	GetSelf(context.Context, *UserGetSelfRequest) (*UserResponse, error)
	All(context.Context, *UserAllRequest) (*UserAllResponse, error)
	UpdateUserGroups(context.Context, *UpdateUserGroupRequest) (*emptypb.Empty, error)
	UpdateUserRoles(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error)
	LoginByLDAP(context.Context, *LoginByLDAPRequest) (*LoginResponse, error)
	LoginByIdaas(context.Context, *LoginByIdaasRequest) (*LoginResponse, error)
	Sync(context.Context, *SyncUserRequest) (*emptypb.Empty, error)
	Ban(context.Context, *UserBanRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) List(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServiceServer) Get(context.Context, *UserGetRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) GetSelf(context.Context, *UserGetSelfRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelf not implemented")
}
func (UnimplementedUserServiceServer) All(context.Context, *UserAllRequest) (*UserAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserGroups(context.Context, *UpdateUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroups not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserRoles(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRoles not implemented")
}
func (UnimplementedUserServiceServer) LoginByLDAP(context.Context, *LoginByLDAPRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByLDAP not implemented")
}
func (UnimplementedUserServiceServer) LoginByIdaas(context.Context, *LoginByIdaasRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByIdaas not implemented")
}
func (UnimplementedUserServiceServer) Sync(context.Context, *SyncUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedUserServiceServer) Ban(context.Context, *UserBanRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetSelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetSelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSelf(ctx, req.(*UserGetSelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).All(ctx, req.(*UserAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserGroups(ctx, req.(*UpdateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserRoles(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginByLDAP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByLDAPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginByLDAP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginByLDAP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginByLDAP(ctx, req.(*LoginByLDAPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginByIdaas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByIdaasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginByIdaas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginByIdaas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginByIdaas(ctx, req.(*LoginByIdaasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Sync(ctx, req.(*SyncUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Ban_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Ban(ctx, req.(*UserBanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetSelf",
			Handler:    _UserService_GetSelf_Handler,
		},
		{
			MethodName: "All",
			Handler:    _UserService_All_Handler,
		},
		{
			MethodName: "UpdateUserGroups",
			Handler:    _UserService_UpdateUserGroups_Handler,
		},
		{
			MethodName: "UpdateUserRoles",
			Handler:    _UserService_UpdateUserRoles_Handler,
		},
		{
			MethodName: "LoginByLDAP",
			Handler:    _UserService_LoginByLDAP_Handler,
		},
		{
			MethodName: "LoginByIdaas",
			Handler:    _UserService_LoginByIdaas_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _UserService_Sync_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _UserService_Ban_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/user.proto",
}
