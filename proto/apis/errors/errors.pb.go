// Code generated by proto-gen-gogo. DO NOT EDIT.
// source: github.com/lack-io/vine/proto/apis/errors/errors.proto

package errors

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

type Error struct {
	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code     int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Detail   string   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Status   string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Position string   `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
	Child    *Child   `protobuf:"bytes,6,opt,name=child,proto3" json:"child,omitempty"`
	Stacks   []*Stack `protobuf:"bytes,7,rep,name=stacks,proto3" json:"stacks,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c33674bd5d75a1e, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Error) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Error) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Error) GetChild() *Child {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *Error) GetStacks() []*Stack {
	if m != nil {
		return m.Stacks
	}
	return nil
}

type Child struct {
	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c33674bd5d75a1e, []int{1}
}
func (m *Child) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Child.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(m, src)
}
func (m *Child) XXX_Size() int {
	return m.Size()
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Child) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Stack struct {
	Code     int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail   string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Position string `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (m *Stack) Reset()         { *m = Stack{} }
func (m *Stack) String() string { return proto.CompactTextString(m) }
func (*Stack) ProtoMessage()    {}
func (*Stack) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c33674bd5d75a1e, []int{2}
}
func (m *Stack) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stack.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stack.Merge(m, src)
}
func (m *Stack) XXX_Size() int {
	return m.Size()
}
func (m *Stack) XXX_DiscardUnknown() {
	xxx_messageInfo_Stack.DiscardUnknown(m)
}

var xxx_messageInfo_Stack proto.InternalMessageInfo

func (m *Stack) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Stack) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Stack) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "errors.Error")
	proto.RegisterType((*Child)(nil), "errors.Child")
	proto.RegisterType((*Stack)(nil), "errors.Stack")
}

func init() {
	proto.RegisterFile("github.com/lack-io/vine/proto/apis/errors/errors.proto", fileDescriptor_1c33674bd5d75a1e)
}

var fileDescriptor_1c33674bd5d75a1e = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x49, 0x4c, 0xce, 0xd6, 0xcd, 0xcc, 0xd7, 0x2f,
	0xcb, 0xcc, 0x4b, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x4f,
	0x2d, 0x2a, 0xca, 0x2f, 0x82, 0x51, 0x7a, 0x60, 0x09, 0x21, 0x36, 0x08, 0x4f, 0xe9, 0x30, 0x23,
	0x17, 0xab, 0x2b, 0x88, 0x29, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x0b, 0x89, 0x71, 0xb1, 0xa5, 0xa4, 0x96, 0x24, 0x66, 0xe6, 0x48,
	0x30, 0x83, 0xd5, 0x41, 0x79, 0x20, 0xf1, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x16, 0x88,
	0x38, 0x84, 0x27, 0x24, 0xc5, 0xc5, 0x51, 0x90, 0x5f, 0x9c, 0x59, 0x92, 0x99, 0x9f, 0x27, 0xc1,
	0x0a, 0x96, 0x81, 0xf3, 0x85, 0x94, 0xb9, 0x58, 0x93, 0x33, 0x32, 0x73, 0x52, 0x24, 0xd8, 0x14,
	0x18, 0x35, 0xb8, 0x8d, 0x78, 0xf5, 0xa0, 0xee, 0x73, 0x06, 0x09, 0x06, 0x41, 0xe4, 0x84, 0x54,
	0xc1, 0x06, 0x27, 0x67, 0x17, 0x4b, 0xb0, 0x2b, 0x30, 0x23, 0xab, 0x0a, 0x06, 0x89, 0x06, 0x41,
	0x25, 0x95, 0x8c, 0xb9, 0x58, 0xc1, 0xda, 0xe0, 0x8e, 0x66, 0xc4, 0xea, 0x68, 0x26, 0x64, 0x47,
	0x2b, 0xf9, 0x73, 0xb1, 0x82, 0x4d, 0x21, 0x45, 0x13, 0x8a, 0x8f, 0x98, 0x51, 0x7d, 0xe4, 0xe4,
	0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x06, 0x44, 0xc7, 0x92, 0x35,
	0x84, 0x4a, 0x62, 0x03, 0xcb, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x50, 0xaf, 0x98, 0x44,
	0xe0, 0x01, 0x00, 0x00,
}

func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Child != nil {
		l = m.Child.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	if len(m.Stacks) > 0 {
		for _, e := range m.Stacks {
			l = e.Size()
			n += 1 + l + sovErrors(uint64(l))
		}
	}
	return n
}

func (m *Child) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func (m *Stack) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Stacks) > 0 {
		for iNdEx := len(m.Stacks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stacks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintErrors(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Child != nil {
		{
			size, err := m.Child.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintErrors(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Position) > 0 {
		i -= len(m.Position)
		copy(dAtA[i:], m.Position)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Position)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Child) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Child) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Child) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Stack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stack) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stack) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Position) > 0 {
		i -= len(m.Position)
		copy(dAtA[i:], m.Position)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Position)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Child", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Child == nil {
				m.Child = &Child{}
			}
			if err := m.Child.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stacks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stacks = append(m.Stacks, &Stack{})
			if err := m.Stacks[len(m.Stacks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Child) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Child: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Child: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Stack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Stack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
