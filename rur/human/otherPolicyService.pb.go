// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: otherPolicyService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

//残联政策落实情况
type OtherPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PersonId      int64   `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id"`
	RealName      string  `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard        string  `protobuf:"bytes,4,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Date          string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date"`
	Type          string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	SubsidyAmount float32 `protobuf:"fixed32,7,opt,name=subsidy_amount,json=subsidyAmount,proto3" json:"subsidy_amount"`
	Status        string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	CreatedAt     string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Person        *Person `protobuf:"bytes,11,opt,name=person,proto3" json:"person"`
}

func (x *OtherPolicy) Reset() {
	*x = OtherPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherPolicyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherPolicy) ProtoMessage() {}

func (x *OtherPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_otherPolicyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherPolicy.ProtoReflect.Descriptor instead.
func (*OtherPolicy) Descriptor() ([]byte, []int) {
	return file_otherPolicyService_proto_rawDescGZIP(), []int{0}
}

func (x *OtherPolicy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OtherPolicy) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *OtherPolicy) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *OtherPolicy) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *OtherPolicy) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *OtherPolicy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OtherPolicy) GetSubsidyAmount() float32 {
	if x != nil {
		return x.SubsidyAmount
	}
	return 0
}

func (x *OtherPolicy) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OtherPolicy) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OtherPolicy) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *OtherPolicy) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type OtherPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	Keywords         string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Status           string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	Ids              []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids"`
	ComparisonStatus string  `protobuf:"bytes,7,opt,name=comparison_status,json=comparisonStatus,proto3" json:"comparison_status"`
	PersonId         int64   `protobuf:"varint,8,opt,name=person_id,json=personId,proto3" json:"person_id"`
	FamilyId         int64   `protobuf:"varint,9,opt,name=family_id,json=familyId,proto3" json:"family_id"`
}

func (x *OtherPolicyRequest) Reset() {
	*x = OtherPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherPolicyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherPolicyRequest) ProtoMessage() {}

func (x *OtherPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otherPolicyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherPolicyRequest.ProtoReflect.Descriptor instead.
func (*OtherPolicyRequest) Descriptor() ([]byte, []int) {
	return file_otherPolicyService_proto_rawDescGZIP(), []int{1}
}

func (x *OtherPolicyRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *OtherPolicyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OtherPolicyRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *OtherPolicyRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *OtherPolicyRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OtherPolicyRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *OtherPolicyRequest) GetComparisonStatus() string {
	if x != nil {
		return x.ComparisonStatus
	}
	return ""
}

func (x *OtherPolicyRequest) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *OtherPolicyRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

type OtherPolicyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OtherPolicy   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*OtherPolicy `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info   `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *OtherPolicyData) Reset() {
	*x = OtherPolicyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherPolicyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherPolicyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherPolicyData) ProtoMessage() {}

