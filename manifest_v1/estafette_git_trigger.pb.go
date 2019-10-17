// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_git_trigger.proto

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

type EstafetteGitTrigger struct {
	Event                string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Repository           string   `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch               string   `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteGitTrigger) Reset()         { *m = EstafetteGitTrigger{} }
func (m *EstafetteGitTrigger) String() string { return proto.CompactTextString(m) }
func (*EstafetteGitTrigger) ProtoMessage()    {}
func (*EstafetteGitTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_006daccc9169db0a, []int{0}
}

func (m *EstafetteGitTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteGitTrigger.Unmarshal(m, b)
}
func (m *EstafetteGitTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteGitTrigger.Marshal(b, m, deterministic)
}
func (m *EstafetteGitTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteGitTrigger.Merge(m, src)
}
func (m *EstafetteGitTrigger) XXX_Size() int {
	return xxx_messageInfo_EstafetteGitTrigger.Size(m)
}
func (m *EstafetteGitTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteGitTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteGitTrigger proto.InternalMessageInfo

func (m *EstafetteGitTrigger) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EstafetteGitTrigger) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *EstafetteGitTrigger) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteGitTrigger)(nil), "manifest.v1.EstafetteGitTrigger")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_git_trigger.proto", fileDescriptor_006daccc9169db0a)
}

var fileDescriptor_006daccc9169db0a = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0x8d, 0x4f, 0xcf, 0x2c, 0x89, 0x2f, 0x29, 0xca, 0x4c, 0x4f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x52, 0xa8, 0x94, 0xcc, 0x25, 0xec, 0x0a, 0x53, 0xeb, 0x9e,
	0x59, 0x12, 0x02, 0x51, 0x29, 0x24, 0xc2, 0xc5, 0x9a, 0x5a, 0x96, 0x9a, 0x57, 0x22, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0xc9, 0x71, 0x71, 0x15, 0xa5, 0x16, 0xe4, 0x17, 0x67,
	0x96, 0xe4, 0x17, 0x55, 0x4a, 0x30, 0x81, 0xa5, 0x90, 0x44, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0x8a,
	0x12, 0xf3, 0x92, 0x33, 0x24, 0x98, 0xc1, 0x72, 0x50, 0x9e, 0x53, 0x74, 0x94, 0x5d, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0xc2, 0x6d, 0x08, 0x96, 0x6e, 0x72, 0xa6, 0x6e,
	0x72, 0x7e, 0x5e, 0x49, 0x51, 0x62, 0x72, 0x49, 0xb1, 0x6e, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba,
	0x3e, 0xcc, 0x99, 0xf1, 0x65, 0x86, 0xab, 0x98, 0x24, 0xe0, 0xae, 0xd4, 0x73, 0xf6, 0xd4, 0xf3,
	0x85, 0xf9, 0x20, 0xcc, 0x30, 0x89, 0x0d, 0xec, 0x2b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xe0, 0xbe, 0xa4, 0x00, 0x01, 0x00, 0x00,
}
