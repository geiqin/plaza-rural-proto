// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personChangeService.proto

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

// Api Endpoints for PersonChangeService service

func NewPersonChangeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonChangeService service

type PersonChangeService interface {
	//创建人员变更
	Create(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error)
	// 编辑人员变更
	Update(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error)
	// 删除人员变更
	Delete(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error)
	// 获取人员变更详情
	Get(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error)
	//分页查询家庭列表
	Search(ctx context.Context, in *PersonChangeRequest, opts ...client.CallOption) (*PersonChangeResponse, error)
}

type personChangeService struct {
	c    client.Client
	name string
}

func NewPersonChangeService(name string, c client.Client) PersonChangeService {
	return &personChangeService{
		c:    c,
		name: name,
	}
}

func (c *personChangeService) Create(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error) {
	req := c.c.NewRequest(c.name, "PersonChangeService.Create", in)
	out := new(PersonChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personChangeService) Update(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error) {
	req := c.c.NewRequest(c.name, "PersonChangeService.Update", in)
	out := new(PersonChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personChangeService) Delete(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error) {
	req := c.c.NewRequest(c.name, "PersonChangeService.Delete", in)
	out := new(PersonChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personChangeService) Get(ctx context.Context, in *PersonChange, opts ...client.CallOption) (*PersonChangeResponse, error) {
	req := c.c.NewRequest(c.name, "PersonChangeService.Get", in)
	out := new(PersonChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personChangeService) Search(ctx context.Context, in *PersonChangeRequest, opts ...client.CallOption) (*PersonChangeResponse, error) {
	req := c.c.NewRequest(c.name, "PersonChangeService.Search", in)
	out := new(PersonChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonChangeService service

type PersonChangeServiceHandler interface {
	//创建人员变更
	Create(context.Context, *PersonChange, *PersonChangeResponse) error
	// 编辑人员变更
	Update(context.Context, *PersonChange, *PersonChangeResponse) error
	// 删除人员变更
	Delete(context.Context, *PersonChange, *PersonChangeResponse) error
	// 获取人员变更详情
	Get(context.Context, *PersonChange, *PersonChangeResponse) error
	//分页查询家庭列表
	Search(context.Context, *PersonChangeRequest, *PersonChangeResponse) error
}

func RegisterPersonChangeServiceHandler(s server.Server, hdlr PersonChangeServiceHandler, opts ...server.HandlerOption) error {
	type personChangeService interface {
		Create(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error
		Update(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error
		Delete(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error
		Get(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error
		Search(ctx context.Context, in *PersonChangeRequest, out *PersonChangeResponse) error
	}
	type PersonChangeService struct {
		personChangeService
	}
	h := &personChangeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonChangeService{h}, opts...))
}

type personChangeServiceHandler struct {
	PersonChangeServiceHandler
}

func (h *personChangeServiceHandler) Create(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error {
	return h.PersonChangeServiceHandler.Create(ctx, in, out)
}

func (h *personChangeServiceHandler) Update(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error {
	return h.PersonChangeServiceHandler.Update(ctx, in, out)
}

func (h *personChangeServiceHandler) Delete(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error {
	return h.PersonChangeServiceHandler.Delete(ctx, in, out)
}

func (h *personChangeServiceHandler) Get(ctx context.Context, in *PersonChange, out *PersonChangeResponse) error {
	return h.PersonChangeServiceHandler.Get(ctx, in, out)
}

func (h *personChangeServiceHandler) Search(ctx context.Context, in *PersonChangeRequest, out *PersonChangeResponse) error {
	return h.PersonChangeServiceHandler.Search(ctx, in, out)
}
