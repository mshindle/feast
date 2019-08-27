// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/FeatureSet.proto

package core // import "github.com/gojek/feast/protos/generated/go/feast/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojek/feast/protos/generated/go/feast/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureSetSpec struct {
	// Name of the featureSet. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// FeatureSet version.
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Features contained within this featureSet.
	Features []*FeatureSpec `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`
	// Source on which feature rows can be found
	Source               *Source  `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSetSpec) Reset()         { *m = FeatureSetSpec{} }
func (m *FeatureSetSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSetSpec) ProtoMessage()    {}
func (*FeatureSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSet_509e8d939816a703, []int{0}
}
func (m *FeatureSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSetSpec.Unmarshal(m, b)
}
func (m *FeatureSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSetSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSetSpec.Merge(dst, src)
}
func (m *FeatureSetSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSetSpec.Size(m)
}
func (m *FeatureSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSetSpec proto.InternalMessageInfo

func (m *FeatureSetSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSetSpec) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FeatureSetSpec) GetFeatures() []*FeatureSpec {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureSetSpec) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type FeatureSpec struct {
	// Name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature.
	ValueType types.ValueType_Enum `protobuf:"varint,2,opt,name=valueType,proto3,enum=feast.types.ValueType_Enum" json:"valueType,omitempty"`
	// Describes whether a feature is also a key.
	// This allows the feature to be used during joins between feature sets.
	// If the featureSet is ingested into a store that supports keys, this value
	// will be made a key.
	IsKey                bool     `protobuf:"varint,3,opt,name=isKey,proto3" json:"isKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSpec) Reset()         { *m = FeatureSpec{} }
func (m *FeatureSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSpec) ProtoMessage()    {}
func (*FeatureSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSet_509e8d939816a703, []int{1}
}
func (m *FeatureSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSpec.Unmarshal(m, b)
}
func (m *FeatureSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSpec.Merge(dst, src)
}
func (m *FeatureSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSpec.Size(m)
}
func (m *FeatureSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSpec proto.InternalMessageInfo

func (m *FeatureSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_UNKNOWN
}

func (m *FeatureSpec) GetIsKey() bool {
	if m != nil {
		return m.IsKey
	}
	return false
}

func init() {
	proto.RegisterType((*FeatureSetSpec)(nil), "feast.core.FeatureSetSpec")
	proto.RegisterType((*FeatureSpec)(nil), "feast.core.FeatureSpec")
}

func init() {
	proto.RegisterFile("feast/core/FeatureSet.proto", fileDescriptor_FeatureSet_509e8d939816a703)
}

var fileDescriptor_FeatureSet_509e8d939816a703 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x89, 0xfb, 0xe3, 0xf6, 0x0e, 0x26, 0x04, 0x61, 0xc1, 0x5d, 0xca, 0x4e, 0xc5, 0x43,
	0x02, 0x1b, 0x1e, 0xbc, 0x0e, 0xf4, 0xe2, 0x45, 0xd2, 0xb1, 0x83, 0xb7, 0xae, 0xbe, 0xab, 0x55,
	0xdb, 0x94, 0x24, 0x1d, 0xf4, 0xd3, 0xf8, 0x55, 0xa5, 0x49, 0x6b, 0x7b, 0xf0, 0xd6, 0xf0, 0x7b,
	0xde, 0x87, 0xa7, 0x3f, 0x58, 0x9f, 0x31, 0x36, 0x56, 0x24, 0x4a, 0xa3, 0x78, 0xc6, 0xd8, 0x56,
	0x1a, 0x23, 0xb4, 0xbc, 0xd4, 0xca, 0x2a, 0x0a, 0x0e, 0xf2, 0x06, 0xde, 0xad, 0x7c, 0xd0, 0xd6,
	0x25, 0x1a, 0x71, 0x8c, 0xbf, 0x2b, 0xf4, 0xa1, 0x0e, 0xb8, 0x86, 0x48, 0x55, 0x3a, 0x69, 0xc1,
	0xe6, 0x87, 0xc0, 0xb2, 0xaf, 0x8c, 0x4a, 0x4c, 0x28, 0x85, 0x71, 0x11, 0xe7, 0xc8, 0x48, 0x40,
	0xc2, 0xb9, 0x74, 0xdf, 0x94, 0xc1, 0xf5, 0x05, 0xb5, 0xc9, 0x54, 0xc1, 0xae, 0x02, 0x12, 0x4e,
	0x64, 0xf7, 0xa4, 0x3b, 0x98, 0x9d, 0xfd, 0xbd, 0x61, 0xa3, 0x60, 0x14, 0x2e, 0xb6, 0x2b, 0xde,
	0x2f, 0xe2, 0x5d, 0x77, 0x89, 0x89, 0xfc, 0x0b, 0xd2, 0x7b, 0x98, 0x1a, 0xb7, 0x82, 0x8d, 0x03,
	0x12, 0x2e, 0xb6, 0x74, 0x78, 0xe2, 0xf7, 0xc9, 0x36, 0xb1, 0xd1, 0xb0, 0x18, 0x94, 0xfc, 0xbb,
	0xee, 0x11, 0xe6, 0x97, 0xe6, 0x67, 0x0f, 0x75, 0x89, 0x6e, 0xdf, 0x72, 0xbb, 0x6e, 0x1b, 0x9d,
	0x0a, 0x7e, 0xec, 0x28, 0x7f, 0x2a, 0xaa, 0x5c, 0xf6, 0x69, 0x7a, 0x0b, 0x93, 0xcc, 0xbc, 0x60,
	0xcd, 0x46, 0x01, 0x09, 0x67, 0xd2, 0x3f, 0xf6, 0x07, 0x18, 0x58, 0xdd, 0xdf, 0xf4, 0x82, 0x5e,
	0x1b, 0x69, 0x6f, 0x0f, 0x69, 0x66, 0x3f, 0xaa, 0x13, 0x4f, 0x54, 0x2e, 0x52, 0xf5, 0x89, 0x5f,
	0xc2, 0xeb, 0x75, 0x4a, 0x8d, 0x48, 0xb1, 0x40, 0x1d, 0x5b, 0x7c, 0x17, 0xa9, 0x12, 0xbd, 0xf8,
	0xd3, 0xd4, 0xf1, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xb9, 0xe1, 0x67, 0xcf, 0x01,
	0x00, 0x00,
}