// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateUserRequest struct {
	NewUser              *CreateUser `protobuf:"bytes,1,opt,name=newUser,proto3" json:"newUser,omitempty"`
	JWT                  string      `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetNewUser() *CreateUser {
	if m != nil {
		return m.NewUser
	}
	return nil
}

func (m *CreateUserRequest) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

type FindByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT                  string   `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIdRequest) Reset()         { *m = FindByIdRequest{} }
func (m *FindByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIdRequest) ProtoMessage()    {}
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *FindByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIdRequest.Unmarshal(m, b)
}
func (m *FindByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIdRequest.Marshal(b, m, deterministic)
}
func (m *FindByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIdRequest.Merge(m, src)
}
func (m *FindByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIdRequest.Size(m)
}
func (m *FindByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIdRequest proto.InternalMessageInfo

func (m *FindByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindByIdRequest) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

type FindByEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	JWT                  string   `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByEmailRequest) Reset()         { *m = FindByEmailRequest{} }
func (m *FindByEmailRequest) String() string { return proto.CompactTextString(m) }
func (*FindByEmailRequest) ProtoMessage()    {}
func (*FindByEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *FindByEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByEmailRequest.Unmarshal(m, b)
}
func (m *FindByEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByEmailRequest.Marshal(b, m, deterministic)
}
func (m *FindByEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByEmailRequest.Merge(m, src)
}
func (m *FindByEmailRequest) XXX_Size() int {
	return xxx_messageInfo_FindByEmailRequest.Size(m)
}
func (m *FindByEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByEmailRequest proto.InternalMessageInfo

func (m *FindByEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *FindByEmailRequest) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	NewPassword          string   `protobuf:"bytes,4,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	JWT                  string   `protobuf:"bytes,5,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateUserRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdateUserRequest) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

type UserReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReply) Reset()         { *m = UserReply{} }
func (m *UserReply) String() string { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()    {}
func (*UserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *UserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReply.Unmarshal(m, b)
}
func (m *UserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReply.Marshal(b, m, deterministic)
}
func (m *UserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReply.Merge(m, src)
}
func (m *UserReply) XXX_Size() int {
	return xxx_messageInfo_UserReply.Size(m)
}
func (m *UserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserReply proto.InternalMessageInfo

func (m *UserReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "users.CreateUserRequest")
	proto.RegisterType((*FindByIdRequest)(nil), "users.FindByIdRequest")
	proto.RegisterType((*FindByEmailRequest)(nil), "users.FindByEmailRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "users.UpdateUserRequest")
	proto.RegisterType((*UserReply)(nil), "users.UserReply")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0xff, 0xa4, 0xcd, 0x2c, 0x68, 0x3b, 0x88, 0xc4, 0x22, 0x58, 0x7a, 0x2a, 0x28,
	0x05, 0x5b, 0xe8, 0xa9, 0xa7, 0x8a, 0x82, 0x1e, 0x44, 0x82, 0x55, 0xf0, 0x16, 0xdd, 0x55, 0x02,
	0x6d, 0x12, 0x77, 0x13, 0x42, 0x1e, 0xc3, 0x27, 0xf5, 0x15, 0x24, 0xbb, 0x9b, 0x3f, 0xa6, 0xf1,
	0x96, 0x99, 0x6f, 0x7f, 0x3b, 0x93, 0xef, 0x5b, 0x20, 0xb1, 0x60, 0x5c, 0xcc, 0x42, 0x1e, 0x44,
	0x01, 0x76, 0x65, 0x31, 0x22, 0x51, 0x1a, 0x32, 0xdd, 0x9b, 0x38, 0x30, 0xbc, 0xe6, 0xcc, 0x8d,
	0xd8, 0x46, 0x30, 0xee, 0xb0, 0xaf, 0x98, 0x89, 0x08, 0x2f, 0xa0, 0xe7, 0xb3, 0x24, 0xeb, 0xd8,
	0xc6, 0xd8, 0x98, 0x92, 0xf9, 0x70, 0xa6, 0x98, 0xca, 0xd1, 0xfc, 0x04, 0x0e, 0xa0, 0x7d, 0xff,
	0xf2, 0x64, 0xb7, 0xc6, 0xc6, 0xd4, 0x72, 0xb2, 0xcf, 0xc9, 0x02, 0x8e, 0x6e, 0x3d, 0x9f, 0xae,
	0xd3, 0x3b, 0x9a, 0xdf, 0x78, 0x08, 0x2d, 0x8f, 0xca, 0xcb, 0xda, 0x4e, 0xcb, 0xa3, 0x0d, 0xd0,
	0x0a, 0x50, 0x41, 0x37, 0x3b, 0xd7, 0xdb, 0xe6, 0xdc, 0x31, 0x74, 0x59, 0x56, 0x4b, 0xd4, 0x72,
	0x54, 0xd1, 0x40, 0x7f, 0x1b, 0x30, 0xdc, 0x84, 0xb4, 0xf6, 0x1f, 0xf5, 0xa9, 0x67, 0x60, 0x7d,
	0x78, 0x5c, 0x44, 0x0f, 0xee, 0x8e, 0x69, 0xba, 0x6c, 0xe0, 0x08, 0xfa, 0x5b, 0x57, 0x8b, 0x6d,
	0x29, 0x16, 0x35, 0x8e, 0x81, 0xf8, 0x2c, 0x79, 0x74, 0x85, 0x48, 0x02, 0x4e, 0xed, 0x8e, 0x94,
	0xab, 0xad, 0x7c, 0xa7, 0x6e, 0xb9, 0xd3, 0x25, 0x58, 0x6a, 0x99, 0x70, 0x9b, 0xe2, 0x39, 0x74,
	0xe2, 0xd2, 0x4f, 0xa2, 0xfd, 0x94, 0xba, 0x14, 0xe6, 0x3f, 0x06, 0xf4, 0x9e, 0xaf, 0xb2, 0x86,
	0xc0, 0x25, 0x98, 0xca, 0x69, 0xb4, 0x67, 0x2a, 0xc0, 0xbd, 0x8c, 0x46, 0x03, 0xad, 0x14, 0x23,
	0x26, 0x07, 0xb8, 0x84, 0x7e, 0x6e, 0x3c, 0x9e, 0x68, 0xbd, 0x96, 0x44, 0x23, 0xb7, 0x02, 0x52,
	0xf1, 0x1e, 0x4f, 0xff, 0xa0, 0xd5, 0x3c, 0xfe, 0x99, 0x6a, 0x2a, 0xeb, 0x8b, 0x6d, 0xf7, 0x92,
	0x68, 0xe2, 0xd6, 0xe6, 0x6b, 0xe7, 0x93, 0x87, 0xef, 0x6f, 0xa6, 0x7c, 0x89, 0x8b, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x11, 0xe5, 0xae, 0x89, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// V1UsersClient is the client API for V1Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type V1UsersClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*UserReply, error)
	FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*UserReply, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type v1UsersClient struct {
	cc grpc.ClientConnInterface
}

func NewV1UsersClient(cc grpc.ClientConnInterface) V1UsersClient {
	return &v1UsersClient{cc}
}

func (c *v1UsersClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/FindByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1UsersServer is the server API for V1Users service.
type V1UsersServer interface {
	Create(context.Context, *CreateUserRequest) (*UserReply, error)
	FindById(context.Context, *FindByIdRequest) (*UserReply, error)
	FindByEmail(context.Context, *FindByEmailRequest) (*UserReply, error)
	Update(context.Context, *UpdateUserRequest) (*UserReply, error)
}

// UnimplementedV1UsersServer can be embedded to have forward compatible implementations.
type UnimplementedV1UsersServer struct {
}

func (*UnimplementedV1UsersServer) Create(ctx context.Context, req *CreateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedV1UsersServer) FindById(ctx context.Context, req *FindByIdRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (*UnimplementedV1UsersServer) FindByEmail(ctx context.Context, req *FindByEmailRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
func (*UnimplementedV1UsersServer) Update(ctx context.Context, req *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterV1UsersServer(s *grpc.Server, srv V1UsersServer) {
	s.RegisterService(&_V1Users_serviceDesc, srv)
}

func _V1Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_FindByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).FindByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/FindByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).FindByEmail(ctx, req.(*FindByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _V1Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.V1Users",
	HandlerType: (*V1UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _V1Users_Create_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _V1Users_FindById_Handler,
		},
		{
			MethodName: "FindByEmail",
			Handler:    _V1Users_FindByEmail_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _V1Users_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
