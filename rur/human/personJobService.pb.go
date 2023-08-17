// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: personJobService.proto

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

//就业监测
type PersonJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PersonId     int64   `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id"`
	RealName     string  `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard       string  `protobuf:"bytes,4,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Date         string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date"`
	Place        string  `protobuf:"bytes,6,opt,name=place,proto3" json:"place"`
	WorkType     string  `protobuf:"bytes,7,opt,name=work_type,json=workType,proto3" json:"work_type"`
	IsPublicPost string  `protobuf:"bytes,8,opt,name=is_public_post,json=isPublicPost,proto3" json:"is_public_post"`
	MonthSalary  float32 `protobuf:"fixed32,9,opt,name=month_salary,json=monthSalary,proto3" json:"month_salary"`
	Status       string  `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	CreatedAt    string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Person       *Person `protobuf:"bytes,13,opt,name=person,proto3" json:"person"`
}

func (x *PersonJob) Reset() {
	*x = PersonJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personJobService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonJob) ProtoMessage() {}

func (x *PersonJob) ProtoReflect() protoreflect.Message {
	mi := &file_personJobService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonJob.ProtoReflect.Descriptor instead.
func (*PersonJob) Descriptor() ([]byte, []int) {
	return file_personJobService_proto_rawDescGZIP(), []int{0}
}

func (x *PersonJob) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonJob) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *PersonJob) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PersonJob) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PersonJob) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PersonJob) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *PersonJob) GetWorkType() string {
	if x != nil {
		return x.WorkType
	}
	return ""
}

func (x *PersonJob) GetIsPublicPost() string {
	if x != nil {
		return x.IsPublicPost
	}
	return ""
}

func (x *PersonJob) GetMonthSalary() float32 {
	if x != nil {
		return x.MonthSalary
	}
	return 0
}

func (x *PersonJob) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PersonJob) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PersonJob) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PersonJob) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type PersonJobRequest struct {
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

func (x *PersonJobRequest) Reset() {
	*x = PersonJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personJobService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonJobRequest) ProtoMessage() {}

func (x *PersonJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personJobService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonJobRequest.ProtoReflect.Descriptor instead.
func (*PersonJobRequest) Descriptor() ([]byte, []int) {
	return file_personJobService_proto_rawDescGZIP(), []int{1}
}

func (x *PersonJobRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PersonJobRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PersonJobRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PersonJobRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PersonJobRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PersonJobRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PersonJobRequest) GetComparisonStatus() string {
	if x != nil {
		return x.ComparisonStatus
	}
	return ""
}

func (x *PersonJobRequest) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *PersonJobRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

type PersonJobData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PersonJob    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PersonJob  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *PersonJobData) Reset() {
	*x = PersonJobData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personJobService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonJobData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonJobData) ProtoMessage() {}

func (x *PersonJobData) ProtoReflect() protoreflect.Message {
	mi := &file_personJobService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonJobData.ProtoReflect.Descriptor instead.
func (*PersonJobData) Descriptor() ([]byte, []int) {
	return file_personJobService_proto_rawDescGZIP(), []int{2}
}

func (x *PersonJobData) GetEntity() *PersonJob {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PersonJobData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PersonJobData) GetItems() []*PersonJob {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PersonJobData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PersonJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PersonJobData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PersonJobResponse) Reset() {
	*x = PersonJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personJobService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonJobResponse) ProtoMessage() {}

func (x *PersonJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personJobService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonJobResponse.ProtoReflect.Descriptor instead.
func (*PersonJobResponse) Descriptor() ([]byte, []int) {
	return file_personJobService_proto_rawDescGZIP(), []int{3}
}

func (x *PersonJobResponse) GetData() *PersonJobData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PersonJobResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_personJobService_proto protoreflect.FileDescriptor

var file_personJobService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x09, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x8c, 0x02, 0x0a, 0x10,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f,
	0x62, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x11, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x91, 0x03, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_personJobService_proto_rawDescOnce sync.Once
	file_personJobService_proto_rawDescData = file_personJobService_proto_rawDesc
)

func file_personJobService_proto_rawDescGZIP() []byte {
	file_personJobService_proto_rawDescOnce.Do(func() {
		file_personJobService_proto_rawDescData = protoimpl.X.CompressGZIP(file_personJobService_proto_rawDescData)
	})
	return file_personJobService_proto_rawDescData
}

var file_personJobService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_personJobService_proto_goTypes = []interface{}{
	(*PersonJob)(nil),         // 0: services.PersonJob
	(*PersonJobRequest)(nil),  // 1: services.PersonJobRequest
	(*PersonJobData)(nil),     // 2: services.PersonJobData
	(*PersonJobResponse)(nil), // 3: services.PersonJobResponse
	(*Person)(nil),            // 4: services.Person
	(*common.Pager)(nil),      // 5: common.Pager
	(*common.Info)(nil),       // 6: common.Info
	(*common.Error)(nil),      // 7: common.Error
}
var file_personJobService_proto_depIdxs = []int32{
	4,  // 0: services.PersonJob.person:type_name -> services.Person
	0,  // 1: services.PersonJobData.entity:type_name -> services.PersonJob
	5,  // 2: services.PersonJobData.pager:type_name -> common.Pager
	0,  // 3: services.PersonJobData.items:type_name -> services.PersonJob
	6,  // 4: services.PersonJobData.info:type_name -> common.Info
	2,  // 5: services.PersonJobResponse.data:type_name -> services.PersonJobData
	7,  // 6: services.PersonJobResponse.error:type_name -> common.Error
	0,  // 7: services.PersonJobService.Create:input_type -> services.PersonJob
	0,  // 8: services.PersonJobService.Update:input_type -> services.PersonJob
	0,  // 9: services.PersonJobService.Delete:input_type -> services.PersonJob
	0,  // 10: services.PersonJobService.Get:input_type -> services.PersonJob
	1,  // 11: services.PersonJobService.Search:input_type -> services.PersonJobRequest
	1,  // 12: services.PersonJobService.Export:input_type -> services.PersonJobRequest
	3,  // 13: services.PersonJobService.Create:output_type -> services.PersonJobResponse
	3,  // 14: services.PersonJobService.Update:output_type -> services.PersonJobResponse
	3,  // 15: services.PersonJobService.Delete:output_type -> services.PersonJobResponse
	3,  // 16: services.PersonJobService.Get:output_type -> services.PersonJobResponse
	3,  // 17: services.PersonJobService.Search:output_type -> services.PersonJobResponse
	3,  // 18: services.PersonJobService.Export:output_type -> services.PersonJobResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_personJobService_proto_init() }
func file_personJobService_proto_init() {
	if File_personJobService_proto != nil {
		return
	}
	file_personService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_personJobService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonJob); i {
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
		file_personJobService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonJobRequest); i {
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
		file_personJobService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonJobData); i {
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
		file_personJobService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonJobResponse); i {
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
			RawDescriptor: file_personJobService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personJobService_proto_goTypes,
		DependencyIndexes: file_personJobService_proto_depIdxs,
		MessageInfos:      file_personJobService_proto_msgTypes,
	}.Build()
	File_personJobService_proto = out.File
	file_personJobService_proto_rawDesc = nil
	file_personJobService_proto_goTypes = nil
	file_personJobService_proto_depIdxs = nil
}
