// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/devgas/v1/event.proto

package types

import (
	fmt "fmt"
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

// ABCI event emitted when a deployer registers a contract to receive fee
// sharing payouts, specifying the deployer, contract, and withdrawer addresses.
type EventRegisterDevGas struct {
	// deployer is the addess of the account that registered the smart contract to
	// receive dev gas royalties.
	Deployer string `protobuf:"bytes,1,opt,name=deployer,proto3" json:"deployer,omitempty"`
	// Address of the smart contract. This identifies the specific contract
	// that will receive fee sharing payouts.
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// The address that will receive the fee sharing payouts for the registered
	// contract. This could be the deployer address or a separate withdrawer
	// address specified.
	Withdrawer string `protobuf:"bytes,3,opt,name=withdrawer,proto3" json:"withdrawer,omitempty"`
}

func (m *EventRegisterDevGas) Reset()         { *m = EventRegisterDevGas{} }
func (m *EventRegisterDevGas) String() string { return proto.CompactTextString(m) }
func (*EventRegisterDevGas) ProtoMessage()    {}
func (*EventRegisterDevGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ce94d3a226edf, []int{0}
}
func (m *EventRegisterDevGas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegisterDevGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegisterDevGas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegisterDevGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegisterDevGas.Merge(m, src)
}
func (m *EventRegisterDevGas) XXX_Size() int {
	return m.Size()
}
func (m *EventRegisterDevGas) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegisterDevGas.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegisterDevGas proto.InternalMessageInfo

func (m *EventRegisterDevGas) GetDeployer() string {
	if m != nil {
		return m.Deployer
	}
	return ""
}

func (m *EventRegisterDevGas) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *EventRegisterDevGas) GetWithdrawer() string {
	if m != nil {
		return m.Withdrawer
	}
	return ""
}

// ABCI event emitted when a deployer cancels fee sharing for a contract,
// specifying the deployer and contract addresses.
type EventCancelDevGas struct {
	// deployer is the addess of the account that registered the smart contract to
	// receive dev gas royalties.
	Deployer string `protobuf:"bytes,1,opt,name=deployer,proto3" json:"deployer,omitempty"`
	// Address of the smart contract. This identifies the specific contract
	// that will receive fee sharing payouts.
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *EventCancelDevGas) Reset()         { *m = EventCancelDevGas{} }
func (m *EventCancelDevGas) String() string { return proto.CompactTextString(m) }
func (*EventCancelDevGas) ProtoMessage()    {}
func (*EventCancelDevGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ce94d3a226edf, []int{1}
}
func (m *EventCancelDevGas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCancelDevGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCancelDevGas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCancelDevGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCancelDevGas.Merge(m, src)
}
func (m *EventCancelDevGas) XXX_Size() int {
	return m.Size()
}
func (m *EventCancelDevGas) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCancelDevGas.DiscardUnknown(m)
}

var xxx_messageInfo_EventCancelDevGas proto.InternalMessageInfo

func (m *EventCancelDevGas) GetDeployer() string {
	if m != nil {
		return m.Deployer
	}
	return ""
}

func (m *EventCancelDevGas) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

// ABCI event emitted when a deployer updates the fee sharing registration for a
// contract, specifying updated deployer, contract, and/or withdrawer addresses.
type EventUpdateDevGas struct {
	// deployer is the addess of the account that registered the smart contract to
	// receive dev gas royalties.
	Deployer string `protobuf:"bytes,1,opt,name=deployer,proto3" json:"deployer,omitempty"`
	// Address of the smart contract. This identifies the specific contract
	// that will receive fee sharing payouts.
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// The address that will receive the fee sharing payouts for the registered
	// contract. This could be the deployer address or a separate withdrawer
	// address specified.
	Withdrawer string `protobuf:"bytes,3,opt,name=withdrawer,proto3" json:"withdrawer,omitempty"`
}

func (m *EventUpdateDevGas) Reset()         { *m = EventUpdateDevGas{} }
func (m *EventUpdateDevGas) String() string { return proto.CompactTextString(m) }
func (*EventUpdateDevGas) ProtoMessage()    {}
func (*EventUpdateDevGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ce94d3a226edf, []int{2}
}
func (m *EventUpdateDevGas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateDevGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateDevGas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateDevGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateDevGas.Merge(m, src)
}
func (m *EventUpdateDevGas) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateDevGas) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateDevGas.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateDevGas proto.InternalMessageInfo

func (m *EventUpdateDevGas) GetDeployer() string {
	if m != nil {
		return m.Deployer
	}
	return ""
}

func (m *EventUpdateDevGas) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *EventUpdateDevGas) GetWithdrawer() string {
	if m != nil {
		return m.Withdrawer
	}
	return ""
}

// ABCI event emitted when fee sharing payouts are made, containing details on
// the payouts in JSON format.
type EventPayoutDevGas struct {
	Payouts string `protobuf:"bytes,1,opt,name=payouts,proto3" json:"payouts,omitempty"`
}

