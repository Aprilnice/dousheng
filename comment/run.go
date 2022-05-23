package comment

import (
	"dousheng/comment/core"
	comment "dousheng/comment/service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)

func Run() {
	// 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.141.101:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("srv.comment"),
		micro.Registry(etcdReg),
		micro.Address("127.0.0.1:8088"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*10),
	)

	microService.Init()

	// 服务注册
	err := comment.RegisterCommentHandler(microService.Server(), new(core.CommentService))
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
