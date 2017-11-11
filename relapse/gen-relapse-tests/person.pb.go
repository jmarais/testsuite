// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	Person
	Address
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

type Person struct {
	Name             *string    `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Addresses        []*Address `protobuf:"bytes,2,rep,name=Addresses" json:"Addresses,omitempty"`
	Telephone        *string    `protobuf:"bytes,3,opt,name=Telephone" json:"Telephone,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptorPerson, []int{0} }

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
	Number           *int64  `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	Street           *string `protobuf:"bytes,2,opt,name=Street" json:"Street,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptorPerson, []int{1} }

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
func (this *Person) Description() (desc *descriptor.FileDescriptorSet) {
	return PersonDescription()
}
func PersonDescription() (desc *descriptor.FileDescriptorSet) {
	d := &descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3755 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x6b, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0x9f, 0x22, 0x0f, 0x29, 0x6a, 0x74, 0x25, 0x6b, 0xb9, 0xf2, 0x63, 0x77, 0x69, 0x3b,
		0x2b, 0xdb, 0x8d, 0x36, 0x58, 0x7b, 0xd7, 0xde, 0xd9, 0x26, 0x2e, 0x45, 0x71, 0x15, 0x6e, 0x25,
		0x91, 0x19, 0x4a, 0xf1, 0x23, 0x28, 0x06, 0xa3, 0x99, 0x4b, 0x72, 0x76, 0x87, 0x33, 0x93, 0x99,
		0xe1, 0xae, 0xb5, 0xe8, 0x0f, 0x17, 0xee, 0x03, 0x41, 0xd1, 0x77, 0x81, 0x24, 0xae, 0xe3, 0xa6,
		0x05, 0x5a, 0xb7, 0xe9, 0x33, 0x4d, 0x9b, 0x3e, 0x7e, 0xf5, 0x4f, 0xda, 0xfe, 0x2a, 0x90, 0xff,
		0xfd, 0xd1, 0xa0, 0x06, 0xfa, 0x72, 0x1b, 0xb7, 0x35, 0xd0, 0x00, 0xfe, 0x13, 0xdc, 0xd7, 0x70,
		0x86, 0xa4, 0x76, 0xa8, 0x00, 0x76, 0x7e, 0x49, 0x73, 0xee, 0xf9, 0xbe, 0x39, 0xf7, 0xdc, 0x73,
		0xcf, 0x39, 0xf7, 0x0e, 0xe1, 0xbb, 0xd7, 0xe0, 0x7c, 0xdf, 0x71, 0xfa, 0x16, 0xbe, 0xe4, 0x7a,
		0x4e, 0xe0, 0x1c, 0x8d, 0x7a, 0x97, 0x0c, 0xec, 0xeb, 0x9e, 0xe9, 0x06, 0x8e, 0xb7, 0x49, 0x65,
		0x68, 0x89, 0x69, 0x6c, 0x0a, 0x8d, 0xda, 0x1e, 0x2c, 0xdf, 0x30, 0x2d, 0xbc, 0x1d, 0x2a, 0x76,
		0x71, 0x80, 0x9e, 0x87, 0x6c, 0xcf, 0xb4, 0x70, 0x35, 0x75, 0x3e, 0xb3, 0x51, 0xba, 0xfc, 0xf8,
		0xe6, 0x04, 0x68, 0x33, 0x8e, 0xe8, 0x10, 0xb1, 0x42, 0x11, 0xb5, 0x77, 0xb2, 0xb0, 0x32, 0x63,
		0x14, 0x21, 0xc8, 0xda, 0xda, 0x90, 0x30, 0xa6, 0x36, 0x8a, 0x0a, 0xfd, 0x1f, 0x55, 0x61, 0xc1,
		0xd5, 0xf4, 0xdb, 0x5a, 0x1f, 0x57, 0xd3, 0x54, 0x2c, 0x1e, 0xd1, 0xa3, 0x00, 0x06, 0x76, 0xb1,
		0x6d, 0x60, 0x5b, 0x3f, 0xae, 0x66, 0xce, 0x67, 0x36, 0x8a, 0x4a, 0x44, 0x82, 0x9e, 0x86, 0x65,
		0x77, 0x74, 0x64, 0x99, 0xba, 0x1a, 0x51, 0x83, 0xf3, 0x99, 0x8d, 0x9c, 0x22, 0xb1, 0x81, 0xed,
		0xb1, 0xf2, 0x45, 0x58, 0xba, 0x8b, 0xb5, 0xdb, 0x51, 0xd5, 0x12, 0x55, 0xad, 0x10, 0x71, 0x44,
		0xb1, 0x01, 0xe5, 0x21, 0xf6, 0x7d, 0xad, 0x8f, 0xd5, 0xe0, 0xd8, 0xc5, 0xd5, 0x2c, 0x9d, 0xfd,
		0xf9, 0xa9, 0xd9, 0x4f, 0xce, 0xbc, 0xc4, 0x51, 0x07, 0xc7, 0x2e, 0x46, 0x75, 0x28, 0x62, 0x7b,
		0x34, 0x64, 0x0c, 0xb9, 0x13, 0xfc, 0xd7, 0xb4, 0x47, 0xc3, 0x49, 0x96, 0x02, 0x81, 0x71, 0x8a,
		0x05, 0x1f, 0x7b, 0x77, 0x4c, 0x1d, 0x57, 0xf3, 0x94, 0xe0, 0xe2, 0x14, 0x41, 0x97, 0x8d, 0x4f,
		0x72, 0x08, 0x1c, 0x6a, 0x40, 0x11, 0xbf, 0x1a, 0x60, 0xdb, 0x37, 0x1d, 0xbb, 0xba, 0x40, 0x49,
		0x9e, 0x98, 0xb1, 0x8a, 0xd8, 0x32, 0x26, 0x29, 0xc6, 0x38, 0x74, 0x15, 0x16, 0x1c, 0x37, 0x30,
		0x1d, 0xdb, 0xaf, 0x16, 0xce, 0xa7, 0x36, 0x4a, 0x97, 0x1f, 0x9e, 0x19, 0x08, 0x6d, 0xa6, 0xa3,
		0x08, 0x65, 0xd4, 0x02, 0xc9, 0x77, 0x46, 0x9e, 0x8e, 0x55, 0xdd, 0x31, 0xb0, 0x6a, 0xda, 0x3d,
		0xa7, 0x5a, 0xa4, 0x04, 0xe7, 0xa6, 0x27, 0x42, 0x15, 0x1b, 0x8e, 0x81, 0x5b, 0x76, 0xcf, 0x51,
		0x2a, 0x7e, 0xec, 0x19, 0xad, 0x41, 0xde, 0x3f, 0xb6, 0x03, 0xed, 0xd5, 0x6a, 0x99, 0x46, 0x08,
		0x7f, 0xaa, 0xfd, 0x75, 0x1e, 0x96, 0xe6, 0x09, 0xb1, 0xeb, 0x90, 0xeb, 0x91, 0x59, 0x56, 0xd3,
		0xa7, 0xf1, 0x01, 0xc3, 0xc4, 0x9d, 0x98, 0xff, 0x01, 0x9d, 0x58, 0x87, 0x92, 0x8d, 0xfd, 0x00,
		0x1b, 0x2c, 0x22, 0x32, 0x73, 0xc6, 0x14, 0x30, 0xd0, 0x74, 0x48, 0x65, 0x7f, 0xa0, 0x90, 0x7a,
		0x09, 0x96, 0x42, 0x93, 0x54, 0x4f, 0xb3, 0xfb, 0x22, 0x36, 0x2f, 0x25, 0x59, 0xb2, 0xd9, 0x14,
		0x38, 0x85, 0xc0, 0x94, 0x0a, 0x8e, 0x3d, 0xa3, 0x6d, 0x00, 0xc7, 0xc6, 0x4e, 0x4f, 0x35, 0xb0,
		0x6e, 0x55, 0x0b, 0x27, 0x78, 0xa9, 0x4d, 0x54, 0xa6, 0xbc, 0xe4, 0x30, 0xa9, 0x6e, 0xa1, 0x6b,
		0xe3, 0x50, 0x5b, 0x38, 0x21, 0x52, 0xf6, 0xd8, 0x26, 0x9b, 0x8a, 0xb6, 0x43, 0xa8, 0x78, 0x98,
		0xc4, 0x3d, 0x36, 0xf8, 0xcc, 0x8a, 0xd4, 0x88, 0xcd, 0xc4, 0x99, 0x29, 0x1c, 0xc6, 0x26, 0xb6,
		0xe8, 0x45, 0x1f, 0xd1, 0x63, 0x10, 0x0a, 0x54, 0x1a, 0x56, 0x40, 0xb3, 0x50, 0x59, 0x08, 0xf7,
		0xb5, 0x21, 0x5e, 0xbf, 0x07, 0x95, 0xb8, 0x7b, 0xd0, 0x2a, 0xe4, 0xfc, 0x40, 0xf3, 0x02, 0x1a,
		0x85, 0x39, 0x85, 0x3d, 0x20, 0x09, 0x32, 0xd8, 0x36, 0x68, 0x96, 0xcb, 0x29, 0xe4, 0x5f, 0xf4,
		0x63, 0xe3, 0x09, 0x67, 0xe8, 0x84, 0x3f, 0x36, 0xbd, 0xa2, 0x31, 0xe6, 0xc9, 0x79, 0xaf, 0x3f,
		0x07, 0x8b, 0xb1, 0x09, 0xcc, 0xfb, 0xea, 0xda, 0x4f, 0xc2, 0x83, 0x33, 0xa9, 0xd1, 0x4b, 0xb0,
		0x3a, 0xb2, 0x4d, 0x3b, 0xc0, 0x9e, 0xeb, 0x61, 0x12, 0xb1, 0xec, 0x55, 0xd5, 0x7f, 0x5d, 0x38,
		0x21, 0xe6, 0x0e, 0xa3, 0xda, 0x8c, 0x45, 0x59, 0x19, 0x4d, 0x0b, 0x9f, 0x2a, 0x16, 0xfe, 0x6d,
		0x41, 0x7a, 0xed, 0xb5, 0xd7, 0x5e, 0x4b, 0xd7, 0xbe, 0x94, 0x87, 0xd5, 0x59, 0x7b, 0x66, 0xe6,
		0xf6, 0x5d, 0x83, 0xbc, 0x3d, 0x1a, 0x1e, 0x61, 0x8f, 0x3a, 0x29, 0xa7, 0xf0, 0x27, 0x54, 0x87,
		0x9c, 0xa5, 0x1d, 0x61, 0xab, 0x9a, 0x3d, 0x9f, 0xda, 0xa8, 0x5c, 0x7e, 0x7a, 0xae, 0x5d, 0xb9,
		0xb9, 0x4b, 0x20, 0x0a, 0x43, 0xa2, 0x4f, 0x41, 0x96, 0xa7, 0x68, 0xc2, 0xf0, 0xd4, 0x7c, 0x0c,
		0x64, 0x2f, 0x29, 0x14, 0x87, 0x1e, 0x82, 0x22, 0xf9, 0xcb, 0x62, 0x23, 0x4f, 0x6d, 0x2e, 0x10,
		0x01, 0x89, 0x0b, 0xb4, 0x0e, 0x05, 0xba, 0x4d, 0x0c, 0x2c, 0x4a, 0x5b, 0xf8, 0x4c, 0x02, 0xcb,
		0xc0, 0x3d, 0x6d, 0x64, 0x05, 0xea, 0x1d, 0xcd, 0x1a, 0x61, 0x1a, 0xf0, 0x45, 0xa5, 0xcc, 0x85,
		0x9f, 0x25, 0x32, 0x74, 0x0e, 0x4a, 0x6c, 0x57, 0x99, 0xb6, 0x81, 0x5f, 0xa5, 0xd9, 0x33, 0xa7,
		0xb0, 0x8d, 0xd6, 0x22, 0x12, 0xf2, 0xfa, 0x5b, 0xbe, 0x63, 0x8b, 0xd0, 0xa4, 0xaf, 0x20, 0x02,
		0xfa, 0xfa, 0xe7, 0x26, 0x13, 0xf7, 0x23, 0xb3, 0xa7, 0x37, 0x19, 0x53, 0xb5, 0x6f, 0xa6, 0x21,
		0x4b, 0xf3, 0xc5, 0x12, 0x94, 0x0e, 0x5e, 0xee, 0x34, 0xd5, 0xed, 0xf6, 0xe1, 0xd6, 0x6e, 0x53,
		0x4a, 0xa1, 0x0a, 0x00, 0x15, 0xdc, 0xd8, 0x6d, 0xd7, 0x0f, 0xa4, 0x74, 0xf8, 0xdc, 0xda, 0x3f,
		0xb8, 0xfa, 0xac, 0x94, 0x09, 0x01, 0x87, 0x4c, 0x90, 0x8d, 0x2a, 0x3c, 0x73, 0x59, 0xca, 0x21,
		0x09, 0xca, 0x8c, 0xa0, 0xf5, 0x52, 0x73, 0xfb, 0xea, 0xb3, 0x52, 0x3e, 0x2e, 0x79, 0xe6, 0xb2,
		0xb4, 0x80, 0x16, 0xa1, 0x48, 0x25, 0x5b, 0xed, 0xf6, 0xae, 0x54, 0x08, 0x39, 0xbb, 0x07, 0x4a,
		0x6b, 0x7f, 0x47, 0x2a, 0x86, 0x9c, 0x3b, 0x4a, 0xfb, 0xb0, 0x23, 0x41, 0xc8, 0xb0, 0xd7, 0xec,
		0x76, 0xeb, 0x3b, 0x4d, 0xa9, 0x14, 0x6a, 0x6c, 0xbd, 0x7c, 0xd0, 0xec, 0x4a, 0xe5, 0x98, 0x59,
		0xcf, 0x5c, 0x96, 0x16, 0xc3, 0x57, 0x34, 0xf7, 0x0f, 0xf7, 0xa4, 0x0a, 0x5a, 0x86, 0x45, 0xf6,
		0x0a, 0x61, 0xc4, 0xd2, 0x84, 0xe8, 0xea, 0xb3, 0x92, 0x34, 0x36, 0x84, 0xb1, 0x2c, 0xc7, 0x04,
		0x57, 0x9f, 0x95, 0x50, 0xad, 0x01, 0x39, 0x1a, 0x5d, 0x08, 0x41, 0x65, 0xb7, 0xbe, 0xd5, 0xdc,
		0x55, 0xdb, 0x9d, 0x83, 0x56, 0x7b, 0xbf, 0xbe, 0x2b, 0xa5, 0xc6, 0x32, 0xa5, 0xf9, 0x99, 0xc3,
		0x96, 0xd2, 0xdc, 0x96, 0xd2, 0x51, 0x59, 0xa7, 0x59, 0x3f, 0x68, 0x6e, 0x4b, 0x99, 0x9a, 0x0e,
		0xab, 0xb3, 0xf2, 0xe4, 0xcc, 0x9d, 0x11, 0x59, 0xe2, 0xf4, 0x09, 0x4b, 0x4c, 0xb9, 0xa6, 0x96,
		0xf8, 0x5f, 0xd2, 0xb0, 0x32, 0xa3, 0x56, 0xcc, 0x7c, 0xc9, 0x0b, 0x90, 0x63, 0x21, 0xca, 0xaa,
		0xe7, 0x93, 0x33, 0x8b, 0x0e, 0x0d, 0xd8, 0xa9, 0x0a, 0x4a, 0x71, 0xd1, 0x0e, 0x22, 0x73, 0x42,
		0x07, 0x41, 0x28, 0xa6, 0x72, 0xfa, 0x4f, 0x4c, 0xe5, 0x74, 0x56, 0xf6, 0xae, 0xce, 0x53, 0xf6,
		0xa8, 0xec, 0x74, 0xb9, 0x3d, 0x37, 0x23, 0xb7, 0x5f, 0x87, 0xe5, 0x29, 0xa2, 0xb9, 0x73, 0xec,
		0xeb, 0x29, 0xa8, 0x9e, 0xe4, 0x9c, 0x84, 0x4c, 0x97, 0x8e, 0x65, 0xba, 0xeb, 0x93, 0x1e, 0xbc,
		0x70, 0xf2, 0x22, 0x4c, 0xad, 0xf5, 0xdb, 0x29, 0x58, 0x9b, 0xdd, 0x29, 0xce, 0xb4, 0xe1, 0x53,
		0x90, 0x1f, 0xe2, 0x60, 0xe0, 0x88, 0x6e, 0xe9, 0x63, 0x33, 0x6a, 0x30, 0x19, 0x9e, 0x5c, 0x6c,
		0x8e, 0x8a, 0x16, 0xf1, 0xcc, 0x49, 0xed, 0x1e, 0xb3, 0x66, 0xca, 0xd2, 0x2f, 0xa4, 0xe1, 0xc1,
		0x99, 0xe4, 0x33, 0x0d, 0x7d, 0x04, 0xc0, 0xb4, 0xdd, 0x51, 0xc0, 0x3a, 0x22, 0x96, 0x60, 0x8b,
		0x54, 0x42, 0x93, 0x17, 0x49, 0x9e, 0xa3, 0x20, 0x1c, 0xcf, 0xd0, 0x71, 0x60, 0x22, 0xaa, 0xf0,
		0xfc, 0xd8, 0xd0, 0x2c, 0x35, 0xf4, 0xd1, 0x13, 0x66, 0x3a, 0x15, 0x98, 0x9f, 0x00, 0x49, 0xb7,
		0x4c, 0x6c, 0x07, 0xaa, 0x1f, 0x78, 0x58, 0x1b, 0x9a, 0x76, 0x9f, 0x56, 0x90, 0x82, 0x9c, 0xeb,
		0x69, 0x96, 0x8f, 0x95, 0x25, 0x36, 0xdc, 0x15, 0xa3, 0x04, 0x41, 0x03, 0xc8, 0x8b, 0x20, 0xf2,
		0x31, 0x04, 0x1b, 0x0e, 0x11, 0xb5, 0x6f, 0x14, 0xa0, 0x14, 0xe9, 0xab, 0xd1, 0x05, 0x28, 0xdf,
		0xd2, 0xee, 0x68, 0xaa, 0x38, 0x2b, 0x31, 0x4f, 0x94, 0x88, 0xac, 0xc3, 0xcf, 0x4b, 0x9f, 0x80,
		0x55, 0xaa, 0xe2, 0x8c, 0x02, 0xec, 0xa9, 0xba, 0xa5, 0xf9, 0x3e, 0x75, 0x5a, 0x81, 0xaa, 0x22,
		0x32, 0xd6, 0x26, 0x43, 0x0d, 0x31, 0x82, 0xae, 0xc0, 0x0a, 0x45, 0x0c, 0x47, 0x56, 0x60, 0xba,
		0x16, 0x56, 0xc9, 0xe9, 0xcd, 0xa7, 0x95, 0x24, 0xb4, 0x6c, 0x99, 0x68, 0xec, 0x71, 0x05, 0x62,
		0x91, 0x8f, 0xb6, 0xe1, 0x11, 0x0a, 0xeb, 0x63, 0x1b, 0x7b, 0x5a, 0x80, 0x55, 0xfc, 0xf9, 0x91,
		0x66, 0xf9, 0xaa, 0x66, 0x1b, 0xea, 0x40, 0xf3, 0x07, 0xd5, 0x55, 0x42, 0xb0, 0x95, 0xae, 0xa6,
		0x94, 0xb3, 0x44, 0x71, 0x87, 0xeb, 0x35, 0xa9, 0x5a, 0xdd, 0x36, 0x3e, 0xad, 0xf9, 0x03, 0x24,
		0xc3, 0x1a, 0x65, 0xf1, 0x03, 0xcf, 0xb4, 0xfb, 0xaa, 0x3e, 0xc0, 0xfa, 0x6d, 0x75, 0x14, 0xf4,
		0x9e, 0xaf, 0x3e, 0x14, 0x7d, 0x3f, 0xb5, 0xb0, 0x4b, 0x75, 0x1a, 0x44, 0xe5, 0x30, 0xe8, 0x3d,
		0x8f, 0xba, 0x50, 0x26, 0x8b, 0x31, 0x34, 0xef, 0x61, 0xb5, 0xe7, 0x78, 0xb4, 0x34, 0x56, 0x66,
		0xa4, 0xa6, 0x88, 0x07, 0x37, 0xdb, 0x1c, 0xb0, 0xe7, 0x18, 0x58, 0xce, 0x75, 0x3b, 0xcd, 0xe6,
		0xb6, 0x52, 0x12, 0x2c, 0x37, 0x1c, 0x8f, 0x04, 0x54, 0xdf, 0x09, 0x1d, 0x5c, 0x62, 0x01, 0xd5,
		0x77, 0x84, 0x7b, 0xaf, 0xc0, 0x8a, 0xae, 0xb3, 0x39, 0x9b, 0xba, 0xca, 0xcf, 0x58, 0x7e, 0x55,
		0x8a, 0x39, 0x4b, 0xd7, 0x77, 0x98, 0x02, 0x8f, 0x71, 0x1f, 0x5d, 0x83, 0x07, 0xc7, 0xce, 0x8a,
		0x02, 0x97, 0xa7, 0x66, 0x39, 0x09, 0xbd, 0x02, 0x2b, 0xee, 0xf1, 0x34, 0x10, 0xc5, 0xde, 0xe8,
		0x1e, 0x4f, 0xc2, 0x9e, 0x83, 0x55, 0x77, 0xe0, 0x4e, 0xe3, 0x9e, 0x8a, 0xe2, 0x90, 0x3b, 0x70,
		0x27, 0x81, 0x4f, 0xd0, 0x03, 0xb7, 0x87, 0x75, 0x2d, 0xc0, 0x46, 0xf5, 0x4c, 0x54, 0x3d, 0x32,
		0x80, 0x2e, 0x81, 0xa4, 0xeb, 0x2a, 0xb6, 0xb5, 0x23, 0x0b, 0xab, 0x9a, 0x87, 0x6d, 0xcd, 0xaf,
		0x9e, 0x8b, 0x2a, 0x57, 0x74, 0xbd, 0x49, 0x47, 0xeb, 0x74, 0x10, 0x3d, 0x05, 0xcb, 0xce, 0xd1,
		0x2d, 0x9d, 0x85, 0xa4, 0xea, 0x7a, 0xb8, 0x67, 0xbe, 0x5a, 0x7d, 0x9c, 0xfa, 0x77, 0x89, 0x0c,
		0xd0, 0x80, 0xec, 0x50, 0x31, 0x7a, 0x12, 0x24, 0xdd, 0x1f, 0x68, 0x9e, 0x4b, 0x73, 0xb2, 0xef,
		0x6a, 0x3a, 0xae, 0x3e, 0xc1, 0x54, 0x99, 0x7c, 0x5f, 0x88, 0xc9, 0x96, 0xf0, 0xef, 0x9a, 0xbd,
		0x40, 0x30, 0x5e, 0x64, 0x5b, 0x82, 0xca, 0x38, 0xdb, 0x06, 0x48, 0xc4, 0x15, 0xb1, 0x17, 0x6f,
		0x50, 0xb5, 0x8a, 0x3b, 0x70, 0xa3, 0xef, 0x7d, 0x0c, 0x16, 0x89, 0xe6, 0xf8, 0xa5, 0x4f, 0xb2,
		0x86, 0xcc, 0x1d, 0x44, 0xde, 0xf8, 0xa1, 0xf5, 0xc6, 0x35, 0x19, 0xca, 0xd1, 0xf8, 0x44, 0x45,
		0x60, 0x11, 0x2a, 0xa5, 0x48, 0xb3, 0xd2, 0x68, 0x6f, 0x93, 0x36, 0xe3, 0x95, 0xa6, 0x94, 0x26,
		0xed, 0xce, 0x6e, 0xeb, 0xa0, 0xa9, 0x2a, 0x87, 0xfb, 0x07, 0xad, 0xbd, 0xa6, 0x94, 0x89, 0xf6,
		0xd5, 0xdf, 0x4a, 0x43, 0x25, 0x7e, 0x44, 0x42, 0x3f, 0x0a, 0x67, 0xc4, 0x7d, 0x86, 0x8f, 0x03,
		0xf5, 0xae, 0xe9, 0xd1, 0x2d, 0x33, 0xd4, 0x58, 0xf9, 0x0a, 0x17, 0x6d, 0x95, 0x6b, 0x75, 0x71,
		0xf0, 0xa2, 0xe9, 0x91, 0x0d, 0x31, 0xd4, 0x02, 0xb4, 0x0b, 0xe7, 0x6c, 0x47, 0xf5, 0x03, 0xcd,
		0x36, 0x34, 0xcf, 0x50, 0xc7, 0x37, 0x49, 0xaa, 0xa6, 0xeb, 0xd8, 0xf7, 0x1d, 0x56, 0xaa, 0x42,
		0x96, 0x87, 0x6d, 0xa7, 0xcb, 0x95, 0xc7, 0x39, 0xbc, 0xce, 0x55, 0x27, 0x02, 0x2c, 0x73, 0x52,
		0x80, 0x3d, 0x04, 0xc5, 0xa1, 0xe6, 0xaa, 0xd8, 0x0e, 0xbc, 0x63, 0xda, 0x18, 0x17, 0x94, 0xc2,
		0x50, 0x73, 0x9b, 0xe4, 0xf9, 0xa3, 0x39, 0x9f, 0xfc, 0x53, 0x06, 0xca, 0xd1, 0xe6, 0x98, 0x9c,
		0x35, 0x74, 0x5a, 0x47, 0x52, 0x34, 0xd3, 0x3c, 0x76, 0xdf, 0x56, 0x7a, 0xb3, 0x41, 0x0a, 0x8c,
		0x9c, 0x67, 0x2d, 0xab, 0xc2, 0x90, 0xa4, 0xb8, 0x93, 0xdc, 0x82, 0x59, 0x8b, 0x50, 0x50, 0xf8,
		0x13, 0xda, 0x81, 0xfc, 0x2d, 0x9f, 0x72, 0xe7, 0x29, 0xf7, 0xe3, 0xf7, 0xe7, 0xbe, 0xd9, 0xa5,
		0xe4, 0xc5, 0x9b, 0x5d, 0x75, 0xbf, 0xad, 0xec, 0xd5, 0x77, 0x15, 0x0e, 0x47, 0x67, 0x21, 0x6b,
		0x69, 0xf7, 0x8e, 0xe3, 0xa5, 0x88, 0x8a, 0xe6, 0x75, 0xfc, 0x59, 0xc8, 0xde, 0xc5, 0xda, 0xed,
		0x78, 0x01, 0xa0, 0xa2, 0x0f, 0x31, 0xf4, 0x2f, 0x41, 0x8e, 0xfa, 0x0b, 0x01, 0x70, 0x8f, 0x49,
		0x0f, 0xa0, 0x02, 0x64, 0x1b, 0x6d, 0x85, 0x84, 0xbf, 0x04, 0x65, 0x26, 0x55, 0x3b, 0xad, 0x66,
		0xa3, 0x29, 0xa5, 0x6b, 0x57, 0x20, 0xcf, 0x9c, 0x40, 0xb6, 0x46, 0xe8, 0x06, 0xe9, 0x01, 0xfe,
		0xc8, 0x39, 0x52, 0x62, 0xf4, 0x70, 0x6f, 0xab, 0xa9, 0x48, 0xe9, 0xe8, 0xf2, 0xfa, 0x50, 0x8e,
		0xf6, 0xc5, 0x1f, 0x4d, 0x4c, 0xfd, 0x4d, 0x0a, 0x4a, 0x91, 0x3e, 0x97, 0x34, 0x28, 0x9a, 0x65,
		0x39, 0x77, 0x55, 0xcd, 0x32, 0x35, 0x9f, 0x07, 0x05, 0x50, 0x51, 0x9d, 0x48, 0xe6, 0x5d, 0xb4,
		0x8f, 0xc4, 0xf8, 0xb7, 0x52, 0x20, 0x4d, 0xb6, 0x98, 0x13, 0x06, 0xa6, 0x7e, 0xa8, 0x06, 0xbe,
		0x99, 0x82, 0x4a, 0xbc, 0xaf, 0x9c, 0x30, 0xef, 0xc2, 0x0f, 0xd5, 0xbc, 0x7f, 0x4e, 0xc3, 0x62,
		0xac, 0x9b, 0x9c, 0xd7, 0xba, 0xcf, 0xc3, 0xb2, 0x69, 0xe0, 0xa1, 0xeb, 0x04, 0xd8, 0xd6, 0x8f,
		0x55, 0x0b, 0xdf, 0xc1, 0x56, 0xb5, 0x46, 0x13, 0xc5, 0xa5, 0xfb, 0xf7, 0xab, 0x9b, 0xad, 0x31,
		0x6e, 0x97, 0xc0, 0xe4, 0x95, 0xd6, 0x76, 0x73, 0xaf, 0xd3, 0x3e, 0x68, 0xee, 0x37, 0x5e, 0x56,
		0x0f, 0xf7, 0x7f, 0x7c, 0xbf, 0xfd, 0xe2, 0xbe, 0x22, 0x99, 0x13, 0x6a, 0x1f, 0xe2, 0x56, 0xef,
		0x80, 0x34, 0x69, 0x14, 0x3a, 0x03, 0xb3, 0xcc, 0x92, 0x1e, 0x40, 0x2b, 0xb0, 0xb4, 0xdf, 0x56,
		0xbb, 0xad, 0xed, 0xa6, 0xda, 0xbc, 0x71, 0xa3, 0xd9, 0x38, 0xe8, 0xb2, 0x1b, 0x88, 0x50, 0xfb,
		0x20, 0xbe, 0xa9, 0xdf, 0xc8, 0xc0, 0xca, 0x0c, 0x4b, 0x50, 0x9d, 0x9f, 0x1d, 0xd8, 0x71, 0xe6,
		0xe3, 0xf3, 0x58, 0xbf, 0x49, 0x4a, 0x7e, 0x47, 0xf3, 0x02, 0x7e, 0xd4, 0x78, 0x12, 0x88, 0x97,
		0xec, 0xc0, 0xec, 0x99, 0xd8, 0xe3, 0x17, 0x36, 0xec, 0x40, 0xb1, 0x34, 0x96, 0xb3, 0x3b, 0x9b,
		0x1f, 0x01, 0xe4, 0x3a, 0xbe, 0x19, 0x98, 0x77, 0xb0, 0x6a, 0xda, 0xe2, 0x76, 0x87, 0x1c, 0x30,
		0xb2, 0x8a, 0x24, 0x46, 0x5a, 0x76, 0x10, 0x6a, 0xdb, 0xb8, 0xaf, 0x4d, 0x68, 0x93, 0x04, 0x9e,
		0x51, 0x24, 0x31, 0x12, 0x6a, 0x5f, 0x80, 0xb2, 0xe1, 0x8c, 0x48, 0xd7, 0xc5, 0xf4, 0x48, 0xbd,
		0x48, 0x29, 0x25, 0x26, 0x0b, 0x55, 0x78, 0x3f, 0x3d, 0xbe, 0x56, 0x2a, 0x2b, 0x25, 0x26, 0x63,
		0x2a, 0x17, 0x61, 0x49, 0xeb, 0xf7, 0x3d, 0x42, 0x2e, 0x88, 0xd8, 0x09, 0xa1, 0x12, 0x8a, 0xa9,
		0xe2, 0xfa, 0x4d, 0x28, 0x08, 0x3f, 0x90, 0x92, 0x4c, 0x3c, 0xa1, 0xba, 0xec, 0xd8, 0x9b, 0xde,
		0x28, 0x2a, 0x05, 0x5b, 0x0c, 0x5e, 0x80, 0xb2, 0xe9, 0xab, 0xe3, 0x5b, 0xf2, 0xf4, 0xf9, 0xf4,
		0x46, 0x41, 0x29, 0x99, 0x7e, 0x78, 0xc3, 0x58, 0x7b, 0x3b, 0x0d, 0x95, 0xf8, 0x2d, 0x3f, 0xda,
		0x86, 0x82, 0xe5, 0xe8, 0x1a, 0x0d, 0x2d, 0xf6, 0x89, 0x69, 0x23, 0xe1, 0xc3, 0xc0, 0xe6, 0x2e,
		0xd7, 0x57, 0x42, 0xe4, 0xfa, 0x3f, 0xa6, 0xa0, 0x20, 0xc4, 0x68, 0x0d, 0xb2, 0xae, 0x16, 0x0c,
		0x28, 0x5d, 0x6e, 0x2b, 0x2d, 0xa5, 0x14, 0xfa, 0x4c, 0xe4, 0xbe, 0xab, 0xd9, 0x34, 0x04, 0xb8,
		0x9c, 0x3c, 0x93, 0x75, 0xb5, 0xb0, 0x66, 0xd0, 0xe3, 0x87, 0x33, 0x1c, 0x62, 0x3b, 0xf0, 0xc5,
		0xba, 0x72, 0x79, 0x83, 0x8b, 0xd1, 0xd3, 0xb0, 0x1c, 0x78, 0x9a, 0x69, 0xc5, 0x74, 0xb3, 0x54,
		0x57, 0x12, 0x03, 0xa1, 0xb2, 0x0c, 0x67, 0x05, 0xaf, 0x81, 0x03, 0x4d, 0x1f, 0x60, 0x63, 0x0c,
		0xca, 0xd3, 0x6b, 0x86, 0x33, 0x5c, 0x61, 0x9b, 0x8f, 0x0b, 0x6c, 0xed, 0xdb, 0x29, 0x58, 0x16,
		0x07, 0x26, 0x23, 0x74, 0xd6, 0x1e, 0x80, 0x66, 0xdb, 0x4e, 0x10, 0x75, 0xd7, 0x74, 0x28, 0x4f,
		0xe1, 0x36, 0xeb, 0x21, 0x48, 0x89, 0x10, 0xac, 0x0f, 0x01, 0xc6, 0x23, 0x27, 0xba, 0xed, 0x1c,
		0x94, 0xf8, 0x27, 0x1c, 0xfa, 0x1d, 0x90, 0x1d, 0xb1, 0x81, 0x89, 0xc8, 0xc9, 0x0a, 0xad, 0x42,
		0xee, 0x08, 0xf7, 0x4d, 0x9b, 0x5f, 0xcc, 0xb2, 0x07, 0x71, 0x11, 0x92, 0x0d, 0x2f, 0x42, 0xb6,
		0x3e, 0x07, 0x2b, 0xba, 0x33, 0x9c, 0x34, 0x77, 0x4b, 0x9a, 0x38, 0xe6, 0xfb, 0x9f, 0x4e, 0xbd,
		0x02, 0xe3, 0x16, 0xf3, 0x7b, 0xa9, 0xd4, 0x6f, 0xa7, 0x33, 0x3b, 0x9d, 0xad, 0xaf, 0xa5, 0xd7,
		0x77, 0x18, 0xb4, 0x23, 0x66, 0xaa, 0xe0, 0x9e, 0x85, 0x75, 0x62, 0x3d, 0x7c, 0xf1, 0x22, 0x7c,
		0xbc, 0x6f, 0x06, 0x83, 0xd1, 0xd1, 0xa6, 0xee, 0x0c, 0x2f, 0xf5, 0x9d, 0xbe, 0x33, 0xfe, 0xf4,
		0x49, 0x9e, 0xe8, 0x03, 0xfd, 0x8f, 0x7f, 0xfe, 0x2c, 0x86, 0xd2, 0xf5, 0xc4, 0x6f, 0xa5, 0xf2,
		0x3e, 0xac, 0x70, 0x65, 0x95, 0x7e, 0x7f, 0x61, 0xa7, 0x08, 0x74, 0xdf, 0x3b, 0xac, 0xea, 0xd7,
		0xdf, 0xa1, 0xe5, 0x5a, 0x59, 0xe6, 0x50, 0x32, 0xc6, 0x0e, 0x1a, 0xb2, 0x02, 0x0f, 0xc6, 0xf8,
		0xd8, 0xd6, 0xc4, 0x5e, 0x02, 0xe3, 0xb7, 0x38, 0xe3, 0x4a, 0x84, 0xb1, 0xcb, 0xa1, 0x72, 0x03,
		0x16, 0x4f, 0xc3, 0xf5, 0x77, 0x9c, 0xab, 0x8c, 0xa3, 0x24, 0x3b, 0xb0, 0x44, 0x49, 0xf4, 0x91,
		0x1f, 0x38, 0x43, 0x9a, 0xf7, 0xee, 0x4f, 0xf3, 0xf7, 0xef, 0xb0, 0xbd, 0x52, 0x21, 0xb0, 0x46,
		0x88, 0x92, 0x65, 0xa0, 0x9f, 0x9c, 0x0c, 0xac, 0x5b, 0x09, 0x0c, 0xff, 0xc0, 0x0d, 0x09, 0xf5,
		0xe5, 0xcf, 0xc2, 0x2a, 0xf9, 0x9f, 0xa6, 0xa5, 0xa8, 0x25, 0xc9, 0x17, 0x5e, 0xd5, 0x6f, 0xbf,
		0xce, 0xb6, 0xe3, 0x4a, 0x48, 0x10, 0xb1, 0x29, 0xb2, 0x8a, 0x7d, 0x1c, 0x04, 0xd8, 0xf3, 0x55,
		0xcd, 0x9a, 0x65, 0x5e, 0xe4, 0xc6, 0xa0, 0xfa, 0xe5, 0x77, 0xe3, 0xab, 0xb8, 0xc3, 0x90, 0x75,
		0xcb, 0x92, 0x0f, 0xe1, 0xcc, 0x8c, 0xa8, 0x98, 0x83, 0xf3, 0x0d, 0xce, 0xb9, 0x3a, 0x15, 0x19,
		0x84, 0xb6, 0x03, 0x42, 0x1e, 0xae, 0xe5, 0x1c, 0x9c, 0xbf, 0xc1, 0x39, 0x11, 0xc7, 0x8a, 0x25,
		0x25, 0x8c, 0x37, 0x61, 0xf9, 0x0e, 0xf6, 0x8e, 0x1c, 0x9f, 0xdf, 0xd2, 0xcc, 0x41, 0xf7, 0x26,
		0xa7, 0x5b, 0xe2, 0x40, 0x7a, 0x6d, 0x43, 0xb8, 0xae, 0x41, 0xa1, 0xa7, 0xe9, 0x78, 0x0e, 0x8a,
		0xaf, 0x70, 0x8a, 0x05, 0xa2, 0x4f, 0xa0, 0x75, 0x28, 0xf7, 0x1d, 0x5e, 0x99, 0x92, 0xe1, 0x6f,
		0x71, 0x78, 0x49, 0x60, 0x38, 0x85, 0xeb, 0xb8, 0x23, 0x8b, 0x94, 0xad, 0x64, 0x8a, 0xdf, 0x14,
		0x14, 0x02, 0xc3, 0x29, 0x4e, 0xe1, 0xd6, 0xaf, 0x0a, 0x0a, 0x3f, 0xe2, 0xcf, 0x17, 0xa0, 0xe4,
		0xd8, 0xd6, 0xb1, 0x63, 0xcf, 0x63, 0xc4, 0x6f, 0x71, 0x06, 0xe0, 0x10, 0x42, 0x70, 0x1d, 0x8a,
		0xf3, 0x2e, 0xc4, 0xef, 0xbc, 0x2b, 0xb6, 0x87, 0x58, 0x81, 0x1d, 0x58, 0x12, 0x09, 0xca, 0x74,
		0xec, 0x39, 0x28, 0x7e, 0x97, 0x53, 0x54, 0x22, 0x30, 0x3e, 0x8d, 0x00, 0xfb, 0x41, 0x1f, 0xcf,
		0x43, 0xf2, 0xb6, 0x98, 0x06, 0x87, 0x70, 0x57, 0x1e, 0x61, 0x5b, 0x1f, 0xcc, 0xc7, 0xf0, 0x7b,
		0xc2, 0x95, 0x02, 0x43, 0x28, 0x1a, 0xb0, 0x38, 0xd4, 0x3c, 0x7f, 0xa0, 0x59, 0x73, 0x2d, 0xc7,
		0xef, 0x73, 0x8e, 0x72, 0x08, 0xe2, 0x1e, 0x19, 0xd9, 0xa7, 0xa1, 0xf9, 0x9a, 0xf0, 0x48, 0x04,
		0xc6, 0xb7, 0x9e, 0x1f, 0xd0, 0x2b, 0xad, 0xd3, 0xb0, 0xfd, 0x81, 0xd8, 0x7a, 0x0c, 0xbb, 0x17,
		0x65, 0xbc, 0x0e, 0x45, 0xdf, 0xbc, 0x37, 0x17, 0xcd, 0x1f, 0x8a, 0x95, 0xa6, 0x00, 0x02, 0x7e,
		0x19, 0xce, 0xce, 0x2c, 0x13, 0x73, 0x90, 0xfd, 0x11, 0x27, 0x5b, 0x9b, 0x51, 0x2a, 0x78, 0x4a,
		0x38, 0x2d, 0xe5, 0x1f, 0x8b, 0x94, 0x80, 0x27, 0xb8, 0x3a, 0xe4, 0xac, 0xe0, 0x6b, 0xbd, 0xd3,
		0x79, 0xed, 0x4f, 0x84, 0xd7, 0x18, 0x36, 0xe6, 0xb5, 0x03, 0x58, 0xe3, 0x8c, 0xa7, 0x5b, 0xd7,
		0x3f, 0x15, 0x89, 0x95, 0xa1, 0x0f, 0xe3, 0xab, 0xfb, 0x39, 0x58, 0x0f, 0xdd, 0x29, 0x9a, 0x52,
		0x5f, 0x1d, 0x6a, 0xee, 0x1c, 0xcc, 0x5f, 0xe7, 0xcc, 0x22, 0xe3, 0x87, 0x5d, 0xad, 0xbf, 0xa7,
		0xb9, 0x84, 0xfc, 0x25, 0xa8, 0x0a, 0xf2, 0x91, 0xed, 0x61, 0xdd, 0xe9, 0xdb, 0xe6, 0x3d, 0x6c,
		0xcc, 0x41, 0xfd, 0x67, 0x13, 0x4b, 0x75, 0x18, 0x81, 0x13, 0xe6, 0x16, 0x48, 0x61, 0xaf, 0xa2,
		0x9a, 0x43, 0xd7, 0xf1, 0x82, 0x04, 0xc6, 0x6f, 0x88, 0x95, 0x0a, 0x71, 0x2d, 0x0a, 0x93, 0x9b,
		0x50, 0xa1, 0x8f, 0xf3, 0x86, 0xe4, 0x9f, 0x73, 0xa2, 0xc5, 0x31, 0x8a, 0x27, 0x0e, 0xdd, 0x19,
		0xba, 0x9a, 0x37, 0x4f, 0xfe, 0xfb, 0x0b, 0x91, 0x38, 0x38, 0x84, 0x27, 0x8e, 0xe0, 0xd8, 0xc5,
		0xa4, 0xda, 0xcf, 0xc1, 0xf0, 0x4d, 0x91, 0x38, 0x04, 0x86, 0x53, 0x88, 0x86, 0x61, 0x0e, 0x8a,
		0xbf, 0x14, 0x14, 0x02, 0x43, 0x28, 0x3e, 0x33, 0x2e, 0xb4, 0x1e, 0xee, 0x9b, 0x7e, 0xe0, 0xb1,
		0x56, 0xf8, 0xfe, 0x54, 0x7f, 0xf5, 0x6e, 0xbc, 0x09, 0x53, 0x22, 0x50, 0xf9, 0x26, 0x2c, 0x4d,
		0xb4, 0x18, 0x28, 0xe9, 0xf7, 0x2b, 0xd5, 0x9f, 0x7a, 0x9f, 0x27, 0xa3, 0x78, 0x87, 0x21, 0xef,
		0x92, 0x75, 0x8f, 0xf7, 0x01, 0xc9, 0x64, 0xaf, 0xbf, 0x1f, 0x2e, 0x7d, 0xac, 0x0d, 0x90, 0x6f,
		0xc0, 0x62, 0xac, 0x07, 0x48, 0xa6, 0xfa, 0x69, 0x4e, 0x55, 0x8e, 0xb6, 0x00, 0xf2, 0x15, 0xc8,
		0x92, 0x7a, 0x9e, 0x0c, 0xff, 0x19, 0x0e, 0xa7, 0xea, 0xf2, 0x27, 0xa1, 0x20, 0xea, 0x78, 0x32,
		0xf4, 0x67, 0x39, 0x34, 0x84, 0x10, 0xb8, 0xa8, 0xe1, 0xc9, 0xf0, 0x9f, 0x13, 0x70, 0x01, 0x21,
		0xf0, 0xf9, 0x5d, 0xf8, 0xb7, 0x3f, 0x9f, 0xe5, 0x79, 0x58, 0xf8, 0xee, 0x3a, 0x2c, 0xf0, 0xe2,
		0x9d, 0x8c, 0xfe, 0x02, 0x7f, 0xb9, 0x40, 0xc8, 0xcf, 0x41, 0x6e, 0x4e, 0x87, 0xff, 0x02, 0x87,
		0x32, 0x7d, 0xb9, 0x01, 0xa5, 0x48, 0xc1, 0x4e, 0x86, 0xff, 0x22, 0x87, 0x47, 0x51, 0xc4, 0x74,
		0x5e, 0xb0, 0x93, 0x09, 0x7e, 0x49, 0x98, 0xce, 0x11, 0xc4, 0x6d, 0xa2, 0x56, 0x27, 0xa3, 0x7f,
		0x59, 0x78, 0x5d, 0x40, 0xe4, 0x17, 0xa0, 0x18, 0xe6, 0xdf, 0x64, 0xfc, 0xaf, 0x70, 0xfc, 0x18,
		0x43, 0x3c, 0x10, 0xc9, 0xff, 0xc9, 0x14, 0xbf, 0x2a, 0x3c, 0x10, 0x41, 0x91, 0x6d, 0x34, 0x59,
		0xd3, 0x93, 0x99, 0x7e, 0x4d, 0x6c, 0xa3, 0x89, 0x92, 0x4e, 0x56, 0x93, 0xa6, 0xc1, 0x64, 0x8a,
		0x5f, 0x17, 0xab, 0x49, 0xf5, 0x89, 0x19, 0x93, 0x45, 0x32, 0x99, 0xe3, 0x8b, 0xc2, 0x8c, 0x89,
		0x1a, 0x29, 0x77, 0x00, 0x4d, 0x17, 0xc8, 0x64, 0xbe, 0x2f, 0x71, 0xbe, 0xe5, 0xa9, 0xfa, 0x28,
		0xbf, 0x08, 0x6b, 0xb3, 0x8b, 0x63, 0x32, 0xeb, 0x97, 0xdf, 0x9f, 0x38, 0xce, 0x44, 0x6b, 0xa3,
		0x7c, 0x30, 0xce, 0xb2, 0xd1, 0xc2, 0x98, 0x4c, 0xfb, 0xc6, 0xfb, 0xf1, 0x44, 0x1b, 0xad, 0x8b,
		0x72, 0x1d, 0x60, 0x5c, 0x93, 0x92, 0xb9, 0xde, 0xe4, 0x5c, 0x11, 0x10, 0xd9, 0x1a, 0xbc, 0x24,
		0x25, 0xe3, 0xbf, 0x22, 0xb6, 0x06, 0x47, 0x90, 0xad, 0x21, 0xaa, 0x51, 0x32, 0xfa, 0x2d, 0xb1,
		0x35, 0x04, 0x44, 0xbe, 0x0e, 0x05, 0x7b, 0x64, 0x59, 0x24, 0xb6, 0xd0, 0xfd, 0x7f, 0x92, 0x55,
		0xfd, 0xf7, 0x0f, 0x38, 0x58, 0x00, 0xe4, 0x2b, 0x90, 0xc3, 0xc3, 0x23, 0x6c, 0x24, 0x21, 0xff,
		0xe3, 0x03, 0x91, 0x4f, 0x88, 0xb6, 0xfc, 0x02, 0x00, 0x3b, 0x4c, 0xd3, 0x0f, 0x45, 0x09, 0xd8,
		0xff, 0xfc, 0x80, 0xff, 0x58, 0x62, 0x0c, 0x19, 0x13, 0xb0, 0x9f, 0x5e, 0xdc, 0x9f, 0xe0, 0xdd,
		0x38, 0x01, 0x3d, 0x80, 0x5f, 0x83, 0x85, 0x5b, 0xbe, 0x63, 0x07, 0x5a, 0x3f, 0x09, 0xfd, 0x5f,
		0x1c, 0x2d, 0xf4, 0x89, 0xc3, 0x86, 0x8e, 0x87, 0x03, 0xad, 0xef, 0x27, 0x61, 0xff, 0x9b, 0x63,
		0x43, 0x00, 0x01, 0xeb, 0x9a, 0x1f, 0xcc, 0x33, 0xef, 0xef, 0x0a, 0xb0, 0x00, 0x10, 0xa3, 0xc9,
		0xff, 0xb7, 0xf1, 0x71, 0x12, 0xf6, 0x3d, 0x61, 0x34, 0xd7, 0x97, 0x3f, 0x09, 0x45, 0xf2, 0x2f,
		0xfb, 0x05, 0x54, 0x02, 0xf8, 0x7f, 0x38, 0x78, 0x8c, 0x20, 0x6f, 0xf6, 0x03, 0x23, 0x30, 0x93,
		0x9d, 0xfd, 0xbf, 0x7c, 0xa5, 0x85, 0xbe, 0x5c, 0x87, 0x92, 0x1f, 0x18, 0xc6, 0x88, 0x77, 0x34,
		0x09, 0xf0, 0xff, 0xfb, 0x20, 0x3c, 0xe4, 0x86, 0x98, 0xad, 0xe6, 0xec, 0xfb, 0x3a, 0xd8, 0x71,
		0x76, 0x1c, 0x76, 0x53, 0xf7, 0x4a, 0x2d, 0xf9, 0xca, 0x0d, 0xfe, 0x3f, 0x05, 0x65, 0x17, 0x7b,
		0xbe, 0x63, 0xf3, 0x8b, 0xb7, 0xec, 0x50, 0x33, 0xed, 0xf5, 0xd3, 0xdd, 0xd6, 0xd5, 0x86, 0x90,
		0xef, 0x50, 0x12, 0x84, 0x20, 0xbb, 0x1f, 0xf9, 0x15, 0x10, 0xfd, 0x95, 0xe3, 0xd3, 0x50, 0xac,
		0x1b, 0x86, 0x87, 0x7d, 0x1f, 0xfb, 0xfc, 0x8a, 0x7f, 0x71, 0x93, 0xbc, 0x66, 0x93, 0x8b, 0x95,
		0xf1, 0x38, 0x7a, 0x18, 0x8a, 0x07, 0xd8, 0xc2, 0xee, 0xc0, 0xb1, 0xc5, 0x05, 0xfe, 0x58, 0x20,
		0x67, 0xdf, 0xfb, 0xea, 0xb9, 0x54, 0xed, 0x1a, 0x2c, 0x70, 0x00, 0x5a, 0x83, 0xfc, 0x3e, 0xfb,
		0x39, 0x56, 0x8a, 0xde, 0xc8, 0xf3, 0x27, 0x22, 0xef, 0x06, 0x1e, 0xc6, 0x01, 0xbf, 0x12, 0xe5,
		0x4f, 0x5b, 0x85, 0xf7, 0xbe, 0xf3, 0x68, 0xea, 0x7b, 0xdf, 0x79, 0x34, 0xf5, 0xfd, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x45, 0x52, 0xe9, 0xa2, 0x86, 0x31, 0x00, 0x00,
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

func init() { proto.RegisterFile("person.proto", fileDescriptorPerson) }

var fileDescriptorPerson = []byte{
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
