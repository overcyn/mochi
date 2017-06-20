// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/image.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Image struct {
	Width  int64  `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Stride int64  `protobuf:"varint,4,opt,name=stride" json:"stride,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Image) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetStride() int64 {
	if m != nil {
		return m.Stride
	}
	return 0
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImageProperties struct {
	Width  int64   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Scale  float64 `protobuf:"fixed64,3,opt,name=scale" json:"scale,omitempty"`
}

func (m *ImageProperties) Reset()                    { *m = ImageProperties{} }
func (m *ImageProperties) String() string            { return proto.CompactTextString(m) }
func (*ImageProperties) ProtoMessage()               {}
func (*ImageProperties) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ImageProperties) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageProperties) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageProperties) GetScale() float64 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func init() {
	proto.RegisterType((*Image)(nil), "matcha.Image")
	proto.RegisterType((*ImageProperties)(nil), "matcha.ImageProperties")
}

func init() { proto.RegisterFile("github.com/overcyn/matcha/pb/image.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xbd, 0x0a, 0xc2, 0x30,
	0x14, 0x46, 0x49, 0xff, 0x90, 0x20, 0x14, 0x42, 0x91, 0x8c, 0xa5, 0x53, 0xa7, 0x66, 0xf0, 0x0d,
	0xba, 0x39, 0x08, 0xa5, 0xe0, 0xe2, 0x96, 0xa4, 0xa1, 0x09, 0x58, 0x13, 0xd2, 0xab, 0xe2, 0xeb,
	0xf8, 0xa4, 0xd2, 0xc4, 0x17, 0x70, 0xbb, 0xe7, 0x0c, 0xf7, 0xe3, 0xe0, 0x76, 0x36, 0xa0, 0x1f,
	0xa2, 0x93, 0x76, 0x61, 0xf6, 0xa9, 0xbc, 0x7c, 0xdf, 0xd9, 0xc2, 0x41, 0x6a, 0xce, 0x9c, 0x60,
	0x66, 0xe1, 0xb3, 0xea, 0x9c, 0xb7, 0x60, 0x49, 0x11, 0x75, 0xc3, 0x71, 0x7e, 0xda, 0x34, 0xa9,
	0x70, 0xfe, 0x32, 0x13, 0x68, 0x8a, 0x6a, 0xd4, 0xa6, 0x63, 0x04, 0x72, 0xc0, 0x85, 0x56, 0x66,
	0xd6, 0x40, 0x93, 0xa0, 0x7f, 0xb4, 0xf9, 0x15, 0xbc, 0x99, 0x14, 0xcd, 0xa2, 0x8f, 0x44, 0x08,
	0xce, 0x26, 0x0e, 0x9c, 0xa6, 0x35, 0x6a, 0xf7, 0x63, 0xb8, 0x9b, 0x0b, 0x2e, 0xc3, 0xc4, 0xe0,
	0xad, 0x53, 0x1e, 0x8c, 0x5a, 0xff, 0x1c, 0xab, 0x70, 0xbe, 0x4a, 0x7e, 0x53, 0xe1, 0x2b, 0x1a,
	0x23, 0xf4, 0xe5, 0x35, 0x71, 0xe2, 0x93, 0xec, 0xce, 0x21, 0x64, 0xe8, 0x45, 0x11, 0xca, 0x8e,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x0e, 0xa8, 0x64, 0x05, 0x01, 0x00, 0x00,
}
