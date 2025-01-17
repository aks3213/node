// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/keygen.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type KeygenStatus int32

const (
	KeygenStatus_PendingKeygen KeygenStatus = 0
	KeygenStatus_KeyGenSuccess KeygenStatus = 1
	KeygenStatus_KeyGenFailed  KeygenStatus = 3
)

var KeygenStatus_name = map[int32]string{
	0: "PendingKeygen",
	1: "KeyGenSuccess",
	3: "KeyGenFailed",
}

var KeygenStatus_value = map[string]int32{
	"PendingKeygen": 0,
	"KeyGenSuccess": 1,
	"KeyGenFailed":  3,
}

func (x KeygenStatus) String() string {
	return proto.EnumName(KeygenStatus_name, int32(x))
}

func (KeygenStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4efb2de738775c96, []int{0}
}

type Keygen struct {
	Status         KeygenStatus `protobuf:"varint,1,opt,name=status,proto3,enum=zetachain.zetacore.observer.KeygenStatus" json:"status,omitempty"`
	GranteePubkeys []string     `protobuf:"bytes,2,rep,name=granteePubkeys,proto3" json:"granteePubkeys,omitempty"`
	BlockNumber    int64        `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
}

func (m *Keygen) Reset()         { *m = Keygen{} }
func (m *Keygen) String() string { return proto.CompactTextString(m) }
func (*Keygen) ProtoMessage()    {}
func (*Keygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_4efb2de738775c96, []int{0}
}
func (m *Keygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keygen.Merge(m, src)
}
func (m *Keygen) XXX_Size() int {
	return m.Size()
}
func (m *Keygen) XXX_DiscardUnknown() {
	xxx_messageInfo_Keygen.DiscardUnknown(m)
}

var xxx_messageInfo_Keygen proto.InternalMessageInfo

func (m *Keygen) GetStatus() KeygenStatus {
	if m != nil {
		return m.Status
	}
	return KeygenStatus_PendingKeygen
}

func (m *Keygen) GetGranteePubkeys() []string {
	if m != nil {
		return m.GranteePubkeys
	}
	return nil
}

func (m *Keygen) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.observer.KeygenStatus", KeygenStatus_name, KeygenStatus_value)
	proto.RegisterType((*Keygen)(nil), "zetachain.zetacore.observer.Keygen")
}

func init() { proto.RegisterFile("observer/keygen.proto", fileDescriptor_4efb2de738775c96) }

var fileDescriptor_4efb2de738775c96 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4f, 0x2a, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0xcf, 0x4e, 0xad, 0x4c, 0x4f, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xae, 0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2,
	0x8b, 0x52, 0xf5, 0x60, 0x2a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xea, 0xf4, 0x41, 0x2c,
	0x88, 0x16, 0xa5, 0xa9, 0x8c, 0x5c, 0x6c, 0xde, 0x60, 0x33, 0x84, 0x1c, 0xb9, 0xd8, 0x8a, 0x4b,
	0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x34, 0xf5, 0xf0, 0x18, 0xa7,
	0x07, 0xd1, 0x14, 0x0c, 0xd6, 0x10, 0x04, 0xd5, 0x28, 0xa4, 0xc6, 0xc5, 0x97, 0x5e, 0x94, 0x98,
	0x57, 0x92, 0x9a, 0x1a, 0x50, 0x9a, 0x94, 0x9d, 0x5a, 0x59, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1,
	0x19, 0x84, 0x26, 0x2a, 0xa4, 0xc0, 0xc5, 0x9d, 0x94, 0x93, 0x9f, 0x9c, 0xed, 0x57, 0x9a, 0x9b,
	0x94, 0x5a, 0x24, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x2c, 0xa4, 0xe5, 0xc3, 0xc5, 0x83,
	0x6c, 0x83, 0x90, 0x20, 0x17, 0x6f, 0x40, 0x6a, 0x5e, 0x4a, 0x66, 0x5e, 0x3a, 0x44, 0x58, 0x80,
	0x01, 0x24, 0xe4, 0x9d, 0x5a, 0xe9, 0x9e, 0x9a, 0x17, 0x5c, 0x9a, 0x9c, 0x9c, 0x5a, 0x5c, 0x2c,
	0xc0, 0x28, 0x24, 0x00, 0xd6, 0xe5, 0x9e, 0x9a, 0xe7, 0x96, 0x98, 0x99, 0x93, 0x9a, 0x22, 0xc0,
	0x2c, 0xc5, 0xb2, 0x62, 0x89, 0x1c, 0xa3, 0x93, 0xe7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83,
	0x3c, 0xa9, 0x0b, 0xf6, 0xaf, 0x3e, 0xcc, 0xbf, 0xfa, 0x15, 0xfa, 0xf0, 0xa0, 0x2e, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xac, 0x1f,
	0xb4, 0x83, 0x01, 0x00, 0x00,
}

func (m *Keygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.GranteePubkeys) > 0 {
		for iNdEx := len(m.GranteePubkeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GranteePubkeys[iNdEx])
			copy(dAtA[i:], m.GranteePubkeys[iNdEx])
			i = encodeVarintKeygen(dAtA, i, uint64(len(m.GranteePubkeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Status != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeygen(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeygen(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Keygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovKeygen(uint64(m.Status))
	}
	if len(m.GranteePubkeys) > 0 {
		for _, s := range m.GranteePubkeys {
			l = len(s)
			n += 1 + l + sovKeygen(uint64(l))
		}
	}
	if m.BlockNumber != 0 {
		n += 1 + sovKeygen(uint64(m.BlockNumber))
	}
	return n
}

func sovKeygen(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeygen(x uint64) (n int) {
	return sovKeygen(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Keygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: Keygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= KeygenStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GranteePubkeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GranteePubkeys = append(m.GranteePubkeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func skipKeygen(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
				return 0, ErrInvalidLengthKeygen
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeygen
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeygen
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeygen        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeygen          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeygen = fmt.Errorf("proto: unexpected end of group")
)
