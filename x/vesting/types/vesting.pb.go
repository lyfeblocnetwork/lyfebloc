// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lyfebloc/vesting/v2/vesting.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	github_com_cosmos_cosmos_sdk_x_auth_vesting_types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// ClawbackVestingAccount implements the VestingAccount interface. It provides
// an account that can hold contributions subject to "lockup" (like a
// PeriodicVestingAccount), or vesting which is subject to clawback
// of unvested tokens, or a combination (tokens vest, but are still locked).
type ClawbackVestingAccount struct {
	// base_vesting_account implements the VestingAccount interface. It contains
	// all the necessary fields needed for any vesting account implementation
	*types.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	// funder_address specifies the account which can perform clawback
	FunderAddress string `protobuf:"bytes,2,opt,name=funder_address,json=funderAddress,proto3" json:"funder_address,omitempty"`
	// start_time defines the time at which the vesting period begins
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// lockup_periods defines the unlocking schedule relative to the start_time
	LockupPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,4,rep,name=lockup_periods,json=lockupPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"lockup_periods"`
	// vesting_periods defines the vesting schedule relative to the start_time
	VestingPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,5,rep,name=vesting_periods,json=vestingPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"vesting_periods"`
}

func (m *ClawbackVestingAccount) Reset()      { *m = ClawbackVestingAccount{} }
func (*ClawbackVestingAccount) ProtoMessage() {}
func (*ClawbackVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0001d894a8ee0c72, []int{0}
}
func (m *ClawbackVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClawbackVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClawbackVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClawbackVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClawbackVestingAccount.Merge(m, src)
}
func (m *ClawbackVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ClawbackVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ClawbackVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ClawbackVestingAccount proto.InternalMessageInfo

// ClawbackProposal is a gov Content type to clawback funds
// from a vesting account that has this functionality enabled.
type ClawbackProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// address is the vesting account address
	// to clawback the funds from
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// destination_address is the address that will receive
	// the clawbacked funds from the given vesting account. When
	// empty, proposal will return the coins to the vesting
	// account funder.
	DestinationAddress string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
}

func (m *ClawbackProposal) Reset()         { *m = ClawbackProposal{} }
func (m *ClawbackProposal) String() string { return proto.CompactTextString(m) }
func (*ClawbackProposal) ProtoMessage()    {}
func (*ClawbackProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_0001d894a8ee0c72, []int{1}
}
func (m *ClawbackProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClawbackProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClawbackProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClawbackProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClawbackProposal.Merge(m, src)
}
func (m *ClawbackProposal) XXX_Size() int {
	return m.Size()
}
func (m *ClawbackProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ClawbackProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ClawbackProposal proto.InternalMessageInfo

func (m *ClawbackProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ClawbackProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ClawbackProposal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClawbackProposal) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*ClawbackVestingAccount)(nil), "lyfebloc.vesting.v2.ClawbackVestingAccount")
	proto.RegisterType((*ClawbackProposal)(nil), "lyfebloc.vesting.v2.ClawbackProposal")
}

func init() { proto.RegisterFile("lyfebloc/vesting/v2/vesting.proto", fileDescriptor_0001d894a8ee0c72) }

