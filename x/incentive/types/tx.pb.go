// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/incentive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgWithdrawReward defines a message for withdrawing reward of a stakeholder.
type MsgWithdrawReward struct {
	// {submitter, reporter, finality_provider, btc_delegation}
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// address is the address of the stakeholder in bech32 string
	// signer of this msg has to be this address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgWithdrawReward) Reset()         { *m = MsgWithdrawReward{} }
func (m *MsgWithdrawReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawReward) ProtoMessage()    {}
func (*MsgWithdrawReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4de6776d39a3a22, []int{0}
}
func (m *MsgWithdrawReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawReward.Merge(m, src)
}
func (m *MsgWithdrawReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawReward proto.InternalMessageInfo

func (m *MsgWithdrawReward) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MsgWithdrawReward) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// MsgWithdrawRewardResponse is the response to the MsgWithdrawReward message
type MsgWithdrawRewardResponse struct {
	// coins is the withdrawed coins
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgWithdrawRewardResponse) Reset()         { *m = MsgWithdrawRewardResponse{} }
func (m *MsgWithdrawRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewardResponse) ProtoMessage()    {}
func (*MsgWithdrawRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4de6776d39a3a22, []int{1}
}
func (m *MsgWithdrawRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewardResponse.Merge(m, src)
}
func (m *MsgWithdrawRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewardResponse proto.InternalMessageInfo

func (m *MsgWithdrawRewardResponse) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// MsgUpdateParams defines a message for updating incentive module parameters.
type MsgUpdateParams struct {
	// authority is the address of the governance account.
	// just FYI: cosmos.AddressString marks that this field should use type alias
	// for AddressString instead of string, but the functionality is not yet implemented
	// in cosmos-proto
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the incentive parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4de6776d39a3a22, []int{2}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse is the response to the MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4de6776d39a3a22, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgWithdrawReward)(nil), "babylon.incentive.MsgWithdrawReward")
	proto.RegisterType((*MsgWithdrawRewardResponse)(nil), "babylon.incentive.MsgWithdrawRewardResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "babylon.incentive.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "babylon.incentive.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("babylon/incentive/tx.proto", fileDescriptor_b4de6776d39a3a22) }

var fileDescriptor_b4de6776d39a3a22 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0xd8, 0x1f, 0xd2, 0xb1, 0x54, 0x3a, 0x14, 0xba, 0xd9, 0xc3, 0xb6, 0x2c, 0x1e, 0x42,
	0x30, 0x3b, 0xa6, 0x05, 0x85, 0xde, 0x8c, 0xe7, 0xa0, 0xac, 0x88, 0xe0, 0x41, 0x99, 0xcd, 0x0e,
	0x93, 0xc1, 0xee, 0xce, 0xb2, 0x6f, 0x9a, 0x36, 0x17, 0x11, 0xff, 0x02, 0xf1, 0xcf, 0xf0, 0xd4,
	0x83, 0x7f, 0x44, 0x8f, 0xa5, 0x27, 0x4f, 0x2a, 0xc9, 0xa1, 0xff, 0x86, 0xcc, 0xce, 0x6c, 0x1b,
	0xbb, 0x05, 0x7b, 0x9a, 0x79, 0xf9, 0xbe, 0xf9, 0xde, 0xf7, 0xbe, 0x97, 0xc5, 0x7e, 0xc2, 0x92,
	0xe9, 0xa1, 0xca, 0xa9, 0xcc, 0x47, 0x3c, 0xd7, 0x72, 0xc2, 0xa9, 0x3e, 0x89, 0x8a, 0x52, 0x69,
	0x45, 0x36, 0x1d, 0x16, 0x5d, 0x61, 0xfe, 0x96, 0x50, 0x42, 0x55, 0x28, 0x35, 0x37, 0x4b, 0xf4,
	0xdb, 0x23, 0x05, 0x99, 0x82, 0x0f, 0x16, 0xb0, 0x85, 0x83, 0xb6, 0x6d, 0x45, 0x33, 0x10, 0x74,
	0xd2, 0x37, 0x87, 0x03, 0x02, 0x07, 0x24, 0x0c, 0x38, 0x9d, 0xf4, 0x13, 0xae, 0x59, 0x9f, 0x8e,
	0x94, 0xcc, 0x6b, 0xbc, 0x69, 0xac, 0x60, 0x25, 0xcb, 0x9c, 0x70, 0xf8, 0x12, 0x6f, 0x0e, 0x41,
	0xbc, 0x95, 0x7a, 0x9c, 0x96, 0xec, 0x38, 0xe6, 0xc7, 0xac, 0x4c, 0x09, 0xc1, 0xcb, 0x7a, 0x5a,
	0x70, 0x0f, 0xed, 0xa2, 0xce, 0x5a, 0x5c, 0xdd, 0x89, 0x87, 0xef, 0xb3, 0x34, 0x2d, 0x39, 0x80,
	0x77, 0xaf, 0xfa, 0xb9, 0x2e, 0x0f, 0xd6, 0xbf, 0x5c, 0x9e, 0x76, 0xeb, 0x2a, 0xfc, 0x84, 0xdb,
	0x0d, 0xc1, 0x98, 0x43, 0xa1, 0x72, 0xe0, 0x84, 0xe1, 0x15, 0xe3, 0x0d, 0x3c, 0xb4, 0xbb, 0xd4,
	0x79, 0xb0, 0xd7, 0x8e, 0xdc, 0x90, 0xc6, 0x7d, 0xe4, 0xdc, 0x47, 0x2f, 0x94, 0xcc, 0x07, 0x4f,
	0xce, 0x7e, 0xed, 0xb4, 0xbe, 0xff, 0xde, 0xe9, 0x08, 0xa9, 0xc7, 0x47, 0x49, 0x34, 0x52, 0x99,
	0x4b, 0xc4, 0x1d, 0x3d, 0x48, 0x3f, 0x52, 0xe3, 0x0c, 0xaa, 0x07, 0x10, 0x5b, 0xe5, 0xf0, 0x1b,
	0xc2, 0x0f, 0x87, 0x20, 0xde, 0x14, 0x29, 0xd3, 0xfc, 0x55, 0x35, 0x2a, 0x79, 0x8a, 0xd7, 0xd8,
	0x91, 0x1e, 0xab, 0x52, 0xea, 0xa9, 0x1d, 0x6a, 0xe0, 0x5d, 0xfc, 0xe8, 0x6d, 0xb9, 0xee, 0xcf,
	0xad, 0xf5, 0xd7, 0xba, 0x94, 0xb9, 0x88, 0xaf, 0xa9, 0xe4, 0x19, 0x5e, 0xb5, 0x61, 0x55, 0x23,
	0x1b, 0xbf, 0x8d, 0x55, 0x46, 0xb6, 0xc5, 0x60, 0xd9, 0xf8, 0x8d, 0x1d, 0xfd, 0x60, 0xc3, 0x44,
	0x72, 0x2d, 0x14, 0xb6, 0xf1, 0xf6, 0x0d, 0x4f, 0x75, 0x24, 0x7b, 0x17, 0x08, 0x2f, 0x0d, 0x41,
	0x90, 0x14, 0x6f, 0xdc, 0xd8, 0xc2, 0xa3, 0x5b, 0xba, 0x35, 0xa2, 0xf5, 0x1f, 0xdf, 0x85, 0x75,
	0xb5, 0x80, 0xf7, 0x78, 0xfd, 0x9f, 0x64, 0xc2, 0xdb, 0x5f, 0x2f, 0x72, 0xfc, 0xee, 0xff, 0x39,
	0xb5, 0xbe, 0xbf, 0xf2, 0xf9, 0xf2, 0xb4, 0x8b, 0x06, 0xc3, 0xb3, 0x59, 0x80, 0xce, 0x67, 0x01,
	0xfa, 0x33, 0x0b, 0xd0, 0xd7, 0x79, 0xd0, 0x3a, 0x9f, 0x07, 0xad, 0x9f, 0xf3, 0xa0, 0xf5, 0x6e,
	0x7f, 0x61, 0x9f, 0x4e, 0xf6, 0x90, 0x25, 0xd0, 0x93, 0xaa, 0x2e, 0xe9, 0xc9, 0xe2, 0x47, 0x64,
	0x16, 0x9c, 0xac, 0x56, 0xff, 0xd5, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x43, 0xa9,
	0xe8, 0x66, 0x03, 0x00, 0x00,
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
	// WithdrawReward defines a method to withdraw rewards of a stakeholder
	WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error)
	// UpdateParams updates the incentive module parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error) {
	out := new(MsgWithdrawRewardResponse)
	err := c.cc.Invoke(ctx, "/babylon.incentive.Msg/WithdrawReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/babylon.incentive.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// WithdrawReward defines a method to withdraw rewards of a stakeholder
	WithdrawReward(context.Context, *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error)
	// UpdateParams updates the incentive module parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) WithdrawReward(ctx context.Context, req *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawReward not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_WithdrawReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.incentive.Msg/WithdrawReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawReward(ctx, req.(*MsgWithdrawReward))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.incentive.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.incentive.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawReward",
			Handler:    _Msg_WithdrawReward_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/incentive/tx.proto",
}

func (m *MsgWithdrawReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgWithdrawReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
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
func (m *MsgWithdrawReward) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgWithdrawRewardResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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