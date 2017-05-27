// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bucket.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	bucket.proto
	registry.proto
	types.proto

It has these top-level messages:
	NewItem
	PutSummary
	Page
	Key
	Item
	NewBucket
	Bucket
	BucketStatus
	Empty
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// NewItem is used to put an item in a bucket.
type NewItem struct {
	// Item to save in the bucket.
	// @inject_tag: valid:"required"
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty" valid:"required"`
	// Bucket name.
	// @inject_tag: valid:"required"
	Bucket string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty" valid:"required"`
}

func (m *NewItem) Reset()                    { *m = NewItem{} }
func (m *NewItem) String() string            { return proto1.CompactTextString(m) }
func (*NewItem) ProtoMessage()               {}
func (*NewItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewItem) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *NewItem) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

// PutSummary is received in response to a Put route.
type PutSummary struct {
	// The number of items received.
	ItemCount int32 `protobuf:"varint,1,opt,name=item_count,json=itemCount" json:"item_count,omitempty" `
}

func (m *PutSummary) Reset()                    { *m = PutSummary{} }
func (m *PutSummary) String() string            { return proto1.CompactTextString(m) }
func (*PutSummary) ProtoMessage()               {}
func (*PutSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PutSummary) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

type Page struct {
	// @inject_tag: valid:"required"
	Bucket  string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty" valid:"required"`
	Page    int32  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	PerPage int32  `protobuf:"varint,3,opt,name=perPage" json:"perPage,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto1.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Page) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Page) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Page) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type Key struct {
	// @inject_tag: valid:"required"
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty" valid:"required"`
	// @inject_tag: valid:"required"
	Key string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty" valid:"required"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto1.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Key) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Item struct {
	// @inject_tag: valid:"required"
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty" valid:"required"`
	// @inject_tag: valid:"required"
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" valid:"required"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto1.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Item) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterType((*NewItem)(nil), "proto.NewItem")
	proto1.RegisterType((*PutSummary)(nil), "proto.PutSummary")
	proto1.RegisterType((*Page)(nil), "proto.Page")
	proto1.RegisterType((*Key)(nil), "proto.Key")
	proto1.RegisterType((*Item)(nil), "proto.Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BucketService service

type BucketServiceClient interface {
	// Put user data
	Put(ctx context.Context, opts ...grpc.CallOption) (BucketService_PutClient, error)
	// Get an item
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Item, error)
	// Delete an item
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Empty, error)
	// List the bucket content
	List(ctx context.Context, in *Page, opts ...grpc.CallOption) (BucketService_ListClient, error)
}

type bucketServiceClient struct {
	cc *grpc.ClientConn
}

func NewBucketServiceClient(cc *grpc.ClientConn) BucketServiceClient {
	return &bucketServiceClient{cc}
}

func (c *bucketServiceClient) Put(ctx context.Context, opts ...grpc.CallOption) (BucketService_PutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BucketService_serviceDesc.Streams[0], c.cc, "/proto.BucketService/Put", opts...)
	if err != nil {
		return nil, err
	}
	x := &bucketServicePutClient{stream}
	return x, nil
}

type BucketService_PutClient interface {
	Send(*NewItem) error
	CloseAndRecv() (*PutSummary, error)
	grpc.ClientStream
}

type bucketServicePutClient struct {
	grpc.ClientStream
}

func (x *bucketServicePutClient) Send(m *NewItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bucketServicePutClient) CloseAndRecv() (*PutSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bucketServiceClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := grpc.Invoke(ctx, "/proto.BucketService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.BucketService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) List(ctx context.Context, in *Page, opts ...grpc.CallOption) (BucketService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BucketService_serviceDesc.Streams[1], c.cc, "/proto.BucketService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &bucketServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BucketService_ListClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type bucketServiceListClient struct {
	grpc.ClientStream
}

func (x *bucketServiceListClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BucketService service

type BucketServiceServer interface {
	// Put user data
	Put(BucketService_PutServer) error
	// Get an item
	Get(context.Context, *Key) (*Item, error)
	// Delete an item
	Delete(context.Context, *Key) (*Empty, error)
	// List the bucket content
	List(*Page, BucketService_ListServer) error
}

func RegisterBucketServiceServer(s *grpc.Server, srv BucketServiceServer) {
	s.RegisterService(&_BucketService_serviceDesc, srv)
}

func _BucketService_Put_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BucketServiceServer).Put(&bucketServicePutServer{stream})
}

type BucketService_PutServer interface {
	SendAndClose(*PutSummary) error
	Recv() (*NewItem, error)
	grpc.ServerStream
}

type bucketServicePutServer struct {
	grpc.ServerStream
}

func (x *bucketServicePutServer) SendAndClose(m *PutSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bucketServicePutServer) Recv() (*NewItem, error) {
	m := new(NewItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BucketService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BucketService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BucketService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Page)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BucketServiceServer).List(m, &bucketServiceListServer{stream})
}

