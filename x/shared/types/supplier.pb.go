// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/shared/supplier.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Supplier represents an actor in Pocket Network that provides RPC services
type Supplier struct {
	// Owner address that controls the staked funds and receives rewards by default
	// Cannot be updated by the operator
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// Operator address managing the offchain server
	// Immutable for supplier's lifespan - requires unstake/re-stake to change.
	// Can update supplier configs except for owner address.
	OperatorAddress string `protobuf:"bytes,2,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// Total amount of staked uPOKT
	Stake *types.Coin `protobuf:"bytes,3,opt,name=stake,proto3" json:"stake,omitempty"`
	// List of service configurations supported by this supplier
	Services []*SupplierServiceConfig `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	// Session end height when supplier initiated unstaking (0 if not unstaking)
	UnstakeSessionEndHeight uint64 `protobuf:"varint,5,opt,name=unstake_session_end_height,json=unstakeSessionEndHeight,proto3" json:"unstake_session_end_height,omitempty"`
	// List of historical service configuration updates, tracking the suppliers
	// services update and corresponding activation heights.
	ServiceConfigHistory []*ServiceConfigUpdate `protobuf:"bytes,6,rep,name=service_config_history,json=serviceConfigHistory,proto3" json:"service_config_history,omitempty"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a189b52ba503cf2, []int{0}
}
func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(m, src)
}
func (m *Supplier) XXX_Size() int {
	return m.Size()
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Supplier) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *Supplier) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *Supplier) GetServices() []*SupplierServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Supplier) GetUnstakeSessionEndHeight() uint64 {
	if m != nil {
		return m.UnstakeSessionEndHeight
	}
	return 0
}

func (m *Supplier) GetServiceConfigHistory() []*ServiceConfigUpdate {
	if m != nil {
		return m.ServiceConfigHistory
	}
	return nil
}

// ServiceConfigUpdate tracks a change in a supplier's service configurations
// at a specific block height, enabling tracking of configuration changes over time.
type ServiceConfigUpdate struct {
	// List of service configurations after the update was applied.
	Services []*SupplierServiceConfig `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	// Block height at which this service configuration update takes effect,
	// aligned with the session start height.
	EffectiveBlockHeight uint64 `protobuf:"varint,2,opt,name=effective_block_height,json=effectiveBlockHeight,proto3" json:"effective_block_height,omitempty"`
}

func (m *ServiceConfigUpdate) Reset()         { *m = ServiceConfigUpdate{} }
func (m *ServiceConfigUpdate) String() string { return proto.CompactTextString(m) }
func (*ServiceConfigUpdate) ProtoMessage()    {}
func (*ServiceConfigUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a189b52ba503cf2, []int{1}
}
func (m *ServiceConfigUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceConfigUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ServiceConfigUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConfigUpdate.Merge(m, src)
}
func (m *ServiceConfigUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ServiceConfigUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConfigUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConfigUpdate proto.InternalMessageInfo

func (m *ServiceConfigUpdate) GetServices() []*SupplierServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ServiceConfigUpdate) GetEffectiveBlockHeight() uint64 {
	if m != nil {
		return m.EffectiveBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Supplier)(nil), "poktroll.shared.Supplier")
	proto.RegisterType((*ServiceConfigUpdate)(nil), "poktroll.shared.ServiceConfigUpdate")
}

func init() { proto.RegisterFile("poktroll/shared/supplier.proto", fileDescriptor_4a189b52ba503cf2) }

