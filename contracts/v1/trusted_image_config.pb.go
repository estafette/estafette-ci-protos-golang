// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/contracts/v1/trusted_image_config.proto

package v1

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

type TrustedImageConfig struct {
	ImagePath               string   `protobuf:"bytes,1,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	RunPrivileged           bool     `protobuf:"varint,2,opt,name=run_privileged,json=runPrivileged,proto3" json:"run_privileged,omitempty"`
	RunDocker               bool     `protobuf:"varint,3,opt,name=run_docker,json=runDocker,proto3" json:"run_docker,omitempty"`
	AllowCommands           bool     `protobuf:"varint,4,opt,name=allow_commands,json=allowCommands,proto3" json:"allow_commands,omitempty"`
	InjectedCredentialTypes []string `protobuf:"bytes,5,rep,name=injected_credential_types,json=injectedCredentialTypes,proto3" json:"injected_credential_types,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *TrustedImageConfig) Reset()         { *m = TrustedImageConfig{} }
func (m *TrustedImageConfig) String() string { return proto.CompactTextString(m) }
func (*TrustedImageConfig) ProtoMessage()    {}
func (*TrustedImageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16176a00c54fd96, []int{0}
}

func (m *TrustedImageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustedImageConfig.Unmarshal(m, b)
}
func (m *TrustedImageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustedImageConfig.Marshal(b, m, deterministic)
}
func (m *TrustedImageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustedImageConfig.Merge(m, src)
}
func (m *TrustedImageConfig) XXX_Size() int {
	return xxx_messageInfo_TrustedImageConfig.Size(m)
}
func (m *TrustedImageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustedImageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TrustedImageConfig proto.InternalMessageInfo

func (m *TrustedImageConfig) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

func (m *TrustedImageConfig) GetRunPrivileged() bool {
	if m != nil {
		return m.RunPrivileged
	}
	return false
}

func (m *TrustedImageConfig) GetRunDocker() bool {
	if m != nil {
		return m.RunDocker
	}
	return false
}

func (m *TrustedImageConfig) GetAllowCommands() bool {
	if m != nil {
		return m.AllowCommands
	}
	return false
}

func (m *TrustedImageConfig) GetInjectedCredentialTypes() []string {
	if m != nil {
		return m.InjectedCredentialTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*TrustedImageConfig)(nil), "estafette.ci.contracts.v1.TrustedImageConfig")
}

func init() {
	proto.RegisterFile("estafette/ci/contracts/v1/trusted_image_config.proto", fileDescriptor_c16176a00c54fd96)
}

var fileDescriptor_c16176a00c54fd96 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xb9, 0x56, 0xc5, 0x0b, 0xe8, 0x70, 0x8b, 0xd7, 0xa1, 0x50, 0x04, 0xa1, 0x4b, 0x73,
	0x14, 0x9d, 0xc4, 0xc9, 0xd3, 0xc1, 0xad, 0x94, 0xe2, 0xe0, 0x72, 0xa4, 0xdf, 0x7d, 0x4d, 0xa3,
	0x77, 0xc9, 0x91, 0x7c, 0x57, 0x11, 0xff, 0x91, 0xbf, 0xc9, 0x1f, 0x23, 0x49, 0x35, 0xe2, 0xe0,
	0x16, 0xde, 0xbc, 0xcf, 0x03, 0xdf, 0xcb, 0xae, 0xd0, 0x91, 0xd8, 0x20, 0x11, 0x16, 0xa0, 0x0a,
	0x30, 0x9a, 0xac, 0x00, 0x72, 0xc5, 0x6e, 0x5e, 0x90, 0xed, 0x1d, 0x61, 0x5d, 0xa9, 0x56, 0x48,
	0xac, 0xc0, 0xe8, 0x8d, 0x92, 0xbc, 0xb3, 0x86, 0x4c, 0x36, 0x8a, 0x14, 0x07, 0xc5, 0x23, 0xc5,
	0x77, 0xf3, 0xf3, 0xcf, 0x84, 0x65, 0xab, 0x3d, 0xf9, 0xe0, 0xc1, 0x32, 0x70, 0xd9, 0x98, 0xb1,
	0xbd, 0xa7, 0x13, 0xb4, 0xcd, 0x93, 0x49, 0x32, 0x4d, 0x97, 0x69, 0x48, 0x16, 0x82, 0xb6, 0xd9,
	0x05, 0x3b, 0xb5, 0xbd, 0xae, 0x3a, 0xab, 0x76, 0xaa, 0x41, 0x89, 0x75, 0x3e, 0x98, 0x24, 0xd3,
	0xe3, 0xe5, 0x89, 0xed, 0xf5, 0x22, 0x86, 0xde, 0xe2, 0x6b, 0xb5, 0x81, 0x17, 0xb4, 0xf9, 0x30,
	0x54, 0x52, 0xdb, 0xeb, 0xbb, 0x10, 0x78, 0x8b, 0x68, 0x1a, 0xf3, 0x5a, 0x81, 0x69, 0x5b, 0xa1,
	0x6b, 0x97, 0x1f, 0xec, 0x2d, 0x21, 0x2d, 0xbf, 0xc3, 0xec, 0x9a, 0x8d, 0x94, 0x7e, 0x46, 0xf0,
	0xc7, 0x81, 0xc5, 0x1a, 0x35, 0x29, 0xd1, 0x54, 0xf4, 0xd6, 0xa1, 0xcb, 0x0f, 0x27, 0xc3, 0x69,
	0xba, 0x3c, 0xfb, 0x29, 0x94, 0xf1, 0x7f, 0xe5, 0xbf, 0x6f, 0xdf, 0xd9, 0x18, 0x4c, 0xcb, 0xff,
	0xbd, 0x7f, 0x91, 0x3c, 0xdd, 0x48, 0x45, 0xdb, 0x7e, 0xcd, 0xc1, 0xb4, 0xc5, 0xef, 0xba, 0xf1,
	0x35, 0x03, 0x35, 0x0b, 0x2b, 0xba, 0x99, 0x34, 0x8d, 0xd0, 0xf2, 0xcf, 0xea, 0x1f, 0x83, 0xd1,
	0x7d, 0x74, 0x97, 0x8a, 0x97, 0xd1, 0xfd, 0x38, 0x5f, 0x1f, 0x05, 0xee, 0xf2, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0x11, 0x90, 0x3b, 0xb5, 0x01, 0x00, 0x00,
}
