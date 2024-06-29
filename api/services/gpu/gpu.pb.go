// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gpu.proto

package gpu

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

type CheckpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *CheckpointRequest) Reset() {
	*x = CheckpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointRequest) ProtoMessage() {}

func (x *CheckpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointRequest.ProtoReflect.Descriptor instead.
func (*CheckpointRequest) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{0}
}

func (x *CheckpointRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type CheckpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	MemPath  string `protobuf:"bytes,2,opt,name=memPath,proto3" json:"memPath,omitempty"`
	CkptPath string `protobuf:"bytes,3,opt,name=ckptPath,proto3" json:"ckptPath,omitempty"`
}

func (x *CheckpointResponse) Reset() {
	*x = CheckpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointResponse) ProtoMessage() {}

func (x *CheckpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointResponse.ProtoReflect.Descriptor instead.
func (*CheckpointResponse) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{1}
}

func (x *CheckpointResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CheckpointResponse) GetMemPath() string {
	if x != nil {
		return x.MemPath
	}
	return ""
}

func (x *CheckpointResponse) GetCkptPath() string {
	if x != nil {
		return x.CkptPath
	}
	return ""
}

type RestoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *RestoreRequest) Reset() {
	*x = RestoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreRequest) ProtoMessage() {}

func (x *RestoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreRequest.ProtoReflect.Descriptor instead.
func (*RestoreRequest) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{2}
}

func (x *RestoreRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type RestoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RestoreResponse) Reset() {
	*x = RestoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreResponse) ProtoMessage() {}

func (x *RestoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreResponse.ProtoReflect.Descriptor instead.
func (*RestoreResponse) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{3}
}

func (x *RestoreResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type StartupPollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartupPollRequest) Reset() {
	*x = StartupPollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartupPollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartupPollRequest) ProtoMessage() {}

func (x *StartupPollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartupPollRequest.ProtoReflect.Descriptor instead.
func (*StartupPollRequest) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{4}
}

type StartupPollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StartupPollResponse) Reset() {
	*x = StartupPollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartupPollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartupPollResponse) ProtoMessage() {}

func (x *StartupPollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartupPollResponse.ProtoReflect.Descriptor instead.
func (*StartupPollResponse) Descriptor() ([]byte, []int) {
	return file_gpu_proto_rawDescGZIP(), []int{5}
}

func (x *StartupPollResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_gpu_proto protoreflect.FileDescriptor

var file_gpu_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x64,
	0x61, 0x6e, 0x61, 0x67, 0x70, 0x75, 0x22, 0x31, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x64, 0x0a, 0x12, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6b, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6b, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x2e, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x14, 0x0a, 0x12,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x6f, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xec, 0x01, 0x0a, 0x09, 0x43, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x47, 0x50,
	0x55, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x67, 0x70, 0x75, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x67, 0x70, 0x75, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x64, 0x61,
	0x6e, 0x61, 0x67, 0x70, 0x75, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x67, 0x70, 0x75,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x6f, 0x6c,
	0x6c, 0x12, 0x1d, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x67, 0x70, 0x75, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x67, 0x70, 0x75, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x70, 0x75, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gpu_proto_rawDescOnce sync.Once
	file_gpu_proto_rawDescData = file_gpu_proto_rawDesc
)

func file_gpu_proto_rawDescGZIP() []byte {
	file_gpu_proto_rawDescOnce.Do(func() {
		file_gpu_proto_rawDescData = protoimpl.X.CompressGZIP(file_gpu_proto_rawDescData)
	})
	return file_gpu_proto_rawDescData
}

var file_gpu_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gpu_proto_goTypes = []interface{}{
	(*CheckpointRequest)(nil),   // 0: cedanagpu.CheckpointRequest
	(*CheckpointResponse)(nil),  // 1: cedanagpu.CheckpointResponse
	(*RestoreRequest)(nil),      // 2: cedanagpu.RestoreRequest
	(*RestoreResponse)(nil),     // 3: cedanagpu.RestoreResponse
	(*StartupPollRequest)(nil),  // 4: cedanagpu.StartupPollRequest
	(*StartupPollResponse)(nil), // 5: cedanagpu.StartupPollResponse
}
var file_gpu_proto_depIdxs = []int32{
	0, // 0: cedanagpu.CedanaGPU.Checkpoint:input_type -> cedanagpu.CheckpointRequest
	2, // 1: cedanagpu.CedanaGPU.Restore:input_type -> cedanagpu.RestoreRequest
	4, // 2: cedanagpu.CedanaGPU.StartupPoll:input_type -> cedanagpu.StartupPollRequest
	1, // 3: cedanagpu.CedanaGPU.Checkpoint:output_type -> cedanagpu.CheckpointResponse
	3, // 4: cedanagpu.CedanaGPU.Restore:output_type -> cedanagpu.RestoreResponse
	5, // 5: cedanagpu.CedanaGPU.StartupPoll:output_type -> cedanagpu.StartupPollResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gpu_proto_init() }
func file_gpu_proto_init() {
	if File_gpu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gpu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointRequest); i {
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
		file_gpu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointResponse); i {
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
		file_gpu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreRequest); i {
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
		file_gpu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreResponse); i {
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
		file_gpu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartupPollRequest); i {
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
		file_gpu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartupPollResponse); i {
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
			RawDescriptor: file_gpu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gpu_proto_goTypes,
		DependencyIndexes: file_gpu_proto_depIdxs,
		MessageInfos:      file_gpu_proto_msgTypes,
	}.Build()
	File_gpu_proto = out.File
	file_gpu_proto_rawDesc = nil
	file_gpu_proto_goTypes = nil
	file_gpu_proto_depIdxs = nil
}
