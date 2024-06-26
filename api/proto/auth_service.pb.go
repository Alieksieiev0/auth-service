// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: api/proto/auth_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_api_proto_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UserClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Algorithm string                 `protobuf:"bytes,4,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Issuer    string                 `protobuf:"bytes,5,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject   string                 `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Audience  []string               `protobuf:"bytes,7,rep,name=audience,proto3" json:"audience,omitempty"`
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	NotBefore *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	IssuedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Id        string                 `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserClaims) Reset() {
	*x = UserClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserClaims) ProtoMessage() {}

func (x *UserClaims) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserClaims.ProtoReflect.Descriptor instead.
func (*UserClaims) Descriptor() ([]byte, []int) {
	return file_api_proto_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserClaims) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserClaims) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserClaims) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserClaims) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *UserClaims) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *UserClaims) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *UserClaims) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *UserClaims) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *UserClaims) GetNotBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *UserClaims) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *UserClaims) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_proto_auth_service_proto protoreflect.FileDescriptor

var file_api_proto_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x82,
	0x03, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x30, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x12, 0x06, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x69, 0x65, 0x6b, 0x73, 0x69, 0x65, 0x69, 0x65, 0x76, 0x30,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_auth_service_proto_rawDescOnce sync.Once
	file_api_proto_auth_service_proto_rawDescData = file_api_proto_auth_service_proto_rawDesc
)

func file_api_proto_auth_service_proto_rawDescGZIP() []byte {
	file_api_proto_auth_service_proto_rawDescOnce.Do(func() {
		file_api_proto_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_auth_service_proto_rawDescData)
	})
	return file_api_proto_auth_service_proto_rawDescData
}

var file_api_proto_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_auth_service_proto_goTypes = []interface{}{
	(*Token)(nil),                 // 0: Token
	(*UserClaims)(nil),            // 1: UserClaims
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_proto_auth_service_proto_depIdxs = []int32{
	2, // 0: UserClaims.expires_at:type_name -> google.protobuf.Timestamp
	2, // 1: UserClaims.not_before:type_name -> google.protobuf.Timestamp
	2, // 2: UserClaims.issued_at:type_name -> google.protobuf.Timestamp
	0, // 3: AuthService.ReadClaims:input_type -> Token
	1, // 4: AuthService.ReadClaims:output_type -> UserClaims
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_auth_service_proto_init() }
func file_api_proto_auth_service_proto_init() {
	if File_api_proto_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_api_proto_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserClaims); i {
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
			RawDescriptor: file_api_proto_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_auth_service_proto_goTypes,
		DependencyIndexes: file_api_proto_auth_service_proto_depIdxs,
		MessageInfos:      file_api_proto_auth_service_proto_msgTypes,
	}.Build()
	File_api_proto_auth_service_proto = out.File
	file_api_proto_auth_service_proto_rawDesc = nil
	file_api_proto_auth_service_proto_goTypes = nil
	file_api_proto_auth_service_proto_depIdxs = nil
}
