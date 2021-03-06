// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/account.proto

package protobuf

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

type Account struct {
	Id                   *AccountID `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Type                 *uint32    `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	Utxos                []*UTXO    `protobuf:"bytes,3,rep,name=utxos" json:"utxos,omitempty"`
	Clear                *ClearTime `protobuf:"bytes,4,opt,name=clear" json:"clear,omitempty"`
	SecurityId           *AccountID `protobuf:"bytes,5,opt,name=securityId" json:"securityId,omitempty"`
	StorageRoot          *Hash      `protobuf:"bytes,6,opt,name=storageRoot" json:"storageRoot,omitempty"`
	CodeHash             *Hash      `protobuf:"bytes,7,opt,name=codeHash" json:"codeHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a8b28a2e7a7402, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() *AccountID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Account) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Account) GetUtxos() []*UTXO {
	if m != nil {
		return m.Utxos
	}
	return nil
}

func (m *Account) GetClear() *ClearTime {
	if m != nil {
		return m.Clear
	}
	return nil
}

func (m *Account) GetSecurityId() *AccountID {
	if m != nil {
		return m.SecurityId
	}
	return nil
}

func (m *Account) GetStorageRoot() *Hash {
	if m != nil {
		return m.StorageRoot
	}
	return nil
}

func (m *Account) GetCodeHash() *Hash {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

type UTXO struct {
	Id                   *Ticket  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	LocatedHeight        *uint32  `protobuf:"varint,2,req,name=locatedHeight" json:"locatedHeight,omitempty"`
	EffectHeight         *uint32  `protobuf:"varint,3,req,name=effectHeight" json:"effectHeight,omitempty"`
	Value                []byte   `protobuf:"bytes,4,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTXO) Reset()         { *m = UTXO{} }
func (m *UTXO) String() string { return proto.CompactTextString(m) }
func (*UTXO) ProtoMessage()    {}
func (*UTXO) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a8b28a2e7a7402, []int{1}
}

func (m *UTXO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTXO.Unmarshal(m, b)
}
func (m *UTXO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTXO.Marshal(b, m, deterministic)
}
func (m *UTXO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTXO.Merge(m, src)
}
func (m *UTXO) XXX_Size() int {
	return xxx_messageInfo_UTXO.Size(m)
}
func (m *UTXO) XXX_DiscardUnknown() {
	xxx_messageInfo_UTXO.DiscardUnknown(m)
}

var xxx_messageInfo_UTXO proto.InternalMessageInfo

func (m *UTXO) GetId() *Ticket {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UTXO) GetLocatedHeight() uint32 {
	if m != nil && m.LocatedHeight != nil {
		return *m.LocatedHeight
	}
	return 0
}

func (m *UTXO) GetEffectHeight() uint32 {
	if m != nil && m.EffectHeight != nil {
		return *m.EffectHeight
	}
	return 0
}

