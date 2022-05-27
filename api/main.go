package main

import (
	"dousheng/api/router"
	"dousheng/api/rpc"
	"dousheng/config"
	"dousheng/pkg/constant"
	"fmt"
	"log"
)

func main() {

	config.Init("../config")
	config.ConfInstance().WithServerConfig(constant.ServerComment)
	config.ConfInstance().WithServerConfig(constant.ServerVideo)
	rpc.InitCommentRPC()
	rpc.InitVideoRPC()
	rpc.InitUserRPC()

	fmt.Println(config.ConfInstance())

	// 路由注册
	r := router.Setup()
	addr := ":" + config.ConfInstance().BaseConfig.Port
	//addr := fmt.Sprintf("%s:%s", config.ConfInstance().BaseConfig.Host, config.ConfInstance().BaseConfig.Port)

	// 启动 API 服务
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
