// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

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

type E int32

const (
	E_ZERO E = 0
)

var E_name = map[int32]string{
	0: "ZERO",
}

var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}

func (x *E) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_value, data, "E")
	if err != nil {
		return err
	}
	*x = E(value)
	return nil
}

func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	p := new(M_Subenum)
	*p = x
	return p
}

func (x M_Subenum) String() string {
	return proto.EnumName(M_Subenum_name, int32(x))
}

func (x *M_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Subenum_value, data, "M_Subenum")
	if err != nil {
		return err
	}
	*x = M_Subenum(value)
	return nil
}

func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	p := new(M_Submessage_Submessage_Subenum)
	*p = x
	return p
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return proto.EnumName(M_Submessage_Submessage_Subenum_name, int32(x))
}

func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Submessage_Submessage_Subenum_value, data, "M_Submessage_Submessage_Subenum")
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(value)
	return nil
}

func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2     `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	S  *string `protobuf:"bytes,4,opt,name=s,def=default" json:"s,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField                   isM_OneofField `protobuf_oneof:"oneof_field"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

var extRange_M = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*M) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_M
}

func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

const Default_M_S string = "default"

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M) GetS() string {
	if m != nil && m.S != nil {
		return *m.S
	}
	return Default_M_S
}

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof"`
}

type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof"`
}

func (*M_OneofInt32) isM_OneofField() {}

func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof"`
}

type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}

func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

var E_ExtensionField = &proto.ExtensionDesc{
	ExtendedType:  (*M)(nil),
	ExtensionType: (*string)(nil),
	Field:         100,
	Name:          "goproto.protoc.import_public.sub.extension_field",
	Tag:           "bytes,100,opt,name=extension_field",
	Filename:      "import_public/sub/a.proto",
}

func init() {
	proto.RegisterEnum("goproto.protoc.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.protoc.import_public.sub.M")
	proto.RegisterType((*M_Submessage)(nil), "goproto.protoc.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
}

func init() { proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_382f7805394b5c4e) }

var fileDescriptor_382f7805394b5c4e = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8f, 0x93, 0x40,
	0x14, 0xc7, 0x77, 0x60, 0x75, 0xeb, 0xdb, 0xac, 0xae, 0x13, 0x6b, 0x70, 0x4f, 0x58, 0x3d, 0x90,
	0x35, 0x65, 0x12, 0x24, 0x1c, 0x7a, 0x93, 0x04, 0x7f, 0xa5, 0xa4, 0x09, 0xc4, 0x4b, 0x2f, 0x84,
	0x81, 0x01, 0x49, 0x80, 0x69, 0x3a, 0x33, 0xc6, 0x23, 0x7f, 0x80, 0xff, 0x95, 0xff, 0x98, 0x29,
	0xf4, 0x67, 0xac, 0x71, 0x6f, 0xbc, 0xc7, 0xe7, 0xf3, 0x7d, 0x93, 0x79, 0x03, 0xaf, 0xaa, 0x66,
	0xc5, 0xd7, 0x32, 0x59, 0x29, 0x5a, 0x57, 0x19, 0x11, 0x8a, 0x92, 0xd4, 0x5e, 0xad, 0xb9, 0xe4,
	0xd8, 0x2c, 0x79, 0xff, 0x31, 0x94, 0x99, 0x7d, 0x42, 0xda, 0x42, 0xd1, 0xbb, 0x33, 0x32, 0x1d,
	0xe8, 0xc9, 0x2f, 0x1d, 0x50, 0x88, 0x5d, 0xd0, 0x1a, 0xc7, 0x40, 0x26, 0xb2, 0xae, 0x9d, 0xb7,
	0xf6, 0xff, 0xf2, 0xec, 0xd0, 0x89, 0xb4, 0xc6, 0xc1, 0x63, 0x40, 0xc2, 0xb8, 0x34, 0x91, 0xf5,
	0x64, 0x76, 0x95, 0xb3, 0x22, 0x55, 0xb5, 0x8c, 0x90, 0xc0, 0xaf, 0xe1, 0x9a, 0xb7, 0x8c, 0x17,
	0x49, 0xd5, 0xca, 0xf7, 0x8e, 0xa1, 0x99, 0xc8, 0x7a, 0xf4, 0xf9, 0x22, 0x82, 0xbe, 0xf9, 0x65,
	0xd3, 0x3b, 0x41, 0x3c, 0xd7, 0xd0, 0x4d, 0x64, 0xe9, 0xc7, 0x88, 0xe7, 0xde, 0xfd, 0x46, 0x00,
	0xb1, 0xa2, 0x0d, 0x13, 0x22, 0x2d, 0x19, 0xf6, 0xe0, 0xa5, 0xd8, 0x57, 0xc9, 0x71, 0x3e, 0xda,
	0xe6, 0xbf, 0x38, 0xfc, 0x5f, 0x1c, 0x26, 0xfd, 0xc3, 0xf3, 0xdc, 0xfe, 0x5c, 0xfa, 0x79, 0xcf,
	0x73, 0x27, 0xef, 0x00, 0x1f, 0xa6, 0x27, 0xb1, 0xa2, 0xac, 0x55, 0x0d, 0x1e, 0xc3, 0xf3, 0x30,
	0x89, 0xbf, 0xf9, 0x61, 0x10, 0xc7, 0x1f, 0x3e, 0x05, 0xc9, 0x32, 0x88, 0x16, 0xb7, 0x17, 0xbe,
	0x71, 0x66, 0x48, 0x51, 0xb1, 0x3a, 0x9f, 0x8c, 0xe1, 0x6a, 0xe7, 0x02, 0x3c, 0x0e, 0xb7, 0xc2,
	0xfd, 0x68, 0x94, 0xdf, 0x76, 0x5d, 0xd7, 0x69, 0xfe, 0xcd, 0xee, 0x26, 0x7a, 0xfe, 0xfe, 0x06,
	0x50, 0x80, 0x47, 0x70, 0x39, 0x70, 0xb3, 0x39, 0x3c, 0x63, 0x3f, 0x25, 0x6b, 0x45, 0xc5, 0xdb,
	0x81, 0xc0, 0x6f, 0x1e, 0xb0, 0x1e, 0x23, 0xdf, 0x2c, 0x25, 0x7a, 0xba, 0x77, 0x3f, 0x6e, 0x54,
	0x7f, 0xbe, 0xfc, 0x5a, 0x56, 0xf2, 0xbb, 0xa2, 0x76, 0xc6, 0x1b, 0x52, 0xf2, 0x3a, 0x6d, 0x4b,
	0xd2, 0xa7, 0x50, 0x55, 0x90, 0x1f, 0x0e, 0xc9, 0x9a, 0x7c, 0xa8, 0xb3, 0x69, 0xc9, 0xda, 0x69,
	0xc9, 0x89, 0x64, 0x42, 0xe6, 0xa9, 0x4c, 0xc9, 0x5f, 0x2f, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0x86, 0xfe, 0xf7, 0x92, 0x02, 0x00, 0x00,
}