// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/container/types.proto

package container

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap"
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
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

// Container is a structure that defines object placement behaviour. Objects
// can be stored only within containers. They define placement rule, attributes
// and access control information. ID of the container is a 32 byte long
// SHA256 hash of stable-marshalled container message.
type Container struct {
	// OwnerID carries identifier of the container owner.
	OwnerId *refs.OwnerID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Nonce is a 16 byte UUID, used to avoid collisions of container id.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// BasicACL contains access control rules for owner, system, others groups and
	// permission bits for bearer token and Extended ACL.
	BasicAcl uint32 `protobuf:"varint,3,opt,name=basic_acl,json=basicAcl,proto3" json:"basic_acl,omitempty"`
	// Attributes define any immutable characteristics of container.
	Attributes []*Container_Attribute `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Placement policy for the object inside the container.
	PlacementPolicy      *netmap.PlacementPolicy `protobuf:"bytes,5,opt,name=placement_policy,json=placementPolicy,proto3" json:"placement_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_ece71f25e6ae314e, []int{0}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Container.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return m.Size()
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetOwnerId() *refs.OwnerID {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *Container) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Container) GetBasicAcl() uint32 {
	if m != nil {
		return m.BasicAcl
	}
	return 0
}

func (m *Container) GetAttributes() []*Container_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Container) GetPlacementPolicy() *netmap.PlacementPolicy {
	if m != nil {
		return m.PlacementPolicy
	}
	return nil
}

// Attribute is a key-value pair of strings.
type Container_Attribute struct {
	// Key of immutable container attribute.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Value of immutable container attribute.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container_Attribute) Reset()         { *m = Container_Attribute{} }
func (m *Container_Attribute) String() string { return proto.CompactTextString(m) }
func (*Container_Attribute) ProtoMessage()    {}
func (*Container_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_ece71f25e6ae314e, []int{0, 0}
}
func (m *Container_Attribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Container_Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Container_Attribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Container_Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container_Attribute.Merge(m, src)
}
func (m *Container_Attribute) XXX_Size() int {
	return m.Size()
}
func (m *Container_Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Container_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Container_Attribute proto.InternalMessageInfo

func (m *Container_Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Container_Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Container)(nil), "neo.fs.v2.container.Container")
	proto.RegisterType((*Container_Attribute)(nil), "neo.fs.v2.container.Container.Attribute")
}

func init() { proto.RegisterFile("v2/container/types.proto", fileDescriptor_ece71f25e6ae314e) }

