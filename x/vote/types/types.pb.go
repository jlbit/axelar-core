// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/vote/v1beta1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// TalliedVote represents a vote for a poll with the accumulated stake of all
// validators voting for the same data
type TalliedVote struct {
	Tally  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=tally,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tally"`
	Voters Voters                                 `protobuf:"bytes,2,rep,name=voters,proto3,castrepeated=Voters" json:"voters,omitempty"`
	Data   *types.Any                             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TalliedVote) Reset()         { *m = TalliedVote{} }
func (m *TalliedVote) String() string { return proto.CompactTextString(m) }
func (*TalliedVote) ProtoMessage()    {}
func (*TalliedVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_584be12bf9f97fd2, []int{0}
}
func (m *TalliedVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalliedVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TalliedVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TalliedVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalliedVote.Merge(m, src)
}
func (m *TalliedVote) XXX_Size() int {
	return m.Size()
}
func (m *TalliedVote) XXX_DiscardUnknown() {
	xxx_messageInfo_TalliedVote.DiscardUnknown(m)
}

var xxx_messageInfo_TalliedVote proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TalliedVote)(nil), "axelar.vote.v1beta1.TalliedVote")
}

func init() { proto.RegisterFile("axelar/vote/v1beta1/types.proto", fileDescriptor_584be12bf9f97fd2) }

var fileDescriptor_584be12bf9f97fd2 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4f, 0x02, 0x31,
	0x18, 0xc6, 0xaf, 0xa2, 0x0c, 0x07, 0xd3, 0xc9, 0x00, 0x0c, 0x3d, 0xc2, 0x60, 0x88, 0x09, 0x6d,
	0xd0, 0x4f, 0xe0, 0xc5, 0xc5, 0xc1, 0x68, 0x08, 0x61, 0x70, 0x31, 0xbd, 0xbb, 0x7a, 0x10, 0xca,
	0xbd, 0xa4, 0x57, 0x90, 0xfb, 0x16, 0x7e, 0x0e, 0x67, 0x3f, 0x04, 0x31, 0x31, 0x61, 0x34, 0x0e,
	0xa8, 0xdc, 0x17, 0x31, 0xd7, 0xd6, 0xc4, 0xc1, 0xa9, 0xef, 0x9f, 0xdf, 0xfb, 0xf6, 0x79, 0x5a,
	0xd7, 0x67, 0x6b, 0x2e, 0x98, 0xa4, 0x2b, 0x50, 0x9c, 0xae, 0x06, 0x21, 0x57, 0x6c, 0x40, 0x55,
	0xbe, 0xe0, 0x19, 0x59, 0x48, 0x50, 0xe0, 0x1d, 0x1b, 0x80, 0x94, 0x00, 0xb1, 0x40, 0xbb, 0x95,
	0x00, 0x24, 0x82, 0x53, 0x8d, 0x84, 0xcb, 0x07, 0xca, 0xd2, 0xdc, 0xf0, 0xed, 0x46, 0x02, 0x09,
	0xe8, 0x90, 0x96, 0x91, 0xad, 0xb6, 0x22, 0xc8, 0xe6, 0x90, 0xdd, 0x9b, 0x86, 0x49, 0x4c, 0xab,
	0xfb, 0x86, 0xdc, 0xda, 0x88, 0x09, 0x31, 0xe5, 0xf1, 0x18, 0x14, 0xf7, 0x2e, 0xdd, 0x23, 0xc5,
	0x84, 0xc8, 0x9b, 0xa8, 0x83, 0x7a, 0xf5, 0x80, 0x6c, 0x76, 0xbe, 0xf3, 0xb1, 0xf3, 0x4f, 0x92,
	0xa9, 0x9a, 0x2c, 0x43, 0x12, 0xc1, 0xdc, 0xce, 0xdb, 0xa3, 0x9f, 0xc5, 0x33, 0xab, 0xf8, 0x2a,
	0x55, 0x43, 0x33, 0xec, 0x75, 0xdd, 0x6a, 0xa9, 0x58, 0x66, 0xcd, 0x83, 0x4e, 0xa5, 0x57, 0x0f,
	0xdc, 0xe7, 0x4f, 0xbf, 0x3a, 0xd6, 0x95, 0xa1, 0xed, 0x78, 0x23, 0xf7, 0x30, 0x66, 0x8a, 0x35,
	0x2b, 0x1d, 0xd4, 0xab, 0x9d, 0x35, 0x88, 0x31, 0x45, 0x7e, 0x4d, 0x91, 0x8b, 0x34, 0x0f, 0x4e,
	0x5f, 0x5f, 0xfa, 0xff, 0x5e, 0x1d, 0xf3, 0x88, 0xde, 0x96, 0xe4, 0x35, 0x93, 0xd9, 0x84, 0x09,
	0x2e, 0x87, 0x7a, 0x5b, 0x70, 0xb3, 0xf9, 0xc6, 0xce, 0x66, 0x8f, 0xd1, 0x76, 0x8f, 0xd1, 0xd7,
	0x1e, 0xa3, 0xa7, 0x02, 0x3b, 0xdb, 0x02, 0x3b, 0xef, 0x05, 0x76, 0xee, 0x06, 0x7f, 0x76, 0x99,
	0x97, 0x4d, 0xb9, 0x7a, 0x04, 0x39, 0xb3, 0x59, 0x3f, 0x02, 0xc9, 0xe9, 0xda, 0xfc, 0x87, 0x76,
	0x15, 0x56, 0xb5, 0xa0, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x02, 0x7a, 0x42, 0xab,
	0x01, 0x00, 0x00,
}

func (m *TalliedVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalliedVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalliedVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.Tally.Size()
		i -= size
		if _, err := m.Tally.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *TalliedVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tally.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Voters) > 0 {
		for _, b := range m.Voters {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TalliedVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TalliedVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalliedVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tally", wireType)
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
			if err := m.Tally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
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
			m.Voters = append(m.Voters, make([]byte, postIndex-iNdEx))
			copy(m.Voters[len(m.Voters)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
