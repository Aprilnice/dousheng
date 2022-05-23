package rpc

import (
	"context"
	"dousheng/comment/service"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
)

var rpcService service.CommentService

// InitCommentRPC 相当于初始化客户端调用
func InitCommentRPC() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	microReg := etcd.NewRegistry(
		registry.Addrs("192.168.141.101:2379"),
	)

	rpcComment := micro.NewService(
		micro.Registry(microReg),
		micro.Name("commentRpcClient"),
	)
	rpcComment.Init()

	rpcService = service.NewCommentService("srv.comment", rpcComment.Client())

}

// CreateComment 具体的调用
func CreateComment(ctx context.Context, req *service.CommentRequest) (resp *service.CommentResponse, err error) {

	resp, err = rpcService.CreateComment(ctx, req)
	return
}

func DeleteComment(ctx context.Context, req *service.CommentRequest) (resp *service.CommentResponse, err error) {

	resp, err = rpcService.DeleteComment(ctx, req)
	return
}
