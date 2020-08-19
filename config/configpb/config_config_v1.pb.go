// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: api/protobuf-spec/config_config_v1.proto

package configpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestAppVersion        string   `protobuf:"bytes,1,opt,name=latest_app_version,json=latestAppVersion,proto3" json:"latest_app_version,omitempty"`
	MinCompatibleAppVersion string   `protobuf:"bytes,2,opt,name=min_compatible_app_version,json=minCompatibleAppVersion,proto3" json:"min_compatible_app_version,omitempty"`
	SupportEmail            string   `protobuf:"bytes,3,opt,name=support_email,json=supportEmail,proto3" json:"support_email,omitempty"`
	ApiBaseUrl              string   `protobuf:"bytes,4,opt,name=api_base_url,json=apiBaseUrl,proto3" json:"api_base_url,omitempty"`
	AvatarPushBaseUrl       string   `protobuf:"bytes,5,opt,name=avatar_push_base_url,json=avatarPushBaseUrl,proto3" json:"avatar_push_base_url,omitempty"`
	IosUpdateUrl            string   `protobuf:"bytes,6,opt,name=ios_update_url,json=iosUpdateUrl,proto3" json:"ios_update_url,omitempty"`
	AndroidUpdateUrl        string   `protobuf:"bytes,7,opt,name=android_update_url,json=androidUpdateUrl,proto3" json:"android_update_url,omitempty"`
	InviteCodeBaseUrl       string   `protobuf:"bytes,8,opt,name=invite_code_base_url,json=inviteCodeBaseUrl,proto3" json:"invite_code_base_url,omitempty"`
	InviteCodeBaseUrls      []string `protobuf:"bytes,9,rep,name=invite_code_base_urls,json=inviteCodeBaseUrls,proto3" json:"invite_code_base_urls,omitempty"`
	OnboardingImages        []string `protobuf:"bytes,10,rep,name=onboarding_images,json=onboardingImages,proto3" json:"onboarding_images,omitempty"`
	SupportedCountries      []string `protobuf:"bytes,11,rep,name=supported_countries,json=supportedCountries,proto3" json:"supported_countries,omitempty"`
	LatestAppUpdated        string   `protobuf:"bytes,12,opt,name=latest_app_updated,json=latestAppUpdated,proto3" json:"latest_app_updated,omitempty"`
	TourSwipeGuide          []string `protobuf:"bytes,13,rep,name=tour_swipe_guide,json=tourSwipeGuide,proto3" json:"tour_swipe_guide,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLatestAppVersion() string {
	if x != nil {
		return x.LatestAppVersion
	}
	return ""
}

func (x *Config) GetMinCompatibleAppVersion() string {
	if x != nil {
		return x.MinCompatibleAppVersion
	}
	return ""
}

func (x *Config) GetSupportEmail() string {
	if x != nil {
		return x.SupportEmail
	}
	return ""
}

func (x *Config) GetApiBaseUrl() string {
	if x != nil {
		return x.ApiBaseUrl
	}
	return ""
}

func (x *Config) GetAvatarPushBaseUrl() string {
	if x != nil {
		return x.AvatarPushBaseUrl
	}
	return ""
}

func (x *Config) GetIosUpdateUrl() string {
	if x != nil {
		return x.IosUpdateUrl
	}
	return ""
}

func (x *Config) GetAndroidUpdateUrl() string {
	if x != nil {
		return x.AndroidUpdateUrl
	}
	return ""
}

func (x *Config) GetInviteCodeBaseUrl() string {
	if x != nil {
		return x.InviteCodeBaseUrl
	}
	return ""
}

func (x *Config) GetInviteCodeBaseUrls() []string {
	if x != nil {
		return x.InviteCodeBaseUrls
	}
	return nil
}

func (x *Config) GetOnboardingImages() []string {
	if x != nil {
		return x.OnboardingImages
	}
	return nil
}

func (x *Config) GetSupportedCountries() []string {
	if x != nil {
		return x.SupportedCountries
	}
	return nil
}

func (x *Config) GetLatestAppUpdated() string {
	if x != nil {
		return x.LatestAppUpdated
	}
	return ""
}

func (x *Config) GetTourSwipeGuide() []string {
	if x != nil {
		return x.TourSwipeGuide
	}
	return nil
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	EnglishName string `protobuf:"bytes,2,opt,name=english_name,json=englishName,proto3" json:"english_name,omitempty"`
	NativeName  string `protobuf:"bytes,3,opt,name=native_name,json=nativeName,proto3" json:"native_name,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{1}
}

