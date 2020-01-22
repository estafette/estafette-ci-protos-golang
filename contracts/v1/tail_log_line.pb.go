// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/contracts/v1/tail_log_line.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type TailLogLine struct {
	Step                 string                   `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	LogLine              *BuildLogLine            `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
	Image                *BuildLogStepDockerImage `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Duration             *duration.Duration       `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	ExitCode             int64                    `protobuf:"varint,5,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Status               string                   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AutoInjected         bool                     `protobuf:"varint,7,opt,name=auto_injected,json=autoInjected,proto3" json:"auto_injected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TailLogLine) Reset()         { *m = TailLogLine{} }
func (m *TailLogLine) String() string { return proto.CompactTextString(m) }
func (*TailLogLine) ProtoMessage()    {}
func (*TailLogLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c65a3f0e93e9918, []int{0}
}

func (m *TailLogLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TailLogLine.Unmarshal(m, b)
}
func (m *TailLogLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TailLogLine.Marshal(b, m, deterministic)
}
func (m *TailLogLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TailLogLine.Merge(m, src)
}
func (m *TailLogLine) XXX_Size() int {
	return xxx_messageInfo_TailLogLine.Size(m)
}
func (m *TailLogLine) XXX_DiscardUnknown() {
	xxx_messageInfo_TailLogLine.DiscardUnknown(m)
}

var xxx_messageInfo_TailLogLine proto.InternalMessageInfo

func (m *TailLogLine) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

func (m *TailLogLine) GetLogLine() *BuildLogLine {
	if m != nil {
		return m.LogLine
	}
	return nil
}

func (m *TailLogLine) GetImage() *BuildLogStepDockerImage {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *TailLogLine) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TailLogLine) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *TailLogLine) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TailLogLine) GetAutoInjected() bool {
	if m != nil {
		return m.AutoInjected
	}
	return false
}

func init() {
	proto.RegisterType((*TailLogLine)(nil), "estafette.ci.contracts.v1.TailLogLine")
}

func init() {
	proto.RegisterFile("estafette/ci/contracts/v1/tail_log_line.proto", fileDescriptor_4c65a3f0e93e9918)
}

var fileDescriptor_4c65a3f0e93e9918 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x6a, 0xe3, 0x30,
	0x14, 0x85, 0x71, 0x7e, 0x1d, 0x65, 0x66, 0xa3, 0xc5, 0xe0, 0x64, 0x98, 0xc1, 0xb4, 0x8b, 0x7a,
	0x63, 0x99, 0xa4, 0x74, 0xd5, 0xae, 0x92, 0x14, 0x1a, 0xc8, 0xa2, 0xb8, 0xa5, 0x8b, 0x6e, 0x8c,
	0x2c, 0x2b, 0xaa, 0x5a, 0xc5, 0x0a, 0xf6, 0x75, 0x28, 0xf4, 0x8d, 0xfa, 0x14, 0x7d, 0xb4, 0x62,
	0xc5, 0x76, 0x69, 0x21, 0x90, 0x9d, 0x74, 0x75, 0xee, 0xc7, 0xf9, 0x10, 0xf2, 0x79, 0x0e, 0x74,
	0xcd, 0x01, 0x78, 0xc0, 0x64, 0xc0, 0x74, 0x0a, 0x19, 0x65, 0x90, 0x07, 0xbb, 0x49, 0x00, 0x54,
	0xaa, 0x48, 0x69, 0x11, 0x29, 0x99, 0x72, 0xb2, 0xcd, 0x34, 0x68, 0x3c, 0x6a, 0xe2, 0x84, 0x49,
	0xd2, 0xc4, 0xc9, 0x6e, 0x32, 0x26, 0x87, 0x49, 0x71, 0x21, 0x55, 0xf2, 0x03, 0x35, 0xbe, 0x3c,
	0x26, 0x9f, 0x03, 0xdf, 0x46, 0x89, 0x66, 0x2f, 0x3c, 0x8b, 0xe4, 0x86, 0x8a, 0x7a, 0xf9, 0xbf,
	0xd0, 0x5a, 0x28, 0x1e, 0x98, 0x5b, 0x5c, 0xac, 0x83, 0xa4, 0xc8, 0x28, 0x48, 0x9d, 0xee, 0xdf,
	0x4f, 0x3e, 0x5a, 0x68, 0x78, 0x4f, 0xa5, 0x5a, 0x69, 0xb1, 0x92, 0x29, 0xc7, 0x18, 0x75, 0x4a,
	0x94, 0x63, 0xb9, 0x96, 0x37, 0x08, 0xcd, 0x19, 0xcf, 0x90, 0x5d, 0x57, 0x72, 0x5a, 0xae, 0xe5,
	0x0d, 0xa7, 0x67, 0xe4, 0xa0, 0x1e, 0x99, 0x95, 0x9d, 0x2a, 0x5c, 0xd8, 0x57, 0x15, 0xf7, 0x06,
	0x75, 0x4d, 0x2d, 0xa7, 0x6d, 0x00, 0xd3, 0x23, 0x00, 0x77, 0xc0, 0xb7, 0x0b, 0x63, 0xb4, 0x2c,
	0x37, 0xc3, 0x3d, 0x00, 0x5f, 0x20, 0xbb, 0x76, 0x70, 0x3a, 0x06, 0x36, 0x22, 0x7b, 0x49, 0x52,
	0x4b, 0x92, 0x45, 0x15, 0x08, 0x9b, 0x28, 0xfe, 0x8b, 0x06, 0xfc, 0x55, 0x42, 0xc4, 0x74, 0xc2,
	0x9d, 0xae, 0x6b, 0x79, 0xed, 0xd0, 0x2e, 0x07, 0x73, 0x9d, 0x70, 0xfc, 0x07, 0xf5, 0x72, 0xa0,
	0x50, 0xe4, 0x4e, 0xcf, 0x78, 0x57, 0x37, 0x7c, 0x8a, 0x7e, 0xd3, 0x02, 0x74, 0x24, 0xd3, 0x67,
	0xce, 0x80, 0x27, 0x4e, 0xdf, 0xb5, 0x3c, 0x3b, 0xfc, 0x55, 0x0e, 0x97, 0xd5, 0x6c, 0xf6, 0x86,
	0xfe, 0x31, 0xbd, 0x39, 0x2c, 0x74, 0x6b, 0x3d, 0x5e, 0x09, 0x09, 0x4f, 0x45, 0x4c, 0x98, 0xde,
	0x04, 0x5f, 0xbf, 0xd9, 0x9c, 0x7c, 0x26, 0x7d, 0xd3, 0x3d, 0xf7, 0x85, 0x56, 0x34, 0x15, 0xdf,
	0x7e, 0xf9, 0xbd, 0x35, 0xba, 0x6e, 0xd8, 0x73, 0x49, 0xe6, 0x0d, 0xfb, 0x61, 0x12, 0xf7, 0xcc,
	0xde, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xa2, 0x77, 0xb1, 0x9f, 0x02, 0x00, 0x00,
}