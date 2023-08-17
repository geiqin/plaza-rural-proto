// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: familyWaterService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FamilyWaterService service

func NewFamilyWaterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FamilyWaterService service

type FamilyWaterService interface {
	//创建饮水情况
	Create(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error)
	// 编辑饮水情况
	Update(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error)
	// 删除饮水情况
	Delete(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error)
	// 获取饮水情况详情
	Get(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error)
	//分页查询家庭列表
	Search(ctx context.Context, in *FamilyWaterRequest, opts ...client.CallOption) (*FamilyWaterResponse, error)
	//导出饮水保障
	Export(ctx context.Context, in *FamilyWaterRequest, opts ...client.CallOption) (*FamilyWaterResponse, error)
}

type familyWaterService struct {
	c    client.Client
	name string
}

func NewFamilyWaterService(name string, c client.Client) FamilyWaterService {
	return &familyWaterService{
		c:    c,
		name: name,
	}
}

func (c *familyWaterService) Create(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Create", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyWaterService) Update(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Update", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyWaterService) Delete(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Delete", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyWaterService) Get(ctx context.Context, in *FamilyWater, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Get", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyWaterService) Search(ctx context.Context, in *FamilyWaterRequest, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Search", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyWaterService) Export(ctx context.Context, in *FamilyWaterRequest, opts ...client.CallOption) (*FamilyWaterResponse, error) {
	req := c.c.NewRequest(c.name, "FamilyWaterService.Export", in)
	out := new(FamilyWaterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FamilyWaterService service

type FamilyWaterServiceHandler interface {
	//创建饮水情况
	Create(context.Context, *FamilyWater, *FamilyWaterResponse) error
	// 编辑饮水情况
	Update(context.Context, *FamilyWater, *FamilyWaterResponse) error
	// 删除饮水情况
	Delete(context.Context, *FamilyWater, *FamilyWaterResponse) error
	// 获取饮水情况详情
	Get(context.Context, *FamilyWater, *FamilyWaterResponse) error
	//分页查询家庭列表
	Search(context.Context, *FamilyWaterRequest, *FamilyWaterResponse) error
	//导出饮水保障
	Export(context.Context, *FamilyWaterRequest, *FamilyWaterResponse) error
}

func RegisterFamilyWaterServiceHandler(s server.Server, hdlr FamilyWaterServiceHandler, opts ...server.HandlerOption) error {
	type familyWaterService interface {
		Create(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error
		Update(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error
		Delete(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error
		Get(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error
		Search(ctx context.Context, in *FamilyWaterRequest, out *FamilyWaterResponse) error
		Export(ctx context.Context, in *FamilyWaterRequest, out *FamilyWaterResponse) error
	}
	type FamilyWaterService struct {
		familyWaterService
	}
	h := &familyWaterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FamilyWaterService{h}, opts...))
}

type familyWaterServiceHandler struct {
	FamilyWaterServiceHandler
}

func (h *familyWaterServiceHandler) Create(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Create(ctx, in, out)
}

func (h *familyWaterServiceHandler) Update(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Update(ctx, in, out)
}

func (h *familyWaterServiceHandler) Delete(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Delete(ctx, in, out)
}

func (h *familyWaterServiceHandler) Get(ctx context.Context, in *FamilyWater, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Get(ctx, in, out)
}

func (h *familyWaterServiceHandler) Search(ctx context.Context, in *FamilyWaterRequest, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Search(ctx, in, out)
}

func (h *familyWaterServiceHandler) Export(ctx context.Context, in *FamilyWaterRequest, out *FamilyWaterResponse) error {
	return h.FamilyWaterServiceHandler.Export(ctx, in, out)
}
