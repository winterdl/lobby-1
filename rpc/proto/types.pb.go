// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Empty response.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto1.RegisterFile("types.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 54 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xec, 0x5c, 0xac, 0xae,
	0xb9, 0x05, 0x25, 0x95, 0x49, 0x6c, 0x60, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0x63,
	0x09, 0x98, 0x25, 0x00, 0x00, 0x00,
}