var fileDescriptor_4a189b52ba503cf2 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0xe3, 0xa6, 0xad, 0x8a, 0x0b, 0x2a, 0x1a, 0xa2, 0x32, 0x8d, 0x84, 0x15, 0x55, 0x08,
	0x65, 0xd3, 0xb1, 0x5a, 0xd8, 0x21, 0x16, 0x24, 0x42, 0xea, 0x86, 0xcd, 0x44, 0x6c, 0xba, 0x19,
	0xcd, 0xcf, 0xcd, 0x8c, 0x95, 0xd4, 0x77, 0x64, 0x3b, 0x29, 0x7d, 0x0a, 0x78, 0x00, 0x1e, 0x83,
	0x87, 0x60, 0x59, 0xb1, 0xea, 0x12, 0x25, 0x2f, 0x82, 0xc6, 0xf6, 0x84, 0x42, 0x91, 0x90, 0xd8,
	0xc5, 0xf7, 0x9c, 0x7b, 0x72, 0xbe, 0xd1, 0xa5, 0xac, 0xc6, 0x99, 0x51, 0x38, 0x9f, 0x73, 0x5d,
	0xa5, 0x0a, 0x0a, 0xae, 0x17, 0x75, 0x3d, 0x17, 0xa0, 0xa2, 0x5a, 0xa1, 0xc1, 0xe0, 0xa0, 0xd5,
	0x23, 0xa7, 0xf7, 0x8f, 0x72, 0xd4, 0x97, 0xa8, 0x13, 0x2b, 0x73, 0xf7, 0x70, 0xde, 0x3e, 0x73,
	0x2f, 0x9e, 0xa5, 0x1a, 0xf8, 0xf2, 0x34, 0x03, 0x93, 0x9e, 0xf2, 0x1c, 0x85, 0xf4, 0xfa, 0xb3,
	0x7b, 0xff, 0x05, 0x6a, 0x29, 0x72, 0xf0, 0x72, 0xaf, 0xc4, 0x12, 0x5d, 0x6c, 0xf3, 0xcb, 0x4d,
	0x8f, 0xbf, 0x74, 0xe9, 0xde, 0xc4, 0x77, 0x0a, 0xde, 0xd0, 0x47, 0x78, 0x25, 0x41, 0x25, 0x69,
	0x51, 0x28, 0xd0, 0x3a, 0x24, 0x03, 0x32, 0x7c, 0x30, 0x0a, 0xbf, 0x7f, 0x3d, 0xe9, 0xf9, 0x2a,
	0x6f, 0x9d, 0x32, 0x31, 0x4a, 0xc8, 0x32, 0x7e, 0x68, 0xed, 0x7e, 0x16, 0x8c, 0xe9, 0x63, 0xac,
	0x41, 0xa5, 0x06, 0x7f, 0x25, 0x6c, 0xfd, 0x23, 0xe1, 0xa0, 0xdd, 0x68, 0x43, 0x38, 0xdd, 0xd1,
	0x26, 0x9d, 0x41, 0xd8, 0x1d, 0x90, 0xe1, 0xfe, 0xd9, 0x51, 0xe4, 0xd7, 0x1a, 0xea, 0xc8, 0x53,
	0x47, 0x63, 0x14, 0x32, 0x76, 0xbe, 0x60, 0x44, 0xf7, 0x3c, 0xa8, 0x0e, 0xb7, 0x07, 0xdd, 0xe1,
	0xfe, 0xd9, 0x8b, 0xe8, 0x8f, 0xaf, 0x1a, 0xb5, 0x84, 0x13, 0x67, 0x1c, 0xa3, 0x9c, 0x8a, 0x32,
	0xde, 0xec, 0x05, 0xaf, 0x69, 0x7f, 0x21, 0x6d, 0x5c, 0xa2, 0x41, 0x6b, 0x81, 0x32, 0x01, 0x59,
	0x24, 0x15, 0x88, 0xb2, 0x32, 0xe1, 0xce, 0x80, 0x0c, 0xb7, 0xe3, 0xa7, 0xde, 0x31, 0x71, 0x86,
	0x77, 0xb2, 0x38, 0xb7, 0x72, 0x70, 0x41, 0x0f, 0x7d, 0x50, 0x92, 0xdb, 0xe0, 0xa4, 0x12, 0xda,
	0xa0, 0xba, 0x0e, 0x77, 0x6d, 0x9d, 0xe7, 0xf7, 0xeb, 0xdc, 0xad, 0xf1, 0xa1, 0x2e, 0x52, 0x03,
	0x71, 0x4f, 0xdf, 0x1d, 0x9e, 0xbb, 0x84, 0xe3, 0x4f, 0x84, 0x3e, 0xf9, 0x8b, 0xfb, 0x37, 0x68,
	0xf2, 0x9f, 0xd0, 0xaf, 0xe8, 0x21, 0x4c, 0xa7, 0x90, 0x1b, 0xb1, 0x84, 0x24, 0x9b, 0x63, 0x3e,
	0x6b, 0x81, 0xb7, 0x2c, 0x70, 0x6f, 0xa3, 0x8e, 0x1a, 0xd1, 0xd1, 0x8e, 0xde, 0x7f, 0x5b, 0x31,
	0x72, 0xb3, 0x62, 0xe4, 0x76, 0xc5, 0xc8, 0x8f, 0x15, 0x23, 0x9f, 0xd7, 0xac, 0x73, 0xb3, 0x66,
	0x9d, 0xdb, 0x35, 0xeb, 0x5c, 0xf0, 0x52, 0x98, 0x6a, 0x91, 0x45, 0x39, 0x5e, 0xf2, 0xa6, 0xcf,
	0x89, 0x04, 0x73, 0x85, 0x6a, 0xc6, 0x37, 0xb7, 0xf9, 0xb1, 0xbd, 0x4e, 0x73, 0x5d, 0x83, 0xce,
	0x76, 0xed, 0x19, 0xbe, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x35, 0x05, 0x68, 0x63, 0x29, 0x03,
	0x00, 0x00,
}

