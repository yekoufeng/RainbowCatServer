// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue260.proto

/*
	Package issue260 is a generated protocol buffer package.

	It is generated from these files:
		issue260.proto

	It has these top-level messages:
		Dropped
		DroppedWithoutGetters
		Kept
*/
package issue260

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Dropped) Reset()                    { *m = Dropped{} }
func (m *Dropped) String() string            { return proto.CompactTextString(m) }
func (*Dropped) ProtoMessage()               {}
func (*Dropped) Descriptor() ([]byte, []int) { return fileDescriptorIssue260, []int{0} }

func (m *Dropped) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dropped) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *DroppedWithoutGetters) Reset()                    { *m = DroppedWithoutGetters{} }
func (m *DroppedWithoutGetters) String() string            { return proto.CompactTextString(m) }
func (*DroppedWithoutGetters) ProtoMessage()               {}
func (*DroppedWithoutGetters) Descriptor() ([]byte, []int) { return fileDescriptorIssue260, []int{1} }

type Kept struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (m *Kept) Reset()                    { *m = Kept{} }
func (m *Kept) String() string            { return proto.CompactTextString(m) }
func (*Kept) ProtoMessage()               {}
func (*Kept) Descriptor() ([]byte, []int) { return fileDescriptorIssue260, []int{2} }

func (m *Kept) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Kept) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Dropped)(nil), "issue260.Dropped")
	proto.RegisterType((*DroppedWithoutGetters)(nil), "issue260.DroppedWithoutGetters")
	proto.RegisterType((*Kept)(nil), "issue260.Kept")
}
func (this *Dropped) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Dropped)
	if !ok {
		that2, ok := that.(Dropped)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Dropped")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Dropped but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Dropped but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Age != that1.Age {
		return fmt.Errorf("Age this(%v) Not Equal that(%v)", this.Age, that1.Age)
	}
	return nil
}
func (this *Dropped) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Dropped)
	if !ok {
		that2, ok := that.(Dropped)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	return true
}
func (this *DroppedWithoutGetters) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*DroppedWithoutGetters)
	if !ok {
		that2, ok := that.(DroppedWithoutGetters)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *DroppedWithoutGetters")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *DroppedWithoutGetters but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *DroppedWithoutGetters but is not nil && this == nil")
	}
	if this.Height != that1.Height {
		return fmt.Errorf("Height this(%v) Not Equal that(%v)", this.Height, that1.Height)
	}
	if this.Width != that1.Width {
		return fmt.Errorf("Width this(%v) Not Equal that(%v)", this.Width, that1.Width)
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	return nil
}
func (this *DroppedWithoutGetters) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DroppedWithoutGetters)
	if !ok {
		that2, ok := that.(DroppedWithoutGetters)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Width != that1.Width {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	return true
}
func (this *Kept) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Kept)
	if !ok {
		that2, ok := that.(Kept)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Kept")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Kept but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Kept but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Age != that1.Age {
		return fmt.Errorf("Age this(%v) Not Equal that(%v)", this.Age, that1.Age)
	}
	return nil
}
func (this *Kept) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Kept)
	if !ok {
		that2, ok := that.(Kept)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	return true
}
func (m *Dropped) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dropped) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Age != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(m.Age))
	}
	return i, nil
}

func (m *DroppedWithoutGetters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DroppedWithoutGetters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(m.Height))
	}
	if m.Width != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(m.Width))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintIssue260(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Kept) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Kept) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Age != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIssue260(dAtA, i, uint64(m.Age))
	}
	return i, nil
}

