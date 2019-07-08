// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/database_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible database errors.
type DatabaseErrorEnum_DatabaseError int32

const (
	// Enum unspecified.
	DatabaseErrorEnum_UNSPECIFIED DatabaseErrorEnum_DatabaseError = 0
	// The received error code is not known in this version.
	DatabaseErrorEnum_UNKNOWN DatabaseErrorEnum_DatabaseError = 1
	// Multiple requests were attempting to modify the same resource at once.
	// Please retry the request.
	DatabaseErrorEnum_CONCURRENT_MODIFICATION DatabaseErrorEnum_DatabaseError = 2
	// The request conflicted with existing data. This error will usually be
	// replaced with a more specific error if the request is retried.
	DatabaseErrorEnum_DATA_CONSTRAINT_VIOLATION DatabaseErrorEnum_DatabaseError = 3
)

var DatabaseErrorEnum_DatabaseError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CONCURRENT_MODIFICATION",
	3: "DATA_CONSTRAINT_VIOLATION",
}

var DatabaseErrorEnum_DatabaseError_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"CONCURRENT_MODIFICATION":   2,
	"DATA_CONSTRAINT_VIOLATION": 3,
}

func (x DatabaseErrorEnum_DatabaseError) String() string {
	return proto.EnumName(DatabaseErrorEnum_DatabaseError_name, int32(x))
}

func (DatabaseErrorEnum_DatabaseError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_114da3527f3b2c1a, []int{0, 0}
}

// Container for enum describing possible database errors.
type DatabaseErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseErrorEnum) Reset()         { *m = DatabaseErrorEnum{} }
func (m *DatabaseErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DatabaseErrorEnum) ProtoMessage()    {}
func (*DatabaseErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_114da3527f3b2c1a, []int{0}
}

func (m *DatabaseErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseErrorEnum.Unmarshal(m, b)
}
func (m *DatabaseErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseErrorEnum.Marshal(b, m, deterministic)
}
func (m *DatabaseErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseErrorEnum.Merge(m, src)
}
func (m *DatabaseErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DatabaseErrorEnum.Size(m)
}
func (m *DatabaseErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.DatabaseErrorEnum_DatabaseError", DatabaseErrorEnum_DatabaseError_name, DatabaseErrorEnum_DatabaseError_value)
	proto.RegisterType((*DatabaseErrorEnum)(nil), "google.ads.googleads.v2.errors.DatabaseErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/database_error.proto", fileDescriptor_114da3527f3b2c1a)
}

var fileDescriptor_114da3527f3b2c1a = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x6a, 0xf3, 0x30,
	0x10, 0xc5, 0xbf, 0x38, 0xf0, 0x15, 0x14, 0x4a, 0x53, 0x6f, 0x4a, 0xff, 0x65, 0xe1, 0x03, 0xc8,
	0xe0, 0xec, 0xd4, 0x95, 0x62, 0x3b, 0x41, 0xb4, 0x95, 0x42, 0xe2, 0xb8, 0x50, 0x0c, 0x46, 0xa9,
	0x8c, 0x30, 0x24, 0x52, 0xb0, 0xdc, 0x2c, 0x7b, 0x98, 0x2e, 0x7b, 0x94, 0x1e, 0xa5, 0xd0, 0x3b,
	0x14, 0x5b, 0x49, 0x20, 0x8b, 0x76, 0xa5, 0xa7, 0xd1, 0xef, 0x8d, 0xde, 0x0c, 0x18, 0x4a, 0xad,
	0xe5, 0xaa, 0xf0, 0xb9, 0x30, 0xbe, 0x95, 0x8d, 0xda, 0x06, 0x7e, 0x51, 0x55, 0xba, 0x32, 0xbe,
	0xe0, 0x35, 0x5f, 0x72, 0x53, 0xe4, 0xed, 0x1d, 0x6e, 0x2a, 0x5d, 0x6b, 0x77, 0x60, 0x49, 0xc8,
	0x85, 0x81, 0x07, 0x13, 0xdc, 0x06, 0xd0, 0x9a, 0xae, 0x6e, 0xf6, 0x4d, 0x37, 0xa5, 0xcf, 0x95,
	0xd2, 0x35, 0xaf, 0x4b, 0xad, 0x8c, 0x75, 0x7b, 0x6f, 0xe0, 0x3c, 0xda, 0x75, 0x8d, 0x1b, 0x3e,
	0x56, 0xaf, 0x6b, 0xaf, 0x04, 0xa7, 0x47, 0x45, 0xf7, 0x0c, 0xf4, 0x16, 0x74, 0x3e, 0x8d, 0x43,
	0x32, 0x26, 0x71, 0xd4, 0xff, 0xe7, 0xf6, 0xc0, 0xc9, 0x82, 0xde, 0x53, 0xf6, 0x44, 0xfb, 0x1d,
	0xf7, 0x1a, 0x5c, 0x84, 0x8c, 0x86, 0x8b, 0xd9, 0x2c, 0xa6, 0x49, 0xfe, 0xc8, 0x22, 0x32, 0x26,
	0x21, 0x4e, 0x08, 0xa3, 0x7d, 0xc7, 0xbd, 0x05, 0x97, 0x11, 0x4e, 0x70, 0x1e, 0x32, 0x3a, 0x4f,
	0x66, 0x98, 0xd0, 0x24, 0x4f, 0x09, 0x7b, 0xb0, 0xcf, 0xdd, 0xd1, 0x77, 0x07, 0x78, 0x2f, 0x7a,
	0x0d, 0xff, 0x1e, 0x62, 0xe4, 0x1e, 0xe5, 0x99, 0x36, 0xd1, 0xa7, 0x9d, 0xe7, 0x68, 0xe7, 0x92,
	0x7a, 0xc5, 0x95, 0x84, 0xba, 0x92, 0xbe, 0x2c, 0x54, 0x3b, 0xd8, 0x7e, 0x7f, 0x9b, 0xd2, 0xfc,
	0xb6, 0xce, 0x3b, 0x7b, 0xbc, 0x3b, 0xdd, 0x09, 0xc6, 0x1f, 0xce, 0x60, 0x62, 0x9b, 0x61, 0x61,
	0xa0, 0x95, 0x8d, 0x4a, 0x03, 0xd8, 0x7e, 0x69, 0x3e, 0xf7, 0x40, 0x86, 0x85, 0xc9, 0x0e, 0x40,
	0x96, 0x06, 0x99, 0x05, 0xbe, 0x1c, 0xcf, 0x56, 0x11, 0xc2, 0xc2, 0x20, 0x74, 0x40, 0x10, 0x4a,
	0x03, 0x84, 0x2c, 0xb4, 0xfc, 0xdf, 0xa6, 0x1b, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x98, 0x2e,
	0x2a, 0x44, 0xeb, 0x01, 0x00, 0x00,
}