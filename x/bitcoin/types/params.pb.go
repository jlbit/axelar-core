// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/bitcoin/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	Network                              Network         `protobuf:"bytes,1,opt,name=network,proto3" json:"network"`
	ConfirmationHeight                   uint64          `protobuf:"varint,2,opt,name=confirmation_height,json=confirmationHeight,proto3" json:"confirmation_height,omitempty"`
	RevoteLockingPeriod                  int64           `protobuf:"varint,3,opt,name=revote_locking_period,json=revoteLockingPeriod,proto3" json:"revote_locking_period,omitempty"`
	SigCheckInterval                     int64           `protobuf:"varint,4,opt,name=sig_check_interval,json=sigCheckInterval,proto3" json:"sig_check_interval,omitempty"`
	MinOutputAmount                      types.DecCoin   `protobuf:"bytes,5,opt,name=min_output_amount,json=minOutputAmount,proto3" json:"min_output_amount"`
	MaxInputCount                        int64           `protobuf:"varint,6,opt,name=max_input_count,json=maxInputCount,proto3" json:"max_input_count,omitempty"`
	MaxSecondaryOutputAmount             types.DecCoin   `protobuf:"bytes,7,opt,name=max_secondary_output_amount,json=maxSecondaryOutputAmount,proto3" json:"max_secondary_output_amount"`
	MasterKeyRetentionPeriod             int64           `protobuf:"varint,8,opt,name=master_key_retention_period,json=masterKeyRetentionPeriod,proto3" json:"master_key_retention_period,omitempty"`
	MasterAddressInternalKeyLockDuration time.Duration   `protobuf:"varint,9,opt,name=master_address_internal_key_lock_duration,json=masterAddressInternalKeyLockDuration,proto3,customtype=time.Duration" json:"master_address_internal_key_lock_duration"`
	MasterAddressExternalKeyLockDuration time.Duration   `protobuf:"varint,10,opt,name=master_address_external_key_lock_duration,json=masterAddressExternalKeyLockDuration,proto3,customtype=time.Duration" json:"master_address_external_key_lock_duration"`
	VotingThreshold                      utils.Threshold `protobuf:"bytes,11,opt,name=voting_threshold,json=votingThreshold,proto3" json:"voting_threshold"`
	MinVoterCount                        int64           `protobuf:"varint,12,opt,name=min_voter_count,json=minVoterCount,proto3" json:"min_voter_count,omitempty"`
	MaxTxSize                            int64           `protobuf:"varint,13,opt,name=max_tx_size,json=maxTxSize,proto3" json:"max_tx_size,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3376470cbac61574, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.bitcoin.v1beta1.Params")
}

func init() {
	proto.RegisterFile("axelar/bitcoin/v1beta1/params.proto", fileDescriptor_3376470cbac61574)
}

var fileDescriptor_3376470cbac61574 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4e, 0x14, 0x4d,
	0x14, 0x9d, 0xfe, 0x98, 0x6f, 0x80, 0x42, 0x04, 0x1b, 0x31, 0x1d, 0x34, 0x0d, 0x41, 0x62, 0x30,
	0xd1, 0xee, 0x80, 0x6e, 0x8d, 0xe1, 0xc7, 0x44, 0xc4, 0x20, 0x19, 0x88, 0x0b, 0x37, 0x95, 0x9a,
	0x9e, 0x6b, 0x4f, 0x65, 0xa6, 0xaa, 0x26, 0x55, 0xd5, 0x63, 0x0f, 0x4f, 0xe1, 0x63, 0xb1, 0x64,
	0x69, 0x5c, 0x10, 0x85, 0x77, 0x70, 0x6d, 0xea, 0xa7, 0x5b, 0x31, 0x90, 0xb0, 0x9b, 0xa9, 0x73,
	0xee, 0xb9, 0xa7, 0xee, 0xb9, 0xd5, 0xe8, 0x31, 0x29, 0x61, 0x40, 0x64, 0xda, 0xa1, 0x3a, 0x13,
	0x94, 0xa7, 0xa3, 0x8d, 0x0e, 0x68, 0xb2, 0x91, 0x0e, 0x89, 0x24, 0x4c, 0x25, 0x43, 0x29, 0xb4,
	0x08, 0x1f, 0x38, 0x52, 0xe2, 0x49, 0x89, 0x27, 0x2d, 0xdd, 0xcf, 0x45, 0x2e, 0x2c, 0x25, 0x35,
	0xbf, 0x1c, 0x7b, 0x69, 0xf5, 0x06, 0x49, 0x3d, 0x1e, 0x82, 0x57, 0x5c, 0x8a, 0x33, 0xa1, 0x98,
	0x50, 0x69, 0x87, 0x28, 0xa8, 0x09, 0x56, 0xdb, 0xe1, 0x6b, 0x5e, 0xa3, 0xd0, 0x74, 0xa0, 0xfe,
	0x28, 0xf4, 0x24, 0xa8, 0x9e, 0x18, 0x74, 0x1d, 0x6b, 0xf5, 0x57, 0x0b, 0xb5, 0x0e, 0xad, 0xd1,
	0xf0, 0x35, 0x9a, 0xe4, 0xa0, 0xbf, 0x08, 0xd9, 0x8f, 0x82, 0x95, 0x60, 0x7d, 0x66, 0x73, 0x39,
	0xb9, 0xde, 0x74, 0x72, 0xe0, 0x68, 0xdb, 0xcd, 0xd3, 0xf3, 0xe5, 0x46, 0xbb, 0xaa, 0x0a, 0x53,
	0xb4, 0x90, 0x09, 0xfe, 0x99, 0x4a, 0x46, 0x34, 0x15, 0x1c, 0xf7, 0x80, 0xe6, 0x3d, 0x1d, 0xfd,
	0xb7, 0x12, 0xac, 0x37, 0xdb, 0xe1, 0xdf, 0xd0, 0x5b, 0x8b, 0x84, 0x9b, 0x68, 0x51, 0xc2, 0x48,
	0x68, 0xc0, 0x03, 0x91, 0xf5, 0x29, 0xcf, 0xf1, 0x10, 0x24, 0x15, 0xdd, 0x68, 0x62, 0x25, 0x58,
	0x9f, 0x68, 0x2f, 0x38, 0xf0, 0xbd, 0xc3, 0x0e, 0x2d, 0x14, 0x3e, 0x43, 0xa1, 0xa2, 0x39, 0xce,
	0x7a, 0x90, 0xf5, 0x31, 0xe5, 0x1a, 0xe4, 0x88, 0x0c, 0xa2, 0xa6, 0x2d, 0x98, 0x57, 0x34, 0xdf,
	0x31, 0xc0, 0x9e, 0x3f, 0x0f, 0x0f, 0xd0, 0x3d, 0x46, 0x39, 0x16, 0x85, 0x1e, 0x16, 0x1a, 0x13,
	0x26, 0x0a, 0xae, 0xa3, 0xff, 0xed, 0xed, 0x1e, 0x25, 0x6e, 0x80, 0x89, 0x19, 0x60, 0x7d, 0xb5,
	0x5d, 0xc8, 0x76, 0x04, 0xe5, 0xfe, 0x6a, 0x73, 0x8c, 0xf2, 0x0f, 0xb6, 0x76, 0xcb, 0x96, 0x86,
	0x4f, 0xd0, 0x1c, 0x23, 0x25, 0xa6, 0xdc, 0xc8, 0x65, 0x56, 0xad, 0x65, 0x5b, 0xcf, 0x32, 0x52,
	0xee, 0x99, 0xd3, 0x1d, 0xcb, 0x23, 0xe8, 0xa1, 0xe1, 0x29, 0xc8, 0x04, 0xef, 0x12, 0x39, 0xfe,
	0xc7, 0xc1, 0xe4, 0xad, 0x1d, 0x44, 0x8c, 0x94, 0x47, 0x95, 0xca, 0x15, 0x2b, 0xaf, 0x4c, 0x0b,
	0xa5, 0x41, 0xe2, 0x3e, 0x8c, 0xb1, 0x04, 0x0d, 0xdc, 0x4e, 0xdd, 0x8f, 0x70, 0xca, 0xda, 0x8a,
	0x1c, 0x65, 0x1f, 0xc6, 0xed, 0x8a, 0xe0, 0xe7, 0xc8, 0xd1, 0x53, 0x5f, 0x4e, 0xba, 0x5d, 0x09,
	0x4a, 0xb9, 0x61, 0x72, 0x32, 0xb0, 0x7a, 0x26, 0x10, 0xdc, 0x2d, 0xa4, 0x8d, 0x2b, 0x9a, 0x36,
	0x62, 0xdb, 0x8b, 0xc6, 0xd1, 0xf7, 0xf3, 0xe5, 0x59, 0x4d, 0x19, 0x24, 0xbb, 0x1e, 0x6c, 0xaf,
	0x39, 0x9d, 0x2d, 0x27, 0xb3, 0xe7, 0x55, 0xf6, 0x61, 0x6c, 0x82, 0xab, 0x58, 0xd7, 0xf4, 0x83,
	0xf2, 0xc6, 0x7e, 0xe8, 0xf6, 0xfd, 0xde, 0x94, 0xd7, 0xf7, 0x3b, 0x44, 0xf3, 0x23, 0xa1, 0xcd,
	0x4e, 0xd5, 0x2b, 0x1f, 0xcd, 0x5c, 0x5d, 0x6b, 0xfb, 0x32, 0xea, 0xb9, 0x1f, 0x57, 0xb4, 0x2a,
	0x7b, 0x57, 0x5e, 0x1f, 0xdb, 0xec, 0x29, 0xc7, 0x66, 0x25, 0xa5, 0xcf, 0xfe, 0x8e, 0xcf, 0x9e,
	0xf2, 0x8f, 0xe6, 0xd4, 0x65, 0x1f, 0xa3, 0x19, 0x93, 0xbd, 0x2e, 0xb1, 0xa2, 0x27, 0x10, 0xcd,
	0x5a, 0xce, 0x34, 0x23, 0xe5, 0x71, 0x79, 0x44, 0x4f, 0xe0, 0x5d, 0x73, 0xea, 0xee, 0xfc, 0xdc,
	0x76, 0xfb, 0xf4, 0x67, 0xdc, 0x38, 0xbd, 0x88, 0x83, 0xb3, 0x8b, 0x38, 0xf8, 0x71, 0x11, 0x07,
	0x5f, 0x2f, 0xe3, 0xc6, 0xd9, 0x65, 0xdc, 0xf8, 0x76, 0x19, 0x37, 0x3e, 0xbd, 0xcc, 0xa9, 0xee,
	0x15, 0x9d, 0x24, 0x13, 0x2c, 0x75, 0x6e, 0xfd, 0x23, 0xf3, 0xff, 0x9e, 0x67, 0x42, 0x42, 0x5a,
	0xd6, 0x1f, 0x08, 0xfb, 0x61, 0xe8, 0xb4, 0xec, 0x9b, 0x7e, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x8e, 0xa2, 0x01, 0x92, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTxSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxTxSize))
		i--
		dAtA[i] = 0x68
	}
	if m.MinVoterCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinVoterCount))
		i--
		dAtA[i] = 0x60
	}
	{
		size, err := m.VotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.MasterAddressExternalKeyLockDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MasterAddressExternalKeyLockDuration))
		i--
		dAtA[i] = 0x50
	}
	if m.MasterAddressInternalKeyLockDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MasterAddressInternalKeyLockDuration))
		i--
		dAtA[i] = 0x48
	}
	if m.MasterKeyRetentionPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MasterKeyRetentionPeriod))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.MaxSecondaryOutputAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.MaxInputCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxInputCount))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.MinOutputAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.SigCheckInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SigCheckInterval))
		i--
		dAtA[i] = 0x20
	}
	if m.RevoteLockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevoteLockingPeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.ConfirmationHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConfirmationHeight))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Network.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Network.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.ConfirmationHeight != 0 {
		n += 1 + sovParams(uint64(m.ConfirmationHeight))
	}
	if m.RevoteLockingPeriod != 0 {
		n += 1 + sovParams(uint64(m.RevoteLockingPeriod))
	}
	if m.SigCheckInterval != 0 {
		n += 1 + sovParams(uint64(m.SigCheckInterval))
	}
	l = m.MinOutputAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxInputCount != 0 {
		n += 1 + sovParams(uint64(m.MaxInputCount))
	}
	l = m.MaxSecondaryOutputAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MasterKeyRetentionPeriod != 0 {
		n += 1 + sovParams(uint64(m.MasterKeyRetentionPeriod))
	}
	if m.MasterAddressInternalKeyLockDuration != 0 {
		n += 1 + sovParams(uint64(m.MasterAddressInternalKeyLockDuration))
	}
	if m.MasterAddressExternalKeyLockDuration != 0 {
		n += 1 + sovParams(uint64(m.MasterAddressExternalKeyLockDuration))
	}
	l = m.VotingThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinVoterCount != 0 {
		n += 1 + sovParams(uint64(m.MinVoterCount))
	}
	if m.MaxTxSize != 0 {
		n += 1 + sovParams(uint64(m.MaxTxSize))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Network.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationHeight", wireType)
			}
			m.ConfirmationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevoteLockingPeriod", wireType)
			}
			m.RevoteLockingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevoteLockingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigCheckInterval", wireType)
			}
			m.SigCheckInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigCheckInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOutputAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinOutputAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxInputCount", wireType)
			}
			m.MaxInputCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxInputCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSecondaryOutputAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSecondaryOutputAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterKeyRetentionPeriod", wireType)
			}
			m.MasterKeyRetentionPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MasterKeyRetentionPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterAddressInternalKeyLockDuration", wireType)
			}
			m.MasterAddressInternalKeyLockDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MasterAddressInternalKeyLockDuration |= time.Duration(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterAddressExternalKeyLockDuration", wireType)
			}
			m.MasterAddressExternalKeyLockDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MasterAddressExternalKeyLockDuration |= time.Duration(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVoterCount", wireType)
			}
			m.MinVoterCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinVoterCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxSize", wireType)
			}
			m.MaxTxSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
