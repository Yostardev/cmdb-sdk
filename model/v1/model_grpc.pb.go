// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: model/v1/model.proto

package model

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
	ModelService_All_FullMethodName       = "/model.ModelService/All"
	ModelService_List_FullMethodName      = "/model.ModelService/List"
	ModelService_Get_FullMethodName       = "/model.ModelService/Get"
	ModelService_GetByName_FullMethodName = "/model.ModelService/GetByName"
	ModelService_Create_FullMethodName    = "/model.ModelService/Create"
	ModelService_Update_FullMethodName    = "/model.ModelService/Update"
	ModelService_Delete_FullMethodName    = "/model.ModelService/Delete"
)

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	All(ctx context.Context, in *ModelAllRequest, opts ...grpc.CallOption) (*ModelAllResponse, error)
	List(ctx context.Context, in *ModelListRequest, opts ...grpc.CallOption) (*ModelListResponse, error)
	Get(ctx context.Context, in *ModelGetRequest, opts ...grpc.CallOption) (*ModelResponse, error)
	GetByName(ctx context.Context, in *ModelGetByNameRequest, opts ...grpc.CallOption) (*ModelResponse, error)
	Create(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *ModelUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ModelDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) All(ctx context.Context, in *ModelAllRequest, opts ...grpc.CallOption) (*ModelAllResponse, error) {
	out := new(ModelAllResponse)
	err := c.cc.Invoke(ctx, ModelService_All_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) List(ctx context.Context, in *ModelListRequest, opts ...grpc.CallOption) (*ModelListResponse, error) {
	out := new(ModelListResponse)
	err := c.cc.Invoke(ctx, ModelService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) Get(ctx context.Context, in *ModelGetRequest, opts ...grpc.CallOption) (*ModelResponse, error) {
	out := new(ModelResponse)
	err := c.cc.Invoke(ctx, ModelService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) GetByName(ctx context.Context, in *ModelGetByNameRequest, opts ...grpc.CallOption) (*ModelResponse, error) {
	out := new(ModelResponse)
	err := c.cc.Invoke(ctx, ModelService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) Create(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) Update(ctx context.Context, in *ModelUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) Delete(ctx context.Context, in *ModelDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServiceServer is the server API for ModelService service.
// All implementations must embed UnimplementedModelServiceServer
// for forward compatibility
type ModelServiceServer interface {
	All(context.Context, *ModelAllRequest) (*ModelAllResponse, error)
	List(context.Context, *ModelListRequest) (*ModelListResponse, error)
	Get(context.Context, *ModelGetRequest) (*ModelResponse, error)
	GetByName(context.Context, *ModelGetByNameRequest) (*ModelResponse, error)
	Create(context.Context, *ModelCreateRequest) (*emptypb.Empty, error)
	Update(context.Context, *ModelUpdateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *ModelDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedModelServiceServer()
}

// UnimplementedModelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModelServiceServer struct {
}

func (UnimplementedModelServiceServer) All(context.Context, *ModelAllRequest) (*ModelAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedModelServiceServer) List(context.Context, *ModelListRequest) (*ModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedModelServiceServer) Get(context.Context, *ModelGetRequest) (*ModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedModelServiceServer) GetByName(context.Context, *ModelGetByNameRequest) (*ModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedModelServiceServer) Create(context.Context, *ModelCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedModelServiceServer) Update(context.Context, *ModelUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedModelServiceServer) Delete(context.Context, *ModelDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedModelServiceServer) mustEmbedUnimplementedModelServiceServer() {}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).All(ctx, req.(*ModelAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).List(ctx, req.(*ModelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Get(ctx, req.(*ModelGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetByName(ctx, req.(*ModelGetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Create(ctx, req.(*ModelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Update(ctx, req.(*ModelUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Delete(ctx, req.(*ModelDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _ModelService_All_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ModelService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ModelService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _ModelService_GetByName_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ModelService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ModelService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ModelService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/v1/model.proto",
}
