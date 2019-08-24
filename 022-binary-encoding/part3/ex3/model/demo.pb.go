// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package model

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

type SearchReq struct {
	Q                    string            `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
	Params               map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id                   int64             `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SearchReq) Reset()         { *m = SearchReq{} }
func (m *SearchReq) String() string { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()    {}
func (*SearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *SearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReq.Unmarshal(m, b)
}
func (m *SearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReq.Marshal(b, m, deterministic)
}
func (m *SearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReq.Merge(m, src)
}
func (m *SearchReq) XXX_Size() int {
	return xxx_messageInfo_SearchReq.Size(m)
}
func (m *SearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReq proto.InternalMessageInfo

func (m *SearchReq) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchReq) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SearchReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SearchResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ans                  string   `protobuf:"bytes,2,opt,name=ans,proto3" json:"ans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResp) Reset()         { *m = SearchResp{} }
func (m *SearchResp) String() string { return proto.CompactTextString(m) }
func (*SearchResp) ProtoMessage()    {}
func (*SearchResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *SearchResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResp.Unmarshal(m, b)
}
func (m *SearchResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResp.Marshal(b, m, deterministic)
}
func (m *SearchResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResp.Merge(m, src)
}
func (m *SearchResp) XXX_Size() int {
	return xxx_messageInfo_SearchResp.Size(m)
}
func (m *SearchResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResp.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResp proto.InternalMessageInfo

func (m *SearchResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SearchResp) GetAns() string {
	if m != nil {
		return m.Ans
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchReq)(nil), "model.other.thing.SearchReq")
	proto.RegisterMapType((map[string]string)(nil), "model.other.thing.SearchReq.ParamsEntry")
	proto.RegisterType((*SearchResp)(nil), "model.other.thing.SearchResp")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0xcb, 0x2f,
	0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0xc9, 0xc8, 0xcc, 0x4b, 0x57, 0x5a, 0xc6, 0xc8, 0xc5, 0x19, 0x9c,
	0x9a, 0x58, 0x94, 0x9c, 0x11, 0x94, 0x5a, 0x28, 0xc4, 0xc3, 0xc5, 0x58, 0x28, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0xc4, 0x58, 0x28, 0xe4, 0xc0, 0xc5, 0x56, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x2c,
	0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa1, 0x87, 0xa1, 0x5f, 0x0f, 0xae, 0x57, 0x2f, 0x00,
	0xac, 0xd4, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0xaa, 0x4f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45,
	0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0x33, 0x45, 0xca, 0x92, 0x8b, 0x1b, 0x49, 0x99,
	0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd4, 0x42, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x2c,
	0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x2a, 0xe9, 0x71,
	0x71, 0xc1, 0xec, 0x2a, 0x2e, 0x80, 0x1a, 0xcc, 0x08, 0x33, 0x18, 0x64, 0x52, 0x62, 0x5e, 0x31,
	0x54, 0x17, 0x88, 0xe9, 0xc4, 0x1e, 0xc5, 0x0a, 0x76, 0x6d, 0x12, 0x1b, 0xd8, 0xef, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x6f, 0xea, 0x6e, 0x09, 0x01, 0x00, 0x00,
}