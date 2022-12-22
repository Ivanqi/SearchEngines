// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

type FalconSearchStoreInfo struct {
	StoreName            string   `protobuf:"bytes,1,opt,name=StoreName,proto3" json:"StoreName,omitempty"`
	StoreLength          int64    `protobuf:"varint,2,opt,name=StoreLength,proto3" json:"StoreLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FalconSearchStoreInfo) Reset()         { *m = FalconSearchStoreInfo{} }
func (m *FalconSearchStoreInfo) String() string { return proto.CompactTextString(m) }
func (*FalconSearchStoreInfo) ProtoMessage()    {}
func (*FalconSearchStoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *FalconSearchStoreInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FalconSearchStoreInfo.Unmarshal(m, b)
}
func (m *FalconSearchStoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FalconSearchStoreInfo.Marshal(b, m, deterministic)
}
func (m *FalconSearchStoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FalconSearchStoreInfo.Merge(m, src)
}
func (m *FalconSearchStoreInfo) XXX_Size() int {
	return xxx_messageInfo_FalconSearchStoreInfo.Size(m)
}
func (m *FalconSearchStoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FalconSearchStoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FalconSearchStoreInfo proto.InternalMessageInfo

func (m *FalconSearchStoreInfo) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *FalconSearchStoreInfo) GetStoreLength() int64 {
	if m != nil {
		return m.StoreLength
	}
	return 0
}

type BinlogMessage struct {
	LogId                int64    `protobuf:"varint,1,opt,name=LogId,proto3" json:"LogId,omitempty"`
	LogTimestamp         int64    `protobuf:"varint,2,opt,name=LogTimestamp,proto3" json:"LogTimestamp,omitempty"`
	LogDetail            []byte   `protobuf:"bytes,3,opt,name=LogDetail,proto3" json:"LogDetail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinlogMessage) Reset()         { *m = BinlogMessage{} }
func (m *BinlogMessage) String() string { return proto.CompactTextString(m) }
func (*BinlogMessage) ProtoMessage()    {}
func (*BinlogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *BinlogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinlogMessage.Unmarshal(m, b)
}
func (m *BinlogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinlogMessage.Marshal(b, m, deterministic)
}
func (m *BinlogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinlogMessage.Merge(m, src)
}
func (m *BinlogMessage) XXX_Size() int {
	return xxx_messageInfo_BinlogMessage.Size(m)
}
func (m *BinlogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BinlogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BinlogMessage proto.InternalMessageInfo

func (m *BinlogMessage) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *BinlogMessage) GetLogTimestamp() int64 {
	if m != nil {
		return m.LogTimestamp
	}
	return 0
}

func (m *BinlogMessage) GetLogDetail() []byte {
	if m != nil {
		return m.LogDetail
	}
	return nil
}

type DictValue struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Length               uint64   `protobuf:"varint,2,opt,name=Length,proto3" json:"Length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DictValue) Reset()         { *m = DictValue{} }
func (m *DictValue) String() string { return proto.CompactTextString(m) }
func (*DictValue) ProtoMessage()    {}
func (*DictValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *DictValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictValue.Unmarshal(m, b)
}
func (m *DictValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictValue.Marshal(b, m, deterministic)
}
func (m *DictValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictValue.Merge(m, src)
}
func (m *DictValue) XXX_Size() int {
	return xxx_messageInfo_DictValue.Size(m)
}
func (m *DictValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DictValue.DiscardUnknown(m)
}

var xxx_messageInfo_DictValue proto.InternalMessageInfo

func (m *DictValue) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DictValue) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type DocId struct {
	DocID                uint32   `protobuf:"varint,1,opt,name=DocID,proto3" json:"DocID,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=Weight,proto3" json:"Weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocId) Reset()         { *m = DocId{} }
func (m *DocId) String() string { return proto.CompactTextString(m) }
func (*DocId) ProtoMessage()    {}
func (*DocId) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *DocId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocId.Unmarshal(m, b)
}
func (m *DocId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocId.Marshal(b, m, deterministic)
}
func (m *DocId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocId.Merge(m, src)
}
func (m *DocId) XXX_Size() int {
	return xxx_messageInfo_DocId.Size(m)
}
func (m *DocId) XXX_DiscardUnknown() {
	xxx_messageInfo_DocId.DiscardUnknown(m)
}

var xxx_messageInfo_DocId proto.InternalMessageInfo

func (m *DocId) GetDocID() uint32 {
	if m != nil {
		return m.DocID
	}
	return 0
}

