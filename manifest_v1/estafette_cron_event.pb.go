// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/manifest/v1/estafette_cron_event.proto

package manifest_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EstafetteCronEvent struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EstafetteCronEvent) Reset()         { *m = EstafetteCronEvent{} }
func (m *EstafetteCronEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteCronEvent) ProtoMessage()    {}
func (*EstafetteCronEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a32598ffc9469427, []int{0}
}

func (m *EstafetteCronEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteCronEvent.Unmarshal(m, b)
}
func (m *EstafetteCronEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteCronEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteCronEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteCronEvent.Merge(m, src)
}
func (m *EstafetteCronEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteCronEvent.Size(m)
}
func (m *EstafetteCronEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteCronEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteCronEvent proto.InternalMessageInfo

func (m *EstafetteCronEvent) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteCronEvent)(nil), "estafette.ci.manifest.v1.EstafetteCronEvent")
}

func init() {
	proto.RegisterFile("estafette/ci/manifest/v1/estafette_cron_event.proto", fileDescriptor_a32598ffc9469427)
}

var fileDescriptor_a32598ffc9469427 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x4a, 0xc0, 0x30,
	0x10, 0x40, 0xa9, 0x88, 0x43, 0xdc, 0x3a, 0x95, 0x22, 0x28, 0x4e, 0x2e, 0xbd, 0x10, 0x3b, 0xba,
	0x59, 0x3b, 0x0a, 0x22, 0xe2, 0xe0, 0x52, 0xda, 0x90, 0xc6, 0x40, 0x93, 0x2b, 0xed, 0x35, 0xf8,
	0x4d, 0x7e, 0xa5, 0x34, 0x25, 0xe9, 0xe4, 0x76, 0x70, 0xef, 0x3d, 0xee, 0x58, 0xad, 0x56, 0xea,
	0x47, 0x45, 0xa4, 0xb8, 0x34, 0xdc, 0xf6, 0xce, 0x8c, 0x6a, 0x25, 0xee, 0x05, 0x4f, 0x8b, 0x4e,
	0x2e, 0xe8, 0x3a, 0xe5, 0x95, 0x23, 0x98, 0x17, 0x24, 0xcc, 0x8b, 0xb4, 0x03, 0x69, 0x20, 0x4a,
	0xe0, 0x45, 0x79, 0xab, 0x11, 0xf5, 0xa4, 0x78, 0xe0, 0x86, 0x6d, 0xe4, 0x64, 0xec, 0x0e, 0xdb,
	0xf9, 0x50, 0xef, 0x5f, 0x58, 0xde, 0x46, 0xb9, 0x59, 0xd0, 0xb5, 0x7b, 0x36, 0x07, 0x76, 0xb9,
	0x83, 0x45, 0x76, 0x97, 0x3d, 0x5c, 0x3f, 0x96, 0x70, 0x54, 0x20, 0x56, 0xe0, 0x23, 0x56, 0xde,
	0x03, 0xf7, 0xfc, 0xc3, 0x6e, 0x24, 0x5a, 0xf8, 0xef, 0x8c, 0xb7, 0xec, 0xeb, 0x49, 0x1b, 0xfa,
	0xde, 0x06, 0x90, 0x68, 0xcf, 0x4f, 0xce, 0xa9, 0x92, 0xa6, 0x0a, 0xf5, 0xb5, 0xd2, 0x38, 0xf5,
	0x4e, 0xa7, 0xd7, 0x3b, 0x2f, 0x7e, 0x2f, 0x8a, 0x74, 0x23, 0x34, 0x06, 0x5e, 0x63, 0xf9, 0x53,
	0x0c, 0x57, 0xc1, 0xaa, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x6e, 0xd7, 0xdd, 0x38, 0x01,
	0x00, 0x00,
}
