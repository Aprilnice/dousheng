package rpc

import (
	video "dousheng/cmd/video/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
)

var VideoRPC video.VideoModuleService

// InitVideoRPC 初始化视频服务客户端
func InitVideoRPC() {

	microVideo := micro.NewService(
		micro.Name(constant.ClientVideo),
	)

	VideoRPC = video.NewVideoModuleService(
		config.Instance().ServerConfig.Server(constant.ServerVideo).Name,
		microVideo.Client(),
	)

}
