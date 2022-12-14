// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smaller/tx.proto

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

type MsgPlayGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	NumPick string `protobuf:"bytes,2,opt,name=numPick,proto3" json:"numPick,omitempty"`
}

func (m *MsgPlayGame) Reset()         { *m = MsgPlayGame{} }
func (m *MsgPlayGame) String() string { return proto.CompactTextString(m) }
func (*MsgPlayGame) ProtoMessage()    {}
func (*MsgPlayGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_81e131a7addccced, []int{0}
}
func (m *MsgPlayGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlayGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlayGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlayGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlayGame.Merge(m, src)
}
func (m *MsgPlayGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlayGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlayGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlayGame proto.InternalMessageInfo

func (m *MsgPlayGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPlayGame) GetNumPick() string {
	if m != nil {
		return m.NumPick
	}
	return ""
}

type MsgPlayGameResponse struct {
	GameIndex string `protobuf:"bytes,1,opt,name=gameIndex,proto3" json:"gameIndex,omitempty"`
	Win       bool   `protobuf:"varint,2,opt,name=win,proto3" json:"win,omitempty"`
}

func (m *MsgPlayGameResponse) Reset()         { *m = MsgPlayGameResponse{} }
func (m *MsgPlayGameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlayGameResponse) ProtoMessage()    {}
func (*MsgPlayGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81e131a7addccced, []int{1}
}
func (m *MsgPlayGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlayGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlayGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlayGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlayGameResponse.Merge(m, src)
}
func (m *MsgPlayGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlayGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlayGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlayGameResponse proto.InternalMessageInfo

func (m *MsgPlayGameResponse) GetGameIndex() string {
	if m != nil {
		return m.GameIndex
	}
	return ""
}

func (m *MsgPlayGameResponse) GetWin() bool {
	if m != nil {
		return m.Win
	}
	return false
}

type MsgSendGameResult struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	GameId           uint64 `protobuf:"varint,5,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (m *MsgSendGameResult) Reset()         { *m = MsgSendGameResult{} }
func (m *MsgSendGameResult) String() string { return proto.CompactTextString(m) }
func (*MsgSendGameResult) ProtoMessage()    {}
func (*MsgSendGameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_81e131a7addccced, []int{2}
}
func (m *MsgSendGameResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendGameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendGameResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendGameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendGameResult.Merge(m, src)
}
func (m *MsgSendGameResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendGameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendGameResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendGameResult proto.InternalMessageInfo

func (m *MsgSendGameResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendGameResult) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendGameResult) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendGameResult) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendGameResult) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type MsgSendGameResultResponse struct {
}

func (m *MsgSendGameResultResponse) Reset()         { *m = MsgSendGameResultResponse{} }
func (m *MsgSendGameResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendGameResultResponse) ProtoMessage()    {}
func (*MsgSendGameResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81e131a7addccced, []int{3}
}
func (m *MsgSendGameResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendGameResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendGameResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendGameResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendGameResultResponse.Merge(m, src)
}
func (m *MsgSendGameResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendGameResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendGameResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendGameResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPlayGame)(nil), "tmsdkeys.smaller3.smaller.MsgPlayGame")
	proto.RegisterType((*MsgPlayGameResponse)(nil), "tmsdkeys.smaller3.smaller.MsgPlayGameResponse")
	proto.RegisterType((*MsgSendGameResult)(nil), "tmsdkeys.smaller3.smaller.MsgSendGameResult")
	proto.RegisterType((*MsgSendGameResultResponse)(nil), "tmsdkeys.smaller3.smaller.MsgSendGameResultResponse")
}

func init() { proto.RegisterFile("smaller/tx.proto", fileDescriptor_81e131a7addccced) }

var fileDescriptor_81e131a7addccced = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x3b, 0xff, 0xf6, 0x5f, 0xdb, 0x2b, 0x48, 0x1d, 0x41, 0xd2, 0x2a, 0xa1, 0x64, 0x21,
	0x45, 0x4a, 0x02, 0xd6, 0x17, 0x50, 0xfc, 0xa0, 0x8b, 0x42, 0x89, 0xae, 0xdc, 0x4d, 0x93, 0x21,
	0x0d, 0xcd, 0xcc, 0x84, 0xcc, 0x04, 0xdb, 0xb7, 0xf0, 0x1d, 0x7c, 0x19, 0x97, 0x5d, 0xba, 0x53,
	0xda, 0x17, 0x91, 0xa4, 0x49, 0xac, 0x16, 0x8b, 0xee, 0xee, 0xc7, 0xe1, 0x77, 0xcf, 0x1c, 0x06,
	0x1a, 0x92, 0x91, 0x20, 0xa0, 0x91, 0xa5, 0xa6, 0x66, 0x18, 0x09, 0x25, 0x70, 0x53, 0x31, 0xe9,
	0x4e, 0xe8, 0x4c, 0x9a, 0xd9, 0xaa, 0x97, 0x17, 0xc6, 0x05, 0xec, 0x0e, 0xa4, 0x37, 0x0c, 0xc8,
	0xec, 0x96, 0x30, 0x8a, 0x35, 0xd8, 0x71, 0x22, 0x4a, 0x94, 0x88, 0x34, 0xd4, 0x46, 0x9d, 0xba,
	0x9d, 0xb7, 0xc9, 0x86, 0xc7, 0x6c, 0xe8, 0x3b, 0x13, 0xed, 0xdf, 0x6a, 0x93, 0xb5, 0xc6, 0x35,
	0x1c, 0xac, 0x21, 0x6c, 0x2a, 0x43, 0xc1, 0x25, 0xc5, 0xc7, 0x50, 0xf7, 0x08, 0xa3, 0x7d, 0xee,
	0xd2, 0x69, 0x06, 0xfb, 0x1c, 0xe0, 0x06, 0x94, 0x1f, 0x7d, 0x9e, 0xa2, 0x6a, 0x76, 0x52, 0x1a,
	0xcf, 0x08, 0xf6, 0x07, 0xd2, 0xbb, 0xa3, 0xdc, 0xcd, 0x38, 0x71, 0xa0, 0xb6, 0x18, 0xc2, 0x50,
	0x09, 0x45, 0xa4, 0x32, 0x37, 0x69, 0x9d, 0xdc, 0x74, 0xc6, 0x84, 0x73, 0x1a, 0xf4, 0xaf, 0xb4,
	0xf2, 0xea, 0x66, 0x31, 0xc0, 0xa7, 0xd0, 0x50, 0x3e, 0xa3, 0x22, 0x56, 0xf7, 0x3e, 0xa3, 0x52,
	0x11, 0x16, 0x6a, 0x95, 0x36, 0xea, 0x54, 0xec, 0x8d, 0x39, 0x3e, 0x84, 0x6a, 0x6a, 0xd6, 0xd5,
	0xfe, 0xa7, 0x8a, 0xac, 0x33, 0x8e, 0xa0, 0xb9, 0x61, 0x32, 0x7f, 0xf2, 0xd9, 0x1b, 0x82, 0xf2,
	0x40, 0x7a, 0x78, 0x04, 0xb5, 0x22, 0xd1, 0x13, 0xf3, 0xc7, 0xf0, 0xcd, 0xb5, 0xd8, 0x5a, 0xe6,
	0xef, 0x74, 0x45, 0xbc, 0x0a, 0xf6, 0xbe, 0x45, 0xd5, 0xdd, 0x4e, 0xf8, 0xaa, 0x6e, 0x9d, 0xff,
	0x45, 0x9d, 0x5f, 0xbd, 0xbc, 0x79, 0x59, 0xe8, 0x68, 0xbe, 0xd0, 0xd1, 0xfb, 0x42, 0x47, 0x4f,
	0x4b, 0xbd, 0x34, 0x5f, 0xea, 0xa5, 0xd7, 0xa5, 0x5e, 0x7a, 0xe8, 0x7a, 0xbe, 0x1a, 0xc7, 0x23,
	0xd3, 0x11, 0xcc, 0xca, 0xc9, 0x56, 0x4e, 0xb6, 0xa6, 0x56, 0xf1, 0x29, 0x67, 0x21, 0x95, 0xa3,
	0x6a, 0xfa, 0x31, 0x7b, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xf9, 0xa8, 0x81, 0xac, 0x02,
	0x00, 0x00,
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
	PlayGame(ctx context.Context, in *MsgPlayGame, opts ...grpc.CallOption) (*MsgPlayGameResponse, error)
	SendGameResult(ctx context.Context, in *MsgSendGameResult, opts ...grpc.CallOption) (*MsgSendGameResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlayGame(ctx context.Context, in *MsgPlayGame, opts ...grpc.CallOption) (*MsgPlayGameResponse, error) {
	out := new(MsgPlayGameResponse)
	err := c.cc.Invoke(ctx, "/tmsdkeys.smaller3.smaller.Msg/PlayGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendGameResult(ctx context.Context, in *MsgSendGameResult, opts ...grpc.CallOption) (*MsgSendGameResultResponse, error) {
	out := new(MsgSendGameResultResponse)
	err := c.cc.Invoke(ctx, "/tmsdkeys.smaller3.smaller.Msg/SendGameResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PlayGame(context.Context, *MsgPlayGame) (*MsgPlayGameResponse, error)
	SendGameResult(context.Context, *MsgSendGameResult) (*MsgSendGameResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlayGame(ctx context.Context, req *MsgPlayGame) (*MsgPlayGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayGame not implemented")
}
func (*UnimplementedMsgServer) SendGameResult(ctx context.Context, req *MsgSendGameResult) (*MsgSendGameResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGameResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlayGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlayGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlayGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tmsdkeys.smaller3.smaller.Msg/PlayGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlayGame(ctx, req.(*MsgPlayGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendGameResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendGameResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendGameResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tmsdkeys.smaller3.smaller.Msg/SendGameResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendGameResult(ctx, req.(*MsgSendGameResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tmsdkeys.smaller3.smaller.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayGame",
			Handler:    _Msg_PlayGame_Handler,
		},
		{
			MethodName: "SendGameResult",
			Handler:    _Msg_SendGameResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smaller/tx.proto",
}

func (m *MsgPlayGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlayGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NumPick) > 0 {
		i -= len(m.NumPick)
		copy(dAtA[i:], m.NumPick)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NumPick)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgPlayGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlayGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Win {
		i--
		if m.Win {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.GameIndex) > 0 {
		i -= len(m.GameIndex)
		copy(dAtA[i:], m.GameIndex)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GameIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendGameResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendGameResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendGameResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x28
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgSendGameResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendGameResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendGameResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgPlayGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NumPick)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPlayGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GameIndex)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Win {
		n += 2
	}
	return n
}

func (m *MsgSendGameResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	return n
}

func (m *MsgSendGameResultResponse) Size() (n int) {
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
func (m *MsgPlayGame) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlayGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayGame: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPick", wireType)
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
			m.NumPick = string(dAtA[iNdEx:postIndex])
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
func (m *MsgPlayGameResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlayGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
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
			m.GameIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Win", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Win = bool(v != 0)
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
func (m *MsgSendGameResult) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendGameResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendGameResult: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
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
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
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
func (m *MsgSendGameResultResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendGameResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendGameResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
