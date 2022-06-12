package rpc

import (
	favorite "dousheng/cmd/favourite/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
)

var FavoriteRPC favorite.FavoriteService

// InitFavoriteRPC 相当于初始化客户端调用
func InitFavoriteRPC() {

	rpcComment := micro.NewService(
		micro.Name(constant.ClientFavorite),
	)
	rpcComment.Init()

	FavoriteRPC = favorite.NewFavoriteService(
		config.Instance().ServerConfig.Server(constant.ServerFavorite).Name,
		rpcComment.Client(),
	)

}
