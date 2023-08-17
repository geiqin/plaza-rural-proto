// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dashboardService.proto

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

// Api Endpoints for DashboardService service

func NewDashboardServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DashboardService service

type DashboardService interface {
	Index(ctx context.Context, in *DashboardRequest, opts ...client.CallOption) (*DashboardResponse, error)
}

type dashboardService struct {
	c    client.Client
	name string
}

func NewDashboardService(name string, c client.Client) DashboardService {
	return &dashboardService{
		c:    c,
		name: name,
	}
}

func (c *dashboardService) Index(ctx context.Context, in *DashboardRequest, opts ...client.CallOption) (*DashboardResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.Index", in)
	out := new(DashboardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DashboardService service

type DashboardServiceHandler interface {
	Index(context.Context, *DashboardRequest, *DashboardResponse) error
}

func RegisterDashboardServiceHandler(s server.Server, hdlr DashboardServiceHandler, opts ...server.HandlerOption) error {
	type dashboardService interface {
		Index(ctx context.Context, in *DashboardRequest, out *DashboardResponse) error
	}
	type DashboardService struct {
		dashboardService
	}
	h := &dashboardServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DashboardService{h}, opts...))
}

type dashboardServiceHandler struct {
	DashboardServiceHandler
}

func (h *dashboardServiceHandler) Index(ctx context.Context, in *DashboardRequest, out *DashboardResponse) error {
	return h.DashboardServiceHandler.Index(ctx, in, out)
}