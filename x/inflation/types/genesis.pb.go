// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	sdkmath "cosmossdk.io/math"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// period is the amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// skipped_epochs is the number of epochs that have passed while inflation is
	// disabled
	SkippedEpochs uint64 `protobuf:"varint,3,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d00e2bb98c08f74, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

// Params holds parameters for the inflation module.
type Params struct {
	// inflation_enabled is the parameter that enables inflation and halts
	// increasing the skipped_epochs
	InflationEnabled bool `protobuf:"varint,1,opt,name=inflation_enabled,json=inflationEnabled,proto3" json:"inflation_enabled,omitempty"`
	// polynomial_factors takes in the variables to calculate polynomial
	// inflation
	PolynomialFactors []sdkmath.LegacyDec `protobuf:"bytes,2,rep,name=polynomial_factors,json=polynomialFactors,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"polynomial_factors"`
	// inflation_distribution of the minted denom
	InflationDistribution InflationDistribution `protobuf:"bytes,3,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// epochs_per_period is the number of epochs that must pass before a new
	// period is created
	EpochsPerPeriod uint64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
	// periods_per_year is the number of periods that occur in a year
	PeriodsPerYear uint64 `protobuf:"varint,5,opt,name=periods_per_year,json=periodsPerYear,proto3" json:"periods_per_year,omitempty"`
	// max_period is the maximum number of periods that have inflation being
	// paid off. After this period, inflation will be disabled.
	MaxPeriod uint64 `protobuf:"varint,6,opt,name=max_period,json=maxPeriod,proto3" json:"max_period,omitempty"`
	// has_inflation_started is the parameter that indicates if inflation has
	// started. It's set to false at the starts, and stays at true when we toggle
	// inflation on. It's used to track num skipped epochs
	HasInflationStarted bool `protobuf:"varint,7,opt,name=has_inflation_started,json=hasInflationStarted,proto3" json:"has_inflation_started,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d00e2bb98c08f74, []int{1}
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

func (m *Params) GetInflationEnabled() bool {
	if m != nil {
		return m.InflationEnabled
	}
	return false
}

func (m *Params) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func (m *Params) GetEpochsPerPeriod() uint64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func (m *Params) GetPeriodsPerYear() uint64 {
	if m != nil {
		return m.PeriodsPerYear
	}
	return 0
}

func (m *Params) GetMaxPeriod() uint64 {
	if m != nil {
		return m.MaxPeriod
	}
	return 0
}

func (m *Params) GetHasInflationStarted() bool {
	if m != nil {
		return m.HasInflationStarted
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nibiru.inflation.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "nibiru.inflation.v1.Params")
}

func init() { proto.RegisterFile("nibiru/inflation/v1/genesis.proto", fileDescriptor_2d00e2bb98c08f74) }

