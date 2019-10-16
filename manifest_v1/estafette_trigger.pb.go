// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_trigger.proto

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

type EstafetteTrigger struct {
	Pipeline             *EstafettePipelineTrigger      `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Release              *EstafetteReleaseTrigger       `protobuf:"bytes,2,opt,name=release,proto3" json:"release,omitempty"`
	Git                  *EstafetteGitTrigger           `protobuf:"bytes,3,opt,name=git,proto3" json:"git,omitempty"`
	Docker               *EstafetteDockerTrigger        `protobuf:"bytes,4,opt,name=docker,proto3" json:"docker,omitempty"`
	Cron                 *EstafetteCronTrigger          `protobuf:"bytes,5,opt,name=cron,proto3" json:"cron,omitempty"`
	PubSub               *EstafettePubSubTrigger        `protobuf:"bytes,6,opt,name=pub_sub,json=pubSub,proto3" json:"pub_sub,omitempty"`
	BuildAction          *EstafetteTriggerBuildAction   `protobuf:"bytes,7,opt,name=build_action,json=buildAction,proto3" json:"build_action,omitempty"`
	ReleaseAction        *EstafetteTriggerReleaseAction `protobuf:"bytes,8,opt,name=release_action,json=releaseAction,proto3" json:"release_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *EstafetteTrigger) Reset()         { *m = EstafetteTrigger{} }
func (m *EstafetteTrigger) String() string { return proto.CompactTextString(m) }
func (*EstafetteTrigger) ProtoMessage()    {}
func (*EstafetteTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ca988580e01498, []int{0}
}

func (m *EstafetteTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteTrigger.Unmarshal(m, b)
}
func (m *EstafetteTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteTrigger.Marshal(b, m, deterministic)
}
func (m *EstafetteTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteTrigger.Merge(m, src)
}
func (m *EstafetteTrigger) XXX_Size() int {
	return xxx_messageInfo_EstafetteTrigger.Size(m)
}
func (m *EstafetteTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteTrigger proto.InternalMessageInfo

func (m *EstafetteTrigger) GetPipeline() *EstafettePipelineTrigger {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *EstafetteTrigger) GetRelease() *EstafetteReleaseTrigger {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *EstafetteTrigger) GetGit() *EstafetteGitTrigger {
	if m != nil {
		return m.Git
	}
	return nil
}

func (m *EstafetteTrigger) GetDocker() *EstafetteDockerTrigger {
	if m != nil {
		return m.Docker
	}
	return nil
}

func (m *EstafetteTrigger) GetCron() *EstafetteCronTrigger {
	if m != nil {
		return m.Cron
	}
	return nil
}

func (m *EstafetteTrigger) GetPubSub() *EstafettePubSubTrigger {
	if m != nil {
		return m.PubSub
	}
	return nil
}

func (m *EstafetteTrigger) GetBuildAction() *EstafetteTriggerBuildAction {
	if m != nil {
		return m.BuildAction
	}
	return nil
}

func (m *EstafetteTrigger) GetReleaseAction() *EstafetteTriggerReleaseAction {
	if m != nil {
		return m.ReleaseAction
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteTrigger)(nil), "manifest.v1.EstafetteTrigger")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_trigger.proto", fileDescriptor_d4ca988580e01498)
}

var fileDescriptor_d4ca988580e01498 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0x83, 0x20, 0x90, 0xa2, 0xc6, 0xf4, 0xd4, 0x70, 0x42, 0xd1, 0x48, 0x50, 0x86, 0x60,
	0x3c, 0x69, 0x88, 0xa0, 0xc6, 0x83, 0x17, 0x44, 0x4f, 0x5e, 0x96, 0x75, 0x94, 0xda, 0x38, 0xda,
	0xa5, 0xeb, 0xf8, 0x33, 0xfe, 0x59, 0x43, 0x69, 0x61, 0x2e, 0xcd, 0xe2, 0x6d, 0x5f, 0xf2, 0x3c,
	0xef, 0xb7, 0x7c, 0x6f, 0x41, 0x7b, 0x19, 0x70, 0xb6, 0x20, 0x89, 0xf2, 0x56, 0x83, 0x3e, 0x49,
	0x54, 0xb0, 0x20, 0x4a, 0x11, 0x5f, 0x49, 0x46, 0x29, 0x91, 0x5e, 0x2c, 0x85, 0x12, 0xb0, 0x91,
	0x81, 0x9a, 0x57, 0x6e, 0x23, 0x66, 0x31, 0x89, 0x18, 0xcf, 0xa9, 0xcd, 0x4b, 0x37, 0x2d, 0x49,
	0x44, 0x82, 0x24, 0x0f, 0x5f, 0xb8, 0x61, 0xca, 0x54, 0x0e, 0xec, 0xba, 0xc1, 0xb9, 0x08, 0xbf,
	0x89, 0xcc, 0xb1, 0x1d, 0x37, 0x1b, 0x4a, 0xc1, 0xff, 0xf7, 0xaf, 0x71, 0x8a, 0xfd, 0x24, 0xc5,
	0x39, 0xf8, 0xba, 0xf0, 0x70, 0x3e, 0x4e, 0x59, 0x34, 0xf7, 0x83, 0x50, 0x31, 0xc1, 0x8d, 0x31,
	0x2c, 0x36, 0xec, 0x49, 0xb2, 0xce, 0xe9, 0x4f, 0x05, 0x1c, 0x3f, 0x5b, 0xf4, 0x63, 0x43, 0xc2,
	0x31, 0xa8, 0xdb, 0x6b, 0xa3, 0x52, 0xab, 0xd4, 0x69, 0x0c, 0xcf, 0xbd, 0x4c, 0xb6, 0xb7, 0x15,
	0xa6, 0x86, 0x32, 0xe2, 0x6c, 0xab, 0xc1, 0x11, 0xa8, 0x99, 0x7d, 0x68, 0x4f, 0x27, 0x9c, 0xb9,
	0x13, 0x66, 0x1b, 0xc8, 0x06, 0x58, 0x09, 0x0e, 0x41, 0x99, 0x32, 0x85, 0xca, 0xda, 0x6d, 0xb9,
	0xdd, 0x17, 0xa6, 0xac, 0xb7, 0x86, 0xe1, 0x1d, 0xa8, 0x6e, 0x0a, 0x42, 0x15, 0xad, 0xb5, 0xdd,
	0xda, 0x93, 0x66, 0xac, 0x69, 0x14, 0x78, 0x0b, 0x2a, 0xeb, 0xc6, 0xd0, 0xbe, 0x56, 0x4f, 0xdc,
	0xea, 0xa3, 0x14, 0xdc, 0x8a, 0x1a, 0x87, 0xf7, 0xa0, 0x66, 0xea, 0x43, 0xd5, 0xa2, 0xa5, 0xd3,
	0x14, 0xbf, 0xa7, 0x78, 0xbb, 0x34, 0xd6, 0x23, 0x7c, 0x05, 0x07, 0xd9, 0x1e, 0x51, 0x4d, 0x47,
	0x74, 0xdc, 0x11, 0x46, 0x9e, 0xac, 0x85, 0xb1, 0xe6, 0x67, 0x0d, 0xbc, 0x1b, 0xe0, 0x1b, 0x38,
	0xfa, 0x5b, 0x31, 0xaa, 0xeb, 0xb8, 0x6e, 0x61, 0x9c, 0x29, 0xc0, 0x04, 0x1e, 0xca, 0xec, 0x38,
	0x79, 0xf8, 0x1c, 0x51, 0xa6, 0xbe, 0x52, 0xec, 0x85, 0x62, 0xb9, 0x7b, 0x52, 0xbb, 0xaf, 0x5e,
	0xc8, 0x7a, 0xa1, 0xe0, 0x4a, 0x06, 0xa1, 0x4a, 0x7a, 0x54, 0x44, 0x01, 0xa7, 0x7d, 0xbb, 0xce,
	0x5f, 0x0d, 0x70, 0x55, 0x3f, 0xb3, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x56, 0x07,
	0xad, 0x07, 0x04, 0x00, 0x00,
}