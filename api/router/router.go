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
		// 评论  从url中获取token
		r.POST("/douyin/comment/action",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.CommentActionHandler)

		// 点赞
		//"/douyin/favorite/action/?token=douyin123456&video_id=2&action_type=1"
		r.POST("/douyin/favorite/action",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.FavoriteActionHandler)
		r.GET("/douyin/favorite/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.FavoriteListHandler)

		// 关注
		//"/douyin/favorite/list/?token=douyin123456"
		r.POST("/douyin/relation/action/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.RelationActionHandler)
		// 关注列表
		r.GET("/douyin/relation/follow/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.FollowListHandler)
		// 粉丝列表
		r.GET("/douyin/relation/follower/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler.FollowerListHandler)

		// 视频发布  从form-data中获取token  formKey : form-data的标识名

		//r.POST("/douyin/publish/action", middlewares.JwtTokenMiddleware(middlewares.FormToken("req")), handler.VideoPublishHandler)
		// 视频发布 test
		r.POST("/douyin/publish/action", handler.VideoPublishHandler)

	}
	// 视频播放
	r.GET("/play", handler.VideoPlayHandler)
	// 获取封面
	r.POST("/cover", handler.GetCoverHandler)
	// 获取视频流
	r.GET("/douyin/feed", handler.GetVideoFeedHandler)
	// 获取评论列表
	r.GET("/douyin/comment/list/", handler.CommentListHandler)

	return r

}
