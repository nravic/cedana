// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: runc.proto

package runc

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

type DumpOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root   string `protobuf:"bytes,1,opt,name=Root,proto3" json:"Root,omitempty"`
	ID     string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Bundle string `protobuf:"bytes,3,opt,name=Bundle,proto3" json:"Bundle,omitempty"`
}

func (x *DumpOpts) Reset() {
	*x = DumpOpts{}
	mi := &file_runc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DumpOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpOpts) ProtoMessage() {}

func (x *DumpOpts) ProtoReflect() protoreflect.Message {
	mi := &file_runc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpOpts.ProtoReflect.Descriptor instead.
func (*DumpOpts) Descriptor() ([]byte, []int) {
	return file_runc_proto_rawDescGZIP(), []int{0}
}

func (x *DumpOpts) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *DumpOpts) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DumpOpts) GetBundle() string {
	if x != nil {
		return x.Bundle
	}
	return ""
}

type RestoreOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root   string `protobuf:"bytes,1,opt,name=Root,proto3" json:"Root,omitempty"`
	ID     string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Bundle string `protobuf:"bytes,3,opt,name=Bundle,proto3" json:"Bundle,omitempty"`
}

func (x *RestoreOpts) Reset() {
	*x = RestoreOpts{}
	mi := &file_runc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestoreOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreOpts) ProtoMessage() {}

func (x *RestoreOpts) ProtoReflect() protoreflect.Message {
	mi := &file_runc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreOpts.ProtoReflect.Descriptor instead.
func (*RestoreOpts) Descriptor() ([]byte, []int) {
	return file_runc_proto_rawDescGZIP(), []int{1}
}

func (x *RestoreOpts) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *RestoreOpts) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RestoreOpts) GetBundle() string {
	if x != nil {
		return x.Bundle
	}
	return ""
}

var File_runc_proto protoreflect.FileDescriptor

var file_runc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x65,
	0x64, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x72, 0x75, 0x6e,
	0x63, 0x22, 0x46, 0x0a, 0x08, 0x44, 0x75, 0x6d, 0x70, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runc_proto_rawDescOnce sync.Once
	file_runc_proto_rawDescData = file_runc_proto_rawDesc
)

func file_runc_proto_rawDescGZIP() []byte {
	file_runc_proto_rawDescOnce.Do(func() {
		file_runc_proto_rawDescData = protoimpl.X.CompressGZIP(file_runc_proto_rawDescData)
	})
	return file_runc_proto_rawDescData
}

var file_runc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_runc_proto_goTypes = []any{
	(*DumpOpts)(nil),    // 0: cedana.plugins.runc.DumpOpts
	(*RestoreOpts)(nil), // 1: cedana.plugins.runc.RestoreOpts
}
var file_runc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runc_proto_init() }
func file_runc_proto_init() {
	if File_runc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runc_proto_goTypes,
		DependencyIndexes: file_runc_proto_depIdxs,
		MessageInfos:      file_runc_proto_msgTypes,
	}.Build()
	File_runc_proto = out.File
	file_runc_proto_rawDesc = nil
	file_runc_proto_goTypes = nil
	file_runc_proto_depIdxs = nil
}
