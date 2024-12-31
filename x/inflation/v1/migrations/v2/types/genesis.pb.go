// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lyfebloc/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// V2GenesisState defines the inflation module's genesis state.
type V2GenesisState struct {
	// V2Params defines all the parameters of the module.
	V2Params V2Params `protobuf:"bytes,1,opt,name=V2Params,proto3" json:"V2Params"`
	// period is the amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// epoch_identifier for inflation
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// epochs_per_period is the number of epochs after which inflation is recalculated
	EpochsPerPeriod int64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
	// skipped_epochs is the number of epochs that have passed while inflation is disabled
	SkippedEpochs uint64 `protobuf:"varint,5,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
}

func (m *V2GenesisState) Reset()         { *m = V2GenesisState{} }
func (m *V2GenesisState) String() string { return proto.CompactTextString(m) }
func (*V2GenesisState) ProtoMessage()    {}
func (*V2GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{0}
}

func (m *V2GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2GenesisState.Merge(m, src)
}

func (m *V2GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *V2GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_V2GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_V2GenesisState proto.InternalMessageInfo

func (m *V2GenesisState) GetV2Params() V2Params {
	if m != nil {
		return m.V2Params
	}
	return V2Params{}
}

func (m *V2GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *V2GenesisState) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *V2GenesisState) GetEpochsPerPeriod() int64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func (m *V2GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

// V2Params holds parameters for the inflation module.
type V2Params struct {
	// mint_denom specifies the type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// exponential_calculation takes in the variables to calculate exponential inflation
	ExponentialCalculation V2ExponentialCalculation `protobuf:"bytes,2,opt,name=exponential_calculation,json=exponentialCalculation,proto3" json:"exponential_calculation"`
	// inflation_distribution of the minted denom
	InflationDistribution V2InflationDistribution `protobuf:"bytes,3,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// enable_inflation is the parameter that enables inflation and halts increasing the skipped_epochs
	EnableInflation bool `protobuf:"varint,4,opt,name=enable_inflation,json=enableInflation,proto3" json:"enable_inflation,omitempty"`
}

func (m *V2Params) Reset()         { *m = V2Params{} }
func (m *V2Params) String() string { return proto.CompactTextString(m) }
func (*V2Params) ProtoMessage()    {}
func (*V2Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{1}
}

func (m *V2Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2Params.Merge(m, src)
}

func (m *V2Params) XXX_Size() int {
	return m.Size()
}

func (m *V2Params) XXX_DiscardUnknown() {
	xxx_messageInfo_V2Params.DiscardUnknown(m)
}

var xxx_messageInfo_V2Params proto.InternalMessageInfo

func (m *V2Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *V2Params) GetExponentialCalculation() V2ExponentialCalculation {
	if m != nil {
		return m.ExponentialCalculation
	}
	return V2ExponentialCalculation{}
}

func (m *V2Params) GetInflationDistribution() V2InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return V2InflationDistribution{}
}

func (m *V2Params) GetEnableInflation() bool {
	if m != nil {
		return m.EnableInflation
	}
	return false
}

func init() {
	proto.RegisterType((*V2GenesisState)(nil), "lyfebloc.inflation.v1.V2GenesisState")
	proto.RegisterType((*V2Params)(nil), "lyfebloc.inflation.v1.V2Params")
}

func init() { proto.RegisterFile("lyfebloc/inflation/v1/genesis.proto", fileDescriptor_1cb8eee530db1235) }

