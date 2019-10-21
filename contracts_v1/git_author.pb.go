// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/git_author.proto

package contracts_v1

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

type GitAuthor struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitAuthor) Reset()         { *m = GitAuthor{} }
func (m *GitAuthor) String() string { return proto.CompactTextString(m) }
func (*GitAuthor) ProtoMessage()    {}
func (*GitAuthor) Descriptor() ([]byte, []int) {
	return fileDescriptor_402389c8bc417805, []int{0}
}

func (m *GitAuthor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitAuthor.Unmarshal(m, b)
}
func (m *GitAuthor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitAuthor.Marshal(b, m, deterministic)
}
func (m *GitAuthor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitAuthor.Merge(m, src)
}
func (m *GitAuthor) XXX_Size() int {
	return xxx_messageInfo_GitAuthor.Size(m)
}
func (m *GitAuthor) XXX_DiscardUnknown() {
	xxx_messageInfo_GitAuthor.DiscardUnknown(m)
}

var xxx_messageInfo_GitAuthor proto.InternalMessageInfo

func (m *GitAuthor) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GitAuthor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GitAuthor) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*GitAuthor)(nil), "contracts.v1.GitAuthor")
}

func init() { proto.RegisterFile("contracts.v1/git_author.proto", fileDescriptor_402389c8bc417805) }

var fileDescriptor_402389c8bc417805 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x33, 0xd4, 0x4f, 0xcf, 0x2c, 0x89, 0x4f, 0x2c, 0x2d,
	0xc9, 0xc8, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x96, 0x56, 0x0a, 0xe4,
	0xe2, 0x74, 0xcf, 0x2c, 0x71, 0x04, 0x2b, 0x10, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x98, 0xc0, 0x82, 0x60, 0xb6, 0x90, 0x14, 0x17, 0x47, 0x69, 0x71, 0x6a, 0x11, 0x58,
	0x9c, 0x19, 0x2c, 0x0e, 0xe7, 0x3b, 0x55, 0x73, 0xc9, 0x64, 0xe6, 0xeb, 0xa5, 0x16, 0x97, 0x24,
	0xa6, 0xa5, 0x96, 0x94, 0xa4, 0xea, 0x25, 0x67, 0xea, 0x21, 0x5b, 0x19, 0x65, 0x9f, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x57, 0x85, 0x60, 0xe9, 0x26, 0x67, 0xea,
	0xc2, 0xd5, 0xeb, 0xa6, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0xc3, 0x05, 0xe2, 0xcb, 0x0c, 0x57,
	0x31, 0x49, 0xba, 0xc2, 0x0d, 0x77, 0xf6, 0xd4, 0x73, 0x86, 0x1b, 0x1e, 0x66, 0x98, 0xc4, 0x06,
	0xf6, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x48, 0x92, 0x26, 0x34, 0x05, 0x01, 0x00, 0x00,
}
