// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model_user.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserModel struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Fullname             string   `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3c508ca69649dd, []int{0}
}

func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModel.Unmarshal(m, b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return xxx_messageInfo_UserModel.Size(m)
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserModel) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserModel) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserModel) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func init() {
	proto.RegisterType((*UserModel)(nil), "api.UserModel")
}

func init() { proto.RegisterFile("model_user.proto", fileDescriptor_7d3c508ca69649dd) }

var fileDescriptor_7d3c508ca69649dd = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcd, 0x4f, 0x49,
	0xcd, 0x89, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x54, 0xca, 0xe7, 0xe2, 0x0c, 0x2d, 0x4e, 0x2d, 0xf2, 0x05, 0x49, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x08, 0x49, 0x71, 0x71,
	0x80, 0xd4, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20,
	0xb9, 0x82, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x66, 0x88, 0x1c, 0x8c, 0x0f, 0x92,
	0x4b, 0x2b, 0xcd, 0xc9, 0x01, 0xeb, 0x63, 0x81, 0xc8, 0xc1, 0xf8, 0x49, 0x6c, 0x60, 0xcb, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x5c, 0x4d, 0xd0, 0x90, 0x00, 0x00, 0x00,
}