func (x *OtherPolicyData) ProtoReflect() protoreflect.Message {
	mi := &file_otherPolicyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherPolicyData.ProtoReflect.Descriptor instead.
func (*OtherPolicyData) Descriptor() ([]byte, []int) {
	return file_otherPolicyService_proto_rawDescGZIP(), []int{2}
}

func (x *OtherPolicyData) GetEntity() *OtherPolicy {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OtherPolicyData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *OtherPolicyData) GetItems() []*OtherPolicy {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OtherPolicyData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type OtherPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *OtherPolicyData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error    `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *OtherPolicyResponse) Reset() {
	*x = OtherPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherPolicyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherPolicyResponse) ProtoMessage() {}

func (x *OtherPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherPolicyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherPolicyResponse.ProtoReflect.Descriptor instead.
func (*OtherPolicyResponse) Descriptor() ([]byte, []int) {
	return file_otherPolicyService_proto_rawDescGZIP(), []int{3}
}

func (x *OtherPolicyResponse) GetData() *OtherPolicyData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OtherPolicyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_otherPolicyService_proto protoreflect.FileDescriptor

var file_otherPolicyService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a,
	0x0b, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x69,
	0x64, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x69, 0x64, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x8e,
	0x02, 0x0a, 0x12, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22,
	0xb4, 0x01, 0x0a, 0x0f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x69, 0x0a, 0x13, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xab, 0x03, 0x0a, 0x12, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otherPolicyService_proto_rawDescOnce sync.Once
	file_otherPolicyService_proto_rawDescData = file_otherPolicyService_proto_rawDesc
)

func file_otherPolicyService_proto_rawDescGZIP() []byte {
	file_otherPolicyService_proto_rawDescOnce.Do(func() {
		file_otherPolicyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_otherPolicyService_proto_rawDescData)
	})
	return file_otherPolicyService_proto_rawDescData
}

var file_otherPolicyService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_otherPolicyService_proto_goTypes = []interface{}{
	(*OtherPolicy)(nil),         // 0: services.OtherPolicy
	(*OtherPolicyRequest)(nil),  // 1: services.OtherPolicyRequest
	(*OtherPolicyData)(nil),     // 2: services.OtherPolicyData
	(*OtherPolicyResponse)(nil), // 3: services.OtherPolicyResponse
	(*Person)(nil),              // 4: services.Person
	(*common.Pager)(nil),        // 5: common.Pager
	(*common.Info)(nil),         // 6: common.Info
	(*common.Error)(nil),        // 7: common.Error
}
var file_otherPolicyService_proto_depIdxs = []int32{
	4,  // 0: services.OtherPolicy.person:type_name -> services.Person
	0,  // 1: services.OtherPolicyData.entity:type_name -> services.OtherPolicy
	5,  // 2: services.OtherPolicyData.pager:type_name -> common.Pager
	0,  // 3: services.OtherPolicyData.items:type_name -> services.OtherPolicy
	6,  // 4: services.OtherPolicyData.info:type_name -> common.Info
	2,  // 5: services.OtherPolicyResponse.data:type_name -> services.OtherPolicyData
	7,  // 6: services.OtherPolicyResponse.error:type_name -> common.Error
	0,  // 7: services.OtherPolicyService.Create:input_type -> services.OtherPolicy
	0,  // 8: services.OtherPolicyService.Update:input_type -> services.OtherPolicy
	0,  // 9: services.OtherPolicyService.Delete:input_type -> services.OtherPolicy
	0,  // 10: services.OtherPolicyService.Get:input_type -> services.OtherPolicy
	1,  // 11: services.OtherPolicyService.Search:input_type -> services.OtherPolicyRequest
	1,  // 12: services.OtherPolicyService.Export:input_type -> services.OtherPolicyRequest
	3,  // 13: services.OtherPolicyService.Create:output_type -> services.OtherPolicyResponse
	3,  // 14: services.OtherPolicyService.Update:output_type -> services.OtherPolicyResponse
	3,  // 15: services.OtherPolicyService.Delete:output_type -> services.OtherPolicyResponse
	3,  // 16: services.OtherPolicyService.Get:output_type -> services.OtherPolicyResponse
	3,  // 17: services.OtherPolicyService.Search:output_type -> services.OtherPolicyResponse
	3,  // 18: services.OtherPolicyService.Export:output_type -> services.OtherPolicyResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_otherPolicyService_proto_init() }
func file_otherPolicyService_proto_init() {
	if File_otherPolicyService_proto != nil {
		return
	}
	file_personService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_otherPolicyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherPolicy); i {
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
		file_otherPolicyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherPolicyRequest); i {
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
		file_otherPolicyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherPolicyData); i {
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
		file_otherPolicyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherPolicyResponse); i {
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
			RawDescriptor: file_otherPolicyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_otherPolicyService_proto_goTypes,
		DependencyIndexes: file_otherPolicyService_proto_depIdxs,
		MessageInfos:      file_otherPolicyService_proto_msgTypes,
	}.Build()
	File_otherPolicyService_proto = out.File
	file_otherPolicyService_proto_rawDesc = nil
	file_otherPolicyService_proto_goTypes = nil
	file_otherPolicyService_proto_depIdxs = nil
}
