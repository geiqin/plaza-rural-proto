// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personTrainService.proto

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

// Api Endpoints for PersonTrainService service

func NewPersonTrainServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonTrainService service

type PersonTrainService interface {
	//创建人员培训
	Create(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error)
	// 编辑人员培训
	Update(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error)
	// 删除人员培训
	Delete(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error)
	// 获取人员培训详情
	Get(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error)
	//分页查询人员培训列表
	Search(ctx context.Context, in *PersonTrainRequest, opts ...client.CallOption) (*PersonTrainResponse, error)
	//导出人员培训
	Export(ctx context.Context, in *PersonTrainRequest, opts ...client.CallOption) (*PersonTrainResponse, error)
}

type personTrainService struct {
	c    client.Client
	name string
}

func NewPersonTrainService(name string, c client.Client) PersonTrainService {
	return &personTrainService{
		c:    c,
		name: name,
	}
}

func (c *personTrainService) Create(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Create", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personTrainService) Update(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Update", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personTrainService) Delete(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Delete", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personTrainService) Get(ctx context.Context, in *PersonTrain, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Get", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personTrainService) Search(ctx context.Context, in *PersonTrainRequest, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Search", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personTrainService) Export(ctx context.Context, in *PersonTrainRequest, opts ...client.CallOption) (*PersonTrainResponse, error) {
	req := c.c.NewRequest(c.name, "PersonTrainService.Export", in)
	out := new(PersonTrainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonTrainService service

type PersonTrainServiceHandler interface {
	//创建人员培训
	Create(context.Context, *PersonTrain, *PersonTrainResponse) error
	// 编辑人员培训
	Update(context.Context, *PersonTrain, *PersonTrainResponse) error
	// 删除人员培训
	Delete(context.Context, *PersonTrain, *PersonTrainResponse) error
	// 获取人员培训详情
	Get(context.Context, *PersonTrain, *PersonTrainResponse) error
	//分页查询人员培训列表
	Search(context.Context, *PersonTrainRequest, *PersonTrainResponse) error
	//导出人员培训
	Export(context.Context, *PersonTrainRequest, *PersonTrainResponse) error
}

func RegisterPersonTrainServiceHandler(s server.Server, hdlr PersonTrainServiceHandler, opts ...server.HandlerOption) error {
	type personTrainService interface {
		Create(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error
		Update(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error
		Delete(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error
		Get(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error
		Search(ctx context.Context, in *PersonTrainRequest, out *PersonTrainResponse) error
		Export(ctx context.Context, in *PersonTrainRequest, out *PersonTrainResponse) error
	}
	type PersonTrainService struct {
		personTrainService
	}
	h := &personTrainServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonTrainService{h}, opts...))
}

type personTrainServiceHandler struct {
	PersonTrainServiceHandler
}

func (h *personTrainServiceHandler) Create(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Create(ctx, in, out)
}

func (h *personTrainServiceHandler) Update(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Update(ctx, in, out)
}

func (h *personTrainServiceHandler) Delete(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Delete(ctx, in, out)
}

func (h *personTrainServiceHandler) Get(ctx context.Context, in *PersonTrain, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Get(ctx, in, out)
}

func (h *personTrainServiceHandler) Search(ctx context.Context, in *PersonTrainRequest, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Search(ctx, in, out)
}

func (h *personTrainServiceHandler) Export(ctx context.Context, in *PersonTrainRequest, out *PersonTrainResponse) error {
	return h.PersonTrainServiceHandler.Export(ctx, in, out)
}