// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: net.proto

package netgrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*TaskMessage_Ping
	Payload       isTaskMessage_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskMessage) Reset() {
	*x = TaskMessage{}
	mi := &file_net_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMessage) ProtoMessage() {}

func (x *TaskMessage) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskMessage.ProtoReflect.Descriptor instead.
func (*TaskMessage) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

func (x *TaskMessage) GetPayload() isTaskMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TaskMessage) GetPing() *Ping {
	if x != nil {
		if x, ok := x.Payload.(*TaskMessage_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

type isTaskMessage_Payload interface {
	isTaskMessage_Payload()
}

type TaskMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,1,opt,name=ping,proto3,oneof"`
}

func (*TaskMessage_Ping) isTaskMessage_Payload() {}

type Ping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int32                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nonce         string                 `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ping) Reset() {
	*x = Ping{}
	mi := &file_net_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[1]
	if x != nil {
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
	return file_net_proto_rawDescGZIP(), []int{1}
}

func (x *Ping) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Ping) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 任务完成通知或其他信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_net_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[2]
	if x != nil {
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
	return file_net_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nonce         string                 `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Salt          string                 `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	Num           int32                  `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Noncs         string                 `protobuf:"bytes,4,opt,name=noncs,proto3" json:"noncs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pong) Reset() {
	*x = Pong{}
	mi := &file_net_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{3}
}

func (x *Pong) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Pong) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *Pong) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Pong) GetNoncs() string {
	if x != nil {
		return x.Noncs
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_net_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_net_proto protoreflect.FileDescriptor

const file_net_proto_rawDesc = "" +
	"\n" +
	"\tnet.proto\x12\anetsend\"=\n" +
	"\vTaskMessage\x12#\n" +
	"\x04ping\x18\x01 \x01(\v2\r.netsend.PingH\x00R\x04pingB\t\n" +
	"\apayload\".\n" +
	"\x04Ping\x12\x10\n" +
	"\x03uid\x18\x01 \x01(\x05R\x03uid\x12\x14\n" +
	"\x05nonce\x18\x02 \x01(\tR\x05nonce\"$\n" +
	"\bResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"X\n" +
	"\x04Pong\x12\x14\n" +
	"\x05nonce\x18\x01 \x01(\tR\x05nonce\x12\x12\n" +
	"\x04salt\x18\x02 \x01(\tR\x04salt\x12\x10\n" +
	"\x03num\x18\x03 \x01(\x05R\x03num\x12\x14\n" +
	"\x05noncs\x18\x04 \x01(\tR\x05noncs\",\n" +
	"\x11DeleteTaskRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId2p\n" +
	"\tTpService\x12(\n" +
	"\bPingPong\x12\r.netsend.Ping\x1a\r.netsend.Pong\x129\n" +
	"\n" +
	"TaskStream\x12\x14.netsend.TaskMessage\x1a\x11.netsend.Response(\x010\x01B\vZ\t./netgrpcb\x06proto3"

var (
	file_net_proto_rawDescOnce sync.Once
	file_net_proto_rawDescData []byte
)

func file_net_proto_rawDescGZIP() []byte {
	file_net_proto_rawDescOnce.Do(func() {
		file_net_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_net_proto_rawDesc), len(file_net_proto_rawDesc)))
	})
	return file_net_proto_rawDescData
}

var file_net_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_net_proto_goTypes = []any{
	(*TaskMessage)(nil),       // 0: netsend.TaskMessage
	(*Ping)(nil),              // 1: netsend.Ping
	(*Response)(nil),          // 2: netsend.Response
	(*Pong)(nil),              // 3: netsend.Pong
	(*DeleteTaskRequest)(nil), // 4: netsend.DeleteTaskRequest
}
var file_net_proto_depIdxs = []int32{
	1, // 0: netsend.TaskMessage.ping:type_name -> netsend.Ping
	1, // 1: netsend.TpService.PingPong:input_type -> netsend.Ping
	0, // 2: netsend.TpService.TaskStream:input_type -> netsend.TaskMessage
	3, // 3: netsend.TpService.PingPong:output_type -> netsend.Pong
	2, // 4: netsend.TpService.TaskStream:output_type -> netsend.Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_net_proto_init() }
func file_net_proto_init() {
	if File_net_proto != nil {
		return
	}
	file_net_proto_msgTypes[0].OneofWrappers = []any{
		(*TaskMessage_Ping)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_net_proto_rawDesc), len(file_net_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_net_proto_goTypes,
		DependencyIndexes: file_net_proto_depIdxs,
		MessageInfos:      file_net_proto_msgTypes,
	}.Build()
	File_net_proto = out.File
	file_net_proto_goTypes = nil
	file_net_proto_depIdxs = nil
}
