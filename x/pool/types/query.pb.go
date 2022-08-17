// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/pool/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("kyve/pool/v1beta1/query.proto", fileDescriptor_9c2f559babbc8665) }

var fileDescriptor_9c2f559babbc8665 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x4a, 0xc7, 0x30,
	0x10, 0x87, 0xdb, 0x41, 0x85, 0x6e, 0x8a, 0x53, 0xd1, 0xec, 0x3a, 0xe4, 0xa8, 0xbe, 0x81, 0xe2,
	0x24, 0x08, 0x2e, 0x82, 0x6e, 0x97, 0x12, 0xd2, 0xd0, 0x36, 0x17, 0x93, 0xb4, 0xda, 0xb7, 0xf0,
	0xb1, 0x1c, 0x3b, 0x3a, 0x4a, 0xfb, 0x22, 0x92, 0x36, 0xb8, 0xfc, 0xb7, 0x83, 0xef, 0xe3, 0xbb,
	0x5f, 0x71, 0xd9, 0x4e, 0xa3, 0x04, 0x4b, 0xd4, 0xc1, 0x58, 0x09, 0x19, 0xb0, 0x82, 0xf7, 0x41,
	0xba, 0x89, 0x5b, 0x47, 0x81, 0xce, 0x4e, 0x23, 0xe6, 0x11, 0xf3, 0x84, 0xcb, 0x73, 0x45, 0x8a,
	0x36, 0x0a, 0xf1, 0xda, 0xc5, 0xf2, 0x42, 0x11, 0xa9, 0x4e, 0x02, 0x5a, 0x0d, 0x68, 0x0c, 0x05,
	0x0c, 0x9a, 0x8c, 0x4f, 0xf4, 0xba, 0x26, 0xdf, 0x93, 0x07, 0x81, 0x5e, 0xee, 0xfd, 0xff, 0x6f,
	0x16, 0x95, 0x36, 0x9b, 0x9c, 0x5c, 0x76, 0xb8, 0xc8, 0xa2, 0xc3, 0x3e, 0xb5, 0x6e, 0x4e, 0x8a,
	0xa3, 0xe7, 0x58, 0xb8, 0xbb, 0xff, 0x5e, 0x58, 0x3e, 0x2f, 0x2c, 0xff, 0x5d, 0x58, 0xfe, 0xb5,
	0xb2, 0x6c, 0x5e, 0x59, 0xf6, 0xb3, 0xb2, 0xec, 0xed, 0x4a, 0xe9, 0xd0, 0x0c, 0x82, 0xd7, 0xd4,
	0xc3, 0xe3, 0xeb, 0xcb, 0xc3, 0x93, 0x0c, 0x1f, 0xe4, 0x5a, 0xa8, 0x1b, 0xd4, 0x06, 0x3e, 0xf7,
	0x78, 0x98, 0xac, 0xf4, 0xe2, 0x78, 0x8b, 0xde, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb4,
	0xf9, 0xdb, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.pool.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "kyve/pool/v1beta1/query.proto",
}
