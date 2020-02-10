// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api_user.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("api_user.proto", fileDescriptor_f72bfd0c56b6607f) }

var fileDescriptor_f72bfd0c56b6607f = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x60, 0xb4, 0xa2, 0x30, 0x87, 0x1e, 0xa2, 0xf4, 0x10, 0xf5, 0x92, 0x53, 0x09, 0x74,
	0x57, 0xed, 0xcd, 0xab, 0x8a, 0x77, 0xa1, 0x67, 0x99, 0xd2, 0x49, 0x58, 0x48, 0x77, 0xb6, 0xbb,
	0x13, 0x51, 0xc4, 0x8b, 0xaf, 0xe0, 0xd3, 0xf8, 0x1c, 0xbe, 0x82, 0x0f, 0x22, 0xd9, 0xa6, 0xc4,
	0x5c, 0xda, 0x5e, 0x42, 0xe6, 0x9f, 0x61, 0x3e, 0x66, 0x61, 0x88, 0xce, 0x3c, 0xd7, 0x81, 0xbc,
	0x72, 0x9e, 0x85, 0x93, 0x01, 0x3a, 0x93, 0x5e, 0x94, 0xcc, 0x65, 0x45, 0x1a, 0x9d, 0xd1, 0x68,
	0x2d, 0x0b, 0x8a, 0x61, 0x1b, 0xd6, 0x23, 0xe9, 0x79, 0xdb, 0x8d, 0xd5, 0xbc, 0x2e, 0x34, 0x2d,
	0x9d, 0xbc, 0xb5, 0xcd, 0xe1, 0x42, 0xf8, 0xdf, 0xbe, 0x9b, 0xef, 0x01, 0x1c, 0xcd, 0x02, 0xf9,
	0xa4, 0x80, 0x93, 0x47, 0x92, 0xf8, 0x7b, 0xaa, 0xd0, 0x19, 0xd5, 0x56, 0x4f, 0xb4, 0xaa, 0x29,
	0x48, 0x7a, 0xd6, 0x0f, 0x83, 0x63, 0x1b, 0x28, 0xbb, 0xfa, 0xfc, 0xf9, 0xfd, 0x3a, 0xcc, 0x93,
	0xb1, 0x2e, 0x79, 0xd2, 0x81, 0xaf, 0x42, 0xde, 0x62, 0xa5, 0x5f, 0xae, 0x75, 0x63, 0xe9, 0xf7,
	0xe6, 0x6b, 0x71, 0x49, 0x1f, 0x49, 0x01, 0x70, 0xe7, 0x09, 0x85, 0x22, 0x35, 0x8a, 0x5b, 0xbb,
	0x60, 0xa3, 0x8d, 0xd4, 0xfa, 0x08, 0xb5, 0xd9, 0xa9, 0x1e, 0x9a, 0x23, 0xb2, 0x71, 0xf4, 0xb2,
	0xec, 0x72, 0xab, 0x77, 0x7b, 0x90, 0x27, 0x2b, 0x80, 0x99, 0x5b, 0xf4, 0x9d, 0x2e, 0xd8, 0xe5,
	0x4c, 0xa3, 0x33, 0x49, 0xf7, 0xbe, 0xab, 0x21, 0x2d, 0xc0, 0x3d, 0x55, 0xd4, 0x23, 0xbb, 0x60,
	0x17, 0xd9, 0x3e, 0x65, 0xbe, 0x37, 0x39, 0x3f, 0x8e, 0x53, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x50, 0xfa, 0xca, 0x24, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*empty.Empty, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_user.proto",
}
