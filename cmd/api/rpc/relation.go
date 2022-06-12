package rpc

import (
	relation "dousheng/cmd/relation/service"
	"dousheng/config"
	"dousheng/pkg/constant"
	"github.com/micro/go-micro/v2"
)

var RelationRPC relation.RelationService

func InitRelationRPC() {

	microRelation := micro.NewService(
		micro.Name(constant.ClientRelation), // 客户端调用
	)

	RelationRPC = relation.NewRelationService(
		config.Instance().ServerConfig.Server(constant.ServerRelation).Name,
		microRelation.Client(),
	)

}
