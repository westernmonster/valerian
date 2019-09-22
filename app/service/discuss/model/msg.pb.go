// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type NotifyDiscussionAdded struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyDiscussionAdded) Reset()         { *m = NotifyDiscussionAdded{} }
func (m *NotifyDiscussionAdded) String() string { return proto.CompactTextString(m) }
func (*NotifyDiscussionAdded) ProtoMessage()    {}
func (*NotifyDiscussionAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}
func (m *NotifyDiscussionAdded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyDiscussionAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyDiscussionAdded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyDiscussionAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyDiscussionAdded.Merge(m, src)
}
func (m *NotifyDiscussionAdded) XXX_Size() int {
	return m.Size()
}
func (m *NotifyDiscussionAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyDiscussionAdded.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyDiscussionAdded proto.InternalMessageInfo

func (m *NotifyDiscussionAdded) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type NotifyDiscussionUpdated struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyDiscussionUpdated) Reset()         { *m = NotifyDiscussionUpdated{} }
func (m *NotifyDiscussionUpdated) String() string { return proto.CompactTextString(m) }
func (*NotifyDiscussionUpdated) ProtoMessage()    {}
func (*NotifyDiscussionUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}
func (m *NotifyDiscussionUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyDiscussionUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyDiscussionUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyDiscussionUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyDiscussionUpdated.Merge(m, src)
}
func (m *NotifyDiscussionUpdated) XXX_Size() int {
	return m.Size()
}
func (m *NotifyDiscussionUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyDiscussionUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyDiscussionUpdated proto.InternalMessageInfo

func (m *NotifyDiscussionUpdated) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type NotifyDiscussionDeleted struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyDiscussionDeleted) Reset()         { *m = NotifyDiscussionDeleted{} }
func (m *NotifyDiscussionDeleted) String() string { return proto.CompactTextString(m) }
func (*NotifyDiscussionDeleted) ProtoMessage()    {}
func (*NotifyDiscussionDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}
func (m *NotifyDiscussionDeleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyDiscussionDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyDiscussionDeleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyDiscussionDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyDiscussionDeleted.Merge(m, src)
}
func (m *NotifyDiscussionDeleted) XXX_Size() int {
	return m.Size()
}
func (m *NotifyDiscussionDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyDiscussionDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyDiscussionDeleted proto.InternalMessageInfo

func (m *NotifyDiscussionDeleted) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*NotifyDiscussionAdded)(nil), "model.NotifyDiscussionAdded")
	proto.RegisterType((*NotifyDiscussionUpdated)(nil), "model.NotifyDiscussionUpdated")
	proto.RegisterType((*NotifyDiscussionDeleted)(nil), "model.NotifyDiscussionDeleted")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb,
	0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5, 0xa4, 0xcf, 0x25, 0xea, 0x97,
	0x5f, 0x92, 0x99, 0x56, 0xe9, 0x92, 0x59, 0x9c, 0x5c, 0x5a, 0x5c, 0x9c, 0x99, 0x9f, 0xe7, 0x98,
	0x92, 0x92, 0x9a, 0x22, 0x24, 0xc6, 0xc5, 0xe4, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xec,
	0xc4, 0xf6, 0xea, 0x9e, 0x3c, 0x53, 0x66, 0x4a, 0x10, 0x93, 0xa7, 0x8b, 0x92, 0x21, 0x97, 0x38,
	0xba, 0x86, 0xd0, 0x82, 0x94, 0xc4, 0x12, 0xd2, 0xb4, 0xb8, 0xa4, 0xe6, 0xa4, 0xe2, 0xd1, 0xe2,
	0xc4, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x26, 0xb1,
	0x81, 0xdd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x65, 0x64, 0x71, 0xc7, 0xee, 0x00, 0x00,
	0x00,
}

func (m *NotifyDiscussionAdded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyDiscussionAdded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyDiscussionAdded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NotifyDiscussionUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyDiscussionUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyDiscussionUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NotifyDiscussionDeleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyDiscussionDeleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyDiscussionDeleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NotifyDiscussionAdded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NotifyDiscussionUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NotifyDiscussionDeleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NotifyDiscussionAdded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: NotifyDiscussionAdded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyDiscussionAdded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *NotifyDiscussionUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: NotifyDiscussionUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyDiscussionUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *NotifyDiscussionDeleted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: NotifyDiscussionDeleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyDiscussionDeleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMsg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsg
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
				next, err := skipMsg(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMsg
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
	ErrInvalidLengthMsg = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg   = fmt.Errorf("proto: integer overflow")
)
