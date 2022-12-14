// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/visitor.proto

package grpc_e2e_tests

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type VisitorCounterRequest struct {
	Visitor              *Visitor `protobuf:"bytes,1,opt,name=visitor,proto3" json:"visitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitorCounterRequest) Reset()         { *m = VisitorCounterRequest{} }
func (m *VisitorCounterRequest) String() string { return proto.CompactTextString(m) }
func (*VisitorCounterRequest) ProtoMessage()    {}
func (*VisitorCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3a50b293680181, []int{0}
}

func (m *VisitorCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitorCounterRequest.Unmarshal(m, b)
}
func (m *VisitorCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitorCounterRequest.Marshal(b, m, deterministic)
}
func (m *VisitorCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitorCounterRequest.Merge(m, src)
}
func (m *VisitorCounterRequest) XXX_Size() int {
	return xxx_messageInfo_VisitorCounterRequest.Size(m)
}
func (m *VisitorCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitorCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VisitorCounterRequest proto.InternalMessageInfo

func (m *VisitorCounterRequest) GetVisitor() *Visitor {
	if m != nil {
		return m.Visitor
	}
	return nil
}

type UpdateVisitorRequest struct {
	Visitor              *Visitor `protobuf:"bytes,1,opt,name=visitor,proto3" json:"visitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVisitorRequest) Reset()         { *m = UpdateVisitorRequest{} }
func (m *UpdateVisitorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVisitorRequest) ProtoMessage()    {}
func (*UpdateVisitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3a50b293680181, []int{1}
}

func (m *UpdateVisitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVisitorRequest.Unmarshal(m, b)
}
func (m *UpdateVisitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVisitorRequest.Marshal(b, m, deterministic)
}
func (m *UpdateVisitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVisitorRequest.Merge(m, src)
}
func (m *UpdateVisitorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVisitorRequest.Size(m)
}
func (m *UpdateVisitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVisitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVisitorRequest proto.InternalMessageInfo

func (m *UpdateVisitorRequest) GetVisitor() *Visitor {
	if m != nil {
		return m.Visitor
	}
	return nil
}

type UpdateVisitorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVisitorResponse) Reset()         { *m = UpdateVisitorResponse{} }
func (m *UpdateVisitorResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateVisitorResponse) ProtoMessage()    {}
func (*UpdateVisitorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3a50b293680181, []int{2}
}

func (m *UpdateVisitorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVisitorResponse.Unmarshal(m, b)
}
func (m *UpdateVisitorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVisitorResponse.Marshal(b, m, deterministic)
}
func (m *UpdateVisitorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVisitorResponse.Merge(m, src)
}
func (m *UpdateVisitorResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateVisitorResponse.Size(m)
}
func (m *UpdateVisitorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVisitorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVisitorResponse proto.InternalMessageInfo

type VisitorCounterResponse struct {
	Visitor              *Visitor `protobuf:"bytes,1,opt,name=visitor,proto3" json:"visitor,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitorCounterResponse) Reset()         { *m = VisitorCounterResponse{} }
func (m *VisitorCounterResponse) String() string { return proto.CompactTextString(m) }
func (*VisitorCounterResponse) ProtoMessage()    {}
func (*VisitorCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3a50b293680181, []int{3}
}

func (m *VisitorCounterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitorCounterResponse.Unmarshal(m, b)
}
func (m *VisitorCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitorCounterResponse.Marshal(b, m, deterministic)
}
func (m *VisitorCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitorCounterResponse.Merge(m, src)
}
func (m *VisitorCounterResponse) XXX_Size() int {
	return xxx_messageInfo_VisitorCounterResponse.Size(m)
}
func (m *VisitorCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitorCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VisitorCounterResponse proto.InternalMessageInfo

func (m *VisitorCounterResponse) GetVisitor() *Visitor {
	if m != nil {
		return m.Visitor
	}
	return nil
}

func (m *VisitorCounterResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Visitor struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Visitor) Reset()         { *m = Visitor{} }
func (m *Visitor) String() string { return proto.CompactTextString(m) }
func (*Visitor) ProtoMessage()    {}
func (*Visitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd3a50b293680181, []int{4}
}

func (m *Visitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Visitor.Unmarshal(m, b)
}
func (m *Visitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Visitor.Marshal(b, m, deterministic)
}
func (m *Visitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Visitor.Merge(m, src)
}
func (m *Visitor) XXX_Size() int {
	return xxx_messageInfo_Visitor.Size(m)
}
func (m *Visitor) XXX_DiscardUnknown() {
	xxx_messageInfo_Visitor.DiscardUnknown(m)
}

var xxx_messageInfo_Visitor proto.InternalMessageInfo

func (m *Visitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*VisitorCounterRequest)(nil), "skaffold.examples.e2e.visitor.VisitorCounterRequest")
	proto.RegisterType((*UpdateVisitorRequest)(nil), "skaffold.examples.e2e.visitor.UpdateVisitorRequest")
	proto.RegisterType((*UpdateVisitorResponse)(nil), "skaffold.examples.e2e.visitor.UpdateVisitorResponse")
	proto.RegisterType((*VisitorCounterResponse)(nil), "skaffold.examples.e2e.visitor.VisitorCounterResponse")
	proto.RegisterType((*Visitor)(nil), "skaffold.examples.e2e.visitor.Visitor")
}

func init() {
	proto.RegisterFile("proto/visitor.proto", fileDescriptor_fd3a50b293680181)
}

var fileDescriptor_fd3a50b293680181 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0x2d, 0x7e, 0x10, 0xc7, 0xe0, 0x61, 0x05, 0x25, 0x24, 0x24, 0xa4, 0x07, 0xc3, 0x85,
	0x6d, 0x52, 0xf0, 0x4e, 0xe0, 0xc0, 0xbd, 0x51, 0xa3, 0xde, 0x4a, 0x19, 0x6a, 0x63, 0xdb, 0x59,
	0xbb, 0x53, 0xe3, 0xc1, 0xf8, 0x08, 0x3e, 0xb3, 0x71, 0xdb, 0x1e, 0x4a, 0x8c, 0x4a, 0xe4, 0x36,
	0x5f, 0xff, 0xdf, 0xcc, 0xce, 0x2c, 0x9c, 0xa9, 0x8c, 0x98, 0x9c, 0x97, 0x48, 0x47, 0x4c, 0x99,
	0x34, 0x9e, 0xe8, 0xeb, 0x27, 0x7f, 0xbd, 0xa6, 0x78, 0x25, 0xf1, 0xd5, 0x4f, 0x54, 0x8c, 0x5a,
	0xa2, 0x8b, 0xb2, 0x2c, 0xb2, 0xef, 0xa1, 0x73, 0x5b, 0x98, 0x73, 0xca, 0x53, 0xc6, 0xcc, 0xc3,
	0xe7, 0x1c, 0x35, 0x8b, 0x29, 0x34, 0xcb, 0x9a, 0xae, 0x35, 0xb0, 0x86, 0x27, 0xee, 0xa5, 0xfc,
	0x91, 0x24, 0x4b, 0x8c, 0x57, 0xc9, 0xec, 0x3b, 0x68, 0xdf, 0xa8, 0x95, 0xcf, 0x58, 0x65, 0x76,
	0x46, 0xbe, 0x80, 0xce, 0x06, 0x59, 0x2b, 0x4a, 0x35, 0xda, 0x0a, 0xce, 0x37, 0x5f, 0x53, 0x64,
	0xfe, 0xdf, 0x54, 0xb4, 0xe1, 0x30, 0xf8, 0x82, 0x76, 0x1b, 0x03, 0x6b, 0xb8, 0xef, 0x15, 0x8e,
	0xdd, 0x87, 0x66, 0x59, 0x29, 0x04, 0x1c, 0xa4, 0x7e, 0x82, 0x86, 0x7f, 0xec, 0x19, 0xdb, 0xfd,
	0x68, 0xc0, 0x69, 0x7d, 0x22, 0xf1, 0x0e, 0xad, 0x05, 0xb2, 0x09, 0x9a, 0x90, 0x98, 0xfc, 0x6d,
	0x92, 0xfa, 0x7d, 0x7a, 0x57, 0x5b, 0xaa, 0xca, 0x0d, 0xed, 0x89, 0x37, 0x68, 0xd5, 0x96, 0x27,
	0xc6, 0xbf, 0x90, 0xbe, 0x3b, 0x62, 0x6f, 0xb2, 0x9d, 0xa8, 0xea, 0x3e, 0x9b, 0x3d, 0x4c, 0xc3,
	0x88, 0x1f, 0xf3, 0xa5, 0x0c, 0x28, 0x71, 0x16, 0x44, 0x61, 0x8c, 0x73, 0x4a, 0xd9, 0x8f, 0x52,
	0xcc, 0xae, 0x89, 0x62, 0xed, 0x54, 0x60, 0xa7, 0x02, 0x3b, 0x61, 0xa6, 0x82, 0x11, 0xba, 0x38,
	0x62, 0xd4, 0xac, 0x97, 0x47, 0xe6, 0x67, 0x8f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x79, 0x64,
	0x1e, 0x54, 0xf0, 0x02, 0x00, 0x00,
}
