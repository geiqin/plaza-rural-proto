// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personHealthyService.proto

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

// Api Endpoints for PersonHealthyService service

func NewPersonHealthyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonHealthyService service

type PersonHealthyService interface {
	//创建健康状况
	Create(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error)
	// 编辑健康状况
	Update(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error)
	// 删除健康状况
	Delete(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error)
	// 获取健康状况详情
	Get(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error)
	//分页查询家庭列表
	Search(ctx context.Context, in *PersonHealthyRequest, opts ...client.CallOption) (*PersonHealthyResponse, error)
	//导出健康状况
	Export(ctx context.Context, in *PersonHealthyRequest, opts ...client.CallOption) (*PersonHealthyResponse, error)
}

type personHealthyService struct {
	c    client.Client
	name string
}

func NewPersonHealthyService(name string, c client.Client) PersonHealthyService {
	return &personHealthyService{
		c:    c,
		name: name,
	}
}

func (c *personHealthyService) Create(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Create", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personHealthyService) Update(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Update", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personHealthyService) Delete(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Delete", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personHealthyService) Get(ctx context.Context, in *PersonHealthy, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Get", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personHealthyService) Search(ctx context.Context, in *PersonHealthyRequest, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Search", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personHealthyService) Export(ctx context.Context, in *PersonHealthyRequest, opts ...client.CallOption) (*PersonHealthyResponse, error) {
	req := c.c.NewRequest(c.name, "PersonHealthyService.Export", in)
	out := new(PersonHealthyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonHealthyService service

type PersonHealthyServiceHandler interface {
	//创建健康状况
	Create(context.Context, *PersonHealthy, *PersonHealthyResponse) error
	// 编辑健康状况
	Update(context.Context, *PersonHealthy, *PersonHealthyResponse) error
	// 删除健康状况
	Delete(context.Context, *PersonHealthy, *PersonHealthyResponse) error
	// 获取健康状况详情
	Get(context.Context, *PersonHealthy, *PersonHealthyResponse) error
	//分页查询家庭列表
	Search(context.Context, *PersonHealthyRequest, *PersonHealthyResponse) error
	//导出健康状况
	Export(context.Context, *PersonHealthyRequest, *PersonHealthyResponse) error
}

func RegisterPersonHealthyServiceHandler(s server.Server, hdlr PersonHealthyServiceHandler, opts ...server.HandlerOption) error {
	type personHealthyService interface {
		Create(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error
		Update(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error
		Delete(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error
		Get(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error
		Search(ctx context.Context, in *PersonHealthyRequest, out *PersonHealthyResponse) error
		Export(ctx context.Context, in *PersonHealthyRequest, out *PersonHealthyResponse) error
	}
	type PersonHealthyService struct {
		personHealthyService
	}
	h := &personHealthyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonHealthyService{h}, opts...))
}

type personHealthyServiceHandler struct {
	PersonHealthyServiceHandler
}

func (h *personHealthyServiceHandler) Create(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Create(ctx, in, out)
}

func (h *personHealthyServiceHandler) Update(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Update(ctx, in, out)
}

func (h *personHealthyServiceHandler) Delete(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Delete(ctx, in, out)
}

func (h *personHealthyServiceHandler) Get(ctx context.Context, in *PersonHealthy, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Get(ctx, in, out)
}

func (h *personHealthyServiceHandler) Search(ctx context.Context, in *PersonHealthyRequest, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Search(ctx, in, out)
}

func (h *personHealthyServiceHandler) Export(ctx context.Context, in *PersonHealthyRequest, out *PersonHealthyResponse) error {
	return h.PersonHealthyServiceHandler.Export(ctx, in, out)
}