func (m *UTXO) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ClearTime struct {
	LastClearTime        *int64   `protobuf:"varint,1,req,name=lastClearTime" json:"lastClearTime,omitempty"`
	LastEffectHeight     *uint32  `protobuf:"varint,2,req,name=lastEffectHeight" json:"lastEffectHeight,omitempty"`
	NextClearTime        *int64   `protobuf:"varint,3,req,name=nextClearTime" json:"nextClearTime,omitempty"`
	NextEffectHeight     *uint32  `protobuf:"varint,4,req,name=nextEffectHeight" json:"nextEffectHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearTime) Reset()         { *m = ClearTime{} }
func (m *ClearTime) String() string { return proto.CompactTextString(m) }
func (*ClearTime) ProtoMessage()    {}
func (*ClearTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a8b28a2e7a7402, []int{2}
}

func (m *ClearTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearTime.Unmarshal(m, b)
}
func (m *ClearTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearTime.Marshal(b, m, deterministic)
}
func (m *ClearTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearTime.Merge(m, src)
}
func (m *ClearTime) XXX_Size() int {
	return xxx_messageInfo_ClearTime.Size(m)
}
func (m *ClearTime) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearTime.DiscardUnknown(m)
}

var xxx_messageInfo_ClearTime proto.InternalMessageInfo

func (m *ClearTime) GetLastClearTime() int64 {
	if m != nil && m.LastClearTime != nil {
		return *m.LastClearTime
	}
	return 0
}

func (m *ClearTime) GetLastEffectHeight() uint32 {
	if m != nil && m.LastEffectHeight != nil {
		return *m.LastEffectHeight
	}
	return 0
}

func (m *ClearTime) GetNextClearTime() int64 {
	if m != nil && m.NextClearTime != nil {
		return *m.NextClearTime
	}
	return 0
}

func (m *ClearTime) GetNextEffectHeight() uint32 {
	if m != nil && m.NextEffectHeight != nil {
		return *m.NextEffectHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "protobuf.Account")
	proto.RegisterType((*UTXO)(nil), "protobuf.UTXO")
	proto.RegisterType((*ClearTime)(nil), "protobuf.ClearTime")
}

func init() { proto.RegisterFile("protobuf/account.proto", fileDescriptor_f3a8b28a2e7a7402) }

var fileDescriptor_f3a8b28a2e7a7402 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0xd3, 0x7f, 0x17, 0xee, 0x01, 0x6e, 0xc8, 0x5c, 0x63, 0x26, 0xac, 0x9a, 0xca, 0xa2,
	0xb2, 0x40, 0x83, 0x4f, 0x60, 0xd4, 0x04, 0x56, 0x26, 0x13, 0x4c, 0xdc, 0x8e, 0xd3, 0x01, 0x1a,
	0xb1, 0x43, 0xda, 0x53, 0x03, 0x6f, 0xe0, 0x93, 0xb8, 0xf5, 0x15, 0xcd, 0xcc, 0x40, 0x69, 0xd1,
	0x5d, 0xfb, 0x9d, 0xef, 0xfc, 0xce, 0x99, 0x03, 0xe7, 0x9b, 0x5c, 0xa1, 0x7a, 0x29, 0x17, 0x57,
	0x5c, 0x08, 0x55, 0x66, 0x38, 0x36, 0x80, 0xb4, 0x0f, 0x7c, 0x30, 0xa8, 0x0c, 0xcc, 0x79, 0x56,
	0x70, 0x81, 0xa9, 0xca, 0xac, 0x15, 0x7d, 0xba, 0xd0, 0xba, 0xb5, 0x7d, 0xe4, 0x02, 0xdc, 0x34,
	0xa1, 0x4e, 0xe8, 0xc6, 0x9d, 0xc9, 0xff, 0xf1, 0xa1, 0x69, 0xbc, 0x2f, 0xcf, 0xee, 0x99, 0x9b,
	0x26, 0x84, 0x80, 0x8f, 0xbb, 0x8d, 0xa4, 0x6e, 0xe8, 0xc6, 0x3d, 0x66, 0xbe, 0xc9, 0x10, 0x82,
	0x12, 0xb7, 0xaa, 0xa0, 0x5e, 0xe8, 0xc5, 0x9d, 0xc9, 0xbf, 0x63, 0xef, 0xd3, 0xfc, 0xf9, 0x91,
	0xd9, 0x22, 0xb9, 0x84, 0x40, 0xac, 0x25, 0xcf, 0xa9, 0x1f, 0x3a, 0xcd, 0x09, 0x77, 0x1a, 0xcf,
	0xd3, 0x37, 0xc9, 0xac, 0x41, 0x6e, 0x00, 0x0a, 0x29, 0xca, 0x3c, 0xc5, 0xdd, 0x2c, 0xa1, 0xc1,
	0xa9, 0x7f, 0xdc, 0xa8, 0xa6, 0x91, 0x6b, 0xe8, 0x14, 0xa8, 0x72, 0xbe, 0x94, 0x4c, 0x29, 0xa4,
	0x7f, 0x4c, 0x57, 0x6d, 0x97, 0x29, 0x2f, 0x56, 0xac, 0xae, 0x90, 0x11, 0xb4, 0x85, 0x4a, 0xa4,
	0x2e, 0xd0, 0xd6, 0xaf, 0x7a, 0x55, 0x8f, 0x3e, 0x1c, 0xf0, 0xf5, 0x6b, 0x48, 0x58, 0xbb, 0x52,
	0xff, 0xa8, 0xcf, 0x53, 0xf1, 0x2a, 0xd1, 0x9c, 0x68, 0x08, 0xbd, 0xb5, 0x12, 0x1c, 0x65, 0x32,
	0x95, 0xe9, 0x72, 0x85, 0xfb, 0x5b, 0x35, 0x21, 0x89, 0xa0, 0x2b, 0x17, 0x0b, 0x29, 0x70, 0x2f,
	0x79, 0x46, 0x6a, 0x30, 0x72, 0x06, 0xc1, 0x3b, 0x5f, 0x97, 0x92, 0xfa, 0xa1, 0x1b, 0x77, 0x99,
	0xfd, 0x89, 0xbe, 0x1c, 0xf8, 0x5b, 0x9d, 0xcc, 0x4c, 0xe3, 0x05, 0x56, 0xc0, 0xac, 0xe6, 0xb1,
	0x26, 0x24, 0x23, 0xe8, 0x6b, 0xf0, 0x50, 0x9f, 0x68, 0xd7, 0xfa, 0xc1, 0x75, 0x62, 0x26, 0xb7,
	0xb5, 0x44, 0xcf, 0x26, 0x36, 0xa0, 0x4e, 0xd4, 0xa0, 0x91, 0xe8, 0xdb, 0xc4, 0x53, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x90, 0x00, 0x0c, 0xa4, 0x02, 0x00, 0x00,
}