func (m *DocId) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type FieldValue struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValue) Reset()         { *m = FieldValue{} }
func (m *FieldValue) String() string { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()    {}
func (*FieldValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *FieldValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValue.Unmarshal(m, b)
}
func (m *FieldValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValue.Marshal(b, m, deterministic)
}
func (m *FieldValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValue.Merge(m, src)
}
func (m *FieldValue) XXX_Size() int {
	return xxx_messageInfo_FieldValue.Size(m)
}
func (m *FieldValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValue.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValue proto.InternalMessageInfo

func (m *FieldValue) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *FieldValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Document struct {
	Val                  []*FieldValue `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetVal() []*FieldValue {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*FalconSearchStoreInfo)(nil), "message.FalconSearchStoreInfo")
	proto.RegisterType((*BinlogMessage)(nil), "message.BinlogMessage")
	proto.RegisterType((*DictValue)(nil), "message.DictValue")
	proto.RegisterType((*DocId)(nil), "message.DocId")
	proto.RegisterType((*FieldValue)(nil), "message.FieldValue")
	proto.RegisterType((*Document)(nil), "message.Document")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0xdd, 0xa6, 0x7b, 0x5b, 0x2f, 0x71, 0x4a, 0x0f, 0x1e, 0x4a, 0x40, 0xe8, 0x69,
	0xa0, 0x22, 0x08, 0xde, 0xa4, 0x0c, 0x06, 0x55, 0x21, 0x13, 0x77, 0x8e, 0xdd, 0x6b, 0x16, 0x48,
	0x9b, 0xd1, 0x66, 0xfb, 0xfb, 0x25, 0x3f, 0xa4, 0xf3, 0x96, 0xcf, 0xe7, 0xf1, 0x7d, 0xef, 0x91,
	0x07, 0x49, 0x83, 0x7d, 0xcf, 0x05, 0x2e, 0x0f, 0x9d, 0x36, 0x9a, 0x5c, 0x06, 0xa4, 0x5b, 0xb8,
	0x59, 0x71, 0x55, 0xe9, 0x76, 0x83, 0xbc, 0xab, 0xf6, 0x1b, 0xa3, 0x3b, 0x5c, 0xb7, 0xb5, 0x26,
	0x77, 0x30, 0x75, 0xf0, 0xc1, 0x1b, 0x4c, 0xa3, 0x2c, 0xca, 0xa7, 0x6c, 0x10, 0x24, 0x83, 0x99,
	0x83, 0x12, 0x5b, 0x61, 0xf6, 0xe9, 0x45, 0x16, 0xe5, 0x31, 0x3b, 0x57, 0x54, 0x40, 0xf2, 0x26,
	0x5b, 0xa5, 0xc5, 0xbb, 0x9f, 0x44, 0x16, 0x30, 0x2e, 0xb5, 0x58, 0xef, 0x5c, 0xb3, 0x98, 0x79,
	0x20, 0x14, 0xe6, 0xa5, 0x16, 0x5f, 0xb2, 0xc1, 0xde, 0xf0, 0xe6, 0x10, 0x3a, 0xfd, 0x73, 0x76,
	0x95, 0x52, 0x8b, 0x02, 0x0d, 0x97, 0x2a, 0x8d, 0xb3, 0x28, 0x9f, 0xb3, 0x41, 0xd0, 0x57, 0x98,
	0x16, 0xb2, 0x32, 0xdf, 0x5c, 0x1d, 0x91, 0xdc, 0xc2, 0xe4, 0xb3, 0xae, 0x7b, 0x34, 0x6e, 0xca,
	0x88, 0x05, 0xb2, 0xfe, 0x6c, 0xd5, 0x11, 0x0b, 0x44, 0x9f, 0x61, 0x5c, 0xe8, 0x6a, 0xbd, 0xb3,
	0xdb, 0xd9, 0x47, 0xe1, 0x72, 0x09, 0xf3, 0x60, 0x63, 0x5b, 0x94, 0x62, 0x6f, 0x5c, 0x2c, 0x61,
	0x81, 0xe8, 0x0b, 0xc0, 0x4a, 0xa2, 0xda, 0xf9, 0xa1, 0x0b, 0x18, 0xd7, 0x96, 0xc2, 0x37, 0x79,
	0xb0, 0xf6, 0x64, 0xcb, 0x2e, 0x3a, 0x67, 0x1e, 0xe8, 0x03, 0x5c, 0x15, 0xba, 0x3a, 0x36, 0xd8,
	0x1a, 0x72, 0x0f, 0xf1, 0x89, 0xab, 0x34, 0xca, 0xe2, 0x7c, 0xf6, 0x78, 0xbd, 0xfc, 0xbb, 0xd0,
	0xd0, 0x99, 0xd9, 0xfa, 0xcf, 0xc4, 0x9d, 0xec, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x34, 0xea,
	0x78, 0x10, 0xc3, 0x01, 0x00, 0x00,
}