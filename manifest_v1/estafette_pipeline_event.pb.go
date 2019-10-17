// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_pipeline_event.proto

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

type EstafettePipelineEvent struct {
	BuildVersion         string   `protobuf:"bytes,1,opt,name=build_version,json=buildVersion,proto3" json:"build_version,omitempty"`
	RepoSource           string   `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3" json:"repo_source,omitempty"`
	RepoOwner            string   `protobuf:"bytes,3,opt,name=repo_owner,json=repoOwner,proto3" json:"repo_owner,omitempty"`
	RepoName             string   `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoBranch           string   `protobuf:"bytes,5,opt,name=repo_branch,json=repoBranch,proto3" json:"repo_branch,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Event                string   `protobuf:"bytes,7,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafettePipelineEvent) Reset()         { *m = EstafettePipelineEvent{} }
func (m *EstafettePipelineEvent) String() string { return proto.CompactTextString(m) }
func (*EstafettePipelineEvent) ProtoMessage()    {}
func (*EstafettePipelineEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cd8e4dfb4a4d7e, []int{0}
}

func (m *EstafettePipelineEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafettePipelineEvent.Unmarshal(m, b)
}
func (m *EstafettePipelineEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafettePipelineEvent.Marshal(b, m, deterministic)
}
func (m *EstafettePipelineEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafettePipelineEvent.Merge(m, src)
}
func (m *EstafettePipelineEvent) XXX_Size() int {
	return xxx_messageInfo_EstafettePipelineEvent.Size(m)
}
func (m *EstafettePipelineEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafettePipelineEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafettePipelineEvent proto.InternalMessageInfo

func (m *EstafettePipelineEvent) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoSource() string {
	if m != nil {
		return m.RepoSource
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoOwner() string {
	if m != nil {
		return m.RepoOwner
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoBranch() string {
	if m != nil {
		return m.RepoBranch
	}
	return ""
}

func (m *EstafettePipelineEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EstafettePipelineEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafettePipelineEvent)(nil), "manifest.v1.EstafettePipelineEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_pipeline_event.proto", fileDescriptor_e0cd8e4dfb4a4d7e)
}

var fileDescriptor_e0cd8e4dfb4a4d7e = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xbf, 0xaf, 0xd1, 0x4e, 0x75, 0x33, 0x48, 0x19, 0x10, 0x51, 0x74, 0x23, 0x42,
	0x52, 0x82, 0x7b, 0x17, 0x95, 0x2e, 0x5c, 0xf8, 0x07, 0x85, 0x2e, 0x74, 0x31, 0x4c, 0xc6, 0xdb,
	0x74, 0x20, 0x99, 0x09, 0x33, 0x37, 0xf1, 0x9d, 0x7c, 0x33, 0xdf, 0x42, 0x72, 0xdb, 0x34, 0xee,
	0xe6, 0xfc, 0xce, 0xe1, 0xc2, 0xfc, 0xd8, 0x4d, 0xa5, 0xac, 0x59, 0x43, 0xc0, 0xb4, 0xcd, 0xe6,
	0x10, 0x50, 0xad, 0x01, 0x11, 0x64, 0x6d, 0x6a, 0x28, 0x8d, 0x05, 0x09, 0x2d, 0x58, 0x4c, 0x6b,
	0xef, 0xd0, 0xf1, 0xe9, 0x9f, 0xed, 0xe5, 0x4f, 0xc4, 0x66, 0xcb, 0x7e, 0xff, 0xb2, 0x9b, 0x2f,
	0xbb, 0x35, 0xbf, 0x62, 0xc7, 0x79, 0x63, 0xca, 0x4f, 0xd9, 0x82, 0x0f, 0xc6, 0x59, 0x11, 0x5d,
	0x44, 0xd7, 0x93, 0xd7, 0x23, 0x82, 0xab, 0x2d, 0xe3, 0xe7, 0x6c, 0xea, 0xa1, 0x76, 0x32, 0xb8,
	0xc6, 0x6b, 0x10, 0x23, 0x9a, 0xb0, 0x0e, 0xbd, 0x11, 0xe1, 0x67, 0x8c, 0x92, 0x74, 0x5f, 0x16,
	0xbc, 0xf8, 0x47, 0xfd, 0xa4, 0x23, 0xcf, 0x1d, 0xe0, 0xa7, 0x8c, 0x82, 0xb4, 0xaa, 0x02, 0xf1,
	0x9f, 0xda, 0xc3, 0x0e, 0x3c, 0xa9, 0x0a, 0xf6, 0xc7, 0x73, 0xaf, 0xac, 0xde, 0x88, 0xf1, 0x70,
	0x7c, 0x41, 0x84, 0xcf, 0x58, 0x1c, 0x50, 0x61, 0x13, 0x44, 0x4c, 0xdd, 0x2e, 0xf1, 0x13, 0x36,
	0xa6, 0x1f, 0x8b, 0x03, 0xc2, 0xdb, 0xb0, 0xf8, 0x78, 0xbf, 0x2b, 0x0c, 0x6e, 0x9a, 0x3c, 0xd5,
	0xae, 0x1a, 0x2c, 0x0d, 0xaf, 0x44, 0x9b, 0x44, 0x3b, 0x8b, 0x5e, 0x69, 0x0c, 0x49, 0xe1, 0x4a,
	0x65, 0x8b, 0x79, 0x6f, 0x4b, 0xb6, 0xd9, 0xf7, 0x48, 0xec, 0x5d, 0xa5, 0xf7, 0x0f, 0xe9, 0x63,
	0x2f, 0x72, 0x95, 0xe5, 0x31, 0xc9, 0xbd, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x85, 0xb9, 0x0c,
	0xa5, 0x8a, 0x01, 0x00, 0x00,
}
