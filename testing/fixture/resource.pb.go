// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource.proto

package fixture

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import v2 "github.com/sensu/core/v2"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Resource struct {
	Foo                  string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	v2.ObjectMeta        `protobuf:"bytes,2,opt,name=metadata,embedded=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_1b8179e05fc21672, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "testing.fixture.Resource")
}
func (this *Resource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Resource)
	if !ok {
		that2, ok := that.(Resource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Foo != that1.Foo {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Foo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResource(dAtA, i, uint64(len(m.Foo)))
		i += copy(dAtA[i:], m.Foo)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintResource(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedResource(r randyResource, easy bool) *Resource {
	this := &Resource{}
	this.Foo = string(randStringResource(r))
	v1 := v2.NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedResource(r, 3)
	}
	return this
}

type randyResource interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneResource(r randyResource) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringResource(r randyResource) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneResource(r)
	}
	return string(tmps)
}
func randUnrecognizedResource(r randyResource, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldResource(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldResource(dAtA []byte, r randyResource, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateResource(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateResource(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateResource(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateResource(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateResource(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateResource(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateResource(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Resource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = m.ObjectMeta.Size()
	n += 1 + l + sovResource(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResource
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
				next, err := skipResource(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthResource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("resource.proto", fileDescriptor_resource_1b8179e05fc21672) }

var fileDescriptor_resource_1b8179e05fc21672 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x49, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0xd7, 0x4b, 0xcb, 0xac, 0x28, 0x29, 0x2d, 0x4a, 0x95, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4b,
	0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x5f, 0xca, 0x00, 0x49, 0x79, 0x71, 0x6a,
	0x5e, 0x71, 0x29, 0x84, 0xd4, 0x4d, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xce, 0x2f, 0x4a,
	0xd5, 0x2f, 0x33, 0xd2, 0xcf, 0x4d, 0x2d, 0x49, 0x84, 0xe8, 0x50, 0xca, 0xe7, 0xe2, 0x08, 0x82,
	0xba, 0x41, 0x48, 0x8c, 0x8b, 0x39, 0x2d, 0x3f, 0x5f, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89,
	0xe5, 0xc4, 0x3d, 0x79, 0xc6, 0x20, 0x90, 0x80, 0x90, 0x37, 0x17, 0x07, 0x48, 0x47, 0x4a, 0x62,
	0x49, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa4, 0x1e, 0xd8, 0x5c, 0x3d, 0x90, 0x81,
	0x7a, 0x65, 0x46, 0x7a, 0xfe, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0xbe, 0xa9, 0x25, 0x89, 0x4e, 0x22,
	0x27, 0xee, 0xc9, 0x33, 0x5c, 0xb8, 0x27, 0xcf, 0xf8, 0xea, 0x9e, 0x3c, 0x5c, 0x5b, 0x10, 0x9c,
	0xe5, 0xa4, 0xfa, 0xe3, 0xa1, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x9e,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31,
	0x44, 0xb1, 0x43, 0x3d, 0x9e, 0xc4, 0x06, 0x76, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x1e, 0xe9, 0xd4, 0x22, 0x01, 0x00, 0x00,
}
