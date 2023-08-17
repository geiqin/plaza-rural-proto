// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: groupService.proto

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

// Api Endpoints for GroupService service

func NewGroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GroupService service

type GroupService interface {
	Create(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error)
	Update(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error)
	Delete(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error)
	Get(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error)
	List(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error)
	Search(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error)
}

type groupService struct {
	c    client.Client
	name string
}

func NewGroupService(name string, c client.Client) GroupService {
	return &groupService{
		c:    c,
		name: name,
	}
}

func (c *groupService) Create(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Create", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) Update(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Update", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) Delete(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Delete", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) Get(ctx context.Context, in *Group, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Get", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) List(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.List", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) Search(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Search", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceHandler interface {
	Create(context.Context, *Group, *GroupResponse) error
	Update(context.Context, *Group, *GroupResponse) error
	Delete(context.Context, *Group, *GroupResponse) error
	Get(context.Context, *Group, *GroupResponse) error
	List(context.Context, *GroupRequest, *GroupResponse) error
	Search(context.Context, *GroupRequest, *GroupResponse) error
}

func RegisterGroupServiceHandler(s server.Server, hdlr GroupServiceHandler, opts ...server.HandlerOption) error {
	type groupService interface {
		Create(ctx context.Context, in *Group, out *GroupResponse) error
		Update(ctx context.Context, in *Group, out *GroupResponse) error
		Delete(ctx context.Context, in *Group, out *GroupResponse) error
		Get(ctx context.Context, in *Group, out *GroupResponse) error
		List(ctx context.Context, in *GroupRequest, out *GroupResponse) error
		Search(ctx context.Context, in *GroupRequest, out *GroupResponse) error
	}
	type GroupService struct {
		groupService
	}
	h := &groupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GroupService{h}, opts...))
}

type groupServiceHandler struct {
	GroupServiceHandler
}

func (h *groupServiceHandler) Create(ctx context.Context, in *Group, out *GroupResponse) error {
	return h.GroupServiceHandler.Create(ctx, in, out)
}

func (h *groupServiceHandler) Update(ctx context.Context, in *Group, out *GroupResponse) error {
	return h.GroupServiceHandler.Update(ctx, in, out)
}

func (h *groupServiceHandler) Delete(ctx context.Context, in *Group, out *GroupResponse) error {
	return h.GroupServiceHandler.Delete(ctx, in, out)
}

func (h *groupServiceHandler) Get(ctx context.Context, in *Group, out *GroupResponse) error {
	return h.GroupServiceHandler.Get(ctx, in, out)
}

func (h *groupServiceHandler) List(ctx context.Context, in *GroupRequest, out *GroupResponse) error {
	return h.GroupServiceHandler.List(ctx, in, out)
}

func (h *groupServiceHandler) Search(ctx context.Context, in *GroupRequest, out *GroupResponse) error {
	return h.GroupServiceHandler.Search(ctx, in, out)
}
