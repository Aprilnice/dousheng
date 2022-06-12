package router

import (
	handler2 "dousheng/cmd/api/handler"
	middlewares "dousheng/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// Setup 路由注册
func Setup() *gin.Engine {

	r := gin.Default()

	r.Use(
		middlewares.SetupServiceMiddleware(),
	)

	// 注册业务路由
	r.POST("/douyin/user/register/", handler2.RegisterHandler)

	// 登录业务路由
	r.POST("/douyin/user/login/", handler2.LoginHandler)

	//个人界面
	r.GET("/douyin/user/", handler2.UserInfoHandler)

	// 这里面的业务路由使用到 jwt 验证 中间件
	{
		// 评论  从url中获取token
		r.POST("/douyin/comment/action/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.CommentActionHandler)

		// 点赞
		//"/douyin/favourite/action/?token=douyin123456&video_id=2&action_type=1"
		r.POST("/douyin/favourite/action/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.FavoriteActionHandler)
		r.GET("/douyin/favourite/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.FavoriteListHandler)

		// 关注
		//"/douyin/favourite/list/?token=douyin123456"
		r.POST("/douyin/relation/action/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.RelationActionHandler)
		// 关注列表
		r.GET("/douyin/relation/follow/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.FollowListHandler)
		// 粉丝列表
		r.GET("/douyin/relation/follower/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.FollowerListHandler)

		// 视频发布  从form-data中获取token  formKey : form-data的标识名
		r.POST("/douyin/publish/action/", middlewares.JwtTokenMiddleware(middlewares.FormToken("token")), handler2.VideoPublishHandler)
		// 视频发布 test
		//r.POST("/douyin/publish/action/", handler2.VideoPublishHandler)

		// 获取用户发布视频列表
		r.GET("/douyin/publish/list/",
			middlewares.JwtTokenMiddleware(middlewares.URLToken("token")),
			handler2.GetVideoListHandler)

	}
	// 视频播放
	r.GET("/play", handler2.VideoPlayHandler)
	// 获取封面
	r.GET("/cover", handler2.GetCoverHandler)
	// 获取视频流
	r.GET("/douyin/feed", handler2.GetVideoFeedHandler)
	// 获取评论列表
	r.GET("/douyin/comment/list/", handler2.CommentListHandler)

	return r

}
