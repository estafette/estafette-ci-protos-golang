// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_cron_event.proto

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
	return fileDescriptor_055a304991ebfb9d, []int{0}
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
	proto.RegisterType((*EstafetteCronEvent)(nil), "manifest.v1.EstafetteCronEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_cron_event.proto", fileDescriptor_055a304991ebfb9d)
}

var fileDescriptor_055a304991ebfb9d = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0x8d, 0x4f, 0x2e, 0xca, 0xcf, 0x8b, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x52, 0x27, 0x25, 0x9f, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x0f,
	0x96, 0x4a, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x05, 0xe9, 0xcd, 0x2d, 0x80, 0xa8, 0x56, 0x72,
	0xe1, 0x12, 0x72, 0x85, 0x99, 0xe5, 0x5c, 0x94, 0x9f, 0xe7, 0x0a, 0x32, 0x49, 0x48, 0x8f, 0x8b,
	0x05, 0xa4, 0x50, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0x62, 0x8a, 0x1e, 0xcc,
	0x14, 0xbd, 0x10, 0x98, 0x29, 0x41, 0x60, 0x75, 0x4e, 0xd1, 0x51, 0x76, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x08, 0xc7, 0x21, 0x58, 0xba, 0xc9, 0x99, 0xba, 0xc9, 0xf9,
	0x79, 0x25, 0x45, 0x89, 0xc9, 0x25, 0xc5, 0xba, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0xfa, 0x30,
	0x87, 0xc6, 0x97, 0x19, 0xae, 0x62, 0x92, 0x80, 0x3b, 0x43, 0xcf, 0xd9, 0x53, 0xcf, 0x17, 0xe6,
	0x87, 0x30, 0xc3, 0x24, 0x36, 0xb0, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x7d,
	0x8f, 0xba, 0x01, 0x01, 0x00, 0x00,
}
