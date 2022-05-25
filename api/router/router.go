package router

import (
	"dousheng/api/handler"
	middlewares "dousheng/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// Setup 路由注册
func Setup() *gin.Engine {

	r := gin.Default()

	// 注册业务路由

	// 登录业务路由

	// 这里面的业务路由使用到 jwt 验证 中间件
	{
		// 评论
		r.POST("/dousheng/comment/action", middlewares.JwtTokenMiddleware(), handler.CommentHandler)

		// 点赞

		// 视频发布
		r.POST("/douyin/publish/action", handler.VideoPublishHandler)

	}
	// 视频播放
	r.POST("/play", handler.VideoPlayHandler)
	// 获取封面
	r.POST("/cover", handler.GetCoverHandler)
	// 获取视频流
	r.POST("/douyin/feed",handler.GetVideoFeedHandler)

	return r

}
