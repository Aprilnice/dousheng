package main

import (
	"dousheng/cmd/user/core"
	"dousheng/cmd/user/dal/mysqldb"
	"dousheng/cmd/user/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"dousheng/pkg/snowflaker"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
)

func Init() {
	// 初始化配置
	config.Init("./config")

	// 初始化 评论服务的相关配置信息
	config.Instance().WithServerConfig(constant.ServerUser)

	// 初始化ID生成器
	if err := snowflaker.Init(
		config.Instance().StartTime, // 起始时间
		config.Instance().ServerConfig.Server(constant.ServerUser).MachineID, // 不同的服务不同的机器id
	); err != nil {

		log.Println("ID 生成器初始化失败")
		log.Fatal(err)
	}

	// 初始化数据库
	if err := mysqldb.Init(config.Instance().MySQLConfig); err != nil {
		log.Println("mysql数据库初始化失败")
		log.Fatal(err)
	}
}

func main() {

	Init()
	// 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.Instance().EtcdConfig.Address),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.Instance().ServerConfig.Server(constant.ServerUser).Name),
		micro.Registry(etcdReg),
		micro.Address(config.Instance().ServerConfig.Server(constant.ServerUser).Address),
	)

	microService.Init()

	// 服务注册
	err := service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
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
