package rpc

import (
	"context"
	"dousheng/comment/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var rpcCommentService service.CommentService

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

	rpcCommentService = service.NewCommentService(
		config.ConfInstance().ServerConfig.Server(constant.ServerComment).Name,
		rpcComment.Client(),
	)

}

// CreateComment 具体的调用
func CreateComment(ctx context.Context, req *service.CommentRequest) (resp *service.CommentResponse, err error) {

	resp, err = rpcCommentService.CreateComment(ctx, req)
	return
}

func DeleteComment(ctx context.Context, req *service.CommentRequest) (resp *service.CommentResponse, err error) {

	resp, err = rpcCommentService.DeleteComment(ctx, req)
	return
}
