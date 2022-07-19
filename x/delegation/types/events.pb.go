// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/delegation/v1beta1/events.proto

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

// EventDelegatePool is an event emitted when someone delegates to a protocol node.
type EventDelegatePool struct {
	// pool_id is the unique ID of the pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// node is the account address of the protocol node.
	Node string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventDelegatePool) Reset()         { *m = EventDelegatePool{} }
func (m *EventDelegatePool) String() string { return proto.CompactTextString(m) }
func (*EventDelegatePool) ProtoMessage()    {}
func (*EventDelegatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{0}
}
func (m *EventDelegatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDelegatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDelegatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDelegatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDelegatePool.Merge(m, src)
}
func (m *EventDelegatePool) XXX_Size() int {
	return m.Size()
}
func (m *EventDelegatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDelegatePool.DiscardUnknown(m)
}

var xxx_messageInfo_EventDelegatePool proto.InternalMessageInfo

func (m *EventDelegatePool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EventDelegatePool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventDelegatePool) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *EventDelegatePool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventUndelegatePool is an event emitted when someone undelegates from a protocol node.
type EventUndelegatePool struct {
	// pool_id is the unique ID of the pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// node is the account address of the protocol node.
	Node string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventUndelegatePool) Reset()         { *m = EventUndelegatePool{} }
func (m *EventUndelegatePool) String() string { return proto.CompactTextString(m) }
func (*EventUndelegatePool) ProtoMessage()    {}
func (*EventUndelegatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{1}
}
func (m *EventUndelegatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUndelegatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUndelegatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUndelegatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUndelegatePool.Merge(m, src)
}
func (m *EventUndelegatePool) XXX_Size() int {
	return m.Size()
}
func (m *EventUndelegatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUndelegatePool.DiscardUnknown(m)
}

var xxx_messageInfo_EventUndelegatePool proto.InternalMessageInfo

