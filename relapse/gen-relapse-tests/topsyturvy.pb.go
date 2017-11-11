// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topsyturvy.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	topsyturvy.proto

It has these top-level messages:
	TopsyTurvy
	Topsy
	Turvy
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import gzip "compress/gzip"
import bytes "bytes"
import ioutil "io/ioutil"

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
	Top              *Topsy `protobuf:"bytes,1,opt,name=Top" json:"Top,omitempty"`
	Turf             *Turvy `protobuf:"bytes,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TopsyTurvy) Reset()                    { *m = TopsyTurvy{} }
func (m *TopsyTurvy) String() string            { return proto.CompactTextString(m) }
func (*TopsyTurvy) ProtoMessage()               {}
func (*TopsyTurvy) Descriptor() ([]byte, []int) { return fileDescriptorTopsyturvy, []int{0} }

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
	Top              *int64 `protobuf:"varint,1,opt,name=Top" json:"Top,omitempty"`
	Turf             *int64 `protobuf:"varint,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Topsy) Reset()                    { *m = Topsy{} }
func (m *Topsy) String() string            { return proto.CompactTextString(m) }
func (*Topsy) ProtoMessage()               {}
func (*Topsy) Descriptor() ([]byte, []int) { return fileDescriptorTopsyturvy, []int{1} }

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
	Turf             *int64 `protobuf:"varint,1,opt,name=Turf" json:"Turf,omitempty"`
	Top              *int64 `protobuf:"varint,2,opt,name=Top" json:"Top,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Turvy) Reset()                    { *m = Turvy{} }
func (m *Turvy) String() string            { return proto.CompactTextString(m) }
func (*Turvy) ProtoMessage()               {}
func (*Turvy) Descriptor() ([]byte, []int) { return fileDescriptorTopsyturvy, []int{2} }

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
func (this *TopsyTurvy) Description() (desc *descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Topsy) Description() (desc *descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Turvy) Description() (desc *descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func TopsyturvyDescription() (desc *descriptor.FileDescriptorSet) {
	d := &descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3744 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x6b, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0x5f, 0x12, 0x79, 0x48, 0x51, 0xa3, 0x2b, 0x79, 0x97, 0x2b, 0xc7, 0xde, 0x5d, 0xda,
		0xce, 0xca, 0x76, 0xcd, 0x0d, 0xd6, 0xde, 0xb5, 0x97, 0xdb, 0xc4, 0xa5, 0xa8, 0x59, 0x85, 0x5b,
		0x49, 0x64, 0x86, 0x54, 0xfc, 0x08, 0x8a, 0xc1, 0x68, 0x78, 0x49, 0xcd, 0xee, 0x70, 0x66, 0x32,
		0x33, 0xdc, 0xb5, 0x16, 0xfd, 0xe1, 0xc2, 0x7d, 0x20, 0x28, 0xfa, 0x2e, 0x90, 0xc4, 0x75, 0xdc,
		0x07, 0xd0, 0xba, 0x4d, 0x5f, 0x49, 0xd3, 0xa6, 0x8f, 0x5f, 0xfd, 0x93, 0xb6, 0xbf, 0x0a, 0xe4,
		0x7f, 0x7f, 0xf4, 0x61, 0xa0, 0x2f, 0xb7, 0x49, 0x5b, 0xff, 0x28, 0xb0, 0x28, 0x50, 0xdc, 0xd7,
		0x70, 0x86, 0xa4, 0x34, 0xa3, 0x00, 0x76, 0x7e, 0x49, 0x73, 0xee, 0xf9, 0xbe, 0x39, 0xf7, 0xdc,
		0x73, 0xcf, 0x39, 0xf7, 0x0e, 0xe1, 0x3b, 0xd7, 0xe1, 0xc2, 0xd0, 0xb6, 0x87, 0x26, 0xbe, 0xec,
		0xb8, 0xb6, 0x6f, 0x1f, 0x8c, 0x07, 0x97, 0xfb, 0xd8, 0xd3, 0x5d, 0xc3, 0xf1, 0x6d, 0xb7, 0x46,
		0x65, 0x68, 0x99, 0x69, 0xd4, 0x84, 0x46, 0x75, 0x17, 0x56, 0x6e, 0x1a, 0x26, 0xde, 0x0a, 0x14,
		0xbb, 0xd8, 0x47, 0x2f, 0x42, 0x76, 0x60, 0x98, 0xb8, 0x92, 0xba, 0x90, 0xd9, 0x28, 0x5e, 0x79,
		0xa2, 0x36, 0x05, 0xaa, 0x45, 0x11, 0x1d, 0x22, 0x56, 0x28, 0xa2, 0xfa, 0x5e, 0x16, 0x56, 0xe7,
		0x8c, 0x22, 0x04, 0x59, 0x4b, 0x1b, 0x11, 0xc6, 0xd4, 0x46, 0x41, 0xa1, 0xff, 0xa3, 0x0a, 0x2c,
		0x3a, 0x9a, 0x7e, 0x47, 0x1b, 0xe2, 0x4a, 0x9a, 0x8a, 0xc5, 0x23, 0x7a, 0x0c, 0xa0, 0x8f, 0x1d,
		0x6c, 0xf5, 0xb1, 0xa5, 0x1f, 0x55, 0x32, 0x17, 0x32, 0x1b, 0x05, 0x25, 0x24, 0x41, 0xcf, 0xc0,
		0x8a, 0x33, 0x3e, 0x30, 0x0d, 0x5d, 0x0d, 0xa9, 0xc1, 0x85, 0xcc, 0x46, 0x4e, 0x91, 0xd8, 0xc0,
		0xd6, 0x44, 0xf9, 0x12, 0x2c, 0xdf, 0xc3, 0xda, 0x9d, 0xb0, 0x6a, 0x91, 0xaa, 0x96, 0x89, 0x38,
		0xa4, 0xd8, 0x84, 0xd2, 0x08, 0x7b, 0x9e, 0x36, 0xc4, 0xaa, 0x7f, 0xe4, 0xe0, 0x4a, 0x96, 0xce,
		0xfe, 0xc2, 0xcc, 0xec, 0xa7, 0x67, 0x5e, 0xe4, 0xa8, 0xde, 0x91, 0x83, 0x51, 0x03, 0x0a, 0xd8,
		0x1a, 0x8f, 0x18, 0x43, 0xee, 0x18, 0xff, 0xc9, 0xd6, 0x78, 0x34, 0xcd, 0x92, 0x27, 0x30, 0x4e,
		0xb1, 0xe8, 0x61, 0xf7, 0xae, 0xa1, 0xe3, 0xca, 0x02, 0x25, 0xb8, 0x34, 0x43, 0xd0, 0x65, 0xe3,
		0xd3, 0x1c, 0x02, 0x87, 0x9a, 0x50, 0xc0, 0xaf, 0xfb, 0xd8, 0xf2, 0x0c, 0xdb, 0xaa, 0x2c, 0x52,
		0x92, 0x27, 0xe7, 0xac, 0x22, 0x36, 0xfb, 0xd3, 0x14, 0x13, 0x1c, 0xba, 0x06, 0x8b, 0xb6, 0xe3,
		0x1b, 0xb6, 0xe5, 0x55, 0xf2, 0x17, 0x52, 0x1b, 0xc5, 0x2b, 0x1f, 0x9b, 0x1b, 0x08, 0x6d, 0xa6,
		0xa3, 0x08, 0x65, 0xd4, 0x02, 0xc9, 0xb3, 0xc7, 0xae, 0x8e, 0x55, 0xdd, 0xee, 0x63, 0xd5, 0xb0,
		0x06, 0x76, 0xa5, 0x40, 0x09, 0xce, 0xcf, 0x4e, 0x84, 0x2a, 0x36, 0xed, 0x3e, 0x6e, 0x59, 0x03,
		0x5b, 0x29, 0x7b, 0x91, 0x67, 0x74, 0x06, 0x16, 0xbc, 0x23, 0xcb, 0xd7, 0x5e, 0xaf, 0x94, 0x68,
		0x84, 0xf0, 0xa7, 0xea, 0x9f, 0x2f, 0xc0, 0x72, 0x92, 0x10, 0xbb, 0x01, 0xb9, 0x01, 0x99, 0x65,
		0x25, 0x7d, 0x1a, 0x1f, 0x30, 0x4c, 0xd4, 0x89, 0x0b, 0xdf, 0xa3, 0x13, 0x1b, 0x50, 0xb4, 0xb0,
		0xe7, 0xe3, 0x3e, 0x8b, 0x88, 0x4c, 0xc2, 0x98, 0x02, 0x06, 0x9a, 0x0d, 0xa9, 0xec, 0xf7, 0x14,
		0x52, 0xaf, 0xc0, 0x72, 0x60, 0x92, 0xea, 0x6a, 0xd6, 0x50, 0xc4, 0xe6, 0xe5, 0x38, 0x4b, 0x6a,
		0xb2, 0xc0, 0x29, 0x04, 0xa6, 0x94, 0x71, 0xe4, 0x19, 0x6d, 0x01, 0xd8, 0x16, 0xb6, 0x07, 0x6a,
		0x1f, 0xeb, 0x66, 0x25, 0x7f, 0x8c, 0x97, 0xda, 0x44, 0x65, 0xc6, 0x4b, 0x36, 0x93, 0xea, 0x26,
		0xba, 0x3e, 0x09, 0xb5, 0xc5, 0x63, 0x22, 0x65, 0x97, 0x6d, 0xb2, 0x99, 0x68, 0xdb, 0x87, 0xb2,
		0x8b, 0x49, 0xdc, 0xe3, 0x3e, 0x9f, 0x59, 0x81, 0x1a, 0x51, 0x8b, 0x9d, 0x99, 0xc2, 0x61, 0x6c,
		0x62, 0x4b, 0x6e, 0xf8, 0x11, 0x3d, 0x0e, 0x81, 0x40, 0xa5, 0x61, 0x05, 0x34, 0x0b, 0x95, 0x84,
		0x70, 0x4f, 0x1b, 0xe1, 0xf5, 0xfb, 0x50, 0x8e, 0xba, 0x07, 0xad, 0x41, 0xce, 0xf3, 0x35, 0xd7,
		0xa7, 0x51, 0x98, 0x53, 0xd8, 0x03, 0x92, 0x20, 0x83, 0xad, 0x3e, 0xcd, 0x72, 0x39, 0x85, 0xfc,
		0x8b, 0x7e, 0x68, 0x32, 0xe1, 0x0c, 0x9d, 0xf0, 0xc7, 0x67, 0x57, 0x34, 0xc2, 0x3c, 0x3d, 0xef,
		0xf5, 0x17, 0x60, 0x29, 0x32, 0x81, 0xa4, 0xaf, 0xae, 0xfe, 0x28, 0x3c, 0x3c, 0x97, 0x1a, 0xbd,
		0x02, 0x6b, 0x63, 0xcb, 0xb0, 0x7c, 0xec, 0x3a, 0x2e, 0x26, 0x11, 0xcb, 0x5e, 0x55, 0xf9, 0xe7,
		0xc5, 0x63, 0x62, 0x6e, 0x3f, 0xac, 0xcd, 0x58, 0x94, 0xd5, 0xf1, 0xac, 0xf0, 0xe9, 0x42, 0xfe,
		0x5f, 0x16, 0xa5, 0x37, 0xde, 0x78, 0xe3, 0x8d, 0x74, 0xf5, 0x4b, 0x0b, 0xb0, 0x36, 0x6f, 0xcf,
		0xcc, 0xdd, 0xbe, 0x67, 0x60, 0xc1, 0x1a, 0x8f, 0x0e, 0xb0, 0x4b, 0x9d, 0x94, 0x53, 0xf8, 0x13,
		0x6a, 0x40, 0xce, 0xd4, 0x0e, 0xb0, 0x59, 0xc9, 0x5e, 0x48, 0x6d, 0x94, 0xaf, 0x3c, 0x93, 0x68,
		0x57, 0xd6, 0x76, 0x08, 0x44, 0x61, 0x48, 0xf4, 0x29, 0xc8, 0xf2, 0x14, 0x4d, 0x18, 0x9e, 0x4e,
		0xc6, 0x40, 0xf6, 0x92, 0x42, 0x71, 0xe8, 0x11, 0x28, 0x90, 0xbf, 0x2c, 0x36, 0x16, 0xa8, 0xcd,
		0x79, 0x22, 0x20, 0x71, 0x81, 0xd6, 0x21, 0x4f, 0xb7, 0x49, 0x1f, 0x8b, 0xd2, 0x16, 0x3c, 0x93,
		0xc0, 0xea, 0xe3, 0x81, 0x36, 0x36, 0x7d, 0xf5, 0xae, 0x66, 0x8e, 0x31, 0x0d, 0xf8, 0x82, 0x52,
		0xe2, 0xc2, 0xcf, 0x12, 0x19, 0x3a, 0x0f, 0x45, 0xb6, 0xab, 0x0c, 0xab, 0x8f, 0x5f, 0xa7, 0xd9,
		0x33, 0xa7, 0xb0, 0x8d, 0xd6, 0x22, 0x12, 0xf2, 0xfa, 0xdb, 0x9e, 0x6d, 0x89, 0xd0, 0xa4, 0xaf,
		0x20, 0x02, 0xfa, 0xfa, 0x17, 0xa6, 0x13, 0xf7, 0xa3, 0xf3, 0xa7, 0x37, 0x1d, 0x53, 0xd5, 0x6f,
		0xa6, 0x21, 0x4b, 0xf3, 0xc5, 0x32, 0x14, 0x7b, 0xaf, 0x76, 0x64, 0x75, 0xab, 0xbd, 0xbf, 0xb9,
		0x23, 0x4b, 0x29, 0x54, 0x06, 0xa0, 0x82, 0x9b, 0x3b, 0xed, 0x46, 0x4f, 0x4a, 0x07, 0xcf, 0xad,
		0xbd, 0xde, 0xb5, 0xe7, 0xa5, 0x4c, 0x00, 0xd8, 0x67, 0x82, 0x6c, 0x58, 0xe1, 0xb9, 0x2b, 0x52,
		0x0e, 0x49, 0x50, 0x62, 0x04, 0xad, 0x57, 0xe4, 0xad, 0x6b, 0xcf, 0x4b, 0x0b, 0x51, 0xc9, 0x73,
		0x57, 0xa4, 0x45, 0xb4, 0x04, 0x05, 0x2a, 0xd9, 0x6c, 0xb7, 0x77, 0xa4, 0x7c, 0xc0, 0xd9, 0xed,
		0x29, 0xad, 0xbd, 0x6d, 0xa9, 0x10, 0x70, 0x6e, 0x2b, 0xed, 0xfd, 0x8e, 0x04, 0x01, 0xc3, 0xae,
		0xdc, 0xed, 0x36, 0xb6, 0x65, 0xa9, 0x18, 0x68, 0x6c, 0xbe, 0xda, 0x93, 0xbb, 0x52, 0x29, 0x62,
		0xd6, 0x73, 0x57, 0xa4, 0xa5, 0xe0, 0x15, 0xf2, 0xde, 0xfe, 0xae, 0x54, 0x46, 0x2b, 0xb0, 0xc4,
		0x5e, 0x21, 0x8c, 0x58, 0x9e, 0x12, 0x5d, 0x7b, 0x5e, 0x92, 0x26, 0x86, 0x30, 0x96, 0x95, 0x88,
		0xe0, 0xda, 0xf3, 0x12, 0xaa, 0x36, 0x21, 0x47, 0xa3, 0x0b, 0x21, 0x28, 0xef, 0x34, 0x36, 0xe5,
		0x1d, 0xb5, 0xdd, 0xe9, 0xb5, 0xda, 0x7b, 0x8d, 0x1d, 0x29, 0x35, 0x91, 0x29, 0xf2, 0x67, 0xf6,
		0x5b, 0x8a, 0xbc, 0x25, 0xa5, 0xc3, 0xb2, 0x8e, 0xdc, 0xe8, 0xc9, 0x5b, 0x52, 0xa6, 0xaa, 0xc3,
		0xda, 0xbc, 0x3c, 0x39, 0x77, 0x67, 0x84, 0x96, 0x38, 0x7d, 0xcc, 0x12, 0x53, 0xae, 0x99, 0x25,
		0xfe, 0xa7, 0x34, 0xac, 0xce, 0xa9, 0x15, 0x73, 0x5f, 0xf2, 0x12, 0xe4, 0x58, 0x88, 0xb2, 0xea,
		0xf9, 0xd4, 0xdc, 0xa2, 0x43, 0x03, 0x76, 0xa6, 0x82, 0x52, 0x5c, 0xb8, 0x83, 0xc8, 0x1c, 0xd3,
		0x41, 0x10, 0x8a, 0x99, 0x9c, 0xfe, 0x23, 0x33, 0x39, 0x9d, 0x95, 0xbd, 0x6b, 0x49, 0xca, 0x1e,
		0x95, 0x9d, 0x2e, 0xb7, 0xe7, 0xe6, 0xe4, 0xf6, 0x1b, 0xb0, 0x32, 0x43, 0x94, 0x38, 0xc7, 0xbe,
		0x99, 0x82, 0xca, 0x71, 0xce, 0x89, 0xc9, 0x74, 0xe9, 0x48, 0xa6, 0xbb, 0x31, 0xed, 0xc1, 0x8b,
		0xc7, 0x2f, 0xc2, 0xcc, 0x5a, 0xbf, 0x9b, 0x82, 0x33, 0xf3, 0x3b, 0xc5, 0xb9, 0x36, 0x7c, 0x0a,
		0x16, 0x46, 0xd8, 0x3f, 0xb4, 0x45, 0xb7, 0xf4, 0xf1, 0x39, 0x35, 0x98, 0x0c, 0x4f, 0x2f, 0x36,
		0x47, 0x85, 0x8b, 0x78, 0xe6, 0xb8, 0x76, 0x8f, 0x59, 0x33, 0x63, 0xe9, 0x17, 0xd2, 0xf0, 0xf0,
		0x5c, 0xf2, 0xb9, 0x86, 0x3e, 0x0a, 0x60, 0x58, 0xce, 0xd8, 0x67, 0x1d, 0x11, 0x4b, 0xb0, 0x05,
		0x2a, 0xa1, 0xc9, 0x8b, 0x24, 0xcf, 0xb1, 0x1f, 0x8c, 0x67, 0xe8, 0x38, 0x30, 0x11, 0x55, 0x78,
		0x71, 0x62, 0x68, 0x96, 0x1a, 0xfa, 0xd8, 0x31, 0x33, 0x9d, 0x09, 0xcc, 0x4f, 0x80, 0xa4, 0x9b,
		0x06, 0xb6, 0x7c, 0xd5, 0xf3, 0x5d, 0xac, 0x8d, 0x0c, 0x6b, 0x48, 0x2b, 0x48, 0xbe, 0x9e, 0x1b,
		0x68, 0xa6, 0x87, 0x95, 0x65, 0x36, 0xdc, 0x15, 0xa3, 0x04, 0x41, 0x03, 0xc8, 0x0d, 0x21, 0x16,
		0x22, 0x08, 0x36, 0x1c, 0x20, 0xaa, 0xdf, 0xc8, 0x43, 0x31, 0xd4, 0x57, 0xa3, 0x8b, 0x50, 0xba,
		0xad, 0xdd, 0xd5, 0x54, 0x71, 0x56, 0x62, 0x9e, 0x28, 0x12, 0x59, 0x87, 0x9f, 0x97, 0x3e, 0x01,
		0x6b, 0x54, 0xc5, 0x1e, 0xfb, 0xd8, 0x55, 0x75, 0x53, 0xf3, 0x3c, 0xea, 0xb4, 0x3c, 0x55, 0x45,
		0x64, 0xac, 0x4d, 0x86, 0x9a, 0x62, 0x04, 0x5d, 0x85, 0x55, 0x8a, 0x18, 0x8d, 0x4d, 0xdf, 0x70,
		0x4c, 0xac, 0x92, 0xd3, 0x9b, 0x47, 0x2b, 0x49, 0x60, 0xd9, 0x0a, 0xd1, 0xd8, 0xe5, 0x0a, 0xc4,
		0x22, 0x0f, 0x6d, 0xc1, 0xa3, 0x14, 0x36, 0xc4, 0x16, 0x76, 0x35, 0x1f, 0xab, 0xf8, 0xf3, 0x63,
		0xcd, 0xf4, 0x54, 0xcd, 0xea, 0xab, 0x87, 0x9a, 0x77, 0x58, 0x59, 0x23, 0x04, 0x9b, 0xe9, 0x4a,
		0x4a, 0x39, 0x47, 0x14, 0xb7, 0xb9, 0x9e, 0x4c, 0xd5, 0x1a, 0x56, 0xff, 0xd3, 0x9a, 0x77, 0x88,
		0xea, 0x70, 0x86, 0xb2, 0x78, 0xbe, 0x6b, 0x58, 0x43, 0x55, 0x3f, 0xc4, 0xfa, 0x1d, 0x75, 0xec,
		0x0f, 0x5e, 0xac, 0x3c, 0x12, 0x7e, 0x3f, 0xb5, 0xb0, 0x4b, 0x75, 0x9a, 0x44, 0x65, 0xdf, 0x1f,
		0xbc, 0x88, 0xba, 0x50, 0x22, 0x8b, 0x31, 0x32, 0xee, 0x63, 0x75, 0x60, 0xbb, 0xb4, 0x34, 0x96,
		0xe7, 0xa4, 0xa6, 0x90, 0x07, 0x6b, 0x6d, 0x0e, 0xd8, 0xb5, 0xfb, 0xb8, 0x9e, 0xeb, 0x76, 0x64,
		0x79, 0x4b, 0x29, 0x0a, 0x96, 0x9b, 0xb6, 0x4b, 0x02, 0x6a, 0x68, 0x07, 0x0e, 0x2e, 0xb2, 0x80,
		0x1a, 0xda, 0xc2, 0xbd, 0x57, 0x61, 0x55, 0xd7, 0xd9, 0x9c, 0x0d, 0x5d, 0xe5, 0x67, 0x2c, 0xaf,
		0x22, 0x45, 0x9c, 0xa5, 0xeb, 0xdb, 0x4c, 0x81, 0xc7, 0xb8, 0x87, 0xae, 0xc3, 0xc3, 0x13, 0x67,
		0x85, 0x81, 0x2b, 0x33, 0xb3, 0x9c, 0x86, 0x5e, 0x85, 0x55, 0xe7, 0x68, 0x16, 0x88, 0x22, 0x6f,
		0x74, 0x8e, 0xa6, 0x61, 0x2f, 0xc0, 0x9a, 0x73, 0xe8, 0xcc, 0xe2, 0x9e, 0x0e, 0xe3, 0x90, 0x73,
		0xe8, 0x4c, 0x03, 0x9f, 0xa4, 0x07, 0x6e, 0x17, 0xeb, 0x9a, 0x8f, 0xfb, 0x95, 0xb3, 0x61, 0xf5,
		0xd0, 0x00, 0xba, 0x0c, 0x92, 0xae, 0xab, 0xd8, 0xd2, 0x0e, 0x4c, 0xac, 0x6a, 0x2e, 0xb6, 0x34,
		0xaf, 0x72, 0x3e, 0xac, 0x5c, 0xd6, 0x75, 0x99, 0x8e, 0x36, 0xe8, 0x20, 0x7a, 0x1a, 0x56, 0xec,
		0x83, 0xdb, 0x3a, 0x0b, 0x49, 0xd5, 0x71, 0xf1, 0xc0, 0x78, 0xbd, 0xf2, 0x04, 0xf5, 0xef, 0x32,
		0x19, 0xa0, 0x01, 0xd9, 0xa1, 0x62, 0xf4, 0x14, 0x48, 0xba, 0x77, 0xa8, 0xb9, 0x0e, 0xcd, 0xc9,
		0x9e, 0xa3, 0xe9, 0xb8, 0xf2, 0x24, 0x53, 0x65, 0xf2, 0x3d, 0x21, 0x26, 0x5b, 0xc2, 0xbb, 0x67,
		0x0c, 0x7c, 0xc1, 0x78, 0x89, 0x6d, 0x09, 0x2a, 0xe3, 0x6c, 0x1b, 0x20, 0x11, 0x57, 0x44, 0x5e,
		0xbc, 0x41, 0xd5, 0xca, 0xce, 0xa1, 0x13, 0x7e, 0xef, 0xe3, 0xb0, 0x44, 0x34, 0x27, 0x2f, 0x7d,
		0x8a, 0x35, 0x64, 0xce, 0x61, 0xe8, 0x8d, 0x1f, 0x5a, 0x6f, 0x5c, 0xad, 0x43, 0x29, 0x1c, 0x9f,
		0xa8, 0x00, 0x2c, 0x42, 0xa5, 0x14, 0x69, 0x56, 0x9a, 0xed, 0x2d, 0xd2, 0x66, 0xbc, 0x26, 0x4b,
		0x69, 0xd2, 0xee, 0xec, 0xb4, 0x7a, 0xb2, 0xaa, 0xec, 0xef, 0xf5, 0x5a, 0xbb, 0xb2, 0x94, 0x09,
		0xf7, 0xd5, 0xdf, 0x4a, 0x43, 0x39, 0x7a, 0x44, 0x42, 0x3f, 0x08, 0x67, 0xc5, 0x7d, 0x86, 0x87,
		0x7d, 0xf5, 0x9e, 0xe1, 0xd2, 0x2d, 0x33, 0xd2, 0x58, 0xf9, 0x0a, 0x16, 0x6d, 0x8d, 0x6b, 0x75,
		0xb1, 0xff, 0xb2, 0xe1, 0x92, 0x0d, 0x31, 0xd2, 0x7c, 0xb4, 0x03, 0xe7, 0x2d, 0x5b, 0xf5, 0x7c,
		0xcd, 0xea, 0x6b, 0x6e, 0x5f, 0x9d, 0xdc, 0x24, 0xa9, 0x9a, 0xae, 0x63, 0xcf, 0xb3, 0x59, 0xa9,
		0x0a, 0x58, 0x3e, 0x66, 0xd9, 0x5d, 0xae, 0x3c, 0xc9, 0xe1, 0x0d, 0xae, 0x3a, 0x15, 0x60, 0x99,
		0xe3, 0x02, 0xec, 0x11, 0x28, 0x8c, 0x34, 0x47, 0xc5, 0x96, 0xef, 0x1e, 0xd1, 0xc6, 0x38, 0xaf,
		0xe4, 0x47, 0x9a, 0x23, 0x93, 0xe7, 0x8f, 0xe6, 0x7c, 0xf2, 0x77, 0x19, 0x28, 0x85, 0x9b, 0x63,
		0x72, 0xd6, 0xd0, 0x69, 0x1d, 0x49, 0xd1, 0x4c, 0xf3, 0xf8, 0x89, 0xad, 0x74, 0xad, 0x49, 0x0a,
		0x4c, 0x7d, 0x81, 0xb5, 0xac, 0x0a, 0x43, 0x92, 0xe2, 0x4e, 0x72, 0x0b, 0x66, 0x2d, 0x42, 0x5e,
		0xe1, 0x4f, 0x68, 0x1b, 0x16, 0x6e, 0x7b, 0x94, 0x7b, 0x81, 0x72, 0x3f, 0x71, 0x32, 0xf7, 0xad,
		0x2e, 0x25, 0x2f, 0xdc, 0xea, 0xaa, 0x7b, 0x6d, 0x65, 0xb7, 0xb1, 0xa3, 0x70, 0x38, 0x3a, 0x07,
		0x59, 0x53, 0xbb, 0x7f, 0x14, 0x2d, 0x45, 0x54, 0x94, 0xd4, 0xf1, 0xe7, 0x20, 0x7b, 0x0f, 0x6b,
		0x77, 0xa2, 0x05, 0x80, 0x8a, 0x3e, 0xc4, 0xd0, 0xbf, 0x0c, 0x39, 0xea, 0x2f, 0x04, 0xc0, 0x3d,
		0x26, 0x3d, 0x84, 0xf2, 0x90, 0x6d, 0xb6, 0x15, 0x12, 0xfe, 0x12, 0x94, 0x98, 0x54, 0xed, 0xb4,
		0xe4, 0xa6, 0x2c, 0xa5, 0xab, 0x57, 0x61, 0x81, 0x39, 0x81, 0x6c, 0x8d, 0xc0, 0x0d, 0xd2, 0x43,
		0xfc, 0x91, 0x73, 0xa4, 0xc4, 0xe8, 0xfe, 0xee, 0xa6, 0xac, 0x48, 0xe9, 0xf0, 0xf2, 0x7a, 0x50,
		0x0a, 0xf7, 0xc5, 0x1f, 0x4d, 0x4c, 0xfd, 0x45, 0x0a, 0x8a, 0xa1, 0x3e, 0x97, 0x34, 0x28, 0x9a,
		0x69, 0xda, 0xf7, 0x54, 0xcd, 0x34, 0x34, 0x8f, 0x07, 0x05, 0x50, 0x51, 0x83, 0x48, 0x92, 0x2e,
		0xda, 0x47, 0x62, 0xfc, 0x3b, 0x29, 0x90, 0xa6, 0x5b, 0xcc, 0x29, 0x03, 0x53, 0xdf, 0x57, 0x03,
		0xdf, 0x4e, 0x41, 0x39, 0xda, 0x57, 0x4e, 0x99, 0x77, 0xf1, 0xfb, 0x6a, 0xde, 0xdf, 0xa7, 0x61,
		0x29, 0xd2, 0x4d, 0x26, 0xb5, 0xee, 0xf3, 0xb0, 0x62, 0xf4, 0xf1, 0xc8, 0xb1, 0x7d, 0x6c, 0xe9,
		0x47, 0xaa, 0x89, 0xef, 0x62, 0xb3, 0x52, 0xa5, 0x89, 0xe2, 0xf2, 0xc9, 0xfd, 0x6a, 0xad, 0x35,
		0xc1, 0xed, 0x10, 0x58, 0x7d, 0xb5, 0xb5, 0x25, 0xef, 0x76, 0xda, 0x3d, 0x79, 0xaf, 0xf9, 0xaa,
		0xba, 0xbf, 0xf7, 0xc3, 0x7b, 0xed, 0x97, 0xf7, 0x14, 0xc9, 0x98, 0x52, 0xfb, 0x10, 0xb7, 0x7a,
		0x07, 0xa4, 0x69, 0xa3, 0xd0, 0x59, 0x98, 0x67, 0x96, 0xf4, 0x10, 0x5a, 0x85, 0xe5, 0xbd, 0xb6,
		0xda, 0x6d, 0x6d, 0xc9, 0xaa, 0x7c, 0xf3, 0xa6, 0xdc, 0xec, 0x75, 0xd9, 0x0d, 0x44, 0xa0, 0xdd,
		0x8b, 0x6e, 0xea, 0xb7, 0x32, 0xb0, 0x3a, 0xc7, 0x12, 0xd4, 0xe0, 0x67, 0x07, 0x76, 0x9c, 0x79,
		0x36, 0x89, 0xf5, 0x35, 0x52, 0xf2, 0x3b, 0x9a, 0xeb, 0xf3, 0xa3, 0xc6, 0x53, 0x40, 0xbc, 0x64,
		0xf9, 0xc6, 0xc0, 0xc0, 0x2e, 0xbf, 0xb0, 0x61, 0x07, 0x8a, 0xe5, 0x89, 0x9c, 0xdd, 0xd9, 0xfc,
		0x00, 0x20, 0xc7, 0xf6, 0x0c, 0xdf, 0xb8, 0x8b, 0x55, 0xc3, 0x12, 0xb7, 0x3b, 0xe4, 0x80, 0x91,
		0x55, 0x24, 0x31, 0xd2, 0xb2, 0xfc, 0x40, 0xdb, 0xc2, 0x43, 0x6d, 0x4a, 0x9b, 0x24, 0xf0, 0x8c,
		0x22, 0x89, 0x91, 0x40, 0xfb, 0x22, 0x94, 0xfa, 0xf6, 0x98, 0x74, 0x5d, 0x4c, 0x8f, 0xd4, 0x8b,
		0x94, 0x52, 0x64, 0xb2, 0x40, 0x85, 0xf7, 0xd3, 0x93, 0x6b, 0xa5, 0x92, 0x52, 0x64, 0x32, 0xa6,
		0x72, 0x09, 0x96, 0xb5, 0xe1, 0xd0, 0x25, 0xe4, 0x82, 0x88, 0x9d, 0x10, 0xca, 0x81, 0x98, 0x2a,
		0xae, 0xdf, 0x82, 0xbc, 0xf0, 0x03, 0x29, 0xc9, 0xc4, 0x13, 0xaa, 0xc3, 0x8e, 0xbd, 0xe9, 0x8d,
		0x82, 0x92, 0xb7, 0xc4, 0xe0, 0x45, 0x28, 0x19, 0x9e, 0x3a, 0xb9, 0x25, 0x4f, 0x5f, 0x48, 0x6f,
		0xe4, 0x95, 0xa2, 0xe1, 0x05, 0x37, 0x8c, 0xd5, 0x77, 0xd3, 0x50, 0x8e, 0xde, 0xf2, 0xa3, 0x2d,
		0xc8, 0x9b, 0xb6, 0xae, 0xd1, 0xd0, 0x62, 0x9f, 0x98, 0x36, 0x62, 0x3e, 0x0c, 0xd4, 0x76, 0xb8,
		0xbe, 0x12, 0x20, 0xd7, 0xff, 0x36, 0x05, 0x79, 0x21, 0x46, 0x67, 0x20, 0xeb, 0x68, 0xfe, 0x21,
		0xa5, 0xcb, 0x6d, 0xa6, 0xa5, 0x94, 0x42, 0x9f, 0x89, 0xdc, 0x73, 0x34, 0x8b, 0x86, 0x00, 0x97,
		0x93, 0x67, 0xb2, 0xae, 0x26, 0xd6, 0xfa, 0xf4, 0xf8, 0x61, 0x8f, 0x46, 0xd8, 0xf2, 0x3d, 0xb1,
		0xae, 0x5c, 0xde, 0xe4, 0x62, 0xf4, 0x0c, 0xac, 0xf8, 0xae, 0x66, 0x98, 0x11, 0xdd, 0x2c, 0xd5,
		0x95, 0xc4, 0x40, 0xa0, 0x5c, 0x87, 0x73, 0x82, 0xb7, 0x8f, 0x7d, 0x4d, 0x3f, 0xc4, 0xfd, 0x09,
		0x68, 0x81, 0x5e, 0x33, 0x9c, 0xe5, 0x0a, 0x5b, 0x7c, 0x5c, 0x60, 0xab, 0xdf, 0x4e, 0xc1, 0x8a,
		0x38, 0x30, 0xf5, 0x03, 0x67, 0xed, 0x02, 0x68, 0x96, 0x65, 0xfb, 0x61, 0x77, 0xcd, 0x86, 0xf2,
		0x0c, 0xae, 0xd6, 0x08, 0x40, 0x4a, 0x88, 0x60, 0x7d, 0x04, 0x30, 0x19, 0x39, 0xd6, 0x6d, 0xe7,
		0xa1, 0xc8, 0x3f, 0xe1, 0xd0, 0xef, 0x80, 0xec, 0x88, 0x0d, 0x4c, 0x44, 0x4e, 0x56, 0x68, 0x0d,
		0x72, 0x07, 0x78, 0x68, 0x58, 0xfc, 0x62, 0x96, 0x3d, 0x88, 0x8b, 0x90, 0x6c, 0x70, 0x11, 0xb2,
		0xf9, 0x39, 0x58, 0xd5, 0xed, 0xd1, 0xb4, 0xb9, 0x9b, 0xd2, 0xd4, 0x31, 0xdf, 0xfb, 0x74, 0xea,
		0x35, 0x98, 0xb4, 0x98, 0xff, 0x9b, 0x4a, 0xfd, 0x46, 0x3a, 0xb3, 0xdd, 0xd9, 0xfc, 0x6a, 0x7a,
		0x7d, 0x9b, 0x41, 0x3b, 0x62, 0xa6, 0x0a, 0x1e, 0x98, 0x58, 0x27, 0xd6, 0xc3, 0x17, 0x2f, 0xc1,
		0xb3, 0x43, 0xc3, 0x3f, 0x1c, 0x1f, 0xd4, 0x74, 0x7b, 0x74, 0x79, 0x68, 0x0f, 0xed, 0xc9, 0xa7,
		0x4f, 0xf2, 0x44, 0x1f, 0xe8, 0x7f, 0xfc, 0xf3, 0x67, 0x21, 0x90, 0xae, 0xc7, 0x7e, 0x2b, 0xad,
		0xef, 0xc1, 0x2a, 0x57, 0x56, 0xe9, 0xf7, 0x17, 0x76, 0x8a, 0x40, 0x27, 0xde, 0x61, 0x55, 0xbe,
		0xfe, 0x1e, 0x2d, 0xd7, 0xca, 0x0a, 0x87, 0x92, 0x31, 0x76, 0xd0, 0xa8, 0x2b, 0xf0, 0x70, 0x84,
		0x8f, 0x6d, 0x4d, 0xec, 0xc6, 0x30, 0x7e, 0x8b, 0x33, 0xae, 0x86, 0x18, 0xbb, 0x1c, 0x5a, 0x6f,
		0xc2, 0xd2, 0x69, 0xb8, 0xfe, 0x8a, 0x73, 0x95, 0x70, 0x98, 0x64, 0x1b, 0x96, 0x29, 0x89, 0x3e,
		0xf6, 0x7c, 0x7b, 0x44, 0xf3, 0xde, 0xc9, 0x34, 0x7f, 0xfd, 0x1e, 0xdb, 0x2b, 0x65, 0x02, 0x6b,
		0x06, 0xa8, 0x7a, 0x1d, 0xe8, 0x27, 0xa7, 0x3e, 0xd6, 0xcd, 0x18, 0x86, 0xbf, 0xe1, 0x86, 0x04,
		0xfa, 0xf5, 0xcf, 0xc2, 0x1a, 0xf9, 0x9f, 0xa6, 0xa5, 0xb0, 0x25, 0xf1, 0x17, 0x5e, 0x95, 0x6f,
		0xbf, 0xc9, 0xb6, 0xe3, 0x6a, 0x40, 0x10, 0xb2, 0x29, 0xb4, 0x8a, 0x43, 0xec, 0xfb, 0xd8, 0xf5,
		0x54, 0xcd, 0x9c, 0x67, 0x5e, 0xe8, 0xc6, 0xa0, 0xf2, 0xe5, 0xf7, 0xa3, 0xab, 0xb8, 0xcd, 0x90,
		0x0d, 0xd3, 0xac, 0xef, 0xc3, 0xd9, 0x39, 0x51, 0x91, 0x80, 0xf3, 0x2d, 0xce, 0xb9, 0x36, 0x13,
		0x19, 0x84, 0xb6, 0x03, 0x42, 0x1e, 0xac, 0x65, 0x02, 0xce, 0x5f, 0xe1, 0x9c, 0x88, 0x63, 0xc5,
		0x92, 0x12, 0xc6, 0x5b, 0xb0, 0x72, 0x17, 0xbb, 0x07, 0xb6, 0xc7, 0x6f, 0x69, 0x12, 0xd0, 0xbd,
		0xcd, 0xe9, 0x96, 0x39, 0x90, 0x5e, 0xdb, 0x10, 0xae, 0xeb, 0x90, 0x1f, 0x68, 0x3a, 0x4e, 0x40,
		0xf1, 0x15, 0x4e, 0xb1, 0x48, 0xf4, 0x09, 0xb4, 0x01, 0xa5, 0xa1, 0xcd, 0x2b, 0x53, 0x3c, 0xfc,
		0x1d, 0x0e, 0x2f, 0x0a, 0x0c, 0xa7, 0x70, 0x6c, 0x67, 0x6c, 0x92, 0xb2, 0x15, 0x4f, 0xf1, 0xab,
		0x82, 0x42, 0x60, 0x38, 0xc5, 0x29, 0xdc, 0xfa, 0x6b, 0x82, 0xc2, 0x0b, 0xf9, 0xf3, 0x25, 0x28,
		0xda, 0x96, 0x79, 0x64, 0x5b, 0x49, 0x8c, 0xf8, 0x75, 0xce, 0x00, 0x1c, 0x42, 0x08, 0x6e, 0x40,
		0x21, 0xe9, 0x42, 0xfc, 0xe6, 0xfb, 0x62, 0x7b, 0x88, 0x15, 0xd8, 0x86, 0x65, 0x91, 0xa0, 0x0c,
		0xdb, 0x4a, 0x40, 0xf1, 0x5b, 0x9c, 0xa2, 0x1c, 0x82, 0xf1, 0x69, 0xf8, 0xd8, 0xf3, 0x87, 0x38,
		0x09, 0xc9, 0xbb, 0x62, 0x1a, 0x1c, 0xc2, 0x5d, 0x79, 0x80, 0x2d, 0xfd, 0x30, 0x19, 0xc3, 0x6f,
		0x0b, 0x57, 0x0a, 0x0c, 0xa1, 0x68, 0xc2, 0xd2, 0x48, 0x73, 0xbd, 0x43, 0xcd, 0x4c, 0xb4, 0x1c,
		0xbf, 0xc3, 0x39, 0x4a, 0x01, 0x88, 0x7b, 0x64, 0x6c, 0x9d, 0x86, 0xe6, 0xab, 0xc2, 0x23, 0x21,
		0x18, 0xdf, 0x7a, 0x9e, 0x4f, 0xaf, 0xb4, 0x4e, 0xc3, 0xf6, 0xbb, 0x62, 0xeb, 0x31, 0xec, 0x6e,
		0x98, 0xf1, 0x06, 0x14, 0x3c, 0xe3, 0x7e, 0x22, 0x9a, 0xdf, 0x13, 0x2b, 0x4d, 0x01, 0x04, 0xfc,
		0x2a, 0x9c, 0x9b, 0x5b, 0x26, 0x12, 0x90, 0xfd, 0x3e, 0x27, 0x3b, 0x33, 0xa7, 0x54, 0xf0, 0x94,
		0x70, 0x5a, 0xca, 0x3f, 0x10, 0x29, 0x01, 0x4f, 0x71, 0x75, 0xc8, 0x59, 0xc1, 0xd3, 0x06, 0xa7,
		0xf3, 0xda, 0x1f, 0x0a, 0xaf, 0x31, 0x6c, 0xc4, 0x6b, 0x3d, 0x38, 0xc3, 0x19, 0x4f, 0xb7, 0xae,
		0x5f, 0x13, 0x89, 0x95, 0xa1, 0xf7, 0xa3, 0xab, 0xfb, 0x39, 0x58, 0x0f, 0xdc, 0x29, 0x9a, 0x52,
		0x4f, 0x1d, 0x69, 0x4e, 0x02, 0xe6, 0xaf, 0x73, 0x66, 0x91, 0xf1, 0x83, 0xae, 0xd6, 0xdb, 0xd5,
		0x1c, 0x42, 0xfe, 0x0a, 0x54, 0x04, 0xf9, 0xd8, 0x72, 0xb1, 0x6e, 0x0f, 0x2d, 0xe3, 0x3e, 0xee,
		0x27, 0xa0, 0xfe, 0xa3, 0xa9, 0xa5, 0xda, 0x0f, 0xc1, 0x09, 0x73, 0x0b, 0xa4, 0xa0, 0x57, 0x51,
		0x8d, 0x91, 0x63, 0xbb, 0x7e, 0x0c, 0xe3, 0x37, 0xc4, 0x4a, 0x05, 0xb8, 0x16, 0x85, 0xd5, 0x65,
		0x28, 0xd3, 0xc7, 0xa4, 0x21, 0xf9, 0xc7, 0x9c, 0x68, 0x69, 0x82, 0xe2, 0x89, 0x43, 0xb7, 0x47,
		0x8e, 0xe6, 0x26, 0xc9, 0x7f, 0x7f, 0x22, 0x12, 0x07, 0x87, 0xf0, 0xc4, 0xe1, 0x1f, 0x39, 0x98,
		0x54, 0xfb, 0x04, 0x0c, 0xdf, 0x14, 0x89, 0x43, 0x60, 0x38, 0x85, 0x68, 0x18, 0x12, 0x50, 0xfc,
		0xa9, 0xa0, 0x10, 0x18, 0x42, 0xf1, 0x99, 0x49, 0xa1, 0x75, 0xf1, 0xd0, 0xf0, 0x7c, 0x97, 0xb5,
		0xc2, 0x27, 0x53, 0xfd, 0xd9, 0xfb, 0xd1, 0x26, 0x4c, 0x09, 0x41, 0xeb, 0xb7, 0x60, 0x79, 0xaa,
		0xc5, 0x40, 0x71, 0xbf, 0x5f, 0xa9, 0xfc, 0xd8, 0x07, 0x3c, 0x19, 0x45, 0x3b, 0x8c, 0xfa, 0x0e,
		0x59, 0xf7, 0x68, 0x1f, 0x10, 0x4f, 0xf6, 0xe6, 0x07, 0xc1, 0xd2, 0x47, 0xda, 0x80, 0xfa, 0x4d,
		0x58, 0x8a, 0xf4, 0x00, 0xf1, 0x54, 0x3f, 0xce, 0xa9, 0x4a, 0xe1, 0x16, 0xa0, 0x7e, 0x15, 0xb2,
		0xa4, 0x9e, 0xc7, 0xc3, 0x7f, 0x82, 0xc3, 0xa9, 0x7a, 0xfd, 0x93, 0x90, 0x17, 0x75, 0x3c, 0x1e,
		0xfa, 0x93, 0x1c, 0x1a, 0x40, 0x08, 0x5c, 0xd4, 0xf0, 0x78, 0xf8, 0x4f, 0x09, 0xb8, 0x80, 0x10,
		0x78, 0x72, 0x17, 0xfe, 0xe5, 0x4f, 0x67, 0x79, 0x1e, 0x16, 0xbe, 0xbb, 0x01, 0x8b, 0xbc, 0x78,
		0xc7, 0xa3, 0xbf, 0xc0, 0x5f, 0x2e, 0x10, 0xf5, 0x17, 0x20, 0x97, 0xd0, 0xe1, 0x3f, 0xc3, 0xa1,
		0x4c, 0xbf, 0xde, 0x84, 0x62, 0xa8, 0x60, 0xc7, 0xc3, 0x7f, 0x96, 0xc3, 0xc3, 0x28, 0x62, 0x3a,
		0x2f, 0xd8, 0xf1, 0x04, 0x3f, 0x27, 0x4c, 0xe7, 0x08, 0xe2, 0x36, 0x51, 0xab, 0xe3, 0xd1, 0x3f,
		0x2f, 0xbc, 0x2e, 0x20, 0xf5, 0x97, 0xa0, 0x10, 0xe4, 0xdf, 0x78, 0xfc, 0x2f, 0x70, 0xfc, 0x04,
		0x43, 0x3c, 0x10, 0xca, 0xff, 0xf1, 0x14, 0xbf, 0x28, 0x3c, 0x10, 0x42, 0x91, 0x6d, 0x34, 0x5d,
		0xd3, 0xe3, 0x99, 0x7e, 0x49, 0x6c, 0xa3, 0xa9, 0x92, 0x4e, 0x56, 0x93, 0xa6, 0xc1, 0x78, 0x8a,
		0x5f, 0x16, 0xab, 0x49, 0xf5, 0x89, 0x19, 0xd3, 0x45, 0x32, 0x9e, 0xe3, 0x8b, 0xc2, 0x8c, 0xa9,
		0x1a, 0x59, 0xef, 0x00, 0x9a, 0x2d, 0x90, 0xf1, 0x7c, 0x5f, 0xe2, 0x7c, 0x2b, 0x33, 0xf5, 0xb1,
		0xfe, 0x32, 0x9c, 0x99, 0x5f, 0x1c, 0xe3, 0x59, 0xbf, 0xfc, 0xc1, 0xd4, 0x71, 0x26, 0x5c, 0x1b,
		0xeb, 0xbd, 0x49, 0x96, 0x0d, 0x17, 0xc6, 0x78, 0xda, 0xb7, 0x3e, 0x88, 0x26, 0xda, 0x70, 0x5d,
		0xac, 0x37, 0x00, 0x26, 0x35, 0x29, 0x9e, 0xeb, 0x6d, 0xce, 0x15, 0x02, 0x91, 0xad, 0xc1, 0x4b,
		0x52, 0x3c, 0xfe, 0x2b, 0x62, 0x6b, 0x70, 0x04, 0xd9, 0x1a, 0xa2, 0x1a, 0xc5, 0xa3, 0xdf, 0x11,
		0x5b, 0x43, 0x40, 0xea, 0x37, 0x20, 0x6f, 0x8d, 0x4d, 0x93, 0xc4, 0x16, 0x3a, 0xf9, 0x27, 0x59,
		0x95, 0x7f, 0x7d, 0xc0, 0xc1, 0x02, 0x50, 0xbf, 0x0a, 0x39, 0x3c, 0x3a, 0xc0, 0xfd, 0x38, 0xe4,
		0xbf, 0x3d, 0x10, 0xf9, 0x84, 0x68, 0xd7, 0x5f, 0x02, 0x60, 0x87, 0x69, 0xfa, 0xa1, 0x28, 0x06,
		0xfb, 0xef, 0x0f, 0xf8, 0x8f, 0x25, 0x26, 0x90, 0x09, 0x01, 0xfb, 0xe9, 0xc5, 0xc9, 0x04, 0xef,
		0x47, 0x09, 0xe8, 0x01, 0xfc, 0x3a, 0x2c, 0xde, 0xf6, 0x6c, 0xcb, 0xd7, 0x86, 0x71, 0xe8, 0xff,
		0xe0, 0x68, 0xa1, 0x4f, 0x1c, 0x36, 0xb2, 0x5d, 0xec, 0x6b, 0x43, 0x2f, 0x0e, 0xfb, 0x9f, 0x1c,
		0x1b, 0x00, 0x08, 0x58, 0xd7, 0x3c, 0x3f, 0xc9, 0xbc, 0xbf, 0x23, 0xc0, 0x02, 0x40, 0x8c, 0x26,
		0xff, 0xdf, 0xc1, 0x47, 0x71, 0xd8, 0xef, 0x0a, 0xa3, 0xb9, 0x7e, 0xfd, 0x93, 0x50, 0x20, 0xff,
		0xb2, 0x5f, 0x40, 0xc5, 0x80, 0xff, 0x8b, 0x83, 0x27, 0x08, 0xf2, 0x66, 0xcf, 0xef, 0xfb, 0x46,
		0xbc, 0xb3, 0xff, 0x9b, 0xaf, 0xb4, 0xd0, 0xaf, 0x37, 0xa0, 0xe8, 0xf9, 0xfd, 0xfe, 0x98, 0x77,
		0x34, 0x31, 0xf0, 0xff, 0x79, 0x10, 0x1c, 0x72, 0x03, 0xcc, 0xa6, 0x3c, 0xff, 0xbe, 0x0e, 0xb6,
		0xed, 0x6d, 0x9b, 0xdd, 0xd4, 0xbd, 0x56, 0x8d, 0xbf, 0x72, 0x83, 0xff, 0x4b, 0x81, 0xe4, 0xdb,
		0x8e, 0x77, 0xe4, 0x8f, 0xdd, 0xbb, 0x47, 0xfc, 0xf2, 0x2d, 0x3b, 0xd2, 0x0c, 0x6b, 0xfd, 0x74,
		0x37, 0x76, 0xd5, 0x1d, 0x80, 0x1e, 0x21, 0xea, 0x11, 0x22, 0xf4, 0x28, 0x64, 0x7a, 0xb6, 0x43,
		0x3f, 0x38, 0x15, 0xaf, 0x14, 0x6b, 0x84, 0xb0, 0x46, 0x87, 0x15, 0x22, 0x47, 0xe7, 0x21, 0xdb,
		0x1b, 0xbb, 0x03, 0xfe, 0x93, 0x38, 0x31, 0x4e, 0x90, 0x0a, 0x1d, 0xa8, 0x3e, 0x0b, 0x39, 0xaa,
		0x8e, 0xa4, 0x09, 0x51, 0x86, 0x61, 0x51, 0x08, 0x9b, 0x09, 0xa9, 0xd3, 0xf7, 0x8a, 0xc1, 0xd4,
		0x64, 0x50, 0x50, 0xa4, 0x03, 0x8a, 0xcd, 0xfc, 0x77, 0xff, 0xe1, 0xb1, 0xd4, 0xd7, 0xfe, 0xf1,
		0xb1, 0xd4, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x83, 0x7c, 0x3b, 0x8c, 0x31, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(ungzipped, d); err != nil {
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

func init() { proto.RegisterFile("topsyturvy.proto", fileDescriptorTopsyturvy) }

var fileDescriptorTopsyturvy = []byte{
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
