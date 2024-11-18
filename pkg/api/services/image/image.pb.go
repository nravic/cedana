//  Copyright 2020 Two Sigma Investments, LP.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: image.proto

package image

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

type Marker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// Types that are assignable to Body:
	//
	//	*Marker_Filename
	//	*Marker_FileData
	//	*Marker_FileEof
	//	*Marker_ImageEof
	Body isMarker_Body `protobuf_oneof:"body"`
}

func (x *Marker) Reset() {
	*x = Marker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Marker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Marker) ProtoMessage() {}

func (x *Marker) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Marker.ProtoReflect.Descriptor instead.
func (*Marker) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *Marker) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (m *Marker) GetBody() isMarker_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Marker) GetFilename() string {
	if x, ok := x.GetBody().(*Marker_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *Marker) GetFileData() uint32 {
	if x, ok := x.GetBody().(*Marker_FileData); ok {
		return x.FileData
	}
	return 0
}

func (x *Marker) GetFileEof() bool {
	if x, ok := x.GetBody().(*Marker_FileEof); ok {
		return x.FileEof
	}
	return false
}

func (x *Marker) GetImageEof() bool {
	if x, ok := x.GetBody().(*Marker_ImageEof); ok {
		return x.ImageEof
	}
	return false
}

type isMarker_Body interface {
	isMarker_Body()
}

type Marker_Filename struct {
	// Denotes the filename of the next upcoming markers (denoted as current file)
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3,oneof"`
}

type Marker_FileData struct {
	// Incoming data for the current file
	FileData uint32 `protobuf:"varint,3,opt,name=file_data,json=fileData,proto3,oneof"`
}

type Marker_FileEof struct {
	// EOF of current file is reached
	FileEof bool `protobuf:"varint,4,opt,name=file_eof,json=fileEof,proto3,oneof"`
}

type Marker_ImageEof struct {
	// EOF of image is reached
	ImageEof bool `protobuf:"varint,5,opt,name=image_eof,json=imageEof,proto3,oneof"`
}

func (*Marker_Filename) isMarker_Body() {}

func (*Marker_FileData) isMarker_Body() {}

func (*Marker_FileEof) isMarker_Body() {}

func (*Marker_ImageEof) isMarker_Body() {}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x65, 0x64, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x6f, 0x66, 0x12, 0x1d, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x6f, 0x66, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x65, 0x64, 0x61, 0x6e, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_image_proto_goTypes = []interface{}{
	(*Marker)(nil), // 0: cedana.services.image.marker
}
var file_image_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Marker); i {
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
	file_image_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Marker_Filename)(nil),
		(*Marker_FileData)(nil),
		(*Marker_FileEof)(nil),
		(*Marker_ImageEof)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
