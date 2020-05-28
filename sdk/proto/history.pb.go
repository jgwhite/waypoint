// Code generated by protoc-gen-go. DO NOT EDIT.
// source: history.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type History_LookupRequest_FilterStatus int32

const (
	History_LookupRequest_INVALID History_LookupRequest_FilterStatus = 0
	History_LookupRequest_SUCCESS History_LookupRequest_FilterStatus = 1
	History_LookupRequest_ERROR   History_LookupRequest_FilterStatus = 2
)

var History_LookupRequest_FilterStatus_name = map[int32]string{
	0: "INVALID",
	1: "SUCCESS",
	2: "ERROR",
}

var History_LookupRequest_FilterStatus_value = map[string]int32{
	"INVALID": 0,
	"SUCCESS": 1,
	"ERROR":   2,
}

func (x History_LookupRequest_FilterStatus) String() string {
	return proto.EnumName(History_LookupRequest_FilterStatus_name, int32(x))
}

func (History_LookupRequest_FilterStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0, 0, 0}
}

type History struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *History) Reset()         { *m = History{} }
func (m *History) String() string { return proto.CompactTextString(m) }
func (*History) ProtoMessage()    {}
func (*History) Descriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0}
}

func (m *History) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_History.Unmarshal(m, b)
}
func (m *History) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_History.Marshal(b, m, deterministic)
}
func (m *History) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History.Merge(m, src)
}
func (m *History) XXX_Size() int {
	return xxx_messageInfo_History.Size(m)
}
func (m *History) XXX_DiscardUnknown() {
	xxx_messageInfo_History.DiscardUnknown(m)
}

var xxx_messageInfo_History proto.InternalMessageInfo

type History_LookupRequest struct {
	Limit                int32                              `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	FilterStatus         History_LookupRequest_FilterStatus `protobuf:"varint,2,opt,name=filter_status,json=filterStatus,proto3,enum=proto.History_LookupRequest_FilterStatus" json:"filter_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *History_LookupRequest) Reset()         { *m = History_LookupRequest{} }
func (m *History_LookupRequest) String() string { return proto.CompactTextString(m) }
func (*History_LookupRequest) ProtoMessage()    {}
func (*History_LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0, 0}
}

