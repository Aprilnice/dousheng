// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: comment.proto

package comment

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

// Api Endpoints for DyComment service

func NewDyCommentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DyComment service

type DyCommentService interface {
	CreateComment(ctx context.Context, in *CommentRequest, opts ...client.CallOption) (*CommentResponse, error)
	DeleteComment(ctx context.Context, in *CommentRequest, opts ...client.CallOption) (*CommentResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...client.CallOption) (*CommentListResponse, error)
}

type dyCommentService struct {
	c    client.Client
	name string
}

func NewDyCommentService(name string, c client.Client) DyCommentService {
	return &dyCommentService{
		c:    c,
		name: name,
	}
}

func (c *dyCommentService) CreateComment(ctx context.Context, in *CommentRequest, opts ...client.CallOption) (*CommentResponse, error) {
	req := c.c.NewRequest(c.name, "DyComment.CreateComment", in)
	out := new(CommentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dyCommentService) DeleteComment(ctx context.Context, in *CommentRequest, opts ...client.CallOption) (*CommentResponse, error) {
	req := c.c.NewRequest(c.name, "DyComment.DeleteComment", in)
	out := new(CommentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dyCommentService) CommentList(ctx context.Context, in *CommentListRequest, opts ...client.CallOption) (*CommentListResponse, error) {
	req := c.c.NewRequest(c.name, "DyComment.CommentList", in)
	out := new(CommentListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DyComment service

type DyCommentHandler interface {
	CreateComment(context.Context, *CommentRequest, *CommentResponse) error
	DeleteComment(context.Context, *CommentRequest, *CommentResponse) error
	CommentList(context.Context, *CommentListRequest, *CommentListResponse) error
}

func RegisterDyCommentHandler(s server.Server, hdlr DyCommentHandler, opts ...server.HandlerOption) error {
	type dyComment interface {
		CreateComment(ctx context.Context, in *CommentRequest, out *CommentResponse) error
		DeleteComment(ctx context.Context, in *CommentRequest, out *CommentResponse) error
		CommentList(ctx context.Context, in *CommentListRequest, out *CommentListResponse) error
	}
	type DyComment struct {
		dyComment
	}
	h := &dyCommentHandler{hdlr}
	return s.Handle(s.NewHandler(&DyComment{h}, opts...))
}

type dyCommentHandler struct {
	DyCommentHandler
}

func (h *dyCommentHandler) CreateComment(ctx context.Context, in *CommentRequest, out *CommentResponse) error {
	return h.DyCommentHandler.CreateComment(ctx, in, out)
}

func (h *dyCommentHandler) DeleteComment(ctx context.Context, in *CommentRequest, out *CommentResponse) error {
	return h.DyCommentHandler.DeleteComment(ctx, in, out)
}

func (h *dyCommentHandler) CommentList(ctx context.Context, in *CommentListRequest, out *CommentListResponse) error {
	return h.DyCommentHandler.CommentList(ctx, in, out)
}
