// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smplsecretsvaultchain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetWalletsRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *QueryGetWalletsRequest) Reset()         { *m = QueryGetWalletsRequest{} }
func (m *QueryGetWalletsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetWalletsRequest) ProtoMessage()    {}
func (*QueryGetWalletsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{2}
}
func (m *QueryGetWalletsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWalletsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWalletsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWalletsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWalletsRequest.Merge(m, src)
}
func (m *QueryGetWalletsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWalletsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWalletsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWalletsRequest proto.InternalMessageInfo

func (m *QueryGetWalletsRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type QueryGetWalletsResponse struct {
	Wallets Wallets `protobuf:"bytes,1,opt,name=wallets,proto3" json:"wallets"`
}

func (m *QueryGetWalletsResponse) Reset()         { *m = QueryGetWalletsResponse{} }
func (m *QueryGetWalletsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetWalletsResponse) ProtoMessage()    {}
func (*QueryGetWalletsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{3}
}
func (m *QueryGetWalletsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWalletsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWalletsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWalletsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWalletsResponse.Merge(m, src)
}
func (m *QueryGetWalletsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWalletsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWalletsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWalletsResponse proto.InternalMessageInfo

func (m *QueryGetWalletsResponse) GetWallets() Wallets {
	if m != nil {
		return m.Wallets
	}
	return Wallets{}
}

type QueryAllWalletsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWalletsRequest) Reset()         { *m = QueryAllWalletsRequest{} }
func (m *QueryAllWalletsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllWalletsRequest) ProtoMessage()    {}
func (*QueryAllWalletsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{4}
}
func (m *QueryAllWalletsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWalletsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWalletsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWalletsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWalletsRequest.Merge(m, src)
}
func (m *QueryAllWalletsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWalletsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWalletsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWalletsRequest proto.InternalMessageInfo

func (m *QueryAllWalletsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllWalletsResponse struct {
	Wallets    []Wallets           `protobuf:"bytes,1,rep,name=wallets,proto3" json:"wallets"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWalletsResponse) Reset()         { *m = QueryAllWalletsResponse{} }
func (m *QueryAllWalletsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllWalletsResponse) ProtoMessage()    {}
func (*QueryAllWalletsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc0433211492f76, []int{5}
}
func (m *QueryAllWalletsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWalletsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWalletsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWalletsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWalletsResponse.Merge(m, src)
}
func (m *QueryAllWalletsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWalletsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWalletsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWalletsResponse proto.InternalMessageInfo

func (m *QueryAllWalletsResponse) GetWallets() []Wallets {
	if m != nil {
		return m.Wallets
	}
	return nil
}

func (m *QueryAllWalletsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryParamsResponse")
	proto.RegisterType((*QueryGetWalletsRequest)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryGetWalletsRequest")
	proto.RegisterType((*QueryGetWalletsResponse)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryGetWalletsResponse")
	proto.RegisterType((*QueryAllWalletsRequest)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryAllWalletsRequest")
	proto.RegisterType((*QueryAllWalletsResponse)(nil), "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.QueryAllWalletsResponse")
}

func init() { proto.RegisterFile("smplsecretsvaultchain/query.proto", fileDescriptor_5dc0433211492f76) }

var fileDescriptor_5dc0433211492f76 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xab, 0xdd, 0xe2, 0x78, 0x1b, 0x17, 0x95, 0x45, 0xa2, 0x8e, 0xa0, 0xe2, 0x61,
	0x86, 0xd6, 0x93, 0xc7, 0xdd, 0xaa, 0xc5, 0x52, 0xa5, 0x4d, 0x41, 0xc1, 0x8b, 0x4e, 0xc2, 0x98,
	0x06, 0x26, 0x99, 0x34, 0x33, 0xbb, 0x75, 0x91, 0x5e, 0x44, 0xef, 0x82, 0xdf, 0x49, 0x7a, 0x2c,
	0x78, 0xf1, 0x24, 0xb2, 0xeb, 0x27, 0xf0, 0x0b, 0x28, 0x99, 0x79, 0xc5, 0xdd, 0x26, 0x2a, 0x98,
	0xed, 0x6d, 0x93, 0x7d, 0xf3, 0xff, 0xff, 0xfe, 0xef, 0xcd, 0x0b, 0xba, 0xae, 0xd3, 0x5c, 0x6a,
	0x11, 0x15, 0xc2, 0xe8, 0x11, 0x1f, 0x4a, 0x13, 0xed, 0xf2, 0x24, 0x63, 0x7b, 0x43, 0x51, 0x8c,
	0x69, 0x5e, 0x28, 0xa3, 0xf0, 0xbd, 0x9d, 0x34, 0x97, 0x0f, 0x22, 0xa5, 0xc7, 0xda, 0x88, 0x94,
	0xd6, 0x1e, 0xa8, 0x7f, 0xdb, 0xeb, 0xc6, 0x2a, 0x56, 0x56, 0x85, 0x95, 0xbf, 0x9c, 0x60, 0xef,
	0x4a, 0xac, 0x54, 0x2c, 0x05, 0xe3, 0x79, 0xc2, 0x78, 0x96, 0x29, 0xc3, 0x4d, 0xa2, 0x32, 0x0d,
	0xff, 0xde, 0x89, 0x94, 0x4e, 0x95, 0x66, 0x21, 0xd7, 0xc2, 0x71, 0xb0, 0xd1, 0x4a, 0x28, 0x0c,
	0x5f, 0x61, 0x39, 0x8f, 0x93, 0xcc, 0x16, 0x43, 0x2d, 0xa9, 0xa7, 0xcf, 0x79, 0xc1, 0xd3, 0x63,
	0xbd, 0x1b, 0xf5, 0x35, 0xfb, 0x5c, 0x4a, 0x61, 0xa0, 0x88, 0x74, 0x11, 0xde, 0x2e, 0xad, 0xb6,
	0xec, 0xc9, 0x40, 0xec, 0x0d, 0x85, 0x36, 0x64, 0x84, 0x2e, 0xcc, 0xbd, 0xd5, 0xb9, 0xca, 0xb4,
	0xc0, 0x2f, 0x50, 0xc7, 0x39, 0x5c, 0xf6, 0xae, 0x79, 0xb7, 0xcf, 0xaf, 0xf6, 0xe9, 0x7f, 0x77,
	0x88, 0x3a, 0xe9, 0xc1, 0xd9, 0xc3, 0xaf, 0x57, 0x5b, 0x01, 0xc8, 0x12, 0x8a, 0x2e, 0x5a, 0xdf,
	0x75, 0x61, 0x9e, 0x39, 0x4c, 0x20, 0xc2, 0x5d, 0xb4, 0xa4, 0xf6, 0x33, 0x51, 0x58, 0xe7, 0x73,
	0x81, 0x7b, 0x20, 0x07, 0xe8, 0x52, 0xa5, 0x1e, 0x58, 0x43, 0xb4, 0x0c, 0x49, 0x01, 0x76, 0xd0,
	0x00, 0x16, 0xc4, 0x81, 0xf6, 0x58, 0x98, 0xbc, 0x04, 0xdc, 0xbe, 0x94, 0x27, 0x70, 0x1f, 0x22,
	0xf4, 0x7b, 0x66, 0x00, 0x70, 0x93, 0xba, 0x01, 0xd3, 0x72, 0xc0, 0xd4, 0x5d, 0x34, 0x18, 0x30,
	0xdd, 0xe2, 0xb1, 0x80, 0xb3, 0xc1, 0xcc, 0x49, 0xf2, 0xc9, 0x83, 0x84, 0xb3, 0x16, 0x75, 0x09,
	0xcf, 0x9c, 0x4a, 0x42, 0xbc, 0x3e, 0x97, 0xa3, 0x6d, 0x73, 0xdc, 0xfa, 0x67, 0x0e, 0x07, 0x38,
	0x1b, 0x64, 0xf5, 0x7d, 0x07, 0x2d, 0xd9, 0x20, 0xf8, 0x87, 0x87, 0x3a, 0x6e, 0xf8, 0xf8, 0x71,
	0x03, 0xe0, 0xea, 0xad, 0xed, 0x3d, 0x59, 0x94, 0x9c, 0xe3, 0x27, 0x8f, 0xde, 0x7e, 0xfe, 0xfe,
	0xb1, 0xbd, 0x86, 0xfb, 0x6c, 0x4e, 0x97, 0xd5, 0xef, 0xd5, 0xdf, 0x36, 0x12, 0xbf, 0x6b, 0xa3,
	0x65, 0x68, 0x31, 0xde, 0x6e, 0x8a, 0x59, 0xd9, 0x8e, 0x5e, 0xb0, 0x48, 0x49, 0x48, 0x1f, 0xd8,
	0xf4, 0x9b, 0x78, 0xa3, 0x41, 0x7a, 0xb8, 0x46, 0xec, 0x8d, 0x5d, 0xd7, 0x03, 0xfc, 0xd3, 0x43,
	0x08, 0x7c, 0xfa, 0x52, 0x36, 0xef, 0x44, 0x65, 0xf1, 0x9a, 0x77, 0xa2, 0xba, 0x68, 0x64, 0xc3,
	0x76, 0xe2, 0x3e, 0x1e, 0x34, 0xef, 0xc4, 0xe0, 0xd5, 0xe1, 0xc4, 0xf7, 0x8e, 0x26, 0xbe, 0xf7,
	0x6d, 0xe2, 0x7b, 0x1f, 0xa6, 0x7e, 0xeb, 0x68, 0xea, 0xb7, 0xbe, 0x4c, 0xfd, 0xd6, 0xf3, 0xcd,
	0x38, 0x31, 0xbb, 0xc3, 0x90, 0x46, 0x2a, 0x3d, 0xe1, 0x53, 0x3e, 0xed, 0x38, 0xc5, 0xa7, 0xa5,
	0xe2, 0x9a, 0x55, 0x7c, 0xfd, 0x07, 0x27, 0x33, 0xce, 0x85, 0x0e, 0x3b, 0xf6, 0xf3, 0x7e, 0xf7,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x6c, 0x7e, 0xe1, 0xe7, 0x06, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Wallets by index.
	Wallets(ctx context.Context, in *QueryGetWalletsRequest, opts ...grpc.CallOption) (*QueryGetWalletsResponse, error)
	// Queries a list of Wallets items.
	WalletsAll(ctx context.Context, in *QueryAllWalletsRequest, opts ...grpc.CallOption) (*QueryAllWalletsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Wallets(ctx context.Context, in *QueryGetWalletsRequest, opts ...grpc.CallOption) (*QueryGetWalletsResponse, error) {
	out := new(QueryGetWalletsResponse)
	err := c.cc.Invoke(ctx, "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/Wallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WalletsAll(ctx context.Context, in *QueryAllWalletsRequest, opts ...grpc.CallOption) (*QueryAllWalletsResponse, error) {
	out := new(QueryAllWalletsResponse)
	err := c.cc.Invoke(ctx, "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/WalletsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Wallets by index.
	Wallets(context.Context, *QueryGetWalletsRequest) (*QueryGetWalletsResponse, error)
	// Queries a list of Wallets items.
	WalletsAll(context.Context, *QueryAllWalletsRequest) (*QueryAllWalletsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Wallets(ctx context.Context, req *QueryGetWalletsRequest) (*QueryGetWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wallets not implemented")
}
func (*UnimplementedQueryServer) WalletsAll(ctx context.Context, req *QueryAllWalletsRequest) (*QueryAllWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletsAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Wallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Wallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/Wallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Wallets(ctx, req.(*QueryGetWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WalletsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WalletsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query/WalletsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WalletsAll(ctx, req.(*QueryAllWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SmplEcosystem.smplsecretsvaultchain.smplsecretsvaultchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Wallets",
			Handler:    _Query_Wallets_Handler,
		},
		{
			MethodName: "WalletsAll",
			Handler:    _Query_WalletsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smplsecretsvaultchain/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetWalletsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWalletsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWalletsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetWalletsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWalletsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWalletsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Wallets.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllWalletsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWalletsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWalletsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllWalletsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWalletsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWalletsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Wallets) > 0 {
		for iNdEx := len(m.Wallets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Wallets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetWalletsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetWalletsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Wallets.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllWalletsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllWalletsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Wallets) > 0 {
		for _, e := range m.Wallets {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetWalletsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetWalletsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWalletsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetWalletsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetWalletsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWalletsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Wallets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllWalletsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllWalletsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWalletsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllWalletsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllWalletsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWalletsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallets = append(m.Wallets, Wallets{})
			if err := m.Wallets[len(m.Wallets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
