// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/bundles/v1beta1/events.proto

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

// BundleStatus ...
type BundleStatus int32

const (
	// BUNDLE_STATUS_UNSPECIFIED ...
	BUNDLE_STATUS_UNSPECIFIED BundleStatus = 0
	// BUNDLE_STATUS_VALID ...
	BUNDLE_STATUS_VALID BundleStatus = 1
	// BUNDLE_STATUS_INVALID ...
	BUNDLE_STATUS_INVALID BundleStatus = 2
	// BUNDLE_STATUS_NO_FUNDS ...
	BUNDLE_STATUS_NO_FUNDS BundleStatus = 3
	// BUNDLE_STATUS_NO_QUORUM ...
	BUNDLE_STATUS_NO_QUORUM BundleStatus = 4
)

var BundleStatus_name = map[int32]string{
	0: "BUNDLE_STATUS_UNSPECIFIED",
	1: "BUNDLE_STATUS_VALID",
	2: "BUNDLE_STATUS_INVALID",
	3: "BUNDLE_STATUS_NO_FUNDS",
	4: "BUNDLE_STATUS_NO_QUORUM",
}

var BundleStatus_value = map[string]int32{
	"BUNDLE_STATUS_UNSPECIFIED": 0,
	"BUNDLE_STATUS_VALID":       1,
	"BUNDLE_STATUS_INVALID":     2,
	"BUNDLE_STATUS_NO_FUNDS":    3,
	"BUNDLE_STATUS_NO_QUORUM":   4,
}

func (x BundleStatus) String() string {
	return proto.EnumName(BundleStatus_name, int32(x))
}

func (BundleStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a02f505e55d81e92, []int{0}
}

