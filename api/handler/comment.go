package handler

import (
	"context"
	"dousheng/api/rpc"
	comment "dousheng/comment/service"
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
	req := comment.CommentListRequest{
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

// CommentActionHandler 创建或删除评论
func CommentActionHandler(c *gin.Context) {
	/*
		"/douyin/comment/action/?token=douyin123456&video_id=1&action_type=1&comment_text=%E6%B3%A5%E5%9A%8E"
	*/

	var commentVar CommentParam
	if err := c.ShouldBindQuery(&commentVar); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	req := comment.CommentRequest{
		UserId:      commentVar.UserId,
		Token:       commentVar.Token,
		VideoId:     commentVar.VideoId,
		ActionType:  commentVar.ActionType,
		CommentText: commentVar.CommentText,
		CommentId:   commentVar.CommentId,
	}

	if req.ActionType == 1 {
		createCommentHandler(c, &req)
	} else if req.ActionType == 2 {
		deleteCommentHandler(c, &req)
	} else {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("请求参数错误"))
		return
	}

}

func createCommentHandler(c *gin.Context, req *comment.CommentRequest) {
	var response *errdeal.Response
	if len(req.CommentText) <= 0 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("评论不能为空"))
		return
	}
	// rpc 调用
	resp, err := rpc.CreateComment(context.Background(), req)

	if err != nil {
		response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode)).WithErr(err)
		HttpResponse(c, response)
		return
	}
	// 成功
	response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode))
	HttpResponse(c, response)
}

func deleteCommentHandler(c *gin.Context, req *comment.CommentRequest) {
	// 声明一个response
	var response *errdeal.Response
	// rpc 调用
	resp, err := rpc.DeleteComment(context.Background(), req)
	if err != nil {
		response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode)).WithErr(err)
		HttpResponse(c, response)
		return
	}
	response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode))
	HttpResponse(c, response)
	fmt.Println(resp)
}
