// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_release.proto

package manifest_v1

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

type EstafetteRelease struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CloneRepository      bool                      `protobuf:"varint,2,opt,name=clone_repository,json=cloneRepository,proto3" json:"clone_repository,omitempty"`
	Actions              []*EstafetteReleaseAction `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Triggers             []*EstafetteTrigger       `protobuf:"bytes,4,rep,name=triggers,proto3" json:"triggers,omitempty"`
	Stages               []*EstafetteStage         `protobuf:"bytes,5,rep,name=stages,proto3" json:"stages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EstafetteRelease) Reset()         { *m = EstafetteRelease{} }
func (m *EstafetteRelease) String() string { return proto.CompactTextString(m) }
func (*EstafetteRelease) ProtoMessage()    {}
func (*EstafetteRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d013379f885dc8f, []int{0}
}

func (m *EstafetteRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteRelease.Unmarshal(m, b)
}
func (m *EstafetteRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteRelease.Marshal(b, m, deterministic)
}
func (m *EstafetteRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteRelease.Merge(m, src)
}
func (m *EstafetteRelease) XXX_Size() int {
	return xxx_messageInfo_EstafetteRelease.Size(m)
}
func (m *EstafetteRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteRelease.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteRelease proto.InternalMessageInfo

func (m *EstafetteRelease) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EstafetteRelease) GetCloneRepository() bool {
	if m != nil {
		return m.CloneRepository
	}
	return false
}

func (m *EstafetteRelease) GetActions() []*EstafetteReleaseAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *EstafetteRelease) GetTriggers() []*EstafetteTrigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func (m *EstafetteRelease) GetStages() []*EstafetteStage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteRelease)(nil), "manifest.v1.EstafetteRelease")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_release.proto", fileDescriptor_0d013379f885dc8f)
}

var fileDescriptor_0d013379f885dc8f = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0xb6, 0x5f, 0xbf, 0x9a, 0x1e, 0x2c, 0x39, 0x85, 0x8a, 0xb0, 0xda, 0xcb, 0x2a,
	0x6c, 0x4a, 0xed, 0xc9, 0x83, 0xa2, 0x82, 0x7f, 0x20, 0x7a, 0xf2, 0xb2, 0xa4, 0x21, 0x8d, 0x81,
	0xdd, 0xa4, 0x64, 0xc6, 0x82, 0x37, 0x7f, 0xba, 0x34, 0xbb, 0xdb, 0x16, 0x59, 0xbd, 0x85, 0xc9,
	0xf3, 0xcc, 0xbc, 0x33, 0x64, 0x56, 0x49, 0x67, 0xd7, 0x1a, 0x90, 0x6f, 0x17, 0x73, 0x0d, 0x28,
	0xd7, 0x1a, 0x51, 0x17, 0x41, 0x97, 0x5a, 0x82, 0xe6, 0x9b, 0xe0, 0xd1, 0xd3, 0xf1, 0x11, 0x34,
	0xbd, 0xfe, 0xd3, 0x28, 0xa4, 0x42, 0xeb, 0x5d, 0x2d, 0x4e, 0x7f, 0xe9, 0x8e, 0xc1, 0x1a, 0xa3,
	0x43, 0x03, 0x5d, 0x74, 0x43, 0x80, 0xd2, 0x34, 0x01, 0x2e, 0xbf, 0x7a, 0x64, 0xf2, 0xdc, 0xfe,
	0x88, 0x7a, 0x12, 0xa5, 0x64, 0xe0, 0x64, 0xa5, 0x59, 0x92, 0x26, 0xd9, 0x89, 0x88, 0x6f, 0x7a,
	0x45, 0x26, 0xaa, 0xf4, 0x6e, 0x17, 0x67, 0xe3, 0xc1, 0xa2, 0x0f, 0x9f, 0xac, 0x97, 0x26, 0xd9,
	0x48, 0x9c, 0xc6, 0xba, 0xd8, 0x97, 0xe9, 0x1d, 0xf9, 0x5f, 0x67, 0x05, 0xd6, 0x4f, 0xfb, 0xd9,
	0xf8, 0x66, 0xc6, 0x8f, 0x82, 0xf0, 0x9f, 0xe3, 0x1e, 0x23, 0x2b, 0x5a, 0x87, 0xde, 0x92, 0x51,
	0xb3, 0x06, 0xb0, 0x41, 0xf4, 0xcf, 0xbb, 0xfd, 0xd7, 0x9a, 0x12, 0x7b, 0x9c, 0x2e, 0xc9, 0x30,
	0x2e, 0x07, 0xec, 0x5f, 0x14, 0xcf, 0xba, 0xc5, 0x97, 0x1d, 0x23, 0x1a, 0xf4, 0xe9, 0xe1, 0xed,
	0xde, 0x58, 0x7c, 0xff, 0x58, 0x71, 0xe5, 0xab, 0xc3, 0x99, 0x0e, 0xaf, 0x5c, 0xd9, 0x5c, 0x79,
	0x87, 0x41, 0x2a, 0x84, 0xdc, 0xf8, 0x52, 0x3a, 0x33, 0x6f, 0x1b, 0x17, 0xdb, 0xc5, 0x6a, 0x18,
	0x6f, 0xb9, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x45, 0x5d, 0xce, 0x37, 0xf3, 0x01, 0x00, 0x00,
}