func (x *Language) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Language) GetEnglishName() string {
	if x != nil {
		return x.EnglishName
	}
	return ""
}

func (x *Language) GetNativeName() string {
	if x != nil {
		return x.NativeName
	}
	return ""
}

type CodeOfConduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headline    string `protobuf:"bytes,1,opt,name=headline,proto3" json:"headline,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IconUrl     string `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
}

func (x *CodeOfConduct) Reset() {
	*x = CodeOfConduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeOfConduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeOfConduct) ProtoMessage() {}

func (x *CodeOfConduct) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeOfConduct.ProtoReflect.Descriptor instead.
func (*CodeOfConduct) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CodeOfConduct) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *CodeOfConduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CodeOfConduct) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

type GetAppConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetAppConfigResponse) Reset() {
	*x = GetAppConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppConfigResponse) ProtoMessage() {}

func (x *GetAppConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppConfigResponse.ProtoReflect.Descriptor instead.
func (*GetAppConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppConfigResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type GetLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*Language `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *GetLanguagesResponse) Reset() {
	*x = GetLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLanguagesResponse) ProtoMessage() {}

func (x *GetLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLanguagesResponse.ProtoReflect.Descriptor instead.
func (*GetLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetLanguagesResponse) GetLanguages() []*Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

type GetCodeOfConductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*CodeOfConduct `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *GetCodeOfConductResponse) Reset() {
	*x = GetCodeOfConductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeOfConductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeOfConductResponse) ProtoMessage() {}

func (x *GetCodeOfConductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_config_config_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeOfConductResponse.ProtoReflect.Descriptor instead.
func (*GetCodeOfConductResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetCodeOfConductResponse) GetCodes() []*CodeOfConduct {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_api_protobuf_spec_config_config_v1_proto protoreflect.FileDescriptor

var file_api_protobuf_spec_config_config_v1_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x0a, 0x1a, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c,
	0x65, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x42, 0x61, 0x73, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6f, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6f,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x41, 0x70,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x75, 0x72,
	0x5f, 0x73, 0x77, 0x69, 0x70, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x53, 0x77, 0x69, 0x70, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x22, 0x62, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x0d, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x66,
	0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c,
	0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x8d, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_spec_config_config_v1_proto_rawDescOnce sync.Once
	file_api_protobuf_spec_config_config_v1_proto_rawDescData = file_api_protobuf_spec_config_config_v1_proto_rawDesc
)

func file_api_protobuf_spec_config_config_v1_proto_rawDescGZIP() []byte {
	file_api_protobuf_spec_config_config_v1_proto_rawDescOnce.Do(func() {
		file_api_protobuf_spec_config_config_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_spec_config_config_v1_proto_rawDescData)
	})
	return file_api_protobuf_spec_config_config_v1_proto_rawDescData
}

var file_api_protobuf_spec_config_config_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_protobuf_spec_config_config_v1_proto_goTypes = []interface{}{
	(*Config)(nil),                   // 0: config.config.Config
	(*Language)(nil),                 // 1: config.config.Language
	(*CodeOfConduct)(nil),            // 2: config.config.CodeOfConduct
	(*GetAppConfigResponse)(nil),     // 3: config.config.GetAppConfigResponse
	(*GetLanguagesResponse)(nil),     // 4: config.config.GetLanguagesResponse
	(*GetCodeOfConductResponse)(nil), // 5: config.config.GetCodeOfConductResponse
	(*empty.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_api_protobuf_spec_config_config_v1_proto_depIdxs = []int32{
	0, // 0: config.config.GetAppConfigResponse.config:type_name -> config.config.Config
	1, // 1: config.config.GetLanguagesResponse.languages:type_name -> config.config.Language
	2, // 2: config.config.GetCodeOfConductResponse.codes:type_name -> config.config.CodeOfConduct
	6, // 3: config.config.ConfigService.GetAppConfig:input_type -> google.protobuf.Empty
	6, // 4: config.config.ConfigService.GetLanguagesSupported:input_type -> google.protobuf.Empty
	6, // 5: config.config.ConfigService.GetCodeOfConduct:input_type -> google.protobuf.Empty
	3, // 6: config.config.ConfigService.GetAppConfig:output_type -> config.config.GetAppConfigResponse
	4, // 7: config.config.ConfigService.GetLanguagesSupported:output_type -> config.config.GetLanguagesResponse
	5, // 8: config.config.ConfigService.GetCodeOfConduct:output_type -> config.config.GetCodeOfConductResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_protobuf_spec_config_config_v1_proto_init() }
func file_api_protobuf_spec_config_config_v1_proto_init() {
	if File_api_protobuf_spec_config_config_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeOfConduct); i {
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
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppConfigResponse); i {
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
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLanguagesResponse); i {
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
		file_api_protobuf_spec_config_config_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeOfConductResponse); i {
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
			RawDescriptor: file_api_protobuf_spec_config_config_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protobuf_spec_config_config_v1_proto_goTypes,
		DependencyIndexes: file_api_protobuf_spec_config_config_v1_proto_depIdxs,
		MessageInfos:      file_api_protobuf_spec_config_config_v1_proto_msgTypes,
	}.Build()
	File_api_protobuf_spec_config_config_v1_proto = out.File
	file_api_protobuf_spec_config_config_v1_proto_rawDesc = nil
	file_api_protobuf_spec_config_config_v1_proto_goTypes = nil
	file_api_protobuf_spec_config_config_v1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	GetAppConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppConfigResponse, error)
	GetLanguagesSupported(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLanguagesResponse, error)
	GetCodeOfConduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCodeOfConductResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) GetAppConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppConfigResponse, error) {
	out := new(GetAppConfigResponse)
	err := c.cc.Invoke(ctx, "/config.config.ConfigService/GetAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetLanguagesSupported(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLanguagesResponse, error) {
	out := new(GetLanguagesResponse)
	err := c.cc.Invoke(ctx, "/config.config.ConfigService/GetLanguagesSupported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetCodeOfConduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCodeOfConductResponse, error) {
	out := new(GetCodeOfConductResponse)
	err := c.cc.Invoke(ctx, "/config.config.ConfigService/GetCodeOfConduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	GetAppConfig(context.Context, *empty.Empty) (*GetAppConfigResponse, error)
	GetLanguagesSupported(context.Context, *empty.Empty) (*GetLanguagesResponse, error)
	GetCodeOfConduct(context.Context, *empty.Empty) (*GetCodeOfConductResponse, error)
}

// UnimplementedConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (*UnimplementedConfigServiceServer) GetAppConfig(context.Context, *empty.Empty) (*GetAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}
func (*UnimplementedConfigServiceServer) GetLanguagesSupported(context.Context, *empty.Empty) (*GetLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguagesSupported not implemented")
}
func (*UnimplementedConfigServiceServer) GetCodeOfConduct(context.Context, *empty.Empty) (*GetCodeOfConductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeOfConduct not implemented")
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.config.ConfigService/GetAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetAppConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetLanguagesSupported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetLanguagesSupported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.config.ConfigService/GetLanguagesSupported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetLanguagesSupported(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetCodeOfConduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetCodeOfConduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.config.ConfigService/GetCodeOfConduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetCodeOfConduct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppConfig",
			Handler:    _ConfigService_GetAppConfig_Handler,
		},
		{
			MethodName: "GetLanguagesSupported",
			Handler:    _ConfigService_GetLanguagesSupported_Handler,
		},
		{
			MethodName: "GetCodeOfConduct",
			Handler:    _ConfigService_GetCodeOfConduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf-spec/config_config_v1.proto",
}
