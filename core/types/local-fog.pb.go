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

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AddrV4 uint32 `protobuf:"fixed32,2,opt,name=addr_v4,json=addrV4,proto3" json:"addr_v4,omitempty"`
	AddrV6 []byte `protobuf:"bytes,3,opt,name=addr_v6,json=addrV6,proto3" json:"addr_v6,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{0}
}

func (x *NodeInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeInfo) GetAddrV4() uint32 {
	if x != nil {
		return x.AddrV4
	}
	return 0
}

func (x *NodeInfo) GetAddrV6() []byte {
	if x != nil {
		return x.AddrV6
	}
	return nil
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{1}
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{2}
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{3}
}

type SyncReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncReply) Reset() {
	*x = SyncReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncReply) ProtoMessage() {}

func (x *SyncReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SyncReply.ProtoReflect.Descriptor instead.
func (*SyncReply) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{4}
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint64 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Body  []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{5}
}

func (x *CallRequest) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *CallRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type CallReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *CallReply) Reset() {
	*x = CallReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallReply) ProtoMessage() {}

func (x *CallReply) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallReply.ProtoReflect.Descriptor instead.
func (*CallReply) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{6}
}

func (x *CallReply) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type GetProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint64 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetProgramRequest) Reset() {
	*x = GetProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgramRequest) ProtoMessage() {}

func (x *GetProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgramRequest.ProtoReflect.Descriptor instead.
func (*GetProgramRequest) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{7}
}

func (x *GetProgramRequest) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type GetProgramReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProgramReply) Reset() {
	*x = GetProgramReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_fog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProgramReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgramReply) ProtoMessage() {}

func (x *GetProgramReply) ProtoReflect() protoreflect.Message {
	mi := &file_local_fog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgramReply.ProtoReflect.Descriptor instead.
func (*GetProgramReply) Descriptor() ([]byte, []int) {
	return file_local_fog_proto_rawDescGZIP(), []int{8}
}

var File_local_fog_proto protoreflect.FileDescriptor

var file_local_fog_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2d, 0x66, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x5f,
	0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x06, 0x61, 0x64, 0x64, 0x72, 0x56, 0x34,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x61, 0x64, 0x64, 0x72, 0x56, 0x36, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x38, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x23, 0x0a, 0x09, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x2a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xf4, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x6f, 0x67, 0x12, 0x34, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x15, 0x2e, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69,
	0x61, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_local_fog_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_local_fog_proto_goTypes = []interface{}{
	(*NodeInfo)(nil),          // 0: tutorial.NodeInfo
	(*PingRequest)(nil),       // 1: tutorial.PingRequest
	(*PingReply)(nil),         // 2: tutorial.PingReply
	(*SyncRequest)(nil),       // 3: tutorial.SyncRequest
	(*SyncReply)(nil),         // 4: tutorial.SyncReply
	(*CallRequest)(nil),       // 5: tutorial.CallRequest
	(*CallReply)(nil),         // 6: tutorial.CallReply
	(*GetProgramRequest)(nil), // 7: tutorial.GetProgramRequest
	(*GetProgramReply)(nil),   // 8: tutorial.GetProgramReply
}
var file_local_fog_proto_depIdxs = []int32{
	1, // 0: tutorial.LocalFog.Ping:input_type -> tutorial.PingRequest
	3, // 1: tutorial.LocalFog.Sync:input_type -> tutorial.SyncRequest
	5, // 2: tutorial.LocalFog.Call:input_type -> tutorial.CallRequest
	7, // 3: tutorial.LocalFog.GetProgram:input_type -> tutorial.GetProgramRequest
	2, // 4: tutorial.LocalFog.Ping:output_type -> tutorial.PingReply
	4, // 5: tutorial.LocalFog.Sync:output_type -> tutorial.SyncReply
	6, // 6: tutorial.LocalFog.Call:output_type -> tutorial.CallReply
	8, // 7: tutorial.LocalFog.GetProgram:output_type -> tutorial.GetProgramReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_local_fog_proto_init() }
func file_local_fog_proto_init() {
	if File_local_fog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_local_fog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
			switch v := v.(*PingRequest); i {
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
			switch v := v.(*PingReply); i {
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
			switch v := v.(*SyncRequest); i {
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
			switch v := v.(*SyncReply); i {
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
		file_local_fog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_local_fog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallReply); i {
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
		file_local_fog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProgramRequest); i {
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
		file_local_fog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProgramReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_local_fog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_local_fog_proto_goTypes,
		DependencyIndexes: file_local_fog_proto_depIdxs,
		MessageInfos:      file_local_fog_proto_msgTypes,
	}.Build()
	File_local_fog_proto = out.File
	file_local_fog_proto_rawDesc = nil
	file_local_fog_proto_goTypes = nil
	file_local_fog_proto_depIdxs = nil
}
