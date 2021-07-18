/// This file has messages for describing gitlab projects

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: external/gitlab/projects/projects.v1.proto

package projects

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

//*
// Represents a Gitlab Project ID
type ProjectID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Gitlab Project ID
}

func (x *ProjectID) Reset() {
	*x = ProjectID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectID) ProtoMessage() {}

func (x *ProjectID) ProtoReflect() protoreflect.Message {
	mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectID.ProtoReflect.Descriptor instead.
func (*ProjectID) Descriptor() ([]byte, []int) {
	return file_external_gitlab_projects_projects_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//*
// Represents a Gitlab Project name
type ProjectName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Gitlab Project name
}

func (x *ProjectName) Reset() {
	*x = ProjectName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectName) ProtoMessage() {}

func (x *ProjectName) ProtoReflect() protoreflect.Message {
	mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectName.ProtoReflect.Descriptor instead.
func (*ProjectName) Descriptor() ([]byte, []int) {
	return file_external_gitlab_projects_projects_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//*
// Represents a Gitlab Project
type ProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               // Gitlab Project ID
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                            // Gitlab Project name
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`              // Gitlab Project description
	WebUrl      string `protobuf:"bytes,4,opt,name=web_url,json=webUrl,proto3" json:"web_url,omitempty"`          // Gitlab Project URL
	AvatarUrl   string `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // Gitlab Project avatar URL
	ReadmeUrl   string `protobuf:"bytes,6,opt,name=readme_url,json=readmeUrl,proto3" json:"readme_url,omitempty"` // // Gitlab Project readme ID
}

func (x *ProjectInfo) Reset() {
	*x = ProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInfo) ProtoMessage() {}

func (x *ProjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_external_gitlab_projects_projects_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInfo.ProtoReflect.Descriptor instead.
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return file_external_gitlab_projects_projects_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectInfo) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *ProjectInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *ProjectInfo) GetReadmeUrl() string {
	if x != nil {
		return x.ReadmeUrl
	}
	return ""
}

var File_external_gitlab_projects_projects_v1_proto protoreflect.FileDescriptor

var file_external_gitlab_projects_projects_v1_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x6e,
	0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x55, 0x72, 0x6c, 0x32, 0xa1, 0x01,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x6e,
	0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x1f, 0x2e, 0x65,
	0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x64, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x65,
	0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_gitlab_projects_projects_v1_proto_rawDescOnce sync.Once
	file_external_gitlab_projects_projects_v1_proto_rawDescData = file_external_gitlab_projects_projects_v1_proto_rawDesc
)

func file_external_gitlab_projects_projects_v1_proto_rawDescGZIP() []byte {
	file_external_gitlab_projects_projects_v1_proto_rawDescOnce.Do(func() {
		file_external_gitlab_projects_projects_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_gitlab_projects_projects_v1_proto_rawDescData)
	})
	return file_external_gitlab_projects_projects_v1_proto_rawDescData
}

var file_external_gitlab_projects_projects_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_gitlab_projects_projects_v1_proto_goTypes = []interface{}{
	(*ProjectID)(nil),   // 0: envspotting_gitlab.ProjectID
	(*ProjectName)(nil), // 1: envspotting_gitlab.ProjectName
	(*ProjectInfo)(nil), // 2: envspotting_gitlab.ProjectInfo
}
var file_external_gitlab_projects_projects_v1_proto_depIdxs = []int32{
	0, // 0: envspotting_gitlab.Projects.Get:input_type -> envspotting_gitlab.ProjectID
	1, // 1: envspotting_gitlab.Projects.List:input_type -> envspotting_gitlab.ProjectName
	2, // 2: envspotting_gitlab.Projects.Get:output_type -> envspotting_gitlab.ProjectInfo
	2, // 3: envspotting_gitlab.Projects.List:output_type -> envspotting_gitlab.ProjectInfo
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_gitlab_projects_projects_v1_proto_init() }
func file_external_gitlab_projects_projects_v1_proto_init() {
	if File_external_gitlab_projects_projects_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_gitlab_projects_projects_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectID); i {
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
		file_external_gitlab_projects_projects_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectName); i {
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
		file_external_gitlab_projects_projects_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInfo); i {
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
			RawDescriptor: file_external_gitlab_projects_projects_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_gitlab_projects_projects_v1_proto_goTypes,
		DependencyIndexes: file_external_gitlab_projects_projects_v1_proto_depIdxs,
		MessageInfos:      file_external_gitlab_projects_projects_v1_proto_msgTypes,
	}.Build()
	File_external_gitlab_projects_projects_v1_proto = out.File
	file_external_gitlab_projects_projects_v1_proto_rawDesc = nil
	file_external_gitlab_projects_projects_v1_proto_goTypes = nil
	file_external_gitlab_projects_projects_v1_proto_depIdxs = nil
}
