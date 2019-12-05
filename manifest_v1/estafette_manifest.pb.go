// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/manifest/v1/estafette_manifest.proto

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

type EstafetteManifest struct {
	Builder              *EstafetteBuilder   `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"`
	Labels               map[string]string   `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Version              *EstafetteVersion   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	GlobalEnvVars        map[string]string   `protobuf:"bytes,4,rep,name=global_env_vars,json=globalEnvVars,proto3" json:"global_env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Triggers             []*EstafetteTrigger `protobuf:"bytes,5,rep,name=triggers,proto3" json:"triggers,omitempty"`
	Stages               []*EstafetteStage   `protobuf:"bytes,6,rep,name=stages,proto3" json:"stages,omitempty"`
	Releases             []*EstafetteRelease `protobuf:"bytes,7,rep,name=releases,proto3" json:"releases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EstafetteManifest) Reset()         { *m = EstafetteManifest{} }
func (m *EstafetteManifest) String() string { return proto.CompactTextString(m) }
func (*EstafetteManifest) ProtoMessage()    {}
func (*EstafetteManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71927f69dda30097, []int{0}
}

func (m *EstafetteManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteManifest.Unmarshal(m, b)
}
func (m *EstafetteManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteManifest.Marshal(b, m, deterministic)
}
func (m *EstafetteManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteManifest.Merge(m, src)
}
func (m *EstafetteManifest) XXX_Size() int {
	return xxx_messageInfo_EstafetteManifest.Size(m)
}
func (m *EstafetteManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteManifest.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteManifest proto.InternalMessageInfo

func (m *EstafetteManifest) GetBuilder() *EstafetteBuilder {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *EstafetteManifest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *EstafetteManifest) GetVersion() *EstafetteVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *EstafetteManifest) GetGlobalEnvVars() map[string]string {
	if m != nil {
		return m.GlobalEnvVars
	}
	return nil
}

func (m *EstafetteManifest) GetTriggers() []*EstafetteTrigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func (m *EstafetteManifest) GetStages() []*EstafetteStage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *EstafetteManifest) GetReleases() []*EstafetteRelease {
	if m != nil {
		return m.Releases
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteManifest)(nil), "estafette.ci.manifest.v1.EstafetteManifest")
	proto.RegisterMapType((map[string]string)(nil), "estafette.ci.manifest.v1.EstafetteManifest.GlobalEnvVarsEntry")
	proto.RegisterMapType((map[string]string)(nil), "estafette.ci.manifest.v1.EstafetteManifest.LabelsEntry")
}

func init() {
	proto.RegisterFile("estafette/ci/manifest/v1/estafette_manifest.proto", fileDescriptor_71927f69dda30097)
}

var fileDescriptor_71927f69dda30097 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0xab, 0xd3, 0x30,
	0x18, 0x87, 0x69, 0xe7, 0x7a, 0x34, 0x43, 0xd4, 0xe0, 0x45, 0x18, 0x5e, 0x1c, 0xbc, 0x1a, 0xc2,
	0x52, 0x7b, 0xbc, 0xf0, 0x1f, 0xc8, 0x61, 0x3a, 0xbd, 0x51, 0x94, 0x28, 0xbb, 0xf0, 0xa6, 0xa4,
	0x35, 0x8b, 0xc1, 0xac, 0x95, 0x24, 0x0b, 0xee, 0x13, 0x09, 0x7e, 0x4a, 0x69, 0x9a, 0x66, 0x83,
	0xc3, 0x20, 0xbd, 0x4b, 0xf3, 0xf2, 0xfc, 0xf2, 0x24, 0xef, 0x5b, 0x50, 0x30, 0x6d, 0xe8, 0x96,
	0x19, 0xc3, 0xf2, 0x5a, 0xe4, 0x3b, 0xda, 0x88, 0x2d, 0xd3, 0x26, 0xb7, 0x45, 0x1e, 0x0a, 0xe5,
	0xb0, 0x8b, 0x7f, 0xab, 0xd6, 0xb4, 0x10, 0x85, 0x0a, 0xae, 0x05, 0x0e, 0x45, 0x5b, 0xcc, 0x9f,
	0x46, 0x84, 0x55, 0x7b, 0x21, 0x7f, 0x30, 0xd5, 0x67, 0x45, 0x11, 0x8a, 0x49, 0x46, 0x35, 0xf3,
	0x04, 0x8e, 0x20, 0xb4, 0xa1, 0x9c, 0x8d, 0x38, 0xc1, 0x28, 0xc1, 0xf9, 0x28, 0x27, 0xcb, 0x94,
	0x16, 0x6d, 0xd3, 0x13, 0x8f, 0xff, 0x4e, 0xc1, 0x83, 0xf5, 0x50, 0xfb, 0xe4, 0x01, 0xf8, 0x0e,
	0x5c, 0xf8, 0xcb, 0xa2, 0xe4, 0x32, 0x59, 0xcc, 0xae, 0x9e, 0xe0, 0x73, 0x2f, 0x87, 0x03, 0xbd,
	0xea, 0x09, 0x32, 0xa0, 0xf0, 0x33, 0xc8, 0x24, 0xad, 0x98, 0xd4, 0x28, 0xbd, 0x9c, 0x2c, 0x66,
	0x57, 0xcf, 0x23, 0x42, 0x06, 0x05, 0xfc, 0xd1, 0x91, 0xeb, 0xc6, 0xa8, 0x03, 0xf1, 0x31, 0x9d,
	0x96, 0xb7, 0x47, 0x93, 0x68, 0xad, 0x4d, 0x4f, 0x90, 0x01, 0x85, 0x5b, 0x70, 0x8f, 0xcb, 0xb6,
	0xa2, 0xb2, 0x64, 0x8d, 0x2d, 0x2d, 0x55, 0x1a, 0xdd, 0x72, 0x7e, 0x6f, 0xc6, 0xf8, 0x7d, 0x70,
	0x11, 0xeb, 0xc6, 0x6e, 0xa8, 0xf2, 0x9a, 0x77, 0xf9, 0xe9, 0x1e, 0x7c, 0x0f, 0x6e, 0xfb, 0xee,
	0x68, 0x34, 0x75, 0x07, 0xc4, 0xe8, 0x7e, 0xeb, 0x11, 0x12, 0x58, 0x78, 0x0d, 0x32, 0x37, 0x15,
	0x1a, 0x65, 0x2e, 0x65, 0x11, 0x91, 0xf2, 0xb5, 0x03, 0x88, 0xe7, 0x3a, 0x13, 0x3f, 0x89, 0x1a,
	0x5d, 0x44, 0x9b, 0x90, 0x1e, 0x21, 0x81, 0x9d, 0xbf, 0x04, 0xb3, 0x93, 0xb6, 0xc0, 0xfb, 0x60,
	0xf2, 0x8b, 0x1d, 0xdc, 0x84, 0xdc, 0x21, 0xdd, 0x12, 0x3e, 0x04, 0x53, 0x4b, 0xe5, 0x9e, 0xa1,
	0xd4, 0xed, 0xf5, 0x1f, 0xaf, 0xd2, 0x17, 0xc9, 0xfc, 0x1a, 0xc0, 0x9b, 0x2f, 0x36, 0x26, 0x61,
	0xf5, 0x07, 0x3c, 0xaa, 0xdb, 0xdd, 0x59, 0xef, 0x2f, 0xc9, 0xf7, 0xd7, 0x5c, 0x98, 0x9f, 0xfb,
	0x0a, 0xd7, 0xed, 0xee, 0x38, 0xef, 0xc7, 0xd5, 0xb2, 0x16, 0x4b, 0x37, 0xf4, 0x7a, 0xc9, 0x5b,
	0x49, 0x1b, 0x1e, 0x7e, 0x90, 0xd2, 0x16, 0xff, 0x52, 0x14, 0x2e, 0x8e, 0xdf, 0x0a, 0x1c, 0xfa,
	0xbc, 0x29, 0xaa, 0xcc, 0x51, 0xcf, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xbf, 0xc9, 0xb2,
	0x71, 0x04, 0x00, 0x00,
}