func (m *EventUndelegatePool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EventUndelegatePool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventUndelegatePool) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *EventUndelegatePool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventRedelegatePool is an event emitted when someone redelegates from one protocol node to another.
type EventRedelegatePool struct {
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// from_pool is the unique ID of the pool the user withdraws its delegation from
	FromPool uint64 `protobuf:"varint,2,opt,name=from_pool,json=fromPool,proto3" json:"from_pool,omitempty"`
	// from_node is the account address of the protocol node the users withdraws from.
	FromNode string `protobuf:"bytes,3,opt,name=from_node,json=fromNode,proto3" json:"from_node,omitempty"`
	// pool_id is the unique ID of the pool of the new pool the user delegates to
	ToPool uint64 `protobuf:"varint,4,opt,name=to_pool,json=toPool,proto3" json:"to_pool,omitempty"`
	// address is the account address of the new staker in the the pool
	ToNode string `protobuf:"bytes,5,opt,name=to_node,json=toNode,proto3" json:"to_node,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventRedelegatePool) Reset()         { *m = EventRedelegatePool{} }
func (m *EventRedelegatePool) String() string { return proto.CompactTextString(m) }
func (*EventRedelegatePool) ProtoMessage()    {}
func (*EventRedelegatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{2}
}
func (m *EventRedelegatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRedelegatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRedelegatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRedelegatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRedelegatePool.Merge(m, src)
}
func (m *EventRedelegatePool) XXX_Size() int {
	return m.Size()
}
func (m *EventRedelegatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRedelegatePool.DiscardUnknown(m)
}

var xxx_messageInfo_EventRedelegatePool proto.InternalMessageInfo

func (m *EventRedelegatePool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventRedelegatePool) GetFromPool() uint64 {
	if m != nil {
		return m.FromPool
	}
	return 0
}

func (m *EventRedelegatePool) GetFromNode() string {
	if m != nil {
		return m.FromNode
	}
	return ""
}

func (m *EventRedelegatePool) GetToPool() uint64 {
	if m != nil {
		return m.ToPool
	}
	return 0
}

func (m *EventRedelegatePool) GetToNode() string {
	if m != nil {
		return m.ToNode
	}
	return ""
}

func (m *EventRedelegatePool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*EventDelegatePool)(nil), "kyve.registry.v1beta1.EventDelegatePool")
	proto.RegisterType((*EventUndelegatePool)(nil), "kyve.registry.v1beta1.EventUndelegatePool")
	proto.RegisterType((*EventRedelegatePool)(nil), "kyve.registry.v1beta1.EventRedelegatePool")
}

func init() {
	proto.RegisterFile("kyve/delegation/v1beta1/events.proto", fileDescriptor_d01988a9108a2e89)
}

var fileDescriptor_d01988a9108a2e89 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xcd, 0x6a, 0xea, 0x40,
	0x14, 0xc7, 0x9d, 0x7b, 0x73, 0xe3, 0x75, 0x76, 0x4d, 0x3f, 0x0c, 0x16, 0x82, 0x48, 0x17, 0xae,
	0x32, 0x48, 0xdf, 0xa0, 0xd4, 0x85, 0x14, 0xa4, 0x04, 0x5a, 0x68, 0x37, 0x12, 0x9d, 0xd3, 0x18,
	0xd4, 0x1c, 0x99, 0x1c, 0xad, 0xbe, 0x45, 0xdf, 0xa5, 0x2f, 0xd1, 0xa5, 0xcb, 0x2e, 0x8b, 0xbe,
	0x48, 0x99, 0x99, 0x48, 0xe3, 0xc2, 0x6d, 0x77, 0x73, 0xce, 0xfc, 0xfe, 0x1f, 0xc3, 0xf0, 0xab,
	0xc9, 0x7a, 0x09, 0x42, 0xc2, 0x14, 0x92, 0x98, 0x52, 0xcc, 0xc4, 0xb2, 0x33, 0x04, 0x8a, 0x3b,
	0x02, 0x96, 0x90, 0x51, 0x1e, 0xce, 0x15, 0x12, 0x7a, 0xe7, 0x9a, 0x0a, 0x15, 0x24, 0x69, 0x4e,
	0x6a, 0x1d, 0x16, 0x4c, 0xe3, 0x2c, 0xc1, 0x04, 0x0d, 0x21, 0xf4, 0xc9, 0xc2, 0x8d, 0xf6, 0x31,
	0xcb, 0x9f, 0x55, 0x41, 0x36, 0x8f, 0x91, 0xb4, 0xb2, 0x44, 0x4b, 0xf1, 0x93, 0xae, 0x2e, 0x72,
	0x6b, 0x19, 0xb8, 0x47, 0x9c, 0x7a, 0x75, 0x5e, 0x9d, 0x23, 0x4e, 0x07, 0xa9, 0xf4, 0x59, 0x93,
	0xb5, 0x9d, 0xc8, 0xd5, 0x63, 0x4f, 0x7a, 0x3e, 0xaf, 0xc6, 0x52, 0x2a, 0xc8, 0x73, 0xff, 0x4f,
	0x93, 0xb5, 0x6b, 0xd1, 0x7e, 0xf4, 0x3c, 0xee, 0x64, 0x28, 0xc1, 0xff, 0x6b, 0xd6, 0xe6, 0xec,
	0x5d, 0x70, 0x37, 0x9e, 0xe1, 0x22, 0x23, 0xdf, 0xb1, 0x2e, 0x76, 0x6a, 0x11, 0x3f, 0x35, 0x99,
	0x0f, 0x99, 0xfc, 0xc5, 0xd4, 0x77, 0x56, 0xc4, 0x46, 0x70, 0x10, 0x5b, 0x72, 0x67, 0x87, 0xee,
	0x97, 0xbc, 0xf6, 0xa2, 0x70, 0x36, 0xd0, 0x35, 0x4c, 0xb2, 0x13, 0xfd, 0xd7, 0x0b, 0x23, 0xdb,
	0x5f, 0x96, 0xf2, 0xcd, 0x65, 0x5f, 0x77, 0xa8, 0xf3, 0x2a, 0xa1, 0xd5, 0x15, 0x25, 0x08, 0xf7,
	0x6f, 0x24, 0xb4, 0x9a, 0x7f, 0x46, 0xe3, 0x12, 0xf6, 0x0f, 0x5b, 0xbb, 0xe5, 0xd6, 0x37, 0xbd,
	0x8f, 0x6d, 0xc0, 0x36, 0xdb, 0x80, 0x7d, 0x6d, 0x03, 0xf6, 0xb6, 0x0b, 0x2a, 0x9b, 0x5d, 0x50,
	0xf9, 0xdc, 0x05, 0x95, 0x67, 0x91, 0xa4, 0x34, 0x5e, 0x0c, 0xc3, 0x11, 0xce, 0xc4, 0xdd, 0xd3,
	0x63, 0xb7, 0x0f, 0xf4, 0x8a, 0x6a, 0x22, 0x46, 0xe3, 0x38, 0xcd, 0xc4, 0xaa, 0xfc, 0xeb, 0xb4,
	0x9e, 0x43, 0x3e, 0x74, 0xcd, 0x8f, 0x5f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x43, 0xeb,
	0x58, 0x92, 0x02, 0x00, 0x00,
}

func (m *EventDelegatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDelegatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDelegatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUndelegatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUndelegatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUndelegatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRedelegatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRedelegatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRedelegatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ToNode) > 0 {
		i -= len(m.ToNode)
		copy(dAtA[i:], m.ToNode)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ToNode)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ToPool != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ToPool))
		i--
		dAtA[i] = 0x20
	}
	if len(m.FromNode) > 0 {
		i -= len(m.FromNode)
		copy(dAtA[i:], m.FromNode)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FromNode)))
		i--
		dAtA[i] = 0x1a
	}
	if m.FromPool != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.FromPool))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventDelegatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovEvents(uint64(m.PoolId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventUndelegatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovEvents(uint64(m.PoolId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventRedelegatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.FromPool != 0 {
		n += 1 + sovEvents(uint64(m.FromPool))
	}
	l = len(m.FromNode)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ToPool != 0 {
		n += 1 + sovEvents(uint64(m.ToPool))
	}
	l = len(m.ToNode)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDelegatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDelegatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDelegatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUndelegatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUndelegatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUndelegatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRedelegatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRedelegatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRedelegatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromPool", wireType)
			}
			m.FromPool = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromPool |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromNode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromNode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToPool", wireType)
			}
			m.ToPool = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToPool |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToNode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToNode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
