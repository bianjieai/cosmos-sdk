// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/distribution/types/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// msg struct for changing the withdraw address for a delegator (or validator self-delegation)
type MsgSetWithdrawAddress struct {
	DelegatorAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"delegator_address,omitempty" yaml:"delegator_address"`
	WithdrawAddress  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"withdraw_address,omitempty" yaml:"withdraw_address"`
}

func (m *MsgSetWithdrawAddress) Reset()         { *m = MsgSetWithdrawAddress{} }
func (m *MsgSetWithdrawAddress) String() string { return proto.CompactTextString(m) }
func (*MsgSetWithdrawAddress) ProtoMessage()    {}
func (*MsgSetWithdrawAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fddf2a8e4a90b09, []int{0}
}
func (m *MsgSetWithdrawAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetWithdrawAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetWithdrawAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetWithdrawAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetWithdrawAddress.Merge(m, src)
}
func (m *MsgSetWithdrawAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetWithdrawAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetWithdrawAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetWithdrawAddress proto.InternalMessageInfo

func (m *MsgSetWithdrawAddress) GetDelegatorAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.DelegatorAddress
	}
	return nil
}

func (m *MsgSetWithdrawAddress) GetWithdrawAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.WithdrawAddress
	}
	return nil
}

// msg struct for delegation withdraw from a single validator
type MsgWithdrawDelegatorReward struct {
	DelegatorAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator_address,omitempty" yaml:"validator_address"`
}

func (m *MsgWithdrawDelegatorReward) Reset()         { *m = MsgWithdrawDelegatorReward{} }
func (m *MsgWithdrawDelegatorReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawDelegatorReward) ProtoMessage()    {}
func (*MsgWithdrawDelegatorReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fddf2a8e4a90b09, []int{1}
}
func (m *MsgWithdrawDelegatorReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawDelegatorReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawDelegatorReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawDelegatorReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawDelegatorReward.Merge(m, src)
}
func (m *MsgWithdrawDelegatorReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawDelegatorReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawDelegatorReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawDelegatorReward proto.InternalMessageInfo

func (m *MsgWithdrawDelegatorReward) GetDelegatorAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.DelegatorAddress
	}
	return nil
}

func (m *MsgWithdrawDelegatorReward) GetValidatorAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

// msg struct for validator withdraw
type MsgWithdrawValidatorCommission struct {
	ValidatorAddress github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator_address,omitempty" yaml:"validator_address"`
}

func (m *MsgWithdrawValidatorCommission) Reset()         { *m = MsgWithdrawValidatorCommission{} }
func (m *MsgWithdrawValidatorCommission) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawValidatorCommission) ProtoMessage()    {}
func (*MsgWithdrawValidatorCommission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fddf2a8e4a90b09, []int{2}
}
func (m *MsgWithdrawValidatorCommission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawValidatorCommission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawValidatorCommission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawValidatorCommission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawValidatorCommission.Merge(m, src)
}
func (m *MsgWithdrawValidatorCommission) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawValidatorCommission) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawValidatorCommission.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawValidatorCommission proto.InternalMessageInfo

func (m *MsgWithdrawValidatorCommission) GetValidatorAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

// MsgFundCommunityPool defines a Msg type that allows an account to directly
// fund the community pool.
type MsgFundCommunityPool struct {
	Amount    github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty" yaml:"depositor"`
}

func (m *MsgFundCommunityPool) Reset()         { *m = MsgFundCommunityPool{} }
func (m *MsgFundCommunityPool) String() string { return proto.CompactTextString(m) }
func (*MsgFundCommunityPool) ProtoMessage()    {}
func (*MsgFundCommunityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fddf2a8e4a90b09, []int{3}
}
func (m *MsgFundCommunityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundCommunityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundCommunityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundCommunityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundCommunityPool.Merge(m, src)
}
func (m *MsgFundCommunityPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundCommunityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundCommunityPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundCommunityPool proto.InternalMessageInfo

func (m *MsgFundCommunityPool) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgFundCommunityPool) GetDepositor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Depositor
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSetWithdrawAddress)(nil), "cosmos_sdk.x.ditribution.v1.MsgSetWithdrawAddress")
	proto.RegisterType((*MsgWithdrawDelegatorReward)(nil), "cosmos_sdk.x.ditribution.v1.MsgWithdrawDelegatorReward")
	proto.RegisterType((*MsgWithdrawValidatorCommission)(nil), "cosmos_sdk.x.ditribution.v1.MsgWithdrawValidatorCommission")
	proto.RegisterType((*MsgFundCommunityPool)(nil), "cosmos_sdk.x.ditribution.v1.MsgFundCommunityPool")
}

