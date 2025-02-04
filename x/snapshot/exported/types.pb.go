// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/snapshot/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type ValidatorIllegibility int32

const (
	// these enum values are used for bitwise operations, therefore they need to
	// be powers of 2
	None                  ValidatorIllegibility = 0
	Tombstoned            ValidatorIllegibility = 1
	Jailed                ValidatorIllegibility = 2
	MissedTooManyBlocks   ValidatorIllegibility = 4
	NoProxyRegistered     ValidatorIllegibility = 8
	TssSuspended          ValidatorIllegibility = 16
	ProxyInsuficientFunds ValidatorIllegibility = 32
)

var ValidatorIllegibility_name = map[int32]string{
	0:  "VALIDATOR_ILLEGIBILITY_UNSPECIFIED",
	1:  "VALIDATOR_ILLEGIBILITY_TOMBSTONED",
	2:  "VALIDATOR_ILLEGIBILITY_JAILED",
	4:  "VALIDATOR_ILLEGIBILITY_MISSED_TOO_MANY_BLOCKS",
	8:  "VALIDATOR_ILLEGIBILITY_NO_PROXY_REGISTERED",
	16: "VALIDATOR_ILLEGIBILITY_TSS_SUSPENDED",
	32: "VALIDATOR_ILLEGIBILITY_PROXY_INSUFICIENT_FUNDS",
}

var ValidatorIllegibility_value = map[string]int32{
	"VALIDATOR_ILLEGIBILITY_UNSPECIFIED":             0,
	"VALIDATOR_ILLEGIBILITY_TOMBSTONED":              1,
	"VALIDATOR_ILLEGIBILITY_JAILED":                  2,
	"VALIDATOR_ILLEGIBILITY_MISSED_TOO_MANY_BLOCKS":  4,
	"VALIDATOR_ILLEGIBILITY_NO_PROXY_REGISTERED":     8,
	"VALIDATOR_ILLEGIBILITY_TSS_SUSPENDED":           16,
	"VALIDATOR_ILLEGIBILITY_PROXY_INSUFICIENT_FUNDS": 32,
}

func (ValidatorIllegibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eabe75f8f44c6f51, []int{0}
}

