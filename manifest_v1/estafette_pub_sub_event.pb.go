// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_pub_sub_event.proto

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

type EstafettePubSubEvent struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafettePubSubEvent) Reset()         { *m = EstafettePubSubEvent{} }
func (m *EstafettePubSubEvent) String() string { return proto.CompactTextString(m) }
func (*EstafettePubSubEvent) ProtoMessage()    {}
func (*EstafettePubSubEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_589afd38cf992176, []int{0}
}

func (m *EstafettePubSubEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafettePubSubEvent.Unmarshal(m, b)
}
func (m *EstafettePubSubEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafettePubSubEvent.Marshal(b, m, deterministic)
}
func (m *EstafettePubSubEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafettePubSubEvent.Merge(m, src)
}
func (m *EstafettePubSubEvent) XXX_Size() int {
	return xxx_messageInfo_EstafettePubSubEvent.Size(m)
}
func (m *EstafettePubSubEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafettePubSubEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafettePubSubEvent proto.InternalMessageInfo

func (m *EstafettePubSubEvent) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *EstafettePubSubEvent) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafettePubSubEvent)(nil), "manifest.v1.EstafettePubSubEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_pub_sub_event.proto", fileDescriptor_589afd38cf992176)
}

var fileDescriptor_589afd38cf992176 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0xa9, 0xa0, 0x62, 0xdc, 0x4a, 0x87, 0x8e, 0xe2, 0xa4, 0x43, 0x5b, 0x8a, 0xbb, 0x88,
	0x50, 0x67, 0xd1, 0xcd, 0x25, 0x24, 0xc7, 0xb5, 0x46, 0x6c, 0x2e, 0x34, 0x97, 0xfe, 0x7e, 0xb1,
	0x52, 0xea, 0xf6, 0xbd, 0x8f, 0x37, 0x3c, 0xb1, 0x6f, 0x95, 0x35, 0x35, 0x7a, 0xce, 0xfb, 0xb2,
	0x40, 0xcf, 0xaa, 0x46, 0x66, 0x94, 0x2e, 0x68, 0xe9, 0x83, 0x96, 0xd8, 0xa3, 0xe5, 0xdc, 0x75,
	0xc4, 0x14, 0xaf, 0xff, 0xd4, 0xed, 0x45, 0x24, 0xd5, 0x68, 0x5f, 0x83, 0xbe, 0x07, 0x5d, 0x7d,
	0xd5, 0x38, 0x15, 0x4b, 0xd7, 0xd1, 0x0b, 0x81, 0xd3, 0x68, 0x13, 0xed, 0x56, 0xb7, 0x11, 0xe3,
	0x44, 0xcc, 0x99, 0x9c, 0x81, 0x74, 0x36, 0xfc, 0x3f, 0x38, 0x9f, 0x1e, 0xc7, 0xc6, 0xf0, 0x33,
	0xe8, 0x1c, 0xa8, 0x9d, 0x02, 0xa6, 0x95, 0x81, 0xc9, 0x80, 0x2c, 0x77, 0x0a, 0xd8, 0x67, 0x0d,
	0xbd, 0x95, 0x6d, 0x8a, 0xb1, 0x44, 0xf6, 0xa5, 0x5e, 0x0c, 0x75, 0x87, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x50, 0xbf, 0xb2, 0xee, 0xca, 0x00, 0x00, 0x00,
}