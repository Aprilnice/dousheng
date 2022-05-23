package main

import (
	"dousheng/api/handler"
	"dousheng/api/rpc"
	"github.com/gin-gonic/gin"
)

func main() {

	rpc.InitCommentRPC()

	r := gin.Default()

	r.POST("/dousheng/api/action/comment", handler.CommentHandler)

	r.Run("127.0.0.1:9010")

}