type Validator struct {
	SDKValidator *types.Any `protobuf:"bytes,1,opt,name=sdk_validator,json=sdkValidator,proto3" json:"sdk_validator,omitempty"`
	ShareCount   int64      `protobuf:"varint,2,opt,name=share_count,json=shareCount,proto3" json:"share_count,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_eabe75f8f44c6f51, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

type Snapshot struct {
	Validators                 []Validator                            `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators"`
	Timestamp                  time.Time                              `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Height                     int64                                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	TotalShareCount            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_share_count,json=totalShareCount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_share_count"`
	Counter                    int64                                  `protobuf:"varint,5,opt,name=counter,proto3" json:"counter,omitempty"`
	KeyShareDistributionPolicy exported.KeyShareDistributionPolicy    `protobuf:"varint,6,opt,name=key_share_distribution_policy,json=keyShareDistributionPolicy,proto3,enum=axelar.tss.exported.v1beta1.KeyShareDistributionPolicy" json:"key_share_distribution_policy,omitempty"`
	CorruptionThreshold        int64                                  `protobuf:"varint,7,opt,name=corruption_threshold,json=corruptionThreshold,proto3" json:"corruption_threshold,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_eabe75f8f44c6f51, []int{1}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("axelar.snapshot.exported.v1beta1.ValidatorIllegibility", ValidatorIllegibility_name, ValidatorIllegibility_value)
	proto.RegisterType((*Validator)(nil), "axelar.snapshot.exported.v1beta1.Validator")
	proto.RegisterType((*Snapshot)(nil), "axelar.snapshot.exported.v1beta1.Snapshot")
}

func init() {
	proto.RegisterFile("axelar/snapshot/exported/v1beta1/types.proto", fileDescriptor_eabe75f8f44c6f51)
}

var fileDescriptor_eabe75f8f44c6f51 = []byte{
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0x37, 0xa5, 0x9b, 0x9d, 0x96, 0xc5, 0xeb, 0x6d, 0x21, 0xb5, 0xb4, 0x8e, 0xa9, 0x56,
	0x10, 0x16, 0x62, 0xd3, 0x22, 0x04, 0xe2, 0x16, 0xd7, 0xee, 0xca, 0x6d, 0xe2, 0x04, 0xdb, 0xad,
	0xe8, 0x5e, 0x2c, 0xc7, 0x9e, 0x75, 0x46, 0x71, 0x66, 0x22, 0xcf, 0x64, 0xa9, 0xf9, 0x05, 0x28,
	0xa7, 0x15, 0xf7, 0x9c, 0xe0, 0xc0, 0x0f, 0xe0, 0x47, 0x54, 0x9c, 0xf6, 0x06, 0xe2, 0x50, 0xa0,
	0xfd, 0x23, 0xa8, 0xb6, 0xe3, 0x46, 0x54, 0x81, 0x93, 0xfd, 0xe6, 0x7d, 0xdf, 0xfb, 0xbe, 0xf7,
	0xde, 0xd8, 0xe0, 0x13, 0xff, 0x1c, 0xc6, 0x7e, 0xa2, 0x52, 0xec, 0x4f, 0xe8, 0x90, 0x30, 0x15,
	0x9e, 0x4f, 0x48, 0xc2, 0x60, 0xa8, 0xbe, 0xda, 0x1b, 0x40, 0xe6, 0xef, 0xa9, 0x2c, 0x9d, 0x40,
	0xaa, 0x4c, 0x12, 0xc2, 0x88, 0x20, 0xe7, 0x68, 0x65, 0x81, 0x56, 0x16, 0x68, 0xa5, 0x40, 0x8b,
	0x5b, 0x11, 0x89, 0x48, 0x06, 0x56, 0x6f, 0xde, 0x72, 0x9e, 0xd8, 0x88, 0x08, 0x89, 0x62, 0xa8,
	0x66, 0xd1, 0x60, 0xfa, 0x52, 0x65, 0x68, 0x0c, 0x29, 0xf3, 0xc7, 0x93, 0x02, 0xb0, 0xf3, 0x6f,
	0x80, 0x8f, 0xd3, 0x22, 0x25, 0x05, 0x84, 0x8e, 0x09, 0x55, 0x07, 0x3e, 0x85, 0xa5, 0xa9, 0x80,
	0x20, 0x5c, 0xe4, 0x3f, 0x2c, 0x3a, 0x60, 0x94, 0xfe, 0xa7, 0x79, 0x71, 0x27, 0x2f, 0xe4, 0xe5,
	0xee, 0xf2, 0xa0, 0x48, 0x3d, 0x2d, 0x34, 0x28, 0xf3, 0x47, 0x08, 0x47, 0x25, 0xbd, 0x88, 0x73,
	0xd4, 0xee, 0x0f, 0x1c, 0x78, 0x70, 0xea, 0xc7, 0x28, 0xf4, 0x19, 0x49, 0x84, 0x10, 0xbc, 0x4d,
	0xc3, 0x91, 0xf7, 0x6a, 0x71, 0x50, 0xe7, 0x64, 0xae, 0xb9, 0xb1, 0xbf, 0xa5, 0xe4, 0xad, 0x28,
	0x8b, 0x56, 0x94, 0x36, 0x4e, 0xb5, 0x8f, 0xae, 0x2e, 0x1b, 0x9b, 0x8e, 0x7e, 0x5c, 0xd2, 0x7f,
	0xfd, 0xa5, 0xb5, 0x5d, 0x0e, 0x6f, 0x39, 0x61, 0x6f, 0xd2, 0x70, 0x74, 0xab, 0xd2, 0x00, 0x1b,
	0x74, 0xe8, 0x27, 0xd0, 0x0b, 0xc8, 0x14, 0xb3, 0xfa, 0x3d, 0x99, 0x6b, 0x56, 0x6d, 0x90, 0x1d,
	0x1d, 0xdc, 0x9c, 0xec, 0x5e, 0x56, 0x41, 0xcd, 0x29, 0xd6, 0x21, 0x7c, 0x0d, 0x40, 0xe9, 0x87,
	0xd6, 0x39, 0xb9, 0xda, 0xdc, 0xd8, 0xff, 0x58, 0xf9, 0xbf, 0xa5, 0x29, 0xa5, 0x9c, 0xb6, 0x76,
	0x71, 0xd9, 0xa8, 0xd8, 0x4b, 0x45, 0x04, 0x0d, 0x3c, 0x28, 0x97, 0x95, 0xc9, 0x6f, 0xec, 0x8b,
	0x77, 0x5a, 0x74, 0x17, 0x08, 0xad, 0x76, 0x53, 0xe0, 0xf5, 0x9f, 0x0d, 0xce, 0xbe, 0xa5, 0x09,
	0xef, 0x82, 0xf5, 0x21, 0x44, 0xd1, 0x90, 0xd5, 0xab, 0x99, 0xff, 0x22, 0x12, 0x5e, 0x80, 0x47,
	0x8c, 0x30, 0x3f, 0xf6, 0x96, 0x5b, 0x5c, 0x93, 0xb9, 0xe6, 0xa6, 0xa6, 0xdc, 0xd4, 0xf9, 0xe3,
	0xb2, 0xf1, 0x41, 0x84, 0xd8, 0x70, 0x3a, 0x50, 0x02, 0x32, 0x2e, 0x56, 0x56, 0x3c, 0x5a, 0x34,
	0x1c, 0x15, 0xeb, 0x35, 0x31, 0xb3, 0xdf, 0xc9, 0x0a, 0x39, 0xe5, 0x5c, 0x84, 0x3a, 0xb8, 0x9f,
	0xd5, 0x83, 0x49, 0xfd, 0xad, 0x4c, 0x74, 0x11, 0x0a, 0xdf, 0x81, 0x27, 0x23, 0x98, 0x16, 0x9a,
	0x21, 0xa2, 0x2c, 0x41, 0x83, 0x29, 0x43, 0x04, 0x7b, 0x13, 0x12, 0xa3, 0x20, 0xad, 0xaf, 0xcb,
	0x5c, 0xf3, 0xe1, 0xfe, 0x17, 0x8b, 0xb9, 0x31, 0x4a, 0xef, 0x8e, 0xec, 0x18, 0xa6, 0x99, 0x98,
	0xbe, 0xc4, 0xef, 0x67, 0x74, 0x5b, 0x1c, 0xad, 0xcc, 0x09, 0x7b, 0x60, 0x2b, 0x20, 0x49, 0x32,
	0x9d, 0x64, 0x7a, 0x6c, 0x98, 0x40, 0x3a, 0x24, 0x71, 0x58, 0xbf, 0x9f, 0x59, 0x7c, 0x7c, 0x9b,
	0x73, 0x17, 0xa9, 0x67, 0xbf, 0x55, 0xc1, 0x76, 0xb9, 0x20, 0x33, 0x8e, 0x61, 0x84, 0x06, 0x28,
	0x46, 0x2c, 0x15, 0x3e, 0x05, 0xbb, 0xa7, 0xed, 0x8e, 0xa9, 0xb7, 0xdd, 0x9e, 0xed, 0x99, 0x9d,
	0x8e, 0xf1, 0xdc, 0xd4, 0xcc, 0x8e, 0xe9, 0x9e, 0x79, 0x27, 0x96, 0xd3, 0x37, 0x0e, 0xcc, 0x43,
	0xd3, 0xd0, 0xf9, 0x8a, 0x58, 0x9b, 0xcd, 0xe5, 0x35, 0x8b, 0x60, 0x28, 0x7c, 0x0e, 0xde, 0x5f,
	0xc1, 0x70, 0x7b, 0x5d, 0xcd, 0x71, 0x7b, 0x96, 0xa1, 0xf3, 0x9c, 0xf8, 0x70, 0x36, 0x97, 0x81,
	0x4b, 0xc6, 0x03, 0xca, 0x08, 0x86, 0xa1, 0xd0, 0x02, 0x4f, 0x56, 0xd0, 0x8e, 0xda, 0x66, 0xc7,
	0xd0, 0xf9, 0x7b, 0x22, 0x98, 0xcd, 0xe5, 0xf5, 0x23, 0x1f, 0xc5, 0x30, 0x14, 0x8e, 0x40, 0x6b,
	0x05, 0xbc, 0x6b, 0x3a, 0x8e, 0xa1, 0x7b, 0x6e, 0xaf, 0xe7, 0x75, 0xdb, 0xd6, 0x99, 0xa7, 0x75,
	0x7a, 0x07, 0xc7, 0x0e, 0xbf, 0x26, 0xbe, 0x37, 0x9b, 0xcb, 0x8f, 0xbb, 0x88, 0x52, 0x18, 0xba,
	0x84, 0x74, 0x7d, 0x9c, 0x6a, 0x31, 0x09, 0x46, 0x54, 0x30, 0xc0, 0xb3, 0x15, 0xb5, 0xac, 0x9e,
	0xd7, 0xb7, 0x7b, 0xdf, 0x9c, 0x79, 0xb6, 0xf1, 0xdc, 0x74, 0x5c, 0xc3, 0x36, 0x74, 0xbe, 0x26,
	0x6e, 0xcf, 0xe6, 0xf2, 0x23, 0x8b, 0xf4, 0x13, 0x72, 0x9e, 0xda, 0x30, 0x42, 0x94, 0xc1, 0x04,
	0x86, 0xc2, 0x57, 0xe0, 0xe9, 0xaa, 0xc6, 0x1d, 0xc7, 0x73, 0x4e, 0x9c, 0xbe, 0x61, 0xe9, 0x86,
	0xce, 0xf3, 0x22, 0x3f, 0x9b, 0xcb, 0x9b, 0x2e, 0xa5, 0xce, 0x94, 0x4e, 0x20, 0x0e, 0x61, 0x28,
	0x74, 0x81, 0xb2, 0x82, 0x9b, 0xeb, 0x9b, 0x96, 0x73, 0x72, 0x68, 0x1e, 0x98, 0x86, 0xe5, 0x7a,
	0x87, 0x27, 0x96, 0xee, 0xf0, 0xb2, 0xb8, 0x33, 0x9b, 0xcb, 0xdb, 0x99, 0x09, 0x13, 0xd3, 0xe9,
	0x4b, 0x14, 0x20, 0x88, 0xd9, 0xe1, 0x14, 0x87, 0x54, 0xac, 0x7d, 0xff, 0xa3, 0x54, 0xf9, 0xf9,
	0x27, 0xa9, 0xa2, 0x9d, 0x5e, 0xfc, 0x2d, 0x55, 0x2e, 0xae, 0x24, 0xee, 0xcd, 0x95, 0xc4, 0xfd,
	0x75, 0x25, 0x71, 0xaf, 0xaf, 0xa5, 0xca, 0x9b, 0x6b, 0xa9, 0xf2, 0xfb, 0xb5, 0x54, 0x79, 0xf1,
	0xe5, 0xd2, 0xcd, 0xcf, 0x6f, 0x22, 0x86, 0xec, 0x5b, 0x92, 0x8c, 0x8a, 0xa8, 0x15, 0x90, 0x04,
	0xaa, 0xe7, 0x77, 0xff, 0xdc, 0x83, 0xf5, 0xec, 0xbb, 0xfc, 0xec, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x77, 0xcf, 0xce, 0xe4, 0xdc, 0x05, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShareCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ShareCount))
		i--
		dAtA[i] = 0x10
	}
	if m.SDKValidator != nil {
		{
			size, err := m.SDKValidator.MarshalToSizedBuffer(dAtA[:i])
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

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CorruptionThreshold != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CorruptionThreshold))
		i--
		dAtA[i] = 0x38
	}
	if m.KeyShareDistributionPolicy != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyShareDistributionPolicy))
		i--
		dAtA[i] = 0x30
	}
	if m.Counter != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.TotalShareCount.Size()
		i -= size
		if _, err := m.TotalShareCount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SDKValidator != nil {
		l = m.SDKValidator.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ShareCount != 0 {
		n += 1 + sovTypes(uint64(m.ShareCount))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = m.TotalShareCount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Counter != 0 {
		n += 1 + sovTypes(uint64(m.Counter))
	}
	if m.KeyShareDistributionPolicy != 0 {
		n += 1 + sovTypes(uint64(m.KeyShareDistributionPolicy))
	}
	if m.CorruptionThreshold != 0 {
		n += 1 + sovTypes(uint64(m.CorruptionThreshold))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SDKValidator", wireType)
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
			if m.SDKValidator == nil {
				m.SDKValidator = &types.Any{}
			}
			if err := m.SDKValidator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCount", wireType)
			}
			m.ShareCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShareCount", wireType)
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
			if err := m.TotalShareCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShareDistributionPolicy", wireType)
			}
			m.KeyShareDistributionPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyShareDistributionPolicy |= exported.KeyShareDistributionPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorruptionThreshold", wireType)
			}
			m.CorruptionThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CorruptionThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
