// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EmptyStruct 空的message，对应真实service只返回error，没有具体返回值
type EmptyStruct struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyStruct) Reset()         { *m = EmptyStruct{} }
func (m *EmptyStruct) String() string { return proto.CompactTextString(m) }
func (*EmptyStruct) ProtoMessage()    {}
func (*EmptyStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}
func (m *EmptyStruct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyStruct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmptyStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyStruct.Merge(m, src)
}
func (m *EmptyStruct) XXX_Size() int {
	return m.Size()
}
func (m *EmptyStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyStruct.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyStruct proto.InternalMessageInfo

type FavInfo struct {
	IsFav                bool     `protobuf:"varint,1,opt,name=IsFav,proto3" json:"is_fav"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavInfo) Reset()         { *m = FavInfo{} }
func (m *FavInfo) String() string { return proto.CompactTextString(m) }
func (*FavInfo) ProtoMessage()    {}
func (*FavInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}
func (m *FavInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FavInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FavInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FavInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavInfo.Merge(m, src)
}
func (m *FavInfo) XXX_Size() int {
	return m.Size()
}
func (m *FavInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FavInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FavInfo proto.InternalMessageInfo

func (m *FavInfo) GetIsFav() bool {
	if m != nil {
		return m.IsFav
	}
	return false
}

type FavReq struct {
	AccountID            int64    `protobuf:"varint,1,opt,name=AccountID,proto3" json:"account_id"`
	TargetID             int64    `protobuf:"varint,2,opt,name=TargetID,proto3" json:"target_id"`
	TargetType           string   `protobuf:"bytes,3,opt,name=TargetType,proto3" json:"target_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavReq) Reset()         { *m = FavReq{} }
func (m *FavReq) String() string { return proto.CompactTextString(m) }
func (*FavReq) ProtoMessage()    {}
func (*FavReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}
func (m *FavReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FavReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FavReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FavReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavReq.Merge(m, src)
}
func (m *FavReq) XXX_Size() int {
	return m.Size()
}
func (m *FavReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FavReq.DiscardUnknown(m)
}

var xxx_messageInfo_FavReq proto.InternalMessageInfo

func (m *FavReq) GetAccountID() int64 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *FavReq) GetTargetID() int64 {
	if m != nil {
		return m.TargetID
	}
	return 0
}

func (m *FavReq) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyStruct)(nil), "service.fav.EmptyStruct")
	proto.RegisterType((*FavInfo)(nil), "service.fav.FavInfo")
	proto.RegisterType((*FavReq)(nil), "service.fav.FavReq")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x4b,
	0x2c, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x57,
	0x89, 0x97, 0x8b, 0xdb, 0x35, 0xb7, 0xa0, 0xa4, 0x32, 0xb8, 0xa4, 0xa8, 0x34, 0xb9, 0x44, 0x49,
	0x9b, 0x8b, 0xdd, 0x2d, 0xb1, 0xcc, 0x33, 0x2f, 0x2d, 0x5f, 0x48, 0x81, 0x8b, 0xd5, 0xb3, 0xd8,
	0x2d, 0xb1, 0x4c, 0x82, 0x51, 0x81, 0x51, 0x83, 0xc3, 0x89, 0xeb, 0xd5, 0x3d, 0x79, 0xb6, 0xcc,
	0xe2, 0xf8, 0xb4, 0xc4, 0xb2, 0x20, 0x88, 0x84, 0xd2, 0x24, 0x46, 0x2e, 0x36, 0xb7, 0xc4, 0xb2,
	0xa0, 0xd4, 0x42, 0x21, 0x1d, 0x2e, 0x4e, 0xc7, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x4f, 0x17,
	0xb0, 0x06, 0x66, 0x27, 0xbe, 0x57, 0xf7, 0xe4, 0xb9, 0x12, 0x21, 0x82, 0xf1, 0x99, 0x29, 0x41,
	0x08, 0x05, 0x42, 0x9a, 0x5c, 0x1c, 0x21, 0x89, 0x45, 0xe9, 0xa9, 0x20, 0xc5, 0x4c, 0x60, 0xc5,
	0xbc, 0xaf, 0xee, 0xc9, 0x73, 0x96, 0x80, 0xc5, 0x40, 0x6a, 0xe1, 0xd2, 0x42, 0xfa, 0x5c, 0x5c,
	0x10, 0x76, 0x48, 0x65, 0x41, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xa7, 0x13, 0xff, 0xab, 0x7b,
	0xf2, 0xdc, 0x50, 0xc5, 0x25, 0x95, 0x05, 0xa9, 0x41, 0x48, 0x4a, 0x8c, 0x96, 0x33, 0x72, 0x31,
	0xbb, 0x25, 0x96, 0x09, 0x19, 0x41, 0x9d, 0x2f, 0x24, 0xac, 0x87, 0x14, 0x3c, 0x7a, 0x10, 0xf7,
	0x4a, 0x89, 0xa0, 0x0b, 0x82, 0xbd, 0x6c, 0x02, 0xd1, 0x8a, 0x55, 0x87, 0x04, 0x8a, 0x20, 0x52,
	0x98, 0x09, 0x99, 0x71, 0xb1, 0x86, 0xe6, 0xa5, 0x91, 0xac, 0xcf, 0x49, 0xf0, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x8c, 0x62, 0x4e, 0x2c, 0xc8, 0x4c, 0x62,
	0x03, 0x47, 0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0xe6, 0x94, 0xc3, 0xdc, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FavClient is the client API for Fav service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FavClient interface {
	IsFav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*FavInfo, error)
	Fav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*EmptyStruct, error)
	Unfav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*EmptyStruct, error)
}

