package main

import (
	"dousheng/cmd/relation/core"
	"dousheng/cmd/relation/dal/mysqldb"
	"dousheng/cmd/relation/dal/redisdb"
	relation "dousheng/cmd/relation/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
)

func Init() {
	// 初始化配置
	config.Init("./config")

	// 初始化 用户服务的相关配置信息
	config.Instance().WithServerConfig(constant.ServerRelation)

	// 初始化数据库
	if err := mysqldb.Init(config.Instance().MySQLConfig); err != nil {
		log.Println("mysql数据库初始化失败")
		log.Fatal(err)
	}

	// 初始化redis
	if err := redisdb.Init(config.Instance().RedisConfig); err != nil {
		log.Println("redis数据库初始化失败")
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
		micro.Name(config.Instance().ServerConfig.Server(constant.ServerRelation).Name),
		micro.Registry(etcdReg),
		micro.Address(config.Instance().ServerConfig.Server(constant.ServerRelation).Address),
	)

	microService.Init()

	// 服务注册
	err := relation.RegisterRelationHandler(microService.Server(), new(core.RelationService))
	if err != nil {
		log.Println("服务注册失败失败")
		log.Fatal(err)
	}
	fmt.Println("[", microService.Name(), "]")
	// 启动服务
	if err = microService.Run(); err != nil {
		log.Println("服务启动失败")
		log.Fatal(err)
	}
}