func (m *History_LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_History_LookupRequest.Unmarshal(m, b)
}
func (m *History_LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_History_LookupRequest.Marshal(b, m, deterministic)
}
func (m *History_LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History_LookupRequest.Merge(m, src)
}
func (m *History_LookupRequest) XXX_Size() int {
	return xxx_messageInfo_History_LookupRequest.Size(m)
}
func (m *History_LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_History_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_History_LookupRequest proto.InternalMessageInfo

func (m *History_LookupRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *History_LookupRequest) GetFilterStatus() History_LookupRequest_FilterStatus {
	if m != nil {
		return m.FilterStatus
	}
	return History_LookupRequest_INVALID
}

type History_LookupResponse struct {
	Results              []*any.Any `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *History_LookupResponse) Reset()         { *m = History_LookupResponse{} }
func (m *History_LookupResponse) String() string { return proto.CompactTextString(m) }
func (*History_LookupResponse) ProtoMessage()    {}
func (*History_LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0, 1}
}

func (m *History_LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_History_LookupResponse.Unmarshal(m, b)
}
func (m *History_LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_History_LookupResponse.Marshal(b, m, deterministic)
}
func (m *History_LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History_LookupResponse.Merge(m, src)
}
func (m *History_LookupResponse) XXX_Size() int {
	return xxx_messageInfo_History_LookupResponse.Size(m)
}
func (m *History_LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_History_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_History_LookupResponse proto.InternalMessageInfo

func (m *History_LookupResponse) GetResults() []*any.Any {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.History_LookupRequest_FilterStatus", History_LookupRequest_FilterStatus_name, History_LookupRequest_FilterStatus_value)
	proto.RegisterType((*History)(nil), "proto.History")
	proto.RegisterType((*History_LookupRequest)(nil), "proto.History.LookupRequest")
	proto.RegisterType((*History_LookupResponse)(nil), "proto.History.LookupResponse")
}

func init() {
	proto.RegisterFile("history.proto", fileDescriptor_454388b49b309873)
}

var fileDescriptor_454388b49b309873 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xe9, 0x4f, 0x0c, 0xde, 0x34, 0x21, 0x0c, 0x5d, 0xc4, 0xa0, 0x10, 0xba, 0x8a,
	0x9b, 0x29, 0xa4, 0x2f, 0x60, 0x69, 0x2b, 0x56, 0x4a, 0x85, 0x09, 0xba, 0x12, 0xa4, 0x95, 0x9b,
	0x1a, 0x4c, 0x33, 0x31, 0x33, 0x11, 0xf2, 0x4a, 0x3e, 0xa1, 0x4b, 0x71, 0x26, 0x81, 0xba, 0xd0,
	0xd5, 0xf0, 0x71, 0xee, 0x39, 0x73, 0x0e, 0xb8, 0x2f, 0xb9, 0x54, 0xa2, 0x6e, 0x59, 0x55, 0x0b,
	0x25, 0xa8, 0xa5, 0x9f, 0xf0, 0x6c, 0x2f, 0xc4, 0xbe, 0xc0, 0x89, 0xa6, 0x5d, 0x93, 0x4d, 0xb6,
	0x65, 0x77, 0x31, 0xfe, 0x24, 0x60, 0xdf, 0x18, 0x4f, 0xf8, 0x41, 0xc0, 0x5d, 0x0b, 0xf1, 0xda,
	0x54, 0x1c, 0xdf, 0x1a, 0x94, 0x8a, 0x8e, 0xc0, 0x2a, 0xf2, 0x43, 0xae, 0x02, 0x12, 0x91, 0xd8,
	0xe2, 0x06, 0xe8, 0x06, 0xdc, 0x2c, 0x2f, 0x14, 0xd6, 0x4f, 0x52, 0x6d, 0x55, 0x23, 0x83, 0x41,
	0x44, 0x62, 0x2f, 0xb9, 0x34, 0x91, 0xac, 0x8b, 0x63, 0x3f, 0xa2, 0xd8, 0xb5, 0x76, 0xa4, 0xda,
	0xc0, 0x87, 0xd9, 0x11, 0x8d, 0xa7, 0x30, 0x3c, 0x56, 0xa9, 0x03, 0xf6, 0x6a, 0xf3, 0x30, 0x5b,
	0xaf, 0x16, 0xfe, 0xbf, 0x6f, 0x48, 0xef, 0xe7, 0xf3, 0x65, 0x9a, 0xfa, 0x84, 0x9e, 0x82, 0xb5,
	0xe4, 0xfc, 0x8e, 0xfb, 0x83, 0xf0, 0x0a, 0xbc, 0xfe, 0x03, 0x59, 0x89, 0x52, 0x22, 0x65, 0x60,
	0xd7, 0x28, 0x9b, 0x42, 0xc9, 0x80, 0x44, 0xff, 0x63, 0x27, 0x19, 0x31, 0xb3, 0x9b, 0xf5, 0xbb,
	0xd9, 0xac, 0x6c, 0x79, 0x7f, 0x94, 0x3c, 0x82, 0xd7, 0x55, 0x4d, 0xb1, 0x7e, 0xcf, 0x9f, 0x91,
	0xde, 0x82, 0xb3, 0xc0, 0xaa, 0x10, 0xed, 0x01, 0x4b, 0x25, 0xe9, 0xf9, 0x5f, 0x83, 0xc2, 0x8b,
	0x5f, 0x54, 0xd3, 0x66, 0x77, 0xa2, 0xd5, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x00,
	0xbd, 0xca, 0x92, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HistoryServiceClient is the client API for HistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HistoryServiceClient interface {
	Deployments(ctx context.Context, in *History_LookupRequest, opts ...grpc.CallOption) (*History_LookupResponse, error)
}

type historyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryServiceClient(cc grpc.ClientConnInterface) HistoryServiceClient {
	return &historyServiceClient{cc}
}

func (c *historyServiceClient) Deployments(ctx context.Context, in *History_LookupRequest, opts ...grpc.CallOption) (*History_LookupResponse, error) {
	out := new(History_LookupResponse)
	err := c.cc.Invoke(ctx, "/proto.HistoryService/Deployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServiceServer is the server API for HistoryService service.
type HistoryServiceServer interface {
	Deployments(context.Context, *History_LookupRequest) (*History_LookupResponse, error)
}

// UnimplementedHistoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHistoryServiceServer struct {
}

func (*UnimplementedHistoryServiceServer) Deployments(ctx context.Context, req *History_LookupRequest) (*History_LookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deployments not implemented")
}

func RegisterHistoryServiceServer(s *grpc.Server, srv HistoryServiceServer) {
	s.RegisterService(&_HistoryService_serviceDesc, srv)
}

func _HistoryService_Deployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(History_LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).Deployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HistoryService/Deployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).Deployments(ctx, req.(*History_LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HistoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HistoryService",
	HandlerType: (*HistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deployments",
			Handler:    _HistoryService_Deployments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}
