// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/stakers/v1beta1/genesis.proto

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

// GenesisState defines the stakers module's genesis state.
type GenesisState struct {
	Params                  Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	StakerList              []Staker                `protobuf:"bytes,2,rep,name=staker_list,json=stakerList,proto3" json:"staker_list"`
	CommissionChangeEntries []CommissionChangeEntry `protobuf:"bytes,3,rep,name=commission_change_entries,json=commissionChangeEntries,proto3" json:"commission_change_entries"`
	UnbondingStakeEntries   []UnbondingStakeEntry   `protobuf:"bytes,4,rep,name=unbonding_stake_entries,json=unbondingStakeEntries,proto3" json:"unbonding_stake_entries"`
	LeavePoolEntries        []LeavePoolEntry        `protobuf:"bytes,5,rep,name=leave_pool_entries,json=leavePoolEntries,proto3" json:"leave_pool_entries"`
	QueueStateUnstaking     QueueState              `protobuf:"bytes,6,opt,name=queue_state_unstaking,json=queueStateUnstaking,proto3" json:"queue_state_unstaking"`
	QueueStateCommission    QueueState              `protobuf:"bytes,7,opt,name=queue_state_commission,json=queueStateCommission,proto3" json:"queue_state_commission"`
	QueueStateLeave         QueueState              `protobuf:"bytes,8,opt,name=queue_state_leave,json=queueStateLeave,proto3" json:"queue_state_leave"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0deb2ee89d595051, []int{0}
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

func (m *GenesisState) GetStakerList() []Staker {
	if m != nil {
		return m.StakerList
	}
	return nil
}

func (m *GenesisState) GetCommissionChangeEntries() []CommissionChangeEntry {
	if m != nil {
		return m.CommissionChangeEntries
	}
	return nil
}

func (m *GenesisState) GetUnbondingStakeEntries() []UnbondingStakeEntry {
	if m != nil {
		return m.UnbondingStakeEntries
	}
	return nil
}

func (m *GenesisState) GetLeavePoolEntries() []LeavePoolEntry {
	if m != nil {
		return m.LeavePoolEntries
	}
	return nil
}

func (m *GenesisState) GetQueueStateUnstaking() QueueState {
	if m != nil {
		return m.QueueStateUnstaking
	}
	return QueueState{}
}

func (m *GenesisState) GetQueueStateCommission() QueueState {
	if m != nil {
		return m.QueueStateCommission
	}
	return QueueState{}
}

func (m *GenesisState) GetQueueStateLeave() QueueState {
	if m != nil {
		return m.QueueStateLeave
	}
	return QueueState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kyve.stakers.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("kyve/stakers/v1beta1/genesis.proto", fileDescriptor_0deb2ee89d595051)
}

var fileDescriptor_0deb2ee89d595051 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0xb6, 0x15, 0xe4, 0x21, 0x01, 0xa6, 0x63, 0xa1, 0x42, 0xa1, 0x4c, 0x1c, 0x40,
	0xa0, 0x44, 0x83, 0x1b, 0xc7, 0x55, 0x83, 0x03, 0x13, 0x1a, 0x9b, 0x86, 0x60, 0x42, 0x8a, 0x9c,
	0xf0, 0xe4, 0x5a, 0x4d, 0xec, 0x2c, 0x76, 0x0a, 0xfd, 0x16, 0x7c, 0xac, 0xde, 0xe8, 0x91, 0x13,
	0x42, 0xed, 0x17, 0x41, 0x76, 0xd2, 0x24, 0x12, 0xde, 0xa1, 0xb7, 0xc4, 0xfe, 0xbf, 0xdf, 0xef,
	0xd9, 0x4f, 0x46, 0x07, 0x93, 0xd9, 0x14, 0x42, 0xa9, 0xc8, 0x04, 0x0a, 0x19, 0x4e, 0x0f, 0x63,
	0x50, 0xe4, 0x30, 0xa4, 0xc0, 0x41, 0x32, 0x19, 0xe4, 0x85, 0x50, 0x02, 0xf7, 0x75, 0x26, 0xa8,
	0x33, 0x41, 0x9d, 0x19, 0xf4, 0xa9, 0xa0, 0xc2, 0x04, 0x42, 0xfd, 0x55, 0x65, 0x07, 0x4f, 0xac,
	0xbc, 0x9c, 0x14, 0x24, 0xab, 0x71, 0x03, 0xbb, 0x72, 0x8d, 0x37, 0x99, 0x83, 0x5f, 0x3b, 0xe8,
	0xf6, 0xbb, 0xaa, 0x89, 0x73, 0x45, 0x14, 0xe0, 0x37, 0xa8, 0x57, 0x41, 0x3c, 0x77, 0xe8, 0x3e,
	0xdb, 0x7d, 0xf5, 0x28, 0xb0, 0x35, 0x15, 0x9c, 0x9a, 0xcc, 0xd1, 0xf6, 0xfc, 0xcf, 0x63, 0xe7,
	0xac, 0xae, 0xc0, 0x23, 0xb4, 0x5b, 0xe5, 0xa2, 0x94, 0x49, 0xe5, 0xdd, 0x18, 0x6e, 0x5d, 0x0f,
	0x38, 0x37, 0xff, 0x35, 0x00, 0x55, 0xbb, 0x27, 0x4c, 0x2a, 0x9c, 0xa1, 0x87, 0x89, 0xc8, 0x32,
	0x26, 0x25, 0x13, 0x3c, 0x4a, 0xc6, 0x84, 0x53, 0x88, 0x80, 0xab, 0x82, 0x81, 0xf4, 0xb6, 0x0c,
	0xf2, 0x85, 0x1d, 0x39, 0x6a, 0xca, 0x46, 0xa6, 0xea, 0x98, 0xab, 0x62, 0x56, 0x1b, 0xf6, 0x13,
	0xcb, 0x26, 0x03, 0x89, 0x29, 0xda, 0x2f, 0x79, 0x2c, 0xf8, 0x37, 0xc6, 0x69, 0x64, 0x88, 0x8d,
	0x6c, 0xdb, 0xc8, 0x9e, 0xdb, 0x65, 0x17, 0xeb, 0x22, 0x73, 0x90, 0xae, 0x6a, 0xaf, 0xfc, 0x6f,
	0x4b, 0x8b, 0x3e, 0x23, 0x9c, 0x02, 0x99, 0x42, 0x94, 0x0b, 0x91, 0x36, 0x8e, 0x1d, 0xe3, 0x78,
	0x6a, 0x77, 0x9c, 0xe8, 0xfc, 0xa9, 0x10, 0x69, 0x17, 0x7f, 0x37, 0xed, 0xae, 0x6a, 0xf2, 0x25,
	0xda, 0xbb, 0x2a, 0xa1, 0x04, 0xdd, 0xbe, 0x82, 0xa8, 0xe4, 0x9a, 0xc3, 0x38, 0xf5, 0x7a, 0x66,
	0x82, 0x43, 0x3b, 0xfc, 0xa3, 0x2e, 0x31, 0x33, 0xaf, 0xc1, 0xf7, 0xaf, 0x9a, 0x95, 0x8b, 0x35,
	0x02, 0x7f, 0x45, 0x0f, 0xba, 0xec, 0xf6, 0x16, 0xbd, 0x9b, 0x1b, 0xc1, 0xfb, 0x2d, 0xbc, 0x1d,
	0x13, 0x3e, 0x43, 0xf7, 0xba, 0x74, 0x73, 0x32, 0xef, 0xd6, 0x46, 0xe0, 0x3b, 0x2d, 0xd8, 0x5c,
	0xd7, 0xd1, 0xdb, 0xf9, 0xd2, 0x77, 0x17, 0x4b, 0xdf, 0xfd, 0xbb, 0xf4, 0xdd, 0x9f, 0x2b, 0xdf,
	0x59, 0xac, 0x7c, 0xe7, 0xf7, 0xca, 0x77, 0x2e, 0x5f, 0x52, 0xa6, 0xc6, 0x65, 0x1c, 0x24, 0x22,
	0x0b, 0xdf, 0x7f, 0xf9, 0x74, 0xfc, 0x01, 0xd4, 0x77, 0x51, 0x4c, 0xc2, 0x64, 0x4c, 0x18, 0x0f,
	0x7f, 0x34, 0x2f, 0x45, 0xcd, 0x72, 0x90, 0x71, 0xcf, 0x3c, 0x90, 0xd7, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa9, 0x8a, 0x00, 0xb6, 0xb9, 0x03, 0x00, 0x00,
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
	{
		size, err := m.QueueStateLeave.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.QueueStateCommission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.QueueStateUnstaking.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.LeavePoolEntries) > 0 {
		for iNdEx := len(m.LeavePoolEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LeavePoolEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UnbondingStakeEntries) > 0 {
		for iNdEx := len(m.UnbondingStakeEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingStakeEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CommissionChangeEntries) > 0 {
		for iNdEx := len(m.CommissionChangeEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommissionChangeEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StakerList) > 0 {
		for iNdEx := len(m.StakerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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
	if len(m.StakerList) > 0 {
		for _, e := range m.StakerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CommissionChangeEntries) > 0 {
		for _, e := range m.CommissionChangeEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbondingStakeEntries) > 0 {
		for _, e := range m.UnbondingStakeEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LeavePoolEntries) > 0 {
		for _, e := range m.LeavePoolEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.QueueStateUnstaking.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.QueueStateCommission.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.QueueStateLeave.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerList", wireType)
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
			m.StakerList = append(m.StakerList, Staker{})
			if err := m.StakerList[len(m.StakerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionChangeEntries", wireType)
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
			m.CommissionChangeEntries = append(m.CommissionChangeEntries, CommissionChangeEntry{})
			if err := m.CommissionChangeEntries[len(m.CommissionChangeEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingStakeEntries", wireType)
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
			m.UnbondingStakeEntries = append(m.UnbondingStakeEntries, UnbondingStakeEntry{})
			if err := m.UnbondingStakeEntries[len(m.UnbondingStakeEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeavePoolEntries", wireType)
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
			m.LeavePoolEntries = append(m.LeavePoolEntries, LeavePoolEntry{})
			if err := m.LeavePoolEntries[len(m.LeavePoolEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStateUnstaking", wireType)
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
			if err := m.QueueStateUnstaking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStateCommission", wireType)
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
			if err := m.QueueStateCommission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStateLeave", wireType)
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
			if err := m.QueueStateLeave.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
