// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/button/button.proto

/*
Package button is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/mochi/pb/button/button.proto

It has these top-level messages:
	Button
*/
package button

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mochi_text "github.com/overcyn/mochi/pb/text"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Button struct {
	StyledText *mochi_text.StyledText `protobuf:"bytes,1,opt,name=styledText" json:"styledText,omitempty"`
	OnPress    int64                  `protobuf:"varint,2,opt,name=onPress" json:"onPress,omitempty"`
}

func (m *Button) Reset()                    { *m = Button{} }
func (m *Button) String() string            { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()               {}
func (*Button) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Button) GetStyledText() *mochi_text.StyledText {
	if m != nil {
		return m.StyledText
	}
	return nil
}

func (m *Button) GetOnPress() int64 {
	if m != nil {
		return m.OnPress
	}
	return 0
}

func init() {
	proto.RegisterType((*Button)(nil), "button.Button")
}

func init() { proto.RegisterFile("github.com/overcyn/mochi/pb/button/button.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x4b, 0x2d, 0x4a, 0xae, 0xcc, 0xd3, 0xcf,
	0xcd, 0x4f, 0xce, 0xc8, 0xd4, 0x2f, 0x48, 0xd2, 0x4f, 0x2a, 0x2d, 0x29, 0xc9, 0xcf, 0x83, 0x52,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94, 0x36, 0x3e, 0x8d, 0x25, 0xa9,
	0x15, 0x25, 0x60, 0x02, 0xa2, 0x49, 0x29, 0x8a, 0x8b, 0xcd, 0x09, 0xac, 0x4d, 0xc8, 0x8c, 0x8b,
	0xab, 0xb8, 0xa4, 0x32, 0x27, 0x35, 0x25, 0x24, 0xb5, 0xa2, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x4c, 0x0f, 0xac, 0x51, 0x0f, 0xac, 0x21, 0x18, 0x2e, 0x1b, 0x84, 0xa4, 0x52, 0x48,
	0x82, 0x8b, 0x3d, 0x3f, 0x2f, 0xa0, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39,
	0x08, 0xc6, 0x75, 0x92, 0x88, 0x82, 0x3a, 0x69, 0x11, 0x13, 0xaf, 0x2f, 0xc8, 0x9c, 0x00, 0x27,
	0x88, 0x5d, 0x49, 0x6c, 0x60, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x7f, 0xb6,
	0x9a, 0xe4, 0x00, 0x00, 0x00,
}
