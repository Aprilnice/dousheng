package rpc

import (
	"context"
	comment "dousheng/comment/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var (
	rpcCommentService comment.DyCommentService
)

// InitCommentRPC 相当于初始化客户端调用
func InitCommentRPC() {

	microReg := etcd.NewRegistry(
		registry.Addrs(config.ConfInstance().EtcdConfig.Address),
	)

	rpcComment := micro.NewService(
		micro.Registry(microReg),
		micro.Name("rpcCommentClient"),
	)
	rpcComment.Init()

	rpcCommentService = comment.NewDyCommentService(
		config.ConfInstance().ServerConfig.Server(constant.ServerComment).Name,
		rpcComment.Client(),
	)

}

// CreateComment 具体的调用
func CreateComment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {

	resp, err = rpcCommentService.CreateComment(ctx, req)
	return
}

func DeleteComment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {

	resp, err = rpcCommentService.DeleteComment(ctx, req)
	return
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp, err = rpcCommentService.CommentList(ctx, req)
	return
}
