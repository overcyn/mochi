// Code generated by protoc-gen-go. DO NOT EDIT.
// source: view.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Node struct {
	Id          int64                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BuildId     int64                `protobuf:"varint,2,opt,name=buildId" json:"buildId,omitempty"`
	LayoutId    int64                `protobuf:"varint,3,opt,name=layoutId" json:"layoutId,omitempty"`
	PaintId     int64                `protobuf:"varint,4,opt,name=paintId" json:"paintId,omitempty"`
	BridgeName  string               `protobuf:"bytes,7,opt,name=bridgeName" json:"bridgeName,omitempty"`
	BridgeValue *google_protobuf.Any `protobuf:"bytes,8,opt,name=bridgeValue" json:"bridgeValue,omitempty"`
	Children    []*Node              `protobuf:"bytes,5,rep,name=children" json:"children,omitempty"`
	LayoutGuide *Guide               `protobuf:"bytes,6,opt,name=layoutGuide" json:"layoutGuide,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *Node) GetLayoutId() int64 {
	if m != nil {
		return m.LayoutId
	}
	return 0
}

func (m *Node) GetPaintId() int64 {
	if m != nil {
		return m.PaintId
	}
	return 0
}

func (m *Node) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *Node) GetBridgeValue() *google_protobuf.Any {
	if m != nil {
		return m.BridgeValue
	}
	return nil
}

func (m *Node) GetChildren() []*Node {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Node) GetLayoutGuide() *Guide {
	if m != nil {
		return m.LayoutGuide
	}
	return nil
}

type Root struct {
	Node *Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Root) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "view.Node")
	proto.RegisterType((*Root)(nil), "view.Root")
}

func init() { proto.RegisterFile("view.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x37, 0x34, 0xe1, 0x02, 0x1d, 0x2c, 0x06, 0x93, 0xa1, 0x8a, 0x3a, 0x54, 0x99,
	0x1c, 0xa9, 0x48, 0xec, 0x74, 0x41, 0x1d, 0xa8, 0x90, 0x07, 0x06, 0xb6, 0xa4, 0x67, 0x52, 0x4b,
	0x21, 0x17, 0x45, 0x09, 0x28, 0xbf, 0x84, 0x9d, 0x5f, 0x8a, 0x62, 0xb7, 0x55, 0xd8, 0xfc, 0xbe,
	0x77, 0x77, 0x7a, 0x7e, 0x00, 0x5f, 0x46, 0x7f, 0xcb, 0xa6, 0xa5, 0x8e, 0xb8, 0x3f, 0xbe, 0xe3,
	0x9b, 0x2a, 0x1f, 0xa8, 0xef, 0x1c, 0x8b, 0xef, 0x4b, 0xa2, 0xb2, 0xd2, 0x99, 0x55, 0x45, 0xff,
	0x91, 0xe5, 0xf5, 0xe0, 0xac, 0xd5, 0x0f, 0x03, 0x7f, 0x4f, 0xa8, 0xf9, 0x02, 0x98, 0x41, 0xe1,
	0x25, 0x5e, 0x3a, 0x53, 0xcc, 0x20, 0x17, 0x10, 0x14, 0xbd, 0xa9, 0x70, 0x87, 0x82, 0x59, 0x78,
	0x96, 0x3c, 0x86, 0xd0, 0x5d, 0xdf, 0xa1, 0x98, 0x59, 0xeb, 0xa2, 0xc7, 0xad, 0x26, 0x37, 0xf5,
	0x68, 0xf9, 0x6e, 0xeb, 0x24, 0xf9, 0x12, 0xa0, 0x68, 0x0d, 0x96, 0x7a, 0x9f, 0x7f, 0x6a, 0x11,
	0x24, 0x5e, 0x7a, 0xad, 0x26, 0x84, 0x3f, 0x42, 0xe4, 0xd4, 0x5b, 0x5e, 0xf5, 0x5a, 0x84, 0x89,
	0x97, 0x46, 0x9b, 0x3b, 0xe9, 0x92, 0xcb, 0x73, 0x72, 0xf9, 0x54, 0x0f, 0x6a, 0x3a, 0xc8, 0xd7,
	0x10, 0x1e, 0x8e, 0xa6, 0xc2, 0x56, 0xd7, 0xe2, 0x2a, 0x99, 0xa5, 0xd1, 0x06, 0xa4, 0xad, 0x63,
	0xfc, 0x95, 0xba, 0x78, 0x3c, 0x83, 0xc8, 0xa5, 0x7c, 0xee, 0x0d, 0x6a, 0x31, 0xb7, 0xf7, 0x6f,
	0xe5, 0xa9, 0x27, 0x0b, 0xd5, 0x74, 0x62, 0xb5, 0x06, 0x5f, 0x11, 0x75, 0x7c, 0x09, 0x7e, 0x4d,
	0xa8, 0x6d, 0x35, 0xff, 0x8f, 0x5b, 0xbe, 0x5d, 0xbc, 0xb3, 0xa6, 0xf8, 0x65, 0xc1, 0x0b, 0x1d,
	0x8e, 0xe6, 0x75, 0x5b, 0xcc, 0x6d, 0xd6, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2c,
	0x59, 0x04, 0x95, 0x01, 0x00, 0x00,
}
