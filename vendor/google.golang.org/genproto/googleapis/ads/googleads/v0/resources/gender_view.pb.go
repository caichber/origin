// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/gender_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A gender view.
type GenderView struct {
	// The resource name of the gender view.
	// Gender view resource names have the form:
	//
	// `customers/{customer_id}/genderViews/{ad_group_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderView) Reset()         { *m = GenderView{} }
func (m *GenderView) String() string { return proto.CompactTextString(m) }
func (*GenderView) ProtoMessage()    {}
func (*GenderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_gender_view_5674b96e81d6b40e, []int{0}
}
func (m *GenderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderView.Unmarshal(m, b)
}
func (m *GenderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderView.Marshal(b, m, deterministic)
}
func (dst *GenderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderView.Merge(dst, src)
}
func (m *GenderView) XXX_Size() int {
	return xxx_messageInfo_GenderView.Size(m)
}
func (m *GenderView) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderView.DiscardUnknown(m)
}

var xxx_messageInfo_GenderView proto.InternalMessageInfo

func (m *GenderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GenderView)(nil), "google.ads.googleads.v0.resources.GenderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/gender_view.proto", fileDescriptor_gender_view_5674b96e81d6b40e)
}

var fileDescriptor_gender_view_5674b96e81d6b40e = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xa2, 0xd4,
	0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0x62, 0xfd, 0xf4, 0xd4, 0xbc, 0x94, 0xd4, 0xa2, 0xf8, 0xb2,
	0xcc, 0xd4, 0x72, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x45, 0x88, 0x4a, 0xbd, 0xc4, 0x94,
	0x62, 0x3d, 0xb8, 0x26, 0xbd, 0x32, 0x03, 0x3d, 0xb8, 0x26, 0x25, 0x43, 0x2e, 0x2e, 0x77, 0xb0,
	0xbe, 0xb0, 0xcc, 0xd4, 0x72, 0x21, 0x65, 0x2e, 0x5e, 0x98, 0x54, 0x7c, 0x5e, 0x62, 0x6e, 0xaa,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f, 0x4c, 0xd0, 0x2f, 0x31, 0x37, 0xd5, 0xe9, 0x0f,
	0x23, 0x97, 0x6a, 0x72, 0x7e, 0xae, 0x1e, 0x41, 0xc3, 0x9d, 0xf8, 0x11, 0x46, 0x07, 0x80, 0x1c,
	0x14, 0xc0, 0x18, 0xe5, 0x05, 0xd5, 0x95, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94,
	0x0e, 0x72, 0x37, 0xd8, 0xb9, 0x30, 0x7f, 0x15, 0x64, 0x16, 0xe3, 0xf1, 0xa6, 0x35, 0x9c, 0xb5,
	0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0xa2, 0x3b, 0xc4, 0x48, 0xc7, 0x94, 0x62, 0x3d,
	0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x0b, 0x82, 0xa9, 0x3c, 0x05, 0x53, 0x13, 0xe3, 0x98, 0x52,
	0x1c, 0x03, 0x57, 0x13, 0x13, 0x66, 0x10, 0x03, 0x57, 0xf3, 0x8a, 0x49, 0x15, 0x22, 0x61, 0x65,
	0xe5, 0x98, 0x52, 0x6c, 0x65, 0x05, 0x57, 0x65, 0x65, 0x15, 0x66, 0x60, 0x65, 0x05, 0x57, 0x97,
	0xc4, 0x06, 0x76, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x78, 0x4a, 0xab, 0x4a, 0x92, 0x01,
	0x00, 0x00,
}