var fileDescriptor_2d00e2bb98c08f74 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x9b, 0x60, 0xe8, 0x16, 0x4a, 0xb3, 0xa5, 0x95, 0x55, 0x84, 0x1b, 0x8a, 0x40, 0x51,
	0x11, 0xb6, 0x6a, 0x4e, 0x5c, 0x43, 0x0b, 0x42, 0x42, 0x28, 0x4a, 0x4f, 0x20, 0x21, 0x6b, 0x6d,
	0x4f, 0xed, 0x55, 0x63, 0xef, 0x6a, 0x77, 0x13, 0x25, 0x6f, 0xc0, 0x91, 0xa7, 0xe0, 0x59, 0x7a,
	0xec, 0x11, 0x71, 0xa8, 0x50, 0xf2, 0x22, 0x55, 0x76, 0x37, 0x3f, 0x07, 0x9f, 0x76, 0x67, 0xbe,
	0x6f, 0xbe, 0x99, 0xf9, 0x34, 0xe8, 0x65, 0x45, 0x13, 0x2a, 0x46, 0x21, 0xad, 0xae, 0x86, 0x44,
	0x51, 0x56, 0x85, 0xe3, 0xb3, 0x30, 0x87, 0x0a, 0x24, 0x95, 0x01, 0x17, 0x4c, 0x31, 0xbc, 0x6f,
	0x28, 0xc1, 0x8a, 0x12, 0x8c, 0xcf, 0x8e, 0x9e, 0xe5, 0x2c, 0x67, 0x1a, 0x0f, 0x17, 0x3f, 0x43,
	0x3d, 0x7a, 0x55, 0xa7, 0xb6, 0xae, 0xd3, 0xa4, 0x93, 0x5f, 0x0e, 0x7a, 0xfc, 0xd9, 0x74, 0xb8,
	0x54, 0x44, 0x01, 0xfe, 0x80, 0x5c, 0x4e, 0x04, 0x29, 0xa5, 0xe7, 0x74, 0x9c, 0xee, 0x4e, 0xf4,
	0x3c, 0xa8, 0xe9, 0x18, 0xf4, 0x35, 0xa5, 0xd7, 0xba, 0xb9, 0x3b, 0x6e, 0x0c, 0x6c, 0x01, 0x3e,
	0x44, 0x2e, 0x07, 0x41, 0x59, 0xe6, 0x6d, 0x75, 0x9c, 0x6e, 0x6b, 0x60, 0x23, 0xfc, 0x1a, 0xed,
	0xca, 0x6b, 0xca, 0x39, 0x64, 0x31, 0x70, 0x96, 0x16, 0xd2, 0x6b, 0x6a, 0xfc, 0x89, 0xcd, 0x5e,
	0xe8, 0xe4, 0xc9, 0x9f, 0x26, 0x72, 0x8d, 0x2e, 0x7e, 0x8b, 0xda, 0xab, 0x76, 0x31, 0x54, 0x24,
	0x19, 0x42, 0xa6, 0xe7, 0x79, 0x34, 0xd8, 0x5b, 0x01, 0x17, 0x26, 0x8f, 0x7f, 0x22, 0xcc, 0xd9,
	0x70, 0x5a, 0xb1, 0x92, 0x92, 0x61, 0x7c, 0x45, 0x52, 0xc5, 0x84, 0xf4, 0xb6, 0x3a, 0xcd, 0xee,
	0x76, 0x2f, 0x58, 0x0c, 0xf8, 0xef, 0xee, 0xf8, 0x4d, 0x4e, 0x55, 0x31, 0x4a, 0x82, 0x94, 0x95,
	0x61, 0xca, 0x64, 0xc9, 0xa4, 0x7d, 0xde, 0xc9, 0xec, 0x3a, 0x54, 0x53, 0x0e, 0x32, 0x38, 0x87,
	0x74, 0xd0, 0x5e, 0x2b, 0x7d, 0x32, 0x42, 0x38, 0x47, 0x87, 0xeb, 0x59, 0x32, 0x2a, 0x95, 0xa0,
	0xc9, 0x68, 0x11, 0xe8, 0x2d, 0x76, 0xa2, 0xd3, 0x5a, 0x83, 0xbe, 0x2c, 0x83, 0xf3, 0x8d, 0x0a,
	0xeb, 0xd7, 0x01, 0xad, 0x03, 0xf1, 0x29, 0x6a, 0x1b, 0x7b, 0x62, 0x0e, 0x22, 0xb6, 0x4e, 0xb6,
	0xb4, 0x53, 0x4f, 0x0d, 0xd0, 0x07, 0xd1, 0x37, 0x96, 0x76, 0xd1, 0x9e, 0x21, 0x18, 0xf2, 0x14,
	0x88, 0xf0, 0x1e, 0x68, 0xea, 0xae, 0xcd, 0xf7, 0x41, 0x7c, 0x07, 0x22, 0xf0, 0x0b, 0x84, 0x4a,
	0x32, 0x59, 0xca, 0xb9, 0x9a, 0xb3, 0x5d, 0x92, 0x89, 0x15, 0x8a, 0xd0, 0x41, 0x41, 0x64, 0xbc,
	0xde, 0x50, 0x2a, 0x22, 0x14, 0x64, 0xde, 0x43, 0xed, 0xf6, 0x7e, 0x41, 0xe4, 0x6a, 0x95, 0x4b,
	0x03, 0xf5, 0xbe, 0xde, 0xcc, 0x7c, 0xe7, 0x76, 0xe6, 0x3b, 0xff, 0x67, 0xbe, 0xf3, 0x7b, 0xee,
	0x37, 0x6e, 0xe7, 0x7e, 0xe3, 0xef, 0xdc, 0x6f, 0xfc, 0x88, 0x36, 0x6c, 0xfe, 0xa6, 0x5d, 0xf9,
	0x58, 0x10, 0x5a, 0x85, 0xf6, 0x12, 0xc7, 0x51, 0x38, 0xd9, 0x38, 0x47, 0x6d, 0x7b, 0xe2, 0xea,
	0x43, 0x7c, 0x7f, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xaa, 0x78, 0xac, 0xfd, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x18
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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
	if m.HasInflationStarted {
		i--
		if m.HasInflationStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.MaxPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxPeriod))
		i--
		dAtA[i] = 0x30
	}
	if m.PeriodsPerYear != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PeriodsPerYear))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
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
	if len(m.PolynomialFactors) > 0 {
		for iNdEx := len(m.PolynomialFactors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.PolynomialFactors[iNdEx].Size()
				i -= size
				if _, err := m.PolynomialFactors[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.InflationEnabled {
		i--
		if m.InflationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InflationEnabled {
		n += 2
	}
	if len(m.PolynomialFactors) > 0 {
		for _, e := range m.PolynomialFactors {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	if m.PeriodsPerYear != 0 {
		n += 1 + sovGenesis(uint64(m.PeriodsPerYear))
	}
	if m.MaxPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.MaxPeriod))
	}
	if m.HasInflationStarted {
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationEnabled", wireType)
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
			m.InflationEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolynomialFactors", wireType)
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
			var v sdkmath.LegacyDec
			m.PolynomialFactors = append(m.PolynomialFactors, v)
			if err := m.PolynomialFactors[len(m.PolynomialFactors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				m.EpochsPerPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodsPerYear", wireType)
			}
			m.PeriodsPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeriodsPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPeriod", wireType)
			}
			m.MaxPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasInflationStarted", wireType)
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
			m.HasInflationStarted = bool(v != 0)
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
