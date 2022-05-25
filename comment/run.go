package comment

import (
	"dousheng/comment/core"
	"dousheng/comment/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)

func Run() {

	// 初始化 评论服务的相关配置信息
	config.ConfInstance().WithServerConfig(constant.ServerComment)
	fmt.Println((*config.ConfInstance().ServerConfig).Server(constant.ServerComment))

	// 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.ConfInstance().EtcdConfig.Address),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.ConfInstance().ServerConfig.Server(constant.ServerComment).Name),
		micro.Registry(etcdReg),
		micro.Address(config.ConfInstance().ServerConfig.Server(constant.ServerComment).Address),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*10),
	)

	microService.Init()

	// 服务注册
	err := service.RegisterCommentHandler(microService.Server(), new(core.CommentService))
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