func encodeFixed64Issue260(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Issue260(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIssue260(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedDropped(r randyIssue260, easy bool) *Dropped {
	this := &Dropped{}
	this.Name = string(randStringIssue260(r))
	this.Age = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Age *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedDroppedWithoutGetters(r randyIssue260, easy bool) *DroppedWithoutGetters {
	this := &DroppedWithoutGetters{}
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	this.Width = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Width *= -1
	}
	v1 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.Timestamp = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedKept(r randyIssue260, easy bool) *Kept {
	this := &Kept{}
	this.Name = string(randStringIssue260(r))
	this.Age = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Age *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyIssue260 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue260(r randyIssue260) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue260(r randyIssue260) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneIssue260(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue260(r randyIssue260, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue260(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue260(dAtA []byte, r randyIssue260, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue260(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue260(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Dropped) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIssue260(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovIssue260(uint64(m.Age))
	}
	return n
}

func (m *DroppedWithoutGetters) Size() (n int) {
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovIssue260(uint64(m.Height))
	}
	if m.Width != 0 {
		n += 1 + sovIssue260(uint64(m.Width))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovIssue260(uint64(l))
	return n
}

func (m *Kept) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIssue260(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovIssue260(uint64(m.Age))
	}
	return n
}

func sovIssue260(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue260(x uint64) (n int) {
	return sovIssue260(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Dropped) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue260
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
			return fmt.Errorf("proto: Dropped: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dropped: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
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
				return ErrInvalidLengthIssue260
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIssue260(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue260
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
func (m *DroppedWithoutGetters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue260
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
			return fmt.Errorf("proto: DroppedWithoutGetters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DroppedWithoutGetters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
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
				return ErrInvalidLengthIssue260
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssue260(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue260
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
func (m *Kept) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue260
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
			return fmt.Errorf("proto: Kept: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Kept: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
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
				return ErrInvalidLengthIssue260
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue260
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIssue260(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue260
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
func skipIssue260(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue260
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
					return 0, ErrIntOverflowIssue260
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
					return 0, ErrIntOverflowIssue260
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
				return 0, ErrInvalidLengthIssue260
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue260
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
				next, err := skipIssue260(dAtA[start:])
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
	ErrInvalidLengthIssue260 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue260   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue260.proto", fileDescriptorIssue260) }

var fileDescriptorIssue260 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0xf3, 0x99, 0xb6, 0xb6, 0x27, 0x88, 0x1c, 0x2a, 0x25, 0xc3, 0x25, 0x74, 0xca, 0xa0,
	0xa9, 0x54, 0x74, 0xe8, 0x18, 0x04, 0x07, 0xb7, 0x20, 0x38, 0x27, 0xf6, 0xbc, 0x1c, 0x18, 0x2f,
	0x24, 0x5f, 0x70, 0x75, 0x74, 0x14, 0xfc, 0x03, 0xba, 0xf9, 0x13, 0x1c, 0x1d, 0x3b, 0xfa, 0x0b,
	0xb4, 0x3d, 0xff, 0x80, 0x63, 0x47, 0xf1, 0xd2, 0xd8, 0xd5, 0xed, 0x7d, 0xe0, 0x7d, 0x3f, 0x9e,
	0x8f, 0x6c, 0xca, 0xb2, 0xac, 0xf8, 0xe8, 0xf8, 0x20, 0xc8, 0x0b, 0x85, 0x8a, 0x76, 0x1b, 0x76,
	0xf6, 0x85, 0xc4, 0xb4, 0x4a, 0x82, 0x4b, 0x95, 0x0d, 0x85, 0x12, 0x6a, 0x68, 0x0a, 0x49, 0x75,
	0x65, 0xc8, 0x80, 0x49, 0xf5, 0xd0, 0x71, 0x85, 0x52, 0xe2, 0x9a, 0xaf, 0x5a, 0x28, 0x33, 0x5e,
	0x62, 0x9c, 0xe5, 0x75, 0x61, 0x70, 0x44, 0xd6, 0x4f, 0x0a, 0x95, 0xe7, 0x7c, 0x42, 0x29, 0x69,
	0xdd, 0xc4, 0x19, 0xef, 0x83, 0x07, 0x7e, 0x2f, 0x32, 0x99, 0x6e, 0x11, 0x3b, 0x16, 0xbc, 0xbf,
	0xe6, 0x81, 0xdf, 0x8e, 0x7e, 0xe3, 0xb8, 0xf5, 0xfd, 0xec, 0x5a, 0x83, 0x47, 0x20, 0x3b, 0xcb,
	0xdd, 0x85, 0xc4, 0x54, 0x55, 0x78, 0xca, 0x11, 0x79, 0x51, 0xd2, 0x5d, 0xd2, 0x49, 0xb9, 0x14,
	0x29, 0x9a, 0x3b, 0x76, 0xb4, 0x24, 0xba, 0x4d, 0xda, 0xb7, 0x72, 0x82, 0xa9, 0xb9, 0x65, 0x47,
	0x35, 0xd0, 0x90, 0xf4, 0xfe, 0x8c, 0xfa, 0xb6, 0x07, 0xfe, 0xc6, 0xc8, 0x09, 0x6a, 0xe7, 0xa0,
	0x71, 0x0e, 0xce, 0x9b, 0x46, 0xd8, 0x9d, 0x7e, 0xb8, 0xd6, 0xc3, 0xa7, 0x0b, 0xd1, 0x6a, 0x36,
	0xee, 0xde, 0x3f, 0xb9, 0x96, 0xb1, 0xda, 0x23, 0xad, 0x33, 0x9e, 0xe3, 0xff, 0x3e, 0x09, 0xfd,
	0xd9, 0x9c, 0xc1, 0x62, 0xce, 0xe0, 0x45, 0x33, 0x78, 0xd5, 0x0c, 0xde, 0x34, 0x83, 0xa9, 0x66,
	0xf0, 0xae, 0x19, 0xcc, 0x34, 0x83, 0x85, 0x66, 0xd6, 0xdd, 0x17, 0xb3, 0x92, 0x8e, 0x51, 0x39,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x67, 0x75, 0x8b, 0x97, 0x01, 0x00, 0x00,
}
