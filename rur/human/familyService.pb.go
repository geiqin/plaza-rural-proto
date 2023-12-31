// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: familyService.proto

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

//家庭信息
type Family struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Code         string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	HouseNumber  string    `protobuf:"bytes,3,opt,name=house_number,json=houseNumber,proto3" json:"house_number"`
	MasterId     int64     `protobuf:"varint,4,opt,name=master_id,json=masterId,proto3" json:"master_id"`
	MasterName   string    `protobuf:"bytes,5,opt,name=master_name,json=masterName,proto3" json:"master_name"`
	MasterIdCard string    `protobuf:"bytes,6,opt,name=master_id_card,json=masterIdCard,proto3" json:"master_id_card"`
	Type         string    `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
	CoverId      int64     `protobuf:"varint,8,opt,name=cover_id,json=coverId,proto3" json:"cover_id"`
	CoverUrl     string    `protobuf:"bytes,9,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
	RegionId     int64     `protobuf:"varint,10,opt,name=region_id,json=regionId,proto3" json:"region_id"`
	GroupId      int64     `protobuf:"varint,11,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Addr         string    `protobuf:"bytes,12,opt,name=addr,proto3" json:"addr"`
	Lng          string    `protobuf:"bytes,13,opt,name=lng,proto3" json:"lng"`
	Lat          string    `protobuf:"bytes,14,opt,name=lat,proto3" json:"lat"`
	Altitude     int32     `protobuf:"varint,15,opt,name=altitude,proto3" json:"altitude"`
	MemberCount  int64     `protobuf:"varint,16,opt,name=member_count,json=memberCount,proto3" json:"member_count"`
	IsPoor       string    `protobuf:"bytes,17,opt,name=is_poor,json=isPoor,proto3" json:"is_poor"`
	Habitation   string    `protobuf:"bytes,18,opt,name=habitation,proto3" json:"habitation"`
	Status       string    `protobuf:"bytes,19,opt,name=status,proto3" json:"status"`
	CreatedAt    string    `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string    `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Master       *Person   `protobuf:"bytes,22,opt,name=master,proto3" json:"master"`
	Region       *Region   `protobuf:"bytes,23,opt,name=region,proto3" json:"region"`
	Group        *Group    `protobuf:"bytes,24,opt,name=group,proto3" json:"group"`
	Membership   []*Person `protobuf:"bytes,25,rep,name=membership,proto3" json:"membership"`
}

func (x *Family) Reset() {
	*x = Family{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Family) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Family) ProtoMessage() {}

func (x *Family) ProtoReflect() protoreflect.Message {
	mi := &file_familyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Family.ProtoReflect.Descriptor instead.
func (*Family) Descriptor() ([]byte, []int) {
	return file_familyService_proto_rawDescGZIP(), []int{0}
}

func (x *Family) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Family) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Family) GetHouseNumber() string {
	if x != nil {
		return x.HouseNumber
	}
	return ""
}

func (x *Family) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *Family) GetMasterName() string {
	if x != nil {
		return x.MasterName
	}
	return ""
}

func (x *Family) GetMasterIdCard() string {
	if x != nil {
		return x.MasterIdCard
	}
	return ""
}

func (x *Family) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Family) GetCoverId() int64 {
	if x != nil {
		return x.CoverId
	}
	return 0
}

func (x *Family) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Family) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *Family) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Family) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Family) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Family) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Family) GetAltitude() int32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Family) GetMemberCount() int64 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *Family) GetIsPoor() string {
	if x != nil {
		return x.IsPoor
	}
	return ""
}

func (x *Family) GetHabitation() string {
	if x != nil {
		return x.Habitation
	}
	return ""
}

func (x *Family) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Family) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Family) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Family) GetMaster() *Person {
	if x != nil {
		return x.Master
	}
	return nil
}

func (x *Family) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *Family) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *Family) GetMembership() []*Person {
	if x != nil {
		return x.Membership
	}
	return nil
}

type FamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	FamilyId    int64   `protobuf:"varint,4,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	PersonId    int64   `protobuf:"varint,5,opt,name=person_id,json=personId,proto3" json:"person_id"`
	MasterId    int64   `protobuf:"varint,6,opt,name=master_id,json=masterId,proto3" json:"master_id"`
	Relation    string  `protobuf:"bytes,7,opt,name=relation,proto3" json:"relation"`
	Status      string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	RegionId    int64   `protobuf:"varint,9,opt,name=region_id,json=regionId,proto3" json:"region_id"`
	GroupId     int64   `protobuf:"varint,10,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Addr        string  `protobuf:"bytes,11,opt,name=addr,proto3" json:"addr"`
	Keywords    string  `protobuf:"bytes,12,opt,name=keywords,proto3" json:"keywords"`
	Ids         []int64 `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids"`
	Type        string  `protobuf:"bytes,14,opt,name=type,proto3" json:"type"`
	TownCode    int64   `protobuf:"varint,15,opt,name=town_code,json=townCode,proto3" json:"town_code"`
	VillageCode int64   `protobuf:"varint,16,opt,name=village_code,json=villageCode,proto3" json:"village_code"`
	GroupCode   int64   `protobuf:"varint,17,opt,name=group_code,json=groupCode,proto3" json:"group_code"`
	IsPoor      string  `protobuf:"bytes,18,opt,name=is_poor,json=isPoor,proto3" json:"is_poor"`
}

func (x *FamilyRequest) Reset() {
	*x = FamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyRequest) ProtoMessage() {}

