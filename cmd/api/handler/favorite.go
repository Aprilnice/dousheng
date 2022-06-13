package handler

import (
	"context"
	favorite "dousheng/cmd/favourite/service"
	"dousheng/pkg/constant"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteActionHandler 点赞视频
func FavoriteActionHandler(c *gin.Context) {
	// "/douyin/favourite/action/?token=douyin123456&video_id=2&action_type=1"
	var favoriteParam FavoriteActionParam
	if err := c.ShouldBindQuery(&favoriteParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}
	// 校验参数 判断视频是否存在

	// 解析token
	token, err := doushengjwt.ParseToken(favoriteParam.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}
	// 绑定参数
	favoriteReq := favorite.FavoriteActionRequest{
		Token:      favoriteParam.Token,
		UserId:     token.UserID,
		VideoId:    favoriteParam.VideoId,
		ActionType: favoriteParam.ActionType,
	}
	// rpc 调用
	serviceRPC := c.Keys[constant.ClientFavorite].(favorite.FavoriteService)
	//ctx, cancelFunc := context.WithTimeout()
	//ctx := cancelFunc
	resp, err := serviceRPC.FavoriteAction(context.Background(), &favoriteReq)
	if err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr))
		return
	}

	HttpResponse(c, resp)

}

// FavoriteListHandler 赞过的视频
func FavoriteListHandler(c *gin.Context) {
	var favoriteParam FavoriteListParam
	if err := c.ShouldBindQuery(&favoriteParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}
	// 校验参数 判断视频是否存在

	// 解析token
	_, err := doushengjwt.ParseToken(favoriteParam.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}

	// 绑定参数
	favoriteReq := favorite.FavoriteListRequest{
		Token:  favoriteParam.Token,
		UserId: favoriteParam.UserId,
	}

	// rpc 调用
	serviceRPC := c.Keys[constant.ClientFavorite].(favorite.FavoriteService)
	resp, err := serviceRPC.FavoriteList(context.Background(), &favoriteReq)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr))
		return
	}

	HttpResponse(c, resp)

}
