package rpc

import (
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
		micro.Name("commentRpcClient"),
	)
	rpcVideo.Init()

	rpcVideoService = service.NewVideoModuleService("srv.comment", rpcVideo.Client())

}