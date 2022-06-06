package main

import (
	"dousheng/cmd/api/router"
	"dousheng/cmd/api/rpc"
	"dousheng/config"
	"dousheng/pkg/constant"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {

	config.Init("./config")

	config.Instance().WithServerConfig(constant.ServerComment)
	config.Instance().WithServerConfig(constant.ServerVideo)
	config.Instance().WithServerConfig(constant.ServerUser)

	rpc.Init()

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//addr := fmt.Sprintf("%s:%s", config.Instance().BaseConfig.Host, config.Instance().BaseConfig.Port)
	addr := fmt.Sprintf(":%s",config.Instance().BaseConfig.Port)

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address(addr),
		//将服务调用实例使用gin处理
		web.Handler(router.Setup()),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
