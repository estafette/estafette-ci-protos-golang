// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/contracts/v1/build_log.proto

package v1

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

type BuildLog struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RepoSource           string               `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3" json:"repo_source,omitempty"`
	RepoOwner            string               `protobuf:"bytes,3,opt,name=repo_owner,json=repoOwner,proto3" json:"repo_owner,omitempty"`
	RepoName             string               `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoBranch           string               `protobuf:"bytes,5,opt,name=repo_branch,json=repoBranch,proto3" json:"repo_branch,omitempty"`
	RepoRevision         string               `protobuf:"bytes,6,opt,name=repo_revision,json=repoRevision,proto3" json:"repo_revision,omitempty"`
	BuildId              string               `protobuf:"bytes,7,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Steps                []*BuildLogStep      `protobuf:"bytes,8,rep,name=steps,proto3" json:"steps,omitempty"`
	InsertedAtTime       *timestamp.Timestamp `protobuf:"bytes,9,opt,name=inserted_at_time,json=insertedAtTime,proto3" json:"inserted_at_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BuildLog) Reset()         { *m = BuildLog{} }
func (m *BuildLog) String() string { return proto.CompactTextString(m) }
func (*BuildLog) ProtoMessage()    {}
func (*BuildLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e291835105cb0da, []int{0}
}

func (m *BuildLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildLog.Unmarshal(m, b)
}
func (m *BuildLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildLog.Marshal(b, m, deterministic)
}
func (m *BuildLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildLog.Merge(m, src)
}
func (m *BuildLog) XXX_Size() int {
	return xxx_messageInfo_BuildLog.Size(m)
}
func (m *BuildLog) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildLog.DiscardUnknown(m)
}

var xxx_messageInfo_BuildLog proto.InternalMessageInfo

func (m *BuildLog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuildLog) GetRepoSource() string {
	if m != nil {
		return m.RepoSource
	}
	return ""
}

func (m *BuildLog) GetRepoOwner() string {
	if m != nil {
		return m.RepoOwner
	}
	return ""
}

func (m *BuildLog) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *BuildLog) GetRepoBranch() string {
	if m != nil {
		return m.RepoBranch
	}
	return ""
}

func (m *BuildLog) GetRepoRevision() string {
	if m != nil {
		return m.RepoRevision
	}
	return ""
}

func (m *BuildLog) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildLog) GetSteps() []*BuildLogStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *BuildLog) GetInsertedAtTime() *timestamp.Timestamp {
	if m != nil {
		return m.InsertedAtTime
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildLog)(nil), "estafette.ci.contracts.v1.BuildLog")
}

func init() {
	proto.RegisterFile("estafette/ci/contracts/v1/build_log.proto", fileDescriptor_3e291835105cb0da)
}

var fileDescriptor_3e291835105cb0da = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xc2, 0x30,
	0x1c, 0xc5, 0xb3, 0x21, 0x30, 0x8a, 0x12, 0xd3, 0x53, 0xc1, 0x10, 0x88, 0x1e, 0xc4, 0x03, 0x5d,
	0xc0, 0xab, 0x1e, 0x04, 0x3d, 0x98, 0x18, 0x35, 0xc3, 0x78, 0xf0, 0xb2, 0x6c, 0x5d, 0x19, 0x4d,
	0xb6, 0x75, 0x69, 0x3b, 0x3c, 0xf8, 0x8d, 0xfc, 0x4e, 0x7e, 0x17, 0xd3, 0x8e, 0x8d, 0x78, 0x20,
	0xf1, 0xb6, 0xff, 0xef, 0xff, 0xde, 0x6b, 0xf6, 0xfe, 0xe0, 0x8a, 0x4a, 0x15, 0xac, 0xa9, 0x52,
	0xd4, 0x25, 0xcc, 0x25, 0x3c, 0x53, 0x22, 0x20, 0x4a, 0xba, 0xdb, 0x99, 0x1b, 0x16, 0x2c, 0x89,
	0xfc, 0x84, 0xc7, 0x38, 0x17, 0x5c, 0x71, 0xd8, 0xaf, 0xa5, 0x98, 0x30, 0x5c, 0x4b, 0xf1, 0x76,
	0x36, 0xc0, 0xff, 0x48, 0xf1, 0xa5, 0xa2, 0x79, 0x19, 0x35, 0x18, 0xc5, 0x9c, 0xc7, 0x09, 0x75,
	0xcd, 0x14, 0x16, 0x6b, 0x57, 0xb1, 0x54, 0x47, 0xa4, 0x3b, 0xc1, 0xf9, 0x8f, 0x0d, 0x9c, 0x85,
	0x76, 0x3e, 0xf1, 0x18, 0xf6, 0x80, 0xcd, 0x22, 0x64, 0x8d, 0xad, 0x49, 0xc7, 0xb3, 0x59, 0x04,
	0x47, 0xa0, 0x2b, 0x68, 0xce, 0x7d, 0xc9, 0x0b, 0x41, 0x28, 0xb2, 0xcd, 0x02, 0x68, 0xb4, 0x32,
	0x04, 0x0e, 0x81, 0x99, 0x7c, 0xfe, 0x99, 0x51, 0x81, 0x1a, 0x66, 0xdf, 0xd1, 0xe4, 0x45, 0x03,
	0x78, 0x06, 0xcc, 0xe0, 0x67, 0x41, 0x4a, 0xd1, 0x91, 0xd9, 0x3a, 0x1a, 0x3c, 0x07, 0x29, 0xad,
	0xc3, 0x43, 0x11, 0x64, 0x64, 0x83, 0x9a, 0xfb, 0xf0, 0x85, 0x21, 0xf0, 0x02, 0x9c, 0x18, 0x81,
	0xa0, 0x5b, 0x26, 0x19, 0xcf, 0x50, 0xcb, 0x48, 0x8e, 0x35, 0xf4, 0x76, 0x0c, 0xf6, 0x81, 0x53,
	0xfe, 0x38, 0x8b, 0x50, 0xdb, 0xec, 0xdb, 0x66, 0x7e, 0x8c, 0xe0, 0x2d, 0x68, 0xea, 0x26, 0x24,
	0x72, 0xc6, 0x8d, 0x49, 0x77, 0x7e, 0x89, 0x0f, 0xd6, 0x8a, 0xab, 0x06, 0x56, 0x8a, 0xe6, 0x5e,
	0xe9, 0x82, 0xf7, 0xe0, 0x94, 0x65, 0x92, 0x0a, 0x45, 0x23, 0x3f, 0x50, 0xbe, 0x2e, 0x0e, 0x75,
	0xc6, 0xd6, 0xa4, 0x3b, 0x1f, 0xe0, 0xb2, 0x55, 0x5c, 0xb5, 0x8a, 0xdf, 0xaa, 0x56, 0xbd, 0x5e,
	0xe5, 0xb9, 0x53, 0x1a, 0x2e, 0xbe, 0xc0, 0x90, 0xf0, 0xf4, 0xf0, 0xd3, 0xaf, 0xd6, 0xc7, 0x4d,
	0xcc, 0xd4, 0xa6, 0x08, 0x31, 0xe1, 0xa9, 0xbb, 0x3f, 0x6f, 0xfd, 0x35, 0x25, 0x6c, 0x6a, 0x1e,
	0x92, 0xd3, 0x98, 0x27, 0x41, 0x16, 0xff, 0x39, 0xfb, 0xb7, 0xdd, 0x7f, 0xa8, 0xb3, 0x97, 0x0c,
	0x2f, 0xeb, 0xec, 0xf7, 0x59, 0xd8, 0x32, 0xbe, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca,
	0x83, 0x95, 0x12, 0x7c, 0x02, 0x00, 0x00,
}
