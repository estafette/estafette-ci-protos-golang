// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/build_log_step_docker_image.proto

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

type BuildLogStepDockerImage struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string             `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	IsPulled             bool               `protobuf:"varint,3,opt,name=is_pulled,json=isPulled,proto3" json:"is_pulled,omitempty"`
	ImageSize            int64              `protobuf:"varint,4,opt,name=image_size,json=imageSize,proto3" json:"image_size,omitempty"`
	PullDuration         *duration.Duration `protobuf:"bytes,5,opt,name=pull_duration,json=pullDuration,proto3" json:"pull_duration,omitempty"`
	Error                string             `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	IsTrusted            bool               `protobuf:"varint,7,opt,name=is_trusted,json=isTrusted,proto3" json:"is_trusted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildLogStepDockerImage) Reset()         { *m = BuildLogStepDockerImage{} }
func (m *BuildLogStepDockerImage) String() string { return proto.CompactTextString(m) }
func (*BuildLogStepDockerImage) ProtoMessage()    {}
func (*BuildLogStepDockerImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b85319001c270955, []int{0}
}

func (m *BuildLogStepDockerImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildLogStepDockerImage.Unmarshal(m, b)
}
func (m *BuildLogStepDockerImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildLogStepDockerImage.Marshal(b, m, deterministic)
}
func (m *BuildLogStepDockerImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildLogStepDockerImage.Merge(m, src)
}
func (m *BuildLogStepDockerImage) XXX_Size() int {
	return xxx_messageInfo_BuildLogStepDockerImage.Size(m)
}
func (m *BuildLogStepDockerImage) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildLogStepDockerImage.DiscardUnknown(m)
}

var xxx_messageInfo_BuildLogStepDockerImage proto.InternalMessageInfo

func (m *BuildLogStepDockerImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildLogStepDockerImage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BuildLogStepDockerImage) GetIsPulled() bool {
	if m != nil {
		return m.IsPulled
	}
	return false
}

func (m *BuildLogStepDockerImage) GetImageSize() int64 {
	if m != nil {
		return m.ImageSize
	}
	return 0
}

func (m *BuildLogStepDockerImage) GetPullDuration() *duration.Duration {
	if m != nil {
		return m.PullDuration
	}
	return nil
}

func (m *BuildLogStepDockerImage) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BuildLogStepDockerImage) GetIsTrusted() bool {
	if m != nil {
		return m.IsTrusted
	}
	return false
}

func init() {
	proto.RegisterType((*BuildLogStepDockerImage)(nil), "contracts.v1.BuildLogStepDockerImage")
}

func init() {
	proto.RegisterFile("contracts.v1/build_log_step_docker_image.proto", fileDescriptor_b85319001c270955)
}

var fileDescriptor_b85319001c270955 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0xde, 0x68, 0x4c, 0x91, 0x90, 0x85, 0x84, 0xcb, 0x4d, 0x11, 0x53, 0x96, 0x3a,
	0x2a, 0xec, 0x20, 0xb5, 0x65, 0xa8, 0xc4, 0x80, 0x52, 0xc4, 0xc0, 0x62, 0x39, 0xc9, 0xa9, 0xb1,
	0x48, 0xe3, 0xc8, 0x76, 0x3a, 0x94, 0x37, 0xe2, 0xe9, 0x78, 0x04, 0x14, 0x87, 0x06, 0xd8, 0xce,
	0xf9, 0xf3, 0xe5, 0xf7, 0x67, 0x23, 0x9a, 0xaa, 0xc2, 0x6a, 0x9e, 0x5a, 0x43, 0xb7, 0xd3, 0x28,
	0xa9, 0x64, 0x9e, 0xb1, 0x5c, 0x09, 0x66, 0x2c, 0x94, 0x2c, 0x53, 0xe9, 0x3b, 0x68, 0x26, 0x37,
	0x5c, 0x00, 0x2d, 0xb5, 0xb2, 0x0a, 0x8f, 0xfe, 0xf2, 0x67, 0x57, 0x42, 0x29, 0x91, 0x43, 0xe4,
	0xbe, 0x25, 0xd5, 0x3a, 0xca, 0x2a, 0xcd, 0xad, 0x54, 0x45, 0x43, 0x5f, 0x7f, 0x79, 0xe8, 0x74,
	0x56, 0x77, 0x3e, 0x2a, 0xb1, 0xb2, 0x50, 0x2e, 0x5c, 0xe1, 0xb2, 0xee, 0xc3, 0x18, 0xf5, 0x0a,
	0xbe, 0x01, 0xe2, 0x05, 0x5e, 0xe8, 0xc7, 0x6e, 0xc6, 0xc7, 0xa8, 0x6b, 0xb9, 0x20, 0x1d, 0x17,
	0xd5, 0x23, 0x3e, 0x47, 0xbe, 0x34, 0xac, 0xac, 0xf2, 0x1c, 0x32, 0xd2, 0x0d, 0xbc, 0x70, 0x18,
	0x0f, 0xa5, 0x79, 0x72, 0x3b, 0xbe, 0x44, 0xc8, 0xb9, 0x31, 0x23, 0x77, 0x40, 0x7a, 0x81, 0x17,
	0x76, 0x63, 0xdf, 0x25, 0x2b, 0xb9, 0x03, 0x7c, 0x87, 0x8e, 0xea, 0x1f, 0xd9, 0x5e, 0x8a, 0xf4,
	0x03, 0x2f, 0x3c, 0xbc, 0x19, 0xd3, 0xc6, 0x9a, 0xee, 0xad, 0xe9, 0xe2, 0x07, 0x88, 0x47, 0x35,
	0xbf, 0xdf, 0xf0, 0x09, 0xea, 0x83, 0xd6, 0x4a, 0x93, 0x81, 0xf3, 0x69, 0x16, 0x77, 0xa8, 0x61,
	0x56, 0x57, 0xc6, 0x42, 0x46, 0x0e, 0x9c, 0x92, 0x2f, 0xcd, 0x73, 0x13, 0xcc, 0x3e, 0xd0, 0x85,
	0x54, 0x14, 0x8c, 0xe5, 0x6b, 0xb0, 0x16, 0x68, 0x2a, 0xff, 0x3d, 0xf1, 0xeb, 0xbd, 0x90, 0xf6,
	0xad, 0x4a, 0x68, 0xaa, 0x36, 0x51, 0x4b, 0xfd, 0x4e, 0x93, 0x54, 0x4e, 0x5a, 0x7e, 0x22, 0x54,
	0xce, 0x0b, 0x11, 0xb5, 0x01, 0xdb, 0x4e, 0x3f, 0x3b, 0xe3, 0x87, 0xb6, 0x7c, 0xbe, 0xa4, 0xf3,
	0xb6, 0xfc, 0x65, 0x9a, 0x0c, 0xdc, 0x95, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x54, 0x52,
	0x1b, 0x0e, 0xd6, 0x01, 0x00, 0x00,
}