var fileDescriptor_ece71f25e6ae314e = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x36, 0xad, 0xd5, 0x66, 0xab, 0x58, 0x52, 0x7f, 0x42, 0x85, 0x10, 0x3d, 0xe5, 0xd2, 0x0d,
	0xa4, 0x17, 0xc1, 0x53, 0x55, 0xc4, 0x82, 0x68, 0x89, 0x37, 0x2f, 0x65, 0xb3, 0x9d, 0xd6, 0xc5,
	0x74, 0x77, 0xc9, 0x6e, 0x23, 0x7d, 0x13, 0x9f, 0xc1, 0x97, 0xf0, 0xea, 0xd1, 0x47, 0x90, 0xfa,
	0x22, 0x92, 0x84, 0xc6, 0x08, 0xde, 0x66, 0xbe, 0xfd, 0xbe, 0x9d, 0xef, 0x9b, 0x41, 0x76, 0x1a,
	0xf8, 0x54, 0x70, 0x4d, 0x18, 0x87, 0xc4, 0xd7, 0x4b, 0x09, 0x0a, 0xcb, 0x44, 0x68, 0x61, 0x75,
	0x38, 0x08, 0x3c, 0x55, 0x38, 0x0d, 0x70, 0x49, 0xe8, 0x1e, 0xa4, 0x81, 0xcf, 0x41, 0xcf, 0x89,
	0xac, 0x72, 0xbb, 0x9d, 0x34, 0xf0, 0x13, 0x98, 0xaa, 0x2a, 0x78, 0xfa, 0x5e, 0x43, 0xe6, 0xe5,
	0x5a, 0x69, 0x05, 0xa8, 0x29, 0x5e, 0x38, 0x24, 0x63, 0x36, 0xb1, 0x0d, 0xd7, 0xf0, 0x5a, 0xc1,
	0x11, 0xfe, 0x9d, 0x90, 0x89, 0xf1, 0x7d, 0xf6, 0x3e, 0xbc, 0x0a, 0xb7, 0x73, 0xe2, 0x70, 0x62,
	0xed, 0xa3, 0x06, 0x17, 0x9c, 0x82, 0x5d, 0x73, 0x0d, 0x6f, 0x27, 0x2c, 0x1a, 0xeb, 0x18, 0x99,
	0x11, 0x51, 0x8c, 0x8e, 0x09, 0x8d, 0xed, 0xba, 0x6b, 0x78, 0xbb, 0x61, 0x33, 0x07, 0x06, 0x34,
	0xb6, 0x6e, 0x10, 0x22, 0x5a, 0x27, 0x2c, 0x5a, 0x68, 0x50, 0xf6, 0xa6, 0x5b, 0xf7, 0x5a, 0x81,
	0x87, 0xff, 0x89, 0x82, 0x4b, 0x6b, 0x78, 0xb0, 0x16, 0x84, 0x15, 0xad, 0x75, 0x8b, 0xda, 0x32,
	0x26, 0x14, 0xe6, 0xc0, 0xf5, 0x58, 0x8a, 0x98, 0xd1, 0xa5, 0xdd, 0xc8, 0x8d, 0x9f, 0x54, 0xfe,
	0x2b, 0x96, 0x81, 0x47, 0x6b, 0xe6, 0x28, 0x27, 0x86, 0x7b, 0xf2, 0x2f, 0xd0, 0xed, 0x23, 0xb3,
	0x1c, 0x63, 0xb5, 0x51, 0xfd, 0x19, 0x96, 0xf9, 0x1a, 0xcc, 0x30, 0x2b, 0xb3, 0xa4, 0x29, 0x89,
	0x17, 0x45, 0x52, 0x33, 0x2c, 0x9a, 0x0b, 0xfa, 0xb1, 0x72, 0x8c, 0xcf, 0x95, 0x63, 0x7c, 0xad,
	0x1c, 0xe3, 0xf5, 0xdb, 0xd9, 0x78, 0x3c, 0x9b, 0x31, 0xfd, 0xb4, 0x88, 0x30, 0x15, 0x73, 0x9f,
	0x2b, 0x49, 0x69, 0x6f, 0x02, 0xa9, 0xcf, 0x41, 0x4c, 0x55, 0x8f, 0x48, 0xd6, 0x9b, 0x09, 0xbf,
	0x7a, 0xd0, 0xf3, 0xb2, 0x7a, 0xab, 0x1d, 0xde, 0x81, 0xb8, 0x7e, 0xc0, 0x83, 0xd1, 0x30, 0xf3,
	0x5d, 0xa6, 0x8f, 0xb6, 0xf2, 0x6b, 0xf5, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xf7, 0x99,
	0x45, 0x0a, 0x02, 0x00, 0x00,
}

func (m *Container) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Container) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Container) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PlacementPolicy != nil {
		{
			size, err := m.PlacementPolicy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BasicAcl != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BasicAcl))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x12
	}
	if m.OwnerId != nil {
		{
			size, err := m.OwnerId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Container_Attribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Container_Attribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Container_Attribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Container) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnerId != nil {
		l = m.OwnerId.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BasicAcl != 0 {
		n += 1 + sovTypes(uint64(m.BasicAcl))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.PlacementPolicy != nil {
		l = m.PlacementPolicy.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Container_Attribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Container) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Container: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Container: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerId == nil {
				m.OwnerId = &refs.OwnerID{}
			}
			if err := m.OwnerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasicAcl", wireType)
			}
			m.BasicAcl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasicAcl |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, &Container_Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlacementPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlacementPolicy == nil {
				m.PlacementPolicy = &netmap.PlacementPolicy{}
			}
			if err := m.PlacementPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Container_Attribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Attribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)