// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/tail_log_line.proto

package contracts_v1

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
	return fileDescriptor_32e41484a0b59d6d, []int{0}
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
	proto.RegisterType((*TailLogLine)(nil), "contracts.v1.TailLogLine")
}

func init() { proto.RegisterFile("contracts.v1/tail_log_line.proto", fileDescriptor_32e41484a0b59d6d) }

var fileDescriptor_32e41484a0b59d6d = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x53, 0xfe, 0xb3, 0xe0, 0xa5, 0x07, 0x53, 0x31, 0x31, 0x55, 0x63, 0xd2, 0x0b, 0xdb,
	0xa0, 0xf1, 0xe4, 0xc1, 0x88, 0x5c, 0x48, 0x38, 0x55, 0x4f, 0x5e, 0x9a, 0xed, 0x76, 0x58, 0x57,
	0x97, 0x0e, 0x69, 0xa7, 0xc4, 0x4f, 0xe4, 0xe7, 0x34, 0x5d, 0x0a, 0x8a, 0xd1, 0xdb, 0xf6, 0xf5,
	0xf7, 0xde, 0x9b, 0x19, 0xe6, 0x4b, 0xcc, 0x28, 0x17, 0x92, 0x0a, 0xbe, 0x99, 0x84, 0x24, 0xb4,
	0x89, 0x0d, 0xaa, 0xd8, 0xe8, 0x0c, 0xf8, 0x3a, 0x47, 0x42, 0x77, 0xf8, 0x93, 0x18, 0x9d, 0x29,
	0x44, 0x65, 0x20, 0xb4, 0xff, 0x92, 0x72, 0x19, 0xa6, 0x65, 0x2e, 0x48, 0x63, 0xb6, 0xa5, 0x47,
	0xe7, 0x07, 0x79, 0x49, 0xa9, 0x4d, 0xfa, 0x2b, 0x70, 0xc4, 0xff, 0x41, 0x0a, 0x82, 0x75, 0x9c,
	0xa2, 0x7c, 0x87, 0x3c, 0xd6, 0x2b, 0xa1, 0x6a, 0xfe, 0xe2, 0xb3, 0xc1, 0x06, 0xcf, 0x42, 0x9b,
	0x05, 0xaa, 0x85, 0xce, 0xc0, 0x75, 0x59, 0xab, 0x42, 0x3d, 0xc7, 0x77, 0x82, 0x7e, 0x64, 0xdf,
	0xee, 0x2d, 0xeb, 0xed, 0x5a, 0xbc, 0x86, 0xef, 0x04, 0x83, 0xeb, 0xc3, 0x1a, 0x3e, 0xad, 0x6a,
	0xea, 0x84, 0xa8, 0x6b, 0xea, 0xa8, 0x3b, 0xd6, 0xb6, 0x4d, 0x5e, 0xd3, 0x7a, 0xae, 0xfe, 0xf6,
	0x3c, 0x11, 0xac, 0x67, 0x76, 0xae, 0x79, 0x05, 0x47, 0x5b, 0x4f, 0xd5, 0xb9, 0x5b, 0xde, 0x6b,
	0x59, 0xff, 0x09, 0xdf, 0x5e, 0x87, 0xef, 0xae, 0xc3, 0x67, 0x35, 0x10, 0xed, 0x51, 0xf7, 0x94,
	0xf5, 0xe1, 0x43, 0x53, 0x2c, 0x31, 0x05, 0xaf, 0xed, 0x3b, 0x41, 0x33, 0xea, 0x55, 0xc2, 0x23,
	0xa6, 0xe0, 0x1e, 0xb3, 0x4e, 0x41, 0x82, 0xca, 0xc2, 0xeb, 0xd8, 0xed, 0xea, 0x2f, 0xf7, 0x92,
	0x1d, 0x89, 0x92, 0x30, 0xd6, 0xd9, 0x1b, 0x48, 0x82, 0xd4, 0xeb, 0xfa, 0x4e, 0xd0, 0x8b, 0x86,
	0x95, 0x38, 0xaf, 0xb5, 0xe9, 0xc3, 0xcb, 0xbd, 0xd2, 0xf4, 0x5a, 0x26, 0x5c, 0xe2, 0x2a, 0x84,
	0x82, 0xc4, 0x12, 0x88, 0xe0, 0xfb, 0x35, 0x96, 0x7a, 0xbc, 0xdf, 0x70, 0xac, 0xd0, 0x88, 0x4c,
	0x85, 0x7b, 0x21, 0xde, 0x4c, 0x92, 0x8e, 0x9d, 0xfc, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x2e, 0xd7, 0x8e, 0x17, 0x02, 0x00, 0x00,
}