// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNodeRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded AppKey
	AppKey        string   `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	RxDelay       uint32   `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset   uint32   `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	ChannelListID int64    `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	RxWindow      RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR         uint32   `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type CreateNodeResponse struct {
}

func (m *CreateNodeResponse) Reset()                    { *m = CreateNodeResponse{} }
func (m *CreateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeResponse) ProtoMessage()               {}
func (*CreateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type GetNodeRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type GetNodeResponse struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded AppKey
	AppKey        string   `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	RxDelay       uint32   `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset   uint32   `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	ChannelListID int64    `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	RxWindow      RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR         uint32   `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
}

func (m *GetNodeResponse) Reset()                    { *m = GetNodeResponse{} }
func (m *GetNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeResponse) ProtoMessage()               {}
func (*GetNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type DeleteNodeRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type ListNodeRequest struct {
	Limit  int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListNodeRequest) Reset()                    { *m = ListNodeRequest{} }
func (m *ListNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodeRequest) ProtoMessage()               {}
func (*ListNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type ListNodeResponse struct {
	TotalCount int64              `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*GetNodeResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNodeResponse) Reset()                    { *m = ListNodeResponse{} }
func (m *ListNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodeResponse) ProtoMessage()               {}
func (*ListNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ListNodeResponse) GetResult() []*GetNodeResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type ListNodeByAppEUIRequest struct {
	Limit  int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,3,opt,name=appEUI" json:"appEUI,omitempty"`
}

func (m *ListNodeByAppEUIRequest) Reset()                    { *m = ListNodeByAppEUIRequest{} }
func (m *ListNodeByAppEUIRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodeByAppEUIRequest) ProtoMessage()               {}
func (*ListNodeByAppEUIRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type UpdateNodeRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded AppKey
	AppKey        string   `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	RxDelay       uint32   `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset   uint32   `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	ChannelListID int64    `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	RxWindow      RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR         uint32   `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
}

func (m *UpdateNodeRequest) Reset()                    { *m = UpdateNodeRequest{} }
func (m *UpdateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeRequest) ProtoMessage()               {}
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type UpdateNodeResponse struct {
}

func (m *UpdateNodeResponse) Reset()                    { *m = UpdateNodeResponse{} }
func (m *UpdateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeResponse) ProtoMessage()               {}
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func init() {
	proto.RegisterType((*CreateNodeRequest)(nil), "api.CreateNodeRequest")
	proto.RegisterType((*CreateNodeResponse)(nil), "api.CreateNodeResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "api.GetNodeRequest")
	proto.RegisterType((*GetNodeResponse)(nil), "api.GetNodeResponse")
	proto.RegisterType((*DeleteNodeRequest)(nil), "api.DeleteNodeRequest")
	proto.RegisterType((*DeleteNodeResponse)(nil), "api.DeleteNodeResponse")
	proto.RegisterType((*ListNodeRequest)(nil), "api.ListNodeRequest")
	proto.RegisterType((*ListNodeResponse)(nil), "api.ListNodeResponse")
	proto.RegisterType((*ListNodeByAppEUIRequest)(nil), "api.ListNodeByAppEUIRequest")
	proto.RegisterType((*UpdateNodeRequest)(nil), "api.UpdateNodeRequest")
	proto.RegisterType((*UpdateNodeResponse)(nil), "api.UpdateNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Node service

type NodeClient interface {
	// Create creates the given node.
	Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error)
	// Get returns the node for the requested DevEUI.
	Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// Delete deletes the node matching the given DevEUI.
	Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error)
	// Update updates the node matching the given DevEUI.
	Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error) {
	out := new(CreateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error) {
	out := new(ListNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// Create creates the given node.
	Create(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
	// Get returns the node for the requested DevEUI.
	Get(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// Delete deletes the node matching the given DevEUI.
	Delete(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(context.Context, *ListNodeRequest) (*ListNodeResponse, error)
	// Update updates the node matching the given DevEUI.
	Update(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Get(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).List(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Update(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Node_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Node_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Node_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("node.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0xc4, 0x2d, 0x53, 0xd2, 0xd0, 0x6d, 0xd2, 0x5a, 0x29, 0x54, 0x91, 0xe1, 0x60,
	0x02, 0x4a, 0x44, 0xb8, 0xf5, 0x82, 0xa0, 0x41, 0x55, 0x55, 0xfe, 0xb4, 0x52, 0xa1, 0x9c, 0x60,
	0xa9, 0xa7, 0xc1, 0x92, 0xb3, 0x6b, 0xec, 0x0d, 0x24, 0xaa, 0x72, 0xe1, 0x15, 0x78, 0x15, 0xde,
	0x84, 0x57, 0xe0, 0x31, 0x10, 0x42, 0xde, 0xdd, 0xa4, 0x6e, 0x9c, 0x43, 0xc5, 0xb5, 0xc7, 0xf9,
	0x3c, 0xfb, 0xed, 0x37, 0x33, 0xdf, 0xac, 0x01, 0xb8, 0x08, 0xb0, 0x13, 0x27, 0x42, 0x0a, 0x62,
	0xb3, 0x38, 0x6c, 0xde, 0x1e, 0x08, 0x31, 0x88, 0xb0, 0xcb, 0xe2, 0xb0, 0xcb, 0x38, 0x17, 0x92,
	0xc9, 0x50, 0xf0, 0x54, 0xa7, 0x34, 0x6f, 0x9e, 0x8a, 0xe1, 0x50, 0x70, 0x1d, 0x79, 0x7f, 0x2d,
	0xd8, 0xd8, 0x4f, 0x90, 0x49, 0x7c, 0x25, 0x02, 0xa4, 0xf8, 0x65, 0x84, 0xa9, 0x24, 0x5b, 0xe0,
	0x04, 0xf8, 0xf5, 0xf9, 0xf1, 0xa1, 0x6b, 0xb5, 0x2c, 0xff, 0x06, 0x35, 0x51, 0x86, 0xb3, 0x38,
	0xce, 0xf0, 0x92, 0xc6, 0x75, 0x64, 0xf0, 0x23, 0x9c, 0xb8, 0xf6, 0x1c, 0x3f, 0xc2, 0x09, 0x71,
	0x61, 0x25, 0x19, 0xf7, 0x31, 0x62, 0x13, 0xb7, 0xdc, 0xb2, 0xfc, 0x2a, 0x9d, 0x85, 0xa4, 0x05,
	0x6b, 0xc9, 0xf8, 0x51, 0x9f, 0xbe, 0x3e, 0x3b, 0x4b, 0x51, 0xba, 0x15, 0xf5, 0x35, 0x0f, 0x91,
	0x7b, 0x50, 0x3d, 0xfd, 0xcc, 0x38, 0xc7, 0xe8, 0x45, 0x98, 0xca, 0xc3, 0xbe, 0xeb, 0xb4, 0x2c,
	0xdf, 0xa6, 0x97, 0x41, 0x72, 0x1f, 0x56, 0x93, 0xf1, 0xbb, 0x90, 0x07, 0xe2, 0x9b, 0xbb, 0xd2,
	0xb2, 0xfc, 0xf5, 0x5e, 0xb5, 0xc3, 0xe2, 0xb0, 0x43, 0x4f, 0x34, 0x48, 0xe7, 0x9f, 0x49, 0x1d,
	0x2a, 0xc9, 0xb8, 0xd7, 0xa7, 0xee, 0xaa, 0xba, 0x4c, 0x07, 0x5e, 0x1d, 0x48, 0xbe, 0xfe, 0x34,
	0x16, 0x3c, 0x45, 0xcf, 0x87, 0xf5, 0x03, 0x94, 0x57, 0x68, 0x89, 0xf7, 0xc7, 0x82, 0xda, 0x3c,
	0x55, 0x9f, 0xbe, 0x4e, 0xed, 0x7b, 0x00, 0x1b, 0x7d, 0x8c, 0xf0, 0x4a, 0xf6, 0xc9, 0x7a, 0x9d,
	0x4f, 0x36, 0xbd, 0x7e, 0x02, 0xb5, 0x4c, 0x4d, 0x9e, 0xa0, 0x0e, 0x95, 0x28, 0x1c, 0x86, 0x52,
	0x9d, 0xb7, 0xa9, 0x0e, 0x32, 0x5a, 0xa1, 0xeb, 0x2d, 0x29, 0xd8, 0x44, 0xde, 0x47, 0xb8, 0x75,
	0x41, 0x60, 0x46, 0xb0, 0x0b, 0x20, 0x85, 0x64, 0xd1, 0xbe, 0x18, 0xf1, 0x19, 0x4d, 0x0e, 0x21,
	0x0f, 0xc1, 0x49, 0x30, 0x1d, 0x45, 0x19, 0x97, 0xed, 0xaf, 0xf5, 0xea, 0xaa, 0xec, 0x85, 0x41,
	0x52, 0x93, 0xe3, 0x7d, 0x80, 0xed, 0xd9, 0x0d, 0xcf, 0x26, 0x4f, 0xd5, 0xd0, 0xfe, 0x4b, 0x6a,
	0xce, 0x01, 0x76, 0xde, 0x01, 0x6a, 0x0d, 0x8f, 0xe3, 0xe0, 0x5a, 0xaf, 0x61, 0xbe, 0x7e, 0xdd,
	0xff, 0xde, 0x4f, 0x1b, 0xca, 0x19, 0x40, 0xde, 0x80, 0xa3, 0xb7, 0x94, 0x6c, 0x29, 0xde, 0xc2,
	0x93, 0xd5, 0xdc, 0x2e, 0xe0, 0xc6, 0x5e, 0x8d, 0xef, 0xbf, 0x7e, 0xff, 0x28, 0xd5, 0x3c, 0x50,
	0xef, 0xe1, 0x20, 0x7b, 0x2c, 0xf7, 0xac, 0x36, 0x79, 0x09, 0xf6, 0x01, 0x4a, 0xb2, 0x79, 0x79,
	0xee, 0x9a, 0x6b, 0xa9, 0x19, 0xbc, 0x1d, 0x45, 0xd4, 0x20, 0x9b, 0x17, 0x44, 0xdd, 0x73, 0x3d,
	0x91, 0x29, 0x79, 0x0b, 0x8e, 0xb6, 0xb6, 0x11, 0x58, 0x58, 0x0a, 0x23, 0x70, 0x89, 0xff, 0x0d,
	0x6f, 0x7b, 0x29, 0xef, 0x09, 0x94, 0xb3, 0x16, 0x13, 0x2d, 0x69, 0x61, 0x4f, 0x9a, 0x8d, 0x05,
	0xd4, 0x30, 0xde, 0x55, 0x8c, 0x77, 0xc8, 0x4e, 0x9e, 0x51, 0x3b, 0x70, 0xda, 0x3d, 0x57, 0x0e,
	0x9d, 0x92, 0xf7, 0xe0, 0xe8, 0x8e, 0x1b, 0xc5, 0x05, 0xfb, 0x19, 0xc5, 0xc5, 0xb1, 0x78, 0xbb,
	0x8a, 0xdf, 0x6d, 0x2e, 0x53, 0xbc, 0x67, 0xb5, 0x3f, 0x39, 0xea, 0xdf, 0xf2, 0xf8, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x89, 0x3e, 0x0d, 0x9a, 0x06, 0x00, 0x00,
}
