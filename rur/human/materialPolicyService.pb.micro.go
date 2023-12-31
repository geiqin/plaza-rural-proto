// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: materialPolicyService.proto

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

// Api Endpoints for MaterialPolicyService service

func NewMaterialPolicyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MaterialPolicyService service

type MaterialPolicyService interface {
	//创建落实农资补助政策
	Create(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error)
	// 编辑落实农资补助政策
	Update(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error)
	// 删除落实农资补助政策
	Delete(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error)
	// 获取落实农资补助政策详情
	Get(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error)
	//分页查询落实农资补助政策列表
	Search(ctx context.Context, in *MaterialPolicyRequest, opts ...client.CallOption) (*MaterialPolicyResponse, error)
	//导出落实农资补助政策
	Export(ctx context.Context, in *MaterialPolicyRequest, opts ...client.CallOption) (*MaterialPolicyResponse, error)
}

type materialPolicyService struct {
	c    client.Client
	name string
}

func NewMaterialPolicyService(name string, c client.Client) MaterialPolicyService {
	return &materialPolicyService{
		c:    c,
		name: name,
	}
}

func (c *materialPolicyService) Create(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Create", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialPolicyService) Update(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Update", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialPolicyService) Delete(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Delete", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialPolicyService) Get(ctx context.Context, in *MaterialPolicy, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Get", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialPolicyService) Search(ctx context.Context, in *MaterialPolicyRequest, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Search", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialPolicyService) Export(ctx context.Context, in *MaterialPolicyRequest, opts ...client.CallOption) (*MaterialPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "MaterialPolicyService.Export", in)
	out := new(MaterialPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MaterialPolicyService service

type MaterialPolicyServiceHandler interface {
	//创建落实农资补助政策
	Create(context.Context, *MaterialPolicy, *MaterialPolicyResponse) error
	// 编辑落实农资补助政策
	Update(context.Context, *MaterialPolicy, *MaterialPolicyResponse) error
	// 删除落实农资补助政策
	Delete(context.Context, *MaterialPolicy, *MaterialPolicyResponse) error
	// 获取落实农资补助政策详情
	Get(context.Context, *MaterialPolicy, *MaterialPolicyResponse) error
	//分页查询落实农资补助政策列表
	Search(context.Context, *MaterialPolicyRequest, *MaterialPolicyResponse) error
	//导出落实农资补助政策
	Export(context.Context, *MaterialPolicyRequest, *MaterialPolicyResponse) error
}

func RegisterMaterialPolicyServiceHandler(s server.Server, hdlr MaterialPolicyServiceHandler, opts ...server.HandlerOption) error {
	type materialPolicyService interface {
		Create(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error
		Update(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error
		Delete(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error
		Get(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error
		Search(ctx context.Context, in *MaterialPolicyRequest, out *MaterialPolicyResponse) error
		Export(ctx context.Context, in *MaterialPolicyRequest, out *MaterialPolicyResponse) error
	}
	type MaterialPolicyService struct {
		materialPolicyService
	}
	h := &materialPolicyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MaterialPolicyService{h}, opts...))
}

type materialPolicyServiceHandler struct {
	MaterialPolicyServiceHandler
}

func (h *materialPolicyServiceHandler) Create(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Create(ctx, in, out)
}

func (h *materialPolicyServiceHandler) Update(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Update(ctx, in, out)
}

func (h *materialPolicyServiceHandler) Delete(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Delete(ctx, in, out)
}

func (h *materialPolicyServiceHandler) Get(ctx context.Context, in *MaterialPolicy, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Get(ctx, in, out)
}

func (h *materialPolicyServiceHandler) Search(ctx context.Context, in *MaterialPolicyRequest, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Search(ctx, in, out)
}

func (h *materialPolicyServiceHandler) Export(ctx context.Context, in *MaterialPolicyRequest, out *MaterialPolicyResponse) error {
	return h.MaterialPolicyServiceHandler.Export(ctx, in, out)
}
