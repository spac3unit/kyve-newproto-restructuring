// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/stakers/v1beta1/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	// vote_slash ...
	VoteSlash string `protobuf:"bytes,1,opt,name=vote_slash,json=voteSlash,proto3" json:"vote_slash,omitempty"`
	// upload_slash ...
	UploadSlash string `protobuf:"bytes,2,opt,name=upload_slash,json=uploadSlash,proto3" json:"upload_slash,omitempty"`
	// timeout_slash ...
	TimeoutSlash string `protobuf:"bytes,3,opt,name=timeout_slash,json=timeoutSlash,proto3" json:"timeout_slash,omitempty"`
	// max_points ...
	MaxPoints uint64 `protobuf:"varint,4,opt,name=max_points,json=maxPoints,proto3" json:"max_points,omitempty"`
	// unbonding_staking_time ...
	UnbondingStakingTime uint64 `protobuf:"varint,5,opt,name=unbonding_staking_time,json=unbondingStakingTime,proto3" json:"unbonding_staking_time,omitempty"`
	// commission_change_time ...
	CommissionChangeTime uint64 `protobuf:"varint,6,opt,name=commission_change_time,json=commissionChangeTime,proto3" json:"commission_change_time,omitempty"`
	// commission_change_time ...
	LeavePoolTime uint64 `protobuf:"varint,7,opt,name=leave_pool_time,json=leavePoolTime,proto3" json:"leave_pool_time,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_405cabd7005fc18b, []int{0}
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

func (m *Params) GetVoteSlash() string {
	if m != nil {
		return m.VoteSlash
	}
	return ""
}

func (m *Params) GetUploadSlash() string {
	if m != nil {
		return m.UploadSlash
	}
	return ""
}

func (m *Params) GetTimeoutSlash() string {
	if m != nil {
		return m.TimeoutSlash
	}
	return ""
}

func (m *Params) GetMaxPoints() uint64 {
	if m != nil {
		return m.MaxPoints
	}
	return 0
}

func (m *Params) GetUnbondingStakingTime() uint64 {
	if m != nil {
		return m.UnbondingStakingTime
	}
	return 0
}

func (m *Params) GetCommissionChangeTime() uint64 {
	if m != nil {
		return m.CommissionChangeTime
	}
	return 0
}

func (m *Params) GetLeavePoolTime() uint64 {
	if m != nil {
		return m.LeavePoolTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "kyve.stakers.v1beta1.Params")
}

func init() { proto.RegisterFile("kyve/stakers/v1beta1/params.proto", fileDescriptor_405cabd7005fc18b) }

var fileDescriptor_405cabd7005fc18b = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xc1, 0x4a, 0xe3, 0x40,
	0x1c, 0x06, 0xf0, 0xa4, 0xdb, 0xed, 0xd2, 0xd9, 0x96, 0x85, 0x50, 0x96, 0x22, 0x34, 0xb6, 0x0a,
	0xd2, 0x83, 0x64, 0x28, 0x7a, 0xf2, 0xa8, 0xe8, 0x45, 0x90, 0xd2, 0x8a, 0xa0, 0x97, 0x30, 0x49,
	0x87, 0x64, 0x68, 0x66, 0xfe, 0x21, 0x33, 0x89, 0xed, 0x5b, 0x78, 0xf4, 0xe8, 0xd1, 0x47, 0xf1,
	0xd8, 0xa3, 0x47, 0x69, 0x5f, 0x44, 0x66, 0x26, 0xd4, 0x53, 0xc2, 0xf7, 0xfd, 0xbe, 0x40, 0x66,
	0xd0, 0x68, 0xb9, 0xae, 0x28, 0x96, 0x8a, 0x2c, 0x69, 0x21, 0x71, 0x35, 0x89, 0xa8, 0x22, 0x13,
	0x9c, 0x93, 0x82, 0x70, 0x19, 0xe4, 0x05, 0x28, 0xf0, 0x7a, 0x9a, 0x04, 0x35, 0x09, 0x6a, 0x72,
	0xd0, 0x4b, 0x20, 0x01, 0x03, 0xb0, 0x7e, 0xb3, 0xf6, 0xe8, 0xbd, 0x81, 0x5a, 0x53, 0x33, 0xf6,
	0x06, 0x08, 0x55, 0xa0, 0x68, 0x28, 0x33, 0x22, 0xd3, 0xbe, 0x3b, 0x74, 0xc7, 0xed, 0x59, 0x5b,
	0x27, 0x73, 0x1d, 0x78, 0x23, 0xd4, 0x29, 0xf3, 0x0c, 0xc8, 0xa2, 0x06, 0x0d, 0x03, 0xfe, 0xda,
	0xcc, 0x92, 0x63, 0xd4, 0x55, 0x8c, 0x53, 0x28, 0x55, 0x6d, 0x7e, 0x19, 0xd3, 0xa9, 0x43, 0x8b,
	0x06, 0x08, 0x71, 0xb2, 0x0a, 0x73, 0x60, 0x42, 0xc9, 0x7e, 0x73, 0xe8, 0x8e, 0x9b, 0xb3, 0x36,
	0x27, 0xab, 0xa9, 0x09, 0xbc, 0x73, 0xf4, 0xbf, 0x14, 0x11, 0x88, 0x05, 0x13, 0x49, 0xa8, 0xff,
	0x41, 0x3f, 0xf5, 0x07, 0xfa, 0xbf, 0x0d, 0xed, 0xed, 0xdb, 0xb9, 0x2d, 0xef, 0x19, 0xa7, 0x7a,
	0x15, 0x03, 0xe7, 0x4c, 0x4a, 0x06, 0x22, 0x8c, 0x53, 0x22, 0x12, 0x6a, 0x57, 0x2d, 0xbb, 0xfa,
	0x69, 0xaf, 0x4c, 0x69, 0x56, 0x27, 0xe8, 0x5f, 0x46, 0x49, 0x45, 0xc3, 0x1c, 0x20, 0xb3, 0xfc,
	0x8f, 0xe1, 0x5d, 0x13, 0x4f, 0x01, 0x32, 0xed, 0x2e, 0x9a, 0xaf, 0x6f, 0x87, 0xce, 0xe5, 0xcd,
	0xc7, 0xd6, 0x77, 0x37, 0x5b, 0xdf, 0xfd, 0xda, 0xfa, 0xee, 0xcb, 0xce, 0x77, 0x36, 0x3b, 0xdf,
	0xf9, 0xdc, 0xf9, 0xce, 0xd3, 0x69, 0xc2, 0x54, 0x5a, 0x46, 0x41, 0x0c, 0x1c, 0xdf, 0x3e, 0x3e,
	0x5c, 0xdf, 0x51, 0xf5, 0x0c, 0xc5, 0x12, 0xc7, 0x29, 0x61, 0x02, 0xaf, 0xf6, 0xb7, 0xa5, 0xd6,
	0x39, 0x95, 0x51, 0xcb, 0x9c, 0xfc, 0xd9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x44, 0xe5,
	0x1f, 0xca, 0x01, 0x00, 0x00,
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
	if m.LeavePoolTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LeavePoolTime))
		i--
		dAtA[i] = 0x38
	}
	if m.CommissionChangeTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommissionChangeTime))
		i--
		dAtA[i] = 0x30
	}
	if m.UnbondingStakingTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondingStakingTime))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxPoints != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPoints))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TimeoutSlash) > 0 {
		i -= len(m.TimeoutSlash)
		copy(dAtA[i:], m.TimeoutSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TimeoutSlash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UploadSlash) > 0 {
		i -= len(m.UploadSlash)
		copy(dAtA[i:], m.UploadSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.UploadSlash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VoteSlash) > 0 {
		i -= len(m.VoteSlash)
		copy(dAtA[i:], m.VoteSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VoteSlash)))
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
	l = len(m.VoteSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.UploadSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TimeoutSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxPoints != 0 {
		n += 1 + sovParams(uint64(m.MaxPoints))
	}
	if m.UnbondingStakingTime != 0 {
		n += 1 + sovParams(uint64(m.UnbondingStakingTime))
	}
	if m.CommissionChangeTime != 0 {
		n += 1 + sovParams(uint64(m.CommissionChangeTime))
	}
	if m.LeavePoolTime != 0 {
		n += 1 + sovParams(uint64(m.LeavePoolTime))
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
				return fmt.Errorf("proto: wrong wireType = %d for field VoteSlash", wireType)
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
			m.VoteSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadSlash", wireType)
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
			m.UploadSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutSlash", wireType)
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
			m.TimeoutSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPoints", wireType)
			}
			m.MaxPoints = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPoints |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingStakingTime", wireType)
			}
			m.UnbondingStakingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingStakingTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionChangeTime", wireType)
			}
			m.CommissionChangeTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionChangeTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeavePoolTime", wireType)
			}
			m.LeavePoolTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeavePoolTime |= uint64(b&0x7F) << shift
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
