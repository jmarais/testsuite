// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topsyturvy.proto

package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TopsyTurvy struct {
	Top                  *Topsy   `protobuf:"bytes,1,opt,name=Top" json:"Top,omitempty"`
	Turf                 *Turvy   `protobuf:"bytes,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopsyTurvy) Reset()         { *m = TopsyTurvy{} }
func (m *TopsyTurvy) String() string { return proto.CompactTextString(m) }
func (*TopsyTurvy) ProtoMessage()    {}
func (*TopsyTurvy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{0}
}
func (m *TopsyTurvy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopsyTurvy.Unmarshal(m, b)
}
func (m *TopsyTurvy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopsyTurvy.Marshal(b, m, deterministic)
}
func (dst *TopsyTurvy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopsyTurvy.Merge(dst, src)
}
func (m *TopsyTurvy) XXX_Size() int {
	return xxx_messageInfo_TopsyTurvy.Size(m)
}
func (m *TopsyTurvy) XXX_DiscardUnknown() {
	xxx_messageInfo_TopsyTurvy.DiscardUnknown(m)
}

var xxx_messageInfo_TopsyTurvy proto.InternalMessageInfo

func (m *TopsyTurvy) GetTop() *Topsy {
	if m != nil {
		return m.Top
	}
	return nil
}

func (m *TopsyTurvy) GetTurf() *Turvy {
	if m != nil {
		return m.Turf
	}
	return nil
}

type Topsy struct {
	Top                  *int64   `protobuf:"varint,1,opt,name=Top" json:"Top,omitempty"`
	Turf                 *int64   `protobuf:"varint,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topsy) Reset()         { *m = Topsy{} }
func (m *Topsy) String() string { return proto.CompactTextString(m) }
func (*Topsy) ProtoMessage()    {}
func (*Topsy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{1}
}
func (m *Topsy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topsy.Unmarshal(m, b)
}
func (m *Topsy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topsy.Marshal(b, m, deterministic)
}
func (dst *Topsy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topsy.Merge(dst, src)
}
func (m *Topsy) XXX_Size() int {
	return xxx_messageInfo_Topsy.Size(m)
}
func (m *Topsy) XXX_DiscardUnknown() {
	xxx_messageInfo_Topsy.DiscardUnknown(m)
}

var xxx_messageInfo_Topsy proto.InternalMessageInfo

func (m *Topsy) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func (m *Topsy) GetTurf() int64 {
	if m != nil && m.Turf != nil {
		return *m.Turf
	}
	return 0
}

type Turvy struct {
	Turf                 *int64   `protobuf:"varint,1,opt,name=Turf" json:"Turf,omitempty"`
	Top                  *int64   `protobuf:"varint,2,opt,name=Top" json:"Top,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Turvy) Reset()         { *m = Turvy{} }
func (m *Turvy) String() string { return proto.CompactTextString(m) }
func (*Turvy) ProtoMessage()    {}
func (*Turvy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{2}
}
func (m *Turvy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Turvy.Unmarshal(m, b)
}
func (m *Turvy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Turvy.Marshal(b, m, deterministic)
}
func (dst *Turvy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Turvy.Merge(dst, src)
}
func (m *Turvy) XXX_Size() int {
	return xxx_messageInfo_Turvy.Size(m)
}
func (m *Turvy) XXX_DiscardUnknown() {
	xxx_messageInfo_Turvy.DiscardUnknown(m)
}

var xxx_messageInfo_Turvy proto.InternalMessageInfo

func (m *Turvy) GetTurf() int64 {
	if m != nil && m.Turf != nil {
		return *m.Turf
	}
	return 0
}

func (m *Turvy) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func init() {
	proto.RegisterType((*TopsyTurvy)(nil), "main.TopsyTurvy")
	proto.RegisterType((*Topsy)(nil), "main.Topsy")
	proto.RegisterType((*Turvy)(nil), "main.Turvy")
}
func (this *TopsyTurvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Topsy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Turvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func TopsyturvyDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3771 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0xff, 0x24, 0xf2, 0x90, 0xa2, 0x46, 0x57, 0xf2, 0x2e, 0x57, 0x8e, 0xbd, 0x5a, 0xda,
		0x8e, 0x65, 0xbb, 0xe6, 0x06, 0x6b, 0xef, 0xda, 0xcb, 0x6d, 0xe2, 0x52, 0xd4, 0xac, 0x42, 0x57,
		0x12, 0x99, 0x21, 0x15, 0xff, 0x04, 0xc5, 0x60, 0x34, 0xbc, 0xa4, 0x66, 0x77, 0x38, 0x33, 0x99,
		0x19, 0xee, 0x5a, 0x8b, 0x3e, 0xb8, 0x70, 0x7f, 0x10, 0x14, 0xfd, 0x2f, 0xd0, 0xc4, 0x75, 0xdc,
		0xa6, 0x40, 0xea, 0x36, 0xfd, 0x4b, 0x9a, 0x36, 0x6d, 0xfa, 0xd4, 0x97, 0xb4, 0x7d, 0x2a, 0x90,
		0xf7, 0x3e, 0xf4, 0xc7, 0x40, 0xff, 0xdc, 0x26, 0x6d, 0xfd, 0x50, 0xc0, 0x28, 0x50, 0xdc, 0xbf,
		0xe1, 0x0c, 0x49, 0x69, 0x46, 0x01, 0x6c, 0x3f, 0x49, 0x73, 0xee, 0xf9, 0xbe, 0x39, 0xf7, 0xdc,
		0x73, 0xcf, 0x39, 0xf7, 0x0e, 0xe1, 0x7b, 0xd7, 0x61, 0x63, 0x68, 0xdb, 0x43, 0x13, 0x5f, 0x76,
		0x5c, 0xdb, 0xb7, 0x0f, 0xc7, 0x83, 0xcb, 0x7d, 0xec, 0xe9, 0xae, 0xe1, 0xf8, 0xb6, 0x5b, 0xa3,
		0x32, 0xb4, 0xcc, 0x34, 0x6a, 0x42, 0xa3, 0xba, 0x07, 0x2b, 0x37, 0x0d, 0x13, 0x6f, 0x07, 0x8a,
		0x5d, 0xec, 0xa3, 0xe7, 0x20, 0x3b, 0x30, 0x4c, 0x5c, 0x49, 0x6d, 0x64, 0x36, 0x8b, 0x57, 0x1e,
		0xa9, 0x4d, 0x81, 0x6a, 0x51, 0x44, 0x87, 0x88, 0x15, 0x8a, 0xa8, 0xbe, 0x93, 0x85, 0xd5, 0x39,
		0xa3, 0x08, 0x41, 0xd6, 0xd2, 0x46, 0x84, 0x31, 0xb5, 0x59, 0x50, 0xe8, 0xff, 0xa8, 0x02, 0x8b,
		0x8e, 0xa6, 0xdf, 0xd6, 0x86, 0xb8, 0x92, 0xa6, 0x62, 0xf1, 0x88, 0x1e, 0x02, 0xe8, 0x63, 0x07,
		0x5b, 0x7d, 0x6c, 0xe9, 0xc7, 0x95, 0xcc, 0x46, 0x66, 0xb3, 0xa0, 0x84, 0x24, 0xe8, 0x49, 0x58,
		0x71, 0xc6, 0x87, 0xa6, 0xa1, 0xab, 0x21, 0x35, 0xd8, 0xc8, 0x6c, 0xe6, 0x14, 0x89, 0x0d, 0x6c,
		0x4f, 0x94, 0x1f, 0x83, 0xe5, 0xbb, 0x58, 0xbb, 0x1d, 0x56, 0x2d, 0x52, 0xd5, 0x32, 0x11, 0x87,
		0x14, 0x9b, 0x50, 0x1a, 0x61, 0xcf, 0xd3, 0x86, 0x58, 0xf5, 0x8f, 0x1d, 0x5c, 0xc9, 0xd2, 0xd9,
		0x6f, 0xcc, 0xcc, 0x7e, 0x7a, 0xe6, 0x45, 0x8e, 0xea, 0x1d, 0x3b, 0x18, 0x35, 0xa0, 0x80, 0xad,
		0xf1, 0x88, 0x31, 0xe4, 0x4e, 0xf0, 0x9f, 0x6c, 0x8d, 0x47, 0xd3, 0x2c, 0x79, 0x02, 0xe3, 0x14,
		0x8b, 0x1e, 0x76, 0xef, 0x18, 0x3a, 0xae, 0x2c, 0x50, 0x82, 0xc7, 0x66, 0x08, 0xba, 0x6c, 0x7c,
		0x9a, 0x43, 0xe0, 0x50, 0x13, 0x0a, 0xf8, 0x55, 0x1f, 0x5b, 0x9e, 0x61, 0x5b, 0x95, 0x45, 0x4a,
		0xf2, 0xe8, 0x9c, 0x55, 0xc4, 0x66, 0x7f, 0x9a, 0x62, 0x82, 0x43, 0xd7, 0x60, 0xd1, 0x76, 0x7c,
		0xc3, 0xb6, 0xbc, 0x4a, 0x7e, 0x23, 0xb5, 0x59, 0xbc, 0xf2, 0xb1, 0xb9, 0x81, 0xd0, 0x66, 0x3a,
		0x8a, 0x50, 0x46, 0x2d, 0x90, 0x3c, 0x7b, 0xec, 0xea, 0x58, 0xd5, 0xed, 0x3e, 0x56, 0x0d, 0x6b,
		0x60, 0x57, 0x0a, 0x94, 0xe0, 0xe2, 0xec, 0x44, 0xa8, 0x62, 0xd3, 0xee, 0xe3, 0x96, 0x35, 0xb0,
		0x95, 0xb2, 0x17, 0x79, 0x46, 0xe7, 0x60, 0xc1, 0x3b, 0xb6, 0x7c, 0xed, 0xd5, 0x4a, 0x89, 0x46,
		0x08, 0x7f, 0xaa, 0x7e, 0x7b, 0x01, 0x96, 0x93, 0x84, 0xd8, 0x0d, 0xc8, 0x0d, 0xc8, 0x2c, 0x2b,
		0xe9, 0xb3, 0xf8, 0x80, 0x61, 0xa2, 0x4e, 0x5c, 0xf8, 0x01, 0x9d, 0xd8, 0x80, 0xa2, 0x85, 0x3d,
		0x1f, 0xf7, 0x59, 0x44, 0x64, 0x12, 0xc6, 0x14, 0x30, 0xd0, 0x6c, 0x48, 0x65, 0x7f, 0xa0, 0x90,
		0x7a, 0x09, 0x96, 0x03, 0x93, 0x54, 0x57, 0xb3, 0x86, 0x22, 0x36, 0x2f, 0xc7, 0x59, 0x52, 0x93,
		0x05, 0x4e, 0x21, 0x30, 0xa5, 0x8c, 0x23, 0xcf, 0x68, 0x1b, 0xc0, 0xb6, 0xb0, 0x3d, 0x50, 0xfb,
		0x58, 0x37, 0x2b, 0xf9, 0x13, 0xbc, 0xd4, 0x26, 0x2a, 0x33, 0x5e, 0xb2, 0x99, 0x54, 0x37, 0xd1,
		0xf5, 0x49, 0xa8, 0x2d, 0x9e, 0x10, 0x29, 0x7b, 0x6c, 0x93, 0xcd, 0x44, 0xdb, 0x01, 0x94, 0x5d,
		0x4c, 0xe2, 0x1e, 0xf7, 0xf9, 0xcc, 0x0a, 0xd4, 0x88, 0x5a, 0xec, 0xcc, 0x14, 0x0e, 0x63, 0x13,
		0x5b, 0x72, 0xc3, 0x8f, 0xe8, 0x61, 0x08, 0x04, 0x2a, 0x0d, 0x2b, 0xa0, 0x59, 0xa8, 0x24, 0x84,
		0xfb, 0xda, 0x08, 0xaf, 0xdf, 0x83, 0x72, 0xd4, 0x3d, 0x68, 0x0d, 0x72, 0x9e, 0xaf, 0xb9, 0x3e,
		0x8d, 0xc2, 0x9c, 0xc2, 0x1e, 0x90, 0x04, 0x19, 0x6c, 0xf5, 0x69, 0x96, 0xcb, 0x29, 0xe4, 0x5f,
		0xf4, 0x23, 0x93, 0x09, 0x67, 0xe8, 0x84, 0x3f, 0x3e, 0xbb, 0xa2, 0x11, 0xe6, 0xe9, 0x79, 0xaf,
		0x3f, 0x0b, 0x4b, 0x91, 0x09, 0x24, 0x7d, 0x75, 0xf5, 0xc7, 0xe1, 0xfe, 0xb9, 0xd4, 0xe8, 0x25,
		0x58, 0x1b, 0x5b, 0x86, 0xe5, 0x63, 0xd7, 0x71, 0x31, 0x89, 0x58, 0xf6, 0xaa, 0xca, 0x3f, 0x2f,
		0x9e, 0x10, 0x73, 0x07, 0x61, 0x6d, 0xc6, 0xa2, 0xac, 0x8e, 0x67, 0x85, 0x4f, 0x14, 0xf2, 0xff,
		0xb2, 0x28, 0xbd, 0xf6, 0xda, 0x6b, 0xaf, 0xa5, 0xab, 0x5f, 0x5c, 0x80, 0xb5, 0x79, 0x7b, 0x66,
		0xee, 0xf6, 0x3d, 0x07, 0x0b, 0xd6, 0x78, 0x74, 0x88, 0x5d, 0xea, 0xa4, 0x9c, 0xc2, 0x9f, 0x50,
		0x03, 0x72, 0xa6, 0x76, 0x88, 0xcd, 0x4a, 0x76, 0x23, 0xb5, 0x59, 0xbe, 0xf2, 0x64, 0xa2, 0x5d,
		0x59, 0xdb, 0x25, 0x10, 0x85, 0x21, 0xd1, 0xa7, 0x20, 0xcb, 0x53, 0x34, 0x61, 0x78, 0x22, 0x19,
		0x03, 0xd9, 0x4b, 0x0a, 0xc5, 0xa1, 0x07, 0xa0, 0x40, 0xfe, 0xb2, 0xd8, 0x58, 0xa0, 0x36, 0xe7,
		0x89, 0x80, 0xc4, 0x05, 0x5a, 0x87, 0x3c, 0xdd, 0x26, 0x7d, 0x2c, 0x4a, 0x5b, 0xf0, 0x4c, 0x02,
		0xab, 0x8f, 0x07, 0xda, 0xd8, 0xf4, 0xd5, 0x3b, 0x9a, 0x39, 0xc6, 0x34, 0xe0, 0x0b, 0x4a, 0x89,
		0x0b, 0x3f, 0x4b, 0x64, 0xe8, 0x22, 0x14, 0xd9, 0xae, 0x32, 0xac, 0x3e, 0x7e, 0x95, 0x66, 0xcf,
		0x9c, 0xc2, 0x36, 0x5a, 0x8b, 0x48, 0xc8, 0xeb, 0x6f, 0x79, 0xb6, 0x25, 0x42, 0x93, 0xbe, 0x82,
		0x08, 0xe8, 0xeb, 0x9f, 0x9d, 0x4e, 0xdc, 0x0f, 0xce, 0x9f, 0xde, 0x74, 0x4c, 0x55, 0xbf, 0x95,
		0x86, 0x2c, 0xcd, 0x17, 0xcb, 0x50, 0xec, 0xbd, 0xdc, 0x91, 0xd5, 0xed, 0xf6, 0xc1, 0xd6, 0xae,
		0x2c, 0xa5, 0x50, 0x19, 0x80, 0x0a, 0x6e, 0xee, 0xb6, 0x1b, 0x3d, 0x29, 0x1d, 0x3c, 0xb7, 0xf6,
		0x7b, 0xd7, 0x9e, 0x91, 0x32, 0x01, 0xe0, 0x80, 0x09, 0xb2, 0x61, 0x85, 0xa7, 0xaf, 0x48, 0x39,
		0x24, 0x41, 0x89, 0x11, 0xb4, 0x5e, 0x92, 0xb7, 0xaf, 0x3d, 0x23, 0x2d, 0x44, 0x25, 0x4f, 0x5f,
		0x91, 0x16, 0xd1, 0x12, 0x14, 0xa8, 0x64, 0xab, 0xdd, 0xde, 0x95, 0xf2, 0x01, 0x67, 0xb7, 0xa7,
		0xb4, 0xf6, 0x77, 0xa4, 0x42, 0xc0, 0xb9, 0xa3, 0xb4, 0x0f, 0x3a, 0x12, 0x04, 0x0c, 0x7b, 0x72,
		0xb7, 0xdb, 0xd8, 0x91, 0xa5, 0x62, 0xa0, 0xb1, 0xf5, 0x72, 0x4f, 0xee, 0x4a, 0xa5, 0x88, 0x59,
		0x4f, 0x5f, 0x91, 0x96, 0x82, 0x57, 0xc8, 0xfb, 0x07, 0x7b, 0x52, 0x19, 0xad, 0xc0, 0x12, 0x7b,
		0x85, 0x30, 0x62, 0x79, 0x4a, 0x74, 0xed, 0x19, 0x49, 0x9a, 0x18, 0xc2, 0x58, 0x56, 0x22, 0x82,
		0x6b, 0xcf, 0x48, 0xa8, 0xda, 0x84, 0x1c, 0x8d, 0x2e, 0x84, 0xa0, 0xbc, 0xdb, 0xd8, 0x92, 0x77,
		0xd5, 0x76, 0xa7, 0xd7, 0x6a, 0xef, 0x37, 0x76, 0xa5, 0xd4, 0x44, 0xa6, 0xc8, 0x9f, 0x39, 0x68,
		0x29, 0xf2, 0xb6, 0x94, 0x0e, 0xcb, 0x3a, 0x72, 0xa3, 0x27, 0x6f, 0x4b, 0x99, 0xaa, 0x0e, 0x6b,
		0xf3, 0xf2, 0xe4, 0xdc, 0x9d, 0x11, 0x5a, 0xe2, 0xf4, 0x09, 0x4b, 0x4c, 0xb9, 0x66, 0x96, 0xf8,
		0x9f, 0xd2, 0xb0, 0x3a, 0xa7, 0x56, 0xcc, 0x7d, 0xc9, 0xf3, 0x90, 0x63, 0x21, 0xca, 0xaa, 0xe7,
		0xe3, 0x73, 0x8b, 0x0e, 0x0d, 0xd8, 0x99, 0x0a, 0x4a, 0x71, 0xe1, 0x0e, 0x22, 0x73, 0x42, 0x07,
		0x41, 0x28, 0x66, 0x72, 0xfa, 0x8f, 0xcd, 0xe4, 0x74, 0x56, 0xf6, 0xae, 0x25, 0x29, 0x7b, 0x54,
		0x76, 0xb6, 0xdc, 0x9e, 0x9b, 0x93, 0xdb, 0x6f, 0xc0, 0xca, 0x0c, 0x51, 0xe2, 0x1c, 0xfb, 0x7a,
		0x0a, 0x2a, 0x27, 0x39, 0x27, 0x26, 0xd3, 0xa5, 0x23, 0x99, 0xee, 0xc6, 0xb4, 0x07, 0x2f, 0x9d,
		0xbc, 0x08, 0x33, 0x6b, 0xfd, 0x76, 0x0a, 0xce, 0xcd, 0xef, 0x14, 0xe7, 0xda, 0xf0, 0x29, 0x58,
		0x18, 0x61, 0xff, 0xc8, 0x16, 0xdd, 0xd2, 0xc7, 0xe7, 0xd4, 0x60, 0x32, 0x3c, 0xbd, 0xd8, 0x1c,
		0x15, 0x2e, 0xe2, 0x99, 0x93, 0xda, 0x3d, 0x66, 0xcd, 0x8c, 0xa5, 0x5f, 0x48, 0xc3, 0xfd, 0x73,
		0xc9, 0xe7, 0x1a, 0xfa, 0x20, 0x80, 0x61, 0x39, 0x63, 0x9f, 0x75, 0x44, 0x2c, 0xc1, 0x16, 0xa8,
		0x84, 0x26, 0x2f, 0x92, 0x3c, 0xc7, 0x7e, 0x30, 0x9e, 0xa1, 0xe3, 0xc0, 0x44, 0x54, 0xe1, 0xb9,
		0x89, 0xa1, 0x59, 0x6a, 0xe8, 0x43, 0x27, 0xcc, 0x74, 0x26, 0x30, 0x3f, 0x01, 0x92, 0x6e, 0x1a,
		0xd8, 0xf2, 0x55, 0xcf, 0x77, 0xb1, 0x36, 0x32, 0xac, 0x21, 0xad, 0x20, 0xf9, 0x7a, 0x6e, 0xa0,
		0x99, 0x1e, 0x56, 0x96, 0xd9, 0x70, 0x57, 0x8c, 0x12, 0x04, 0x0d, 0x20, 0x37, 0x84, 0x58, 0x88,
		0x20, 0xd8, 0x70, 0x80, 0xa8, 0x7e, 0x33, 0x0f, 0xc5, 0x50, 0x5f, 0x8d, 0x2e, 0x41, 0xe9, 0x96,
		0x76, 0x47, 0x53, 0xc5, 0x59, 0x89, 0x79, 0xa2, 0x48, 0x64, 0x1d, 0x7e, 0x5e, 0xfa, 0x04, 0xac,
		0x51, 0x15, 0x7b, 0xec, 0x63, 0x57, 0xd5, 0x4d, 0xcd, 0xf3, 0xa8, 0xd3, 0xf2, 0x54, 0x15, 0x91,
		0xb1, 0x36, 0x19, 0x6a, 0x8a, 0x11, 0x74, 0x15, 0x56, 0x29, 0x62, 0x34, 0x36, 0x7d, 0xc3, 0x31,
		0xb1, 0x4a, 0x4e, 0x6f, 0x1e, 0xad, 0x24, 0x81, 0x65, 0x2b, 0x44, 0x63, 0x8f, 0x2b, 0x10, 0x8b,
		0x3c, 0xb4, 0x0d, 0x0f, 0x52, 0xd8, 0x10, 0x5b, 0xd8, 0xd5, 0x7c, 0xac, 0xe2, 0xcf, 0x8f, 0x35,
		0xd3, 0x53, 0x35, 0xab, 0xaf, 0x1e, 0x69, 0xde, 0x51, 0x65, 0x8d, 0x10, 0x6c, 0xa5, 0x2b, 0x29,
		0xe5, 0x02, 0x51, 0xdc, 0xe1, 0x7a, 0x32, 0x55, 0x6b, 0x58, 0xfd, 0x4f, 0x6b, 0xde, 0x11, 0xaa,
		0xc3, 0x39, 0xca, 0xe2, 0xf9, 0xae, 0x61, 0x0d, 0x55, 0xfd, 0x08, 0xeb, 0xb7, 0xd5, 0xb1, 0x3f,
		0x78, 0xae, 0xf2, 0x40, 0xf8, 0xfd, 0xd4, 0xc2, 0x2e, 0xd5, 0x69, 0x12, 0x95, 0x03, 0x7f, 0xf0,
		0x1c, 0xea, 0x42, 0x89, 0x2c, 0xc6, 0xc8, 0xb8, 0x87, 0xd5, 0x81, 0xed, 0xd2, 0xd2, 0x58, 0x9e,
		0x93, 0x9a, 0x42, 0x1e, 0xac, 0xb5, 0x39, 0x60, 0xcf, 0xee, 0xe3, 0x7a, 0xae, 0xdb, 0x91, 0xe5,
		0x6d, 0xa5, 0x28, 0x58, 0x6e, 0xda, 0x2e, 0x09, 0xa8, 0xa1, 0x1d, 0x38, 0xb8, 0xc8, 0x02, 0x6a,
		0x68, 0x0b, 0xf7, 0x5e, 0x85, 0x55, 0x5d, 0x67, 0x73, 0x36, 0x74, 0x95, 0x9f, 0xb1, 0xbc, 0x8a,
		0x14, 0x71, 0x96, 0xae, 0xef, 0x30, 0x05, 0x1e, 0xe3, 0x1e, 0xba, 0x0e, 0xf7, 0x4f, 0x9c, 0x15,
		0x06, 0xae, 0xcc, 0xcc, 0x72, 0x1a, 0x7a, 0x15, 0x56, 0x9d, 0xe3, 0x59, 0x20, 0x8a, 0xbc, 0xd1,
		0x39, 0x9e, 0x86, 0x3d, 0x0b, 0x6b, 0xce, 0x91, 0x33, 0x8b, 0x7b, 0x22, 0x8c, 0x43, 0xce, 0x91,
		0x33, 0x0d, 0x7c, 0x94, 0x1e, 0xb8, 0x5d, 0xac, 0x6b, 0x3e, 0xee, 0x57, 0xce, 0x87, 0xd5, 0x43,
		0x03, 0xe8, 0x32, 0x48, 0xba, 0xae, 0x62, 0x4b, 0x3b, 0x34, 0xb1, 0xaa, 0xb9, 0xd8, 0xd2, 0xbc,
		0xca, 0xc5, 0xb0, 0x72, 0x59, 0xd7, 0x65, 0x3a, 0xda, 0xa0, 0x83, 0xe8, 0x09, 0x58, 0xb1, 0x0f,
		0x6f, 0xe9, 0x2c, 0x24, 0x55, 0xc7, 0xc5, 0x03, 0xe3, 0xd5, 0xca, 0x23, 0xd4, 0xbf, 0xcb, 0x64,
		0x80, 0x06, 0x64, 0x87, 0x8a, 0xd1, 0xe3, 0x20, 0xe9, 0xde, 0x91, 0xe6, 0x3a, 0x34, 0x27, 0x7b,
		0x8e, 0xa6, 0xe3, 0xca, 0xa3, 0x4c, 0x95, 0xc9, 0xf7, 0x85, 0x98, 0x6c, 0x09, 0xef, 0xae, 0x31,
		0xf0, 0x05, 0xe3, 0x63, 0x6c, 0x4b, 0x50, 0x19, 0x67, 0xdb, 0x04, 0x89, 0xb8, 0x22, 0xf2, 0xe2,
		0x4d, 0xaa, 0x56, 0x76, 0x8e, 0x9c, 0xf0, 0x7b, 0x1f, 0x86, 0x25, 0xa2, 0x39, 0x79, 0xe9, 0xe3,
		0xac, 0x21, 0x73, 0x8e, 0x42, 0x6f, 0xfc, 0xc0, 0x7a, 0xe3, 0x6a, 0x1d, 0x4a, 0xe1, 0xf8, 0x44,
		0x05, 0x60, 0x11, 0x2a, 0xa5, 0x48, 0xb3, 0xd2, 0x6c, 0x6f, 0x93, 0x36, 0xe3, 0x15, 0x59, 0x4a,
		0x93, 0x76, 0x67, 0xb7, 0xd5, 0x93, 0x55, 0xe5, 0x60, 0xbf, 0xd7, 0xda, 0x93, 0xa5, 0x4c, 0xb8,
		0xaf, 0xfe, 0x4e, 0x1a, 0xca, 0xd1, 0x23, 0x12, 0xfa, 0x61, 0x38, 0x2f, 0xee, 0x33, 0x3c, 0xec,
		0xab, 0x77, 0x0d, 0x97, 0x6e, 0x99, 0x91, 0xc6, 0xca, 0x57, 0xb0, 0x68, 0x6b, 0x5c, 0xab, 0x8b,
		0xfd, 0x17, 0x0d, 0x97, 0x6c, 0x88, 0x91, 0xe6, 0xa3, 0x5d, 0xb8, 0x68, 0xd9, 0xaa, 0xe7, 0x6b,
		0x56, 0x5f, 0x73, 0xfb, 0xea, 0xe4, 0x26, 0x49, 0xd5, 0x74, 0x1d, 0x7b, 0x9e, 0xcd, 0x4a, 0x55,
		0xc0, 0xf2, 0x31, 0xcb, 0xee, 0x72, 0xe5, 0x49, 0x0e, 0x6f, 0x70, 0xd5, 0xa9, 0x00, 0xcb, 0x9c,
		0x14, 0x60, 0x0f, 0x40, 0x61, 0xa4, 0x39, 0x2a, 0xb6, 0x7c, 0xf7, 0x98, 0x36, 0xc6, 0x79, 0x25,
		0x3f, 0xd2, 0x1c, 0x99, 0x3c, 0x7f, 0x38, 0xe7, 0x93, 0xbf, 0xcb, 0x40, 0x29, 0xdc, 0x1c, 0x93,
		0xb3, 0x86, 0x4e, 0xeb, 0x48, 0x8a, 0x66, 0x9a, 0x87, 0x4f, 0x6d, 0xa5, 0x6b, 0x4d, 0x52, 0x60,
		0xea, 0x0b, 0xac, 0x65, 0x55, 0x18, 0x92, 0x14, 0x77, 0x92, 0x5b, 0x30, 0x6b, 0x11, 0xf2, 0x0a,
		0x7f, 0x42, 0x3b, 0xb0, 0x70, 0xcb, 0xa3, 0xdc, 0x0b, 0x94, 0xfb, 0x91, 0xd3, 0xb9, 0x5f, 0xe8,
		0x52, 0xf2, 0xc2, 0x0b, 0x5d, 0x75, 0xbf, 0xad, 0xec, 0x35, 0x76, 0x15, 0x0e, 0x47, 0x17, 0x20,
		0x6b, 0x6a, 0xf7, 0x8e, 0xa3, 0xa5, 0x88, 0x8a, 0x92, 0x3a, 0xfe, 0x02, 0x64, 0xef, 0x62, 0xed,
		0x76, 0xb4, 0x00, 0x50, 0xd1, 0x07, 0x18, 0xfa, 0x97, 0x21, 0x47, 0xfd, 0x85, 0x00, 0xb8, 0xc7,
		0xa4, 0xfb, 0x50, 0x1e, 0xb2, 0xcd, 0xb6, 0x42, 0xc2, 0x5f, 0x82, 0x12, 0x93, 0xaa, 0x9d, 0x96,
		0xdc, 0x94, 0xa5, 0x74, 0xf5, 0x2a, 0x2c, 0x30, 0x27, 0x90, 0xad, 0x11, 0xb8, 0x41, 0xba, 0x8f,
		0x3f, 0x72, 0x8e, 0x94, 0x18, 0x3d, 0xd8, 0xdb, 0x92, 0x15, 0x29, 0x1d, 0x5e, 0x5e, 0x0f, 0x4a,
		0xe1, 0xbe, 0xf8, 0xc3, 0x89, 0xa9, 0xbf, 0x48, 0x41, 0x31, 0xd4, 0xe7, 0x92, 0x06, 0x45, 0x33,
		0x4d, 0xfb, 0xae, 0xaa, 0x99, 0x86, 0xe6, 0xf1, 0xa0, 0x00, 0x2a, 0x6a, 0x10, 0x49, 0xd2, 0x45,
		0xfb, 0x50, 0x8c, 0x7f, 0x2b, 0x05, 0xd2, 0x74, 0x8b, 0x39, 0x65, 0x60, 0xea, 0x23, 0x35, 0xf0,
		0xcd, 0x14, 0x94, 0xa3, 0x7d, 0xe5, 0x94, 0x79, 0x97, 0x3e, 0x52, 0xf3, 0xfe, 0x3e, 0x0d, 0x4b,
		0x91, 0x6e, 0x32, 0xa9, 0x75, 0x9f, 0x87, 0x15, 0xa3, 0x8f, 0x47, 0x8e, 0xed, 0x63, 0x4b, 0x3f,
		0x56, 0x4d, 0x7c, 0x07, 0x9b, 0x95, 0x2a, 0x4d, 0x14, 0x97, 0x4f, 0xef, 0x57, 0x6b, 0xad, 0x09,
		0x6e, 0x97, 0xc0, 0xea, 0xab, 0xad, 0x6d, 0x79, 0xaf, 0xd3, 0xee, 0xc9, 0xfb, 0xcd, 0x97, 0xd5,
		0x83, 0xfd, 0x1f, 0xdd, 0x6f, 0xbf, 0xb8, 0xaf, 0x48, 0xc6, 0x94, 0xda, 0x07, 0xb8, 0xd5, 0x3b,
		0x20, 0x4d, 0x1b, 0x85, 0xce, 0xc3, 0x3c, 0xb3, 0xa4, 0xfb, 0xd0, 0x2a, 0x2c, 0xef, 0xb7, 0xd5,
		0x6e, 0x6b, 0x5b, 0x56, 0xe5, 0x9b, 0x37, 0xe5, 0x66, 0xaf, 0xcb, 0x6e, 0x20, 0x02, 0xed, 0x5e,
		0x74, 0x53, 0xbf, 0x91, 0x81, 0xd5, 0x39, 0x96, 0xa0, 0x06, 0x3f, 0x3b, 0xb0, 0xe3, 0xcc, 0x53,
		0x49, 0xac, 0xaf, 0x91, 0x92, 0xdf, 0xd1, 0x5c, 0x9f, 0x1f, 0x35, 0x1e, 0x07, 0xe2, 0x25, 0xcb,
		0x37, 0x06, 0x06, 0x76, 0xf9, 0x85, 0x0d, 0x3b, 0x50, 0x2c, 0x4f, 0xe4, 0xec, 0xce, 0xe6, 0x87,
		0x00, 0x39, 0xb6, 0x67, 0xf8, 0xc6, 0x1d, 0xac, 0x1a, 0x96, 0xb8, 0xdd, 0x21, 0x07, 0x8c, 0xac,
		0x22, 0x89, 0x91, 0x96, 0xe5, 0x07, 0xda, 0x16, 0x1e, 0x6a, 0x53, 0xda, 0x24, 0x81, 0x67, 0x14,
		0x49, 0x8c, 0x04, 0xda, 0x97, 0xa0, 0xd4, 0xb7, 0xc7, 0xa4, 0xeb, 0x62, 0x7a, 0xa4, 0x5e, 0xa4,
		0x94, 0x22, 0x93, 0x05, 0x2a, 0xbc, 0x9f, 0x9e, 0x5c, 0x2b, 0x95, 0x94, 0x22, 0x93, 0x31, 0x95,
		0xc7, 0x60, 0x59, 0x1b, 0x0e, 0x5d, 0x42, 0x2e, 0x88, 0xd8, 0x09, 0xa1, 0x1c, 0x88, 0xa9, 0xe2,
		0xfa, 0x0b, 0x90, 0x17, 0x7e, 0x20, 0x25, 0x99, 0x78, 0x42, 0x75, 0xd8, 0xb1, 0x37, 0xbd, 0x59,
		0x50, 0xf2, 0x96, 0x18, 0xbc, 0x04, 0x25, 0xc3, 0x53, 0x27, 0xb7, 0xe4, 0xe9, 0x8d, 0xf4, 0x66,
		0x5e, 0x29, 0x1a, 0x5e, 0x70, 0xc3, 0x58, 0x7d, 0x3b, 0x0d, 0xe5, 0xe8, 0x2d, 0x3f, 0xda, 0x86,
		0xbc, 0x69, 0xeb, 0x1a, 0x0d, 0x2d, 0xf6, 0x89, 0x69, 0x33, 0xe6, 0xc3, 0x40, 0x6d, 0x97, 0xeb,
		0x2b, 0x01, 0x72, 0xfd, 0x6f, 0x53, 0x90, 0x17, 0x62, 0x74, 0x0e, 0xb2, 0x8e, 0xe6, 0x1f, 0x51,
		0xba, 0xdc, 0x56, 0x5a, 0x4a, 0x29, 0xf4, 0x99, 0xc8, 0x3d, 0x47, 0xb3, 0x68, 0x08, 0x70, 0x39,
		0x79, 0x26, 0xeb, 0x6a, 0x62, 0xad, 0x4f, 0x8f, 0x1f, 0xf6, 0x68, 0x84, 0x2d, 0xdf, 0x13, 0xeb,
		0xca, 0xe5, 0x4d, 0x2e, 0x46, 0x4f, 0xc2, 0x8a, 0xef, 0x6a, 0x86, 0x19, 0xd1, 0xcd, 0x52, 0x5d,
		0x49, 0x0c, 0x04, 0xca, 0x75, 0xb8, 0x20, 0x78, 0xfb, 0xd8, 0xd7, 0xf4, 0x23, 0xdc, 0x9f, 0x80,
		0x16, 0xe8, 0x35, 0xc3, 0x79, 0xae, 0xb0, 0xcd, 0xc7, 0x05, 0xb6, 0xfa, 0xdd, 0x14, 0xac, 0x88,
		0x03, 0x53, 0x3f, 0x70, 0xd6, 0x1e, 0x80, 0x66, 0x59, 0xb6, 0x1f, 0x76, 0xd7, 0x6c, 0x28, 0xcf,
		0xe0, 0x6a, 0x8d, 0x00, 0xa4, 0x84, 0x08, 0xd6, 0x47, 0x00, 0x93, 0x91, 0x13, 0xdd, 0x76, 0x11,
		0x8a, 0xfc, 0x13, 0x0e, 0xfd, 0x0e, 0xc8, 0x8e, 0xd8, 0xc0, 0x44, 0xe4, 0x64, 0x85, 0xd6, 0x20,
		0x77, 0x88, 0x87, 0x86, 0xc5, 0x2f, 0x66, 0xd9, 0x83, 0xb8, 0x08, 0xc9, 0x06, 0x17, 0x21, 0x5b,
		0x9f, 0x83, 0x55, 0xdd, 0x1e, 0x4d, 0x9b, 0xbb, 0x25, 0x4d, 0x1d, 0xf3, 0xbd, 0x4f, 0xa7, 0x5e,
		0x81, 0x49, 0x8b, 0xf9, 0xbf, 0xa9, 0xd4, 0x6f, 0xa5, 0x33, 0x3b, 0x9d, 0xad, 0xaf, 0xa5, 0xd7,
		0x77, 0x18, 0xb4, 0x23, 0x66, 0xaa, 0xe0, 0x81, 0x89, 0x75, 0x62, 0x3d, 0x7c, 0x75, 0x13, 0x9e,
		0x1a, 0x1a, 0xfe, 0xd1, 0xf8, 0xb0, 0xa6, 0xdb, 0xa3, 0xcb, 0x43, 0x7b, 0x68, 0x4f, 0x3e, 0x7d,
		0x92, 0x27, 0xfa, 0x40, 0xff, 0xe3, 0x9f, 0x3f, 0x0b, 0x81, 0x74, 0x3d, 0xf6, 0x5b, 0x69, 0x7d,
		0x1f, 0x56, 0xb9, 0xb2, 0x4a, 0xbf, 0xbf, 0xb0, 0x53, 0x04, 0x3a, 0xf5, 0x0e, 0xab, 0xf2, 0x8d,
		0x77, 0x68, 0xb9, 0x56, 0x56, 0x38, 0x94, 0x8c, 0xb1, 0x83, 0x46, 0x5d, 0x81, 0xfb, 0x23, 0x7c,
		0x6c, 0x6b, 0x62, 0x37, 0x86, 0xf1, 0x3b, 0x9c, 0x71, 0x35, 0xc4, 0xd8, 0xe5, 0xd0, 0x7a, 0x13,
		0x96, 0xce, 0xc2, 0xf5, 0x57, 0x9c, 0xab, 0x84, 0xc3, 0x24, 0x3b, 0xb0, 0x4c, 0x49, 0xf4, 0xb1,
		0xe7, 0xdb, 0x23, 0x9a, 0xf7, 0x4e, 0xa7, 0xf9, 0xeb, 0x77, 0xd8, 0x5e, 0x29, 0x13, 0x58, 0x33,
		0x40, 0xd5, 0xeb, 0x40, 0x3f, 0x39, 0xf5, 0xb1, 0x6e, 0xc6, 0x30, 0xfc, 0x0d, 0x37, 0x24, 0xd0,
		0xaf, 0x7f, 0x16, 0xd6, 0xc8, 0xff, 0x34, 0x2d, 0x85, 0x2d, 0x89, 0xbf, 0xf0, 0xaa, 0x7c, 0xf7,
		0x75, 0xb6, 0x1d, 0x57, 0x03, 0x82, 0x90, 0x4d, 0xa1, 0x55, 0x1c, 0x62, 0xdf, 0xc7, 0xae, 0xa7,
		0x6a, 0xe6, 0x3c, 0xf3, 0x42, 0x37, 0x06, 0x95, 0x2f, 0xbd, 0x1b, 0x5d, 0xc5, 0x1d, 0x86, 0x6c,
		0x98, 0x66, 0xfd, 0x00, 0xce, 0xcf, 0x89, 0x8a, 0x04, 0x9c, 0x6f, 0x70, 0xce, 0xb5, 0x99, 0xc8,
		0x20, 0xb4, 0x1d, 0x10, 0xf2, 0x60, 0x2d, 0x13, 0x70, 0xfe, 0x3a, 0xe7, 0x44, 0x1c, 0x2b, 0x96,
		0x94, 0x30, 0xbe, 0x00, 0x2b, 0x77, 0xb0, 0x7b, 0x68, 0x7b, 0xfc, 0x96, 0x26, 0x01, 0xdd, 0x9b,
		0x9c, 0x6e, 0x99, 0x03, 0xe9, 0xb5, 0x0d, 0xe1, 0xba, 0x0e, 0xf9, 0x81, 0xa6, 0xe3, 0x04, 0x14,
		0x5f, 0xe6, 0x14, 0x8b, 0x44, 0x9f, 0x40, 0x1b, 0x50, 0x1a, 0xda, 0xbc, 0x32, 0xc5, 0xc3, 0xdf,
		0xe2, 0xf0, 0xa2, 0xc0, 0x70, 0x0a, 0xc7, 0x76, 0xc6, 0x26, 0x29, 0x5b, 0xf1, 0x14, 0xbf, 0x21,
		0x28, 0x04, 0x86, 0x53, 0x9c, 0xc1, 0xad, 0xbf, 0x29, 0x28, 0xbc, 0x90, 0x3f, 0x9f, 0x87, 0xa2,
		0x6d, 0x99, 0xc7, 0xb6, 0x95, 0xc4, 0x88, 0xaf, 0x70, 0x06, 0xe0, 0x10, 0x42, 0x70, 0x03, 0x0a,
		0x49, 0x17, 0xe2, 0xab, 0xef, 0x8a, 0xed, 0x21, 0x56, 0x60, 0x07, 0x96, 0x45, 0x82, 0x32, 0x6c,
		0x2b, 0x01, 0xc5, 0x6f, 0x73, 0x8a, 0x72, 0x08, 0xc6, 0xa7, 0xe1, 0x63, 0xcf, 0x1f, 0xe2, 0x24,
		0x24, 0x6f, 0x8b, 0x69, 0x70, 0x08, 0x77, 0xe5, 0x21, 0xb6, 0xf4, 0xa3, 0x64, 0x0c, 0xbf, 0x23,
		0x5c, 0x29, 0x30, 0x84, 0xa2, 0x09, 0x4b, 0x23, 0xcd, 0xf5, 0x8e, 0x34, 0x33, 0xd1, 0x72, 0xfc,
		0x2e, 0xe7, 0x28, 0x05, 0x20, 0xee, 0x91, 0xb1, 0x75, 0x16, 0x9a, 0xaf, 0x09, 0x8f, 0x84, 0x60,
		0x7c, 0xeb, 0x79, 0x3e, 0xbd, 0xd2, 0x3a, 0x0b, 0xdb, 0xef, 0x89, 0xad, 0xc7, 0xb0, 0x7b, 0x61,
		0xc6, 0x1b, 0x50, 0xf0, 0x8c, 0x7b, 0x89, 0x68, 0x7e, 0x5f, 0xac, 0x34, 0x05, 0x10, 0xf0, 0xcb,
		0x70, 0x61, 0x6e, 0x99, 0x48, 0x40, 0xf6, 0x07, 0x9c, 0xec, 0xdc, 0x9c, 0x52, 0xc1, 0x53, 0xc2,
		0x59, 0x29, 0xff, 0x50, 0xa4, 0x04, 0x3c, 0xc5, 0xd5, 0x21, 0x67, 0x05, 0x4f, 0x1b, 0x9c, 0xcd,
		0x6b, 0x7f, 0x24, 0xbc, 0xc6, 0xb0, 0x11, 0xaf, 0xf5, 0xe0, 0x1c, 0x67, 0x3c, 0xdb, 0xba, 0x7e,
		0x5d, 0x24, 0x56, 0x86, 0x3e, 0x88, 0xae, 0xee, 0xe7, 0x60, 0x3d, 0x70, 0xa7, 0x68, 0x4a, 0x3d,
		0x75, 0xa4, 0x39, 0x09, 0x98, 0xbf, 0xc1, 0x99, 0x45, 0xc6, 0x0f, 0xba, 0x5a, 0x6f, 0x4f, 0x73,
		0x08, 0xf9, 0x4b, 0x50, 0x11, 0xe4, 0x63, 0xcb, 0xc5, 0xba, 0x3d, 0xb4, 0x8c, 0x7b, 0xb8, 0x9f,
		0x80, 0xfa, 0x8f, 0xa7, 0x96, 0xea, 0x20, 0x04, 0x27, 0xcc, 0x2d, 0x90, 0x82, 0x5e, 0x45, 0x35,
		0x46, 0x8e, 0xed, 0xfa, 0x31, 0x8c, 0xdf, 0x14, 0x2b, 0x15, 0xe0, 0x5a, 0x14, 0x56, 0x97, 0xa1,
		0x4c, 0x1f, 0x93, 0x86, 0xe4, 0x9f, 0x70, 0xa2, 0xa5, 0x09, 0x8a, 0x27, 0x0e, 0xdd, 0x1e, 0x39,
		0x9a, 0x9b, 0x24, 0xff, 0xfd, 0xa9, 0x48, 0x1c, 0x1c, 0xc2, 0x13, 0x87, 0x7f, 0xec, 0x60, 0x52,
		0xed, 0x13, 0x30, 0x7c, 0x4b, 0x24, 0x0e, 0x81, 0xe1, 0x14, 0xa2, 0x61, 0x48, 0x40, 0xf1, 0x67,
		0x82, 0x42, 0x60, 0x08, 0xc5, 0x67, 0x26, 0x85, 0xd6, 0xc5, 0x43, 0xc3, 0xf3, 0x5d, 0xd6, 0x0a,
		0x9f, 0x4e, 0xf5, 0xe7, 0xef, 0x46, 0x9b, 0x30, 0x25, 0x04, 0x25, 0x99, 0x88, 0x5f, 0xa1, 0xd2,
		0x93, 0x52, 0xbc, 0x61, 0xdf, 0x16, 0x99, 0x28, 0x04, 0x63, 0xfb, 0x73, 0x79, 0xaa, 0x57, 0x41,
		0x71, 0x3f, 0x84, 0xa9, 0xfc, 0xc4, 0x7b, 0x9c, 0x2b, 0xda, 0xaa, 0xd4, 0x77, 0x49, 0x00, 0x45,
		0x1b, 0x8a, 0x78, 0xb2, 0xd7, 0xdf, 0x0b, 0x62, 0x28, 0xd2, 0x4f, 0xd4, 0x6f, 0xc2, 0x52, 0xa4,
		0x99, 0x88, 0xa7, 0xfa, 0x49, 0x4e, 0x55, 0x0a, 0xf7, 0x12, 0xf5, 0xab, 0x90, 0x25, 0x8d, 0x41,
		0x3c, 0xfc, 0xa7, 0x38, 0x9c, 0xaa, 0xd7, 0x3f, 0x09, 0x79, 0xd1, 0x10, 0xc4, 0x43, 0x7f, 0x9a,
		0x43, 0x03, 0x08, 0x81, 0x8b, 0x66, 0x20, 0x1e, 0xfe, 0x33, 0x02, 0x2e, 0x20, 0x04, 0x9e, 0xdc,
		0x85, 0x7f, 0xf9, 0xb3, 0x59, 0x9e, 0xd0, 0x85, 0xef, 0x6e, 0xc0, 0x22, 0xef, 0x02, 0xe2, 0xd1,
		0x5f, 0xe0, 0x2f, 0x17, 0x88, 0xfa, 0xb3, 0x90, 0x4b, 0xe8, 0xf0, 0x9f, 0xe3, 0x50, 0xa6, 0x5f,
		0x6f, 0x42, 0x31, 0x54, 0xf9, 0xe3, 0xe1, 0x3f, 0xcf, 0xe1, 0x61, 0x14, 0x31, 0x9d, 0x57, 0xfe,
		0x78, 0x82, 0x5f, 0x10, 0xa6, 0x73, 0x04, 0x71, 0x9b, 0x28, 0xfa, 0xf1, 0xe8, 0x5f, 0x14, 0x5e,
		0x17, 0x90, 0xfa, 0xf3, 0x50, 0x08, 0x12, 0x79, 0x3c, 0xfe, 0x97, 0x38, 0x7e, 0x82, 0x21, 0x1e,
		0x08, 0x15, 0x92, 0x78, 0x8a, 0x5f, 0x16, 0x1e, 0x08, 0xa1, 0xc8, 0x36, 0x9a, 0x6e, 0x0e, 0xe2,
		0x99, 0x7e, 0x45, 0x6c, 0xa3, 0xa9, 0xde, 0x80, 0xac, 0x26, 0xcd, 0xa7, 0xf1, 0x14, 0xbf, 0x2a,
		0x56, 0x93, 0xea, 0x13, 0x33, 0xa6, 0xab, 0x6d, 0x3c, 0xc7, 0xaf, 0x09, 0x33, 0xa6, 0x8a, 0x6d,
		0xbd, 0x03, 0x68, 0xb6, 0xd2, 0xc6, 0xf3, 0x7d, 0x91, 0xf3, 0xad, 0xcc, 0x14, 0xda, 0xfa, 0x8b,
		0x70, 0x6e, 0x7e, 0x95, 0x8d, 0x67, 0xfd, 0xd2, 0x7b, 0x53, 0xe7, 0xa2, 0x70, 0x91, 0xad, 0xf7,
		0x26, 0xe9, 0x3a, 0x5c, 0x61, 0xe3, 0x69, 0xdf, 0x78, 0x2f, 0x9a, 0xb1, 0xc3, 0x05, 0xb6, 0xde,
		0x00, 0x98, 0x14, 0xb7, 0x78, 0xae, 0x37, 0x39, 0x57, 0x08, 0x44, 0xb6, 0x06, 0xaf, 0x6d, 0xf1,
		0xf8, 0x2f, 0x8b, 0xad, 0xc1, 0x11, 0x64, 0x6b, 0x88, 0xb2, 0x16, 0x8f, 0x7e, 0x4b, 0x6c, 0x0d,
		0x01, 0x21, 0x91, 0x1d, 0xaa, 0x1c, 0xf1, 0x0c, 0x5f, 0x11, 0x91, 0x1d, 0x42, 0xd5, 0x6f, 0x40,
		0xde, 0x1a, 0x9b, 0x26, 0x09, 0x50, 0x74, 0xfa, 0x0f, 0xc4, 0x2a, 0xff, 0xfa, 0x3e, 0xb7, 0x40,
		0x00, 0xea, 0x57, 0x21, 0x87, 0x47, 0x87, 0xb8, 0x1f, 0x87, 0xfc, 0xb7, 0xf7, 0x45, 0x52, 0x22,
		0xda, 0xf5, 0xe7, 0x01, 0xd8, 0xd1, 0x9e, 0x7e, 0xb6, 0x8a, 0xc1, 0xfe, 0xfb, 0xfb, 0xfc, 0xa7,
		0x1b, 0x13, 0xc8, 0x84, 0x80, 0xfd, 0x10, 0xe4, 0x74, 0x82, 0x77, 0xa3, 0x04, 0x74, 0xd6, 0xd7,
		0x61, 0xf1, 0x96, 0x67, 0x5b, 0xbe, 0x36, 0x8c, 0x43, 0xff, 0x07, 0x47, 0x0b, 0x7d, 0xe2, 0xb0,
		0x91, 0xed, 0x62, 0x5f, 0x1b, 0x7a, 0x71, 0xd8, 0xff, 0xe4, 0xd8, 0x00, 0x40, 0xc0, 0xba, 0xe6,
		0xf9, 0x49, 0xe6, 0xfd, 0x3d, 0x01, 0x16, 0x00, 0x62, 0x34, 0xf9, 0xff, 0x36, 0x3e, 0x8e, 0xc3,
		0x7e, 0x5f, 0x18, 0xcd, 0xf5, 0xeb, 0x9f, 0x84, 0x02, 0xf9, 0x97, 0xfd, 0x1e, 0x2b, 0x06, 0xfc,
		0x5f, 0x1c, 0x3c, 0x41, 0x90, 0x37, 0x7b, 0x7e, 0xdf, 0x37, 0xe2, 0x9d, 0xfd, 0xdf, 0x7c, 0xa5,
		0x85, 0x7e, 0xbd, 0x01, 0x45, 0xcf, 0xef, 0xf7, 0xc7, 0xbc, 0xbf, 0x8a, 0x81, 0xff, 0xcf, 0xfb,
		0xc1, 0x91, 0x3b, 0xc0, 0x6c, 0xc9, 0xf3, 0x6f, 0x0f, 0x61, 0xc7, 0xde, 0xb1, 0xd9, 0xbd, 0xe1,
		0x2b, 0xd5, 0xf8, 0x0b, 0x40, 0xf8, 0xbf, 0x14, 0x48, 0xbe, 0xed, 0x78, 0xc7, 0xfe, 0xd8, 0xbd,
		0x73, 0xcc, 0xaf, 0x02, 0xb3, 0x23, 0xcd, 0xb0, 0xd6, 0xcf, 0x76, 0x7f, 0x58, 0xdd, 0x05, 0xe8,
		0x11, 0xa2, 0x1e, 0x21, 0x42, 0x0f, 0x42, 0xa6, 0x67, 0x3b, 0xf4, 0xf3, 0x57, 0xf1, 0x4a, 0xb1,
		0x46, 0x08, 0x6b, 0x74, 0x58, 0x21, 0x72, 0x74, 0x11, 0xb2, 0xbd, 0xb1, 0x3b, 0xe0, 0x3f, 0xd0,
		0x13, 0xe3, 0x04, 0xa9, 0xd0, 0x81, 0xea, 0x53, 0x90, 0xa3, 0xea, 0x48, 0x9a, 0x10, 0x65, 0x18,
		0x16, 0x85, 0xb0, 0x99, 0x90, 0x3a, 0x7d, 0xaf, 0x18, 0x4c, 0x4d, 0x06, 0x05, 0x45, 0x3a, 0xa0,
		0xd8, 0xca, 0x7f, 0xff, 0x1f, 0x1e, 0x4a, 0x7d, 0xfd, 0x1f, 0x1f, 0x4a, 0xfd, 0x7f, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xf5, 0xf6, 0x2f, 0x10, 0x1a, 0x32, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *TopsyTurvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.TopsyTurvy{")
	if this.Top != nil {
		s = append(s, "Top: "+fmt.Sprintf("%#v", this.Top)+",\n")
	}
	if this.Turf != nil {
		s = append(s, "Turf: "+fmt.Sprintf("%#v", this.Turf)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Topsy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Topsy{")
	if this.Top != nil {
		s = append(s, "Top: "+valueToGoStringTopsyturvy(this.Top, "int64")+",\n")
	}
	if this.Turf != nil {
		s = append(s, "Turf: "+valueToGoStringTopsyturvy(this.Turf, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Turvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Turvy{")
	if this.Turf != nil {
		s = append(s, "Turf: "+valueToGoStringTopsyturvy(this.Turf, "int64")+",\n")
	}
	if this.Top != nil {
		s = append(s, "Top: "+valueToGoStringTopsyturvy(this.Top, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTopsyturvy(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("topsyturvy.proto", fileDescriptor_topsyturvy_7bb8c7da1903a704) }

var fileDescriptor_topsyturvy_7bb8c7da1903a704 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc9, 0x2f, 0x28,
	0xae, 0x2c, 0x29, 0x2d, 0x2a, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0xa4, 0xe4, 0xc3, 0xc5, 0x15, 0x02, 0x32, 0x28, 0x04, 0x64, 0x90, 0x90, 0x2c, 0x17, 0x73, 0x48,
	0x7e, 0x81, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xb7, 0x1e, 0xc8, 0x40, 0x3d, 0xb0, 0x74,
	0x10, 0x48, 0x5c, 0x48, 0x9e, 0x8b, 0x25, 0xa4, 0xb4, 0x28, 0x4d, 0x82, 0x09, 0x45, 0x1e, 0xa4,
	0x33, 0x08, 0x2c, 0xa1, 0xa4, 0xcb, 0xc5, 0x0a, 0x56, 0x2e, 0x24, 0x80, 0x30, 0x88, 0x19, 0xa2,
	0x57, 0x08, 0x49, 0x2f, 0x33, 0x92, 0x72, 0xb0, 0xbd, 0x30, 0x49, 0x46, 0x84, 0x24, 0xcc, 0x08,
	0x26, 0xb8, 0x11, 0x4e, 0x1c, 0x1f, 0x1e, 0xca, 0x31, 0x6e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x03, 0xf9, 0xbd, 0x05, 0xfd, 0x00, 0x00, 0x00,
}
