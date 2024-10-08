// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/application/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_sortkeys "github.com/cosmos/gogoproto/sortkeys"
	types1 "github.com/pokt-network/poktroll/x/shared/types"
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

// Application defines the type used to store an on-chain definition and state for an application
type Application struct {
	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stake   *types.Coin `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	// TODO_BETA(@red-0ne, @olshansk): Limit this to one service_config.
	//    Remove `repeated`, drop the `s` from service_configs and document why
	//    this is the case in the app config (and here) per this discussion:
	//    https://github.com/pokt-network/poktroll/pull/750#discussion_r1735025033
	ServiceConfigs []*types1.ApplicationServiceConfig `protobuf:"bytes,3,rep,name=service_configs,json=serviceConfigs,proto3" json:"service_configs,omitempty"`
	// TODO_BETA: Rename `delegatee_gateway_addresses` to `gateway_addresses_delegated_to`.
	// Ensure to rename all relevant configs, comments, variables, function names, etc as well.
	DelegateeGatewayAddresses []string `protobuf:"bytes,4,rep,name=delegatee_gateway_addresses,json=delegateeGatewayAddresses,proto3" json:"delegatee_gateway_addresses,omitempty"`
	// A map from sessionEndHeights to a list of Gateways.
	// The key is the height of the last block of the session during which the
	// respective undelegation was committed.
	// The value is a list of gateways being undelegated from.
	// TODO_DOCUMENT(@red-0ne): Need to document the flow from this comment
	// so its clear to everyone why this is necessary; https://github.com/pokt-network/poktroll/issues/476#issuecomment-2052639906.
	PendingUndelegations map[uint64]UndelegatingGatewayList `protobuf:"bytes,5,rep,name=pending_undelegations,json=pendingUndelegations,proto3" json:"pending_undelegations" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The end height of the session at which an application initiated its unstaking process.
	// If the application did not unstake, this value will be 0.
	UnstakeSessionEndHeight uint64 `protobuf:"varint,6,opt,name=unstake_session_end_height,json=unstakeSessionEndHeight,proto3" json:"unstake_session_end_height,omitempty"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_1899440439257283, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return m.Size()
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Application) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *Application) GetServiceConfigs() []*types1.ApplicationServiceConfig {
	if m != nil {
		return m.ServiceConfigs
	}
	return nil
}

func (m *Application) GetDelegateeGatewayAddresses() []string {
	if m != nil {
		return m.DelegateeGatewayAddresses
	}
	return nil
}

func (m *Application) GetPendingUndelegations() map[uint64]UndelegatingGatewayList {
	if m != nil {
		return m.PendingUndelegations
	}
	return nil
}

func (m *Application) GetUnstakeSessionEndHeight() uint64 {
	if m != nil {
		return m.UnstakeSessionEndHeight
	}
	return 0
}

// UndelegatingGatewayList is used as the Value of `pending_undelegations`.
// It is required to store a repeated list of strings as a map value.
type UndelegatingGatewayList struct {
	GatewayAddresses []string `protobuf:"bytes,2,rep,name=gateway_addresses,json=gatewayAddresses,proto3" json:"gateway_addresses,omitempty"`
}

func (m *UndelegatingGatewayList) Reset()         { *m = UndelegatingGatewayList{} }
func (m *UndelegatingGatewayList) String() string { return proto.CompactTextString(m) }
func (*UndelegatingGatewayList) ProtoMessage()    {}
func (*UndelegatingGatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1899440439257283, []int{1}
}
func (m *UndelegatingGatewayList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UndelegatingGatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *UndelegatingGatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndelegatingGatewayList.Merge(m, src)
}
func (m *UndelegatingGatewayList) XXX_Size() int {
	return m.Size()
}
func (m *UndelegatingGatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_UndelegatingGatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_UndelegatingGatewayList proto.InternalMessageInfo

func (m *UndelegatingGatewayList) GetGatewayAddresses() []string {
	if m != nil {
		return m.GatewayAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Application)(nil), "poktroll.application.Application")
	proto.RegisterMapType((map[uint64]UndelegatingGatewayList)(nil), "poktroll.application.Application.PendingUndelegationsEntry")
	proto.RegisterType((*UndelegatingGatewayList)(nil), "poktroll.application.UndelegatingGatewayList")
}

func init() { proto.RegisterFile("poktroll/application/types.proto", fileDescriptor_1899440439257283) }

var fileDescriptor_1899440439257283 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x1c, 0xc6, 0x9b, 0xbe, 0x0c, 0xcd, 0x95, 0x60, 0x58, 0x45, 0x4b, 0x0b, 0x84, 0x68, 0xa7, 0x72,
	0xa8, 0xad, 0x15, 0x0e, 0x88, 0x9d, 0xda, 0x6a, 0x02, 0x24, 0x0e, 0x28, 0x15, 0x17, 0x84, 0x14,
	0xb9, 0x89, 0x71, 0xad, 0x76, 0x76, 0x14, 0xbb, 0x19, 0xfd, 0x16, 0x7c, 0x18, 0x3e, 0xc4, 0x8e,
	0x83, 0xd3, 0x4e, 0x08, 0xb5, 0x5f, 0x04, 0x25, 0x76, 0xbb, 0x88, 0xb5, 0xd2, 0x2e, 0x51, 0xac,
	0xe7, 0xf7, 0xd8, 0x8f, 0xfd, 0xe8, 0x0f, 0xfc, 0x44, 0xce, 0x74, 0x2a, 0xe7, 0x73, 0x4c, 0x92,
	0x64, 0xce, 0x23, 0xa2, 0xb9, 0x14, 0x58, 0x2f, 0x13, 0xaa, 0x50, 0x92, 0x4a, 0x2d, 0x61, 0x6b,
	0x43, 0xa0, 0x12, 0xd1, 0x69, 0x31, 0xc9, 0x64, 0x01, 0xe0, 0xfc, 0xcf, 0xb0, 0x1d, 0x2f, 0x92,
	0xea, 0x42, 0x2a, 0x3c, 0x21, 0x8a, 0xe2, 0xec, 0x74, 0x42, 0x35, 0x39, 0xc5, 0x91, 0xe4, 0xc2,
	0xea, 0x6d, 0xa3, 0x87, 0xc6, 0x68, 0x16, 0x56, 0x7a, 0xbe, 0x0d, 0xa2, 0xa6, 0x24, 0xa5, 0x31,
	0x56, 0x34, 0xcd, 0x78, 0x44, 0x8d, 0x7c, 0xf2, 0xab, 0x0e, 0x9a, 0x83, 0xdb, 0xf3, 0x61, 0x1f,
	0x3c, 0x20, 0x71, 0x9c, 0x52, 0xa5, 0x5c, 0xc7, 0x77, 0xba, 0x87, 0x43, 0xf7, 0xf7, 0xcf, 0x5e,
	0xcb, 0xee, 0x38, 0x30, 0xca, 0x58, 0xa7, 0x5c, 0xb0, 0x60, 0x03, 0x42, 0x0c, 0x1a, 0x4a, 0x93,
	0x19, 0x75, 0xab, 0xbe, 0xd3, 0x6d, 0xf6, 0xdb, 0xc8, 0xe2, 0x79, 0x5a, 0x64, 0xd3, 0xa2, 0x91,
	0xe4, 0x22, 0x30, 0x1c, 0x0c, 0xc0, 0x23, 0x9b, 0x22, 0x8c, 0xa4, 0xf8, 0xc6, 0x99, 0x72, 0x6b,
	0x7e, 0xad, 0xdb, 0xec, 0xbf, 0x44, 0xdb, 0x47, 0x31, 0x69, 0x51, 0x29, 0xdb, 0xd8, 0x58, 0x46,
	0x85, 0x23, 0x78, 0xa8, 0xca, 0x4b, 0x05, 0xbf, 0x82, 0xa7, 0x31, 0x9d, 0x53, 0x46, 0x34, 0xa5,
	0x61, 0xfe, 0xbd, 0x24, 0xcb, 0xd0, 0x26, 0xa4, 0xca, 0xad, 0xfb, 0xb5, 0xee, 0xe1, 0xf0, 0xd9,
	0xd5, 0x9f, 0x17, 0x95, 0xbd, 0x17, 0x6a, 0x6f, 0x37, 0x78, 0x67, 0xfc, 0x83, 0x8d, 0x1d, 0x66,
	0xe0, 0x49, 0x42, 0x45, 0xcc, 0x05, 0x0b, 0x17, 0xc2, 0x62, 0x5c, 0x0a, 0xe5, 0x36, 0x8a, 0xdc,
	0x67, 0x68, 0x57, 0x99, 0xe5, 0xf0, 0xe8, 0x93, 0xb1, 0x7f, 0x2e, 0xbb, 0xcf, 0x85, 0x4e, 0x97,
	0xc3, 0x7a, 0x1e, 0x2a, 0x68, 0x25, 0x3b, 0x00, 0x78, 0x06, 0x3a, 0x0b, 0x51, 0x3c, 0x5a, 0xa8,
	0xa8, 0x52, 0x5c, 0x8a, 0x90, 0x8a, 0x38, 0x9c, 0x52, 0xce, 0xa6, 0xda, 0x3d, 0xf0, 0x9d, 0x6e,
	0x3d, 0x38, 0xb6, 0xc4, 0xd8, 0x00, 0xe7, 0x22, 0x7e, 0x5f, 0xc8, 0x9d, 0x0c, 0xb4, 0xf7, 0x9e,
	0x0a, 0x8f, 0x40, 0x6d, 0x46, 0x97, 0x45, 0xc9, 0xf5, 0x20, 0xff, 0x85, 0x23, 0xd0, 0xc8, 0xc8,
	0x7c, 0xb1, 0xa9, 0xb1, 0xb7, 0xfb, 0x4e, 0xb7, 0x5b, 0x09, 0x66, 0x9f, 0xe9, 0x23, 0x57, 0x3a,
	0x30, 0xde, 0xb7, 0xd5, 0x37, 0xce, 0x49, 0x0c, 0x8e, 0xf7, 0x50, 0xf0, 0x03, 0x78, 0x7c, 0xb7,
	0x9b, 0xea, 0x3d, 0xba, 0x39, 0x62, 0xff, 0x55, 0x32, 0x0c, 0xae, 0x56, 0x9e, 0x73, 0xbd, 0xf2,
	0x9c, 0x9b, 0x95, 0xe7, 0xfc, 0x5d, 0x79, 0xce, 0x8f, 0xb5, 0x57, 0xb9, 0x5e, 0x7b, 0x95, 0x9b,
	0xb5, 0x57, 0xf9, 0xf2, 0x9a, 0x71, 0x3d, 0x5d, 0x4c, 0x50, 0x24, 0x2f, 0x70, 0x7e, 0x8f, 0x9e,
	0xa0, 0xfa, 0x52, 0xa6, 0x33, 0xbc, 0x1d, 0x87, 0xef, 0x77, 0x27, 0x73, 0x72, 0x50, 0x0c, 0xc5,
	0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x59, 0x36, 0xdf, 0xbe, 0x03, 0x00, 0x00,
}

func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Application) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnstakeSessionEndHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.UnstakeSessionEndHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PendingUndelegations) > 0 {
		keysForPendingUndelegations := make([]uint64, 0, len(m.PendingUndelegations))
		for k := range m.PendingUndelegations {
			keysForPendingUndelegations = append(keysForPendingUndelegations, uint64(k))
		}
		github_com_cosmos_gogoproto_sortkeys.Uint64s(keysForPendingUndelegations)
		for iNdEx := len(keysForPendingUndelegations) - 1; iNdEx >= 0; iNdEx-- {
			v := m.PendingUndelegations[uint64(keysForPendingUndelegations[iNdEx])]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i = encodeVarintTypes(dAtA, i, uint64(keysForPendingUndelegations[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DelegateeGatewayAddresses) > 0 {
		for iNdEx := len(m.DelegateeGatewayAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DelegateeGatewayAddresses[iNdEx])
			copy(dAtA[i:], m.DelegateeGatewayAddresses[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.DelegateeGatewayAddresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ServiceConfigs) > 0 {
		for iNdEx := len(m.ServiceConfigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceConfigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UndelegatingGatewayList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UndelegatingGatewayList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UndelegatingGatewayList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GatewayAddresses) > 0 {
		for iNdEx := len(m.GatewayAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GatewayAddresses[iNdEx])
			copy(dAtA[i:], m.GatewayAddresses[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.GatewayAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
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
func (m *Application) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.ServiceConfigs) > 0 {
		for _, e := range m.ServiceConfigs {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.DelegateeGatewayAddresses) > 0 {
		for _, s := range m.DelegateeGatewayAddresses {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.PendingUndelegations) > 0 {
		for k, v := range m.PendingUndelegations {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + sovTypes(uint64(k)) + 1 + l + sovTypes(uint64(l))
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		n += 1 + sovTypes(uint64(m.UnstakeSessionEndHeight))
	}
	return n
}

func (m *UndelegatingGatewayList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GatewayAddresses) > 0 {
		for _, s := range m.GatewayAddresses {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Application) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
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
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceConfigs", wireType)
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
			m.ServiceConfigs = append(m.ServiceConfigs, &types1.ApplicationServiceConfig{})
			if err := m.ServiceConfigs[len(m.ServiceConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateeGatewayAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegateeGatewayAddresses = append(m.DelegateeGatewayAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingUndelegations", wireType)
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
			if m.PendingUndelegations == nil {
				m.PendingUndelegations = make(map[uint64]UndelegatingGatewayList)
			}
			var mapkey uint64
			mapvalue := &UndelegatingGatewayList{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTypes
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTypes
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &UndelegatingGatewayList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PendingUndelegations[mapkey] = *mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeSessionEndHeight", wireType)
			}
			m.UnstakeSessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
func (m *UndelegatingGatewayList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UndelegatingGatewayList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UndelegatingGatewayList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayAddresses = append(m.GatewayAddresses, string(dAtA[iNdEx:postIndex]))
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
