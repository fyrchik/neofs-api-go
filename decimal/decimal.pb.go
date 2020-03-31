// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/decimal.proto

package decimal

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Decimal is a structure used for representation of assets amount
type Decimal struct {
	// Value is value number
	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	// Precision is precision number
	Precision            uint32   `protobuf:"varint,2,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal) Reset()      { *m = Decimal{} }
func (*Decimal) ProtoMessage() {}
func (*Decimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7e70e1773836c80, []int{0}
}
func (m *Decimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Decimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Decimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal.Merge(m, src)
}
func (m *Decimal) XXX_Size() int {
	return m.Size()
}
func (m *Decimal) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal proto.InternalMessageInfo

func (m *Decimal) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Decimal) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func init() {
	proto.RegisterType((*Decimal)(nil), "decimal.Decimal")
}

func init() { proto.RegisterFile("decimal/decimal.proto", fileDescriptor_e7e70e1773836c80) }

var fileDescriptor_e7e70e1773836c80 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x49, 0x4d, 0xce,
	0xcc, 0x4d, 0xcc, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a,
	0xbe, 0x3e, 0x58, 0x3e, 0xa9, 0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0x88, 0x3e, 0x25, 0x67,
	0x2e, 0x76, 0x17, 0x88, 0x4e, 0x21, 0x11, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x48, 0x86, 0x8b, 0x33, 0xa0, 0x28, 0x35, 0x39, 0xb3,
	0x38, 0x33, 0x3f, 0x4f, 0x82, 0x49, 0x81, 0x51, 0x83, 0x37, 0x08, 0x21, 0x60, 0xc5, 0x32, 0x63,
	0x81, 0x3c, 0x83, 0x53, 0xf0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x9a, 0x48, 0x2e, 0xc9, 0x2b,
	0x2e, 0x48, 0x4e, 0xd6, 0x4d, 0x49, 0x2d, 0xd3, 0xcf, 0x4b, 0xcd, 0x4f, 0x2b, 0xd6, 0x4d, 0x2c,
	0xc8, 0xd4, 0x4d, 0xcf, 0x87, 0xf9, 0x61, 0x15, 0x93, 0xa0, 0x5f, 0x6a, 0xbe, 0x5b, 0xb0, 0x9e,
	0x63, 0x80, 0xa7, 0x1e, 0xd4, 0x39, 0x49, 0x6c, 0x60, 0x07, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x41, 0x4c, 0x9f, 0xf1, 0x00, 0x00, 0x00,
}

func (m *Decimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Decimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Decimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Precision != 0 {
		i = encodeVarintDecimal(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i = encodeVarintDecimal(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDecimal(dAtA []byte, offset int, v uint64) int {
	offset -= sovDecimal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Decimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovDecimal(uint64(m.Value))
	}
	if m.Precision != 0 {
		n += 1 + sovDecimal(uint64(m.Precision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDecimal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDecimal(x uint64) (n int) {
	return sovDecimal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Decimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDecimal
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
			return fmt.Errorf("proto: Decimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Decimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecimal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecimal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDecimal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDecimal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDecimal
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
func skipDecimal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDecimal
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
					return 0, ErrIntOverflowDecimal
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
					return 0, ErrIntOverflowDecimal
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
				return 0, ErrInvalidLengthDecimal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDecimal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDecimal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDecimal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDecimal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDecimal = fmt.Errorf("proto: unexpected end of group")
)
