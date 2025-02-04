// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/vote/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/permission/exported"
	exported "github.com/axelarnetwork/axelar-core/x/vote/exported"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type VoteRequest struct {
	Sender  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	PollKey exported.PollKey                              `protobuf:"bytes,2,opt,name=poll_key,json=pollKey,proto3" json:"poll_key"`
	Vote    exported.Vote                                 `protobuf:"bytes,3,opt,name=vote,proto3" json:"vote"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5a35666e11406c5, []int{0}
}
func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(m, src)
}
func (m *VoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

type VoteResponse struct {
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *VoteResponse) Reset()         { *m = VoteResponse{} }
func (m *VoteResponse) String() string { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()    {}
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5a35666e11406c5, []int{1}
}
func (m *VoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteResponse.Merge(m, src)
}
func (m *VoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *VoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VoteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VoteRequest)(nil), "axelar.vote.v1beta1.VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "axelar.vote.v1beta1.VoteResponse")
}

func init() { proto.RegisterFile("axelar/vote/v1beta1/tx.proto", fileDescriptor_b5a35666e11406c5) }

var fileDescriptor_b5a35666e11406c5 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbd, 0x4e, 0x32, 0x41,
	0x14, 0xdd, 0xf9, 0x20, 0x7c, 0xdf, 0xb7, 0x50, 0x98, 0xd5, 0x62, 0x21, 0x66, 0x21, 0x24, 0x26,
	0x34, 0xec, 0x66, 0xb5, 0x33, 0x36, 0x50, 0x98, 0x18, 0x0b, 0xcd, 0x16, 0x16, 0x36, 0x64, 0x7f,
	0xae, 0x2b, 0x61, 0xd9, 0xbb, 0xce, 0x0c, 0xb8, 0xdb, 0xf9, 0x08, 0xbe, 0x8c, 0xef, 0x40, 0x49,
	0x69, 0x45, 0x14, 0x7c, 0x0a, 0x2b, 0xb3, 0x33, 0x43, 0xa4, 0xa2, 0x9a, 0x7b, 0xe7, 0x9c, 0x39,
	0xf7, 0x9e, 0x33, 0xfa, 0xb1, 0x9f, 0x43, 0xe2, 0x53, 0x67, 0x8e, 0x1c, 0x9c, 0xb9, 0x1b, 0x00,
	0xf7, 0x5d, 0x87, 0xe7, 0x76, 0x46, 0x91, 0xa3, 0x71, 0x28, 0x51, 0xbb, 0x44, 0x6d, 0x85, 0xb6,
	0x9a, 0x31, 0x62, 0x9c, 0x80, 0x23, 0x28, 0xc1, 0xec, 0xc1, 0xf1, 0xd3, 0x42, 0xf2, 0x5b, 0x47,
	0x31, 0xc6, 0x28, 0x4a, 0xa7, 0xac, 0xd4, 0x6d, 0x33, 0x44, 0x36, 0x45, 0x36, 0x92, 0x80, 0x6c,
	0x14, 0xd4, 0xdb, 0x1d, 0x0f, 0x79, 0x86, 0x94, 0x43, 0xf4, 0xbb, 0x47, 0x91, 0xc1, 0x96, 0x69,
	0x2b, 0x66, 0x06, 0x74, 0x3a, 0x66, 0x6c, 0x8c, 0xe9, 0x5e, 0x7e, 0xf7, 0x8b, 0xe8, 0xf5, 0x3b,
	0xe4, 0xe0, 0xc1, 0xd3, 0x0c, 0x18, 0x37, 0xae, 0xf4, 0x1a, 0x83, 0x34, 0x02, 0x6a, 0x92, 0x0e,
	0xe9, 0x35, 0x86, 0xee, 0xf7, 0xaa, 0xdd, 0x8f, 0xc7, 0xfc, 0x71, 0x16, 0xd8, 0x21, 0x4e, 0xd5,
	0x5a, 0xea, 0xe8, 0xb3, 0x68, 0xa2, 0xd4, 0x06, 0x61, 0x38, 0x88, 0x22, 0x0a, 0x8c, 0x79, 0x4a,
	0xc0, 0xb8, 0xd4, 0xff, 0x65, 0x98, 0x24, 0xa3, 0x09, 0x14, 0xe6, 0x9f, 0x0e, 0xe9, 0xd5, 0x4f,
	0x4f, 0xec, 0xdd, 0xa0, 0xb6, 0x7b, 0x6d, 0x13, 0xb3, 0x6f, 0x31, 0x49, 0xae, 0xa1, 0x18, 0x56,
	0x17, 0xab, 0xb6, 0xe6, 0xfd, 0xcd, 0x64, 0x6b, 0x5c, 0xe8, 0xd5, 0x92, 0x6f, 0x56, 0x84, 0x46,
	0x77, 0xbf, 0x46, 0xe9, 0x45, 0x09, 0x88, 0x57, 0xe7, 0xd5, 0x97, 0x37, 0x93, 0x74, 0x3b, 0x7a,
	0x43, 0xba, 0x64, 0x19, 0xa6, 0x0c, 0x8c, 0x03, 0xbd, 0x92, 0x60, 0x2c, 0x3c, 0xfe, 0xf7, 0xca,
	0x72, 0x78, 0xb3, 0xf8, 0xb4, 0xb4, 0xc5, 0xda, 0x22, 0xcb, 0xb5, 0x45, 0x3e, 0xd6, 0x16, 0x79,
	0xdd, 0x58, 0xda, 0x72, 0x63, 0x69, 0xef, 0x1b, 0x4b, 0xbb, 0x77, 0x77, 0x22, 0x90, 0xf3, 0x53,
	0xe0, 0xcf, 0x48, 0x27, 0xaa, 0xeb, 0x87, 0x48, 0xc1, 0xc9, 0xe5, 0x07, 0x89, 0x44, 0x82, 0x9a,
	0x08, 0xf8, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xe6, 0x66, 0x43, 0x3b, 0x02, 0x00, 0x00,
}

func (m *VoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PollKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.PollKey.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Vote.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *VoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: VoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PollKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *VoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: VoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
