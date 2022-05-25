package video

import (
	"dousheng/config"
	"dousheng/pkg/constant"
	"dousheng/video/core"
	"dousheng/video/service"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)

func VideoRun() {
	// 初始化 视频服务的相关配置信息
	config.ConfInstance().WithServerConfig(constant.ServerVideo)
	fmt.Println((*config.ConfInstance().ServerConfig).Server(constant.ServerVideo))
	// 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.ConfInstance().EtcdConfig.Address),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.ConfInstance().ServerConfig.Server(constant.ServerVideo).Name),
		micro.Registry(etcdReg),
		micro.Address(config.ConfInstance().ServerConfig.Server(constant.ServerVideo).Address),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*10),
	)

	microService.Init()

	// 服务注册
	err := service.RegisterVideoModuleHandler(microService.Server(), new(core.VideoModuleService))
	if err != nil {
		log.Println("服务注册失败失败")
		log.Fatal(err)
	}

	// 启动服务
	if err = microService.Run(); err != nil {
		log.Println("服务启动失败")
		log.Fatal(err)
	}
}