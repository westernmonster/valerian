// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ecode.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Error struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMessage           string   `protobuf:"bytes,2,opt,name=err_message,json=errMessage,proto3" json:"err_message,omitempty"`
	ErrDetail            *any.Any `protobuf:"bytes,3,opt,name=err_detail,json=errDetail,proto3" json:"err_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_c62007a7b6fefd8a, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Error) GetErrMessage() string {
	if m != nil {
		return m.ErrMessage
	}
	return ""
}

func (m *Error) GetErrDetail() *any.Any {
	if m != nil {
		return m.ErrDetail
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "pb.Error")
}

func init() { proto.RegisterFile("ecode.proto", fileDescriptor_c62007a7b6fefd8a) }

var fileDescriptor_c62007a7b6fefd8a = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xbd, 0x8e, 0x83, 0x30,
	0x10, 0x84, 0x65, 0x4e, 0xdc, 0x1d, 0xa6, 0xb3, 0xae, 0x80, 0x4b, 0x11, 0x94, 0x8a, 0xca, 0x96,
	0xc2, 0x13, 0xe4, 0xaf, 0x4c, 0x43, 0x99, 0x06, 0xd9, 0x61, 0x83, 0x90, 0x1c, 0x8c, 0x06, 0x12,
	0x85, 0xb7, 0x8f, 0x30, 0x4a, 0xb9, 0x33, 0xfb, 0xe9, 0x1b, 0x1e, 0xd3, 0xd5, 0xd5, 0x24, 0x7b,
	0xb8, 0xd1, 0x89, 0xa0, 0x37, 0xff, 0x69, 0xe3, 0x5c, 0x63, 0x49, 0xf9, 0xc4, 0x3c, 0x6e, 0x4a,
	0x77, 0xd3, 0x52, 0x6f, 0x5e, 0x3c, 0x3c, 0x01, 0x0e, 0x22, 0xe5, 0xbf, 0x04, 0x54, 0x33, 0x99,
	0xb0, 0x8c, 0xe5, 0x61, 0xf9, 0x43, 0xc0, 0xc1, 0xd5, 0x24, 0xd6, 0x3c, 0x9e, 0xab, 0x3b, 0x0d,
	0x83, 0x6e, 0x28, 0x09, 0x32, 0x96, 0x47, 0x25, 0x27, 0xe0, 0xbc, 0x24, 0xa2, 0xe0, 0xf3, 0x55,
	0xd5, 0x34, 0xea, 0xd6, 0x26, 0x5f, 0x19, 0xcb, 0xe3, 0xed, 0x9f, 0x5c, 0xa4, 0xf2, 0x23, 0x95,
	0xbb, 0x6e, 0x2a, 0x23, 0x02, 0x8e, 0xfe, 0x6d, 0xbf, 0xba, 0xa4, 0x4f, 0x6d, 0x09, 0xad, 0xee,
	0x94, 0x6d, 0x0d, 0x34, 0x26, 0xe5, 0x87, 0xab, 0xde, 0x98, 0x6f, 0x4f, 0x15, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x88, 0x24, 0x55, 0xcb, 0x00, 0x00, 0x00,
}
