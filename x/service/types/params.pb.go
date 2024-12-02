// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/service/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	// The amount of uPOKT required to add a new service.
	// This will be deducted from the signer's account balance,
	// and transferred to the pocket network foundation.
	AddServiceFee *types.Coin `protobuf:"bytes,1,opt,name=add_service_fee,json=addServiceFee,proto3" json:"add_service_fee" yaml:"add_service_fee"`
	// target_num_relays is the target for the EMA of the number of relays per session.
	// Per service, on-chain relay mining difficulty will be adjusted to maintain this target.
	TargetNumRelays uint64 `protobuf:"varint,2,opt,name=target_num_relays,json=targetNumRelays,proto3" json:"target_num_relays" yaml:"target_num_relays"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b5d0104478b383, []int{0}
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

func (m *Params) GetAddServiceFee() *types.Coin {
	if m != nil {
		return m.AddServiceFee
	}
	return nil
}

func (m *Params) GetTargetNumRelays() uint64 {
	if m != nil {
		return m.TargetNumRelays
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "poktroll.service.Params")
}

func init() { proto.RegisterFile("poktroll/service/params.proto", fileDescriptor_69b5d0104478b383) }

var fileDescriptor_69b5d0104478b383 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x7b, 0x22, 0x1d, 0x22, 0x52, 0x5b, 0x44, 0xda, 0x82, 0x97, 0x92, 0xa9, 0x08, 0xcd,
	0x59, 0xdd, 0x3a, 0x56, 0x70, 0x2c, 0x52, 0x37, 0x41, 0xc2, 0x25, 0x39, 0x63, 0x68, 0x2e, 0x5f,
	0xb8, 0xbb, 0x56, 0xfb, 0x0a, 0x4e, 0x3e, 0x82, 0x8f, 0xe0, 0x63, 0x38, 0x76, 0xec, 0x14, 0x24,
	0x1d, 0x94, 0x8e, 0xdd, 0x05, 0x69, 0x2e, 0x76, 0x68, 0x5d, 0x8e, 0x8f, 0xff, 0xef, 0xee, 0xf7,
	0x71, 0x7f, 0xe3, 0x34, 0x81, 0x91, 0x12, 0x10, 0x45, 0x44, 0x32, 0x31, 0x09, 0x3d, 0x46, 0x12,
	0x2a, 0x28, 0x97, 0x76, 0x22, 0x40, 0x41, 0xed, 0xe8, 0x0f, 0xdb, 0x05, 0x6e, 0x56, 0x29, 0x0f,
	0x63, 0x20, 0xf9, 0xa9, 0x2f, 0x35, 0x8f, 0x03, 0x08, 0x20, 0x1f, 0xc9, 0x7a, 0x2a, 0x52, 0xec,
	0x81, 0xe4, 0x20, 0x89, 0x4b, 0x25, 0x23, 0x93, 0xae, 0xcb, 0x14, 0xed, 0x12, 0x0f, 0xc2, 0x58,
	0x73, 0xeb, 0x07, 0x19, 0xe5, 0x9b, 0x7c, 0x57, 0x2d, 0x31, 0x2a, 0xd4, 0xf7, 0x9d, 0x62, 0x85,
	0xf3, 0xc0, 0x58, 0x1d, 0xb5, 0x50, 0xfb, 0xe0, 0xa2, 0x61, 0x6b, 0x89, 0xbd, 0x96, 0xd8, 0x85,
	0xc4, 0xbe, 0x82, 0x30, 0xee, 0x77, 0x96, 0xa9, 0xb9, 0xfd, 0x6a, 0x95, 0x9a, 0x27, 0x53, 0xca,
	0xa3, 0x9e, 0xb5, 0x05, 0xac, 0xe1, 0x21, 0xf5, 0xfd, 0x5b, 0x1d, 0x5c, 0x33, 0x56, 0xbb, 0x37,
	0xaa, 0x8a, 0x8a, 0x80, 0x29, 0x27, 0x1e, 0x73, 0x47, 0xb0, 0x88, 0x4e, 0x65, 0x7d, 0xaf, 0x85,
	0xda, 0xfb, 0xfd, 0xee, 0x32, 0x35, 0x77, 0xe1, 0x2a, 0x35, 0xeb, 0x5a, 0xbd, 0x83, 0xac, 0x61,
	0x45, 0x67, 0x83, 0x31, 0x1f, 0xe6, 0x49, 0xcf, 0xfa, 0x7e, 0x33, 0xd1, 0xcb, 0xd7, 0xfb, 0x59,
	0x63, 0x53, 0xef, 0xf3, 0xa6, 0x60, 0xfd, 0xe9, 0xfe, 0xe0, 0x23, 0xc3, 0x68, 0x96, 0x61, 0x34,
	0xcf, 0x30, 0xfa, 0xcc, 0x30, 0x7a, 0x5d, 0xe0, 0xd2, 0x6c, 0x81, 0x4b, 0xf3, 0x05, 0x2e, 0xdd,
	0x9d, 0x07, 0xa1, 0x7a, 0x1c, 0xbb, 0xb6, 0x07, 0x9c, 0xac, 0x1d, 0x9d, 0x98, 0xa9, 0x27, 0x10,
	0x23, 0xf2, 0x8f, 0x50, 0x4d, 0x13, 0x26, 0xdd, 0x72, 0x5e, 0xeb, 0xe5, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcf, 0xf3, 0x1c, 0x51, 0xd2, 0x01, 0x00, 0x00,
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
	if !this.AddServiceFee.Equal(that1.AddServiceFee) {
		return false
	}
	if this.TargetNumRelays != that1.TargetNumRelays {
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
	if m.TargetNumRelays != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TargetNumRelays))
		i--
		dAtA[i] = 0x10
	}
	if m.AddServiceFee != nil {
		{
			size, err := m.AddServiceFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
	if m.AddServiceFee != nil {
		l = m.AddServiceFee.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.TargetNumRelays != 0 {
		n += 1 + sovParams(uint64(m.TargetNumRelays))
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
				return fmt.Errorf("proto: wrong wireType = %d for field AddServiceFee", wireType)
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
			if m.AddServiceFee == nil {
				m.AddServiceFee = &types.Coin{}
			}
			if err := m.AddServiceFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetNumRelays", wireType)
			}
			m.TargetNumRelays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetNumRelays |= uint64(b&0x7F) << shift
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
