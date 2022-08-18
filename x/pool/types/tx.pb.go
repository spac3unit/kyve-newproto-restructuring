// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/pool/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgFundPool defines a SDK message for funding a pool.
type MsgFundPool struct {
	// creator ...
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// id ...
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgFundPool) Reset()         { *m = MsgFundPool{} }
func (m *MsgFundPool) String() string { return proto.CompactTextString(m) }
func (*MsgFundPool) ProtoMessage()    {}
func (*MsgFundPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ddefdf83388ddc, []int{0}
}
func (m *MsgFundPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundPool.Merge(m, src)
}
func (m *MsgFundPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundPool proto.InternalMessageInfo

func (m *MsgFundPool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgFundPool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgFundPool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgFundPoolResponse defines the Msg/DefundPool response type.
type MsgFundPoolResponse struct {
}

func (m *MsgFundPoolResponse) Reset()         { *m = MsgFundPoolResponse{} }
func (m *MsgFundPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundPoolResponse) ProtoMessage()    {}
func (*MsgFundPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ddefdf83388ddc, []int{1}
}
func (m *MsgFundPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundPoolResponse.Merge(m, src)
}
func (m *MsgFundPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundPoolResponse proto.InternalMessageInfo

// MsgDefundPool defines a SDK message for defunding a pool.
type MsgDefundPool struct {
	// creator ...
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// id ...
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgDefundPool) Reset()         { *m = MsgDefundPool{} }
func (m *MsgDefundPool) String() string { return proto.CompactTextString(m) }
func (*MsgDefundPool) ProtoMessage()    {}
func (*MsgDefundPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ddefdf83388ddc, []int{2}
}
func (m *MsgDefundPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDefundPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDefundPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDefundPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDefundPool.Merge(m, src)
}
func (m *MsgDefundPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgDefundPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDefundPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDefundPool proto.InternalMessageInfo

func (m *MsgDefundPool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDefundPool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgDefundPool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgDefundPoolResponse defines the Msg/DefundPool response type.
type MsgDefundPoolResponse struct {
}

func (m *MsgDefundPoolResponse) Reset()         { *m = MsgDefundPoolResponse{} }
func (m *MsgDefundPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDefundPoolResponse) ProtoMessage()    {}
func (*MsgDefundPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ddefdf83388ddc, []int{3}
}
func (m *MsgDefundPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDefundPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDefundPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDefundPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDefundPoolResponse.Merge(m, src)
}
func (m *MsgDefundPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDefundPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDefundPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDefundPoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFundPool)(nil), "kyve.pool.v1beta1.MsgFundPool")
	proto.RegisterType((*MsgFundPoolResponse)(nil), "kyve.pool.v1beta1.MsgFundPoolResponse")
	proto.RegisterType((*MsgDefundPool)(nil), "kyve.pool.v1beta1.MsgDefundPool")
	proto.RegisterType((*MsgDefundPoolResponse)(nil), "kyve.pool.v1beta1.MsgDefundPoolResponse")
}

func init() { proto.RegisterFile("kyve/pool/v1beta1/tx.proto", fileDescriptor_20ddefdf83388ddc) }

var fileDescriptor_20ddefdf83388ddc = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xae, 0x2c, 0x4b,
	0xd5, 0x2f, 0xc8, 0xcf, 0xcf, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0xc9, 0xe9, 0x81, 0xe4, 0xf4, 0xa0, 0x72,
	0x52, 0xd2, 0x98, 0xca, 0xd3, 0xf3, 0xcb, 0x20, 0xea, 0x95, 0xfc, 0xb9, 0xb8, 0x7d, 0x8b, 0xd3,
	0xdd, 0x4a, 0xf3, 0x52, 0x02, 0xf2, 0xf3, 0x73, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13,
	0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x3e, 0x2e, 0xa6,
	0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x31, 0x2e, 0xb6,
	0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66, 0xb0, 0x18, 0x94, 0xa7, 0x24, 0xca, 0x25, 0x8c,
	0x64, 0x60, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x52, 0x20, 0x17, 0xaf, 0x6f, 0x71,
	0xba, 0x4b, 0x6a, 0x1a, 0xf5, 0x6c, 0x12, 0xe7, 0x12, 0x45, 0x31, 0x12, 0x66, 0x97, 0xd1, 0x67,
	0x26, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x20, 0x2e, 0x0e, 0xb8, 0xc7, 0xe4, 0xf4, 0x30, 0x02,
	0x46, 0x0f, 0xc9, 0x9d, 0x52, 0x6a, 0xf8, 0xe5, 0x61, 0x66, 0x0b, 0x45, 0x70, 0x71, 0x21, 0x79,
	0x42, 0x01, 0xbb, 0x2e, 0x84, 0x0a, 0x29, 0x0d, 0x42, 0x2a, 0xe0, 0x26, 0xc7, 0x71, 0x71, 0x39,
	0x83, 0x42, 0x20, 0x15, 0x6c, 0xb2, 0x32, 0x16, 0x7d, 0xee, 0xf9, 0x65, 0xbe, 0xc5, 0xe9, 0x08,
	0x45, 0x52, 0xda, 0x44, 0x28, 0x42, 0x36, 0x3f, 0xb4, 0x20, 0x85, 0xb0, 0xf9, 0x08, 0x45, 0x78,
	0xcc, 0x47, 0x28, 0x82, 0x99, 0xef, 0xe4, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xde, 0x91,
	0x61, 0xae, 0x7e, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0xfa, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa,
	0x15, 0x90, 0xa4, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0x95, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xbb, 0x8c, 0xd5, 0xe3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// FundPool ...
	FundPool(ctx context.Context, in *MsgFundPool, opts ...grpc.CallOption) (*MsgFundPoolResponse, error)
	// DefundPool ...
	DefundPool(ctx context.Context, in *MsgDefundPool, opts ...grpc.CallOption) (*MsgDefundPoolResponse, error)
	// CreatePool ...
	CreatePool(ctx context.Context, in *GovMsgCreatePool, opts ...grpc.CallOption) (*GovMsgCreatePoolResponse, error)
	// UpdatePool ...
	UpdatePool(ctx context.Context, in *GovMsgUpdatePool, opts ...grpc.CallOption) (*GovMsgUpdatePoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FundPool(ctx context.Context, in *MsgFundPool, opts ...grpc.CallOption) (*MsgFundPoolResponse, error) {
	out := new(MsgFundPoolResponse)
	err := c.cc.Invoke(ctx, "/kyve.pool.v1beta1.Msg/FundPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DefundPool(ctx context.Context, in *MsgDefundPool, opts ...grpc.CallOption) (*MsgDefundPoolResponse, error) {
	out := new(MsgDefundPoolResponse)
	err := c.cc.Invoke(ctx, "/kyve.pool.v1beta1.Msg/DefundPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePool(ctx context.Context, in *GovMsgCreatePool, opts ...grpc.CallOption) (*GovMsgCreatePoolResponse, error) {
	out := new(GovMsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/kyve.pool.v1beta1.Msg/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePool(ctx context.Context, in *GovMsgUpdatePool, opts ...grpc.CallOption) (*GovMsgUpdatePoolResponse, error) {
	out := new(GovMsgUpdatePoolResponse)
	err := c.cc.Invoke(ctx, "/kyve.pool.v1beta1.Msg/UpdatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// FundPool ...
	FundPool(context.Context, *MsgFundPool) (*MsgFundPoolResponse, error)
	// DefundPool ...
	DefundPool(context.Context, *MsgDefundPool) (*MsgDefundPoolResponse, error)
	// CreatePool ...
	CreatePool(context.Context, *GovMsgCreatePool) (*GovMsgCreatePoolResponse, error)
	// UpdatePool ...
	UpdatePool(context.Context, *GovMsgUpdatePool) (*GovMsgUpdatePoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) FundPool(ctx context.Context, req *MsgFundPool) (*MsgFundPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundPool not implemented")
}
func (*UnimplementedMsgServer) DefundPool(ctx context.Context, req *MsgDefundPool) (*MsgDefundPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefundPool not implemented")
}
func (*UnimplementedMsgServer) CreatePool(ctx context.Context, req *GovMsgCreatePool) (*GovMsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (*UnimplementedMsgServer) UpdatePool(ctx context.Context, req *GovMsgUpdatePool) (*GovMsgUpdatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_FundPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.pool.v1beta1.Msg/FundPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundPool(ctx, req.(*MsgFundPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DefundPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDefundPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DefundPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.pool.v1beta1.Msg/DefundPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DefundPool(ctx, req.(*MsgDefundPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovMsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.pool.v1beta1.Msg/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*GovMsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovMsgUpdatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.pool.v1beta1.Msg/UpdatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePool(ctx, req.(*GovMsgUpdatePool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.pool.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FundPool",
			Handler:    _Msg_FundPool_Handler,
		},
		{
			MethodName: "DefundPool",
			Handler:    _Msg_DefundPool_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "UpdatePool",
			Handler:    _Msg_UpdatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kyve/pool/v1beta1/tx.proto",
}

func (m *MsgFundPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDefundPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDefundPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDefundPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDefundPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDefundPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDefundPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFundPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	return n
}

func (m *MsgFundPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDefundPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	return n
}

func (m *MsgDefundPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFundPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgFundPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDefundPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDefundPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDefundPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDefundPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDefundPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDefundPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
