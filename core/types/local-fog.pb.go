// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: local-fog.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of the requset.
type Request_Type int32

const (
	Request_RESERVED    Request_Type = 0
	Request_PING        Request_Type = 1
	Request_SYNC        Request_Type = 2
	Request_CALL        Request_Type = 3
	Request_GET_REQUEST Request_Type = 4
)

// Enum value maps for Request_Type.
var (
	Request_Type_name = map[int32]string{
		0: "RESERVED",
		1: "PING",
		2: "SYNC",
		3: "CALL",
		4: "GET_REQUEST",
	}
	Request_Type_value = map[string]int32{
		"RESERVED":    0,
		"PING":        1,
		"SYNC":        2,
		"CALL":        3,
		"GET_REQUEST": 4,
	}
)

func (x Request_Type) Enum() *Request_Type {
	p := new(Request_Type)
	*p = x
	return p
}

func (x Request_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_local_fog_proto_enumTypes[0].Descriptor()
}

func (Request_Type) Type() protoreflect.EnumType {
	return &file_local_fog_proto_enumTypes[0]
}

func (x Request_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Type.Descriptor instead.
func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Conetnt:
	//
	//	*Request_Ping
	//	*Request_Sync
	//	*Request_Call
	//	*Request_GetProgram
	Conetnt isRequest_Conetnt `protobuf_oneof:"conetnt"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{0}
}

func (m *Request) GetConetnt() isRequest_Conetnt {
	if m != nil {
		return m.Conetnt
	}
	return nil
}

func (x *Request) GetPing() *Ping {
	if x, ok := x.GetConetnt().(*Request_Ping); ok {
		return x.Ping
	}
	return nil
}

func (x *Request) GetSync() *Sync {
	if x, ok := x.GetConetnt().(*Request_Sync); ok {
		return x.Sync
	}
	return nil
}

func (x *Request) GetCall() *Call {
	if x, ok := x.GetConetnt().(*Request_Call); ok {
		return x.Call
	}
	return nil
}

func (x *Request) GetGetProgram() *GetProgram {
	if x, ok := x.GetConetnt().(*Request_GetProgram); ok {
		return x.GetProgram
	}
	return nil
}

type isRequest_Conetnt interface {
	isRequest_Conetnt()
}

type Request_Ping struct {
	Ping *Ping `protobuf:"bytes,1,opt,name=ping,proto3,oneof"`
}

type Request_Sync struct {
	Sync *Sync `protobuf:"bytes,2,opt,name=sync,proto3,oneof"`
}

type Request_Call struct {
	Call *Call `protobuf:"bytes,3,opt,name=call,proto3,oneof"`
}

type Request_GetProgram struct {
	GetProgram *GetProgram `protobuf:"bytes,4,opt,name=get_program,json=getProgram,proto3,oneof"`
}

func (*Request_Ping) isRequest_Conetnt() {}

func (*Request_Sync) isRequest_Conetnt() {}

func (*Request_Call) isRequest_Conetnt() {}

func (*Request_GetProgram) isRequest_Conetnt() {}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{1}
}

type Sync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sync) Reset() {
	*x = Sync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync) ProtoMessage() {}

func (x *Sync) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync.ProtoReflect.Descriptor instead.
func (*Sync) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{2}
}

type Call struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Body  []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Call) Reset() {
	*x = Call{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call) ProtoMessage() {}

func (x *Call) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call.ProtoReflect.Descriptor instead.
func (*Call) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{3}
}

func (x *Call) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *Call) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetProgram) Reset() {
	*x = GetProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgram) ProtoMessage() {}

func (x *GetProgram) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgram.ProtoReflect.Descriptor instead.
func (*GetProgram) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{4}
}

func (x *GetProgram) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

var File_local_fog_proto protoreflect.FileDescriptor

var file_local_fog_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2d, 0x66, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x84, 0x02, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a,
	0x04, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x00, 0x52, 0x04, 0x73,
	0x79, 0x6e, 0x63, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x48, 0x00, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x37, 0x0a, 0x0b, 0x67, 0x65, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x22, 0x43, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x43, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x65, 0x74,
	0x6e, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x06, 0x0a, 0x04, 0x53, 0x79,
	0x6e, 0x63, 0x22, 0x31, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x23, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_local_fog_proto_rawDescOnce sync.Once
	file_local_fog_proto_rawDescData = file_local_fog_proto_rawDesc
)

func file_local_fog_proto_rawDescGZIP() []byte {
	file_local_fog_proto_rawDescOnce.Do(func() {
		file_local_fog_proto_rawDescData = protoimpl.X.CompressGZIP(file_local_fog_proto_rawDescData)
	})
	return file_local_fog_proto_rawDescData
}

var file_local_fog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_local_fog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_local_fog_proto_goTypes = []interface{}{
	(Request_Type)(0),  // 0: tutorial.Request.Type
	(*Request)(nil),    // 1: tutorial.Request
	(*Ping)(nil),       // 2: tutorial.Ping
	(*Sync)(nil),       // 3: tutorial.Sync
	(*Call)(nil),       // 4: tutorial.Call
	(*GetProgram)(nil), // 5: tutorial.GetProgram
}
var file_local_fog_proto_depIdxs = []int32{
	2, // 0: tutorial.Request.ping:type_name -> tutorial.Ping
	3, // 1: tutorial.Request.sync:type_name -> tutorial.Sync
	4, // 2: tutorial.Request.call:type_name -> tutorial.Call
	5, // 3: tutorial.Request.get_program:type_name -> tutorial.GetProgram
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_local_fog_proto_init() }
func file_local_fog_proto_init() {
	if File_local_fog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_local_fog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_local_fog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_local_fog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync); i {
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
		file_local_fog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call); i {
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
		file_local_fog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProgram); i {
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
	file_local_fog_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_Ping)(nil),
		(*Request_Sync)(nil),
		(*Request_Call)(nil),
		(*Request_GetProgram)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_local_fog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_local_fog_proto_goTypes,
		DependencyIndexes: file_local_fog_proto_depIdxs,
		EnumInfos:         file_local_fog_proto_enumTypes,
		MessageInfos:      file_local_fog_proto_msgTypes,
	}.Build()
	File_local_fog_proto = out.File
	file_local_fog_proto_rawDesc = nil
	file_local_fog_proto_goTypes = nil
	file_local_fog_proto_depIdxs = nil
}
