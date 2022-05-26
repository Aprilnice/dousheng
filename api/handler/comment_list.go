package handler

import (
	"context"
	"dousheng/api/rpc"
	comments "dousheng/comment/service"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CommentListHandler 获取评论列表
func CommentListHandler(c *gin.Context) {

	// 绑定请求参数
	var clp CommentListParam
	if err := c.ShouldBindQuery(&clp); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token
	token, err := doushengjwt.ParseToken(clp.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}

	// 构造请求数据
	req := comments.CommentListRequest{
		UserId:  token.UserID,
		Token:   clp.Token,
		VideoId: clp.VideoId,
	}
	resp, err := rpc.CommentList(context.Background(), &req)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err))
		return
	}

	HttpResponse(c, errdeal.NewResponse(errdeal.CodeSuccess).WithData(resp.CommentList))

}
