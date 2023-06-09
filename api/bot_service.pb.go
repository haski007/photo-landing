// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: bot_service.proto

package __

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

// The message containing user's name, email, and the contact message.
type ContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ContactRequest) Reset() {
	*x = ContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRequest) ProtoMessage() {}

func (x *ContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRequest.ProtoReflect.Descriptor instead.
func (*ContactRequest) Descriptor() ([]byte, []int) {
	return file_bot_service_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The response message containing the result of sending the contact request.
type ContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ContactResponse) Reset() {
	*x = ContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactResponse) ProtoMessage() {}

func (x *ContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactResponse.ProtoReflect.Descriptor instead.
func (*ContactResponse) Descriptor() ([]byte, []int) {
	return file_bot_service_proto_rawDescGZIP(), []int{1}
}

func (x *ContactResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ContactResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_bot_service_proto protoreflect.FileDescriptor

var file_bot_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x54, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0x4d, 0x0a, 0x0a, 0x42, 0x6f, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_service_proto_rawDescOnce sync.Once
	file_bot_service_proto_rawDescData = file_bot_service_proto_rawDesc
)

func file_bot_service_proto_rawDescGZIP() []byte {
	file_bot_service_proto_rawDescOnce.Do(func() {
		file_bot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_service_proto_rawDescData)
	})
	return file_bot_service_proto_rawDescData
}

var file_bot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bot_service_proto_goTypes = []interface{}{
	(*ContactRequest)(nil),  // 0: api.ContactRequest
	(*ContactResponse)(nil), // 1: api.ContactResponse
}
var file_bot_service_proto_depIdxs = []int32{
	0, // 0: api.BotService.SendContactRequest:input_type -> api.ContactRequest
	1, // 1: api.BotService.SendContactRequest:output_type -> api.ContactResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bot_service_proto_init() }
func file_bot_service_proto_init() {
	if File_bot_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bot_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRequest); i {
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
		file_bot_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactResponse); i {
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
			RawDescriptor: file_bot_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bot_service_proto_goTypes,
		DependencyIndexes: file_bot_service_proto_depIdxs,
		MessageInfos:      file_bot_service_proto_msgTypes,
	}.Build()
	File_bot_service_proto = out.File
	file_bot_service_proto_rawDesc = nil
	file_bot_service_proto_goTypes = nil
	file_bot_service_proto_depIdxs = nil
}
