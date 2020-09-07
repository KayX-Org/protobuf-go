// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: api/protobuf-spec/notification_notification_v1.proto

package notificationpb

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

type SendEmailVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendEmailVerificationRequest) Reset() {
	*x = SendEmailVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailVerificationRequest) ProtoMessage() {}

func (x *SendEmailVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailVerificationRequest.ProtoReflect.Descriptor instead.
func (*SendEmailVerificationRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_notification_notification_v1_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailVerificationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailVerificationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SendSMSVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendSMSVerificationRequest) Reset() {
	*x = SendSMSVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSVerificationRequest) ProtoMessage() {}

func (x *SendSMSVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSVerificationRequest.ProtoReflect.Descriptor instead.
func (*SendSMSVerificationRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_notification_notification_v1_proto_rawDescGZIP(), []int{1}
}

func (x *SendSMSVerificationRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SendSMSVerificationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UpdateFirebaseTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UpdateFirebaseTokenRequest) Reset() {
	*x = UpdateFirebaseTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFirebaseTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFirebaseTokenRequest) ProtoMessage() {}

func (x *UpdateFirebaseTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFirebaseTokenRequest.ProtoReflect.Descriptor instead.
func (*UpdateFirebaseTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_notification_notification_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFirebaseTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateFirebaseTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_api_protobuf_spec_notification_notification_v1_proto protoreflect.FileDescriptor

var file_api_protobuf_spec_notification_notification_v1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48,
	0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x4d, 0x53, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a,
	0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd1, 0x02, 0x0a, 0x13, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6a, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66,
	0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_spec_notification_notification_v1_proto_rawDescOnce sync.Once
	file_api_protobuf_spec_notification_notification_v1_proto_rawDescData = file_api_protobuf_spec_notification_notification_v1_proto_rawDesc
)

func file_api_protobuf_spec_notification_notification_v1_proto_rawDescGZIP() []byte {
	file_api_protobuf_spec_notification_notification_v1_proto_rawDescOnce.Do(func() {
		file_api_protobuf_spec_notification_notification_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_spec_notification_notification_v1_proto_rawDescData)
	})
	return file_api_protobuf_spec_notification_notification_v1_proto_rawDescData
}

var file_api_protobuf_spec_notification_notification_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_protobuf_spec_notification_notification_v1_proto_goTypes = []interface{}{
	(*SendEmailVerificationRequest)(nil), // 0: notification.notification.SendEmailVerificationRequest
	(*SendSMSVerificationRequest)(nil),   // 1: notification.notification.SendSMSVerificationRequest
	(*UpdateFirebaseTokenRequest)(nil),   // 2: notification.notification.UpdateFirebaseTokenRequest
	(*empty.Empty)(nil),                  // 3: google.protobuf.Empty
}
var file_api_protobuf_spec_notification_notification_v1_proto_depIdxs = []int32{
	0, // 0: notification.notification.NotificationService.SendEmailVerification:input_type -> notification.notification.SendEmailVerificationRequest
	1, // 1: notification.notification.NotificationService.SendSMSVerification:input_type -> notification.notification.SendSMSVerificationRequest
	2, // 2: notification.notification.NotificationService.UpdateFirebaseToken:input_type -> notification.notification.UpdateFirebaseTokenRequest
	3, // 3: notification.notification.NotificationService.SendEmailVerification:output_type -> google.protobuf.Empty
	3, // 4: notification.notification.NotificationService.SendSMSVerification:output_type -> google.protobuf.Empty
	3, // 5: notification.notification.NotificationService.UpdateFirebaseToken:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protobuf_spec_notification_notification_v1_proto_init() }
func file_api_protobuf_spec_notification_notification_v1_proto_init() {
	if File_api_protobuf_spec_notification_notification_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailVerificationRequest); i {
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
		file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSVerificationRequest); i {
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
		file_api_protobuf_spec_notification_notification_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFirebaseTokenRequest); i {
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
			RawDescriptor: file_api_protobuf_spec_notification_notification_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protobuf_spec_notification_notification_v1_proto_goTypes,
		DependencyIndexes: file_api_protobuf_spec_notification_notification_v1_proto_depIdxs,
		MessageInfos:      file_api_protobuf_spec_notification_notification_v1_proto_msgTypes,
	}.Build()
	File_api_protobuf_spec_notification_notification_v1_proto = out.File
	file_api_protobuf_spec_notification_notification_v1_proto_rawDesc = nil
	file_api_protobuf_spec_notification_notification_v1_proto_goTypes = nil
	file_api_protobuf_spec_notification_notification_v1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	SendEmailVerification(ctx context.Context, in *SendEmailVerificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendSMSVerification(ctx context.Context, in *SendSMSVerificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateFirebaseToken(ctx context.Context, in *UpdateFirebaseTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SendEmailVerification(ctx context.Context, in *SendEmailVerificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notification.notification.NotificationService/SendEmailVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendSMSVerification(ctx context.Context, in *SendSMSVerificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notification.notification.NotificationService/SendSMSVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) UpdateFirebaseToken(ctx context.Context, in *UpdateFirebaseTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notification.notification.NotificationService/UpdateFirebaseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	SendEmailVerification(context.Context, *SendEmailVerificationRequest) (*empty.Empty, error)
	SendSMSVerification(context.Context, *SendSMSVerificationRequest) (*empty.Empty, error)
	UpdateFirebaseToken(context.Context, *UpdateFirebaseTokenRequest) (*empty.Empty, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) SendEmailVerification(context.Context, *SendEmailVerificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailVerification not implemented")
}
func (*UnimplementedNotificationServiceServer) SendSMSVerification(context.Context, *SendSMSVerificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMSVerification not implemented")
}
func (*UnimplementedNotificationServiceServer) UpdateFirebaseToken(context.Context, *UpdateFirebaseTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFirebaseToken not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_SendEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.notification.NotificationService/SendEmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendEmailVerification(ctx, req.(*SendEmailVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendSMSVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendSMSVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.notification.NotificationService/SendSMSVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendSMSVerification(ctx, req.(*SendSMSVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_UpdateFirebaseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFirebaseTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).UpdateFirebaseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.notification.NotificationService/UpdateFirebaseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).UpdateFirebaseToken(ctx, req.(*UpdateFirebaseTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailVerification",
			Handler:    _NotificationService_SendEmailVerification_Handler,
		},
		{
			MethodName: "SendSMSVerification",
			Handler:    _NotificationService_SendSMSVerification_Handler,
		},
		{
			MethodName: "UpdateFirebaseToken",
			Handler:    _NotificationService_UpdateFirebaseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf-spec/notification_notification_v1.proto",
}