func init() { proto.RegisterFile("x/distribution/types/types.proto", fileDescriptor_9fddf2a8e4a90b09) }

var fileDescriptor_9fddf2a8e4a90b09 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x4d, 0x8b, 0xd3, 0x40,
	0x1c, 0xc6, 0x33, 0x11, 0x17, 0x1c, 0x05, 0xdb, 0xa8, 0x58, 0x2a, 0x24, 0x25, 0x07, 0xe9, 0x65,
	0x13, 0x57, 0x6f, 0xde, 0xb6, 0xab, 0x9e, 0x0c, 0x48, 0x85, 0x15, 0x04, 0x29, 0xd3, 0x4c, 0x48,
	0x87, 0x4d, 0xf2, 0x0f, 0x33, 0x93, 0x66, 0xb3, 0x9f, 0xc1, 0x83, 0x07, 0x8f, 0x7e, 0x02, 0x3f,
	0xc9, 0x1e, 0xf7, 0x28, 0x1e, 0xa2, 0xb4, 0xdf, 0xa0, 0x47, 0x4f, 0xb2, 0x79, 0xeb, 0x9b, 0x87,
	0xd2, 0x83, 0x7b, 0x49, 0x42, 0xe6, 0x99, 0xe7, 0xf7, 0xe4, 0x99, 0xfc, 0x71, 0xef, 0xdc, 0xa6,
	0x4c, 0x48, 0xce, 0xc6, 0x89, 0x64, 0x10, 0xd9, 0x32, 0x8b, 0x3d, 0x51, 0x5e, 0xad, 0x98, 0x83,
	0x04, 0xed, 0x89, 0x0b, 0x22, 0x04, 0x31, 0x12, 0xf4, 0xcc, 0x3a, 0xb7, 0x28, 0x6b, 0xb4, 0xd6,
	0xf4, 0xa8, 0xfb, 0x54, 0x4e, 0x18, 0xa7, 0xa3, 0x98, 0x70, 0x99, 0xd9, 0x85, 0xde, 0xf6, 0xc1,
	0x87, 0xe5, 0x53, 0x69, 0xd2, 0x6d, 0x6f, 0xf9, 0x9a, 0x9f, 0x55, 0xfc, 0xc8, 0x11, 0xfe, 0x7b,
	0x4f, 0x7e, 0x60, 0x72, 0x42, 0x39, 0x49, 0x8f, 0x29, 0xe5, 0x9e, 0x10, 0xda, 0x05, 0x6e, 0x53,
	0x2f, 0xf0, 0x7c, 0x22, 0x81, 0x8f, 0x48, 0xf9, 0xb2, 0x83, 0x7a, 0xa8, 0x7f, 0x6f, 0xe0, 0x2c,
	0x72, 0xa3, 0x93, 0x91, 0x30, 0x78, 0x69, 0x6e, 0x49, 0xcc, 0x3f, 0xb9, 0x71, 0xe8, 0x33, 0x39,
	0x49, 0xc6, 0x96, 0x0b, 0xa1, 0x5d, 0xe6, 0xae, 0x6e, 0x87, 0x82, 0x9e, 0x55, 0xf8, 0x63, 0xd7,
	0xad, 0x48, 0xc3, 0x56, 0x63, 0x52, 0xb3, 0x53, 0xdc, 0x4a, 0xab, 0x38, 0x0d, 0x5a, 0x2d, 0xd0,
	0x6f, 0x17, 0xb9, 0xf1, 0xb8, 0x44, 0x6f, 0x2a, 0xf6, 0x20, 0xdf, 0x4f, 0xd7, 0x3f, 0xda, 0xfc,
	0xaa, 0xe2, 0xae, 0x23, 0xfc, 0xba, 0x8b, 0x57, 0x75, 0xb0, 0xa1, 0x97, 0x12, 0x4e, 0x6f, 0xb4,
	0x93, 0x0b, 0xdc, 0x9e, 0x92, 0x80, 0xd1, 0x35, 0xb6, 0xba, 0xc9, 0xde, 0x92, 0xec, 0xca, 0x3e,
	0x25, 0x41, 0xc3, 0x6e, 0x4c, 0xea, 0x5a, 0xbe, 0x21, 0xac, 0xaf, 0xd4, 0x72, 0x5a, 0xaf, 0x9f,
	0x40, 0x18, 0x32, 0x21, 0x18, 0x44, 0xff, 0x8e, 0x87, 0xfe, 0x4f, 0xbc, 0x9f, 0x08, 0x3f, 0x74,
	0x84, 0xff, 0x26, 0x89, 0xe8, 0x75, 0xa2, 0x24, 0x62, 0x32, 0x7b, 0x07, 0x10, 0x68, 0x9f, 0xf0,
	0x01, 0x09, 0x21, 0x89, 0x64, 0x07, 0xf5, 0x6e, 0xf5, 0xef, 0x3e, 0x7f, 0x60, 0xad, 0x8c, 0xd1,
	0xf4, 0xc8, 0x3a, 0x01, 0x16, 0x0d, 0x9e, 0x5d, 0xe6, 0x86, 0xf2, 0xfd, 0x97, 0xd1, 0xdf, 0x21,
	0xc6, 0xf5, 0x06, 0x31, 0xac, 0x4c, 0x35, 0x17, 0xdf, 0xa1, 0x5e, 0x0c, 0x82, 0x49, 0xe0, 0xd5,
	0x51, 0xbc, 0x5e, 0xe4, 0x46, 0xab, 0xfe, 0x0d, 0xaa, 0xa5, 0x3d, 0x8e, 0x7f, 0xe9, 0x3b, 0x30,
	0x2e, 0x67, 0x3a, 0xba, 0x9a, 0xe9, 0xe8, 0xf7, 0x4c, 0x47, 0x5f, 0xe6, 0xba, 0x72, 0x35, 0xd7,
	0x95, 0x1f, 0x73, 0x5d, 0xf9, 0x78, 0xbb, 0xd8, 0x36, 0x3e, 0x28, 0x26, 0xf9, 0xc5, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbb, 0x7b, 0xe0, 0xd1, 0x45, 0x04, 0x00, 0x00,
}

func (m *MsgSetWithdrawAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetWithdrawAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetWithdrawAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawDelegatorReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawDelegatorReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawDelegatorReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawValidatorCommission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawValidatorCommission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawValidatorCommission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundCommunityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundCommunityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundCommunityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *MsgSetWithdrawAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgWithdrawDelegatorReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgWithdrawValidatorCommission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgFundCommunityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Depositor)
	if l > 0 {
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
func (m *MsgSetWithdrawAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetWithdrawAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetWithdrawAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = append(m.DelegatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DelegatorAddress == nil {
				m.DelegatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
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
			m.WithdrawAddress = append(m.WithdrawAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.WithdrawAddress == nil {
				m.WithdrawAddress = []byte{}
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
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWithdrawDelegatorReward) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawDelegatorReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawDelegatorReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = append(m.DelegatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DelegatorAddress == nil {
				m.DelegatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = append(m.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddress == nil {
				m.ValidatorAddress = []byte{}
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
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWithdrawValidatorCommission) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawValidatorCommission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawValidatorCommission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = append(m.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddress == nil {
				m.ValidatorAddress = []byte{}
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
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgFundCommunityPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFundCommunityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundCommunityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
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
