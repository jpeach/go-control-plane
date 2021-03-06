// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/regex.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType           isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	Regex                string                    `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

type RegexMatcher_GoogleRE2 struct {
	MaxProgramSize       *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

type RegexMatchAndSubstitute struct {
	Pattern              *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Substitution         string        `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RegexMatchAndSubstitute) Reset()         { *m = RegexMatchAndSubstitute{} }
func (m *RegexMatchAndSubstitute) String() string { return proto.CompactTextString(m) }
func (*RegexMatchAndSubstitute) ProtoMessage()    {}
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{1}
}

func (m *RegexMatchAndSubstitute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatchAndSubstitute.Unmarshal(m, b)
}
func (m *RegexMatchAndSubstitute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatchAndSubstitute.Marshal(b, m, deterministic)
}
func (m *RegexMatchAndSubstitute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatchAndSubstitute.Merge(m, src)
}
func (m *RegexMatchAndSubstitute) XXX_Size() int {
	return xxx_messageInfo_RegexMatchAndSubstitute.Size(m)
}
func (m *RegexMatchAndSubstitute) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatchAndSubstitute.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatchAndSubstitute proto.InternalMessageInfo

func (m *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *RegexMatchAndSubstitute) GetSubstitution() string {
	if m != nil {
		return m.Substitution
	}
	return ""
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.v3.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.v3.RegexMatcher.GoogleRE2")
	proto.RegisterType((*RegexMatchAndSubstitute)(nil), "envoy.type.matcher.v3.RegexMatchAndSubstitute")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/regex.proto", fileDescriptor_7de4f215ebc85482) }

var fileDescriptor_7de4f215ebc85482 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0xcd, 0xe8, 0xee, 0xda, 0x77, 0x17, 0x29, 0x03, 0xb2, 0xa5, 0xb8, 0x6b, 0xdb, 0x05,
	0x59, 0xff, 0x25, 0x3a, 0xbd, 0x55, 0x3c, 0x18, 0x10, 0xff, 0x80, 0x50, 0xb2, 0x28, 0xde, 0x4a,
	0x6a, 0x5f, 0xc7, 0x40, 0x9b, 0x84, 0x4c, 0x66, 0x6c, 0xf7, 0xe4, 0x51, 0xfd, 0x08, 0x7e, 0x09,
	0xef, 0xde, 0x05, 0x6f, 0xe2, 0xb7, 0x91, 0x3d, 0xc9, 0x4c, 0x3a, 0xea, 0xe2, 0xa0, 0xde, 0xc2,
	0xfb, 0x3e, 0xcf, 0xc3, 0xef, 0x7d, 0x02, 0x7d, 0xd4, 0x85, 0x59, 0x31, 0xbf, 0xb2, 0xc8, 0x16,
	0xd2, 0xbf, 0x78, 0x85, 0x8e, 0x15, 0x43, 0xe6, 0x30, 0xc5, 0x25, 0xb5, 0xce, 0x78, 0x13, 0x5f,
	0xac, 0x24, 0xb4, 0x94, 0xd0, 0xb5, 0x84, 0x16, 0xc3, 0xee, 0x7e, 0x6a, 0x4c, 0x3a, 0x47, 0x56,
	0x89, 0xa6, 0xf9, 0x4b, 0xf6, 0xda, 0x49, 0x6b, 0xd1, 0x65, 0xc1, 0xd6, 0xdd, 0xcb, 0x67, 0x56,
	0x32, 0xa9, 0xb5, 0xf1, 0xd2, 0x2b, 0xa3, 0x33, 0x96, 0x79, 0xe9, 0xf3, 0x7a, 0xdd, 0xff, 0x63,
	0x5d, 0xa0, 0xcb, 0x94, 0xd1, 0x4a, 0xa7, 0x6b, 0xc9, 0x6e, 0x21, 0xe7, 0x6a, 0x26, 0x3d, 0xb2,
	0xfa, 0x11, 0x16, 0x83, 0xaf, 0x11, 0xec, 0x88, 0x92, 0xf0, 0x49, 0xc0, 0x89, 0x9f, 0x03, 0x04,
	0x9a, 0x89, 0xc3, 0xa4, 0x43, 0x7a, 0xe4, 0x70, 0x3b, 0xb9, 0x49, 0x1b, 0xb9, 0xe9, 0xef, 0x46,
	0xfa, 0xa0, 0x72, 0x89, 0xfb, 0x09, 0x3f, 0x7f, 0xc2, 0x37, 0xde, 0x93, 0xa8, 0x4d, 0x1e, 0x9e,
	0x11, 0xad, 0x10, 0x26, 0x30, 0x89, 0xf7, 0x60, 0xa3, 0xea, 0xa2, 0x13, 0xf5, 0xc8, 0x61, 0x8b,
	0x6f, 0x9d, 0xf0, 0x73, 0x2e, 0xea, 0x11, 0x11, 0xa6, 0xdd, 0x77, 0x04, 0x5a, 0x3f, 0x33, 0xe2,
	0xc7, 0xd0, 0x5e, 0xc8, 0xe5, 0xc4, 0x3a, 0x93, 0x3a, 0xb9, 0x98, 0x64, 0xea, 0x18, 0xd7, 0x30,
	0x97, 0x68, 0x88, 0xa4, 0x75, 0x5b, 0xf4, 0xe9, 0x23, 0xed, 0x87, 0xc9, 0x33, 0x39, 0xcf, 0x91,
	0x47, 0x1d, 0x22, 0x2e, 0x2c, 0xe4, 0x72, 0x1c, 0x8c, 0x47, 0xea, 0x18, 0x47, 0xb7, 0x3e, 0x7c,
	0x7e, 0xbb, 0x7f, 0x1d, 0xae, 0x36, 0x1c, 0xd1, 0x7c, 0xc1, 0xe8, 0x4a, 0xe9, 0xe8, 0xc3, 0xe5,
	0x7f, 0x38, 0x78, 0x0c, 0xdb, 0xa8, 0x53, 0xa5, 0x71, 0x52, 0x6a, 0xe2, 0xb3, 0xdf, 0x39, 0x19,
	0x7c, 0x24, 0xb0, 0xfb, 0x4b, 0x74, 0x4f, 0xcf, 0x8e, 0xf2, 0x69, 0xe6, 0x95, 0xcf, 0x3d, 0xc6,
	0x77, 0x61, 0xcb, 0x4a, 0xef, 0xd1, 0xe9, 0xf5, 0x31, 0x07, 0xff, 0xd1, 0xac, 0xa8, 0x3d, 0xf1,
	0x00, 0x76, 0xb2, 0x3a, 0x4c, 0x19, 0x1d, 0x8a, 0x14, 0xa7, 0x66, 0xa3, 0xdb, 0x25, 0xfa, 0x0d,
	0xb8, 0xf6, 0x57, 0xf4, 0x53, 0x54, 0xfc, 0xce, 0xa7, 0x37, 0x5f, 0xbe, 0x6d, 0x46, 0xed, 0x08,
	0x0e, 0x94, 0x09, 0x40, 0xd6, 0x99, 0xe5, 0xaa, 0x99, 0x8d, 0x43, 0x95, 0x33, 0x2e, 0xdb, 0x1f,
	0x93, 0xe9, 0x66, 0xf5, 0x0d, 0xc3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x64, 0xee, 0x2b,
	0xfe, 0x02, 0x00, 0x00,
}
