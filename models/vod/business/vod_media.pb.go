// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: vod/business/vod_media.proto

package business

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type VodMediaBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName     string   `protobuf:"bytes,1,opt,name=SpaceName,proto3" json:"SpaceName,omitempty"`         //空间名
	Vid           string   `protobuf:"bytes,2,opt,name=Vid,proto3" json:"Vid,omitempty"`                     //视频ID
	Title         string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`                 //视频名称
	Description   string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`     //视频描述
	PosterUri     string   `protobuf:"bytes,5,opt,name=PosterUri,proto3" json:"PosterUri,omitempty"`         //封面图对象地址
	PublishStatus string   `protobuf:"bytes,6,opt,name=PublishStatus,proto3" json:"PublishStatus,omitempty"` //发布状态
	Tags          []string `protobuf:"bytes,7,rep,name=Tags,proto3" json:"Tags,omitempty"`                   //标签列表
	CreateTime    string   `protobuf:"bytes,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`       //创建时间
}

func (x *VodMediaBasicInfo) Reset() {
	*x = VodMediaBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodMediaBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodMediaBasicInfo) ProtoMessage() {}

func (x *VodMediaBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodMediaBasicInfo.ProtoReflect.Descriptor instead.
func (*VodMediaBasicInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_media_proto_rawDescGZIP(), []int{0}
}

func (x *VodMediaBasicInfo) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *VodMediaBasicInfo) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

func (x *VodMediaBasicInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VodMediaBasicInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VodMediaBasicInfo) GetPosterUri() string {
	if x != nil {
		return x.PosterUri
	}
	return ""
}

func (x *VodMediaBasicInfo) GetPublishStatus() string {
	if x != nil {
		return x.PublishStatus
	}
	return ""
}

func (x *VodMediaBasicInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *VodMediaBasicInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type VodMediaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicInfo      *VodMediaBasicInfo  `protobuf:"bytes,1,opt,name=BasicInfo,proto3" json:"BasicInfo,omitempty"`           //视频基础信息
	SourceInfo     *VodSourceInfo      `protobuf:"bytes,2,opt,name=SourceInfo,proto3" json:"SourceInfo,omitempty"`         //原视频信息
	TranscodeInfos []*VodTranscodeInfo `protobuf:"bytes,3,rep,name=TranscodeInfos,proto3" json:"TranscodeInfos,omitempty"` //转码视频信息列表
}

func (x *VodMediaInfo) Reset() {
	*x = VodMediaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodMediaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodMediaInfo) ProtoMessage() {}

func (x *VodMediaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodMediaInfo.ProtoReflect.Descriptor instead.
func (*VodMediaInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_media_proto_rawDescGZIP(), []int{1}
}

func (x *VodMediaInfo) GetBasicInfo() *VodMediaBasicInfo {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *VodMediaInfo) GetSourceInfo() *VodSourceInfo {
	if x != nil {
		return x.SourceInfo
	}
	return nil
}

func (x *VodMediaInfo) GetTranscodeInfos() []*VodTranscodeInfo {
	if x != nil {
		return x.TranscodeInfos
	}
	return nil
}

type VodGetMediaInfosData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaInfoList []*VodMediaInfo `protobuf:"bytes,1,rep,name=MediaInfoList,proto3" json:"MediaInfoList,omitempty"` //视频信息列表
	NotExistVids  []string        `protobuf:"bytes,2,rep,name=NotExistVids,proto3" json:"NotExistVids,omitempty"`   //不存在的视频VID列表
}

func (x *VodGetMediaInfosData) Reset() {
	*x = VodGetMediaInfosData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetMediaInfosData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetMediaInfosData) ProtoMessage() {}

func (x *VodGetMediaInfosData) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetMediaInfosData.ProtoReflect.Descriptor instead.
func (*VodGetMediaInfosData) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_media_proto_rawDescGZIP(), []int{2}
}

func (x *VodGetMediaInfosData) GetMediaInfoList() []*VodMediaInfo {
	if x != nil {
		return x.MediaInfoList
	}
	return nil
}

func (x *VodGetMediaInfosData) GetNotExistVids() []string {
	if x != nil {
		return x.NotExistVids
	}
	return nil
}

type VodStoreUriGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid       string   `protobuf:"bytes,1,opt,name=Vid,proto3" json:"Vid,omitempty"`             //视频ID
	StoreUris []string `protobuf:"bytes,2,rep,name=StoreUris,proto3" json:"StoreUris,omitempty"` //封面图对象地址列表
}

func (x *VodStoreUriGroup) Reset() {
	*x = VodStoreUriGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodStoreUriGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodStoreUriGroup) ProtoMessage() {}

func (x *VodStoreUriGroup) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodStoreUriGroup.ProtoReflect.Descriptor instead.
func (*VodStoreUriGroup) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_media_proto_rawDescGZIP(), []int{3}
}

func (x *VodStoreUriGroup) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

func (x *VodStoreUriGroup) GetStoreUris() []string {
	if x != nil {
		return x.StoreUris
	}
	return nil
}

type VodGetRecPosterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreUriGroups []*VodStoreUriGroup `protobuf:"bytes,1,rep,name=StoreUriGroups,proto3" json:"StoreUriGroups,omitempty"` //封面图信息
	NotExistVids   []string            `protobuf:"bytes,2,rep,name=NotExistVids,proto3" json:"NotExistVids,omitempty"`     //不存在的视频VID列表
}

func (x *VodGetRecPosterData) Reset() {
	*x = VodGetRecPosterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetRecPosterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetRecPosterData) ProtoMessage() {}

func (x *VodGetRecPosterData) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetRecPosterData.ProtoReflect.Descriptor instead.
func (*VodGetRecPosterData) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_media_proto_rawDescGZIP(), []int{4}
}

func (x *VodGetRecPosterData) GetStoreUriGroups() []*VodStoreUriGroup {
	if x != nil {
		return x.StoreUriGroups
	}
	return nil
}

func (x *VodGetRecPosterData) GetNotExistVids() []string {
	if x != nil {
		return x.NotExistVids
	}
	return nil
}

var File_vod_business_vod_media_proto protoreflect.FileDescriptor

var file_vod_business_vod_media_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x1d,
	0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01,
	0x0a, 0x11, 0x56, 0x6f, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x56, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x0c, 0x56, 0x6f, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f,
	0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x58, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x8e,
	0x01, 0x0a, 0x14, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x56, 0x6f, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x22,
	0x42, 0x0a, 0x10, 0x56, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x56, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x72,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55,
	0x72, 0x69, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x0e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x56, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x4e, 0x6f, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x73, 0x42, 0xac, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42,
	0x08, 0x56, 0x6f, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca,
	0x02, 0x18, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x6f,
	0x64, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x1b, 0x56, 0x6f, 0x6c,
	0x63, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_media_proto_rawDescOnce sync.Once
	file_vod_business_vod_media_proto_rawDescData = file_vod_business_vod_media_proto_rawDesc
)

func file_vod_business_vod_media_proto_rawDescGZIP() []byte {
	file_vod_business_vod_media_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_media_proto_rawDescData)
	})
	return file_vod_business_vod_media_proto_rawDescData
}

var file_vod_business_vod_media_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vod_business_vod_media_proto_goTypes = []interface{}{
	(*VodMediaBasicInfo)(nil),    // 0: Volcengine.Models.Vod.Business.VodMediaBasicInfo
	(*VodMediaInfo)(nil),         // 1: Volcengine.Models.Vod.Business.VodMediaInfo
	(*VodGetMediaInfosData)(nil), // 2: Volcengine.Models.Vod.Business.VodGetMediaInfosData
	(*VodStoreUriGroup)(nil),     // 3: Volcengine.Models.Vod.Business.VodStoreUriGroup
	(*VodGetRecPosterData)(nil),  // 4: Volcengine.Models.Vod.Business.VodGetRecPosterData
	(*VodSourceInfo)(nil),        // 5: Volcengine.Models.Vod.Business.VodSourceInfo
	(*VodTranscodeInfo)(nil),     // 6: Volcengine.Models.Vod.Business.VodTranscodeInfo
}
var file_vod_business_vod_media_proto_depIdxs = []int32{
	0, // 0: Volcengine.Models.Vod.Business.VodMediaInfo.BasicInfo:type_name -> Volcengine.Models.Vod.Business.VodMediaBasicInfo
	5, // 1: Volcengine.Models.Vod.Business.VodMediaInfo.SourceInfo:type_name -> Volcengine.Models.Vod.Business.VodSourceInfo
	6, // 2: Volcengine.Models.Vod.Business.VodMediaInfo.TranscodeInfos:type_name -> Volcengine.Models.Vod.Business.VodTranscodeInfo
	1, // 3: Volcengine.Models.Vod.Business.VodGetMediaInfosData.MediaInfoList:type_name -> Volcengine.Models.Vod.Business.VodMediaInfo
	3, // 4: Volcengine.Models.Vod.Business.VodGetRecPosterData.StoreUriGroups:type_name -> Volcengine.Models.Vod.Business.VodStoreUriGroup
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vod_business_vod_media_proto_init() }
func file_vod_business_vod_media_proto_init() {
	if File_vod_business_vod_media_proto != nil {
		return
	}
	file_vod_business_vod_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodMediaBasicInfo); i {
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
		file_vod_business_vod_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodMediaInfo); i {
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
		file_vod_business_vod_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetMediaInfosData); i {
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
		file_vod_business_vod_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodStoreUriGroup); i {
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
		file_vod_business_vod_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetRecPosterData); i {
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
			RawDescriptor: file_vod_business_vod_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_media_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_media_proto_depIdxs,
		MessageInfos:      file_vod_business_vod_media_proto_msgTypes,
	}.Build()
	File_vod_business_vod_media_proto = out.File
	file_vod_business_vod_media_proto_rawDesc = nil
	file_vod_business_vod_media_proto_goTypes = nil
	file_vod_business_vod_media_proto_depIdxs = nil
}
