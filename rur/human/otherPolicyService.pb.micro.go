// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: otherPolicyService.proto

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

// Api Endpoints for OtherPolicyService service

func NewOtherPolicyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OtherPolicyService service

type OtherPolicyService interface {
	//创建其他部门政策落实情况
	Create(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error)
	// 编辑其他部门政策落实情况
	Update(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error)
	// 删除其他部门政策落实情况
	Delete(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error)
	// 获取其他部门政策落实情况详情
	Get(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error)
	//分页查询其他部门政策落实情况列表
	Search(ctx context.Context, in *OtherPolicyRequest, opts ...client.CallOption) (*OtherPolicyResponse, error)
	//导出其他部门政策落实情况
	Export(ctx context.Context, in *OtherPolicyRequest, opts ...client.CallOption) (*OtherPolicyResponse, error)
}

type otherPolicyService struct {
	c    client.Client
	name string
}

func NewOtherPolicyService(name string, c client.Client) OtherPolicyService {
	return &otherPolicyService{
		c:    c,
		name: name,
	}
}

func (c *otherPolicyService) Create(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Create", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherPolicyService) Update(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Update", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherPolicyService) Delete(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Delete", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherPolicyService) Get(ctx context.Context, in *OtherPolicy, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Get", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherPolicyService) Search(ctx context.Context, in *OtherPolicyRequest, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Search", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherPolicyService) Export(ctx context.Context, in *OtherPolicyRequest, opts ...client.CallOption) (*OtherPolicyResponse, error) {
	req := c.c.NewRequest(c.name, "OtherPolicyService.Export", in)
	out := new(OtherPolicyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OtherPolicyService service

type OtherPolicyServiceHandler interface {
	//创建其他部门政策落实情况
	Create(context.Context, *OtherPolicy, *OtherPolicyResponse) error
	// 编辑其他部门政策落实情况
	Update(context.Context, *OtherPolicy, *OtherPolicyResponse) error
	// 删除其他部门政策落实情况
	Delete(context.Context, *OtherPolicy, *OtherPolicyResponse) error
	// 获取其他部门政策落实情况详情
	Get(context.Context, *OtherPolicy, *OtherPolicyResponse) error
	//分页查询其他部门政策落实情况列表
	Search(context.Context, *OtherPolicyRequest, *OtherPolicyResponse) error
	//导出其他部门政策落实情况
	Export(context.Context, *OtherPolicyRequest, *OtherPolicyResponse) error
}

func RegisterOtherPolicyServiceHandler(s server.Server, hdlr OtherPolicyServiceHandler, opts ...server.HandlerOption) error {
	type otherPolicyService interface {
		Create(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error
		Update(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error
		Delete(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error
		Get(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error
		Search(ctx context.Context, in *OtherPolicyRequest, out *OtherPolicyResponse) error
		Export(ctx context.Context, in *OtherPolicyRequest, out *OtherPolicyResponse) error
	}
	type OtherPolicyService struct {
		otherPolicyService
	}
	h := &otherPolicyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OtherPolicyService{h}, opts...))
}

type otherPolicyServiceHandler struct {
	OtherPolicyServiceHandler
}

func (h *otherPolicyServiceHandler) Create(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Create(ctx, in, out)
}

func (h *otherPolicyServiceHandler) Update(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Update(ctx, in, out)
}

func (h *otherPolicyServiceHandler) Delete(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Delete(ctx, in, out)
}

func (h *otherPolicyServiceHandler) Get(ctx context.Context, in *OtherPolicy, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Get(ctx, in, out)
}

func (h *otherPolicyServiceHandler) Search(ctx context.Context, in *OtherPolicyRequest, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Search(ctx, in, out)
}

func (h *otherPolicyServiceHandler) Export(ctx context.Context, in *OtherPolicyRequest, out *OtherPolicyResponse) error {
	return h.OtherPolicyServiceHandler.Export(ctx, in, out)
}
