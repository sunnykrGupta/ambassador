// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tap/v3/tap.proto

package envoy_extensions_transport_sockets_tap_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/extensions/common/tap/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Tap struct {
	CommonConfig         *v3.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	TransportSocket      *v31.TransportSocket      `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e115889d24f02d1, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v3.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *v31.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.extensions.transport_sockets.tap.v3.Tap")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/tap/v3/tap.proto", fileDescriptor_5e115889d24f02d1)
}

var fileDescriptor_5e115889d24f02d1 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x3d, 0x4b, 0xc3, 0x40,
	0x18, 0x26, 0x11, 0x4b, 0x89, 0x8a, 0xa5, 0x8b, 0xa5, 0xe0, 0x27, 0x08, 0x8a, 0x78, 0x07, 0x8d,
	0x22, 0x38, 0xa6, 0x38, 0x0a, 0x45, 0x33, 0xb9, 0x94, 0xb7, 0xe9, 0x59, 0x83, 0xed, 0xbd, 0x47,
	0xee, 0x1a, 0xda, 0xcd, 0x51, 0xfc, 0x09, 0xfe, 0x14, 0x77, 0xc1, 0xd5, 0xbf, 0x23, 0x08, 0x72,
	0xf7, 0x26, 0x68, 0x9b, 0xc5, 0x29, 0xb9, 0x3c, 0x9f, 0xf7, 0x24, 0x08, 0x85, 0xcc, 0x71, 0xce,
	0xc5, 0xcc, 0x08, 0xa9, 0x53, 0x94, 0x9a, 0x9b, 0x0c, 0xa4, 0x56, 0x98, 0x99, 0xbe, 0xc6, 0xe4,
	0x51, 0x18, 0xcd, 0x0d, 0x28, 0x9e, 0x87, 0xf6, 0xc1, 0x54, 0x86, 0x06, 0x9b, 0xc7, 0x4e, 0xc4,
	0x7e, 0x45, 0xac, 0x22, 0x62, 0x96, 0x9d, 0x87, 0xed, 0x5d, 0xf2, 0x4f, 0x50, 0xde, 0xa7, 0x23,
	0x9e, 0x60, 0x26, 0xac, 0xd5, 0x00, 0xb4, 0x20, 0xaf, 0xf6, 0x49, 0xa5, 0x40, 0x82, 0x93, 0x09,
	0xca, 0x32, 0x95, 0x4e, 0x05, 0x79, 0x7b, 0x3a, 0x54, 0xc0, 0x41, 0x4a, 0x34, 0x60, 0x1c, 0x59,
	0x1b, 0x30, 0x53, 0x5d, 0xc0, 0xfb, 0x15, 0x38, 0x17, 0x99, 0x35, 0x4d, 0xe5, 0xa8, 0xa0, 0x6c,
	0xe5, 0x30, 0x4e, 0x87, 0x60, 0x04, 0x2f, 0x5f, 0x08, 0x38, 0xf8, 0xf6, 0x82, 0x95, 0x18, 0x54,
	0x73, 0x18, 0x6c, 0x50, 0x64, 0x9f, 0x3a, 0xb7, 0xbc, 0x3d, 0xef, 0x68, 0xad, 0x73, 0xce, 0x2a,
	0x77, 0x2e, 0x9a, 0xd1, 0x45, 0x59, 0xd7, 0x9d, 0xae, 0x4a, 0xb8, 0xeb, 0xc4, 0x51, 0xfd, 0x2b,
	0x5a, 0x7d, 0xf1, 0xfc, 0x86, 0x77, 0xb3, 0x4e, 0x74, 0xfa, 0xde, 0xbc, 0x0b, 0x1a, 0xcb, 0x93,
	0xb5, 0x7c, 0x17, 0x74, 0x58, 0x04, 0x51, 0x3a, 0xb3, 0x8b, 0x59, 0xfb, 0xb8, 0x64, 0xdf, 0x3a,
	0xf2, 0x1f, 0xe3, 0x4d, 0xb3, 0x08, 0x5d, 0x9e, 0xbd, 0xbe, 0x3f, 0xef, 0xf0, 0xe0, 0x74, 0xc1,
	0x67, 0x39, 0x8d, 0x6a, 0x77, 0x60, 0xac, 0x1e, 0x80, 0xc5, 0xa0, 0xa2, 0xeb, 0xb7, 0xa7, 0x8f,
	0xcf, 0x9a, 0xdf, 0xf0, 0x83, 0x8b, 0x14, 0xa9, 0x83, 0xca, 0x70, 0x36, 0x67, 0xff, 0xfe, 0xd7,
	0x51, 0x3d, 0x06, 0xd5, 0xb3, 0x63, 0xf6, 0xbc, 0x41, 0xcd, 0xad, 0x1a, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x56, 0xf9, 0x21, 0x22, 0x60, 0x02, 0x00, 0x00,
}