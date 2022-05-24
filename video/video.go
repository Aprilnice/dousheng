package main

import (
	"dousheng/config"
	"dousheng/video/core"
	"dousheng/video/service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)

func VideoRun() {
	// 获取配置文件
	commentConf := *(config.ConfInstance())
	(&commentConf).WithServerConfig("video")
	// 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(commentConf.EtcdConfig.Address),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(commentConf.ServerConfig.Name),
		micro.Registry(etcdReg),
		micro.Address(commentConf.ServerConfig.Address),
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