func (m *EventPayoutDevGas) Reset()         { *m = EventPayoutDevGas{} }
func (m *EventPayoutDevGas) String() string { return proto.CompactTextString(m) }
func (*EventPayoutDevGas) ProtoMessage()    {}
func (*EventPayoutDevGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ce94d3a226edf, []int{3}
}
func (m *EventPayoutDevGas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPayoutDevGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPayoutDevGas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPayoutDevGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPayoutDevGas.Merge(m, src)
}
func (m *EventPayoutDevGas) XXX_Size() int {
	return m.Size()
}
func (m *EventPayoutDevGas) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPayoutDevGas.DiscardUnknown(m)
}

var xxx_messageInfo_EventPayoutDevGas proto.InternalMessageInfo

func (m *EventPayoutDevGas) GetPayouts() string {
	if m != nil {
		return m.Payouts
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRegisterDevGas)(nil), "nibiru.devgas.v1.EventRegisterDevGas")
	proto.RegisterType((*EventCancelDevGas)(nil), "nibiru.devgas.v1.EventCancelDevGas")
	proto.RegisterType((*EventUpdateDevGas)(nil), "nibiru.devgas.v1.EventUpdateDevGas")
	proto.RegisterType((*EventPayoutDevGas)(nil), "nibiru.devgas.v1.EventPayoutDevGas")
}

func init() { proto.RegisterFile("nibiru/devgas/v1/event.proto", fileDescriptor_dd3ce94d3a226edf) }

var fileDescriptor_dd3ce94d3a226edf = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xcb, 0x4c, 0xca,
	0x2c, 0x2a, 0xd5, 0x4f, 0x49, 0x2d, 0x4b, 0x4f, 0x2c, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xc8, 0xea, 0x41, 0x64, 0xf5,
	0xca, 0x0c, 0x95, 0x72, 0xb9, 0x84, 0x5d, 0x41, 0x0a, 0x82, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52,
	0x8b, 0x5c, 0x52, 0xcb, 0xdc, 0x13, 0x8b, 0x85, 0xa4, 0xb8, 0x38, 0x52, 0x52, 0x0b, 0x72, 0xf2,
	0x2b, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x72, 0x7e,
	0x5e, 0x49, 0x51, 0x62, 0x72, 0x89, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x17, 0x92, 0xe3, 0xe2, 0x2a,
	0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a, 0x2c, 0x4f, 0x2d, 0x92, 0x60, 0x06, 0xcb, 0x22, 0x89, 0x28,
	0x79, 0x73, 0x09, 0x82, 0xad, 0x73, 0x4e, 0xcc, 0x4b, 0x4e, 0xcd, 0xa1, 0xcc, 0x32, 0xa5, 0x6c,
	0xa8, 0x61, 0xa1, 0x05, 0x29, 0x89, 0x25, 0xa9, 0x34, 0x76, 0xb9, 0x2e, 0xd4, 0xb2, 0x80, 0xc4,
	0xca, 0xfc, 0xd2, 0x12, 0xa8, 0x65, 0x12, 0x5c, 0xec, 0x05, 0x60, 0x7e, 0x31, 0xd4, 0x2e, 0x18,
	0xd7, 0xc9, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xfd, 0xc0, 0xd1, 0xe1, 0x9c, 0x91, 0x98,
	0x99, 0xa7, 0x0f, 0x8d, 0xb8, 0x32, 0x23, 0xfd, 0x0a, 0xa4, 0xd8, 0x2b, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0x47, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x57, 0x64, 0x78, 0xde,
	0x01, 0x00, 0x00,
}

func (m *EventRegisterDevGas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegisterDevGas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegisterDevGas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Withdrawer) > 0 {
		i -= len(m.Withdrawer)
		copy(dAtA[i:], m.Withdrawer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Withdrawer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Deployer) > 0 {
		i -= len(m.Deployer)
		copy(dAtA[i:], m.Deployer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Deployer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCancelDevGas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCancelDevGas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCancelDevGas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Deployer) > 0 {
		i -= len(m.Deployer)
		copy(dAtA[i:], m.Deployer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Deployer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateDevGas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateDevGas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateDevGas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Withdrawer) > 0 {
		i -= len(m.Withdrawer)
		copy(dAtA[i:], m.Withdrawer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Withdrawer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Deployer) > 0 {
		i -= len(m.Deployer)
		copy(dAtA[i:], m.Deployer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Deployer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventPayoutDevGas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPayoutDevGas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPayoutDevGas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payouts) > 0 {
		i -= len(m.Payouts)
		copy(dAtA[i:], m.Payouts)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Payouts)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRegisterDevGas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Deployer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Withdrawer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventCancelDevGas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Deployer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventUpdateDevGas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Deployer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Withdrawer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventPayoutDevGas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payouts)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRegisterDevGas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRegisterDevGas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegisterDevGas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deployer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Withdrawer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventCancelDevGas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventCancelDevGas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCancelDevGas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deployer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventUpdateDevGas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventUpdateDevGas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateDevGas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deployer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Withdrawer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventPayoutDevGas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventPayoutDevGas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPayoutDevGas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payouts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payouts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
