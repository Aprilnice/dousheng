package rpc

import (
	"context"
	"dousheng/user/service"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
)

var rpcUserService service.UserService

func InitUserRPC() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	microReg := etcd.NewRegistry(
		registry.Addrs("192.168.141.101:2379"),
	)

	rpcUser := micro.NewService(
		micro.Registry(microReg),
		micro.Name("userRpcClient"),
	)
	rpcUser.Init()

	rpcUserService = service.NewUserService("srv.user", rpcUser.Client())

}
func Register(ctx context.Context, req *service.DouyinUserRegisterRequest) (resp *service.DouyinUserRegisterResponse, err error) {
	resp, err = rpcUserService.Register(ctx, req)
	return
}
func Login(ctx context.Context, req *service.DouyinUserLoginRequest) (resp *service.DouyinUserLoginResponse, err error) {
	resp, err = rpcUserService.Login(ctx, req)
	return
}
