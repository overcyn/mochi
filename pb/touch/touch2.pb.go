// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/touch/touch2.proto

/*
Package touch is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/mochi/pb/touch/touch2.proto

It has these top-level messages:
	ButtonRecognizer
	ButtonEvent
*/
package touch

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import mochi_touch "github.com/overcyn/mochi/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ButtonRecognizer struct {
	OnEvent int64 `protobuf:"varint,1,opt,name=onEvent" json:"onEvent,omitempty"`
}

func (m *ButtonRecognizer) Reset()                    { *m = ButtonRecognizer{} }
func (m *ButtonRecognizer) String() string            { return proto.CompactTextString(m) }
func (*ButtonRecognizer) ProtoMessage()               {}
func (*ButtonRecognizer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ButtonRecognizer) GetOnEvent() int64 {
	if m != nil {
		return m.OnEvent
	}
	return 0
}

type ButtonEvent struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Inside    bool                       `protobuf:"varint,3,opt,name=inside" json:"inside,omitempty"`
	Kind      mochi_touch.EventKind      `protobuf:"varint,4,opt,name=kind,enum=mochi.touch.EventKind" json:"kind,omitempty"`
}

func (m *ButtonEvent) Reset()                    { *m = ButtonEvent{} }
func (m *ButtonEvent) String() string            { return proto.CompactTextString(m) }
func (*ButtonEvent) ProtoMessage()               {}
func (*ButtonEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ButtonEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ButtonEvent) GetInside() bool {
	if m != nil {
		return m.Inside
	}
	return false
}

func (m *ButtonEvent) GetKind() mochi_touch.EventKind {
	if m != nil {
		return m.Kind
	}
	return mochi_touch.EventKind_EVENT_KIND_BEGAN
}

func init() {
	proto.RegisterType((*ButtonRecognizer)(nil), "touch.ButtonRecognizer")
	proto.RegisterType((*ButtonEvent)(nil), "touch.ButtonEvent")
}

func init() { proto.RegisterFile("github.com/overcyn/mochi/pb/touch/touch2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x89, 0x77, 0x9e, 0x9a, 0x13, 0x91, 0x0c, 0xa5, 0x74, 0xb1, 0xdc, 0x62, 0x11, 0x49,
	0xa0, 0x2e, 0xce, 0x05, 0x27, 0x11, 0x24, 0xdc, 0xe4, 0x66, 0xd3, 0xd8, 0x06, 0x6d, 0xbe, 0xd2,
	0xfb, 0x72, 0xa0, 0x7f, 0xc1, 0x7f, 0xe1, 0x2f, 0x95, 0x7e, 0xb1, 0x3a, 0xba, 0x04, 0x9e, 0xf0,
	0x90, 0xf7, 0x21, 0x5c, 0xb6, 0x0e, 0xbb, 0x50, 0x4b, 0x03, 0xbd, 0x82, 0xbd, 0x1d, 0xcd, 0xbb,
	0x57, 0x3d, 0x98, 0xce, 0xa9, 0xa1, 0x56, 0x08, 0xc1, 0x74, 0xf1, 0x2c, 0xe5, 0x30, 0x02, 0x82,
	0x38, 0x24, 0xca, 0x2e, 0x5a, 0x80, 0xf6, 0xcd, 0x2a, 0xba, 0xac, 0xc3, 0x8b, 0x42, 0xd7, 0xdb,
	0x1d, 0x3e, 0xf7, 0x43, 0xf4, 0xb2, 0xcb, 0x7f, 0xdf, 0x8d, 0xe2, 0xe6, 0x9a, 0x9f, 0x57, 0x01,
	0x11, 0xbc, 0xb6, 0x06, 0x5a, 0xef, 0x3e, 0xec, 0x28, 0x52, 0x7e, 0x04, 0xfe, 0x6e, 0x6f, 0x3d,
	0xa6, 0x2c, 0x67, 0xc5, 0x42, 0xcf, 0xb8, 0xf9, 0x64, 0x7c, 0x1d, 0x75, 0x62, 0x71, 0xcb, 0x4f,
	0x7e, 0x97, 0xc9, 0x5d, 0x97, 0x99, 0x8c, 0x6d, 0x72, 0x6e, 0x93, 0xdb, 0xd9, 0xd0, 0x7f, 0xb2,
	0x48, 0xf8, 0xca, 0xf9, 0x9d, 0x6b, 0x6c, 0xba, 0xc8, 0x59, 0x71, 0xac, 0x7f, 0x48, 0x5c, 0xf1,
	0xe5, 0xab, 0xf3, 0x4d, 0xba, 0xcc, 0x59, 0x71, 0x56, 0x26, 0x92, 0xa2, 0x65, 0x2c, 0xa6, 0xcd,
	0x7b, 0xe7, 0x1b, 0x4d, 0x4e, 0x95, 0x3c, 0xc5, 0xef, 0xf8, 0x3a, 0x38, 0x7d, 0x98, 0xb4, 0xc7,
	0x6a, 0x3b, 0x61, 0xbd, 0xa2, 0xe9, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xdb, 0xa6,
	0x77, 0x5d, 0x01, 0x00, 0x00,
}
