package rpc

import (
	"context"
	"dousheng/config"
	"dousheng/pkg/constant"
	"dousheng/user/service"
	user "dousheng/user/service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var rpcUserService service.UserService

func InitUserRPC() {
	microReg := etcd.NewRegistry(
		registry.Addrs(config.ConfInstance().EtcdConfig.Address),
	)

	rpcUser := micro.NewService(
		micro.Registry(microReg),
		micro.Name("rpcUserClient"),
	)
	rpcUser.Init()

	rpcUserService = user.NewUserService(
		config.ConfInstance().ServerConfig.Server(constant.ServerUser).Name,
		rpcUser.Client(),
	)

}
func Register(ctx context.Context, req *service.DouyinUserRegisterRequest) (resp *service.DouyinUserRegisterResponse, err error) {
	resp, err = rpcUserService.Register(ctx, req)
	return
}
func Login(ctx context.Context, req *service.DouyinUserLoginRequest) (resp *service.DouyinUserLoginResponse, err error) {
	resp, err = rpcUserService.Login(ctx, req)
	return
}

func UserInfo(ctx context.Context, req *service.DouyinUserRequest) (resp *service.DouyinUserResponse, err error) {
	resp, err = rpcUserService.UserInfo(ctx, req)
	return
}
