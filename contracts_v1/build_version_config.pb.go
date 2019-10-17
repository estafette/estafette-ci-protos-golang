// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/build_version_config.proto

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

type BuildVersionConfig struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Major                int64    `protobuf:"varint,2,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int64    `protobuf:"varint,3,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                string   `protobuf:"bytes,4,opt,name=patch,proto3" json:"patch,omitempty"`
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	AutoIncrement        int64    `protobuf:"varint,6,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildVersionConfig) Reset()         { *m = BuildVersionConfig{} }
func (m *BuildVersionConfig) String() string { return proto.CompactTextString(m) }
func (*BuildVersionConfig) ProtoMessage()    {}
func (*BuildVersionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd2ac76605caf77, []int{0}
}

func (m *BuildVersionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildVersionConfig.Unmarshal(m, b)
}
func (m *BuildVersionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildVersionConfig.Marshal(b, m, deterministic)
}
func (m *BuildVersionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildVersionConfig.Merge(m, src)
}
func (m *BuildVersionConfig) XXX_Size() int {
	return xxx_messageInfo_BuildVersionConfig.Size(m)
}
func (m *BuildVersionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildVersionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BuildVersionConfig proto.InternalMessageInfo

func (m *BuildVersionConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *BuildVersionConfig) GetMajor() int64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *BuildVersionConfig) GetMinor() int64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *BuildVersionConfig) GetPatch() string {
	if m != nil {
		return m.Patch
	}
	return ""
}

func (m *BuildVersionConfig) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *BuildVersionConfig) GetAutoIncrement() int64 {
	if m != nil {
		return m.AutoIncrement
	}
	return 0
}

func init() {
	proto.RegisterType((*BuildVersionConfig)(nil), "contracts.v1.BuildVersionConfig")
}

func init() {
	proto.RegisterFile("contracts.v1/build_version_config.proto", fileDescriptor_bdd2ac76605caf77)
}

var fileDescriptor_bdd2ac76605caf77 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xe9, 0xe6, 0x26, 0x06, 0xf5, 0x10, 0x3c, 0xc4, 0xdb, 0x10, 0xc4, 0x5d, 0x9a, 0x52,
	0xfc, 0x01, 0xc2, 0x8a, 0x87, 0x5d, 0x77, 0xd8, 0x41, 0x90, 0x92, 0xc6, 0x2c, 0x8b, 0xa4, 0x79,
	0x23, 0x7d, 0xed, 0x8f, 0x12, 0xfc, 0x8f, 0x92, 0x64, 0xad, 0xbb, 0xe5, 0xfb, 0xf8, 0xf2, 0xe0,
	0x3d, 0xf2, 0x22, 0xc1, 0xa1, 0x17, 0x12, 0x3b, 0x3e, 0x94, 0x45, 0xd3, 0x1b, 0xfb, 0x55, 0x0f,
	0xca, 0x77, 0x06, 0x5c, 0x2d, 0xc1, 0x1d, 0x8c, 0xe6, 0x27, 0x0f, 0x08, 0xf4, 0xf6, 0x32, 0x7c,
	0xfa, 0xcd, 0x08, 0xdd, 0x84, 0x78, 0x9f, 0xda, 0x2a, 0xa6, 0x94, 0x91, 0xeb, 0xf3, 0x67, 0x96,
	0xad, 0xb2, 0xf5, 0xcd, 0x6e, 0x44, 0xfa, 0x40, 0x16, 0xad, 0xf8, 0x06, 0xcf, 0x66, 0xab, 0x6c,
	0x3d, 0xdf, 0x25, 0x88, 0xd6, 0x38, 0xf0, 0x6c, 0x7e, 0xb6, 0x01, 0x82, 0x3d, 0x09, 0x94, 0x47,
	0x76, 0x15, 0x67, 0x24, 0x08, 0xd6, 0x8a, 0x46, 0x59, 0xb6, 0x48, 0x36, 0x02, 0x7d, 0x26, 0xf7,
	0xa2, 0x47, 0xa8, 0x8d, 0x93, 0x5e, 0xb5, 0xca, 0x21, 0x5b, 0xc6, 0x51, 0x77, 0xc1, 0x6e, 0x47,
	0xb9, 0xf9, 0xfc, 0x78, 0xd3, 0x06, 0x8f, 0x7d, 0xc3, 0x25, 0xb4, 0x85, 0xea, 0x50, 0x1c, 0x14,
	0xa2, 0xfa, 0x7f, 0xe5, 0xd2, 0xe4, 0xd3, 0x86, 0xb9, 0x06, 0x2b, 0x9c, 0x2e, 0x26, 0x51, 0x0f,
	0xe5, 0xcf, 0xec, 0xf1, 0x7d, 0x8c, 0x79, 0xb5, 0xe5, 0xd5, 0x74, 0x8e, 0x7d, 0xd9, 0x2c, 0xe3,
	0x8d, 0x5e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0x22, 0x9d, 0x89, 0x4e, 0x01, 0x00, 0x00,
}
