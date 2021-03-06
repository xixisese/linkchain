// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}

var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

type Test struct {
	Label                *string  `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type                 *int32   `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps                 []int64  `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

type TestStringArray struct {
	S1                   *string  `protobuf:"bytes,1,req,name=s1" json:"s1,omitempty"`
	S2                   *string  `protobuf:"bytes,2,req,name=s2" json:"s2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestStringArray) Reset()         { *m = TestStringArray{} }
func (m *TestStringArray) String() string { return proto.CompactTextString(m) }
func (*TestStringArray) ProtoMessage()    {}
func (*TestStringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *TestStringArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestStringArray.Unmarshal(m, b)
}
func (m *TestStringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestStringArray.Marshal(b, m, deterministic)
}
func (m *TestStringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestStringArray.Merge(m, src)
}
func (m *TestStringArray) XXX_Size() int {
	return xxx_messageInfo_TestStringArray.Size(m)
}
func (m *TestStringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_TestStringArray.DiscardUnknown(m)
}

var xxx_messageInfo_TestStringArray proto.InternalMessageInfo

func (m *TestStringArray) GetS1() string {
	if m != nil && m.S1 != nil {
		return *m.S1
	}
	return ""
}

func (m *TestStringArray) GetS2() string {
	if m != nil && m.S2 != nil {
		return *m.S2
	}
	return ""
}

type TestUint struct {
	U                    *uint32  `protobuf:"varint,1,req,name=u" json:"u,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestUint) Reset()         { *m = TestUint{} }
func (m *TestUint) String() string { return proto.CompactTextString(m) }
func (*TestUint) ProtoMessage()    {}
func (*TestUint) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *TestUint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUint.Unmarshal(m, b)
}
func (m *TestUint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUint.Marshal(b, m, deterministic)
}
func (m *TestUint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUint.Merge(m, src)
}
func (m *TestUint) XXX_Size() int {
	return xxx_messageInfo_TestUint.Size(m)
}
func (m *TestUint) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUint.DiscardUnknown(m)
}

var xxx_messageInfo_TestUint proto.InternalMessageInfo

func (m *TestUint) GetU() uint32 {
	if m != nil && m.U != nil {
		return *m.U
	}
	return 0
}

func init() {
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterType((*TestStringArray)(nil), "example.TestStringArray")
	proto.RegisterType((*TestUint)(nil), "example.TestUint")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x3c, 0xb8,
	0x58, 0x42, 0x52, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x73, 0x12, 0x93, 0x52, 0x73, 0x24, 0x18,
	0x15, 0x98, 0x34, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x56, 0x2b, 0x26, 0x73, 0xf3, 0x20, 0x30, 0x5f, 0x48, 0x88, 0x8b, 0xa5,
	0x28, 0xb5, 0xa0, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0x39, 0x08, 0xcc, 0x56, 0x32, 0xe4, 0xe2,
	0x07, 0x99, 0x14, 0x5c, 0x52, 0x94, 0x99, 0x97, 0xee, 0x58, 0x54, 0x94, 0x58, 0x29, 0xc4, 0xc7,
	0xc5, 0x54, 0x6c, 0x08, 0x35, 0x91, 0xa9, 0xd8, 0x10, 0xcc, 0x37, 0x92, 0x60, 0x82, 0xf2, 0x8d,
	0x94, 0x24, 0xb8, 0x38, 0x40, 0x5a, 0x42, 0x33, 0xf3, 0x4a, 0x84, 0x78, 0xb8, 0x18, 0x4b, 0xc1,
	0x4a, 0x79, 0x83, 0x18, 0x4b, 0xb5, 0x78, 0xb8, 0x98, 0xdd, 0xfc, 0xfd, 0x85, 0x58, 0xb9, 0x18,
	0x23, 0x04, 0x04, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xe5, 0x5a, 0x88, 0xbd, 0x00, 0x00,
	0x00,
}
