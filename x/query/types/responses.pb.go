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
	// data ...
	Data *types.Pool `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// bundle_proposal ...
	BundleProposal *types1.BundleProposal `protobuf:"bytes,3,opt,name=bundle_proposal,json=bundleProposal,proto3" json:"bundle_proposal,omitempty"`
	// stakers ...
	Stakers []string `protobuf:"bytes,4,rep,name=stakers,proto3" json:"stakers,omitempty"`
	// total_stake ...
	TotalStake uint64 `protobuf:"varint,5,opt,name=total_stake,json=totalStake,proto3" json:"total_stake,omitempty"`
	// status ...
	Status types.PoolStatus `protobuf:"varint,6,opt,name=status,proto3,enum=kyve.pool.v1beta1.PoolStatus" json:"status,omitempty"`
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

func (m *PoolResponse) GetData() *types.Pool {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PoolResponse) GetBundleProposal() *types1.BundleProposal {
	if m != nil {
		return m.BundleProposal
	}
	return nil
}

func (m *PoolResponse) GetStakers() []string {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func (m *PoolResponse) GetTotalStake() uint64 {
	if m != nil {
		return m.TotalStake
	}
	return 0
}

func (m *PoolResponse) GetStatus() types.PoolStatus {
	if m != nil {
		return m.Status
	}
	return types.POOL_STATUS_UNSPECIFIED
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

type StakerPoolResponse struct {
	// staker ...
	Staker *types2.Staker `protobuf:"bytes,1,opt,name=staker,proto3" json:"staker,omitempty"`
	// valaccount ...
	Valaccount *types2.Valaccount `protobuf:"bytes,2,opt,name=valaccount,proto3" json:"valaccount,omitempty"`
}

func (m *StakerPoolResponse) Reset()         { *m = StakerPoolResponse{} }
func (m *StakerPoolResponse) String() string { return proto.CompactTextString(m) }
func (*StakerPoolResponse) ProtoMessage()    {}
func (*StakerPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3263586247a64a9, []int{2}
}
func (m *StakerPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakerPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakerPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakerPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerPoolResponse.Merge(m, src)
}
func (m *StakerPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *StakerPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakerPoolResponse proto.InternalMessageInfo

func (m *StakerPoolResponse) GetStaker() *types2.Staker {
	if m != nil {
		return m.Staker
	}
	return nil
}

func (m *StakerPoolResponse) GetValaccount() *types2.Valaccount {
	if m != nil {
		return m.Valaccount
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolResponse)(nil), "KYVENetwork.chain.query.PoolResponse")
	proto.RegisterType((*StakerResponse)(nil), "KYVENetwork.chain.query.StakerResponse")
	proto.RegisterType((*StakerPoolResponse)(nil), "KYVENetwork.chain.query.StakerPoolResponse")
}

func init() {
	proto.RegisterFile("kyve/query/v1beta1/responses.proto", fileDescriptor_c3263586247a64a9)
}

var fileDescriptor_c3263586247a64a9 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd4, 0x30,
	0x1c, 0xc5, 0x27, 0x9d, 0x71, 0xc4, 0x54, 0x46, 0x08, 0xc2, 0x86, 0x65, 0xad, 0x65, 0xf0, 0x50,
	0x14, 0x1a, 0x76, 0xd4, 0xbb, 0x0c, 0x7a, 0x12, 0x65, 0xe9, 0xc2, 0x82, 0x5e, 0x96, 0x74, 0x26,
	0x74, 0xcb, 0xd4, 0xfe, 0x6b, 0x92, 0x8e, 0xce, 0x57, 0x90, 0x3d, 0xf8, 0xb1, 0x3c, 0xee, 0xd1,
	0xa3, 0xcc, 0x7c, 0x11, 0xe9, 0xbf, 0x69, 0x77, 0x06, 0x14, 0xc1, 0x5b, 0x5f, 0xf2, 0x7b, 0x2f,
	0x2f, 0x49, 0x43, 0xa7, 0xab, 0xcd, 0x5a, 0x89, 0xcf, 0xb5, 0xd2, 0x1b, 0xb1, 0x3e, 0x4d, 0x95,
	0x95, 0xa7, 0x42, 0x2b, 0x53, 0x41, 0x69, 0x94, 0x89, 0x2b, 0x0d, 0x16, 0xd8, 0xd1, 0xdb, 0x0f,
	0x17, 0x6f, 0xde, 0x2b, 0xfb, 0x05, 0xf4, 0x2a, 0x5e, 0x5c, 0xc9, 0xbc, 0x8c, 0xd1, 0x70, 0xfc,
	0x30, 0x83, 0x0c, 0x90, 0x11, 0xcd, 0x57, 0x8b, 0x1f, 0x9f, 0x64, 0x00, 0x59, 0xa1, 0x84, 0xac,
	0x72, 0x21, 0xcb, 0x12, 0xac, 0xb4, 0x39, 0x94, 0xa6, 0x9b, 0xc5, 0x05, 0x2b, 0x80, 0xa2, 0x5f,
	0xaf, 0x11, 0x6e, 0xb6, 0xad, 0x63, 0xac, 0x5c, 0x29, 0x6d, 0x7a, 0xc0, 0xe9, 0x03, 0x26, 0xad,
	0xcb, 0x65, 0xa1, 0x6e, 0x19, 0xa7, 0x5b, 0x66, 0x7a, 0xed, 0xd1, 0xfb, 0x67, 0x00, 0x45, 0xe2,
	0xb6, 0xc2, 0x26, 0xd4, 0xcb, 0x97, 0x9c, 0x84, 0x24, 0x1a, 0x25, 0x5e, 0xbe, 0x64, 0xcf, 0xe8,
	0x68, 0x29, 0xad, 0xe4, 0x5e, 0x48, 0x22, 0x7f, 0x76, 0x14, 0x37, 0x99, 0x31, 0x16, 0x71, 0x81,
	0x31, 0xda, 0x11, 0x62, 0xef, 0xe8, 0x83, 0x36, 0xfe, 0xb2, 0xd2, 0x50, 0x81, 0x91, 0x05, 0x1f,
	0xa2, 0xef, 0x49, 0xeb, 0xeb, 0xd6, 0xee, 0xac, 0x73, 0xd4, 0x67, 0x8e, 0x4d, 0x26, 0xe9, 0x81,
	0x66, 0x9c, 0xde, 0x75, 0x3b, 0xe2, 0xa3, 0x70, 0x18, 0xdd, 0x4b, 0x3a, 0xc9, 0x1e, 0x53, 0xdf,
	0x82, 0x95, 0xc5, 0x25, 0x0e, 0xf0, 0x3b, 0x58, 0x97, 0xe2, 0xd0, 0x79, 0x33, 0xc2, 0x5e, 0xd2,
	0xb1, 0xb1, 0xd2, 0xd6, 0x86, 0x8f, 0x43, 0x12, 0x4d, 0x66, 0x8f, 0xfe, 0x52, 0xfc, 0x1c, 0xa1,
	0xc4, 0xc1, 0xd3, 0x6f, 0x84, 0x4e, 0x30, 0x40, 0xf7, 0x07, 0xf2, 0x02, 0x93, 0x56, 0x4a, 0xe3,
	0xa1, 0xf8, 0xb3, 0x93, 0x36, 0xa9, 0x3b, 0xea, 0x2e, 0xcc, 0xb9, 0x1c, 0xcb, 0xe6, 0xd4, 0x5f,
	0xcb, 0x42, 0x2e, 0x16, 0x50, 0x97, 0xd6, 0x70, 0x2f, 0x1c, 0x46, 0xfe, 0x2c, 0xfc, 0xb3, 0xf5,
	0xa2, 0x07, 0x93, 0x7d, 0xd3, 0xf4, 0x9a, 0x50, 0xd6, 0xc6, 0x1e, 0xdc, 0xd0, 0xff, 0x15, 0x7a,
	0x45, 0xe9, 0x6d, 0xb6, 0xbb, 0xcd, 0x7f, 0xf7, 0xd9, 0xf3, 0xcc, 0x5f, 0xff, 0xd8, 0x06, 0xe4,
	0x66, 0x1b, 0x90, 0x5f, 0xdb, 0x80, 0x7c, 0xdf, 0x05, 0x83, 0x9b, 0x5d, 0x30, 0xf8, 0xb9, 0x0b,
	0x06, 0x1f, 0x9f, 0x66, 0xb9, 0xbd, 0xaa, 0xd3, 0x78, 0x01, 0x9f, 0xc4, 0xde, 0x13, 0x10, 0xf8,
	0x04, 0xc4, 0x57, 0xf7, 0x6a, 0xec, 0xa6, 0x52, 0x26, 0x1d, 0xe3, 0x7f, 0xf7, 0xfc, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0xaf, 0xbf, 0x9d, 0x50, 0x03, 0x00, 0x00,
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
	if m.Status != 0 {
		i = encodeVarintResponses(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.TotalStake != 0 {
		i = encodeVarintResponses(dAtA, i, uint64(m.TotalStake))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Stakers) > 0 {
		for iNdEx := len(m.Stakers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Stakers[iNdEx])
			copy(dAtA[i:], m.Stakers[iNdEx])
			i = encodeVarintResponses(dAtA, i, uint64(len(m.Stakers[iNdEx])))
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
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
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

func (m *StakerPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakerPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakerPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Valaccount != nil {
		{
			size, err := m.Valaccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResponses(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if m.BundleProposal != nil {
		l = m.BundleProposal.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if len(m.Stakers) > 0 {
		for _, s := range m.Stakers {
			l = len(s)
			n += 1 + l + sovResponses(uint64(l))
		}
	}
	if m.TotalStake != 0 {
		n += 1 + sovResponses(uint64(m.TotalStake))
	}
	if m.Status != 0 {
		n += 1 + sovResponses(uint64(m.Status))
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

func (m *StakerPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Staker != nil {
		l = m.Staker.Size()
		n += 1 + l + sovResponses(uint64(l))
	}
	if m.Valaccount != nil {
		l = m.Valaccount.Size()
		n += 1 + l + sovResponses(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			if m.Data == nil {
				m.Data = &types.Pool{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Stakers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
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
				return ErrInvalidLengthResponses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stakers = append(m.Stakers, string(dAtA[iNdEx:postIndex]))
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponses
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.PoolStatus(b&0x7F) << shift
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
func (m *StakerPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StakerPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakerPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Valaccount", wireType)
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
			if m.Valaccount == nil {
				m.Valaccount = &types2.Valaccount{}
			}
			if err := m.Valaccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
