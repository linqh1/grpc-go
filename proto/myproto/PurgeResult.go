// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto/PurgeResult.proto

package myproto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PurgeReuslt struct {
	Total                int32                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Results              []*SinglePurgeResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurgeReuslt) Reset()         { *m = PurgeReuslt{} }
func (m *PurgeReuslt) String() string { return proto.CompactTextString(m) }
func (*PurgeReuslt) ProtoMessage()    {}
func (*PurgeReuslt) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc0b6b18da1e2e0, []int{0}
}

func (m *PurgeReuslt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeReuslt.Unmarshal(m, b)
}
func (m *PurgeReuslt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeReuslt.Marshal(b, m, deterministic)
}
func (m *PurgeReuslt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeReuslt.Merge(m, src)
}
func (m *PurgeReuslt) XXX_Size() int {
	return xxx_messageInfo_PurgeReuslt.Size(m)
}
func (m *PurgeReuslt) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeReuslt.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeReuslt proto.InternalMessageInfo

func (m *PurgeReuslt) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PurgeReuslt) GetResults() []*SinglePurgeResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SinglePurgeResult struct {
	PurgeId              string        `protobuf:"bytes,1,opt,name=purgeId,proto3" json:"purgeId,omitempty"`
	Status               string        `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Errors               []*PurgeError `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SinglePurgeResult) Reset()         { *m = SinglePurgeResult{} }
func (m *SinglePurgeResult) String() string { return proto.CompactTextString(m) }
func (*SinglePurgeResult) ProtoMessage()    {}
func (*SinglePurgeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc0b6b18da1e2e0, []int{1}
}

func (m *SinglePurgeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SinglePurgeResult.Unmarshal(m, b)
}
func (m *SinglePurgeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SinglePurgeResult.Marshal(b, m, deterministic)
}
func (m *SinglePurgeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SinglePurgeResult.Merge(m, src)
}
func (m *SinglePurgeResult) XXX_Size() int {
	return xxx_messageInfo_SinglePurgeResult.Size(m)
}
func (m *SinglePurgeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SinglePurgeResult.DiscardUnknown(m)
}

var xxx_messageInfo_SinglePurgeResult proto.InternalMessageInfo

func (m *SinglePurgeResult) GetPurgeId() string {
	if m != nil {
		return m.PurgeId
	}
	return ""
}

func (m *SinglePurgeResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SinglePurgeResult) GetErrors() []*PurgeError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type PurgeError struct {
	ErrorCode            string   `protobuf:"bytes,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeError) Reset()         { *m = PurgeError{} }
func (m *PurgeError) String() string { return proto.CompactTextString(m) }
func (*PurgeError) ProtoMessage()    {}
func (*PurgeError) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc0b6b18da1e2e0, []int{2}
}

func (m *PurgeError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeError.Unmarshal(m, b)
}
func (m *PurgeError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeError.Marshal(b, m, deterministic)
}
func (m *PurgeError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeError.Merge(m, src)
}
func (m *PurgeError) XXX_Size() int {
	return xxx_messageInfo_PurgeError.Size(m)
}
func (m *PurgeError) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeError.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeError proto.InternalMessageInfo

func (m *PurgeError) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *PurgeError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*PurgeReuslt)(nil), "myproto.PurgeReuslt")
	proto.RegisterType((*SinglePurgeResult)(nil), "myproto.SinglePurgeResult")
	proto.RegisterType((*PurgeError)(nil), "myproto.PurgeError")
}

func init() { proto.RegisterFile("myproto/PurgeResult.proto", fileDescriptor_0cc0b6b18da1e2e0) }

var fileDescriptor_0cc0b6b18da1e2e0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8d, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0x47, 0x71, 0x8d, 0x6d, 0x7c, 0xee, 0xd2, 0x6b, 0x29, 0x6a, 0xe9, 0x60, 0x34, 0x19, 0x0a,
	0x2e, 0xb4, 0xfd, 0x06, 0xa5, 0x43, 0x87, 0x86, 0xa0, 0x4c, 0x19, 0x1d, 0x2c, 0x4c, 0x40, 0x89,
	0x8c, 0xfe, 0x0c, 0xf9, 0xf6, 0x41, 0x17, 0x39, 0x4e, 0xc8, 0xf8, 0x7b, 0x4f, 0xa7, 0x07, 0x2f,
	0xbb, 0xc3, 0x68, 0xb4, 0xd3, 0x1f, 0x4b, 0x6f, 0x06, 0x29, 0xa4, 0xf5, 0xca, 0xb5, 0x44, 0xb0,
	0x88, 0x8a, 0xaf, 0xa1, 0x8a, 0xd6, 0x5b, 0xe5, 0xf0, 0x09, 0x32, 0xa7, 0x5d, 0xa7, 0x58, 0x52,
	0x27, 0x4d, 0x26, 0x4e, 0x03, 0xbf, 0xa1, 0x30, 0x74, 0x6d, 0xd9, 0x5d, 0x9d, 0x36, 0xd5, 0xe7,
	0x6b, 0x1b, 0xef, 0xdb, 0xd5, 0x76, 0x3f, 0x28, 0x79, 0x11, 0x10, 0xd3, 0x53, 0x6e, 0xe0, 0xe1,
	0xc6, 0x22, 0x83, 0x62, 0x0c, 0xf3, 0xaf, 0xa7, 0x44, 0x29, 0xa6, 0x89, 0xcf, 0x90, 0x5b, 0xd7,
	0x39, 0x1f, 0x1a, 0x41, 0xc4, 0x85, 0xef, 0x90, 0x4b, 0x63, 0xb4, 0xb1, 0x2c, 0xa5, 0xf6, 0xe3,
	0xb9, 0x4d, 0xff, 0xfe, 0x06, 0x27, 0xe2, 0x13, 0xbe, 0x00, 0x98, 0x29, 0xbe, 0x41, 0x49, 0xfc,
	0x47, 0xf7, 0x32, 0xe6, 0x66, 0x80, 0x1c, 0xee, 0x69, 0xfc, 0x4b, 0x6b, 0xbb, 0x41, 0xc6, 0xec,
	0x15, 0xdb, 0xe4, 0x54, 0xfa, 0x3a, 0x06, 0x00, 0x00, 0xff, 0xff, 0x86, 0x7d, 0x02, 0xfa, 0x4b,
	0x01, 0x00, 0x00,
}