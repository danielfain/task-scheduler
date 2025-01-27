// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/cron.proto

package proto

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

type AddTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    string                 `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"` // Cron expression
	Command       string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`       // Shell command to execute
	Metadata      []byte                 `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`     // Optional JSON metadata
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	mi := &file_proto_cron_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cron_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_cron_proto_rawDescGZIP(), []int{0}
}

func (x *AddTaskRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *AddTaskRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *AddTaskRequest) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Expression    string                 `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	Command       string                 `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	NextRun       string                 `protobuf:"bytes,4,opt,name=next_run,json=nextRun,proto3" json:"next_run,omitempty"` // timestamp
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_proto_cron_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cron_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_cron_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *Task) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Task) GetNextRun() string {
	if x != nil {
		return x.NextRun
	}
	return ""
}

func (x *Task) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type TaskId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskId) Reset() {
	*x = TaskId{}
	mi := &file_proto_cron_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskId) ProtoMessage() {}

func (x *TaskId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cron_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskId.ProtoReflect.Descriptor instead.
func (*TaskId) Descriptor() ([]byte, []int) {
	return file_proto_cron_proto_rawDescGZIP(), []int{2}
}

func (x *TaskId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_cron_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cron_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_cron_proto_rawDescGZIP(), []int{3}
}

var File_proto_cron_proto protoreflect.FileDescriptor

var file_proto_cron_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x88, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x95,
	0x01, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x63, 0x72,
	0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x0c, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x1a, 0x0b, 0x2e,
	0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x63, 0x72, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_cron_proto_rawDescOnce sync.Once
	file_proto_cron_proto_rawDescData []byte
)

func file_proto_cron_proto_rawDescGZIP() []byte {
	file_proto_cron_proto_rawDescOnce.Do(func() {
		file_proto_cron_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cron_proto_rawDesc), len(file_proto_cron_proto_rawDesc)))
	})
	return file_proto_cron_proto_rawDescData
}

var file_proto_cron_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_cron_proto_goTypes = []any{
	(*AddTaskRequest)(nil), // 0: cron.AddTaskRequest
	(*Task)(nil),           // 1: cron.Task
	(*TaskId)(nil),         // 2: cron.TaskId
	(*Empty)(nil),          // 3: cron.Empty
}
var file_proto_cron_proto_depIdxs = []int32{
	0, // 0: cron.TaskScheduler.AddTask:input_type -> cron.AddTaskRequest
	2, // 1: cron.TaskScheduler.DeleteTask:input_type -> cron.TaskId
	3, // 2: cron.TaskScheduler.ListTasks:input_type -> cron.Empty
	2, // 3: cron.TaskScheduler.AddTask:output_type -> cron.TaskId
	3, // 4: cron.TaskScheduler.DeleteTask:output_type -> cron.Empty
	1, // 5: cron.TaskScheduler.ListTasks:output_type -> cron.Task
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cron_proto_init() }
func file_proto_cron_proto_init() {
	if File_proto_cron_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cron_proto_rawDesc), len(file_proto_cron_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cron_proto_goTypes,
		DependencyIndexes: file_proto_cron_proto_depIdxs,
		MessageInfos:      file_proto_cron_proto_msgTypes,
	}.Build()
	File_proto_cron_proto = out.File
	file_proto_cron_proto_goTypes = nil
	file_proto_cron_proto_depIdxs = nil
}
