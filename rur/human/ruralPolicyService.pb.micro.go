// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ruralPolicyService.proto

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

// Api Endpoints for RuralPolicyService service

func NewRuralPolicyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RuralPolicyService service

type RuralPolicyService interface {
	//创建乡村振兴补助政策
	Create(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error)
	// 编辑乡村振兴补助政策
	Update(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error)
	// 删除乡村振兴补助政策
	Delete(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error)
	// 获取乡村振兴补助政策详情
	Get(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error)
	//分页查询乡村振兴补助政策列表
	Search(ctx context.Context, in *RuralPolicyRequest, opts ...client.CallOption) (*RuralPolicyResponse, error)
	//导出乡村振兴补助政策
	Export(ctx context.Context, in *RuralPolicyRequest, opts ...client.CallOption) (*RuralPolicyResponse, error)
}

type ruralPolicyService struct {
	c    client.Client
	name string
}

func NewRuralPolicyService(name string, c client.Client) RuralPolicyService {
	return &ruralPolicyService{
		c:    c,
		name: name,
	}
}

func (c *ruralPolicyService) Create(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Create", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruralPolicyService) Update(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Update", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruralPolicyService) Delete(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Delete", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruralPolicyService) Get(ctx context.Context, in *RuralPolicy, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Get", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruralPolicyService) Search(ctx context.Context, in *RuralPolicyRequest, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Search", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruralPolicyService) Export(ctx context.Context, in *RuralPolicyRequest, opts ...client.CallOption) (*RuralPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "RuralPolicyService.Export", in)
	out := new(RuralPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RuralPolicyService service

type RuralPolicyServiceHandler interface {
	//创建乡村振兴补助政策
	Create(context.Context, *RuralPolicy, *RuralPolicyResponse) error
	// 编辑乡村振兴补助政策
	Update(context.Context, *RuralPolicy, *RuralPolicyResponse) error
	// 删除乡村振兴补助政策
	Delete(context.Context, *RuralPolicy, *RuralPolicyResponse) error
	// 获取乡村振兴补助政策详情
	Get(context.Context, *RuralPolicy, *RuralPolicyResponse) error
	//分页查询乡村振兴补助政策列表
	Search(context.Context, *RuralPolicyRequest, *RuralPolicyResponse) error
	//导出乡村振兴补助政策
	Export(context.Context, *RuralPolicyRequest, *RuralPolicyResponse) error
}

func RegisterRuralPolicyServiceHandler(s server.Server, hdlr RuralPolicyServiceHandler, opts ...server.HandlerOption) error {
	type ruralPolicyService interface {
		Create(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error
		Update(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error
		Delete(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error
		Get(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error
		Search(ctx context.Context, in *RuralPolicyRequest, out *RuralPolicyResponse) error
		Export(ctx context.Context, in *RuralPolicyRequest, out *RuralPolicyResponse) error
	}
	type RuralPolicyService struct {
		ruralPolicyService
	}
	h := &ruralPolicyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RuralPolicyService{h}, opts...))
}

type ruralPolicyServiceHandler struct {
	RuralPolicyServiceHandler
}

func (h *ruralPolicyServiceHandler) Create(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Create(ctx, in, out)
}

func (h *ruralPolicyServiceHandler) Update(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Update(ctx, in, out)
}

func (h *ruralPolicyServiceHandler) Delete(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Delete(ctx, in, out)
}

func (h *ruralPolicyServiceHandler) Get(ctx context.Context, in *RuralPolicy, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Get(ctx, in, out)
}

func (h *ruralPolicyServiceHandler) Search(ctx context.Context, in *RuralPolicyRequest, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Search(ctx, in, out)
}

func (h *ruralPolicyServiceHandler) Export(ctx context.Context, in *RuralPolicyRequest, out *RuralPolicyResponse) error {
	return h.RuralPolicyServiceHandler.Export(ctx, in, out)
}