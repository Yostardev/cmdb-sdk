// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: auth/v1/group.proto

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
	GroupService_All_FullMethodName             = "/auth.GroupService/All"
	GroupService_List_FullMethodName            = "/auth.GroupService/List"
	GroupService_Get_FullMethodName             = "/auth.GroupService/Get"
	GroupService_Create_FullMethodName          = "/auth.GroupService/Create"
	GroupService_Update_FullMethodName          = "/auth.GroupService/Update"
	GroupService_Delete_FullMethodName          = "/auth.GroupService/Delete"
	GroupService_UpdateGroupRole_FullMethodName = "/auth.GroupService/UpdateGroupRole"
	GroupService_GetGroupUserIDs_FullMethodName = "/auth.GroupService/GetGroupUserIDs"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	All(ctx context.Context, in *GroupAllRequest, opts ...grpc.CallOption) (*GroupAllResponse, error)
	List(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
	Get(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	Create(ctx context.Context, in *GroupCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *GroupUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *GroupDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateGroupRole(ctx context.Context, in *UpdateGroupRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetGroupUserIDs(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupUserIDsResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) All(ctx context.Context, in *GroupAllRequest, opts ...grpc.CallOption) (*GroupAllResponse, error) {
	out := new(GroupAllResponse)
	err := c.cc.Invoke(ctx, GroupService_All_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) List(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := c.cc.Invoke(ctx, GroupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Get(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, GroupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Create(ctx context.Context, in *GroupCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Update(ctx context.Context, in *GroupUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Delete(ctx context.Context, in *GroupDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroupRole(ctx context.Context, in *UpdateGroupRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupService_UpdateGroupRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupUserIDs(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupUserIDsResponse, error) {
	out := new(GroupUserIDsResponse)
	err := c.cc.Invoke(ctx, GroupService_GetGroupUserIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	All(context.Context, *GroupAllRequest) (*GroupAllResponse, error)
	List(context.Context, *GroupListRequest) (*GroupListResponse, error)
	Get(context.Context, *GroupGetRequest) (*GroupResponse, error)
	Create(context.Context, *GroupCreateRequest) (*emptypb.Empty, error)
	Update(context.Context, *GroupUpdateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *GroupDeleteRequest) (*emptypb.Empty, error)
	UpdateGroupRole(context.Context, *UpdateGroupRoleRequest) (*emptypb.Empty, error)
	GetGroupUserIDs(context.Context, *GroupGetRequest) (*GroupUserIDsResponse, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) All(context.Context, *GroupAllRequest) (*GroupAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedGroupServiceServer) List(context.Context, *GroupListRequest) (*GroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGroupServiceServer) Get(context.Context, *GroupGetRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupServiceServer) Create(context.Context, *GroupCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupServiceServer) Update(context.Context, *GroupUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupServiceServer) Delete(context.Context, *GroupDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroupRole(context.Context, *UpdateGroupRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupRole not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupUserIDs(context.Context, *GroupGetRequest) (*GroupUserIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupUserIDs not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).All(ctx, req.(*GroupAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).List(ctx, req.(*GroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Get(ctx, req.(*GroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Create(ctx, req.(*GroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Update(ctx, req.(*GroupUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Delete(ctx, req.(*GroupDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroupRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroupRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_UpdateGroupRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroupRole(ctx, req.(*UpdateGroupRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_GetGroupUserIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupUserIDs(ctx, req.(*GroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _GroupService_All_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GroupService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GroupService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupService_Delete_Handler,
		},
		{
			MethodName: "UpdateGroupRole",
			Handler:    _GroupService_UpdateGroupRole_Handler,
		},
		{
			MethodName: "GetGroupUserIDs",
			Handler:    _GroupService_GetGroupUserIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/group.proto",
}
