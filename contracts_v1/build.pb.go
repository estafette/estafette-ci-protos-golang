// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/build.proto

package contracts_v1

import (
	fmt "fmt"
	manifest_v1 "github.com/estafette/estafette-ci-contracts-golang/manifest_v1"
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
	InsertedAt           *timestamp.Timestamp            `protobuf:"bytes,16,opt,name=inserted_at,json=insertedAt,proto3" json:"inserted_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp            `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Duration             *duration.Duration              `protobuf:"bytes,18,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_a232016f98093c2b, []int{0}
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

func (m *Build) GetInsertedAt() *timestamp.Timestamp {
	if m != nil {
		return m.InsertedAt
	}
	return nil
}

func (m *Build) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
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
	proto.RegisterType((*Build)(nil), "contracts.v1.Build")
}

func init() { proto.RegisterFile("contracts.v1/build.proto", fileDescriptor_a232016f98093c2b) }

var fileDescriptor_a232016f98093c2b = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0xb5, 0xee, 0xbf, 0xae, 0x75, 0xf7, 0xf2, 0xc7, 0x20, 0xf0, 0x3a, 0xc6, 0x36, 0x76,
	0x33, 0x09, 0xcd, 0x55, 0x37, 0xb8, 0x98, 0xb8, 0x40, 0x7b, 0x13, 0x42, 0x42, 0x20, 0x65, 0xd3,
	0x90, 0xb8, 0xb1, 0x9c, 0xc4, 0x4d, 0x2d, 0x25, 0x71, 0x65, 0x9f, 0x74, 0x17, 0x7c, 0x12, 0xbe,
	0x02, 0x9f, 0x12, 0xe5, 0x24, 0x4e, 0x1b, 0x98, 0xc4, 0x5d, 0xce, 0xf3, 0xfc, 0x1e, 0xdb, 0xf1,
	0x39, 0x26, 0x2c, 0x32, 0x39, 0x58, 0x19, 0x81, 0xe3, 0xf3, 0xf1, 0x28, 0x2c, 0x74, 0x1a, 0xf3,
	0x99, 0x35, 0x60, 0xe8, 0xc6, 0xb2, 0x33, 0xdc, 0x6b, 0x71, 0x89, 0x06, 0x11, 0x99, 0x2c, 0xd3,
	0x50, 0xc1, 0xc3, 0xf6, 0x32, 0xa9, 0x0c, 0x55, 0x5a, 0x3b, 0x87, 0x2d, 0xc7, 0xaa, 0x54, 0x49,
	0xa7, 0x04, 0x48, 0x9b, 0x28, 0x1f, 0x7e, 0x95, 0x18, 0x93, 0xa4, 0x6a, 0x84, 0x55, 0x58, 0x4c,
	0x46, 0x71, 0x61, 0x25, 0x68, 0x93, 0xd7, 0xfe, 0xfe, 0x9f, 0x3e, 0xe8, 0x4c, 0x39, 0x90, 0xd9,
	0xcc, 0xef, 0x91, 0xc9, 0x5c, 0x4f, 0x94, 0x83, 0x72, 0x8b, 0xd2, 0x99, 0x28, 0x00, 0x25, 0xd4,
	0x5c, 0xe5, 0x7e, 0x8f, 0xa3, 0xc7, 0x11, 0xb0, 0x3a, 0x49, 0x94, 0xad, 0xa0, 0xd7, 0x3f, 0xbb,
	0x64, 0xed, 0xb2, 0xbc, 0x02, 0xba, 0x45, 0x3a, 0x3a, 0x66, 0x2b, 0x07, 0x2b, 0xc7, 0xfd, 0xa0,
	0xa3, 0x63, 0xba, 0x4f, 0x06, 0x56, 0xcd, 0x8c, 0x70, 0xa6, 0xb0, 0x91, 0x62, 0x1d, 0x34, 0x48,
	0x29, 0xdd, 0xa2, 0x42, 0xf7, 0x08, 0x56, 0xc2, 0x3c, 0xe4, 0xca, 0xb2, 0x55, 0xf4, 0xfb, 0xa5,
	0xf2, 0xb5, 0x14, 0xe8, 0x2e, 0xc1, 0x42, 0xe4, 0x32, 0x53, 0xec, 0x3f, 0x74, 0x7b, 0xa5, 0xf0,
	0x45, 0x66, 0xaa, 0x59, 0x3c, 0xb4, 0x32, 0x8f, 0xa6, 0x6c, 0x6d, 0xb1, 0xf8, 0x25, 0x2a, 0xf4,
	0x88, 0x6c, 0x22, 0x60, 0xd5, 0x5c, 0x3b, 0x6d, 0x72, 0xd6, 0x45, 0x64, 0xa3, 0x14, 0x83, 0x5a,
	0x2b, 0x21, 0x6c, 0x9f, 0x98, 0x2b, 0x8b, 0xd0, 0x7a, 0x05, 0xa1, 0x78, 0x5f, 0x69, 0xf4, 0x90,
	0x54, 0xb5, 0x70, 0x20, 0xa1, 0x70, 0xac, 0x87, 0xcc, 0x00, 0xb5, 0x5b, 0x94, 0xe8, 0x1b, 0xd2,
	0xc5, 0xfe, 0x39, 0xd6, 0x3f, 0x58, 0x3d, 0x1e, 0x9c, 0x3e, 0xe5, 0xcb, 0x1d, 0xe4, 0x9f, 0x4b,
	0x2f, 0xa8, 0x11, 0x7a, 0x4d, 0xb6, 0xdb, 0x2d, 0x75, 0x8c, 0x60, 0x6a, 0xb7, 0x9d, 0x0a, 0x2a,
	0xe8, 0x0e, 0x99, 0x60, 0xcb, 0x2e, 0x97, 0x8e, 0x0e, 0x49, 0xcf, 0xb7, 0x87, 0x0d, 0xaa, 0xcb,
	0xf1, 0x35, 0x7d, 0x4b, 0x9e, 0xfb, 0x6f, 0xf1, 0xa0, 0x61, 0x2a, 0x62, 0x35, 0x91, 0x45, 0x0a,
	0x8e, 0x6d, 0x20, 0xf9, 0xcc, 0xbb, 0xdf, 0x34, 0x4c, 0xaf, 0x6b, 0x8f, 0x8e, 0xc9, 0x7a, 0x35,
	0x9f, 0x8e, 0x6d, 0xe2, 0x79, 0x5e, 0xb4, 0xcf, 0xf3, 0x51, 0xc3, 0x15, 0xfa, 0x81, 0xe7, 0xe8,
	0x39, 0xe9, 0xd5, 0xd3, 0xe0, 0xd8, 0x16, 0x66, 0xf6, 0xf8, 0xd2, 0xd0, 0xf0, 0x1b, 0x3f, 0x34,
	0x77, 0x15, 0x15, 0x34, 0x38, 0x3d, 0x23, 0x5d, 0x9c, 0x35, 0xc7, 0xb6, 0xeb, 0x9f, 0x7f, 0x34,
	0x78, 0x53, 0x32, 0x41, 0x8d, 0xd2, 0xf7, 0x64, 0xa0, 0x73, 0xa7, 0x2c, 0xa8, 0x58, 0x48, 0x60,
	0xff, 0x1f, 0xac, 0x1c, 0x0f, 0x4e, 0x87, 0xbc, 0x9a, 0x75, 0xee, 0x67, 0x9d, 0xdf, 0xf9, 0x59,
	0x0f, 0x88, 0xc7, 0x2f, 0x80, 0x9e, 0x13, 0x52, 0xcc, 0x62, 0x59, 0x67, 0x9f, 0xfc, 0x33, 0xdb,
	0xaf, 0xe9, 0x0b, 0xa0, 0xef, 0x48, 0xcf, 0xbf, 0x2f, 0x46, 0x31, 0xb8, 0xf3, 0x57, 0xf0, 0xba,
	0x06, 0x82, 0x06, 0xbd, 0xfc, 0x41, 0x5e, 0x6a, 0xc3, 0x9b, 0x97, 0xc3, 0x23, 0xdd, 0xba, 0xd1,
	0xef, 0x1f, 0x12, 0x0d, 0xd3, 0x22, 0xe4, 0x91, 0xc9, 0x16, 0xef, 0x6b, 0xf1, 0x75, 0x12, 0xe9,
	0x93, 0x86, 0x3f, 0x49, 0x4c, 0x2a, 0xf3, 0x64, 0xd4, 0x08, 0x62, 0x3e, 0xfe, 0xd5, 0xd9, 0x69,
	0x2e, 0x8a, 0x5f, 0x7d, 0xe2, 0x57, 0xcd, 0xe2, 0xf7, 0xe3, 0xb0, 0x8b, 0x27, 0x3b, 0xfb, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x8d, 0x01, 0xf3, 0xae, 0x04, 0x00, 0x00,
}
