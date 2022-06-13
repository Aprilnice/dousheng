package main

import (
	"dousheng/cmd/favourite/core"
	"dousheng/cmd/favourite/dal/mysqldb"
	"dousheng/cmd/favourite/dal/redisdb"
	favorite "dousheng/cmd/favourite/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
)

func Init() {
	// 初始化配置
	config.Init("./config")

	// 初始化 评论服务的相关配置信息
	config.Instance().WithServerConfig(constant.ServerFavorite)

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
		micro.Name(config.Instance().ServerConfig.Server(constant.ServerFavorite).Name),
		micro.Address(config.Instance().ServerConfig.Server(constant.ServerFavorite).Address),
		micro.Registry(etcdReg),
	)

	microService.Init()
	// 服务注册
	err := favorite.RegisterFavoriteHandler(microService.Server(), new(core.FavouriteService))
	if err != nil {
		log.Println("点赞服务注册失败失败")
		log.Fatal(err)
	}

	println("[", microService.Name(), "]")

	// 启动服务
	if err = microService.Run(); err != nil {
		log.Println("服务启动失败")
		log.Fatal(err)
	}
}
