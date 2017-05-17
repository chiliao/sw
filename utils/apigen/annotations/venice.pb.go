// Code generated by protoc-gen-gogo.
// source: venice.proto
// DO NOT EDIT!

/*
	Package venice is a generated protocol buffer package.

	It is generated from these files:
		venice.proto

	It has these top-level messages:
*/
package venice

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_FileGrpcDest = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51235,
	Name:          "venice.fileGrpcDest",
	Tag:           "bytes,51235,opt,name=fileGrpcDest",
	Filename:      "venice.proto",
}

var E_ApiVersion = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51250,
	Name:          "venice.apiVersion",
	Tag:           "bytes,51250,opt,name=apiVersion",
	Filename:      "venice.proto",
}

var E_ApiPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51251,
	Name:          "venice.apiPrefix",
	Tag:           "bytes,51251,opt,name=apiPrefix",
	Filename:      "venice.proto",
}

var E_MethodOper = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51270,
	Name:          "venice.methodOper",
	Tag:           "bytes,51270,opt,name=methodOper",
	Filename:      "venice.proto",
}

var E_ObjectPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51290,
	Name:          "venice.objectPrefix",
	Tag:           "bytes,51290,opt,name=objectPrefix",
	Filename:      "venice.proto",
}

func init() {
	proto.RegisterExtension(E_FileGrpcDest)
	proto.RegisterExtension(E_ApiVersion)
	proto.RegisterExtension(E_ApiPrefix)
	proto.RegisterExtension(E_MethodOper)
	proto.RegisterExtension(E_ObjectPrefix)
}

func init() { proto.RegisterFile("venice.proto", fileDescriptorVenice) }

var fileDescriptorVenice = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4b, 0xcd, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x14, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0x95, 0x56, 0x4e, 0x5c, 0x3c, 0x69, 0x99, 0x39, 0xa9, 0xee,
	0x45, 0x05, 0xc9, 0x2e, 0xa9, 0xc5, 0x25, 0x42, 0x32, 0x7a, 0x10, 0x2d, 0x7a, 0x30, 0x2d, 0x7a,
	0x6e, 0x99, 0x39, 0xa9, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12, 0x8b, 0x27, 0x30, 0x2b,
	0x30, 0x6a, 0x70, 0x06, 0xa1, 0xe8, 0xb1, 0x72, 0xe4, 0xe2, 0x4a, 0x2c, 0xc8, 0x0c, 0x4b, 0x2d,
	0x2a, 0xce, 0xcc, 0xcf, 0x13, 0x92, 0xc7, 0x30, 0x21, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x19, 0x6e,
	0xc8, 0x26, 0xa8, 0x21, 0x48, 0x9a, 0xac, 0xec, 0xb9, 0x38, 0x13, 0x0b, 0x32, 0x03, 0x8a, 0x52,
	0xd3, 0x32, 0x2b, 0x08, 0x9b, 0xb0, 0x19, 0x6a, 0x02, 0x42, 0x8f, 0x95, 0x03, 0x17, 0x57, 0x6e,
	0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x7f, 0x41, 0x6a, 0x91, 0x90, 0x1c, 0x86, 0x09, 0xbe, 0x50, 0x49,
	0x88, 0x01, 0xc7, 0x60, 0x4e, 0x40, 0xe8, 0xb1, 0x72, 0xe5, 0xe2, 0xc9, 0x4f, 0xca, 0x4a, 0x4d,
	0x2e, 0xc1, 0xe9, 0x0a, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0xb8, 0x2b, 0x6e, 0xc1, 0x02, 0x03,
	0x59, 0x9b, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x31, 0xfb, 0x7a, 0xd9, 0xa3, 0x01, 0x00, 0x00,
}
