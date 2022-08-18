// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/pool/v1beta1/gov.proto

package types

import (
	fmt "fmt"
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

// GovMsgCreatePool defines a SDK message for creating a pool.
type GovMsgCreatePool struct {
	// title ...
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// name ...
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// runtime ...
	Runtime string `protobuf:"bytes,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// logo ...
	Logo string `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
	// config ...
	Config string `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
	// start_key ...
	StartKey string `protobuf:"bytes,7,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	// upload_interval ...
	UploadInterval uint64 `protobuf:"varint,8,opt,name=upload_interval,json=uploadInterval,proto3" json:"upload_interval,omitempty"`
	// operating_cost ...
	OperatingCost uint64 `protobuf:"varint,9,opt,name=operating_cost,json=operatingCost,proto3" json:"operating_cost,omitempty"`
	// min_stake ...
	MinStake uint64 `protobuf:"varint,10,opt,name=min_stake,json=minStake,proto3" json:"min_stake,omitempty"`
	// max_bundle_size ...
	MaxBundleSize uint64 `protobuf:"varint,11,opt,name=max_bundle_size,json=maxBundleSize,proto3" json:"max_bundle_size,omitempty"`
	// version ...
	Version string `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
	// binaries ...
	Binaries string `protobuf:"bytes,13,opt,name=binaries,proto3" json:"binaries,omitempty"`
}

func (m *GovMsgCreatePool) Reset()         { *m = GovMsgCreatePool{} }
func (m *GovMsgCreatePool) String() string { return proto.CompactTextString(m) }
func (*GovMsgCreatePool) ProtoMessage()    {}
func (*GovMsgCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_adce52e9478669ec, []int{0}
}
func (m *GovMsgCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GovMsgCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GovMsgCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GovMsgCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovMsgCreatePool.Merge(m, src)
}
func (m *GovMsgCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *GovMsgCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_GovMsgCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_GovMsgCreatePool proto.InternalMessageInfo

func (m *GovMsgCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GovMsgCreatePool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GovMsgCreatePool) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *GovMsgCreatePool) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *GovMsgCreatePool) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *GovMsgCreatePool) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GovMsgCreatePool) GetUploadInterval() uint64 {
	if m != nil {
		return m.UploadInterval
	}
	return 0
}

func (m *GovMsgCreatePool) GetOperatingCost() uint64 {
	if m != nil {
		return m.OperatingCost
	}
	return 0
}

func (m *GovMsgCreatePool) GetMinStake() uint64 {
	if m != nil {
		return m.MinStake
	}
	return 0
}

func (m *GovMsgCreatePool) GetMaxBundleSize() uint64 {
	if m != nil {
		return m.MaxBundleSize
	}
	return 0
}

func (m *GovMsgCreatePool) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GovMsgCreatePool) GetBinaries() string {
	if m != nil {
		return m.Binaries
	}
	return ""
}

type GovMsgCreatePoolResponse struct {
}

func (m *GovMsgCreatePoolResponse) Reset()         { *m = GovMsgCreatePoolResponse{} }
func (m *GovMsgCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*GovMsgCreatePoolResponse) ProtoMessage()    {}
func (*GovMsgCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adce52e9478669ec, []int{1}
}
func (m *GovMsgCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GovMsgCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GovMsgCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GovMsgCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovMsgCreatePoolResponse.Merge(m, src)
}
func (m *GovMsgCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *GovMsgCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GovMsgCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GovMsgCreatePoolResponse proto.InternalMessageInfo

// GovMsgUpdatePool is a gov Content type for updating a pool.
type GovMsgUpdatePool struct {
	// creator ...
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// id ...
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// payload
	Payload string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *GovMsgUpdatePool) Reset()         { *m = GovMsgUpdatePool{} }
func (m *GovMsgUpdatePool) String() string { return proto.CompactTextString(m) }
func (*GovMsgUpdatePool) ProtoMessage()    {}
func (*GovMsgUpdatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_adce52e9478669ec, []int{2}
}
func (m *GovMsgUpdatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GovMsgUpdatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GovMsgUpdatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GovMsgUpdatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovMsgUpdatePool.Merge(m, src)
}
func (m *GovMsgUpdatePool) XXX_Size() int {
	return m.Size()
}
func (m *GovMsgUpdatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_GovMsgUpdatePool.DiscardUnknown(m)
}

var xxx_messageInfo_GovMsgUpdatePool proto.InternalMessageInfo

func (m *GovMsgUpdatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GovMsgUpdatePool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GovMsgUpdatePool) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type GovMsgUpdatePoolResponse struct {
}

func (m *GovMsgUpdatePoolResponse) Reset()         { *m = GovMsgUpdatePoolResponse{} }
func (m *GovMsgUpdatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*GovMsgUpdatePoolResponse) ProtoMessage()    {}
func (*GovMsgUpdatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adce52e9478669ec, []int{3}
}
func (m *GovMsgUpdatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GovMsgUpdatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GovMsgUpdatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GovMsgUpdatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovMsgUpdatePoolResponse.Merge(m, src)
}
func (m *GovMsgUpdatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *GovMsgUpdatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GovMsgUpdatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GovMsgUpdatePoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GovMsgCreatePool)(nil), "kyve.pool.v1beta1.GovMsgCreatePool")
	proto.RegisterType((*GovMsgCreatePoolResponse)(nil), "kyve.pool.v1beta1.GovMsgCreatePoolResponse")
	proto.RegisterType((*GovMsgUpdatePool)(nil), "kyve.pool.v1beta1.GovMsgUpdatePool")
	proto.RegisterType((*GovMsgUpdatePoolResponse)(nil), "kyve.pool.v1beta1.GovMsgUpdatePoolResponse")
}

func init() { proto.RegisterFile("kyve/pool/v1beta1/gov.proto", fileDescriptor_adce52e9478669ec) }

var fileDescriptor_adce52e9478669ec = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x52, 0xba, 0xd6, 0xb0, 0x0e, 0x7c, 0x40, 0xd6, 0x26, 0x45, 0x53, 0x25, 0x60,
	0x5c, 0x1a, 0x4d, 0x7c, 0x83, 0x55, 0x08, 0xa1, 0x09, 0x84, 0x3a, 0x31, 0x09, 0x2e, 0x91, 0x93,
	0x3c, 0x32, 0x2b, 0x89, 0x9f, 0x65, 0x3b, 0xa1, 0xd9, 0xa7, 0xe0, 0x63, 0x71, 0xdc, 0x91, 0x23,
	0x6a, 0xef, 0x7c, 0x06, 0x64, 0x27, 0x61, 0x12, 0x17, 0x6e, 0xfe, 0xfd, 0xff, 0x7f, 0xf9, 0xf9,
	0xf9, 0x3d, 0x72, 0x52, 0xb4, 0x0d, 0x44, 0x0a, 0xb1, 0x8c, 0x9a, 0xf3, 0x04, 0x2c, 0x3f, 0x8f,
	0x72, 0x6c, 0x56, 0x4a, 0xa3, 0x45, 0xfa, 0xd4, 0x99, 0x2b, 0x67, 0xae, 0x7a, 0x73, 0xf9, 0x7b,
	0x4c, 0x9e, 0xbc, 0xc5, 0xe6, 0xbd, 0xc9, 0xd7, 0x1a, 0xb8, 0x85, 0x8f, 0x88, 0x25, 0x65, 0xe4,
	0x20, 0x75, 0x84, 0x9a, 0x05, 0xa7, 0xc1, 0xd9, 0x7c, 0x33, 0x20, 0xa5, 0x64, 0x22, 0x79, 0x05,
	0xec, 0x81, 0x97, 0xfd, 0xd9, 0xa5, 0x75, 0x2d, 0xad, 0xa8, 0x80, 0x4d, 0xba, 0x74, 0x8f, 0x2e,
	0x5d, 0x62, 0x8e, 0xec, 0x61, 0x97, 0x76, 0x67, 0xfa, 0x8c, 0x4c, 0x53, 0x94, 0x5f, 0x45, 0xce,
	0xa6, 0x5e, 0xed, 0x89, 0x9e, 0x90, 0xb9, 0xb1, 0x5c, 0xdb, 0xb8, 0x80, 0x96, 0x1d, 0x78, 0x6b,
	0xe6, 0x85, 0x4b, 0x68, 0xe9, 0x4b, 0x72, 0x54, 0xab, 0x12, 0x79, 0x16, 0x0b, 0x69, 0x41, 0x37,
	0xbc, 0x64, 0xb3, 0xd3, 0xe0, 0x6c, 0xb2, 0x59, 0x74, 0xf2, 0xbb, 0x5e, 0xa5, 0xcf, 0xc9, 0x02,
	0x15, 0x68, 0x6e, 0x85, 0xcc, 0xe3, 0x14, 0x8d, 0x65, 0x73, 0x9f, 0x3b, 0xfc, 0xab, 0xae, 0xd1,
	0x58, 0x57, 0xac, 0x12, 0x32, 0x36, 0x96, 0x17, 0xc0, 0x88, 0x4f, 0xcc, 0x2a, 0x21, 0xaf, 0x1c,
	0xd3, 0x17, 0xe4, 0xa8, 0xe2, 0xdb, 0x38, 0xa9, 0x65, 0x56, 0x42, 0x6c, 0xc4, 0x2d, 0xb0, 0x47,
	0xdd, 0x25, 0x15, 0xdf, 0x5e, 0x78, 0xf5, 0x4a, 0xdc, 0xfa, 0xbe, 0x1b, 0xd0, 0x46, 0xa0, 0x64,
	0x8f, 0xbb, 0xbe, 0x7b, 0xa4, 0xc7, 0x64, 0x96, 0x08, 0xc9, 0xb5, 0x00, 0xc3, 0x0e, 0xbb, 0x56,
	0x06, 0x5e, 0x1e, 0x13, 0xf6, 0xef, 0x7f, 0x6f, 0xc0, 0x28, 0x94, 0x06, 0x96, 0xd7, 0xc3, 0x2c,
	0x3e, 0xa9, 0xec, 0xff, 0xb3, 0x58, 0x90, 0xb1, 0xc8, 0xd8, 0xd8, 0x3f, 0x6d, 0x2c, 0x32, 0x97,
	0x54, 0xbc, 0x75, 0xdf, 0xd1, 0x8f, 0x67, 0xc0, 0xfb, 0x9a, 0xf7, 0xf7, 0x0e, 0x35, 0x2f, 0xd6,
	0x3f, 0x76, 0x61, 0x70, 0xb7, 0x0b, 0x83, 0x5f, 0xbb, 0x30, 0xf8, 0xbe, 0x0f, 0x47, 0x77, 0xfb,
	0x70, 0xf4, 0x73, 0x1f, 0x8e, 0xbe, 0xbc, 0xca, 0x85, 0xbd, 0xa9, 0x93, 0x55, 0x8a, 0x55, 0x74,
	0xf9, 0xf9, 0xfa, 0xcd, 0x07, 0xb0, 0xdf, 0x50, 0x17, 0x51, 0x7a, 0xc3, 0x85, 0x8c, 0xb6, 0xdd,
	0x92, 0xd9, 0x56, 0x81, 0x49, 0xa6, 0x7e, 0xbf, 0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x13, 0x72, 0xc8, 0x7e, 0x02, 0x00, 0x00,
}

func (m *GovMsgCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GovMsgCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GovMsgCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Binaries) > 0 {
		i -= len(m.Binaries)
		copy(dAtA[i:], m.Binaries)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Binaries)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x62
	}
	if m.MaxBundleSize != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.MaxBundleSize))
		i--
		dAtA[i] = 0x58
	}
	if m.MinStake != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.MinStake))
		i--
		dAtA[i] = 0x50
	}
	if m.OperatingCost != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.OperatingCost))
		i--
		dAtA[i] = 0x48
	}
	if m.UploadInterval != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.UploadInterval))
		i--
		dAtA[i] = 0x40
	}
	if len(m.StartKey) > 0 {
		i -= len(m.StartKey)
		copy(dAtA[i:], m.StartKey)
		i = encodeVarintGov(dAtA, i, uint64(len(m.StartKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Config) > 0 {
		i -= len(m.Config)
		copy(dAtA[i:], m.Config)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Config)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Logo) > 0 {
		i -= len(m.Logo)
		copy(dAtA[i:], m.Logo)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Logo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Runtime) > 0 {
		i -= len(m.Runtime)
		copy(dAtA[i:], m.Runtime)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Runtime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GovMsgCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GovMsgCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GovMsgCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GovMsgUpdatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GovMsgUpdatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GovMsgUpdatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GovMsgUpdatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GovMsgUpdatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GovMsgUpdatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GovMsgCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Runtime)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Logo)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.UploadInterval != 0 {
		n += 1 + sovGov(uint64(m.UploadInterval))
	}
	if m.OperatingCost != 0 {
		n += 1 + sovGov(uint64(m.OperatingCost))
	}
	if m.MinStake != 0 {
		n += 1 + sovGov(uint64(m.MinStake))
	}
	if m.MaxBundleSize != 0 {
		n += 1 + sovGov(uint64(m.MaxBundleSize))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Binaries)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *GovMsgCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GovMsgUpdatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovGov(uint64(m.Id))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *GovMsgUpdatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GovMsgCreatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: GovMsgCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GovMsgCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Runtime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadInterval", wireType)
			}
			m.UploadInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UploadInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatingCost", wireType)
			}
			m.OperatingCost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OperatingCost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStake", wireType)
			}
			m.MinStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBundleSize", wireType)
			}
			m.MaxBundleSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBundleSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Binaries", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Binaries = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *GovMsgCreatePoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: GovMsgCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GovMsgCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *GovMsgUpdatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: GovMsgUpdatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GovMsgUpdatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *GovMsgUpdatePoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: GovMsgUpdatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GovMsgUpdatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
