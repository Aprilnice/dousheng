package rpc

import (
	"context"
	"dousheng/video/service"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
)

var rpcVideoService service.VideoModuleService

// InitVideoRPC 初始化视频服务客户端
func InitVideoRPC() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	microReg := etcd.NewRegistry(
		registry.Addrs("192.168.141.101:2379"),
	)

	rpcVideo := micro.NewService(
		micro.Registry(microReg),
		micro.Name("videoRpcClient"),
	)
	rpcVideo.Init()

	rpcVideoService = service.NewVideoModuleService("srv.video", rpcVideo.Client())

}

// VideoPublish 调用视频发布
func VideoPublish(c context.Context, req *service.DouyinPublishActionRequest) (resp *service.DouyinPublishActionResponse, err error) {
	resp, err = rpcVideoService.VideoPublish(c, req)
	return
}

// VideoPlay 调用视频播放
func VideoPlay(c context.Context, req *service.PlayVideoReq) (resp *service.PlayVideoResp, err error) {
	resp, err = rpcVideoService.PlayVideo(c, req)
	return
}

// CoverGet 调用封面下载
func CoverGet(c context.Context, req *service.GetCoverReq) (resp *service.GetCoverResp, err error) {
	resp, err = rpcVideoService.GetCover(c, req)
	return
}