func (m *Supplier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supplier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supplier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceConfigHistory) > 0 {
		for iNdEx := len(m.ServiceConfigHistory) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceConfigHistory[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupplier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		i = encodeVarintSupplier(dAtA, i, uint64(m.UnstakeSessionEndHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupplier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSupplier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintSupplier(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintSupplier(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceConfigUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceConfigUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceConfigUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EffectiveBlockHeight != 0 {
		i = encodeVarintSupplier(dAtA, i, uint64(m.EffectiveBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupplier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSupplier(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupplier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Supplier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovSupplier(uint64(l))
	}
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovSupplier(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovSupplier(uint64(l))
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSupplier(uint64(l))
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		n += 1 + sovSupplier(uint64(m.UnstakeSessionEndHeight))
	}
	if len(m.ServiceConfigHistory) > 0 {
		for _, e := range m.ServiceConfigHistory {
			l = e.Size()
			n += 1 + l + sovSupplier(uint64(l))
		}
	}
	return n
}

func (m *ServiceConfigUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSupplier(uint64(l))
		}
	}
	if m.EffectiveBlockHeight != 0 {
		n += 1 + sovSupplier(uint64(m.EffectiveBlockHeight))
	}
	return n
}

func sovSupplier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupplier(x uint64) (n int) {
	return sovSupplier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Supplier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplier
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
			return fmt.Errorf("proto: Supplier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supplier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &SupplierServiceConfig{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeSessionEndHeight", wireType)
			}
			m.UnstakeSessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeSessionEndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceConfigHistory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceConfigHistory = append(m.ServiceConfigHistory, &ServiceConfigUpdate{})
			if err := m.ServiceConfigHistory[len(m.ServiceConfigHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupplier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupplier
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
func (m *ServiceConfigUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplier
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
			return fmt.Errorf("proto: ServiceConfigUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceConfigUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &SupplierServiceConfig{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveBlockHeight", wireType)
			}
			m.EffectiveBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EffectiveBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSupplier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupplier
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
func skipSupplier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
				return 0, ErrInvalidLengthSupplier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupplier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupplier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupplier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupplier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupplier = fmt.Errorf("proto: unexpected end of group")
)
