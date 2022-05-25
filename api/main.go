package main

import (
	"dousheng/api/router"
	"dousheng/api/rpc"
	"dousheng/config"
	"dousheng/pkg/constant"
	"log"
)

func main() {

	config.Init()
	config.ConfInstance().WithServerConfig(constant.ServerComment)
	rpc.InitCommentRPC()
	// 路由注册
	r := router.Setup()

	addr := config.ConfInstance().BaseConfig.Host + ":" + config.ConfInstance().BaseConfig.Port
	//addr := fmt.Sprintf("%s:%s", config.ConfInstance().BaseConfig.Host, config.ConfInstance().BaseConfig.Port)

	// 启动 API 服务
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}

}
