// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: relation.proto

package relation

import (
	fmt "fmt"
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

// Api Endpoints for Relation service

func NewRelationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Relation service

type RelationService interface {
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...client.CallOption) (*RelationActionResponse, error)
	FollowList(ctx context.Context, in *FollowListRequest, opts ...client.CallOption) (*FollowerListResponse, error)
	FollowerList(ctx context.Context, in *FollowerListRequest, opts ...client.CallOption) (*FollowerListResponse, error)
}

type relationService struct {
	c    client.Client
	name string
}

func NewRelationService(name string, c client.Client) RelationService {
	return &relationService{
		c:    c,
		name: name,
	}
}

func (c *relationService) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...client.CallOption) (*RelationActionResponse, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationAction", in)
	out := new(RelationActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) FollowList(ctx context.Context, in *FollowListRequest, opts ...client.CallOption) (*FollowerListResponse, error) {
	req := c.c.NewRequest(c.name, "Relation.FollowList", in)
	out := new(FollowerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) FollowerList(ctx context.Context, in *FollowerListRequest, opts ...client.CallOption) (*FollowerListResponse, error) {
	req := c.c.NewRequest(c.name, "Relation.FollowerList", in)
	out := new(FollowerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Relation service

type RelationHandler interface {
	RelationAction(context.Context, *RelationActionRequest, *RelationActionResponse) error
	FollowList(context.Context, *FollowListRequest, *FollowerListResponse) error
	FollowerList(context.Context, *FollowerListRequest, *FollowerListResponse) error
}

func RegisterRelationHandler(s server.Server, hdlr RelationHandler, opts ...server.HandlerOption) error {
	type relation interface {
		RelationAction(ctx context.Context, in *RelationActionRequest, out *RelationActionResponse) error
		FollowList(ctx context.Context, in *FollowListRequest, out *FollowerListResponse) error
		FollowerList(ctx context.Context, in *FollowerListRequest, out *FollowerListResponse) error
	}
	type Relation struct {
		relation
	}
	h := &relationHandler{hdlr}
	return s.Handle(s.NewHandler(&Relation{h}, opts...))
}

type relationHandler struct {
	RelationHandler
}

func (h *relationHandler) RelationAction(ctx context.Context, in *RelationActionRequest, out *RelationActionResponse) error {
	return h.RelationHandler.RelationAction(ctx, in, out)
}

func (h *relationHandler) FollowList(ctx context.Context, in *FollowListRequest, out *FollowerListResponse) error {
	return h.RelationHandler.FollowList(ctx, in, out)
}

func (h *relationHandler) FollowerList(ctx context.Context, in *FollowerListRequest, out *FollowerListResponse) error {
	return h.RelationHandler.FollowerList(ctx, in, out)
}