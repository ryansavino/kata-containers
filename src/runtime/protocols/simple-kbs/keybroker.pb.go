// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: keybroker.proto

package simple_kbs

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

type BundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// for platform verification
	// the cert chain includes the PDH
	CertificateChain string `protobuf:"bytes,1,opt,name=CertificateChain,proto3" json:"CertificateChain,omitempty"`
	// Required to construct launch blob
	Policy uint32 `protobuf:"varint,2,opt,name=Policy,proto3" json:"Policy,omitempty"`
}

func (x *BundleRequest) Reset() {
	*x = BundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keybroker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleRequest) ProtoMessage() {}

func (x *BundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keybroker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleRequest.ProtoReflect.Descriptor instead.
func (*BundleRequest) Descriptor() ([]byte, []int) {
	return file_keybroker_proto_rawDescGZIP(), []int{0}
}

func (x *BundleRequest) GetCertificateChain() string {
	if x != nil {
		return x.CertificateChain
	}
	return ""
}

func (x *BundleRequest) GetPolicy() uint32 {
	if x != nil {
		return x.Policy
	}
	return 0
}

type BundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestOwnerPublicKey string `protobuf:"bytes,1,opt,name=GuestOwnerPublicKey,proto3" json:"GuestOwnerPublicKey,omitempty"`
	LaunchBlob          string `protobuf:"bytes,2,opt,name=LaunchBlob,proto3" json:"LaunchBlob,omitempty"`
	// GUID
	LaunchId string `protobuf:"bytes,3,opt,name=LaunchId,proto3" json:"LaunchId,omitempty"`
}

func (x *BundleResponse) Reset() {
	*x = BundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keybroker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleResponse) ProtoMessage() {}

