// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: incomeService.proto

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

// Api Endpoints for IncomeService service

func NewIncomeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IncomeService service

type IncomeService interface {
	//批量创建家庭收入
	BatchCreate(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error)
	//创建家庭收入
	Create(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error)
	// 编辑家庭收入
	Update(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error)
	// 删除家庭收入
	Delete(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error)
	// 获取家庭收入
	Get(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error)
	//分页查询家庭列表
	Search(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error)
	//获取家庭列表
	List(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error)
}

type incomeService struct {
	c    client.Client
	name string
}

func NewIncomeService(name string, c client.Client) IncomeService {
	return &incomeService{
		c:    c,
		name: name,
	}
}

func (c *incomeService) BatchCreate(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.BatchCreate", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) Create(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.Create", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) Update(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.Update", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) Delete(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.Delete", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) Get(ctx context.Context, in *Income, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.Get", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) Search(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.Search", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeService) List(ctx context.Context, in *IncomeRequest, opts ...client.CallOption) (*IncomeResponse, error) {
	req := c.c.NewRequest(c.name, "IncomeService.List", in)
	out := new(IncomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IncomeService service

type IncomeServiceHandler interface {
	//批量创建家庭收入
	BatchCreate(context.Context, *IncomeRequest, *IncomeResponse) error
	//创建家庭收入
	Create(context.Context, *Income, *IncomeResponse) error
	// 编辑家庭收入
	Update(context.Context, *Income, *IncomeResponse) error
	// 删除家庭收入
	Delete(context.Context, *Income, *IncomeResponse) error
	// 获取家庭收入
	Get(context.Context, *Income, *IncomeResponse) error
	//分页查询家庭列表
	Search(context.Context, *IncomeRequest, *IncomeResponse) error
	//获取家庭列表
	List(context.Context, *IncomeRequest, *IncomeResponse) error
}

func RegisterIncomeServiceHandler(s server.Server, hdlr IncomeServiceHandler, opts ...server.HandlerOption) error {
	type incomeService interface {
		BatchCreate(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error
		Create(ctx context.Context, in *Income, out *IncomeResponse) error
		Update(ctx context.Context, in *Income, out *IncomeResponse) error
		Delete(ctx context.Context, in *Income, out *IncomeResponse) error
		Get(ctx context.Context, in *Income, out *IncomeResponse) error
		Search(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error
		List(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error
	}
	type IncomeService struct {
		incomeService
	}
	h := &incomeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IncomeService{h}, opts...))
}

type incomeServiceHandler struct {
	IncomeServiceHandler
}

func (h *incomeServiceHandler) BatchCreate(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error {
	return h.IncomeServiceHandler.BatchCreate(ctx, in, out)
}

func (h *incomeServiceHandler) Create(ctx context.Context, in *Income, out *IncomeResponse) error {
	return h.IncomeServiceHandler.Create(ctx, in, out)
}

func (h *incomeServiceHandler) Update(ctx context.Context, in *Income, out *IncomeResponse) error {
	return h.IncomeServiceHandler.Update(ctx, in, out)
}

func (h *incomeServiceHandler) Delete(ctx context.Context, in *Income, out *IncomeResponse) error {
	return h.IncomeServiceHandler.Delete(ctx, in, out)
}

func (h *incomeServiceHandler) Get(ctx context.Context, in *Income, out *IncomeResponse) error {
	return h.IncomeServiceHandler.Get(ctx, in, out)
}

func (h *incomeServiceHandler) Search(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error {
	return h.IncomeServiceHandler.Search(ctx, in, out)
}

func (h *incomeServiceHandler) List(ctx context.Context, in *IncomeRequest, out *IncomeResponse) error {
	return h.IncomeServiceHandler.List(ctx, in, out)
}
