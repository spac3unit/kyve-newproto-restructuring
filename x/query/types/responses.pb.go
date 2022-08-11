// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/query/v1beta1/responses.proto

package types

import (
	fmt "fmt"
	types1 "github.com/KYVENetwork/chain/x/bundles/types"
	types "github.com/KYVENetwork/chain/x/pool/types"
	types2 "github.com/KYVENetwork/chain/x/stakers/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PoolResponse struct {
	// id ...
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// pool ...
	Pool *types.Pool `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
	// bundle_proposal ...
	BundleProposal *types1.BundleProposal `protobuf:"bytes,3,opt,name=bundle_proposal,json=bundleProposal,proto3" json:"bundle_proposal,omitempty"`
	// valaccounts ...
	Valaccounts []*types2.Valaccount `protobuf:"bytes,4,rep,name=valaccounts,proto3" json:"valaccounts,omitempty"`
	// total_stake ...
	TotalStake uint64 `protobuf:"varint,5,opt,name=total_stake,json=totalStake,proto3" json:"total_stake,omitempty"`
}

func (m *PoolResponse) Reset()         { *m = PoolResponse{} }
func (m *PoolResponse) String() string { return proto.CompactTextString(m) }
func (*PoolResponse) ProtoMessage()    {}
func (*PoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3263586247a64a9, []int{0}
}
func (m *PoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolResponse.Merge(m, src)
}
func (m *PoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *PoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PoolResponse proto.InternalMessageInfo

func (m *PoolResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PoolResponse) GetPool() *types.Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

func (m *PoolResponse) GetBundleProposal() *types1.BundleProposal {
	if m != nil {
		return m.BundleProposal
	}
	return nil
}

func (m *PoolResponse) GetValaccounts() []*types2.Valaccount {
	if m != nil {
		return m.Valaccounts
	}
	return nil
}

func (m *PoolResponse) GetTotalStake() uint64 {
	if m != nil {
		return m.TotalStake
	}
	return 0
}

type StakerResponse struct {
	// staker ...
	Staker *types2.Staker `protobuf:"bytes,1,opt,name=staker,proto3" json:"staker,omitempty"`
	// valaccounts ...
	Valaccounts []*types2.Valaccount `protobuf:"bytes,2,rep,name=valaccounts,proto3" json:"valaccounts,omitempty"`
}

func (m *StakerResponse) Reset()         { *m = StakerResponse{} }
func (m *StakerResponse) String() string { return proto.CompactTextString(m) }
func (*StakerResponse) ProtoMessage()    {}
func (*StakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3263586247a64a9, []int{1}
}
func (m *StakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerResponse.Merge(m, src)
}
func (m *StakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *StakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakerResponse proto.InternalMessageInfo

func (m *StakerResponse) GetStaker() *types2.Staker {
	if m != nil {
		return m.Staker
	}
	return nil
}

func (m *StakerResponse) GetValaccounts() []*types2.Valaccount {
	if m != nil {
		return m.Valaccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolResponse)(nil), "KYVENetwork.chain.query.PoolResponse")
	proto.RegisterType((*StakerResponse)(nil), "KYVENetwork.chain.query.StakerResponse")
}

func init() {
	proto.RegisterFile("kyve/query/v1beta1/responses.proto", fileDescriptor_c3263586247a64a9)
}

var fileDescriptor_c3263586247a64a9 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0x69, 0xff, 0x2e, 0x26, 0x3f, 0x15, 0x82, 0xd0, 0x50, 0x4a, 0x0c, 0xc5, 0x45,
	0x51, 0xc8, 0xd0, 0xea, 0x13, 0x14, 0x5d, 0x89, 0x52, 0x22, 0x14, 0x74, 0x53, 0x26, 0xe9, 0x90,
	0x86, 0xc6, 0xdc, 0x98, 0x99, 0x54, 0xfb, 0x0a, 0xae, 0x7c, 0x2c, 0x97, 0x5d, 0xba, 0x94, 0xf6,
	0x41, 0x94, 0x4c, 0x26, 0xb1, 0x15, 0x57, 0xee, 0x72, 0x67, 0xbe, 0x73, 0xce, 0xbd, 0x93, 0x8b,
	0x7b, 0x8b, 0xd5, 0x92, 0x91, 0xc7, 0x8c, 0xa5, 0x2b, 0xb2, 0x1c, 0x78, 0x4c, 0xd0, 0x01, 0x49,
	0x19, 0x4f, 0x20, 0xe6, 0x8c, 0x3b, 0x49, 0x0a, 0x02, 0x8c, 0xf6, 0xd5, 0xdd, 0xe4, 0xf2, 0x86,
	0x89, 0x27, 0x48, 0x17, 0x8e, 0x3f, 0xa7, 0x61, 0xec, 0x48, 0x41, 0xe7, 0x30, 0x80, 0x00, 0x24,
	0x43, 0xf2, 0xaf, 0x02, 0xef, 0x74, 0x03, 0x80, 0x20, 0x62, 0x84, 0x26, 0x21, 0xa1, 0x71, 0x0c,
	0x82, 0x8a, 0x10, 0x62, 0x5e, 0xde, 0xca, 0xc0, 0x04, 0x20, 0xaa, 0xf2, 0xf2, 0x42, 0xdd, 0x16,
	0xed, 0x70, 0x41, 0x17, 0x2c, 0xe5, 0x15, 0xa0, 0xea, 0x3d, 0xc6, 0xcb, 0xe2, 0x59, 0xc4, 0xbe,
	0x19, 0x55, 0x17, 0x4c, 0xef, 0x13, 0xe1, 0xff, 0x63, 0x80, 0xc8, 0x55, 0xa3, 0x18, 0x2d, 0xac,
	0x85, 0x33, 0x13, 0xd9, 0xa8, 0xdf, 0x70, 0xb5, 0x70, 0x66, 0x9c, 0xe2, 0x46, 0x1e, 0x6b, 0x6a,
	0x36, 0xea, 0xeb, 0xc3, 0xb6, 0x93, 0x7b, 0x3a, 0xb2, 0x11, 0x65, 0xe8, 0x48, 0xb9, 0x84, 0x8c,
	0x6b, 0x7c, 0x50, 0xd8, 0x4f, 0x93, 0x14, 0x12, 0xe0, 0x34, 0x32, 0xeb, 0x52, 0x77, 0x5c, 0xe8,
	0xca, 0xec, 0x52, 0x3a, 0x92, 0xf5, 0x58, 0xb1, 0x6e, 0xcb, 0xdb, 0xab, 0x8d, 0x11, 0xd6, 0x97,
	0x34, 0xa2, 0xbe, 0x0f, 0x59, 0x2c, 0xb8, 0xd9, 0xb0, 0xeb, 0x7d, 0x7d, 0x68, 0x17, 0x56, 0xe5,
	0xa8, 0xa5, 0xd5, 0xa4, 0x02, 0xdd, 0x5d, 0x91, 0x71, 0x84, 0x75, 0x01, 0x82, 0x46, 0x53, 0x29,
	0x30, 0xff, 0xc9, 0xc1, 0xb0, 0x3c, 0xba, 0xcd, 0x4f, 0x7a, 0x2f, 0x08, 0xb7, 0xe4, 0x57, 0x5a,
	0xbd, 0xc1, 0x39, 0x6e, 0x16, 0xf6, 0xf2, 0x1d, 0xf4, 0x61, 0xf7, 0xf7, 0x48, 0xa5, 0x52, 0xec,
	0xcf, 0x6e, 0xb5, 0x3f, 0x74, 0x3b, 0xba, 0x78, 0xdb, 0x58, 0x68, 0xbd, 0xb1, 0xd0, 0xc7, 0xc6,
	0x42, 0xaf, 0x5b, 0xab, 0xb6, 0xde, 0x5a, 0xb5, 0xf7, 0xad, 0x55, 0xbb, 0x3f, 0x09, 0x42, 0x31,
	0xcf, 0x3c, 0xc7, 0x87, 0x07, 0xb2, 0xb3, 0x66, 0x44, 0xae, 0x19, 0x79, 0x56, 0x9b, 0x29, 0x56,
	0x09, 0xe3, 0x5e, 0x53, 0xfe, 0xdb, 0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xe2, 0xc9,
	0x79, 0xb4, 0x02, 0x00, 0x00,
}

func (m *PoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalStake != 0 {
		i = encodeVarintResponses(dAtA, i, uint64(m.TotalStake))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Valaccounts) > 0 {
		for iNdEx := len(m.Valaccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valaccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResponses(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BundleProposal != nil {
		{
			size, err := m.BundleProposal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResponses(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Pool != nil {
		{
			size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResponses(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintResponses(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Valaccounts) > 0 {
		for iNdEx := len(m.Valaccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valaccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResponses(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Staker != nil {
		{
			size, err := m.Staker.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResponses(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponses(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponses(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovResponses(uint64(m.Id))
	}
	if m.Pool != nil {
		l = m.Pool.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if m.BundleProposal != nil {
		l = m.BundleProposal.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if len(m.Valaccounts) > 0 {
		for _, e := range m.Valaccounts {
			l = e.Size()
			n += 1 + l + sovResponses(uint64(l))
		}
	}
	if m.TotalStake != 0 {
		n += 1 + sovResponses(uint64(m.TotalStake))
	}
	return n
}

func (m *StakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Staker != nil {
		l = m.Staker.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if len(m.Valaccounts) > 0 {
		for _, e := range m.Valaccounts {
			l = e.Size()
			n += 1 + l + sovResponses(uint64(l))
		}
	}
	return n
}

func sovResponses(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponses(x uint64) (n int) {
	return sovResponses(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponses
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
			return fmt.Errorf("proto: PoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pool == nil {
				m.Pool = &types.Pool{}
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleProposal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BundleProposal == nil {
				m.BundleProposal = &types1.BundleProposal{}
			}
			if err := m.BundleProposal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valaccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Valaccounts = append(m.Valaccounts, &types2.Valaccount{})
			if err := m.Valaccounts[len(m.Valaccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStake", wireType)
			}
			m.TotalStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResponses(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponses
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
func (m *StakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponses
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
			return fmt.Errorf("proto: StakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Staker == nil {
				m.Staker = &types2.Staker{}
			}
			if err := m.Staker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valaccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Valaccounts = append(m.Valaccounts, &types2.Valaccount{})
			if err := m.Valaccounts[len(m.Valaccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponses(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponses
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
func skipResponses(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponses
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
					return 0, ErrIntOverflowResponses
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
					return 0, ErrIntOverflowResponses
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
				return 0, ErrInvalidLengthResponses
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponses
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponses
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponses        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponses          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponses = fmt.Errorf("proto: unexpected end of group")
)
