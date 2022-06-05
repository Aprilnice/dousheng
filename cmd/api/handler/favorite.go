package handler

import (
	"dousheng/pkg/errdeal"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteActionHandler 点赞视频
func FavoriteActionHandler(c *gin.Context) {
	// "/douyin/favorite/action/?token=douyin123456&video_id=2&action_type=1"
	var favoriteParam FavoriteActionParam
	if err := c.ShouldBindQuery(&favoriteParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}
	// 校验参数 判断视频是否存在
	//// 绑定参数
	//favoriteReq := favorite.FavoriteActionRequest{
	//	Token:      favoriteParam.Token,
	//	VideoId:    favoriteParam.VideoId,
	//	ActionType: favoriteParam.ActionType,
	//}
	//// rpc 调用
	//resp, err := rpc.FavoriteAction(context.Background(), &favoriteReq)
	//
	//println(resp, err)
	HttpResponse(c, errdeal.NewResponse(errdeal.CodeSuccess))

}

// FavoriteListHandler 赞过的视频
func FavoriteListHandler(c *gin.Context) {

}