var fileDescriptor_1cb8eee530db1235 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x36, 0xc1, 0x22, 0x5b, 0xa0, 0xb0, 0x82, 0x60, 0x45, 0xc2, 0x58, 0x91, 0x90, 0xdc,
	0x0a, 0xd9, 0xa4, 0x5c, 0x38, 0x97, 0x16, 0xd4, 0x5b, 0x64, 0x6e, 0x5c, 0x56, 0xfe, 0x98, 0xb8,
	0x2b, 0xec, 0xdd, 0x95, 0x77, 0x63, 0x95, 0x7f, 0xc1, 0x9f, 0xe0, 0xbf, 0xf4, 0xd8, 0x23, 0xa7,
	0x0a, 0x25, 0x7f, 0x04, 0x79, 0xd7, 0x38, 0x95, 0xea, 0x8b, 0xe5, 0x7d, 0xef, 0xcd, 0x7b, 0x33,
	0xa3, 0xc1, 0x3e, 0x34, 0x95, 0x50, 0x11, 0xe3, 0xeb, 0x32, 0xd1, 0x4c, 0xf0, 0xa8, 0x59, 0x46,
	0x05, 0x70, 0x50, 0x4c, 0x85, 0xb2, 0x16, 0x5a, 0x10, 0x62, 0x14, 0x61, 0xaf, 0x08, 0x9b, 0xe5,
	0xfc, 0x65, 0x21, 0x0a, 0x61, 0xe8, 0xa8, 0xfd, 0xb3, 0xca, 0xf9, 0x62, 0xc0, 0x6b, 0x5f, 0x66,
	0x34, 0x8b, 0x3b, 0x84, 0x9f, 0x7c, 0xb5, 0xfe, 0xdf, 0x74, 0xa2, 0x81, 0x7c, 0xc2, 0x8e, 0x4c,
	0xea, 0xa4, 0x52, 0x2e, 0xf2, 0x51, 0x70, 0x78, 0x3a, 0x0f, 0x1f, 0xe6, 0x85, 0x2b, 0xa3, 0x38,
	0x9b, 0xdc, 0xdc, 0xbd, 0x1d, 0xc5, 0x9d, 0x9e, 0xcc, 0xb0, 0x23, 0xa1, 0x66, 0x22, 0x77, 0x0f,
	0x7c, 0x14, 0x4c, 0xe2, 0xee, 0x45, 0x8e, 0xf1, 0x73, 0x90, 0x22, 0xbb, 0xa2, 0x2c, 0x07, 0xae,
	0xd9, 0x9a, 0x41, 0xed, 0x8e, 0x7d, 0x14, 0x4c, 0xe3, 0x23, 0x83, 0x5f, 0xf6, 0x30, 0x39, 0xc1,
	0x2f, 0x0c, 0xa4, 0xa8, 0x84, 0x9a, 0x76, 0x6e, 0x13, 0x1f, 0x05, 0xe3, 0x4e, 0xab, 0x56, 0x50,
	0xaf, 0xac, 0xed, 0x3b, 0xfc, 0x4c, 0xfd, 0x60, 0x52, 0x42, 0x4e, 0x2d, 0xe5, 0x3e, 0x32, 0xb1,
	0x4f, 0x3b, 0xf4, 0xc2, 0x80, 0x8b, 0xdf, 0x07, 0xd8, 0xb1, 0xed, 0x92, 0x37, 0x18, 0x57, 0x8c,
	0x6b, 0x9a, 0x03, 0x17, 0x95, 0x19, 0x6f, 0x1a, 0x4f, 0x5b, 0xe4, 0xbc, 0x05, 0x08, 0xc3, 0xaf,
	0xe1, 0x5a, 0x0a, 0xde, 0x76, 0x93, 0x94, 0x34, 0x4b, 0xca, 0x6c, 0x63, 0x47, 0x36, 0x03, 0x1d,
	0x9e, 0x9e, 0x0c, 0xad, 0xe2, 0x62, 0x5f, 0xf2, 0x79, 0x5f, 0xd1, 0xad, 0x66, 0x06, 0x83, 0x2c,
	0x59, 0xe3, 0x59, 0x6f, 0x42, 0x73, 0xa6, 0x74, 0xcd, 0xd2, 0x8d, 0x49, 0x1a, 0x9b, 0xa4, 0xe3,
	0xa1, 0xa4, 0xcb, 0xff, 0x8f, 0xf3, 0x7b, 0x05, 0x5d, 0xd0, 0x2b, 0x36, 0x44, 0x9a, 0xd5, 0xf3,
	0x24, 0x2d, 0x81, 0xf6, 0xbc, 0x59, 0xe7, 0xe3, 0xf8, 0xc8, 0xe2, 0xbd, 0xe7, 0xd9, 0x97, 0x9b,
	0xad, 0x87, 0x6e, 0xb7, 0x1e, 0xfa, 0xbb, 0xf5, 0xd0, 0xaf, 0x9d, 0x37, 0xba, 0xdd, 0x79, 0xa3,
	0x3f, 0x3b, 0x6f, 0xf4, 0xfd, 0x7d, 0xc1, 0xf4, 0xd5, 0x26, 0x0d, 0x33, 0x51, 0x45, 0xf6, 0xa2,
	0xec, 0xb7, 0x59, 0x7e, 0x88, 0xae, 0xef, 0x5d, 0x97, 0xfe, 0x29, 0x41, 0xa5, 0x8e, 0xb9, 0xab,
	0x8f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xbd, 0xac, 0x94, 0xc9, 0x02, 0x00, 0x00,
}

func (m *V2GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.V2Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *V2Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableInflation {
		i--
		if m.EnableInflation {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ExponentialCalculation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *V2GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.V2Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	return n
}

func (m *V2Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ExponentialCalculation.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EnableInflation {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V2GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V2GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.V2Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsPerPeriod", wireType)
			}
			m.EpochsPerPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsPerPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedEpochs", wireType)
			}
			m.SkippedEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *V2Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V2Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentialCalculation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExponentialCalculation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableInflation", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableInflation = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
