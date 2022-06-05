package rpc

import (
	service2 "dousheng/cmd/user/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
)

var UserRPC service2.UserService

func InitUserRPC() {

	microUser := micro.NewService(
		micro.Name(constant.ClientUser), // 客户端调用
	)

	UserRPC = service2.NewUserService(
		config.Instance().ServerConfig.Server(constant.ServerUser).Name,
		microUser.Client(),
	)

}
