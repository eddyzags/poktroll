// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/tokenomics/params.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Params defines the parameters for the tokenomics module.
type Params struct {
	// mint_allocation_dao is the percentage of the minted tokens which are sent
	// to the DAO reward address during claim settlement.
	MintAllocationDao float64 `protobuf:"fixed64,1,opt,name=mint_allocation_dao,json=mintAllocationDao,proto3" json:"mint_allocation_dao" yaml:"mint_allocation_dao"`
	// mint_allocation_proposer is the percentage of the minted tokens which are sent
	// to the block proposer account address during claim settlement.
	MintAllocationProposer float64 `protobuf:"fixed64,2,opt,name=mint_allocation_proposer,json=mintAllocationProposer,proto3" json:"mint_allocation_proposer" yaml:"mint_allocation_proposer"`
	// mint_allocation_supplier is the percentage of the minted tokens which are sent
	// to the block supplier account address during claim settlement.
	MintAllocationSupplier float64 `protobuf:"fixed64,3,opt,name=mint_allocation_supplier,json=mintAllocationSupplier,proto3" json:"mint_allocation_supplier" yaml:"mint_allocation_supplier"`
	// mint_allocation_source_owner is the percentage of the minted tokens which are sent
	// to the service source owner account address during claim settlement.
	MintAllocationSourceOwner float64 `protobuf:"fixed64,4,opt,name=mint_allocation_source_owner,json=mintAllocationSourceOwner,proto3" json:"mint_allocation_source_owner" yaml:"mint_allocation_source_owner"`
	// mint_allocation_application is the percentage of the minted tokens which are sent
	// to the application account address during claim settlement.
	MintAllocationApplication float64 `protobuf:"fixed64,5,opt,name=mint_allocation_application,json=mintAllocationApplication,proto3" json:"mint_allocation_application" yaml:"mint_allocation_application"`
	// dao_reward_address is the address to which mint_allocation_dao percentage of the
	// minted tokens are at the end of claim settlement.
	DaoRewardAddress string `protobuf:"bytes,6,opt,name=dao_reward_address,json=daoRewardAddress,proto3" json:"dao_reward_address" yaml:"dao_reward_address"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_df10a06914fc6eee, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
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

func (m *Params) GetMintAllocationDao() float64 {
	if m != nil {
		return m.MintAllocationDao
	}
	return 0
}

func (m *Params) GetMintAllocationProposer() float64 {
	if m != nil {
		return m.MintAllocationProposer
	}
	return 0
}

func (m *Params) GetMintAllocationSupplier() float64 {
	if m != nil {
		return m.MintAllocationSupplier
	}
	return 0
}

func (m *Params) GetMintAllocationSourceOwner() float64 {
	if m != nil {
		return m.MintAllocationSourceOwner
	}
	return 0
}

func (m *Params) GetMintAllocationApplication() float64 {
	if m != nil {
		return m.MintAllocationApplication
	}
	return 0
}

func (m *Params) GetDaoRewardAddress() string {
	if m != nil {
		return m.DaoRewardAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "poktroll.tokenomics.Params")
}

func init() { proto.RegisterFile("poktroll/tokenomics/params.proto", fileDescriptor_df10a06914fc6eee) }

var fileDescriptor_df10a06914fc6eee = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0xfc, 0x89, 0x84, 0x27, 0xea, 0x56, 0xc8, 0x09, 0x95, 0x2f, 0x32, 0x42, 0xaa,
	0x90, 0x1a, 0x0f, 0x15, 0x4b, 0x17, 0x94, 0x08, 0xc4, 0xc0, 0x40, 0x71, 0x37, 0x16, 0xeb, 0x6a,
	0x9f, 0x82, 0x15, 0xdb, 0xef, 0xe9, 0xee, 0xa2, 0x90, 0x0f, 0x80, 0x84, 0x98, 0x98, 0x99, 0xf8,
	0x08, 0x0c, 0x7c, 0x08, 0xc6, 0x8a, 0xa9, 0xd3, 0x09, 0x25, 0x03, 0xc8, 0xa3, 0x3f, 0x01, 0xf2,
	0x9d, 0xd3, 0xa6, 0x89, 0x93, 0x25, 0xba, 0xbc, 0xbf, 0xdf, 0xdd, 0xf3, 0x0c, 0x7e, 0xad, 0x1e,
	0x83, 0xb1, 0xe4, 0x90, 0xa6, 0xbe, 0x84, 0x31, 0xcd, 0x21, 0x4b, 0x22, 0xe1, 0x33, 0xc2, 0x49,
	0x26, 0xfa, 0x8c, 0x83, 0x04, 0x7b, 0x7f, 0x69, 0xf4, 0x6f, 0x8c, 0xee, 0x1e, 0xc9, 0x92, 0x1c,
	0x7c, 0xfd, 0x6b, 0xbc, 0xee, 0xc1, 0x08, 0x46, 0xa0, 0x8f, 0x7e, 0x75, 0xaa, 0xa7, 0x9d, 0x08,
	0x44, 0x06, 0x22, 0x34, 0xc0, 0xfc, 0x31, 0xc8, 0xfb, 0xd6, 0xb6, 0xda, 0x67, 0x3a, 0xc9, 0xa6,
	0xd6, 0x7e, 0x96, 0xe4, 0x32, 0x24, 0x69, 0x0a, 0x11, 0x91, 0x09, 0xe4, 0x61, 0x4c, 0xc0, 0x41,
	0x3d, 0x74, 0x84, 0x86, 0xcf, 0x0b, 0x85, 0x9b, 0x70, 0xa9, 0x70, 0x77, 0x46, 0xb2, 0xf4, 0xd4,
	0x6b, 0x80, 0x5e, 0xb0, 0x57, 0x4d, 0x07, 0xd7, 0xc3, 0x97, 0x04, 0xec, 0x99, 0xe5, 0xac, 0xab,
	0x8c, 0x03, 0x03, 0x41, 0xb9, 0x73, 0x47, 0x67, 0xbd, 0x28, 0x14, 0xde, 0xea, 0x94, 0x0a, 0xe3,
	0xe6, 0xc0, 0xa5, 0xe1, 0x05, 0x8f, 0x6e, 0xa7, 0x9e, 0xd5, 0xa0, 0x29, 0x5a, 0x4c, 0x18, 0x4b,
	0x13, 0xca, 0x9d, 0xbb, 0xdb, 0xa3, 0x97, 0xce, 0xf6, 0xe8, 0xa5, 0xb1, 0x11, 0x7d, 0x5e, 0x03,
	0xfb, 0x33, 0xb2, 0x0e, 0x37, 0x6e, 0xc1, 0x84, 0x47, 0x34, 0x84, 0x69, 0x4e, 0xb9, 0x73, 0x4f,
	0xe7, 0xbf, 0x2e, 0x14, 0xde, 0xe9, 0x95, 0x0a, 0x3f, 0xd9, 0xd2, 0x61, 0xc5, 0xf2, 0x82, 0xce,
	0x5a, 0x0f, 0x0d, 0xdf, 0x56, 0xcc, 0xfe, 0x84, 0xac, 0xc7, 0xeb, 0x97, 0x49, 0x55, 0xd3, 0x9c,
	0x9d, 0xfb, 0xba, 0xc9, 0xab, 0x42, 0xe1, 0x5d, 0x5a, 0xa9, 0xb0, 0xd7, 0x5c, 0x64, 0x45, 0xda,
	0xe8, 0x31, 0xb8, 0x61, 0xf6, 0xcc, 0xb2, 0x63, 0x02, 0x21, 0xa7, 0x53, 0xc2, 0xe3, 0x90, 0xc4,
	0x31, 0xa7, 0x42, 0x38, 0xed, 0x1e, 0x3a, 0x7a, 0x30, 0x7c, 0x53, 0x28, 0xdc, 0x40, 0x4b, 0x85,
	0x3b, 0x26, 0x74, 0x93, 0x79, 0xbf, 0x7f, 0x1e, 0x1f, 0xd4, 0xdf, 0xf6, 0xc0, 0x8c, 0xce, 0x25,
	0x4f, 0xf2, 0x51, 0xf0, 0x30, 0x26, 0x10, 0x68, 0xb7, 0x9e, 0x9f, 0x3e, 0xfd, 0xf7, 0x1d, 0xa3,
	0x2f, 0x7f, 0x7f, 0x3c, 0x3b, 0xbc, 0xde, 0xbc, 0x8f, 0xab, 0xbb, 0x67, 0x36, 0x62, 0xf8, 0xee,
	0xd7, 0xdc, 0x45, 0x97, 0x73, 0x17, 0x5d, 0xcd, 0x5d, 0xf4, 0x67, 0xee, 0xa2, 0xaf, 0x0b, 0xb7,
	0x75, 0xb9, 0x70, 0x5b, 0x57, 0x0b, 0xb7, 0xf5, 0xfe, 0x64, 0x94, 0xc8, 0x0f, 0x93, 0x8b, 0x7e,
	0x04, 0x99, 0x5f, 0x3d, 0x73, 0x9c, 0x53, 0x39, 0x05, 0x3e, 0xf6, 0x9b, 0xdf, 0x94, 0x33, 0x46,
	0xc5, 0x45, 0x5b, 0xaf, 0xdd, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x02, 0xe9, 0xcd,
	0xf3, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MintAllocationDao != that1.MintAllocationDao {
		return false
	}
	if this.MintAllocationProposer != that1.MintAllocationProposer {
		return false
	}
	if this.MintAllocationSupplier != that1.MintAllocationSupplier {
		return false
	}
	if this.MintAllocationSourceOwner != that1.MintAllocationSourceOwner {
		return false
	}
	if this.MintAllocationApplication != that1.MintAllocationApplication {
		return false
	}
	if this.DaoRewardAddress != that1.DaoRewardAddress {
		return false
	}
	return true
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
	if len(m.DaoRewardAddress) > 0 {
		i -= len(m.DaoRewardAddress)
		copy(dAtA[i:], m.DaoRewardAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DaoRewardAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.MintAllocationApplication != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MintAllocationApplication))))
		i--
		dAtA[i] = 0x29
	}
	if m.MintAllocationSourceOwner != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MintAllocationSourceOwner))))
		i--
		dAtA[i] = 0x21
	}
	if m.MintAllocationSupplier != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MintAllocationSupplier))))
		i--
		dAtA[i] = 0x19
	}
	if m.MintAllocationProposer != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MintAllocationProposer))))
		i--
		dAtA[i] = 0x11
	}
	if m.MintAllocationDao != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MintAllocationDao))))
		i--
		dAtA[i] = 0x9
	}
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
	if m.MintAllocationDao != 0 {
		n += 9
	}
	if m.MintAllocationProposer != 0 {
		n += 9
	}
	if m.MintAllocationSupplier != 0 {
		n += 9
	}
	if m.MintAllocationSourceOwner != 0 {
		n += 9
	}
	if m.MintAllocationApplication != 0 {
		n += 9
	}
	l = len(m.DaoRewardAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAllocationDao", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MintAllocationDao = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAllocationProposer", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MintAllocationProposer = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAllocationSupplier", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MintAllocationSupplier = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAllocationSourceOwner", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MintAllocationSourceOwner = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAllocationApplication", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MintAllocationApplication = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoRewardAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaoRewardAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
