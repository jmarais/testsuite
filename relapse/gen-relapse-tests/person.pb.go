// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person.proto

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

type Person struct {
	Name                 *string    `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Addresses            []*Address `protobuf:"bytes,2,rep,name=Addresses" json:"Addresses,omitempty"`
	Telephone            *string    `protobuf:"bytes,3,opt,name=Telephone" json:"Telephone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_134438c8f87a957d, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Person) GetTelephone() string {
	if m != nil && m.Telephone != nil {
		return *m.Telephone
	}
	return ""
}

type Address struct {
	Number               *int64   `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	Street               *string  `protobuf:"bytes,2,opt,name=Street" json:"Street,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_134438c8f87a957d, []int{1}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Address) GetStreet() string {
	if m != nil && m.Street != nil {
		return *m.Street
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "main.Person")
	proto.RegisterType((*Address)(nil), "main.Address")
}
func (this *Person) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return PersonDescription()
}
func PersonDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3782 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0x7f, 0x45, 0x1e, 0x52, 0xd4, 0xe8, 0x4a, 0xd6, 0x72, 0xe5, 0x9f, 0xdd, 0xa5, 0xed,
		0x58, 0xb6, 0x1b, 0x6d, 0xb0, 0xf6, 0xae, 0xbd, 0xdc, 0x26, 0x2e, 0x45, 0x71, 0x15, 0xba, 0x92,
		0xc8, 0x0c, 0xa5, 0xf8, 0x27, 0x28, 0x06, 0xa3, 0x99, 0x4b, 0x72, 0x76, 0x87, 0x33, 0x93, 0x99,
		0xe1, 0xae, 0xb5, 0xe8, 0x83, 0x0b, 0xf7, 0x07, 0x41, 0xd1, 0xff, 0x02, 0x4d, 0x5c, 0xc7, 0x4d,
		0x0a, 0xa4, 0x6e, 0xd3, 0xdf, 0x34, 0x6d, 0xda, 0xf4, 0xa9, 0x2f, 0x69, 0xfb, 0x54, 0x20, 0xef,
		0x7d, 0x68, 0x50, 0x03, 0xfd, 0x73, 0x1b, 0xb7, 0x35, 0xd0, 0x00, 0xfb, 0x12, 0xdc, 0xbf, 0xe1,
		0x0c, 0x49, 0xed, 0x8c, 0x02, 0xd8, 0x7e, 0x92, 0xe6, 0xdc, 0xf3, 0x7d, 0x73, 0xee, 0xb9, 0xe7,
		0x9e, 0x73, 0xee, 0x1d, 0xc2, 0xf7, 0xaf, 0xc2, 0xf9, 0x81, 0x6d, 0x0f, 0x4c, 0x7c, 0xd1, 0x71,
		0x6d, 0xdf, 0x3e, 0x1a, 0xf7, 0x2f, 0xea, 0xd8, 0xd3, 0x5c, 0xc3, 0xf1, 0x6d, 0x77, 0x93, 0xca,
		0xd0, 0x12, 0xd3, 0xd8, 0x14, 0x1a, 0xb5, 0x3d, 0x58, 0xbe, 0x6e, 0x98, 0x78, 0x3b, 0x50, 0xec,
		0x61, 0x1f, 0x3d, 0x07, 0xd9, 0xbe, 0x61, 0xe2, 0x6a, 0xea, 0x7c, 0x66, 0xa3, 0x74, 0xe9, 0xd1,
		0xcd, 0x29, 0xd0, 0x66, 0x14, 0xd1, 0x25, 0x62, 0x99, 0x22, 0x6a, 0xef, 0x64, 0x61, 0x65, 0xce,
		0x28, 0x42, 0x90, 0xb5, 0xd4, 0x11, 0x61, 0x4c, 0x6d, 0x14, 0x65, 0xfa, 0x3f, 0xaa, 0xc2, 0x82,
		0xa3, 0x6a, 0x37, 0xd5, 0x01, 0xae, 0xa6, 0xa9, 0x58, 0x3c, 0xa2, 0x87, 0x01, 0x74, 0xec, 0x60,
		0x4b, 0xc7, 0x96, 0x76, 0x5c, 0xcd, 0x9c, 0xcf, 0x6c, 0x14, 0xe5, 0x90, 0x04, 0x3d, 0x05, 0xcb,
		0xce, 0xf8, 0xc8, 0x34, 0x34, 0x25, 0xa4, 0x06, 0xe7, 0x33, 0x1b, 0x39, 0x59, 0x62, 0x03, 0xdb,
		0x13, 0xe5, 0xc7, 0x61, 0xe9, 0x36, 0x56, 0x6f, 0x86, 0x55, 0x4b, 0x54, 0xb5, 0x42, 0xc4, 0x21,
		0xc5, 0x26, 0x94, 0x47, 0xd8, 0xf3, 0xd4, 0x01, 0x56, 0xfc, 0x63, 0x07, 0x57, 0xb3, 0x74, 0xf6,
		0xe7, 0x67, 0x66, 0x3f, 0x3d, 0xf3, 0x12, 0x47, 0x1d, 0x1c, 0x3b, 0x18, 0x35, 0xa0, 0x88, 0xad,
		0xf1, 0x88, 0x31, 0xe4, 0x4e, 0xf0, 0x5f, 0xcb, 0x1a, 0x8f, 0xa6, 0x59, 0x0a, 0x04, 0xc6, 0x29,
		0x16, 0x3c, 0xec, 0xde, 0x32, 0x34, 0x5c, 0xcd, 0x53, 0x82, 0xc7, 0x67, 0x08, 0x7a, 0x6c, 0x7c,
		0x9a, 0x43, 0xe0, 0x50, 0x13, 0x8a, 0xf8, 0x55, 0x1f, 0x5b, 0x9e, 0x61, 0x5b, 0xd5, 0x05, 0x4a,
		0xf2, 0xd8, 0x9c, 0x55, 0xc4, 0xa6, 0x3e, 0x4d, 0x31, 0xc1, 0xa1, 0x2b, 0xb0, 0x60, 0x3b, 0xbe,
		0x61, 0x5b, 0x5e, 0xb5, 0x70, 0x3e, 0xb5, 0x51, 0xba, 0xf4, 0xe0, 0xdc, 0x40, 0xe8, 0x30, 0x1d,
		0x59, 0x28, 0xa3, 0x36, 0x48, 0x9e, 0x3d, 0x76, 0x35, 0xac, 0x68, 0xb6, 0x8e, 0x15, 0xc3, 0xea,
		0xdb, 0xd5, 0x22, 0x25, 0x38, 0x37, 0x3b, 0x11, 0xaa, 0xd8, 0xb4, 0x75, 0xdc, 0xb6, 0xfa, 0xb6,
		0x5c, 0xf1, 0x22, 0xcf, 0x68, 0x0d, 0xf2, 0xde, 0xb1, 0xe5, 0xab, 0xaf, 0x56, 0xcb, 0x34, 0x42,
		0xf8, 0x53, 0xed, 0xdb, 0x79, 0x58, 0x4a, 0x12, 0x62, 0xd7, 0x20, 0xd7, 0x27, 0xb3, 0xac, 0xa6,
		0x4f, 0xe3, 0x03, 0x86, 0x89, 0x3a, 0x31, 0xff, 0x23, 0x3a, 0xb1, 0x01, 0x25, 0x0b, 0x7b, 0x3e,
		0xd6, 0x59, 0x44, 0x64, 0x12, 0xc6, 0x14, 0x30, 0xd0, 0x6c, 0x48, 0x65, 0x7f, 0xa4, 0x90, 0x7a,
		0x09, 0x96, 0x02, 0x93, 0x14, 0x57, 0xb5, 0x06, 0x22, 0x36, 0x2f, 0xc6, 0x59, 0xb2, 0xd9, 0x12,
		0x38, 0x99, 0xc0, 0xe4, 0x0a, 0x8e, 0x3c, 0xa3, 0x6d, 0x00, 0xdb, 0xc2, 0x76, 0x5f, 0xd1, 0xb1,
		0x66, 0x56, 0x0b, 0x27, 0x78, 0xa9, 0x43, 0x54, 0x66, 0xbc, 0x64, 0x33, 0xa9, 0x66, 0xa2, 0xab,
		0x93, 0x50, 0x5b, 0x38, 0x21, 0x52, 0xf6, 0xd8, 0x26, 0x9b, 0x89, 0xb6, 0x43, 0xa8, 0xb8, 0x98,
		0xc4, 0x3d, 0xd6, 0xf9, 0xcc, 0x8a, 0xd4, 0x88, 0xcd, 0xd8, 0x99, 0xc9, 0x1c, 0xc6, 0x26, 0xb6,
		0xe8, 0x86, 0x1f, 0xd1, 0x23, 0x10, 0x08, 0x14, 0x1a, 0x56, 0x40, 0xb3, 0x50, 0x59, 0x08, 0xf7,
		0xd5, 0x11, 0x5e, 0xbf, 0x03, 0x95, 0xa8, 0x7b, 0xd0, 0x2a, 0xe4, 0x3c, 0x5f, 0x75, 0x7d, 0x1a,
		0x85, 0x39, 0x99, 0x3d, 0x20, 0x09, 0x32, 0xd8, 0xd2, 0x69, 0x96, 0xcb, 0xc9, 0xe4, 0x5f, 0xf4,
		0x13, 0x93, 0x09, 0x67, 0xe8, 0x84, 0x3f, 0x36, 0xbb, 0xa2, 0x11, 0xe6, 0xe9, 0x79, 0xaf, 0x3f,
		0x0b, 0x8b, 0x91, 0x09, 0x24, 0x7d, 0x75, 0xed, 0xa7, 0xe1, 0xfe, 0xb9, 0xd4, 0xe8, 0x25, 0x58,
		0x1d, 0x5b, 0x86, 0xe5, 0x63, 0xd7, 0x71, 0x31, 0x89, 0x58, 0xf6, 0xaa, 0xea, 0xbf, 0x2e, 0x9c,
		0x10, 0x73, 0x87, 0x61, 0x6d, 0xc6, 0x22, 0xaf, 0x8c, 0x67, 0x85, 0x4f, 0x16, 0x0b, 0xff, 0xb6,
		0x20, 0xbd, 0xf6, 0xda, 0x6b, 0xaf, 0xa5, 0x6b, 0x5f, 0xcc, 0xc3, 0xea, 0xbc, 0x3d, 0x33, 0x77,
		0xfb, 0xae, 0x41, 0xde, 0x1a, 0x8f, 0x8e, 0xb0, 0x4b, 0x9d, 0x94, 0x93, 0xf9, 0x13, 0x6a, 0x40,
		0xce, 0x54, 0x8f, 0xb0, 0x59, 0xcd, 0x9e, 0x4f, 0x6d, 0x54, 0x2e, 0x3d, 0x95, 0x68, 0x57, 0x6e,
		0xee, 0x12, 0x88, 0xcc, 0x90, 0xe8, 0x53, 0x90, 0xe5, 0x29, 0x9a, 0x30, 0x3c, 0x99, 0x8c, 0x81,
		0xec, 0x25, 0x99, 0xe2, 0xd0, 0x03, 0x50, 0x24, 0x7f, 0x59, 0x6c, 0xe4, 0xa9, 0xcd, 0x05, 0x22,
		0x20, 0x71, 0x81, 0xd6, 0xa1, 0x40, 0xb7, 0x89, 0x8e, 0x45, 0x69, 0x0b, 0x9e, 0x49, 0x60, 0xe9,
		0xb8, 0xaf, 0x8e, 0x4d, 0x5f, 0xb9, 0xa5, 0x9a, 0x63, 0x4c, 0x03, 0xbe, 0x28, 0x97, 0xb9, 0xf0,
		0xb3, 0x44, 0x86, 0xce, 0x41, 0x89, 0xed, 0x2a, 0xc3, 0xd2, 0xf1, 0xab, 0x34, 0x7b, 0xe6, 0x64,
		0xb6, 0xd1, 0xda, 0x44, 0x42, 0x5e, 0x7f, 0xc3, 0xb3, 0x2d, 0x11, 0x9a, 0xf4, 0x15, 0x44, 0x40,
		0x5f, 0xff, 0xec, 0x74, 0xe2, 0x7e, 0x68, 0xfe, 0xf4, 0xa6, 0x63, 0xaa, 0xf6, 0xad, 0x34, 0x64,
		0x69, 0xbe, 0x58, 0x82, 0xd2, 0xc1, 0xcb, 0xdd, 0x96, 0xb2, 0xdd, 0x39, 0xdc, 0xda, 0x6d, 0x49,
		0x29, 0x54, 0x01, 0xa0, 0x82, 0xeb, 0xbb, 0x9d, 0xc6, 0x81, 0x94, 0x0e, 0x9e, 0xdb, 0xfb, 0x07,
		0x57, 0x9e, 0x91, 0x32, 0x01, 0xe0, 0x90, 0x09, 0xb2, 0x61, 0x85, 0xa7, 0x2f, 0x49, 0x39, 0x24,
		0x41, 0x99, 0x11, 0xb4, 0x5f, 0x6a, 0x6d, 0x5f, 0x79, 0x46, 0xca, 0x47, 0x25, 0x4f, 0x5f, 0x92,
		0x16, 0xd0, 0x22, 0x14, 0xa9, 0x64, 0xab, 0xd3, 0xd9, 0x95, 0x0a, 0x01, 0x67, 0xef, 0x40, 0x6e,
		0xef, 0xef, 0x48, 0xc5, 0x80, 0x73, 0x47, 0xee, 0x1c, 0x76, 0x25, 0x08, 0x18, 0xf6, 0x5a, 0xbd,
		0x5e, 0x63, 0xa7, 0x25, 0x95, 0x02, 0x8d, 0xad, 0x97, 0x0f, 0x5a, 0x3d, 0xa9, 0x1c, 0x31, 0xeb,
		0xe9, 0x4b, 0xd2, 0x62, 0xf0, 0x8a, 0xd6, 0xfe, 0xe1, 0x9e, 0x54, 0x41, 0xcb, 0xb0, 0xc8, 0x5e,
		0x21, 0x8c, 0x58, 0x9a, 0x12, 0x5d, 0x79, 0x46, 0x92, 0x26, 0x86, 0x30, 0x96, 0xe5, 0x88, 0xe0,
		0xca, 0x33, 0x12, 0xaa, 0x35, 0x21, 0x47, 0xa3, 0x0b, 0x21, 0xa8, 0xec, 0x36, 0xb6, 0x5a, 0xbb,
		0x4a, 0xa7, 0x7b, 0xd0, 0xee, 0xec, 0x37, 0x76, 0xa5, 0xd4, 0x44, 0x26, 0xb7, 0x3e, 0x73, 0xd8,
		0x96, 0x5b, 0xdb, 0x52, 0x3a, 0x2c, 0xeb, 0xb6, 0x1a, 0x07, 0xad, 0x6d, 0x29, 0x53, 0xd3, 0x60,
		0x75, 0x5e, 0x9e, 0x9c, 0xbb, 0x33, 0x42, 0x4b, 0x9c, 0x3e, 0x61, 0x89, 0x29, 0xd7, 0xcc, 0x12,
		0xff, 0x4b, 0x1a, 0x56, 0xe6, 0xd4, 0x8a, 0xb9, 0x2f, 0x79, 0x1e, 0x72, 0x2c, 0x44, 0x59, 0xf5,
		0x7c, 0x62, 0x6e, 0xd1, 0xa1, 0x01, 0x3b, 0x53, 0x41, 0x29, 0x2e, 0xdc, 0x41, 0x64, 0x4e, 0xe8,
		0x20, 0x08, 0xc5, 0x4c, 0x4e, 0xff, 0xa9, 0x99, 0x9c, 0xce, 0xca, 0xde, 0x95, 0x24, 0x65, 0x8f,
		0xca, 0x4e, 0x97, 0xdb, 0x73, 0x73, 0x72, 0xfb, 0x35, 0x58, 0x9e, 0x21, 0x4a, 0x9c, 0x63, 0x5f,
		0x4f, 0x41, 0xf5, 0x24, 0xe7, 0xc4, 0x64, 0xba, 0x74, 0x24, 0xd3, 0x5d, 0x9b, 0xf6, 0xe0, 0x85,
		0x93, 0x17, 0x61, 0x66, 0xad, 0xdf, 0x4e, 0xc1, 0xda, 0xfc, 0x4e, 0x71, 0xae, 0x0d, 0x9f, 0x82,
		0xfc, 0x08, 0xfb, 0x43, 0x5b, 0x74, 0x4b, 0x1f, 0x9b, 0x53, 0x83, 0xc9, 0xf0, 0xf4, 0x62, 0x73,
		0x54, 0xb8, 0x88, 0x67, 0x4e, 0x6a, 0xf7, 0x98, 0x35, 0x33, 0x96, 0x7e, 0x21, 0x0d, 0xf7, 0xcf,
		0x25, 0x9f, 0x6b, 0xe8, 0x43, 0x00, 0x86, 0xe5, 0x8c, 0x7d, 0xd6, 0x11, 0xb1, 0x04, 0x5b, 0xa4,
		0x12, 0x9a, 0xbc, 0x48, 0xf2, 0x1c, 0xfb, 0xc1, 0x78, 0x86, 0x8e, 0x03, 0x13, 0x51, 0x85, 0xe7,
		0x26, 0x86, 0x66, 0xa9, 0xa1, 0x0f, 0x9f, 0x30, 0xd3, 0x99, 0xc0, 0xfc, 0x04, 0x48, 0x9a, 0x69,
		0x60, 0xcb, 0x57, 0x3c, 0xdf, 0xc5, 0xea, 0xc8, 0xb0, 0x06, 0xb4, 0x82, 0x14, 0xea, 0xb9, 0xbe,
		0x6a, 0x7a, 0x58, 0x5e, 0x62, 0xc3, 0x3d, 0x31, 0x4a, 0x10, 0x34, 0x80, 0xdc, 0x10, 0x22, 0x1f,
		0x41, 0xb0, 0xe1, 0x00, 0x51, 0xfb, 0x66, 0x01, 0x4a, 0xa1, 0xbe, 0x1a, 0x5d, 0x80, 0xf2, 0x0d,
		0xf5, 0x96, 0xaa, 0x88, 0xb3, 0x12, 0xf3, 0x44, 0x89, 0xc8, 0xba, 0xfc, 0xbc, 0xf4, 0x09, 0x58,
		0xa5, 0x2a, 0xf6, 0xd8, 0xc7, 0xae, 0xa2, 0x99, 0xaa, 0xe7, 0x51, 0xa7, 0x15, 0xa8, 0x2a, 0x22,
		0x63, 0x1d, 0x32, 0xd4, 0x14, 0x23, 0xe8, 0x32, 0xac, 0x50, 0xc4, 0x68, 0x6c, 0xfa, 0x86, 0x63,
		0x62, 0x85, 0x9c, 0xde, 0x3c, 0x5a, 0x49, 0x02, 0xcb, 0x96, 0x89, 0xc6, 0x1e, 0x57, 0x20, 0x16,
		0x79, 0x68, 0x1b, 0x1e, 0xa2, 0xb0, 0x01, 0xb6, 0xb0, 0xab, 0xfa, 0x58, 0xc1, 0x9f, 0x1f, 0xab,
		0xa6, 0xa7, 0xa8, 0x96, 0xae, 0x0c, 0x55, 0x6f, 0x58, 0x5d, 0x25, 0x04, 0x5b, 0xe9, 0x6a, 0x4a,
		0x3e, 0x4b, 0x14, 0x77, 0xb8, 0x5e, 0x8b, 0xaa, 0x35, 0x2c, 0xfd, 0xd3, 0xaa, 0x37, 0x44, 0x75,
		0x58, 0xa3, 0x2c, 0x9e, 0xef, 0x1a, 0xd6, 0x40, 0xd1, 0x86, 0x58, 0xbb, 0xa9, 0x8c, 0xfd, 0xfe,
		0x73, 0xd5, 0x07, 0xc2, 0xef, 0xa7, 0x16, 0xf6, 0xa8, 0x4e, 0x93, 0xa8, 0x1c, 0xfa, 0xfd, 0xe7,
		0x50, 0x0f, 0xca, 0x64, 0x31, 0x46, 0xc6, 0x1d, 0xac, 0xf4, 0x6d, 0x97, 0x96, 0xc6, 0xca, 0x9c,
		0xd4, 0x14, 0xf2, 0xe0, 0x66, 0x87, 0x03, 0xf6, 0x6c, 0x1d, 0xd7, 0x73, 0xbd, 0x6e, 0xab, 0xb5,
		0x2d, 0x97, 0x04, 0xcb, 0x75, 0xdb, 0x25, 0x01, 0x35, 0xb0, 0x03, 0x07, 0x97, 0x58, 0x40, 0x0d,
		0x6c, 0xe1, 0xde, 0xcb, 0xb0, 0xa2, 0x69, 0x6c, 0xce, 0x86, 0xa6, 0xf0, 0x33, 0x96, 0x57, 0x95,
		0x22, 0xce, 0xd2, 0xb4, 0x1d, 0xa6, 0xc0, 0x63, 0xdc, 0x43, 0x57, 0xe1, 0xfe, 0x89, 0xb3, 0xc2,
		0xc0, 0xe5, 0x99, 0x59, 0x4e, 0x43, 0x2f, 0xc3, 0x8a, 0x73, 0x3c, 0x0b, 0x44, 0x91, 0x37, 0x3a,
		0xc7, 0xd3, 0xb0, 0x67, 0x61, 0xd5, 0x19, 0x3a, 0xb3, 0xb8, 0x27, 0xc3, 0x38, 0xe4, 0x0c, 0x9d,
		0x69, 0xe0, 0x63, 0xf4, 0xc0, 0xed, 0x62, 0x4d, 0xf5, 0xb1, 0x5e, 0x3d, 0x13, 0x56, 0x0f, 0x0d,
		0xa0, 0x8b, 0x20, 0x69, 0x9a, 0x82, 0x2d, 0xf5, 0xc8, 0xc4, 0x8a, 0xea, 0x62, 0x4b, 0xf5, 0xaa,
		0xe7, 0xc2, 0xca, 0x15, 0x4d, 0x6b, 0xd1, 0xd1, 0x06, 0x1d, 0x44, 0x4f, 0xc2, 0xb2, 0x7d, 0x74,
		0x43, 0x63, 0x21, 0xa9, 0x38, 0x2e, 0xee, 0x1b, 0xaf, 0x56, 0x1f, 0xa5, 0xfe, 0x5d, 0x22, 0x03,
		0x34, 0x20, 0xbb, 0x54, 0x8c, 0x9e, 0x00, 0x49, 0xf3, 0x86, 0xaa, 0xeb, 0xd0, 0x9c, 0xec, 0x39,
		0xaa, 0x86, 0xab, 0x8f, 0x31, 0x55, 0x26, 0xdf, 0x17, 0x62, 0xb2, 0x25, 0xbc, 0xdb, 0x46, 0xdf,
		0x17, 0x8c, 0x8f, 0xb3, 0x2d, 0x41, 0x65, 0x9c, 0x6d, 0x03, 0x24, 0xe2, 0x8a, 0xc8, 0x8b, 0x37,
		0xa8, 0x5a, 0xc5, 0x19, 0x3a, 0xe1, 0xf7, 0x3e, 0x02, 0x8b, 0x44, 0x73, 0xf2, 0xd2, 0x27, 0x58,
		0x43, 0xe6, 0x0c, 0x43, 0x6f, 0xfc, 0xc0, 0x7a, 0xe3, 0x5a, 0x1d, 0xca, 0xe1, 0xf8, 0x44, 0x45,
		0x60, 0x11, 0x2a, 0xa5, 0x48, 0xb3, 0xd2, 0xec, 0x6c, 0x93, 0x36, 0xe3, 0x95, 0x96, 0x94, 0x26,
		0xed, 0xce, 0x6e, 0xfb, 0xa0, 0xa5, 0xc8, 0x87, 0xfb, 0x07, 0xed, 0xbd, 0x96, 0x94, 0x09, 0xf7,
		0xd5, 0xdf, 0x49, 0x43, 0x25, 0x7a, 0x44, 0x42, 0x3f, 0x0e, 0x67, 0xc4, 0x7d, 0x86, 0x87, 0x7d,
		0xe5, 0xb6, 0xe1, 0xd2, 0x2d, 0x33, 0x52, 0x59, 0xf9, 0x0a, 0x16, 0x6d, 0x95, 0x6b, 0xf5, 0xb0,
		0xff, 0xa2, 0xe1, 0x92, 0x0d, 0x31, 0x52, 0x7d, 0xb4, 0x0b, 0xe7, 0x2c, 0x5b, 0xf1, 0x7c, 0xd5,
		0xd2, 0x55, 0x57, 0x57, 0x26, 0x37, 0x49, 0x8a, 0xaa, 0x69, 0xd8, 0xf3, 0x6c, 0x56, 0xaa, 0x02,
		0x96, 0x07, 0x2d, 0xbb, 0xc7, 0x95, 0x27, 0x39, 0xbc, 0xc1, 0x55, 0xa7, 0x02, 0x2c, 0x73, 0x52,
		0x80, 0x3d, 0x00, 0xc5, 0x91, 0xea, 0x28, 0xd8, 0xf2, 0xdd, 0x63, 0xda, 0x18, 0x17, 0xe4, 0xc2,
		0x48, 0x75, 0x5a, 0xe4, 0xf9, 0xc3, 0x39, 0x9f, 0xfc, 0x53, 0x06, 0xca, 0xe1, 0xe6, 0x98, 0x9c,
		0x35, 0x34, 0x5a, 0x47, 0x52, 0x34, 0xd3, 0x3c, 0x72, 0xcf, 0x56, 0x7a, 0xb3, 0x49, 0x0a, 0x4c,
		0x3d, 0xcf, 0x5a, 0x56, 0x99, 0x21, 0x49, 0x71, 0x27, 0xb9, 0x05, 0xb3, 0x16, 0xa1, 0x20, 0xf3,
		0x27, 0xb4, 0x03, 0xf9, 0x1b, 0x1e, 0xe5, 0xce, 0x53, 0xee, 0x47, 0xef, 0xcd, 0xfd, 0x42, 0x8f,
		0x92, 0x17, 0x5f, 0xe8, 0x29, 0xfb, 0x1d, 0x79, 0xaf, 0xb1, 0x2b, 0x73, 0x38, 0x3a, 0x0b, 0x59,
		0x53, 0xbd, 0x73, 0x1c, 0x2d, 0x45, 0x54, 0x94, 0xd4, 0xf1, 0x67, 0x21, 0x7b, 0x1b, 0xab, 0x37,
		0xa3, 0x05, 0x80, 0x8a, 0x3e, 0xc0, 0xd0, 0xbf, 0x08, 0x39, 0xea, 0x2f, 0x04, 0xc0, 0x3d, 0x26,
		0xdd, 0x87, 0x0a, 0x90, 0x6d, 0x76, 0x64, 0x12, 0xfe, 0x12, 0x94, 0x99, 0x54, 0xe9, 0xb6, 0x5b,
		0xcd, 0x96, 0x94, 0xae, 0x5d, 0x86, 0x3c, 0x73, 0x02, 0xd9, 0x1a, 0x81, 0x1b, 0xa4, 0xfb, 0xf8,
		0x23, 0xe7, 0x48, 0x89, 0xd1, 0xc3, 0xbd, 0xad, 0x96, 0x2c, 0xa5, 0xc3, 0xcb, 0xeb, 0x41, 0x39,
		0xdc, 0x17, 0x7f, 0x38, 0x31, 0xf5, 0x37, 0x29, 0x28, 0x85, 0xfa, 0x5c, 0xd2, 0xa0, 0xa8, 0xa6,
		0x69, 0xdf, 0x56, 0x54, 0xd3, 0x50, 0x3d, 0x1e, 0x14, 0x40, 0x45, 0x0d, 0x22, 0x49, 0xba, 0x68,
		0x1f, 0x8a, 0xf1, 0x6f, 0xa5, 0x40, 0x9a, 0x6e, 0x31, 0xa7, 0x0c, 0x4c, 0x7d, 0xa4, 0x06, 0xbe,
		0x99, 0x82, 0x4a, 0xb4, 0xaf, 0x9c, 0x32, 0xef, 0xc2, 0x47, 0x6a, 0xde, 0x3f, 0xa7, 0x61, 0x31,
		0xd2, 0x4d, 0x26, 0xb5, 0xee, 0xf3, 0xb0, 0x6c, 0xe8, 0x78, 0xe4, 0xd8, 0x3e, 0xb6, 0xb4, 0x63,
		0xc5, 0xc4, 0xb7, 0xb0, 0x59, 0xad, 0xd1, 0x44, 0x71, 0xf1, 0xde, 0xfd, 0xea, 0x66, 0x7b, 0x82,
		0xdb, 0x25, 0xb0, 0xfa, 0x4a, 0x7b, 0xbb, 0xb5, 0xd7, 0xed, 0x1c, 0xb4, 0xf6, 0x9b, 0x2f, 0x2b,
		0x87, 0xfb, 0x3f, 0xb9, 0xdf, 0x79, 0x71, 0x5f, 0x96, 0x8c, 0x29, 0xb5, 0x0f, 0x70, 0xab, 0x77,
		0x41, 0x9a, 0x36, 0x0a, 0x9d, 0x81, 0x79, 0x66, 0x49, 0xf7, 0xa1, 0x15, 0x58, 0xda, 0xef, 0x28,
		0xbd, 0xf6, 0x76, 0x4b, 0x69, 0x5d, 0xbf, 0xde, 0x6a, 0x1e, 0xf4, 0xd8, 0x0d, 0x44, 0xa0, 0x7d,
		0x10, 0xdd, 0xd4, 0x6f, 0x64, 0x60, 0x65, 0x8e, 0x25, 0xa8, 0xc1, 0xcf, 0x0e, 0xec, 0x38, 0xf3,
		0xf1, 0x24, 0xd6, 0x6f, 0x92, 0x92, 0xdf, 0x55, 0x5d, 0x9f, 0x1f, 0x35, 0x9e, 0x00, 0xe2, 0x25,
		0xcb, 0x37, 0xfa, 0x06, 0x76, 0xf9, 0x85, 0x0d, 0x3b, 0x50, 0x2c, 0x4d, 0xe4, 0xec, 0xce, 0xe6,
		0xc7, 0x00, 0x39, 0xb6, 0x67, 0xf8, 0xc6, 0x2d, 0xac, 0x18, 0x96, 0xb8, 0xdd, 0x21, 0x07, 0x8c,
		0xac, 0x2c, 0x89, 0x91, 0xb6, 0xe5, 0x07, 0xda, 0x16, 0x1e, 0xa8, 0x53, 0xda, 0x24, 0x81, 0x67,
		0x64, 0x49, 0x8c, 0x04, 0xda, 0x17, 0xa0, 0xac, 0xdb, 0x63, 0xd2, 0x75, 0x31, 0x3d, 0x52, 0x2f,
		0x52, 0x72, 0x89, 0xc9, 0x02, 0x15, 0xde, 0x4f, 0x4f, 0xae, 0x95, 0xca, 0x72, 0x89, 0xc9, 0x98,
		0xca, 0xe3, 0xb0, 0xa4, 0x0e, 0x06, 0x2e, 0x21, 0x17, 0x44, 0xec, 0x84, 0x50, 0x09, 0xc4, 0x54,
		0x71, 0xfd, 0x05, 0x28, 0x08, 0x3f, 0x90, 0x92, 0x4c, 0x3c, 0xa1, 0x38, 0xec, 0xd8, 0x9b, 0xde,
		0x28, 0xca, 0x05, 0x4b, 0x0c, 0x5e, 0x80, 0xb2, 0xe1, 0x29, 0x93, 0x5b, 0xf2, 0xf4, 0xf9, 0xf4,
		0x46, 0x41, 0x2e, 0x19, 0x5e, 0x70, 0xc3, 0x58, 0x7b, 0x3b, 0x0d, 0x95, 0xe8, 0x2d, 0x3f, 0xda,
		0x86, 0x82, 0x69, 0x6b, 0x2a, 0x0d, 0x2d, 0xf6, 0x89, 0x69, 0x23, 0xe6, 0xc3, 0xc0, 0xe6, 0x2e,
		0xd7, 0x97, 0x03, 0xe4, 0xfa, 0x3f, 0xa6, 0xa0, 0x20, 0xc4, 0x68, 0x0d, 0xb2, 0x8e, 0xea, 0x0f,
		0x29, 0x5d, 0x6e, 0x2b, 0x2d, 0xa5, 0x64, 0xfa, 0x4c, 0xe4, 0x9e, 0xa3, 0x5a, 0x34, 0x04, 0xb8,
		0x9c, 0x3c, 0x93, 0x75, 0x35, 0xb1, 0xaa, 0xd3, 0xe3, 0x87, 0x3d, 0x1a, 0x61, 0xcb, 0xf7, 0xc4,
		0xba, 0x72, 0x79, 0x93, 0x8b, 0xd1, 0x53, 0xb0, 0xec, 0xbb, 0xaa, 0x61, 0x46, 0x74, 0xb3, 0x54,
		0x57, 0x12, 0x03, 0x81, 0x72, 0x1d, 0xce, 0x0a, 0x5e, 0x1d, 0xfb, 0xaa, 0x36, 0xc4, 0xfa, 0x04,
		0x94, 0xa7, 0xd7, 0x0c, 0x67, 0xb8, 0xc2, 0x36, 0x1f, 0x17, 0xd8, 0xda, 0x77, 0x53, 0xb0, 0x2c,
		0x0e, 0x4c, 0x7a, 0xe0, 0xac, 0x3d, 0x00, 0xd5, 0xb2, 0x6c, 0x3f, 0xec, 0xae, 0xd9, 0x50, 0x9e,
		0xc1, 0x6d, 0x36, 0x02, 0x90, 0x1c, 0x22, 0x58, 0x1f, 0x01, 0x4c, 0x46, 0x4e, 0x74, 0xdb, 0x39,
		0x28, 0xf1, 0x4f, 0x38, 0xf4, 0x3b, 0x20, 0x3b, 0x62, 0x03, 0x13, 0x91, 0x93, 0x15, 0x5a, 0x85,
		0xdc, 0x11, 0x1e, 0x18, 0x16, 0xbf, 0x98, 0x65, 0x0f, 0xe2, 0x22, 0x24, 0x1b, 0x5c, 0x84, 0x6c,
		0x7d, 0x0e, 0x56, 0x34, 0x7b, 0x34, 0x6d, 0xee, 0x96, 0x34, 0x75, 0xcc, 0xf7, 0x3e, 0x9d, 0x7a,
		0x05, 0x26, 0x2d, 0xe6, 0x0f, 0x52, 0xa9, 0xdf, 0x4d, 0x67, 0x76, 0xba, 0x5b, 0x5f, 0x4f, 0xaf,
		0xef, 0x30, 0x68, 0x57, 0xcc, 0x54, 0xc6, 0x7d, 0x13, 0x6b, 0xc4, 0x7a, 0xf8, 0xda, 0x06, 0x7c,
		0x7c, 0x60, 0xf8, 0xc3, 0xf1, 0xd1, 0xa6, 0x66, 0x8f, 0x2e, 0x0e, 0xec, 0x81, 0x3d, 0xf9, 0xf4,
		0x49, 0x9e, 0xe8, 0x03, 0xfd, 0x8f, 0x7f, 0xfe, 0x2c, 0x06, 0xd2, 0xf5, 0xd8, 0x6f, 0xa5, 0xf5,
		0x7d, 0x58, 0xe1, 0xca, 0x0a, 0xfd, 0xfe, 0xc2, 0x4e, 0x11, 0xe8, 0x9e, 0x77, 0x58, 0xd5, 0x6f,
		0xbc, 0x43, 0xcb, 0xb5, 0xbc, 0xcc, 0xa1, 0x64, 0x8c, 0x1d, 0x34, 0xea, 0x32, 0xdc, 0x1f, 0xe1,
		0x63, 0x5b, 0x13, 0xbb, 0x31, 0x8c, 0xdf, 0xe1, 0x8c, 0x2b, 0x21, 0xc6, 0x1e, 0x87, 0xd6, 0x9b,
		0xb0, 0x78, 0x1a, 0xae, 0xbf, 0xe3, 0x5c, 0x65, 0x1c, 0x26, 0xd9, 0x81, 0x25, 0x4a, 0xa2, 0x8d,
		0x3d, 0xdf, 0x1e, 0xd1, 0xbc, 0x77, 0x6f, 0x9a, 0xbf, 0x7f, 0x87, 0xed, 0x95, 0x0a, 0x81, 0x35,
		0x03, 0x54, 0xbd, 0x0e, 0xf4, 0x93, 0x93, 0x8e, 0x35, 0x33, 0x86, 0xe1, 0x1f, 0xb8, 0x21, 0x81,
		0x7e, 0xfd, 0xb3, 0xb0, 0x4a, 0xfe, 0xa7, 0x69, 0x29, 0x6c, 0x49, 0xfc, 0x85, 0x57, 0xf5, 0xbb,
		0xaf, 0xb3, 0xed, 0xb8, 0x12, 0x10, 0x84, 0x6c, 0x0a, 0xad, 0xe2, 0x00, 0xfb, 0x3e, 0x76, 0x3d,
		0x45, 0x35, 0xe7, 0x99, 0x17, 0xba, 0x31, 0xa8, 0x7e, 0xe9, 0xdd, 0xe8, 0x2a, 0xee, 0x30, 0x64,
		0xc3, 0x34, 0xeb, 0x87, 0x70, 0x66, 0x4e, 0x54, 0x24, 0xe0, 0x7c, 0x83, 0x73, 0xae, 0xce, 0x44,
		0x06, 0xa1, 0xed, 0x82, 0x90, 0x07, 0x6b, 0x99, 0x80, 0xf3, 0xb7, 0x39, 0x27, 0xe2, 0x58, 0xb1,
		0xa4, 0x84, 0xf1, 0x05, 0x58, 0xbe, 0x85, 0xdd, 0x23, 0xdb, 0xe3, 0xb7, 0x34, 0x09, 0xe8, 0xde,
		0xe4, 0x74, 0x4b, 0x1c, 0x48, 0xaf, 0x6d, 0x08, 0xd7, 0x55, 0x28, 0xf4, 0x55, 0x0d, 0x27, 0xa0,
		0xf8, 0x32, 0xa7, 0x58, 0x20, 0xfa, 0x04, 0xda, 0x80, 0xf2, 0xc0, 0xe6, 0x95, 0x29, 0x1e, 0xfe,
		0x16, 0x87, 0x97, 0x04, 0x86, 0x53, 0x38, 0xb6, 0x33, 0x36, 0x49, 0xd9, 0x8a, 0xa7, 0xf8, 0x1d,
		0x41, 0x21, 0x30, 0x9c, 0xe2, 0x14, 0x6e, 0xfd, 0x8a, 0xa0, 0xf0, 0x42, 0xfe, 0x7c, 0x1e, 0x4a,
		0xb6, 0x65, 0x1e, 0xdb, 0x56, 0x12, 0x23, 0xbe, 0xca, 0x19, 0x80, 0x43, 0x08, 0xc1, 0x35, 0x28,
		0x26, 0x5d, 0x88, 0xaf, 0xbd, 0x2b, 0xb6, 0x87, 0x58, 0x81, 0x1d, 0x58, 0x12, 0x09, 0xca, 0xb0,
		0xad, 0x04, 0x14, 0xbf, 0xc7, 0x29, 0x2a, 0x21, 0x18, 0x9f, 0x86, 0x8f, 0x3d, 0x7f, 0x80, 0x93,
		0x90, 0xbc, 0x2d, 0xa6, 0xc1, 0x21, 0xdc, 0x95, 0x47, 0xd8, 0xd2, 0x86, 0xc9, 0x18, 0x7e, 0x5f,
		0xb8, 0x52, 0x60, 0x08, 0x45, 0x13, 0x16, 0x47, 0xaa, 0xeb, 0x0d, 0x55, 0x33, 0xd1, 0x72, 0xfc,
		0x01, 0xe7, 0x28, 0x07, 0x20, 0xee, 0x91, 0xb1, 0x75, 0x1a, 0x9a, 0xaf, 0x0b, 0x8f, 0x84, 0x60,
		0x7c, 0xeb, 0x79, 0x3e, 0xbd, 0xd2, 0x3a, 0x0d, 0xdb, 0x1f, 0x8a, 0xad, 0xc7, 0xb0, 0x7b, 0x61,
		0xc6, 0x6b, 0x50, 0xf4, 0x8c, 0x3b, 0x89, 0x68, 0xfe, 0x48, 0xac, 0x34, 0x05, 0x10, 0xf0, 0xcb,
		0x70, 0x76, 0x6e, 0x99, 0x48, 0x40, 0xf6, 0xc7, 0x9c, 0x6c, 0x6d, 0x4e, 0xa9, 0xe0, 0x29, 0xe1,
		0xb4, 0x94, 0x7f, 0x22, 0x52, 0x02, 0x9e, 0xe2, 0xea, 0x92, 0xb3, 0x82, 0xa7, 0xf6, 0x4f, 0xe7,
		0xb5, 0x3f, 0x15, 0x5e, 0x63, 0xd8, 0x88, 0xd7, 0x0e, 0x60, 0x8d, 0x33, 0x9e, 0x6e, 0x5d, 0xff,
		0x4c, 0x24, 0x56, 0x86, 0x3e, 0x8c, 0xae, 0xee, 0xe7, 0x60, 0x3d, 0x70, 0xa7, 0x68, 0x4a, 0x3d,
		0x65, 0xa4, 0x3a, 0x09, 0x98, 0xbf, 0xc1, 0x99, 0x45, 0xc6, 0x0f, 0xba, 0x5a, 0x6f, 0x4f, 0x75,
		0x08, 0xf9, 0x4b, 0x50, 0x15, 0xe4, 0x63, 0xcb, 0xc5, 0x9a, 0x3d, 0xb0, 0x8c, 0x3b, 0x58, 0x4f,
		0x40, 0xfd, 0xe7, 0x53, 0x4b, 0x75, 0x18, 0x82, 0x13, 0xe6, 0x36, 0x48, 0x41, 0xaf, 0xa2, 0x18,
		0x23, 0xc7, 0x76, 0xfd, 0x18, 0xc6, 0x6f, 0x8a, 0x95, 0x0a, 0x70, 0x6d, 0x0a, 0xab, 0xb7, 0xa0,
		0x42, 0x1f, 0x93, 0x86, 0xe4, 0x5f, 0x70, 0xa2, 0xc5, 0x09, 0x8a, 0x27, 0x0e, 0xcd, 0x1e, 0x39,
		0xaa, 0x9b, 0x24, 0xff, 0xfd, 0xa5, 0x48, 0x1c, 0x1c, 0xc2, 0x13, 0x87, 0x7f, 0xec, 0x60, 0x52,
		0xed, 0x13, 0x30, 0x7c, 0x4b, 0x24, 0x0e, 0x81, 0xe1, 0x14, 0xa2, 0x61, 0x48, 0x40, 0xf1, 0x57,
		0x82, 0x42, 0x60, 0x08, 0xc5, 0x67, 0x26, 0x85, 0xd6, 0xc5, 0x03, 0xc3, 0xf3, 0x5d, 0xd6, 0x0a,
		0xdf, 0x9b, 0xea, 0xaf, 0xdf, 0x8d, 0x36, 0x61, 0x72, 0x08, 0x4a, 0x32, 0x11, 0xbf, 0x42, 0xa5,
		0x27, 0xa5, 0x78, 0xc3, 0xbe, 0x2d, 0x32, 0x51, 0x08, 0xc6, 0xf6, 0xe7, 0xd2, 0x54, 0xaf, 0x82,
		0xe2, 0x7e, 0x08, 0x53, 0xfd, 0x99, 0xf7, 0x39, 0x57, 0xb4, 0x55, 0xa9, 0xef, 0x92, 0x00, 0x8a,
		0x36, 0x14, 0xf1, 0x64, 0xaf, 0xbf, 0x1f, 0xc4, 0x50, 0xa4, 0x9f, 0xa8, 0x5f, 0x87, 0xc5, 0x48,
		0x33, 0x11, 0x4f, 0xf5, 0xb3, 0x9c, 0xaa, 0x1c, 0xee, 0x25, 0xea, 0x97, 0x21, 0x4b, 0x1a, 0x83,
		0x78, 0xf8, 0xcf, 0x71, 0x38, 0x55, 0xaf, 0x7f, 0x12, 0x0a, 0xa2, 0x21, 0x88, 0x87, 0xfe, 0x3c,
		0x87, 0x06, 0x10, 0x02, 0x17, 0xcd, 0x40, 0x3c, 0xfc, 0x17, 0x04, 0x5c, 0x40, 0x08, 0x3c, 0xb9,
		0x0b, 0xff, 0xf6, 0x17, 0xb3, 0x3c, 0xa1, 0x0b, 0xdf, 0x5d, 0x83, 0x05, 0xde, 0x05, 0xc4, 0xa3,
		0xbf, 0xc0, 0x5f, 0x2e, 0x10, 0xf5, 0x67, 0x21, 0x97, 0xd0, 0xe1, 0xbf, 0xc4, 0xa1, 0x4c, 0xbf,
		0xde, 0x84, 0x52, 0xa8, 0xf2, 0xc7, 0xc3, 0x7f, 0x99, 0xc3, 0xc3, 0x28, 0x62, 0x3a, 0xaf, 0xfc,
		0xf1, 0x04, 0xbf, 0x22, 0x4c, 0xe7, 0x08, 0xe2, 0x36, 0x51, 0xf4, 0xe3, 0xd1, 0xbf, 0x2a, 0xbc,
		0x2e, 0x20, 0xf5, 0xe7, 0xa1, 0x18, 0x24, 0xf2, 0x78, 0xfc, 0xaf, 0x71, 0xfc, 0x04, 0x43, 0x3c,
		0x10, 0x2a, 0x24, 0xf1, 0x14, 0xbf, 0x2e, 0x3c, 0x10, 0x42, 0x91, 0x6d, 0x34, 0xdd, 0x1c, 0xc4,
		0x33, 0xfd, 0x86, 0xd8, 0x46, 0x53, 0xbd, 0x01, 0x59, 0x4d, 0x9a, 0x4f, 0xe3, 0x29, 0x7e, 0x53,
		0xac, 0x26, 0xd5, 0x27, 0x66, 0x4c, 0x57, 0xdb, 0x78, 0x8e, 0xdf, 0x12, 0x66, 0x4c, 0x15, 0xdb,
		0x7a, 0x17, 0xd0, 0x6c, 0xa5, 0x8d, 0xe7, 0xfb, 0x22, 0xe7, 0x5b, 0x9e, 0x29, 0xb4, 0xf5, 0x17,
		0x61, 0x6d, 0x7e, 0x95, 0x8d, 0x67, 0xfd, 0xd2, 0xfb, 0x53, 0xe7, 0xa2, 0x70, 0x91, 0xad, 0x1f,
		0x4c, 0xd2, 0x75, 0xb8, 0xc2, 0xc6, 0xd3, 0xbe, 0xf1, 0x7e, 0x34, 0x63, 0x87, 0x0b, 0x6c, 0xbd,
		0x01, 0x30, 0x29, 0x6e, 0xf1, 0x5c, 0x6f, 0x72, 0xae, 0x10, 0x88, 0x6c, 0x0d, 0x5e, 0xdb, 0xe2,
		0xf1, 0x5f, 0x16, 0x5b, 0x83, 0x23, 0xc8, 0xd6, 0x10, 0x65, 0x2d, 0x1e, 0xfd, 0x96, 0xd8, 0x1a,
		0x02, 0x42, 0x22, 0x3b, 0x54, 0x39, 0xe2, 0x19, 0xbe, 0x2a, 0x22, 0x3b, 0x84, 0xaa, 0x5f, 0x83,
		0x82, 0x35, 0x36, 0x4d, 0x12, 0xa0, 0xe8, 0xde, 0x3f, 0x10, 0xab, 0xfe, 0xfb, 0x5d, 0x6e, 0x81,
		0x00, 0xd4, 0x2f, 0x43, 0x0e, 0x8f, 0x8e, 0xb0, 0x1e, 0x87, 0xfc, 0x8f, 0xbb, 0x22, 0x29, 0x11,
		0xed, 0xfa, 0xf3, 0x00, 0xec, 0x68, 0x4f, 0x3f, 0x5b, 0xc5, 0x60, 0xff, 0xf3, 0x2e, 0xff, 0xe9,
		0xc6, 0x04, 0x32, 0x21, 0x60, 0x3f, 0x04, 0xb9, 0x37, 0xc1, 0xbb, 0x51, 0x02, 0x3a, 0xeb, 0xab,
		0xb0, 0x70, 0xc3, 0xb3, 0x2d, 0x5f, 0x1d, 0xc4, 0xa1, 0xff, 0x8b, 0xa3, 0x85, 0x3e, 0x71, 0xd8,
		0xc8, 0x76, 0xb1, 0xaf, 0x0e, 0xbc, 0x38, 0xec, 0x7f, 0x73, 0x6c, 0x00, 0x20, 0x60, 0x4d, 0xf5,
		0xfc, 0x24, 0xf3, 0xfe, 0xbe, 0x00, 0x0b, 0x00, 0x31, 0x9a, 0xfc, 0x7f, 0x13, 0x1f, 0xc7, 0x61,
		0xdf, 0x13, 0x46, 0x73, 0xfd, 0xfa, 0x27, 0xa1, 0x48, 0xfe, 0x65, 0xbf, 0xc7, 0x8a, 0x01, 0xff,
		0x0f, 0x07, 0x4f, 0x10, 0xe4, 0xcd, 0x9e, 0xaf, 0xfb, 0x46, 0xbc, 0xb3, 0xff, 0x97, 0xaf, 0xb4,
		0xd0, 0xaf, 0x37, 0xa0, 0xe4, 0xf9, 0xba, 0x3e, 0xe6, 0xfd, 0x55, 0x0c, 0xfc, 0xff, 0xee, 0x06,
		0x47, 0xee, 0x00, 0xb3, 0xd5, 0x9a, 0x7f, 0x7b, 0x08, 0x3b, 0xf6, 0x8e, 0xcd, 0xee, 0x0d, 0x5f,
		0xa9, 0xc5, 0x5f, 0x00, 0xc2, 0xff, 0xa7, 0xa0, 0xec, 0x60, 0xd7, 0xb3, 0x2d, 0x7e, 0x0d, 0x98,
		0x1d, 0xa9, 0x86, 0xb5, 0x7e, 0xba, 0xbb, 0xc3, 0xda, 0x08, 0xf2, 0x5d, 0x4a, 0x82, 0x10, 0x64,
		0xf7, 0x43, 0xbf, 0x49, 0xa2, 0xbf, 0xb9, 0x7c, 0x0a, 0x8a, 0x0d, 0x5d, 0x77, 0xb1, 0xe7, 0x61,
		0x8f, 0x7f, 0x70, 0x58, 0xdc, 0x24, 0xaf, 0xd9, 0xe4, 0x62, 0x79, 0x32, 0x8e, 0x1e, 0x84, 0xe2,
		0x01, 0x36, 0xb1, 0x33, 0xb4, 0x2d, 0xf1, 0x39, 0x61, 0x22, 0xa8, 0x67, 0xdf, 0xfb, 0xca, 0xb9,
		0x54, 0xed, 0x2a, 0x2c, 0x70, 0x00, 0x5a, 0x83, 0xfc, 0x3e, 0xfb, 0x71, 0x58, 0x8a, 0x7e, 0x1f,
		0xe0, 0x4f, 0x44, 0xde, 0xf3, 0x5d, 0x8c, 0x7d, 0x7e, 0x41, 0xcb, 0x9f, 0xb6, 0x0a, 0xef, 0x7d,
		0xef, 0xe1, 0xd4, 0x0f, 0xbe, 0xf7, 0x70, 0xea, 0x87, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x27,
		0x4d, 0x75, 0x14, 0x32, 0x00, 0x00,
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
func (this *Person) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&main.Person{")
	if this.Name != nil {
		s = append(s, "Name: "+valueToGoStringPerson(this.Name, "string")+",\n")
	}
	if this.Addresses != nil {
		s = append(s, "Addresses: "+fmt.Sprintf("%#v", this.Addresses)+",\n")
	}
	if this.Telephone != nil {
		s = append(s, "Telephone: "+valueToGoStringPerson(this.Telephone, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Address) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Address{")
	if this.Number != nil {
		s = append(s, "Number: "+valueToGoStringPerson(this.Number, "int64")+",\n")
	}
	if this.Street != nil {
		s = append(s, "Street: "+valueToGoStringPerson(this.Street, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPerson(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedPerson(r randyPerson, easy bool) *Person {
	this := &Person{}
	if r.Intn(10) != 0 {
		v1 := string(randStringPerson(r))
		this.Name = &v1
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Addresses = make([]*Address, v2)
		for i := 0; i < v2; i++ {
			this.Addresses[i] = NewPopulatedAddress(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v3 := string(randStringPerson(r))
		this.Telephone = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPerson(r, 4)
	}
	return this
}

func NewPopulatedAddress(r randyPerson, easy bool) *Address {
	this := &Address{}
	if r.Intn(10) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.Number = &v4
	}
	if r.Intn(10) != 0 {
		v5 := string(randStringPerson(r))
		this.Street = &v5
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPerson(r, 3)
	}
	return this
}

type randyPerson interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePerson(r randyPerson) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPerson(r randyPerson) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RunePerson(r)
	}
	return string(tmps)
}
func randUnrecognizedPerson(r randyPerson, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPerson(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPerson(dAtA []byte, r randyPerson, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePerson(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_person_134438c8f87a957d) }

var fileDescriptor_person_134438c8f87a957d = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0x94, 0xcb, 0xc5,
	0x16, 0x00, 0x36, 0x44, 0x48, 0x88, 0x8b, 0xc5, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xcc, 0x16, 0xd2, 0xe6, 0xe2, 0x74, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x4e,
	0x2d, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd5, 0x03, 0x59, 0xa3, 0x07, 0x15, 0x0e,
	0x42, 0xc8, 0x0b, 0xc9, 0x70, 0x71, 0x86, 0xa4, 0xe6, 0xa4, 0x16, 0x64, 0xe4, 0xe7, 0xa5, 0x4a,
	0x30, 0x83, 0x4d, 0x41, 0x08, 0x58, 0xb1, 0x7c, 0x58, 0x20, 0xcf, 0xa8, 0x64, 0xc9, 0xc5, 0x0e,
	0xd5, 0x20, 0x24, 0xc6, 0xc5, 0xe6, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x04, 0xb6, 0x91, 0x39, 0x08,
	0xca, 0x03, 0x89, 0x07, 0x97, 0x14, 0xa5, 0xa6, 0x96, 0x48, 0x30, 0x81, 0xcd, 0x80, 0xf2, 0x9c,
	0x38, 0x3e, 0x3c, 0x94, 0x63, 0xfc, 0xf1, 0x50, 0x8e, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0xe6, 0x3e, 0x1b, 0xf7, 0x00, 0x00, 0x00,
}