// EventBundleVote is an event emitted when a protocol node votes on a bundle.
type EventBundleVote struct {
	// pool_id is the unique ID of the pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// address is the account address of the protocol node.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// storage_id is the unique ID of the bundle.
	StorageId string `protobuf:"bytes,3,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	// vote is the vote type of the protocol node.
	Vote VoteType `protobuf:"varint,4,opt,name=vote,proto3,enum=kyve.bundles.v1beta1.VoteType" json:"vote,omitempty"`
}

func (m *EventBundleVote) Reset()         { *m = EventBundleVote{} }
func (m *EventBundleVote) String() string { return proto.CompactTextString(m) }
func (*EventBundleVote) ProtoMessage()    {}
func (*EventBundleVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02f505e55d81e92, []int{0}
}
func (m *EventBundleVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBundleVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBundleVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBundleVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBundleVote.Merge(m, src)
}
func (m *EventBundleVote) XXX_Size() int {
	return m.Size()
}
func (m *EventBundleVote) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBundleVote.DiscardUnknown(m)
}

var xxx_messageInfo_EventBundleVote proto.InternalMessageInfo

func (m *EventBundleVote) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EventBundleVote) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventBundleVote) GetStorageId() string {
	if m != nil {
		return m.StorageId
	}
	return ""
}

func (m *EventBundleVote) GetVote() VoteType {
	if m != nil {
		return m.Vote
	}
	return VOTE_TYPE_UNSPECIFIED
}

// EventBundleFinalised is an event emitted when a bundle is finalised.
type EventBundleFinalised struct {
	// pool_id is the unique ID of the pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// storage_id ...
	StorageId string `protobuf:"bytes,2,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	// byte_size ...
	ByteSize uint64 `protobuf:"varint,3,opt,name=byte_size,json=byteSize,proto3" json:"byte_size,omitempty"`
	// uploader ...
	Uploader string `protobuf:"bytes,4,opt,name=uploader,proto3" json:"uploader,omitempty"`
	// next_uploader ...
	NextUploader string `protobuf:"bytes,5,opt,name=next_uploader,json=nextUploader,proto3" json:"next_uploader,omitempty"`
	// reward ...
	Reward uint64 `protobuf:"varint,6,opt,name=reward,proto3" json:"reward,omitempty"`
	// valid ...
	Valid uint64 `protobuf:"varint,7,opt,name=valid,proto3" json:"valid,omitempty"`
	// invalid ...
	Invalid uint64 `protobuf:"varint,8,opt,name=invalid,proto3" json:"invalid,omitempty"`
	// from_height ...
	FromHeight uint64 `protobuf:"varint,9,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	// to_height ...
	ToHeight uint64 `protobuf:"varint,10,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	// status ...
	Status BundleStatus `protobuf:"varint,11,opt,name=status,proto3,enum=kyve.bundles.v1beta1.BundleStatus" json:"status,omitempty"`
	// to_key ...
	ToKey string `protobuf:"bytes,12,opt,name=to_key,json=toKey,proto3" json:"to_key,omitempty"`
	// to_value ...
	ToValue string `protobuf:"bytes,13,opt,name=to_value,json=toValue,proto3" json:"to_value,omitempty"`
	// id ...
	Id uint64 `protobuf:"varint,14,opt,name=id,proto3" json:"id,omitempty"`
	// bundle_hash ...
	BundleHash string `protobuf:"bytes,15,opt,name=bundle_hash,json=bundleHash,proto3" json:"bundle_hash,omitempty"`
	// abstain ...
	Abstain uint64 `protobuf:"varint,16,opt,name=abstain,proto3" json:"abstain,omitempty"`
	// total ...
	Total uint64 `protobuf:"varint,17,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *EventBundleFinalised) Reset()         { *m = EventBundleFinalised{} }
func (m *EventBundleFinalised) String() string { return proto.CompactTextString(m) }
func (*EventBundleFinalised) ProtoMessage()    {}
func (*EventBundleFinalised) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02f505e55d81e92, []int{1}
}
func (m *EventBundleFinalised) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBundleFinalised) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBundleFinalised.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBundleFinalised) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBundleFinalised.Merge(m, src)
}
func (m *EventBundleFinalised) XXX_Size() int {
	return m.Size()
}
func (m *EventBundleFinalised) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBundleFinalised.DiscardUnknown(m)
}

var xxx_messageInfo_EventBundleFinalised proto.InternalMessageInfo

func (m *EventBundleFinalised) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EventBundleFinalised) GetStorageId() string {
	if m != nil {
		return m.StorageId
	}
	return ""
}

func (m *EventBundleFinalised) GetByteSize() uint64 {
	if m != nil {
		return m.ByteSize
	}
	return 0
}

func (m *EventBundleFinalised) GetUploader() string {
	if m != nil {
		return m.Uploader
	}
	return ""
}

func (m *EventBundleFinalised) GetNextUploader() string {
	if m != nil {
		return m.NextUploader
	}
	return ""
}

func (m *EventBundleFinalised) GetReward() uint64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *EventBundleFinalised) GetValid() uint64 {
	if m != nil {
		return m.Valid
	}
	return 0
}

func (m *EventBundleFinalised) GetInvalid() uint64 {
	if m != nil {
		return m.Invalid
	}
	return 0
}

func (m *EventBundleFinalised) GetFromHeight() uint64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *EventBundleFinalised) GetToHeight() uint64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

func (m *EventBundleFinalised) GetStatus() BundleStatus {
	if m != nil {
		return m.Status
	}
	return BUNDLE_STATUS_UNSPECIFIED
}

func (m *EventBundleFinalised) GetToKey() string {
	if m != nil {
		return m.ToKey
	}
	return ""
}

func (m *EventBundleFinalised) GetToValue() string {
	if m != nil {
		return m.ToValue
	}
	return ""
}

func (m *EventBundleFinalised) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventBundleFinalised) GetBundleHash() string {
	if m != nil {
		return m.BundleHash
	}
	return ""
}

func (m *EventBundleFinalised) GetAbstain() uint64 {
	if m != nil {
		return m.Abstain
	}
	return 0
}

func (m *EventBundleFinalised) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterEnum("kyve.bundles.v1beta1.BundleStatus", BundleStatus_name, BundleStatus_value)
	proto.RegisterType((*EventBundleVote)(nil), "kyve.bundles.v1beta1.EventBundleVote")
	proto.RegisterType((*EventBundleFinalised)(nil), "kyve.bundles.v1beta1.EventBundleFinalised")
}

func init() { proto.RegisterFile("kyve/bundles/v1beta1/events.proto", fileDescriptor_a02f505e55d81e92) }

var fileDescriptor_a02f505e55d81e92 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x4f, 0xdb, 0x3e,
	0x1c, 0x6e, 0x4a, 0x28, 0xed, 0x8f, 0xb7, 0xfe, 0xfd, 0x2f, 0x60, 0x8a, 0xc8, 0x18, 0xbb, 0xa0,
	0x69, 0x6a, 0x04, 0xbb, 0xed, 0x06, 0x6b, 0x2b, 0x2a, 0x58, 0xd9, 0x52, 0x52, 0x69, 0xbb, 0x44,
	0x2e, 0xf1, 0x1a, 0x8b, 0x10, 0x57, 0xb1, 0x5b, 0x28, 0xb7, 0xdd, 0x76, 0xdc, 0x6d, 0xd7, 0x49,
	0xfb, 0x32, 0x3b, 0x72, 0xdc, 0x71, 0x82, 0x2f, 0x32, 0xd9, 0x4e, 0x11, 0xdd, 0xd8, 0x2d, 0xcf,
	0x8b, 0xed, 0xe7, 0xf7, 0xc4, 0x86, 0xa7, 0xe7, 0xe3, 0x11, 0x75, 0x7b, 0xc3, 0x24, 0x8c, 0xa9,
	0x70, 0x47, 0xbb, 0x3d, 0x2a, 0xc9, 0xae, 0x4b, 0x47, 0x34, 0x91, 0xa2, 0x36, 0x48, 0xb9, 0xe4,
	0xa8, 0xa2, 0x2c, 0xb5, 0xcc, 0x52, 0xcb, 0x2c, 0xd5, 0x4a, 0x9f, 0xf7, 0xb9, 0x36, 0xb8, 0xea,
	0xcb, 0x78, 0xab, 0x9b, 0x8f, 0x6e, 0x27, 0xaf, 0x8c, 0xbc, 0xfd, 0xd5, 0x82, 0xe5, 0x86, 0xda,
	0xfb, 0x40, 0x3b, 0xba, 0x5c, 0x52, 0xb4, 0x06, 0x73, 0x03, 0xce, 0xe3, 0x80, 0x85, 0xd8, 0xda,
	0xb2, 0x76, 0x6c, 0xaf, 0xa0, 0x60, 0x2b, 0x44, 0x18, 0xe6, 0x48, 0x18, 0xa6, 0x54, 0x08, 0x9c,
	0xdf, 0xb2, 0x76, 0x4a, 0xde, 0x04, 0xa2, 0x4d, 0x00, 0x21, 0x79, 0x4a, 0xfa, 0x54, 0xad, 0x9a,
	0xd1, 0x62, 0x29, 0x63, 0x5a, 0x21, 0xda, 0x03, 0x7b, 0xc4, 0x25, 0xc5, 0xf6, 0x96, 0xb5, 0xb3,
	0xb4, 0xe7, 0xd4, 0x1e, 0xcb, 0x5f, 0x53, 0x67, 0x9f, 0x8e, 0x07, 0xd4, 0xd3, 0xde, 0xed, 0x4f,
	0x36, 0x54, 0x1e, 0x24, 0x6b, 0xb2, 0x84, 0xc4, 0x4c, 0xd0, 0xf0, 0xdf, 0xf1, 0xa6, 0x43, 0xe4,
	0xff, 0x0c, 0xb1, 0x01, 0xa5, 0xde, 0x58, 0xd2, 0x40, 0xb0, 0x6b, 0xaa, 0x23, 0xda, 0x5e, 0x51,
	0x11, 0x1d, 0x76, 0x4d, 0x51, 0x15, 0x8a, 0xc3, 0x41, 0xcc, 0x49, 0x48, 0x53, 0x9d, 0xb2, 0xe4,
	0xdd, 0x63, 0xf4, 0x0c, 0x16, 0x13, 0x7a, 0x25, 0x83, 0x7b, 0xc3, 0xac, 0x36, 0x2c, 0x28, 0xd2,
	0x9f, 0x98, 0x56, 0xa1, 0x90, 0xd2, 0x4b, 0x92, 0x86, 0xb8, 0x60, 0x42, 0x19, 0x84, 0x2a, 0x30,
	0x3b, 0x22, 0x31, 0x0b, 0xf1, 0x9c, 0xa6, 0x0d, 0x50, 0x4d, 0xb2, 0xc4, 0xf0, 0x45, 0xcd, 0x4f,
	0x20, 0x7a, 0x02, 0xf3, 0x1f, 0x53, 0x7e, 0x11, 0x44, 0x94, 0xf5, 0x23, 0x89, 0x4b, 0x5a, 0x05,
	0x45, 0x1d, 0x6a, 0x46, 0x8d, 0x21, 0xf9, 0x44, 0x06, 0x33, 0x86, 0xe4, 0x99, 0xf8, 0x0a, 0x0a,
	0x42, 0x12, 0x39, 0x14, 0x78, 0x5e, 0x57, 0xbd, 0xfd, 0x78, 0xd5, 0xa6, 0xd2, 0x8e, 0x76, 0x7a,
	0xd9, 0x0a, 0xb4, 0x02, 0x05, 0xc9, 0x83, 0x73, 0x3a, 0xc6, 0x0b, 0x7a, 0xbe, 0x59, 0xc9, 0x8f,
	0xe8, 0x18, 0xad, 0x43, 0x51, 0xf2, 0x60, 0x44, 0xe2, 0x21, 0xc5, 0x8b, 0xe6, 0xaf, 0x4b, 0xde,
	0x55, 0x10, 0x2d, 0x41, 0x9e, 0x85, 0x78, 0x49, 0x67, 0xc8, 0x9b, 0xec, 0xe6, 0xa4, 0x20, 0x22,
	0x22, 0xc2, 0xcb, 0xda, 0x0d, 0x86, 0x3a, 0x24, 0x22, 0xd2, 0x17, 0xa8, 0x27, 0x24, 0x61, 0x09,
	0x2e, 0x9b, 0xb1, 0x33, 0xa8, 0x6a, 0x92, 0x5c, 0x92, 0x18, 0xff, 0x67, 0x6a, 0xd2, 0xe0, 0xf9,
	0x37, 0x0b, 0x16, 0x1e, 0x66, 0x45, 0x9b, 0xb0, 0x7e, 0xe0, 0xb7, 0xeb, 0xc7, 0x8d, 0xa0, 0x73,
	0xba, 0x7f, 0xea, 0x77, 0x02, 0xbf, 0xdd, 0x79, 0xdb, 0x78, 0xdd, 0x6a, 0xb6, 0x1a, 0xf5, 0x72,
	0x0e, 0xad, 0xc1, 0xff, 0xd3, 0x72, 0x77, 0xff, 0xb8, 0x55, 0x2f, 0x5b, 0x68, 0x1d, 0x56, 0xa6,
	0x85, 0x56, 0xdb, 0x48, 0x79, 0x54, 0x85, 0xd5, 0x69, 0xa9, 0x7d, 0x12, 0x34, 0xfd, 0x76, 0xbd,
	0x53, 0x9e, 0x41, 0x1b, 0xb0, 0xf6, 0x97, 0xf6, 0xce, 0x3f, 0xf1, 0xfc, 0x37, 0x65, 0xbb, 0x6a,
	0x7f, 0xfe, 0xee, 0xe4, 0x0e, 0x9a, 0x3f, 0x6e, 0x1d, 0xeb, 0xe6, 0xd6, 0xb1, 0x7e, 0xdd, 0x3a,
	0xd6, 0x97, 0x3b, 0x27, 0x77, 0x73, 0xe7, 0xe4, 0x7e, 0xde, 0x39, 0xb9, 0x0f, 0x2f, 0xfa, 0x4c,
	0x46, 0xc3, 0x5e, 0xed, 0x8c, 0x5f, 0xb8, 0x47, 0xef, 0xbb, 0x8d, 0x36, 0x95, 0x97, 0x3c, 0x3d,
	0x77, 0xcf, 0x22, 0xc2, 0x12, 0xf7, 0xea, 0xfe, 0x4d, 0xca, 0xf1, 0x80, 0x8a, 0x5e, 0x41, 0xbf,
	0xc7, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xcc, 0x5a, 0xfe, 0xff, 0x03, 0x00, 0x00,
}

func (m *EventBundleVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBundleVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBundleVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x20
	}
	if len(m.StorageId) > 0 {
		i -= len(m.StorageId)
		copy(dAtA[i:], m.StorageId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.StorageId)))
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

func (m *EventBundleFinalised) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBundleFinalised) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBundleFinalised) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.Abstain != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Abstain))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.BundleHash) > 0 {
		i -= len(m.BundleHash)
		copy(dAtA[i:], m.BundleHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BundleHash)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x70
	}
	if len(m.ToValue) > 0 {
		i -= len(m.ToValue)
		copy(dAtA[i:], m.ToValue)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ToValue)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ToKey) > 0 {
		i -= len(m.ToKey)
		copy(dAtA[i:], m.ToKey)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ToKey)))
		i--
		dAtA[i] = 0x62
	}
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x58
	}
	if m.ToHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ToHeight))
		i--
		dAtA[i] = 0x50
	}
	if m.FromHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.FromHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.Invalid != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Invalid))
		i--
		dAtA[i] = 0x40
	}
	if m.Valid != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Valid))
		i--
		dAtA[i] = 0x38
	}
	if m.Reward != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Reward))
		i--
		dAtA[i] = 0x30
	}
	if len(m.NextUploader) > 0 {
		i -= len(m.NextUploader)
		copy(dAtA[i:], m.NextUploader)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NextUploader)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Uploader) > 0 {
		i -= len(m.Uploader)
		copy(dAtA[i:], m.Uploader)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Uploader)))
		i--
		dAtA[i] = 0x22
	}
	if m.ByteSize != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ByteSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.StorageId) > 0 {
		i -= len(m.StorageId)
		copy(dAtA[i:], m.StorageId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.StorageId)))
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
func (m *EventBundleVote) Size() (n int) {
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
	l = len(m.StorageId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Vote != 0 {
		n += 1 + sovEvents(uint64(m.Vote))
	}
	return n
}

func (m *EventBundleFinalised) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovEvents(uint64(m.PoolId))
	}
	l = len(m.StorageId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ByteSize != 0 {
		n += 1 + sovEvents(uint64(m.ByteSize))
	}
	l = len(m.Uploader)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NextUploader)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Reward != 0 {
		n += 1 + sovEvents(uint64(m.Reward))
	}
	if m.Valid != 0 {
		n += 1 + sovEvents(uint64(m.Valid))
	}
	if m.Invalid != 0 {
		n += 1 + sovEvents(uint64(m.Invalid))
	}
	if m.FromHeight != 0 {
		n += 1 + sovEvents(uint64(m.FromHeight))
	}
	if m.ToHeight != 0 {
		n += 1 + sovEvents(uint64(m.ToHeight))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	l = len(m.ToKey)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ToValue)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.BundleHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Abstain != 0 {
		n += 2 + sovEvents(uint64(m.Abstain))
	}
	if m.Total != 0 {
		n += 2 + sovEvents(uint64(m.Total))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventBundleVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventBundleVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBundleVote: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field StorageId", wireType)
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
			m.StorageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= VoteType(b&0x7F) << shift
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
func (m *EventBundleFinalised) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventBundleFinalised: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBundleFinalised: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field StorageId", wireType)
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
			m.StorageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByteSize", wireType)
			}
			m.ByteSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ByteSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uploader", wireType)
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
			m.Uploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextUploader", wireType)
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
			m.NextUploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			m.Reward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			m.Valid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Valid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invalid", wireType)
			}
			m.Invalid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Invalid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromHeight", wireType)
			}
			m.FromHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToHeight", wireType)
			}
			m.ToHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BundleStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToKey", wireType)
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
			m.ToKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToValue", wireType)
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
			m.ToValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleHash", wireType)
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
			m.BundleHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abstain", wireType)
			}
			m.Abstain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Abstain |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
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