type favClient struct {
	cc *grpc.ClientConn
}

func NewFavClient(cc *grpc.ClientConn) FavClient {
	return &favClient{cc}
}

func (c *favClient) IsFav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*FavInfo, error) {
	out := new(FavInfo)
	err := c.cc.Invoke(ctx, "/service.fav.Fav/IsFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favClient) Fav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*EmptyStruct, error) {
	out := new(EmptyStruct)
	err := c.cc.Invoke(ctx, "/service.fav.Fav/Fav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favClient) Unfav(ctx context.Context, in *FavReq, opts ...grpc.CallOption) (*EmptyStruct, error) {
	out := new(EmptyStruct)
	err := c.cc.Invoke(ctx, "/service.fav.Fav/Unfav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavServer is the server API for Fav service.
type FavServer interface {
	IsFav(context.Context, *FavReq) (*FavInfo, error)
	Fav(context.Context, *FavReq) (*EmptyStruct, error)
	Unfav(context.Context, *FavReq) (*EmptyStruct, error)
}

// UnimplementedFavServer can be embedded to have forward compatible implementations.
type UnimplementedFavServer struct {
}

func (*UnimplementedFavServer) IsFav(ctx context.Context, req *FavReq) (*FavInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFav not implemented")
}
func (*UnimplementedFavServer) Fav(ctx context.Context, req *FavReq) (*EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fav not implemented")
}
func (*UnimplementedFavServer) Unfav(ctx context.Context, req *FavReq) (*EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfav not implemented")
}

func RegisterFavServer(s *grpc.Server, srv FavServer) {
	s.RegisterService(&_Fav_serviceDesc, srv)
}

func _Fav_IsFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).IsFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.fav.Fav/IsFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).IsFav(ctx, req.(*FavReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fav_Fav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).Fav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.fav.Fav/Fav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).Fav(ctx, req.(*FavReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fav_Unfav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).Unfav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.fav.Fav/Unfav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).Unfav(ctx, req.(*FavReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fav_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.fav.Fav",
	HandlerType: (*FavServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsFav",
			Handler:    _Fav_IsFav_Handler,
		},
		{
			MethodName: "Fav",
			Handler:    _Fav_Fav_Handler,
		},
		{
			MethodName: "Unfav",
			Handler:    _Fav_Unfav_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func (m *EmptyStruct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyStruct) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmptyStruct) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *FavInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FavInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FavInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsFav {
		i--
		if m.IsFav {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FavReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FavReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FavReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TargetType) > 0 {
		i -= len(m.TargetType)
		copy(dAtA[i:], m.TargetType)
		i = encodeVarintPb(dAtA, i, uint64(len(m.TargetType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TargetID != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.TargetID))
		i--
		dAtA[i] = 0x10
	}
	if m.AccountID != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.AccountID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	offset -= sovPb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EmptyStruct) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FavInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsFav {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FavReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccountID != 0 {
		n += 1 + sovPb(uint64(m.AccountID))
	}
	if m.TargetID != 0 {
		n += 1 + sovPb(uint64(m.TargetID))
	}
	l = len(m.TargetType)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmptyStruct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyStruct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyStruct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FavInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FavInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FavInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFav", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsFav = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FavReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FavReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FavReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			m.AccountID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetID", wireType)
			}
			m.TargetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPb
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPb
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb   = fmt.Errorf("proto: integer overflow")
)
