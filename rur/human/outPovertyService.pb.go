// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: outPovertyService.proto

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

//脱贫监测对象
type OutPoverty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	FamilyId          int64   `protobuf:"varint,2,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	RealName          string  `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard            string  `protobuf:"bytes,4,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	FamilyType        string  `protobuf:"bytes,5,opt,name=family_type,json=familyType,proto3" json:"family_type"`
	MonitoringType    string  `protobuf:"bytes,6,opt,name=monitoring_type,json=monitoringType,proto3" json:"monitoring_type"`
	RiskIsEliminate   string  `protobuf:"bytes,7,opt,name=risk_is_eliminate,json=riskIsEliminate,proto3" json:"risk_is_eliminate"`
	RiskIdentifyTime  string  `protobuf:"bytes,8,opt,name=risk_identify_time,json=riskIdentifyTime,proto3" json:"risk_identify_time"`
	RiskEliminateTime string  `protobuf:"bytes,9,opt,name=risk_eliminate_time,json=riskEliminateTime,proto3" json:"risk_eliminate_time"`
	Status            string  `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	CreatedAt         string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Family            *Family `protobuf:"bytes,13,opt,name=family,proto3" json:"family"`
	Ids               []int64 `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *OutPoverty) Reset() {
	*x = OutPoverty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outPovertyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPoverty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPoverty) ProtoMessage() {}

func (x *OutPoverty) ProtoReflect() protoreflect.Message {
	mi := &file_outPovertyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPoverty.ProtoReflect.Descriptor instead.
func (*OutPoverty) Descriptor() ([]byte, []int) {
	return file_outPovertyService_proto_rawDescGZIP(), []int{0}
}

func (x *OutPoverty) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OutPoverty) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *OutPoverty) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *OutPoverty) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *OutPoverty) GetFamilyType() string {
	if x != nil {
		return x.FamilyType
	}
	return ""
}

func (x *OutPoverty) GetMonitoringType() string {
	if x != nil {
		return x.MonitoringType
	}
	return ""
}

func (x *OutPoverty) GetRiskIsEliminate() string {
	if x != nil {
		return x.RiskIsEliminate
	}
	return ""
}

func (x *OutPoverty) GetRiskIdentifyTime() string {
	if x != nil {
		return x.RiskIdentifyTime
	}
	return ""
}

func (x *OutPoverty) GetRiskEliminateTime() string {
	if x != nil {
		return x.RiskEliminateTime
	}
	return ""
}

func (x *OutPoverty) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OutPoverty) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OutPoverty) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *OutPoverty) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *OutPoverty) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type OutPovertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	FamilyId         int64  `protobuf:"varint,4,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	PersonId         int64  `protobuf:"varint,5,opt,name=person_id,json=personId,proto3" json:"person_id"`
	MasterId         int64  `protobuf:"varint,6,opt,name=master_id,json=masterId,proto3" json:"master_id"`
	RealName         string `protobuf:"bytes,7,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard           string `protobuf:"bytes,8,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Status           string `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	ComparisonStatus string `protobuf:"bytes,10,opt,name=comparison_status,json=comparisonStatus,proto3" json:"comparison_status"`
}

func (x *OutPovertyRequest) Reset() {
	*x = OutPovertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outPovertyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPovertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPovertyRequest) ProtoMessage() {}

func (x *OutPovertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outPovertyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPovertyRequest.ProtoReflect.Descriptor instead.
func (*OutPovertyRequest) Descriptor() ([]byte, []int) {
	return file_outPovertyService_proto_rawDescGZIP(), []int{1}
}

func (x *OutPovertyRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *OutPovertyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OutPovertyRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *OutPovertyRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *OutPovertyRequest) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *OutPovertyRequest) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *OutPovertyRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *OutPovertyRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *OutPovertyRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OutPovertyRequest) GetComparisonStatus() string {
	if x != nil {
		return x.ComparisonStatus
	}
	return ""
}

type OutPovertyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OutPoverty   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*OutPoverty `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *OutPovertyData) Reset() {
	*x = OutPovertyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outPovertyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPovertyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPovertyData) ProtoMessage() {}

func (x *OutPovertyData) ProtoReflect() protoreflect.Message {
	mi := &file_outPovertyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPovertyData.ProtoReflect.Descriptor instead.
func (*OutPovertyData) Descriptor() ([]byte, []int) {
	return file_outPovertyService_proto_rawDescGZIP(), []int{2}
}

func (x *OutPovertyData) GetEntity() *OutPoverty {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OutPovertyData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *OutPovertyData) GetItems() []*OutPoverty {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OutPovertyData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type OutPovertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *OutPovertyData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *OutPovertyResponse) Reset() {
	*x = OutPovertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outPovertyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPovertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPovertyResponse) ProtoMessage() {}

func (x *OutPovertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outPovertyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPovertyResponse.ProtoReflect.Descriptor instead.
func (*OutPovertyResponse) Descriptor() ([]byte, []int) {
	return file_outPovertyService_proto_rawDescGZIP(), []int{3}
}

func (x *OutPovertyResponse) GetData() *OutPovertyData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OutPovertyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_outPovertyService_proto protoreflect.FileDescriptor

var file_outPovertyService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x03, 0x0a, 0x0a,
	0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x69, 0x73, 0x6b, 0x5f,
	0x69, 0x73, 0x5f, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x69, 0x73, 0x6b, 0x49, 0x73, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x72, 0x69, 0x73, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x72, 0x69, 0x73, 0x6b, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0xb2, 0x02, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x4f, 0x75, 0x74,
	0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65,
	0x72, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a, 0x12,
	0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50,
	0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9e, 0x03, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76,
	0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75,
	0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75,
	0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f,
	0x76, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outPovertyService_proto_rawDescOnce sync.Once
	file_outPovertyService_proto_rawDescData = file_outPovertyService_proto_rawDesc
)

func file_outPovertyService_proto_rawDescGZIP() []byte {
	file_outPovertyService_proto_rawDescOnce.Do(func() {
		file_outPovertyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_outPovertyService_proto_rawDescData)
	})
	return file_outPovertyService_proto_rawDescData
}

var file_outPovertyService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_outPovertyService_proto_goTypes = []interface{}{
	(*OutPoverty)(nil),         // 0: services.OutPoverty
	(*OutPovertyRequest)(nil),  // 1: services.OutPovertyRequest
	(*OutPovertyData)(nil),     // 2: services.OutPovertyData
	(*OutPovertyResponse)(nil), // 3: services.OutPovertyResponse
	(*Family)(nil),             // 4: services.Family
	(*common.Pager)(nil),       // 5: common.Pager
	(*common.Info)(nil),        // 6: common.Info
	(*common.Error)(nil),       // 7: common.Error
}
var file_outPovertyService_proto_depIdxs = []int32{
	4,  // 0: services.OutPoverty.family:type_name -> services.Family
	0,  // 1: services.OutPovertyData.entity:type_name -> services.OutPoverty
	5,  // 2: services.OutPovertyData.pager:type_name -> common.Pager
	0,  // 3: services.OutPovertyData.items:type_name -> services.OutPoverty
	6,  // 4: services.OutPovertyData.info:type_name -> common.Info
	2,  // 5: services.OutPovertyResponse.data:type_name -> services.OutPovertyData
	7,  // 6: services.OutPovertyResponse.error:type_name -> common.Error
	0,  // 7: services.OutPovertyService.Create:input_type -> services.OutPoverty
	0,  // 8: services.OutPovertyService.Update:input_type -> services.OutPoverty
	0,  // 9: services.OutPovertyService.Delete:input_type -> services.OutPoverty
	0,  // 10: services.OutPovertyService.Get:input_type -> services.OutPoverty
	1,  // 11: services.OutPovertyService.Search:input_type -> services.OutPovertyRequest
	1,  // 12: services.OutPovertyService.Export:input_type -> services.OutPovertyRequest
	3,  // 13: services.OutPovertyService.Create:output_type -> services.OutPovertyResponse
	3,  // 14: services.OutPovertyService.Update:output_type -> services.OutPovertyResponse
	3,  // 15: services.OutPovertyService.Delete:output_type -> services.OutPovertyResponse
	3,  // 16: services.OutPovertyService.Get:output_type -> services.OutPovertyResponse
	3,  // 17: services.OutPovertyService.Search:output_type -> services.OutPovertyResponse
	3,  // 18: services.OutPovertyService.Export:output_type -> services.OutPovertyResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_outPovertyService_proto_init() }
func file_outPovertyService_proto_init() {
	if File_outPovertyService_proto != nil {
		return
	}
	file_familyService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_outPovertyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPoverty); i {
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
		file_outPovertyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPovertyRequest); i {
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
		file_outPovertyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPovertyData); i {
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
		file_outPovertyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPovertyResponse); i {
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
			RawDescriptor: file_outPovertyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outPovertyService_proto_goTypes,
		DependencyIndexes: file_outPovertyService_proto_depIdxs,
		MessageInfos:      file_outPovertyService_proto_msgTypes,
	}.Build()
	File_outPovertyService_proto = out.File
	file_outPovertyService_proto_rawDesc = nil
	file_outPovertyService_proto_goTypes = nil
	file_outPovertyService_proto_depIdxs = nil
}