func (x *FamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_familyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyRequest.ProtoReflect.Descriptor instead.
func (*FamilyRequest) Descriptor() ([]byte, []int) {
	return file_familyService_proto_rawDescGZIP(), []int{1}
}

func (x *FamilyRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FamilyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FamilyRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FamilyRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *FamilyRequest) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FamilyRequest) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *FamilyRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *FamilyRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FamilyRequest) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *FamilyRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *FamilyRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *FamilyRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *FamilyRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FamilyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FamilyRequest) GetTownCode() int64 {
	if x != nil {
		return x.TownCode
	}
	return 0
}

func (x *FamilyRequest) GetVillageCode() int64 {
	if x != nil {
		return x.VillageCode
	}
	return 0
}

func (x *FamilyRequest) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *FamilyRequest) GetIsPoor() string {
	if x != nil {
		return x.IsPoor
	}
	return ""
}

type FamilyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Family       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Family     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *FamilyData) Reset() {
	*x = FamilyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyData) ProtoMessage() {}

func (x *FamilyData) ProtoReflect() protoreflect.Message {
	mi := &file_familyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyData.ProtoReflect.Descriptor instead.
func (*FamilyData) Descriptor() ([]byte, []int) {
	return file_familyService_proto_rawDescGZIP(), []int{2}
}

func (x *FamilyData) GetEntity() *Family {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FamilyData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FamilyData) GetItems() []*Family {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FamilyData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type FamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *FamilyData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *FamilyResponse) Reset() {
	*x = FamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyResponse) ProtoMessage() {}

func (x *FamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_familyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyResponse.ProtoReflect.Descriptor instead.
func (*FamilyResponse) Descriptor() ([]byte, []int) {
	return file_familyService_proto_rawDescGZIP(), []int{3}
}

func (x *FamilyResponse) GetData() *FamilyData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FamilyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_familyService_proto protoreflect.FileDescriptor

var file_familyService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x05, 0x0a, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x50, 0x6f, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x0a, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x22, 0xed, 0x03,
	0x0a, 0x0d, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x77, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6f, 0x72, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x50, 0x6f, 0x6f, 0x72, 0x22, 0xa5, 0x01,
	0x0a, 0x0a, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe8, 0x02, 0x0a, 0x0d, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_familyService_proto_rawDescOnce sync.Once
	file_familyService_proto_rawDescData = file_familyService_proto_rawDesc
)

func file_familyService_proto_rawDescGZIP() []byte {
	file_familyService_proto_rawDescOnce.Do(func() {
		file_familyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_familyService_proto_rawDescData)
	})
	return file_familyService_proto_rawDescData
}

var file_familyService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_familyService_proto_goTypes = []interface{}{
	(*Family)(nil),         // 0: services.Family
	(*FamilyRequest)(nil),  // 1: services.FamilyRequest
	(*FamilyData)(nil),     // 2: services.FamilyData
	(*FamilyResponse)(nil), // 3: services.FamilyResponse
	(*Person)(nil),         // 4: services.Person
	(*Region)(nil),         // 5: services.Region
	(*Group)(nil),          // 6: services.Group
	(*common.Pager)(nil),   // 7: common.Pager
	(*common.Info)(nil),    // 8: common.Info
	(*common.Error)(nil),   // 9: common.Error
}
var file_familyService_proto_depIdxs = []int32{
	4,  // 0: services.Family.master:type_name -> services.Person
	5,  // 1: services.Family.region:type_name -> services.Region
	6,  // 2: services.Family.group:type_name -> services.Group
	4,  // 3: services.Family.membership:type_name -> services.Person
	0,  // 4: services.FamilyData.entity:type_name -> services.Family
	7,  // 5: services.FamilyData.pager:type_name -> common.Pager
	0,  // 6: services.FamilyData.items:type_name -> services.Family
	8,  // 7: services.FamilyData.info:type_name -> common.Info
	2,  // 8: services.FamilyResponse.data:type_name -> services.FamilyData
	9,  // 9: services.FamilyResponse.error:type_name -> common.Error
	0,  // 10: services.FamilyService.Create:input_type -> services.Family
	0,  // 11: services.FamilyService.Update:input_type -> services.Family
	0,  // 12: services.FamilyService.Delete:input_type -> services.Family
	0,  // 13: services.FamilyService.Get:input_type -> services.Family
	1,  // 14: services.FamilyService.Search:input_type -> services.FamilyRequest
	1,  // 15: services.FamilyService.List:input_type -> services.FamilyRequest
	3,  // 16: services.FamilyService.Create:output_type -> services.FamilyResponse
	3,  // 17: services.FamilyService.Update:output_type -> services.FamilyResponse
	3,  // 18: services.FamilyService.Delete:output_type -> services.FamilyResponse
	3,  // 19: services.FamilyService.Get:output_type -> services.FamilyResponse
	3,  // 20: services.FamilyService.Search:output_type -> services.FamilyResponse
	3,  // 21: services.FamilyService.List:output_type -> services.FamilyResponse
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_familyService_proto_init() }
func file_familyService_proto_init() {
	if File_familyService_proto != nil {
		return
	}
	file_personService_proto_init()
	file_regionService_proto_init()
	file_groupService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_familyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Family); i {
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
		file_familyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyRequest); i {
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
		file_familyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyData); i {
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
		file_familyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyResponse); i {
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
			RawDescriptor: file_familyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_familyService_proto_goTypes,
		DependencyIndexes: file_familyService_proto_depIdxs,
		MessageInfos:      file_familyService_proto_msgTypes,
	}.Build()
	File_familyService_proto = out.File
	file_familyService_proto_rawDesc = nil
	file_familyService_proto_goTypes = nil
	file_familyService_proto_depIdxs = nil
}
