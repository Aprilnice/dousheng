package handler

import (
	"dousheng/pkg/errdeal"
	"github.com/gin-gonic/gin"
)

// FavoriteActionHandler 点赞视频
func FavoriteActionHandler(c *gin.Context) {

	HttpResponse(c, errdeal.NewResponse(errdeal.CodeSuccess))

}

// FavoriteListHandler 赞过的视频
func FavoriteListHandler(c *gin.Context) {

}
