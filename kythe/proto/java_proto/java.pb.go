// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/java.proto

/*
Package java_go_proto is a generated protocol buffer package.

It is generated from these files:
	kythe/proto/java.proto

It has these top-level messages:
	JavaDetails
*/
package java_go_proto

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

type JavaDetails struct {
	Classpath      []string `protobuf:"bytes,1,rep,name=classpath" json:"classpath,omitempty"`
	Sourcepath     []string `protobuf:"bytes,2,rep,name=sourcepath" json:"sourcepath,omitempty"`
	Bootclasspath  []string `protobuf:"bytes,3,rep,name=bootclasspath" json:"bootclasspath,omitempty"`
	ExtraJavacopts []string `protobuf:"bytes,10,rep,name=extra_javacopts,json=extraJavacopts" json:"extra_javacopts,omitempty"`
}

func (m *JavaDetails) Reset()                    { *m = JavaDetails{} }
func (m *JavaDetails) String() string            { return proto.CompactTextString(m) }
func (*JavaDetails) ProtoMessage()               {}
func (*JavaDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JavaDetails) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaDetails) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaDetails) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

func (m *JavaDetails) GetExtraJavacopts() []string {
	if m != nil {
		return m.ExtraJavacopts
	}
	return nil
}

func init() {
	proto.RegisterType((*JavaDetails)(nil), "kythe.proto.JavaDetails")
}

func init() { proto.RegisterFile("kythe/proto/java.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xae, 0x2c, 0xc9,
	0x48, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4a, 0x2c, 0x4b, 0xd4, 0x03, 0x33, 0x85,
	0xb8, 0xc1, 0xe2, 0x10, 0x8e, 0xd2, 0x2c, 0x46, 0x2e, 0x6e, 0xaf, 0xc4, 0xb2, 0x44, 0x97, 0xd4,
	0x92, 0xc4, 0xcc, 0x9c, 0x62, 0x21, 0x19, 0x2e, 0xce, 0xe4, 0x9c, 0xc4, 0xe2, 0xe2, 0x82, 0xc4,
	0x92, 0x0c, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x90, 0x1c, 0x17, 0x57, 0x71,
	0x7e, 0x69, 0x51, 0x72, 0x2a, 0x58, 0x9a, 0x09, 0x2c, 0x8d, 0x24, 0x22, 0xa4, 0xc2, 0xc5, 0x9b,
	0x94, 0x9f, 0x5f, 0x82, 0x30, 0x81, 0x19, 0xac, 0x04, 0x55, 0x50, 0x48, 0x9d, 0x8b, 0x3f, 0xb5,
	0xa2, 0xa4, 0x28, 0x31, 0x1e, 0xe4, 0xa8, 0xe4, 0xfc, 0x82, 0x92, 0x62, 0x09, 0x2e, 0xb0, 0x3a,
	0x3e, 0xb0, 0xb0, 0x17, 0x4c, 0xd4, 0x49, 0x91, 0x4b, 0x3e, 0x39, 0x3f, 0x57, 0x2f, 0x3d, 0x3f,
	0x3f, 0x3d, 0x27, 0x55, 0x2f, 0x25, 0xb5, 0xac, 0x24, 0x3f, 0x3f, 0xa7, 0x58, 0x0f, 0xc9, 0xfd,
	0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x15, 0x22, 0x18, 0xed, 0x00,
	0x00, 0x00,
}
