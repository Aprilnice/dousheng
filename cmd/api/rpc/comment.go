package rpc

import (
	comment "dousheng/cmd/comment/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
)

// CommentRPC RPC 调用
var CommentRPC comment.DyCommentService

// InitCommentRPC 服务发现
func InitCommentRPC() {

	// 评论
	microComment := micro.NewService(
		micro.Name(constant.ClientComment), // 评论服务
	)

	// 调用 实例
	CommentRPC = comment.NewDyCommentService(
		config.Instance().ServerConfig.Server(constant.ServerComment).Name,
		microComment.Client(),
	)
}
