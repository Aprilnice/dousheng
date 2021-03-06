// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: favourite.proto

package favourite

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

// Api Endpoints for Favorite service

func NewFavoriteEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Favorite service

type FavoriteService interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...client.CallOption) (*FavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...client.CallOption) (*FavoriteListResponse, error)
}

type favoriteService struct {
	c    client.Client
	name string
}

func NewFavoriteService(name string, c client.Client) FavoriteService {
	return &favoriteService{
		c:    c,
		name: name,
	}
}

func (c *favoriteService) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...client.CallOption) (*FavoriteActionResponse, error) {
	req := c.c.NewRequest(c.name, "Favorite.FavoriteAction", in)
	out := new(FavoriteActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...client.CallOption) (*FavoriteListResponse, error) {
	req := c.c.NewRequest(c.name, "Favorite.FavoriteList", in)
	out := new(FavoriteListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Favorite service

type FavoriteHandler interface {
	FavoriteAction(context.Context, *FavoriteActionRequest, *FavoriteActionResponse) error
	FavoriteList(context.Context, *FavoriteListRequest, *FavoriteListResponse) error
}

func RegisterFavoriteHandler(s server.Server, hdlr FavoriteHandler, opts ...server.HandlerOption) error {
	type favorite interface {
		FavoriteAction(ctx context.Context, in *FavoriteActionRequest, out *FavoriteActionResponse) error
		FavoriteList(ctx context.Context, in *FavoriteListRequest, out *FavoriteListResponse) error
	}
	type Favorite struct {
		favorite
	}
	h := &favoriteHandler{hdlr}
	return s.Handle(s.NewHandler(&Favorite{h}, opts...))
}

type favoriteHandler struct {
	FavoriteHandler
}

func (h *favoriteHandler) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, out *FavoriteActionResponse) error {
	return h.FavoriteHandler.FavoriteAction(ctx, in, out)
}

func (h *favoriteHandler) FavoriteList(ctx context.Context, in *FavoriteListRequest, out *FavoriteListResponse) error {
	return h.FavoriteHandler.FavoriteList(ctx, in, out)
}
