package comment

import (
	"dousheng/comment/core"
	"dousheng/comment/service"
	"dousheng/config"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)

func Run() {
	// 获取配置文件
	commentConf := *(config.ConfInstance())
	(&commentConf).WithServerConfig("comment")
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
