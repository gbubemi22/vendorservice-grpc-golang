// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/vendor.proto

package vendor

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

// Messages
type VendorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorId string `protobuf:"bytes,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
}

func (x *VendorRequest) Reset() {
	*x = VendorRequest{}
	mi := &file_proto_vendor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorRequest) ProtoMessage() {}

func (x *VendorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vendor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorRequest.ProtoReflect.Descriptor instead.
func (*VendorRequest) Descriptor() ([]byte, []int) {
	return file_proto_vendor_proto_rawDescGZIP(), []int{0}
}

func (x *VendorRequest) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

type VendorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VendorResponse) Reset() {
	*x = VendorResponse{}
	mi := &file_proto_vendor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorResponse) ProtoMessage() {}

func (x *VendorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vendor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorResponse.ProtoReflect.Descriptor instead.
func (*VendorResponse) Descriptor() ([]byte, []int) {
	return file_proto_vendor_proto_rawDescGZIP(), []int{1}
}

func (x *VendorResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VendorResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateVendorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateVendorRequest) Reset() {
	*x = CreateVendorRequest{}
	mi := &file_proto_vendor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVendorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVendorRequest) ProtoMessage() {}

func (x *CreateVendorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vendor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVendorRequest.ProtoReflect.Descriptor instead.
func (*CreateVendorRequest) Descriptor() ([]byte, []int) {
	return file_proto_vendor_proto_rawDescGZIP(), []int{2}
}

func (x *CreateVendorRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateVendorRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateVendorRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateVendorRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateVendorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateVendorResponse) Reset() {
	*x = CreateVendorResponse{}
	mi := &file_proto_vendor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVendorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVendorResponse) ProtoMessage() {}

func (x *CreateVendorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vendor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVendorResponse.ProtoReflect.Descriptor instead.
func (*CreateVendorResponse) Descriptor() ([]byte, []int) {
	return file_proto_vendor_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVendorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_vendor_proto protoreflect.FileDescriptor

var file_proto_vendor_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x0d,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x96, 0x01, 0x0a,
	0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vendor_proto_rawDescOnce sync.Once
	file_proto_vendor_proto_rawDescData = file_proto_vendor_proto_rawDesc
)

func file_proto_vendor_proto_rawDescGZIP() []byte {
	file_proto_vendor_proto_rawDescOnce.Do(func() {
		file_proto_vendor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vendor_proto_rawDescData)
	})
	return file_proto_vendor_proto_rawDescData
}

var file_proto_vendor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_vendor_proto_goTypes = []any{
	(*VendorRequest)(nil),        // 0: vendor.VendorRequest
	(*VendorResponse)(nil),       // 1: vendor.VendorResponse
	(*CreateVendorRequest)(nil),  // 2: vendor.CreateVendorRequest
	(*CreateVendorResponse)(nil), // 3: vendor.CreateVendorResponse
}
var file_proto_vendor_proto_depIdxs = []int32{
	0, // 0: vendor.VendorService.GetVendor:input_type -> vendor.VendorRequest
	2, // 1: vendor.VendorService.CreateVendor:input_type -> vendor.CreateVendorRequest
	1, // 2: vendor.VendorService.GetVendor:output_type -> vendor.VendorResponse
	3, // 3: vendor.VendorService.CreateVendor:output_type -> vendor.CreateVendorResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_vendor_proto_init() }
func file_proto_vendor_proto_init() {
	if File_proto_vendor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_vendor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vendor_proto_goTypes,
		DependencyIndexes: file_proto_vendor_proto_depIdxs,
		MessageInfos:      file_proto_vendor_proto_msgTypes,
	}.Build()
	File_proto_vendor_proto = out.File
	file_proto_vendor_proto_rawDesc = nil
	file_proto_vendor_proto_goTypes = nil
	file_proto_vendor_proto_depIdxs = nil
}
