// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/manifest/v1/estafette_release.proto

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
	return fileDescriptor_68a41aaae1211360, []int{0}
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
	proto.RegisterType((*EstafetteRelease)(nil), "estafette.ci.manifest.v1.EstafetteRelease")
}

func init() {
	proto.RegisterFile("estafette/ci/manifest/v1/estafette_release.proto", fileDescriptor_68a41aaae1211360)
}

var fileDescriptor_68a41aaae1211360 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x36, 0xe7, 0x8c, 0x07, 0x47, 0x4e, 0xa1, 0x78, 0x28, 0x9e, 0xaa, 0xd0, 0x74,
	0xd5, 0x83, 0x07, 0x2f, 0xfe, 0x40, 0x0f, 0x82, 0x20, 0x51, 0x3c, 0x78, 0x29, 0x59, 0xc8, 0x62,
	0xa0, 0x6d, 0x46, 0xf3, 0x2c, 0xfa, 0x17, 0x09, 0xfe, 0x95, 0xb2, 0xb4, 0xcd, 0x40, 0x18, 0xf4,
	0x16, 0xde, 0xe3, 0xf3, 0x79, 0xdf, 0x7e, 0x8b, 0x16, 0xd2, 0x02, 0x5f, 0x49, 0x00, 0x99, 0x0a,
	0x9d, 0x96, 0xbc, 0xd2, 0x2b, 0x69, 0x21, 0x6d, 0xb2, 0xd4, 0x2f, 0xf2, 0x5a, 0x16, 0x92, 0x5b,
	0x49, 0xd7, 0xb5, 0x01, 0x83, 0x89, 0x5f, 0x50, 0xa1, 0x69, 0x4f, 0xd0, 0x26, 0x0b, 0x2f, 0x87,
	0xbb, 0x72, 0x2e, 0x40, 0x9b, 0xaa, 0x55, 0x86, 0x74, 0x00, 0x68, 0x81, 0xab, 0x2e, 0x42, 0x38,
	0x24, 0x34, 0xd4, 0x5a, 0x29, 0x59, 0xb7, 0xc4, 0xc9, 0xcf, 0x08, 0xcd, 0xef, 0xfb, 0x1d, 0x6b,
	0x33, 0x60, 0x8c, 0x26, 0x15, 0x2f, 0x25, 0x09, 0xa2, 0x20, 0x3e, 0x60, 0xee, 0x8d, 0x4f, 0xd1,
	0x5c, 0x14, 0xa6, 0xda, 0x04, 0x5d, 0x1b, 0xab, 0xc1, 0xd4, 0xdf, 0x64, 0x14, 0x05, 0xf1, 0x8c,
	0x1d, 0xb9, 0x39, 0xf3, 0x63, 0xfc, 0x88, 0xf6, 0xdb, 0xaf, 0xb0, 0x64, 0x1c, 0x8d, 0xe3, 0xc3,
	0xf3, 0x05, 0xdd, 0x55, 0x0d, 0xfd, 0x7f, 0xfb, 0xc6, 0x81, 0xac, 0x17, 0xe0, 0x07, 0x34, 0xeb,
	0x02, 0x5b, 0x32, 0x71, 0xb2, 0xb3, 0x01, 0xb2, 0xd7, 0x16, 0x61, 0x9e, 0xc5, 0xd7, 0x68, 0xea,
	0x8a, 0xb2, 0x64, 0xcf, 0x59, 0xe2, 0x01, 0x96, 0x97, 0x0d, 0xc0, 0x3a, 0xee, 0xf6, 0x0b, 0x1d,
	0x0b, 0x53, 0xee, 0xc4, 0x9e, 0x83, 0xf7, 0x2b, 0xa5, 0xe1, 0xe3, 0x73, 0x49, 0x85, 0x29, 0xb7,
	0x7d, 0x6f, 0x5f, 0x89, 0xd0, 0x89, 0x2b, 0xdd, 0x26, 0xca, 0x14, 0xbc, 0x52, 0xfe, 0x07, 0xe5,
	0x4d, 0xf6, 0x3b, 0x22, 0xfe, 0x2e, 0xbd, 0xd3, 0xf4, 0xa9, 0x37, 0xbf, 0x65, 0xcb, 0xa9, 0xa3,
	0x2e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x33, 0xb4, 0xe3, 0x93, 0x02, 0x00, 0x00,
}