var fileDescriptor_0001d894a8ee0c72 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xcd, 0xd8, 0xae, 0xd2, 0x29, 0xad, 0xeb, 0x58, 0xa4, 0xf4, 0x90, 0x94, 0x45, 0xa1, 0x16,
	0xcc, 0xb8, 0x15, 0x04, 0xf7, 0xb6, 0x51, 0xc4, 0xe3, 0x52, 0xc4, 0x83, 0x97, 0x30, 0x49, 0x66,
	0xb3, 0xa1, 0x49, 0x26, 0x64, 0x26, 0x75, 0xfd, 0x07, 0x22, 0x88, 0x7b, 0x14, 0xbc, 0xf4, 0x28,
	0x9e, 0xf6, 0x5f, 0xb8, 0xc7, 0x1e, 0x3d, 0xed, 0x4a, 0x7b, 0x58, 0x7f, 0x86, 0x64, 0x66, 0xb2,
	0xa6, 0xa8, 0x57, 0x2f, 0xc9, 0xcc, 0x7b, 0x5f, 0xbe, 0x79, 0x6f, 0xbe, 0x17, 0x68, 0xd2, 0x79,
	0xc2, 0x38, 0x9e, 0x53, 0x2e, 0xa2, 0x34, 0xc4, 0xf3, 0x49, 0xb5, 0xb4, 0xb3, 0x9c, 0x09, 0x86,
	0xb6, 0x25, 0x6f, 0x57, 0xe0, 0x7c, 0x32, 0xb8, 0x45, 0x92, 0x28, 0x65, 0x58, 0x3e, 0x55, 0xd1,
	0xe0, 0xae, 0xcf, 0xf8, 0x46, 0x97, 0x5d, 0x8f, 0x0a, 0xb2, 0xbb, 0xd9, 0x6a, 0xd0, 0x0b, 0x59,
	0xc8, 0xe4, 0x12, 0x97, 0x2b, 0x8d, 0x5a, 0x21, 0x63, 0x61, 0x4c, 0xb1, 0xdc, 0x79, 0xc5, 0x21,
	0x16, 0x51, 0x42, 0xb9, 0x20, 0x49, 0xa6, 0x0a, 0x76, 0xbe, 0x35, 0xe1, 0x9d, 0xa7, 0x31, 0x79,
	0xe3, 0x11, 0x7f, 0xf6, 0x4a, 0x35, 0xdc, 0xf7, 0x7d, 0x56, 0xa4, 0x02, 0x79, 0xb0, 0xe7, 0x11,
	0x4e, 0x5d, 0x7d, 0x8e, 0x4b, 0x14, 0xde, 0x07, 0x43, 0x30, 0x6a, 0x4f, 0xc6, 0xb6, 0x92, 0xf5,
	0x5b, 0xbc, 0x92, 0x65, 0x3b, 0x84, 0xd3, 0xcd, 0x4e, 0x4e, 0x73, 0x79, 0x6e, 0x81, 0x29, 0xf2,
	0xfe, 0x60, 0xd0, 0x3d, 0xd8, 0x3d, 0x2c, 0xd2, 0x80, 0xe6, 0x2e, 0x09, 0x82, 0x9c, 0x72, 0xde,
	0xbf, 0x36, 0x04, 0xa3, 0xd6, 0xb4, 0xa3, 0xd0, 0x7d, 0x05, 0xa2, 0x17, 0x10, 0x72, 0x41, 0x72,
	0xe1, 0x96, 0xf2, 0xfb, 0x0d, 0x29, 0x60, 0x60, 0x2b, 0x6f, 0x76, 0xe5, 0xcd, 0x7e, 0x59, 0x79,
	0x73, 0x3a, 0x67, 0xe7, 0x96, 0x71, 0x72, 0x61, 0x81, 0x2f, 0x97, 0xa7, 0x63, 0x30, 0x6d, 0xc9,
	0x8f, 0x4b, 0x1a, 0x7d, 0x00, 0xb0, 0x1b, 0x33, 0x7f, 0x56, 0x64, 0x6e, 0x46, 0xf3, 0x88, 0x05,
	0xbc, 0xdf, 0x1c, 0x36, 0x46, 0xed, 0x89, 0xf9, 0x2f, 0x3f, 0x07, 0xb2, 0xcc, 0x79, 0x5e, 0xb6,
	0xfc, 0x7a, 0x61, 0x3d, 0x09, 0x23, 0x71, 0x54, 0x78, 0xb6, 0xcf, 0x12, 0xac, 0x07, 0xa3, 0x5e,
	0x0f, 0x78, 0x30, 0xc3, 0xc7, 0x98, 0x14, 0xe2, 0xe8, 0x6a, 0x54, 0xe2, 0x6d, 0x46, 0xb9, 0xee,
	0xc0, 0x95, 0x96, 0x8e, 0x3a, 0x5d, 0x63, 0xe8, 0x23, 0x80, 0x37, 0xab, 0x0b, 0xae, 0x04, 0x6d,
	0xfd, 0x57, 0x41, 0x5d, 0xcd, 0x69, 0x70, 0xef, 0xf1, 0xbb, 0x85, 0x65, 0x7c, 0x5a, 0x58, 0xc6,
	0xfb, 0xcb, 0xd3, 0xf1, 0x7d, 0x15, 0xe0, 0xe3, 0x7a, 0x84, 0xff, 0x1e, 0x97, 0x9d, 0xcf, 0x00,
	0x6e, 0x57, 0xd4, 0x41, 0xce, 0x32, 0xc6, 0x49, 0x8c, 0x7a, 0x70, 0x4b, 0x44, 0x22, 0xa6, 0x32,
	0x34, 0xad, 0xa9, 0xda, 0xa0, 0x21, 0x6c, 0x07, 0x94, 0xfb, 0x79, 0x94, 0x89, 0x88, 0xa5, 0x7a,
	0xe4, 0x75, 0x08, 0xf5, 0xe1, 0x8d, 0x2a, 0x10, 0x0d, 0xc9, 0x56, 0x5b, 0x84, 0xe1, 0xed, 0x40,
	0x1e, 0x4c, 0xca, 0xc2, 0xab, 0xd8, 0x34, 0x65, 0x15, 0xaa, 0x51, 0x3a, 0x3b, 0x7b, 0xcd, 0x9f,
	0x0b, 0xcb, 0x70, 0x9e, 0x9d, 0xad, 0x4c, 0xb0, 0x5c, 0x99, 0xe0, 0xc7, 0xca, 0x04, 0x27, 0x6b,
	0xd3, 0x58, 0xae, 0x4d, 0xe3, 0xfb, 0xda, 0x34, 0x5e, 0x8f, 0x6b, 0xf7, 0xa7, 0xdc, 0xea, 0x9f,
	0x76, 0xf2, 0xb0, 0xe6, 0x5b, 0x5e, 0x9c, 0x77, 0x5d, 0x66, 0xed, 0xd1, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x47, 0xd1, 0xdd, 0x7b, 0xd8, 0x03, 0x00, 0x00,
}

func (m *ClawbackVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClawbackVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClawbackVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LockupPeriods) > 0 {
		for iNdEx := len(m.LockupPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockupPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVesting(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.FunderAddress) > 0 {
		i -= len(m.FunderAddress)
		copy(dAtA[i:], m.FunderAddress)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.FunderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClawbackProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClawbackProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClawbackProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClawbackVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.FunderAddress)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovVesting(uint64(l))
	if len(m.LockupPeriods) > 0 {
		for _, e := range m.LockupPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *ClawbackProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClawbackVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ClawbackVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClawbackVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockupPeriods = append(m.LockupPeriods, types.Period{})
			if err := m.LockupPeriods[len(m.LockupPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, types.Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *ClawbackProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ClawbackProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClawbackProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)