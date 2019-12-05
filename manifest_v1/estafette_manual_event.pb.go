// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/manifest/v1/estafette_manual_event.proto

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

type EstafetteManualEvent struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteManualEvent) Reset()         { *m = EstafetteManualEvent{} }
func (m *EstafetteManualEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteManualEvent) ProtoMessage()    {}
func (*EstafetteManualEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a34e12f6e00bc49, []int{0}
}

func (m *EstafetteManualEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteManualEvent.Unmarshal(m, b)
}
func (m *EstafetteManualEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteManualEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteManualEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteManualEvent.Merge(m, src)
}
func (m *EstafetteManualEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteManualEvent.Size(m)
}
func (m *EstafetteManualEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteManualEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteManualEvent proto.InternalMessageInfo

func (m *EstafetteManualEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteManualEvent)(nil), "estafette.ci.manifest.v1.EstafetteManualEvent")
}

func init() {
	proto.RegisterFile("estafette/ci/manifest/v1/estafette_manual_event.proto", fileDescriptor_1a34e12f6e00bc49)
}

var fileDescriptor_1a34e12f6e00bc49 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0x2d, 0x2e, 0x49,
	0x4c, 0x4b, 0x2d, 0x29, 0x49, 0xd5, 0x4f, 0xce, 0xd4, 0xcf, 0x4d, 0xcc, 0xcb, 0x4c, 0x4b, 0x2d,
	0x2e, 0xd1, 0x2f, 0x33, 0xd4, 0x87, 0x4b, 0xc4, 0xe7, 0x26, 0xe6, 0x95, 0x26, 0xe6, 0xc4, 0xa7,
	0x96, 0xa5, 0xe6, 0x95, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0xc0, 0x65, 0xf5, 0x92,
	0x33, 0xf5, 0x60, 0xda, 0xf4, 0xca, 0x0c, 0x95, 0xf4, 0xb9, 0x44, 0x5c, 0x61, 0x72, 0xbe, 0x60,
	0x8d, 0xae, 0x20, 0x7d, 0x42, 0xe2, 0x5c, 0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x53, 0x05, 0x97, 0x4c, 0x72,
	0x7e, 0xae, 0x1e, 0x2e, 0x03, 0x03, 0x18, 0xa3, 0xac, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0x11, 0xae, 0x42, 0xb0, 0x74, 0x93, 0x33, 0x75, 0xc1, 0xae, 0x2a, 0xd6, 0x4d,
	0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0x87, 0x7b, 0x23, 0xbe, 0xcc, 0x70, 0x15, 0x93, 0x04, 0xdc, 0x39,
	0x7a, 0xce, 0x99, 0x7a, 0xbe, 0x30, 0x93, 0xc3, 0x0c, 0x93, 0xd8, 0xc0, 0xba, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0x20, 0xc5, 0xc8, 0x04, 0x01, 0x00, 0x00,
}
