// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/tokenfactory/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{0}
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

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// Module parameters stored in state
	Params ModuleParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{1}
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

func (m *QueryParamsResponse) GetParams() ModuleParams {
	if m != nil {
		return m.Params
	}
	return ModuleParams{}
}

// QueryDenomsRequest: gRPC query for all denoms registered for a creator
type QueryDenomsRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *QueryDenomsRequest) Reset()         { *m = QueryDenomsRequest{} }
func (m *QueryDenomsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsRequest) ProtoMessage()    {}
func (*QueryDenomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{2}
}
func (m *QueryDenomsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsRequest.Merge(m, src)
}
func (m *QueryDenomsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsRequest proto.InternalMessageInfo

func (m *QueryDenomsRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// QueryDenomsResponse: All registered denoms for a creator
type QueryDenomsResponse struct {
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *QueryDenomsResponse) Reset()         { *m = QueryDenomsResponse{} }
func (m *QueryDenomsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsResponse) ProtoMessage()    {}
func (*QueryDenomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{3}
}
func (m *QueryDenomsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsResponse.Merge(m, src)
}
func (m *QueryDenomsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsResponse proto.InternalMessageInfo

func (m *QueryDenomsResponse) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

// QueryDenomInfoRequest: gRPC query for the denom admin and x/bank metadata
type QueryDenomInfoRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryDenomInfoRequest) Reset()         { *m = QueryDenomInfoRequest{} }
func (m *QueryDenomInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomInfoRequest) ProtoMessage()    {}
func (*QueryDenomInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{4}
}
func (m *QueryDenomInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomInfoRequest.Merge(m, src)
}
func (m *QueryDenomInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomInfoRequest proto.InternalMessageInfo

func (m *QueryDenomInfoRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryDenomInfoResponse: All registered denoms for a creator
type QueryDenomInfoResponse struct {
	// Admin of the token factory denom
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// Metadata: Official x/bank metadata for the denom. All token factory denoms
	// are standard, native assets.
	Metadata types.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
}

func (m *QueryDenomInfoResponse) Reset()         { *m = QueryDenomInfoResponse{} }
func (m *QueryDenomInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomInfoResponse) ProtoMessage()    {}
func (*QueryDenomInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d8bbc34d6c2a91, []int{5}
}
func (m *QueryDenomInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomInfoResponse.Merge(m, src)
}
func (m *QueryDenomInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomInfoResponse proto.InternalMessageInfo

func (m *QueryDenomInfoResponse) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *QueryDenomInfoResponse) GetMetadata() types.Metadata {
	if m != nil {
		return m.Metadata
	}
	return types.Metadata{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nibiru.tokenfactory.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nibiru.tokenfactory.v1.QueryParamsResponse")
	proto.RegisterType((*QueryDenomsRequest)(nil), "nibiru.tokenfactory.v1.QueryDenomsRequest")
	proto.RegisterType((*QueryDenomsResponse)(nil), "nibiru.tokenfactory.v1.QueryDenomsResponse")
	proto.RegisterType((*QueryDenomInfoRequest)(nil), "nibiru.tokenfactory.v1.QueryDenomInfoRequest")
	proto.RegisterType((*QueryDenomInfoResponse)(nil), "nibiru.tokenfactory.v1.QueryDenomInfoResponse")
}

func init() {
	proto.RegisterFile("nibiru/tokenfactory/v1/query.proto", fileDescriptor_b7d8bbc34d6c2a91)
}

var fileDescriptor_b7d8bbc34d6c2a91 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xbf, 0x6f, 0x13, 0x3f,
	0x18, 0xc6, 0xe3, 0xf6, 0x9b, 0xfb, 0x12, 0xb3, 0x99, 0x10, 0x45, 0x11, 0x1c, 0x91, 0x85, 0x50,
	0x14, 0x88, 0x4d, 0x82, 0x98, 0x91, 0x02, 0x0b, 0x43, 0xf9, 0x91, 0x0d, 0x36, 0x5f, 0xe2, 0x5c,
	0x4f, 0xed, 0xf9, 0xbd, 0x9e, 0x7d, 0x11, 0x51, 0xd5, 0x85, 0x8d, 0x0d, 0xa9, 0x3b, 0x7f, 0x4f,
	0xc5, 0x54, 0x89, 0x85, 0x09, 0xa1, 0x84, 0x3f, 0x04, 0xd5, 0xf6, 0x85, 0x06, 0x08, 0xe9, 0xe6,
	0xd7, 0xfe, 0xf8, 0x7d, 0x1e, 0xbf, 0xcf, 0x1d, 0xa6, 0x2a, 0x89, 0x92, 0xbc, 0xe0, 0x06, 0x0e,
	0xa4, 0x9a, 0x8a, 0xb1, 0x81, 0x7c, 0xce, 0x67, 0x7d, 0x7e, 0x54, 0xc8, 0x7c, 0xce, 0xb2, 0x1c,
	0x0c, 0x90, 0x86, 0x63, 0xd8, 0x65, 0x86, 0xcd, 0xfa, 0xad, 0x7a, 0x0c, 0x31, 0x58, 0x84, 0x5f,
	0xac, 0x1c, 0xdd, 0xba, 0x15, 0x03, 0xc4, 0x87, 0x92, 0x8b, 0x2c, 0xe1, 0x42, 0x29, 0x30, 0xc2,
	0x24, 0xa0, 0xb4, 0x3f, 0x0d, 0xc7, 0xa0, 0x53, 0xd0, 0x3c, 0x12, 0xea, 0x80, 0xcf, 0xfa, 0x91,
	0x34, 0xa2, 0x6f, 0x0b, 0x7f, 0xbe, 0xc9, 0x8f, 0x36, 0xc2, 0x48, 0xc7, 0xd0, 0x3a, 0x26, 0xaf,
	0x2f, 0xec, 0xbd, 0x12, 0xb9, 0x48, 0xf5, 0x48, 0x1e, 0x15, 0x52, 0x1b, 0xfa, 0x06, 0xdf, 0x58,
	0xdb, 0xd5, 0x19, 0x28, 0x2d, 0xc9, 0x10, 0x07, 0x99, 0xdd, 0x69, 0xa2, 0x36, 0xea, 0x5c, 0x1f,
	0xdc, 0x65, 0x7f, 0x7f, 0x0d, 0xdb, 0x83, 0x49, 0x71, 0x28, 0xdd, 0xed, 0xe1, 0x7f, 0x67, 0xdf,
	0xee, 0x54, 0x46, 0xfe, 0x26, 0x65, 0x5e, 0xf0, 0x99, 0x54, 0xb0, 0x12, 0x24, 0x4d, 0xfc, 0xff,
	0x38, 0x97, 0xc2, 0x40, 0x6e, 0x5b, 0xd7, 0x46, 0x65, 0x49, 0x7b, 0xde, 0x4a, 0xc9, 0x7b, 0x2b,
	0x0d, 0x1c, 0x4c, 0xec, 0x4e, 0x13, 0xb5, 0x77, 0x3b, 0xb5, 0x91, 0xaf, 0x68, 0x0f, 0xdf, 0xfc,
	0x85, 0x3f, 0x57, 0x53, 0x28, 0x15, 0xea, 0xb8, 0x6a, 0x11, 0xdf, 0xdf, 0x15, 0x14, 0x70, 0xe3,
	0x77, 0xdc, 0x0b, 0xd4, 0x71, 0x55, 0x4c, 0xd2, 0x44, 0x95, 0xbc, 0x2d, 0xc8, 0x13, 0x7c, 0x2d,
	0x95, 0x46, 0x4c, 0x84, 0x11, 0xcd, 0x1d, 0x3b, 0x83, 0xdb, 0xcc, 0xa5, 0xc0, 0xec, 0xe0, 0x7d,
	0x0a, 0x6c, 0xcf, 0x43, 0xfe, 0xf1, 0xab, 0x4b, 0x83, 0xcf, 0xbb, 0xb8, 0x6a, 0x15, 0xc9, 0x07,
	0x84, 0x03, 0x37, 0x21, 0xd2, 0xdd, 0x34, 0xc7, 0x3f, 0xa3, 0x69, 0xdd, 0xbf, 0x12, 0xeb, 0x1e,
	0x41, 0xef, 0xbd, 0xff, 0xf2, 0xe3, 0x74, 0xa7, 0x4d, 0x42, 0xbe, 0xe1, 0x53, 0x70, 0xa1, 0x90,
	0x53, 0x84, 0x03, 0x37, 0xe0, 0x2d, 0x5e, 0xd6, 0x52, 0xdb, 0xe2, 0x65, 0x3d, 0x31, 0xfa, 0xd0,
	0x7a, 0xe9, 0x92, 0xce, 0x26, 0x2f, 0x2e, 0x41, 0x7e, 0xec, 0x93, 0x3f, 0x21, 0x9f, 0x10, 0xae,
	0xad, 0x82, 0x21, 0xbd, 0xed, 0x62, 0x97, 0xf2, 0x6e, 0xb1, 0xab, 0xe2, 0xde, 0xde, 0xc0, 0xda,
	0x7b, 0x40, 0xba, 0xff, 0xb4, 0xd7, 0x4b, 0xd4, 0x14, 0xf8, 0xb1, 0x5d, 0x9f, 0x0c, 0x5f, 0x9e,
	0x2d, 0x42, 0x74, 0xbe, 0x08, 0xd1, 0xf7, 0x45, 0x88, 0x3e, 0x2e, 0xc3, 0xca, 0xf9, 0x32, 0xac,
	0x7c, 0x5d, 0x86, 0x95, 0xb7, 0x8f, 0xe3, 0xc4, 0xec, 0x17, 0x11, 0x1b, 0x43, 0xca, 0x5f, 0xd8,
	0x7e, 0x4f, 0xf7, 0x45, 0xa2, 0xca, 0xde, 0xb3, 0x01, 0x7f, 0xb7, 0x2e, 0x60, 0xe6, 0x99, 0xd4,
	0x51, 0x60, 0x7f, 0xca, 0x47, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x55, 0x22, 0x33, 0xf5, 0x4a,
	0x04, 0x00, 0x00,
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
	// Params retrieves the module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Denoms retrieves all registered denoms for a given creator
	Denoms(ctx context.Context, in *QueryDenomsRequest, opts ...grpc.CallOption) (*QueryDenomsResponse, error)
	// DenomInfo retrieves the denom metadata and admin info
	DenomInfo(ctx context.Context, in *QueryDenomInfoRequest, opts ...grpc.CallOption) (*QueryDenomInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nibiru.tokenfactory.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denoms(ctx context.Context, in *QueryDenomsRequest, opts ...grpc.CallOption) (*QueryDenomsResponse, error) {
	out := new(QueryDenomsResponse)
	err := c.cc.Invoke(ctx, "/nibiru.tokenfactory.v1.Query/Denoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomInfo(ctx context.Context, in *QueryDenomInfoRequest, opts ...grpc.CallOption) (*QueryDenomInfoResponse, error) {
	out := new(QueryDenomInfoResponse)
	err := c.cc.Invoke(ctx, "/nibiru.tokenfactory.v1.Query/DenomInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params retrieves the module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Denoms retrieves all registered denoms for a given creator
	Denoms(context.Context, *QueryDenomsRequest) (*QueryDenomsResponse, error)
	// DenomInfo retrieves the denom metadata and admin info
	DenomInfo(context.Context, *QueryDenomInfoRequest) (*QueryDenomInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Denoms(ctx context.Context, req *QueryDenomsRequest) (*QueryDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denoms not implemented")
}
func (*UnimplementedQueryServer) DenomInfo(ctx context.Context, req *QueryDenomInfoRequest) (*QueryDenomInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomInfo not implemented")
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
		FullMethod: "/nibiru.tokenfactory.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nibiru.tokenfactory.v1.Query/Denoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denoms(ctx, req.(*QueryDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nibiru.tokenfactory.v1.Query/DenomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomInfo(ctx, req.(*QueryDenomInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nibiru.tokenfactory.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Denoms",
			Handler:    _Query_Denoms_Handler,
		},
		{
			MethodName: "DenomInfo",
			Handler:    _Query_DenomInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nibiru/tokenfactory/v1/query.proto",
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

func (m *QueryDenomsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
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

func (m *QueryDenomsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDenomInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovQuery(uint64(l))
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
func (m *QueryDenomsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDenomsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
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
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
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
func (m *QueryDenomInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDenomInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
