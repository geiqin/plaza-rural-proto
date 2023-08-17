// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personMedicalService.proto

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

// Api Endpoints for PersonMedicalService service

func NewPersonMedicalServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonMedicalService service

type PersonMedicalService interface {
	//创建医疗状况
	Create(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error)
	// 编辑医疗状况
	Update(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error)
	// 删除医疗状况
	Delete(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error)
	// 获取医疗状况详情
	Get(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error)
	//分页查询医疗状况列表
	Search(ctx context.Context, in *PersonMedicalRequest, opts ...client.CallOption) (*PersonMedicalResponse, error)
	//导出医疗保障
	Export(ctx context.Context, in *PersonMedicalRequest, opts ...client.CallOption) (*PersonMedicalResponse, error)
}

type personMedicalService struct {
	c    client.Client
	name string
}

func NewPersonMedicalService(name string, c client.Client) PersonMedicalService {
	return &personMedicalService{
		c:    c,
		name: name,
	}
}

func (c *personMedicalService) Create(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Create", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personMedicalService) Update(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Update", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personMedicalService) Delete(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Delete", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personMedicalService) Get(ctx context.Context, in *PersonMedical, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Get", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personMedicalService) Search(ctx context.Context, in *PersonMedicalRequest, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Search", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personMedicalService) Export(ctx context.Context, in *PersonMedicalRequest, opts ...client.CallOption) (*PersonMedicalResponse, error) {
	req := c.c.NewRequest(c.name, "PersonMedicalService.Export", in)
	out := new(PersonMedicalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonMedicalService service

type PersonMedicalServiceHandler interface {
	//创建医疗状况
	Create(context.Context, *PersonMedical, *PersonMedicalResponse) error
	// 编辑医疗状况
	Update(context.Context, *PersonMedical, *PersonMedicalResponse) error
	// 删除医疗状况
	Delete(context.Context, *PersonMedical, *PersonMedicalResponse) error
	// 获取医疗状况详情
	Get(context.Context, *PersonMedical, *PersonMedicalResponse) error
	//分页查询医疗状况列表
	Search(context.Context, *PersonMedicalRequest, *PersonMedicalResponse) error
	//导出医疗保障
	Export(context.Context, *PersonMedicalRequest, *PersonMedicalResponse) error
}

func RegisterPersonMedicalServiceHandler(s server.Server, hdlr PersonMedicalServiceHandler, opts ...server.HandlerOption) error {
	type personMedicalService interface {
		Create(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error
		Update(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error
		Delete(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error
		Get(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error
		Search(ctx context.Context, in *PersonMedicalRequest, out *PersonMedicalResponse) error
		Export(ctx context.Context, in *PersonMedicalRequest, out *PersonMedicalResponse) error
	}
	type PersonMedicalService struct {
		personMedicalService
	}
	h := &personMedicalServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonMedicalService{h}, opts...))
}

type personMedicalServiceHandler struct {
	PersonMedicalServiceHandler
}

func (h *personMedicalServiceHandler) Create(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Create(ctx, in, out)
}

func (h *personMedicalServiceHandler) Update(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Update(ctx, in, out)
}

func (h *personMedicalServiceHandler) Delete(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Delete(ctx, in, out)
}

func (h *personMedicalServiceHandler) Get(ctx context.Context, in *PersonMedical, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Get(ctx, in, out)
}

func (h *personMedicalServiceHandler) Search(ctx context.Context, in *PersonMedicalRequest, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Search(ctx, in, out)
}

func (h *personMedicalServiceHandler) Export(ctx context.Context, in *PersonMedicalRequest, out *PersonMedicalResponse) error {
	return h.PersonMedicalServiceHandler.Export(ctx, in, out)
}
