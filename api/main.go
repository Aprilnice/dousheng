package main

import (
	"dousheng/api/handler"
	"dousheng/api/rpc"
	"dousheng/config"
	"github.com/gin-gonic/gin"
)

func main() {

	rpc.InitCommentRPC()

	r := gin.Default()

	r.POST("/dousheng/api/action/comment", handler.CommentHandler)

	add := config.ConfInstance().BaseConfig.Host+":"+config.ConfInstance().BaseConfig.Port
	r.Run(add)

}
