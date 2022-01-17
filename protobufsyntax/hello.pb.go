// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: hello.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Option int32

const (
	Option_option_a Option = 0
	Option_option_b Option = 1
	Option_option_c Option = 2
)

// Enum value maps for Option.
var (
	Option_name = map[int32]string{
		0: "option_a",
		1: "option_b",
		2: "option_c",
	}
	Option_value = map[string]int32{
		"option_a": 0,
		"option_b": 1,
		"option_c": 2,
	}
)

func (x Option) Enum() *Option {
	p := new(Option)
	*p = x
	return p
}

func (x Option) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Option) Descriptor() protoreflect.EnumDescriptor {
	return file_hello_proto_enumTypes[0].Descriptor()
}

func (Option) Type() protoreflect.EnumType {
	return &file_hello_proto_enumTypes[0]
}

func (x Option) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Option.Descriptor instead.
func (Option) EnumDescriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     string             `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key       *string            `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Code      int32              `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Response  []*Result          `protobuf:"bytes,4,rep,name=response,proto3" json:"response,omitempty"`
	ResultMap map[string]*Result `protobuf:"bytes,5,rep,name=result_map,json=resultMap,proto3" json:"result_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Result    *Result            `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
	// Types that are assignable to OneofSample:
	//	*Response_Stub1
	//	*Response_Stub2
	OneofSample isResponse_OneofSample `protobuf_oneof:"oneof_sample"`
	Details     *anypb.Any             `protobuf:"bytes,9,opt,name=details,proto3" json:"details,omitempty"` // 需要先import "google/protobuf/any.proto"; 然后再生成代码时通过-I 导入
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Response) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetResponse() []*Result {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Response) GetResultMap() map[string]*Result {
	if x != nil {
		return x.ResultMap
	}
	return nil
}

func (x *Response) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (m *Response) GetOneofSample() isResponse_OneofSample {
	if m != nil {
		return m.OneofSample
	}
	return nil
}

func (x *Response) GetStub1() *Stub1 {
	if x, ok := x.GetOneofSample().(*Response_Stub1); ok {
		return x.Stub1
	}
	return nil
}

func (x *Response) GetStub2() *Stub2 {
	if x, ok := x.GetOneofSample().(*Response_Stub2); ok {
		return x.Stub2
	}
	return nil
}

func (x *Response) GetDetails() *anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type isResponse_OneofSample interface {
	isResponse_OneofSample()
}

type Response_Stub1 struct {
	Stub1 *Stub1 `protobuf:"bytes,7,opt,name=stub1,proto3,oneof"`
}

type Response_Stub2 struct {
	Stub2 *Stub2 `protobuf:"bytes,8,opt,name=stub2,proto3,oneof"`
}

func (*Response_Stub1) isResponse_OneofSample() {}

func (*Response_Stub2) isResponse_OneofSample() {}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Stub1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Stub1) Reset() {
	*x = Stub1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stub1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stub1) ProtoMessage() {}

func (x *Stub1) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stub1.ProtoReflect.Descriptor instead.
func (*Stub1) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{2}
}

func (x *Stub1) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Stub2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Stub2) Reset() {
	*x = Stub2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stub2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stub2) ProtoMessage() {}

func (x *Stub2) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stub2.ProtoReflect.Descriptor instead.
func (*Stub2) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{3}
}

func (x *Stub2) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x04, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x15, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x62, 0x31, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x62, 0x31, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x62, 0x31, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x62, 0x32, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x62, 0x32, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x62, 0x32, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x55, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6b, 0x65, 0x79, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0f, 0x10,
	0x15, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x32, 0x22, 0x24, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a,
	0x05, 0x53, 0x74, 0x75, 0x62, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x05,
	0x53, 0x74, 0x75, 0x62, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x32, 0x0a, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x10, 0x02, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x41,
	0x4f, 0x6c, 0x6f, 0x77, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hello_proto_goTypes = []interface{}{
	(Option)(0),       // 0: grpc.demo.hello.option
	(*Response)(nil),  // 1: grpc.demo.hello.Response
	(*Result)(nil),    // 2: grpc.demo.hello.Result
	(*Stub1)(nil),     // 3: grpc.demo.hello.Stub1
	(*Stub2)(nil),     // 4: grpc.demo.hello.Stub2
	nil,               // 5: grpc.demo.hello.Response.ResultMapEntry
	(*anypb.Any)(nil), // 6: google.protobuf.Any
}
var file_hello_proto_depIdxs = []int32{
	2, // 0: grpc.demo.hello.Response.response:type_name -> grpc.demo.hello.Result
	5, // 1: grpc.demo.hello.Response.result_map:type_name -> grpc.demo.hello.Response.ResultMapEntry
	2, // 2: grpc.demo.hello.Response.result:type_name -> grpc.demo.hello.Result
	3, // 3: grpc.demo.hello.Response.stub1:type_name -> grpc.demo.hello.Stub1
	4, // 4: grpc.demo.hello.Response.stub2:type_name -> grpc.demo.hello.Stub2
	6, // 5: grpc.demo.hello.Response.details:type_name -> google.protobuf.Any
	2, // 6: grpc.demo.hello.Response.ResultMapEntry.value:type_name -> grpc.demo.hello.Result
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stub1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stub2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_hello_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Response_Stub1)(nil),
		(*Response_Stub2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		EnumInfos:         file_hello_proto_enumTypes,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}
