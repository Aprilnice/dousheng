package main

import (
	"dousheng/api/handler"
	"dousheng/api/rpc"
	"dousheng/config"
	"github.com/gin-gonic/gin"
)

func main() {

	rpc.InitCommentRPC()
	rpc.InitVideoRPC()

	r := gin.Default()

	// 评论
	r.POST("/dousheng/api/action/comment", handler.CommentHandler)
	// 视频发布
	r.POST("/douyin/publish/action", handler.VideoPublishHandler)
	// 视频播放
	r.POST("/play/:video_id", handler.VideoPlayHandler)
	// 获取封面
	r.POST("/cover/:cover_id", handler.GetCoverHandler)

	add := config.ConfInstance().BaseConfig.Host+":"+config.ConfInstance().BaseConfig.Port
	r.Run(add)

}
