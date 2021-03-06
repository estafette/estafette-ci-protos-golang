// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estafette/ci/contracts/v1/build.proto

package v1

import (
	fmt "fmt"
	manifest_v1 "github.com/estafette/estafette-ci-protos-golang/manifest_v1"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Build struct {
	Id                   string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RepoSource           string                          `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3" json:"repo_source,omitempty"`
	RepoOwner            string                          `protobuf:"bytes,3,opt,name=repo_owner,json=repoOwner,proto3" json:"repo_owner,omitempty"`
	RepoName             string                          `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoBranch           string                          `protobuf:"bytes,5,opt,name=repo_branch,json=repoBranch,proto3" json:"repo_branch,omitempty"`
	RepoRevision         string                          `protobuf:"bytes,6,opt,name=repo_revision,json=repoRevision,proto3" json:"repo_revision,omitempty"`
	BuildVersion         string                          `protobuf:"bytes,7,opt,name=build_version,json=buildVersion,proto3" json:"build_version,omitempty"`
	BuildStatus          string                          `protobuf:"bytes,8,opt,name=build_status,json=buildStatus,proto3" json:"build_status,omitempty"`
	Labels               []*Label                        `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty"`
	ReleaseTargets       []*ReleaseTarget                `protobuf:"bytes,10,rep,name=release_targets,json=releaseTargets,proto3" json:"release_targets,omitempty"`
	Manifest             string                          `protobuf:"bytes,11,opt,name=manifest,proto3" json:"manifest,omitempty"`
	ManifestWithDefaults string                          `protobuf:"bytes,12,opt,name=manifest_with_defaults,json=manifestWithDefaults,proto3" json:"manifest_with_defaults,omitempty"`
	Commits              []*GitCommit                    `protobuf:"bytes,13,rep,name=commits,proto3" json:"commits,omitempty"`
	Triggers             []*manifest_v1.EstafetteTrigger `protobuf:"bytes,14,rep,name=triggers,proto3" json:"triggers,omitempty"`
	Events               []*manifest_v1.EstafetteEvent   `protobuf:"bytes,15,rep,name=events,proto3" json:"events,omitempty"`
	InsertedAtTime       *timestamp.Timestamp            `protobuf:"bytes,16,opt,name=inserted_at_time,json=insertedAtTime,proto3" json:"inserted_at_time,omitempty"`
	UpdatedAtTime        *timestamp.Timestamp            `protobuf:"bytes,17,opt,name=updated_at_time,json=updatedAtTime,proto3" json:"updated_at_time,omitempty"`
	Duration             *duration.Duration              `protobuf:"bytes,18,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a2e57e9bcc69b6, []int{0}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Build) GetRepoSource() string {
	if m != nil {
		return m.RepoSource
	}
	return ""
}

func (m *Build) GetRepoOwner() string {
	if m != nil {
		return m.RepoOwner
	}
	return ""
}

func (m *Build) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *Build) GetRepoBranch() string {
	if m != nil {
		return m.RepoBranch
	}
	return ""
}

func (m *Build) GetRepoRevision() string {
	if m != nil {
		return m.RepoRevision
	}
	return ""
}

func (m *Build) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

func (m *Build) GetBuildStatus() string {
	if m != nil {
		return m.BuildStatus
	}
	return ""
}

func (m *Build) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Build) GetReleaseTargets() []*ReleaseTarget {
	if m != nil {
		return m.ReleaseTargets
	}
	return nil
}

func (m *Build) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *Build) GetManifestWithDefaults() string {
	if m != nil {
		return m.ManifestWithDefaults
	}
	return ""
}

func (m *Build) GetCommits() []*GitCommit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *Build) GetTriggers() []*manifest_v1.EstafetteTrigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func (m *Build) GetEvents() []*manifest_v1.EstafetteEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Build) GetInsertedAtTime() *timestamp.Timestamp {
	if m != nil {
		return m.InsertedAtTime
	}
	return nil
}

func (m *Build) GetUpdatedAtTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAtTime
	}
	return nil
}

func (m *Build) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*Build)(nil), "estafette.ci.contracts.v1.Build")
}

func init() {
	proto.RegisterFile("estafette/ci/contracts/v1/build.proto", fileDescriptor_a1a2e57e9bcc69b6)
}

var fileDescriptor_a1a2e57e9bcc69b6 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x3e,
	0x14, 0x57, 0xbb, 0xff, 0xba, 0xce, 0xdd, 0xba, 0xfd, 0x2d, 0x84, 0xbc, 0xa2, 0xb1, 0xf2, 0x25,
	0x55, 0x93, 0xe6, 0xb0, 0x01, 0x12, 0x17, 0x08, 0x41, 0xb7, 0xc1, 0x0d, 0xe2, 0x23, 0x9b, 0x86,
	0xc4, 0x8d, 0xe5, 0x24, 0x6e, 0x6a, 0x29, 0x89, 0x2b, 0xdb, 0xe9, 0x2e, 0x78, 0x23, 0x9e, 0x8d,
	0x87, 0x40, 0x3e, 0x89, 0xb3, 0x0d, 0xd6, 0xc2, 0x5d, 0xcf, 0xef, 0xeb, 0x44, 0xc7, 0xe7, 0x14,
	0x3d, 0x11, 0xc6, 0xf2, 0x89, 0xb0, 0x56, 0x04, 0xb1, 0x0c, 0x62, 0x55, 0x58, 0xcd, 0x63, 0x6b,
	0x82, 0xf9, 0x61, 0x10, 0x95, 0x32, 0x4b, 0xe8, 0x4c, 0x2b, 0xab, 0xf0, 0x4e, 0x23, 0xa3, 0xb1,
	0xa4, 0x8d, 0x8c, 0xce, 0x0f, 0x07, 0xfb, 0x8b, 0x13, 0x52, 0x69, 0x59, 0xac, 0xf2, 0x5c, 0xda,
	0x2a, 0x66, 0xb0, 0xa4, 0x5b, 0xc6, 0x23, 0x91, 0xd5, 0x32, 0xba, 0x58, 0xa6, 0x45, 0x26, 0xb8,
	0x11, 0xcc, 0x72, 0x9d, 0x0a, 0x1f, 0x7b, 0x3f, 0x55, 0x2a, 0xcd, 0x44, 0x00, 0x55, 0x54, 0x4e,
	0x82, 0xa4, 0xd4, 0xdc, 0x4a, 0x55, 0xd4, 0xfc, 0xde, 0xef, 0xbc, 0x95, 0xb9, 0x6b, 0x91, 0xcf,
	0x6e, 0x6d, 0x98, 0xf3, 0x42, 0x4e, 0x84, 0xb1, 0xae, 0x5f, 0x43, 0x30, 0x31, 0x17, 0x85, 0x6f,
	0xf8, 0xf4, 0x1f, 0xf4, 0x56, 0xcb, 0x34, 0x15, 0xba, 0x72, 0x3c, 0xfc, 0xd9, 0x41, 0xab, 0x63,
	0x37, 0x50, 0xdc, 0x47, 0x6d, 0x99, 0x90, 0xd6, 0xb0, 0x35, 0x5a, 0x0f, 0xdb, 0x32, 0xc1, 0x7b,
	0xa8, 0xa7, 0xc5, 0x4c, 0x31, 0xa3, 0x4a, 0x1d, 0x0b, 0xd2, 0x06, 0x02, 0x39, 0xe8, 0x0c, 0x10,
	0xbc, 0x8b, 0xa0, 0x62, 0xea, 0xb2, 0x10, 0x9a, 0xac, 0x00, 0xbf, 0xee, 0x90, 0x4f, 0x0e, 0xc0,
	0xf7, 0x10, 0x14, 0xac, 0xe0, 0xb9, 0x20, 0xff, 0x01, 0xdb, 0x75, 0xc0, 0x47, 0x9e, 0x8b, 0x26,
	0x3c, 0xd2, 0xbc, 0x88, 0xa7, 0x64, 0xf5, 0x2a, 0x7c, 0x0c, 0x08, 0x7e, 0x84, 0x36, 0x41, 0xa0,
	0xc5, 0x5c, 0x1a, 0xa9, 0x0a, 0xd2, 0x01, 0xc9, 0x86, 0x03, 0xc3, 0x1a, 0x73, 0x22, 0x58, 0x06,
	0x36, 0x17, 0x1a, 0x44, 0x6b, 0x95, 0x08, 0xc0, 0x8b, 0x0a, 0xc3, 0x0f, 0x50, 0x55, 0x33, 0x63,
	0xb9, 0x2d, 0x0d, 0xe9, 0x82, 0xa6, 0x07, 0xd8, 0x19, 0x40, 0xf8, 0x25, 0xea, 0xc0, 0x33, 0x1b,
	0xb2, 0x3e, 0x5c, 0x19, 0xf5, 0x8e, 0x86, 0x74, 0xe1, 0x5a, 0xd1, 0x0f, 0x4e, 0x18, 0xd6, 0x7a,
	0xfc, 0x05, 0x6d, 0xdd, 0x7c, 0x79, 0x43, 0x10, 0x44, 0x8c, 0x96, 0x44, 0x84, 0x95, 0xe3, 0x1c,
	0x0c, 0x61, 0x5f, 0x5f, 0x2f, 0x0d, 0x1e, 0xa0, 0xae, 0x7f, 0x38, 0xd2, 0xab, 0xc6, 0xe6, 0x6b,
	0xfc, 0x1c, 0xdd, 0xf5, 0xbf, 0xd9, 0xa5, 0xb4, 0x53, 0x96, 0x88, 0x09, 0x2f, 0x33, 0x6b, 0xc8,
	0x06, 0x28, 0xef, 0x78, 0xf6, 0xab, 0xb4, 0xd3, 0x93, 0x9a, 0xc3, 0xaf, 0xd1, 0x5a, 0xb5, 0xed,
	0x86, 0x6c, 0xc2, 0xc7, 0x3d, 0x5e, 0xf2, 0x71, 0xef, 0xa5, 0x3d, 0x06, 0x71, 0xe8, 0x4d, 0xf8,
	0x1d, 0xea, 0xd6, 0x4b, 0x63, 0x48, 0x1f, 0x02, 0xf6, 0x6f, 0x06, 0xf8, 0xae, 0xce, 0x7f, 0xea,
	0x89, 0xf3, 0xca, 0x12, 0x36, 0x5e, 0xfc, 0x06, 0x75, 0x60, 0x59, 0x0d, 0xd9, 0xba, 0x6d, 0x46,
	0xb7, 0xa6, 0x9c, 0x3a, 0x43, 0x58, 0xfb, 0xf0, 0x09, 0xda, 0x96, 0x85, 0x11, 0xda, 0x8a, 0x84,
	0x71, 0xcb, 0xdc, 0xb9, 0x90, 0xed, 0x61, 0x6b, 0xd4, 0x3b, 0x1a, 0xd0, 0xea, 0x96, 0xa8, 0xbf,
	0x25, 0x7a, 0xee, 0x6f, 0x29, 0xec, 0x7b, 0xcf, 0x5b, 0xeb, 0x40, 0x3c, 0x46, 0x5b, 0xe5, 0x2c,
	0xe1, 0xd7, 0x43, 0xfe, 0xff, 0x6b, 0xc8, 0x66, 0x6d, 0xa9, 0x33, 0x5e, 0xa0, 0xae, 0x3f, 0x66,
	0x82, 0xc1, 0xbc, 0xf3, 0x87, 0xf9, 0xa4, 0x16, 0x84, 0x8d, 0x74, 0xfc, 0x1d, 0xed, 0xc6, 0x2a,
	0x5f, 0x3c, 0xfe, 0xcf, 0xad, 0x6f, 0xaf, 0x52, 0x69, 0xa7, 0x65, 0x44, 0x63, 0x95, 0x5f, 0x5d,
	0xed, 0xd5, 0xaf, 0x83, 0x58, 0x1e, 0x40, 0x07, 0x73, 0x90, 0xaa, 0x8c, 0x17, 0xe9, 0x8d, 0xff,
	0xa1, 0x1f, 0xed, 0x9d, 0x66, 0x74, 0xf4, 0x58, 0xd2, 0xe3, 0x26, 0xfb, 0xe2, 0x30, 0xea, 0x80,
	0xef, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xf6, 0x07, 0x47, 0x5c, 0x05, 0x00, 0x00,
}
