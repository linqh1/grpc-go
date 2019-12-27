// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/DispatchTask.proto

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

type DispatchTask struct {
	TaskId               string       `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Details              []*PurgeTask `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DispatchTask) Reset()         { *m = DispatchTask{} }
func (m *DispatchTask) String() string { return proto.CompactTextString(m) }
func (*DispatchTask) ProtoMessage()    {}
func (*DispatchTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_5709154d174c5578, []int{0}
}

func (m *DispatchTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DispatchTask.Unmarshal(m, b)
}
func (m *DispatchTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DispatchTask.Marshal(b, m, deterministic)
}
func (m *DispatchTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DispatchTask.Merge(m, src)
}
func (m *DispatchTask) XXX_Size() int {
	return xxx_messageInfo_DispatchTask.Size(m)
}
func (m *DispatchTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DispatchTask.DiscardUnknown(m)
}

var xxx_messageInfo_DispatchTask proto.InternalMessageInfo

func (m *DispatchTask) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DispatchTask) GetDetails() []*PurgeTask {
	if m != nil {
		return m.Details
	}
	return nil
}

type PurgeTask struct {
	PurgeId              string            `protobuf:"bytes,1,opt,name=purgeId,proto3" json:"purgeId,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content              []string          `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty"`
	Headers              map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Action               string            `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PurgeTask) Reset()         { *m = PurgeTask{} }
func (m *PurgeTask) String() string { return proto.CompactTextString(m) }
func (*PurgeTask) ProtoMessage()    {}
func (*PurgeTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_5709154d174c5578, []int{1}
}

func (m *PurgeTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeTask.Unmarshal(m, b)
}
func (m *PurgeTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeTask.Marshal(b, m, deterministic)
}
func (m *PurgeTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeTask.Merge(m, src)
}
func (m *PurgeTask) XXX_Size() int {
	return xxx_messageInfo_PurgeTask.Size(m)
}
func (m *PurgeTask) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeTask.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeTask proto.InternalMessageInfo

func (m *PurgeTask) GetPurgeId() string {
	if m != nil {
		return m.PurgeId
	}
	return ""
}

func (m *PurgeTask) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PurgeTask) GetContent() []string {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PurgeTask) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *PurgeTask) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*DispatchTask)(nil), "myproto.DispatchTask")
	proto.RegisterType((*PurgeTask)(nil), "myproto.PurgeTask")
	proto.RegisterMapType((map[string]string)(nil), "myproto.PurgeTask.HeadersEntry")
}

func init() { proto.RegisterFile("proto/DispatchTask.proto", fileDescriptor_5709154d174c5578) }

var fileDescriptor_5709154d174c5578 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xd2, 0x36, 0x74, 0xec, 0x41, 0x06, 0x91, 0xc1, 0x8b, 0xa1, 0xa7, 0x1e, 0x24,
	0x82, 0x5e, 0xb4, 0x67, 0x05, 0xbd, 0x49, 0xe8, 0x17, 0x58, 0x93, 0xc5, 0x86, 0xd4, 0x64, 0xd9,
	0x9d, 0x0a, 0xfb, 0x75, 0xfd, 0x24, 0xb2, 0xff, 0x4a, 0xc1, 0xdb, 0xfb, 0xbd, 0x7d, 0x33, 0xfb,
	0x06, 0x48, 0xe9, 0x89, 0xa7, 0xfb, 0x97, 0xde, 0x28, 0xc1, 0xed, 0x7e, 0x27, 0xcc, 0x50, 0x7b,
	0x0b, 0xcb, 0x6f, 0xeb, 0xc5, 0x7a, 0x07, 0xab, 0xf3, 0x67, 0xbc, 0x86, 0x05, 0x0b, 0x33, 0xbc,
	0x77, 0x94, 0x55, 0xd9, 0x66, 0xd9, 0x44, 0xc2, 0x3b, 0x28, 0x3b, 0xc9, 0xa2, 0x3f, 0x18, 0xca,
	0xab, 0x62, 0x73, 0xf1, 0x80, 0x75, 0x5c, 0x51, 0x7f, 0x1c, 0xf5, 0x97, 0x74, 0xc3, 0x4d, 0x8a,
	0xac, 0x7f, 0x33, 0x58, 0x9e, 0x6c, 0x24, 0x28, 0x95, 0x83, 0xd3, 0xd2, 0x84, 0x88, 0x30, 0x63,
	0xab, 0x24, 0xe5, 0xde, 0xf6, 0xda, 0xa5, 0xdb, 0x69, 0x64, 0x39, 0x32, 0x15, 0x55, 0xe1, 0xd2,
	0x11, 0xf1, 0x19, 0xca, 0xbd, 0x14, 0x9d, 0xd4, 0x86, 0x66, 0xbe, 0xc3, 0xed, 0xff, 0x0e, 0xf5,
	0x5b, 0x48, 0xbc, 0x8e, 0xac, 0x6d, 0x93, 0xf2, 0xee, 0x2c, 0xd1, 0x72, 0x3f, 0x8d, 0x34, 0x0f,
	0x67, 0x05, 0xba, 0xd9, 0xc2, 0xea, 0x7c, 0x00, 0x2f, 0xa1, 0x18, 0xa4, 0x8d, 0x35, 0x9d, 0xc4,
	0x2b, 0x98, 0xff, 0x88, 0xc3, 0x31, 0x75, 0x0c, 0xb0, 0xcd, 0x9f, 0xb2, 0xcf, 0x85, 0xff, 0xfa,
	0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x20, 0x41, 0x47, 0x66, 0x01, 0x00, 0x00,
}
