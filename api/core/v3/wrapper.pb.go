// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v3/wrapper.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	v2 "github.com/sensu/sensu-go/api/core/v2"
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

// Wrapper represents a wrapped resource
type Wrapper struct {
	// TypeMeta contains the type and the API version of the resource
	*v2.TypeMeta `protobuf:"bytes,1,opt,name=typemeta,proto3,embedded=typemeta" json:""`
	// ObjectMeta contains the name, namespace, labels and annotations of the
	// resource
	Metadata *v2.ObjectMeta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Value contains the encoded resource value
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wrapper) Reset()         { *m = Wrapper{} }
func (m *Wrapper) String() string { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()    {}
func (*Wrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4e62a25c4eb29c4, []int{0}
}
func (m *Wrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wrapper.Merge(m, src)
}
func (m *Wrapper) XXX_Size() int {
	return m.Size()
}
func (m *Wrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_Wrapper.DiscardUnknown(m)
}

var xxx_messageInfo_Wrapper proto.InternalMessageInfo

func (m *Wrapper) GetMetadata() *v2.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Wrapper) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Wrapper)(nil), "sensu.core.v3.Wrapper")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v3/wrapper.proto", fileDescriptor_e4e62a25c4eb29c4)
}

var fileDescriptor_e4e62a25c4eb29c4 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x4e, 0xcd, 0x2b, 0x2e, 0x85, 0x90, 0xba, 0xe9,
	0xf9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x65, 0xc6, 0xfa, 0xe5, 0x45,
	0x89, 0x05, 0x05, 0xa9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x60, 0x35, 0x7a,
	0x20, 0x49, 0xbd, 0x32, 0x63, 0x29, 0x13, 0x24, 0x33, 0xd2, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xaa,
	0x92, 0x4a, 0xd3, 0x1c, 0xca, 0x0c, 0xf5, 0x8c, 0xf5, 0x0c, 0xc1, 0x82, 0x60, 0x31, 0x30, 0x0b,
	0x62, 0x88, 0x94, 0x01, 0x31, 0x36, 0x1b, 0xe9, 0xe7, 0xa6, 0x96, 0x24, 0x42, 0x74, 0x28, 0xad,
	0x66, 0xe4, 0x62, 0x0f, 0x87, 0x38, 0x44, 0xc8, 0x91, 0x8b, 0xa3, 0xa4, 0xb2, 0x20, 0x15, 0x24,
	0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xae, 0x87, 0xec, 0x2a, 0x23, 0xbd, 0x90, 0xca,
	0x82, 0x54, 0xdf, 0xd4, 0x92, 0x44, 0x27, 0x9e, 0x13, 0xf7, 0xe4, 0x19, 0x2f, 0xdc, 0x93, 0x67,
	0x7c, 0x75, 0x4f, 0x9e, 0x21, 0x08, 0xae, 0x4d, 0xc8, 0x9d, 0x8b, 0x03, 0x44, 0xa7, 0x24, 0x96,
	0x24, 0x4a, 0x30, 0x81, 0x8d, 0x90, 0x44, 0x33, 0xc2, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0x04, 0x6c,
	0x88, 0xc0, 0x09, 0x88, 0x01, 0x70, 0x2d, 0x41, 0x70, 0x96, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62,
	0x4e, 0x69, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84, 0xe3, 0xa4, 0xf0, 0xe3, 0xa1,
	0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x31, 0x95, 0x19, 0x27,
	0xb1, 0x81, 0xbd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x36, 0xc2, 0x94, 0x84, 0x01,
	0x00, 0x00,
}

func (this *Wrapper) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Wrapper)
	if !ok {
		that2, ok := that.(Wrapper)
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
	if !this.TypeMeta.Equal(that1.TypeMeta) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Wrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintWrapper(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWrapper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TypeMeta != nil {
		{
			size, err := m.TypeMeta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWrapper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWrapper(dAtA []byte, offset int, v uint64) int {
	offset -= sovWrapper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedWrapper(r randyWrapper, easy bool) *Wrapper {
	this := &Wrapper{}
	if r.Intn(5) != 0 {
		this.TypeMeta = v2.NewPopulatedTypeMeta(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Metadata = v2.NewPopulatedObjectMeta(r, easy)
	}
	v1 := r.Intn(100)
	this.Value = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedWrapper(r, 4)
	}
	return this
}

type randyWrapper interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneWrapper(r randyWrapper) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringWrapper(r randyWrapper) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneWrapper(r)
	}
	return string(tmps)
}
func randUnrecognizedWrapper(r randyWrapper, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldWrapper(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldWrapper(dAtA []byte, r randyWrapper, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateWrapper(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateWrapper(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Wrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypeMeta != nil {
		l = m.TypeMeta.Size()
		n += 1 + l + sovWrapper(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovWrapper(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWrapper(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWrapper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWrapper(x uint64) (n int) {
	return sovWrapper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWrapper
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
			return fmt.Errorf("proto: Wrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWrapper
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
				return ErrInvalidLengthWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TypeMeta == nil {
				m.TypeMeta = &v2.TypeMeta{}
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWrapper
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
				return ErrInvalidLengthWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &v2.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWrapper
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
				return ErrInvalidLengthWrapper
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWrapper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWrapper
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWrapper
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
func skipWrapper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWrapper
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
					return 0, ErrIntOverflowWrapper
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
					return 0, ErrIntOverflowWrapper
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
				return 0, ErrInvalidLengthWrapper
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWrapper
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWrapper
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWrapper        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWrapper          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWrapper = fmt.Errorf("proto: unexpected end of group")
)
