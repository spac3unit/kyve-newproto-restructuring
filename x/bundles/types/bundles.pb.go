// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/bundles/v1beta1/bundles.proto

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
	return fileDescriptor_889cf76d77a4de2b, []int{0}
}

// BundleProposal ...
type BundleProposal struct {
	// pool_id ...
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// storage_id ...
	StorageId string `protobuf:"bytes,2,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	// uploader ...
	Uploader string `protobuf:"bytes,3,opt,name=uploader,proto3" json:"uploader,omitempty"`
	// next_uploader ...
	NextUploader string `protobuf:"bytes,4,opt,name=next_uploader,json=nextUploader,proto3" json:"next_uploader,omitempty"`
	// byte_size ...
	ByteSize uint64 `protobuf:"varint,5,opt,name=byte_size,json=byteSize,proto3" json:"byte_size,omitempty"`
	// to_height ...
	ToHeight uint64 `protobuf:"varint,6,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	// to_key ...
	ToKey string `protobuf:"bytes,7,opt,name=to_key,json=toKey,proto3" json:"to_key,omitempty"`
	// to_value ...
	ToValue string `protobuf:"bytes,8,opt,name=to_value,json=toValue,proto3" json:"to_value,omitempty"`
	// bundle_hash ...
	BundleHash string `protobuf:"bytes,9,opt,name=bundle_hash,json=bundleHash,proto3" json:"bundle_hash,omitempty"`
	// created_at ...
	CreatedAt uint64 `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// voters_valid ...
	VotersValid []string `protobuf:"bytes,11,rep,name=voters_valid,json=votersValid,proto3" json:"voters_valid,omitempty"`
	// voters_invalid ...
	VotersInvalid []string `protobuf:"bytes,12,rep,name=voters_invalid,json=votersInvalid,proto3" json:"voters_invalid,omitempty"`
	// voters_abstain ...
	VotersAbstain []string `protobuf:"bytes,13,rep,name=voters_abstain,json=votersAbstain,proto3" json:"voters_abstain,omitempty"`
}

func (m *BundleProposal) Reset()         { *m = BundleProposal{} }
func (m *BundleProposal) String() string { return proto.CompactTextString(m) }
func (*BundleProposal) ProtoMessage()    {}
func (*BundleProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_889cf76d77a4de2b, []int{0}
}
func (m *BundleProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BundleProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BundleProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BundleProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleProposal.Merge(m, src)
}
func (m *BundleProposal) XXX_Size() int {
	return m.Size()
}
func (m *BundleProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleProposal.DiscardUnknown(m)
}

var xxx_messageInfo_BundleProposal proto.InternalMessageInfo

func (m *BundleProposal) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *BundleProposal) GetStorageId() string {
	if m != nil {
		return m.StorageId
	}
	return ""
}

func (m *BundleProposal) GetUploader() string {
	if m != nil {
		return m.Uploader
	}
	return ""
}

func (m *BundleProposal) GetNextUploader() string {
	if m != nil {
		return m.NextUploader
	}
	return ""
}

func (m *BundleProposal) GetByteSize() uint64 {
	if m != nil {
		return m.ByteSize
	}
	return 0
}

func (m *BundleProposal) GetToHeight() uint64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

func (m *BundleProposal) GetToKey() string {
	if m != nil {
		return m.ToKey
	}
	return ""
}

func (m *BundleProposal) GetToValue() string {
	if m != nil {
		return m.ToValue
	}
	return ""
}

func (m *BundleProposal) GetBundleHash() string {
	if m != nil {
		return m.BundleHash
	}
	return ""
}

func (m *BundleProposal) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *BundleProposal) GetVotersValid() []string {
	if m != nil {
		return m.VotersValid
	}
	return nil
}

func (m *BundleProposal) GetVotersInvalid() []string {
	if m != nil {
		return m.VotersInvalid
	}
	return nil
}

func (m *BundleProposal) GetVotersAbstain() []string {
	if m != nil {
		return m.VotersAbstain
	}
	return nil
}

// Proposal ...
type FinalizedBundle struct {
	// pool_id ...
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// id ...
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// storage_id ...
	StorageId string `protobuf:"bytes,3,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	// uploader ...
	Uploader string `protobuf:"bytes,4,opt,name=uploader,proto3" json:"uploader,omitempty"`
	// from_height ...
	FromHeight uint64 `protobuf:"varint,5,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	// to_height ...
	ToHeight uint64 `protobuf:"varint,6,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	// key ...
	Key string `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	// value ...
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	// bundle_hash ...
	BundleHash string `protobuf:"bytes,9,opt,name=bundle_hash,json=bundleHash,proto3" json:"bundle_hash,omitempty"`
	// finalized_at ...
	FinalizedAt uint64 `protobuf:"varint,10,opt,name=finalized_at,json=finalizedAt,proto3" json:"finalized_at,omitempty"`
}

func (m *FinalizedBundle) Reset()         { *m = FinalizedBundle{} }
func (m *FinalizedBundle) String() string { return proto.CompactTextString(m) }
func (*FinalizedBundle) ProtoMessage()    {}
func (*FinalizedBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_889cf76d77a4de2b, []int{1}
}
func (m *FinalizedBundle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalizedBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalizedBundle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalizedBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalizedBundle.Merge(m, src)
}
func (m *FinalizedBundle) XXX_Size() int {
	return m.Size()
}
func (m *FinalizedBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalizedBundle.DiscardUnknown(m)
}

var xxx_messageInfo_FinalizedBundle proto.InternalMessageInfo

func (m *FinalizedBundle) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *FinalizedBundle) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FinalizedBundle) GetStorageId() string {
	if m != nil {
		return m.StorageId
	}
	return ""
}

func (m *FinalizedBundle) GetUploader() string {
	if m != nil {
		return m.Uploader
	}
	return ""
}

func (m *FinalizedBundle) GetFromHeight() uint64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *FinalizedBundle) GetToHeight() uint64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

func (m *FinalizedBundle) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FinalizedBundle) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FinalizedBundle) GetBundleHash() string {
	if m != nil {
		return m.BundleHash
	}
	return ""
}

func (m *FinalizedBundle) GetFinalizedAt() uint64 {
	if m != nil {
		return m.FinalizedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("kyve.bundles.v1beta1.BundleStatus", BundleStatus_name, BundleStatus_value)
	proto.RegisterType((*BundleProposal)(nil), "kyve.bundles.v1beta1.BundleProposal")
	proto.RegisterType((*FinalizedBundle)(nil), "kyve.bundles.v1beta1.FinalizedBundle")
}

func init() {
	proto.RegisterFile("kyve/bundles/v1beta1/bundles.proto", fileDescriptor_889cf76d77a4de2b)
}

var fileDescriptor_889cf76d77a4de2b = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x4e, 0xdb, 0x3c,
	0x18, 0xc6, 0x9b, 0xb6, 0x94, 0xf6, 0x6d, 0xe1, 0xab, 0xfc, 0xc1, 0x08, 0x45, 0x84, 0x3f, 0xd3,
	0x24, 0x34, 0x4d, 0x44, 0x68, 0x57, 0x50, 0xd6, 0x56, 0x44, 0xb0, 0xc0, 0x1a, 0x52, 0x69, 0x3b,
	0x89, 0x1c, 0x62, 0x1a, 0x8b, 0x10, 0x57, 0x89, 0xdb, 0x51, 0xae, 0x60, 0x87, 0xbb, 0x83, 0x4d,
	0x9a, 0xb4, 0x6b, 0xd9, 0x21, 0x87, 0x3b, 0x9c, 0xe0, 0x46, 0x26, 0xdb, 0x69, 0xa1, 0x4c, 0x63,
	0x67, 0x79, 0x7f, 0xcf, 0x4f, 0xb1, 0xfc, 0x3e, 0x32, 0x6c, 0x5f, 0x8c, 0x47, 0xc4, 0xf4, 0x87,
	0x71, 0x10, 0x91, 0xd4, 0x1c, 0xed, 0xf9, 0x84, 0xe3, 0xbd, 0xc9, 0xbc, 0x3b, 0x48, 0x18, 0x67,
	0x68, 0x49, 0x38, 0xbb, 0x13, 0x96, 0x39, 0x8d, 0xa5, 0x3e, 0xeb, 0x33, 0x29, 0x98, 0xe2, 0x4b,
	0xb9, 0xdb, 0xdf, 0x0b, 0xb0, 0xb8, 0x2f, 0xcd, 0x93, 0x84, 0x0d, 0x58, 0x8a, 0x23, 0xb4, 0x02,
	0xf3, 0x03, 0xc6, 0x22, 0x8f, 0x06, 0xba, 0xb6, 0xa9, 0xed, 0x14, 0xbb, 0x25, 0x31, 0x5a, 0x01,
	0x5a, 0x07, 0x48, 0x39, 0x4b, 0x70, 0x9f, 0x88, 0x2c, 0xbf, 0xa9, 0xed, 0x54, 0xba, 0x95, 0x8c,
	0x58, 0x01, 0x6a, 0x40, 0x79, 0x38, 0x88, 0x18, 0x0e, 0x48, 0xa2, 0x17, 0x64, 0x38, 0x9d, 0xd1,
	0x73, 0x58, 0x88, 0xc9, 0x15, 0xf7, 0xa6, 0x42, 0x51, 0x0a, 0x35, 0x01, 0xdd, 0x89, 0xb4, 0x06,
	0x15, 0x7f, 0xcc, 0x89, 0x97, 0xd2, 0x6b, 0xa2, 0xcf, 0xc9, 0xa3, 0xcb, 0x02, 0x38, 0xf4, 0x9a,
	0x88, 0x90, 0x33, 0x2f, 0x24, 0xb4, 0x1f, 0x72, 0xbd, 0xa4, 0x42, 0xce, 0x0e, 0xe4, 0x8c, 0x96,
	0xa1, 0xc4, 0x99, 0x77, 0x41, 0xc6, 0xfa, 0xbc, 0xfc, 0xef, 0x1c, 0x67, 0x87, 0x64, 0x8c, 0x56,
	0xa1, 0xcc, 0x99, 0x37, 0xc2, 0xd1, 0x90, 0xe8, 0x65, 0x19, 0xcc, 0x73, 0xd6, 0x13, 0x23, 0xda,
	0x80, 0xaa, 0x5a, 0x90, 0x17, 0xe2, 0x34, 0xd4, 0x2b, 0x32, 0x05, 0x85, 0x0e, 0x70, 0x1a, 0x8a,
	0xcb, 0x9e, 0x25, 0x04, 0x73, 0x12, 0x78, 0x98, 0xeb, 0x20, 0x0f, 0xac, 0x64, 0xa4, 0xc9, 0xd1,
	0x16, 0xd4, 0x46, 0x8c, 0x93, 0x24, 0x15, 0xbf, 0xa7, 0x81, 0x5e, 0xdd, 0x2c, 0xec, 0x54, 0xba,
	0x55, 0xc5, 0x7a, 0x02, 0xa1, 0x17, 0xb0, 0x98, 0x29, 0x34, 0x56, 0x52, 0x4d, 0x4a, 0x0b, 0x8a,
	0x5a, 0x0a, 0x3e, 0xd0, 0xb0, 0x9f, 0x72, 0x4c, 0x63, 0x7d, 0xe1, 0xa1, 0xd6, 0x54, 0x70, 0xfb,
	0x4b, 0x1e, 0xfe, 0xeb, 0xd0, 0x18, 0x47, 0xf4, 0x9a, 0x04, 0xaa, 0xb1, 0xbf, 0x37, 0xb5, 0x08,
	0xf9, 0xac, 0xa1, 0x62, 0x37, 0x4f, 0x1f, 0x37, 0x57, 0x78, 0xaa, 0xb9, 0xe2, 0xa3, 0xe6, 0x36,
	0xa0, 0x7a, 0x9e, 0xb0, 0xcb, 0xc9, 0xe6, 0x55, 0x2d, 0x20, 0x50, 0xb6, 0xfb, 0x27, 0x8b, 0xa9,
	0x43, 0xe1, 0xbe, 0x15, 0xf1, 0x89, 0x96, 0x60, 0xee, 0x61, 0x21, 0x6a, 0xf8, 0x77, 0x1d, 0x5b,
	0x50, 0x3b, 0x9f, 0xdc, 0xfe, 0xbe, 0x90, 0xea, 0x94, 0x35, 0xf9, 0xcb, 0xaf, 0x1a, 0xd4, 0xd4,
	0x62, 0x1c, 0x8e, 0xf9, 0x30, 0x45, 0xeb, 0xb0, 0xba, 0xef, 0xda, 0xad, 0xa3, 0xb6, 0xe7, 0x9c,
	0x36, 0x4f, 0x5d, 0xc7, 0x73, 0x6d, 0xe7, 0xa4, 0xfd, 0xc6, 0xea, 0x58, 0xed, 0x56, 0x3d, 0x87,
	0x56, 0xe0, 0xff, 0xd9, 0xb8, 0xd7, 0x3c, 0xb2, 0x5a, 0x75, 0x0d, 0xad, 0xc2, 0xf2, 0x6c, 0x60,
	0xd9, 0x2a, 0xca, 0xa3, 0x06, 0x3c, 0x9b, 0x8d, 0xec, 0x63, 0xaf, 0xe3, 0xda, 0x2d, 0xa7, 0x5e,
	0x40, 0x6b, 0xb0, 0xf2, 0x47, 0xf6, 0xce, 0x3d, 0xee, 0xba, 0x6f, 0xeb, 0xc5, 0x46, 0xf1, 0xd3,
	0x37, 0x23, 0xb7, 0xdf, 0xf9, 0x71, 0x6b, 0x68, 0x37, 0xb7, 0x86, 0xf6, 0xeb, 0xd6, 0xd0, 0x3e,
	0xdf, 0x19, 0xb9, 0x9b, 0x3b, 0x23, 0xf7, 0xf3, 0xce, 0xc8, 0x7d, 0x78, 0xd5, 0xa7, 0x3c, 0x1c,
	0xfa, 0xbb, 0x67, 0xec, 0xd2, 0x3c, 0x7c, 0xdf, 0x6b, 0xdb, 0x84, 0x7f, 0x64, 0xc9, 0x85, 0x79,
	0x16, 0x62, 0x1a, 0x9b, 0x57, 0xd3, 0x17, 0xcf, 0xc7, 0x03, 0x92, 0xfa, 0x25, 0xf9, 0x78, 0x5f,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xf6, 0x51, 0x7a, 0x0e, 0x04, 0x00, 0x00,
}

func (m *BundleProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BundleProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BundleProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VotersAbstain) > 0 {
		for iNdEx := len(m.VotersAbstain) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotersAbstain[iNdEx])
			copy(dAtA[i:], m.VotersAbstain[iNdEx])
			i = encodeVarintBundles(dAtA, i, uint64(len(m.VotersAbstain[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.VotersInvalid) > 0 {
		for iNdEx := len(m.VotersInvalid) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotersInvalid[iNdEx])
			copy(dAtA[i:], m.VotersInvalid[iNdEx])
			i = encodeVarintBundles(dAtA, i, uint64(len(m.VotersInvalid[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.VotersValid) > 0 {
		for iNdEx := len(m.VotersValid) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotersValid[iNdEx])
			copy(dAtA[i:], m.VotersValid[iNdEx])
			i = encodeVarintBundles(dAtA, i, uint64(len(m.VotersValid[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.CreatedAt != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x50
	}
	if len(m.BundleHash) > 0 {
		i -= len(m.BundleHash)
		copy(dAtA[i:], m.BundleHash)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.BundleHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ToValue) > 0 {
		i -= len(m.ToValue)
		copy(dAtA[i:], m.ToValue)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.ToValue)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ToKey) > 0 {
		i -= len(m.ToKey)
		copy(dAtA[i:], m.ToKey)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.ToKey)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ToHeight != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.ToHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.ByteSize != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.ByteSize))
		i--
		dAtA[i] = 0x28
	}
	if len(m.NextUploader) > 0 {
		i -= len(m.NextUploader)
		copy(dAtA[i:], m.NextUploader)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.NextUploader)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uploader) > 0 {
		i -= len(m.Uploader)
		copy(dAtA[i:], m.Uploader)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.Uploader)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StorageId) > 0 {
		i -= len(m.StorageId)
		copy(dAtA[i:], m.StorageId)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.StorageId)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FinalizedBundle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalizedBundle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalizedBundle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalizedAt != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.FinalizedAt))
		i--
		dAtA[i] = 0x50
	}
	if len(m.BundleHash) > 0 {
		i -= len(m.BundleHash)
		copy(dAtA[i:], m.BundleHash)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.BundleHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ToHeight != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.ToHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.FromHeight != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.FromHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Uploader) > 0 {
		i -= len(m.Uploader)
		copy(dAtA[i:], m.Uploader)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.Uploader)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StorageId) > 0 {
		i -= len(m.StorageId)
		copy(dAtA[i:], m.StorageId)
		i = encodeVarintBundles(dAtA, i, uint64(len(m.StorageId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintBundles(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBundles(dAtA []byte, offset int, v uint64) int {
	offset -= sovBundles(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BundleProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovBundles(uint64(m.PoolId))
	}
	l = len(m.StorageId)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.Uploader)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.NextUploader)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	if m.ByteSize != 0 {
		n += 1 + sovBundles(uint64(m.ByteSize))
	}
	if m.ToHeight != 0 {
		n += 1 + sovBundles(uint64(m.ToHeight))
	}
	l = len(m.ToKey)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.ToValue)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.BundleHash)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovBundles(uint64(m.CreatedAt))
	}
	if len(m.VotersValid) > 0 {
		for _, s := range m.VotersValid {
			l = len(s)
			n += 1 + l + sovBundles(uint64(l))
		}
	}
	if len(m.VotersInvalid) > 0 {
		for _, s := range m.VotersInvalid {
			l = len(s)
			n += 1 + l + sovBundles(uint64(l))
		}
	}
	if len(m.VotersAbstain) > 0 {
		for _, s := range m.VotersAbstain {
			l = len(s)
			n += 1 + l + sovBundles(uint64(l))
		}
	}
	return n
}

func (m *FinalizedBundle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovBundles(uint64(m.PoolId))
	}
	if m.Id != 0 {
		n += 1 + sovBundles(uint64(m.Id))
	}
	l = len(m.StorageId)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.Uploader)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	if m.FromHeight != 0 {
		n += 1 + sovBundles(uint64(m.FromHeight))
	}
	if m.ToHeight != 0 {
		n += 1 + sovBundles(uint64(m.ToHeight))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	l = len(m.BundleHash)
	if l > 0 {
		n += 1 + l + sovBundles(uint64(l))
	}
	if m.FinalizedAt != 0 {
		n += 1 + sovBundles(uint64(m.FinalizedAt))
	}
	return n
}

func sovBundles(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBundles(x uint64) (n int) {
	return sovBundles(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BundleProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBundles
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
			return fmt.Errorf("proto: BundleProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BundleProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uploader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextUploader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextUploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByteSize", wireType)
			}
			m.ByteSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToHeight", wireType)
			}
			m.ToHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundleHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotersValid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotersValid = append(m.VotersValid, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotersInvalid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotersInvalid = append(m.VotersInvalid, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotersAbstain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotersAbstain = append(m.VotersAbstain, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBundles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBundles
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
func (m *FinalizedBundle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBundles
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
			return fmt.Errorf("proto: FinalizedBundle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalizedBundle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uploader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromHeight", wireType)
			}
			m.FromHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToHeight", wireType)
			}
			m.ToHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
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
				return ErrInvalidLengthBundles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBundles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundleHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedAt", wireType)
			}
			m.FinalizedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBundles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBundles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBundles
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
func skipBundles(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBundles
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
					return 0, ErrIntOverflowBundles
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
					return 0, ErrIntOverflowBundles
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
				return 0, ErrInvalidLengthBundles
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBundles
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBundles
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBundles        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBundles          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBundles = fmt.Errorf("proto: unexpected end of group")
)