type BucketService_ListServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type bucketServiceListServer struct {
	grpc.ServerStream
}

func (x *bucketServiceListServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

var _BucketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BucketService",
	HandlerType: (*BucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BucketService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BucketService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Put",
			Handler:       _BucketService_Put_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _BucketService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bucket.proto",
}

func init() { proto1.RegisterFile("bucket.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0xed, 0x7e, 0x49, 0x5a, 0x3a, 0xed, 0x27, 0x3a, 0x88, 0x84, 0x82, 0x58, 0x16, 0x2f, 0x0a,
	0x4a, 0x94, 0xfa, 0x06, 0x55, 0x11, 0x69, 0x91, 0x90, 0x3e, 0x80, 0xa4, 0x61, 0x28, 0xa1, 0x8d,
	0x09, 0xe9, 0x6c, 0x65, 0x9f, 0xc9, 0x97, 0x94, 0xdd, 0x6c, 0x35, 0x0a, 0x5e, 0x65, 0xce, 0xcf,
	0x9e, 0xc9, 0x1c, 0x18, 0xae, 0x54, 0xb6, 0x21, 0x8e, 0xaa, 0xba, 0xe4, 0x12, 0x03, 0xfb, 0x19,
	0x0d, 0x58, 0x57, 0xb4, 0x6b, 0x38, 0x39, 0x83, 0xde, 0x0b, 0xbd, 0x3f, 0x33, 0x15, 0x78, 0x01,
	0x7e, 0xce, 0x54, 0x84, 0x62, 0x2c, 0x26, 0x83, 0xe9, 0xa0, 0x31, 0x44, 0x46, 0x4a, 0xac, 0x80,
	0x67, 0xd0, 0x6d, 0xf2, 0xc2, 0x7f, 0x63, 0x31, 0xe9, 0x27, 0x0e, 0xc9, 0x2b, 0x80, 0x58, 0xf1,
	0x52, 0x15, 0x45, 0x5a, 0x6b, 0x3c, 0x07, 0x30, 0xee, 0xd7, 0xac, 0x54, 0x6f, 0x6c, 0xc3, 0x82,
	0xa4, 0x6f, 0x98, 0x7b, 0x43, 0xc8, 0x05, 0xf8, 0x71, 0xba, 0xa6, 0x56, 0x98, 0x68, 0x87, 0x21,
	0x82, 0x5f, 0xa5, 0x6b, 0xb2, 0x2b, 0x82, 0xc4, 0xce, 0x18, 0x42, 0xaf, 0xa2, 0xda, 0x3c, 0x0b,
	0x3d, 0x4b, 0x1f, 0xa0, 0xbc, 0x01, 0x6f, 0x4e, 0xfa, 0xcf, 0xb0, 0x63, 0xf0, 0x36, 0xa4, 0xdd,
	0xef, 0x9a, 0x51, 0x46, 0xe0, 0xdb, 0x63, 0x9d, 0x22, 0xbe, 0x14, 0x3c, 0x85, 0x60, 0x9f, 0x6e,
	0x55, 0xb3, 0x79, 0x98, 0x34, 0x60, 0xfa, 0x21, 0xe0, 0xff, 0xcc, 0x86, 0x2d, 0xa9, 0xde, 0xe7,
	0x19, 0xe1, 0x35, 0x78, 0xb1, 0x62, 0x3c, 0x72, 0xfd, 0xb8, 0xf6, 0x46, 0x27, 0x0e, 0x7f, 0x37,
	0x21, 0x3b, 0x13, 0x81, 0x63, 0xf0, 0x9e, 0x88, 0x11, 0x9c, 0x3a, 0x27, 0x3d, 0x6a, 0x37, 0x2b,
	0x3b, 0x78, 0x09, 0xdd, 0x07, 0xda, 0x12, 0xd3, 0x0f, 0xd3, 0xd0, 0xcd, 0x8f, 0x45, 0xc5, 0xda,
	0xba, 0xfc, 0x45, 0xbe, 0x63, 0x3c, 0x3c, 0x36, 0xd7, 0xff, 0x4a, 0xba, 0x15, 0xab, 0xae, 0xc5,
	0x77, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x58, 0x31, 0x25, 0xf8, 0x01, 0x00, 0x00,
}