func (x *BundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keybroker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleResponse.ProtoReflect.Descriptor instead.
func (*BundleResponse) Descriptor() ([]byte, []int) {
	return file_keybroker_proto_rawDescGZIP(), []int{1}
}

func (x *BundleResponse) GetGuestOwnerPublicKey() string {
	if x != nil {
		return x.GuestOwnerPublicKey
	}
	return ""
}

func (x *BundleResponse) GetLaunchBlob() string {
	if x != nil {
		return x.LaunchBlob
	}
	return ""
}

func (x *BundleResponse) GetLaunchId() string {
	if x != nil {
		return x.LaunchId
	}
	return ""
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid       string `protobuf:"bytes,1,opt,name=Guid,proto3" json:"Guid,omitempty"`
	Format     string `protobuf:"bytes,2,opt,name=Format,proto3" json:"Format,omitempty"`
	SecretType string `protobuf:"bytes,3,opt,name=SecretType,proto3" json:"SecretType,omitempty"`
	Id         string `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keybroker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_keybroker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDetails.ProtoReflect.Descriptor instead.
func (*RequestDetails) Descriptor() ([]byte, []int) {
	return file_keybroker_proto_rawDescGZIP(), []int{2}
}

func (x *RequestDetails) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *RequestDetails) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *RequestDetails) GetSecretType() string {
	if x != nil {
		return x.SecretType
	}
	return ""
}

func (x *RequestDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// b64 encoded string
	LaunchMeasurement string `protobuf:"bytes,1,opt,name=LaunchMeasurement,proto3" json:"LaunchMeasurement,omitempty"`
	LaunchId          string `protobuf:"bytes,2,opt,name=LaunchId,proto3" json:"LaunchId,omitempty"`
	Policy            uint32 `protobuf:"varint,3,opt,name=Policy,proto3" json:"Policy,omitempty"`
	// Hints for secret validation,
	ApiMajor uint32 `protobuf:"varint,4,opt,name=ApiMajor,proto3" json:"ApiMajor,omitempty"`
	ApiMinor uint32 `protobuf:"varint,5,opt,name=ApiMinor,proto3" json:"ApiMinor,omitempty"`
	BuildId  uint32 `protobuf:"varint,6,opt,name=BuildId,proto3" json:"BuildId,omitempty"`
	// The fw digest that the guest was launched with. Hopefully we can
	// get this from QEMU. Pass this in as a base64 string.
	FwDigest string `protobuf:"bytes,7,opt,name=FwDigest,proto3" json:"FwDigest,omitempty"`
	// Flexible description of launch provided by orchestrator
	// and logged by KBS.
	LaunchDescription string            `protobuf:"bytes,8,opt,name=LaunchDescription,proto3" json:"LaunchDescription,omitempty"`
	SecretRequests    []*RequestDetails `protobuf:"bytes,9,rep,name=SecretRequests,proto3" json:"SecretRequests,omitempty"`
}

func (x *SecretRequest) Reset() {
	*x = SecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keybroker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretRequest) ProtoMessage() {}

func (x *SecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keybroker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretRequest.ProtoReflect.Descriptor instead.
func (*SecretRequest) Descriptor() ([]byte, []int) {
	return file_keybroker_proto_rawDescGZIP(), []int{3}
}

func (x *SecretRequest) GetLaunchMeasurement() string {
	if x != nil {
		return x.LaunchMeasurement
	}
	return ""
}

func (x *SecretRequest) GetLaunchId() string {
	if x != nil {
		return x.LaunchId
	}
	return ""
}

func (x *SecretRequest) GetPolicy() uint32 {
	if x != nil {
		return x.Policy
	}
	return 0
}

func (x *SecretRequest) GetApiMajor() uint32 {
	if x != nil {
		return x.ApiMajor
	}
	return 0
}

func (x *SecretRequest) GetApiMinor() uint32 {
	if x != nil {
		return x.ApiMinor
	}
	return 0
}

func (x *SecretRequest) GetBuildId() uint32 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *SecretRequest) GetFwDigest() string {
	if x != nil {
		return x.FwDigest
	}
	return ""
}

func (x *SecretRequest) GetLaunchDescription() string {
	if x != nil {
		return x.LaunchDescription
	}
	return ""
}

func (x *SecretRequest) GetSecretRequests() []*RequestDetails {
	if x != nil {
		return x.SecretRequests
	}
	return nil
}

type SecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaunchSecretHeader string `protobuf:"bytes,1,opt,name=LaunchSecretHeader,proto3" json:"LaunchSecretHeader,omitempty"`
	LaunchSecretData   string `protobuf:"bytes,2,opt,name=LaunchSecretData,proto3" json:"LaunchSecretData,omitempty"`
}

func (x *SecretResponse) Reset() {
	*x = SecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keybroker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretResponse) ProtoMessage() {}

func (x *SecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keybroker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretResponse.ProtoReflect.Descriptor instead.
func (*SecretResponse) Descriptor() ([]byte, []int) {
	return file_keybroker_proto_rawDescGZIP(), []int{4}
}

func (x *SecretResponse) GetLaunchSecretHeader() string {
	if x != nil {
		return x.LaunchSecretHeader
	}
	return ""
}

func (x *SecretResponse) GetLaunchSecretData() string {
	if x != nil {
		return x.LaunchSecretData
	}
	return ""
}

var File_keybroker_proto protoreflect.FileDescriptor

var file_keybroker_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x0d,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x22, 0x7e, 0x0a, 0x0e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x42,
	0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x47, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22,
	0xd0, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x41, 0x70, 0x69, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x41, 0x70, 0x69, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x77, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x77, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x6c, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x32, 0x9a, 0x01, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6b,
	0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6b, 0x65, 0x79, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6b, 0x62, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keybroker_proto_rawDescOnce sync.Once
	file_keybroker_proto_rawDescData = file_keybroker_proto_rawDesc
)

func file_keybroker_proto_rawDescGZIP() []byte {
	file_keybroker_proto_rawDescOnce.Do(func() {
		file_keybroker_proto_rawDescData = protoimpl.X.CompressGZIP(file_keybroker_proto_rawDescData)
	})
	return file_keybroker_proto_rawDescData
}

var file_keybroker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_keybroker_proto_goTypes = []interface{}{
	(*BundleRequest)(nil),  // 0: keybroker.BundleRequest
	(*BundleResponse)(nil), // 1: keybroker.BundleResponse
	(*RequestDetails)(nil), // 2: keybroker.RequestDetails
	(*SecretRequest)(nil),  // 3: keybroker.SecretRequest
	(*SecretResponse)(nil), // 4: keybroker.SecretResponse
}
var file_keybroker_proto_depIdxs = []int32{
	2, // 0: keybroker.SecretRequest.SecretRequests:type_name -> keybroker.RequestDetails
	0, // 1: keybroker.KeyBrokerService.GetBundle:input_type -> keybroker.BundleRequest
	3, // 2: keybroker.KeyBrokerService.GetSecret:input_type -> keybroker.SecretRequest
	1, // 3: keybroker.KeyBrokerService.GetBundle:output_type -> keybroker.BundleResponse
	4, // 4: keybroker.KeyBrokerService.GetSecret:output_type -> keybroker.SecretResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keybroker_proto_init() }
func file_keybroker_proto_init() {
	if File_keybroker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keybroker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleRequest); i {
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
		file_keybroker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleResponse); i {
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
		file_keybroker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDetails); i {
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
		file_keybroker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretRequest); i {
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
		file_keybroker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretResponse); i {
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
			RawDescriptor: file_keybroker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keybroker_proto_goTypes,
		DependencyIndexes: file_keybroker_proto_depIdxs,
		MessageInfos:      file_keybroker_proto_msgTypes,
	}.Build()
	File_keybroker_proto = out.File
	file_keybroker_proto_rawDesc = nil
	file_keybroker_proto_goTypes = nil
	file_keybroker_proto_depIdxs = nil
}