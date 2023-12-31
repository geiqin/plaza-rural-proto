// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: landResourceService.proto

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

// Api Endpoints for LandResourceService service

func NewLandResourceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LandResourceService service

type LandResourceService interface {
	//创建土地资源
	Create(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error)
	// 编辑土地资源
	Update(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error)
	// 删除土地资源
	Delete(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error)
	// 获取土地资源详情
	Get(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error)
	//分页查询土地资源
	Search(ctx context.Context, in *LandResourceRequest, opts ...client.CallOption) (*LandResourceResponse, error)
	//导出土地资源
	Export(ctx context.Context, in *LandResourceRequest, opts ...client.CallOption) (*LandResourceResponse, error)
}

type landResourceService struct {
	c    client.Client
	name string
}

func NewLandResourceService(name string, c client.Client) LandResourceService {
	return &landResourceService{
		c:    c,
		name: name,
	}
}

func (c *landResourceService) Create(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Create", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *landResourceService) Update(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Update", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *landResourceService) Delete(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Delete", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *landResourceService) Get(ctx context.Context, in *LandResource, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Get", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *landResourceService) Search(ctx context.Context, in *LandResourceRequest, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Search", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *landResourceService) Export(ctx context.Context, in *LandResourceRequest, opts ...client.CallOption) (*LandResourceResponse, error) {
	req := c.c.NewRequest(c.name, "LandResourceService.Export", in)
	out := new(LandResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LandResourceService service

type LandResourceServiceHandler interface {
	//创建土地资源
	Create(context.Context, *LandResource, *LandResourceResponse) error
	// 编辑土地资源
	Update(context.Context, *LandResource, *LandResourceResponse) error
	// 删除土地资源
	Delete(context.Context, *LandResource, *LandResourceResponse) error
	// 获取土地资源详情
	Get(context.Context, *LandResource, *LandResourceResponse) error
	//分页查询土地资源
	Search(context.Context, *LandResourceRequest, *LandResourceResponse) error
	//导出土地资源
	Export(context.Context, *LandResourceRequest, *LandResourceResponse) error
}

func RegisterLandResourceServiceHandler(s server.Server, hdlr LandResourceServiceHandler, opts ...server.HandlerOption) error {
	type landResourceService interface {
		Create(ctx context.Context, in *LandResource, out *LandResourceResponse) error
		Update(ctx context.Context, in *LandResource, out *LandResourceResponse) error
		Delete(ctx context.Context, in *LandResource, out *LandResourceResponse) error
		Get(ctx context.Context, in *LandResource, out *LandResourceResponse) error
		Search(ctx context.Context, in *LandResourceRequest, out *LandResourceResponse) error
		Export(ctx context.Context, in *LandResourceRequest, out *LandResourceResponse) error
	}
	type LandResourceService struct {
		landResourceService
	}
	h := &landResourceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LandResourceService{h}, opts...))
}

type landResourceServiceHandler struct {
	LandResourceServiceHandler
}

func (h *landResourceServiceHandler) Create(ctx context.Context, in *LandResource, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Create(ctx, in, out)
}

func (h *landResourceServiceHandler) Update(ctx context.Context, in *LandResource, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Update(ctx, in, out)
}

func (h *landResourceServiceHandler) Delete(ctx context.Context, in *LandResource, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Delete(ctx, in, out)
}

func (h *landResourceServiceHandler) Get(ctx context.Context, in *LandResource, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Get(ctx, in, out)
}

func (h *landResourceServiceHandler) Search(ctx context.Context, in *LandResourceRequest, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Search(ctx, in, out)
}

func (h *landResourceServiceHandler) Export(ctx context.Context, in *LandResourceRequest, out *LandResourceResponse) error {
	return h.LandResourceServiceHandler.Export(ctx, in, out)
}
