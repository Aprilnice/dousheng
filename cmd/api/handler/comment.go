package handler

import (
	"context"
	"dousheng/cmd/comment/service"
	"dousheng/pkg/constant"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	middlewares "dousheng/pkg/middleware"
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
	fmt.Println("------------评论列表请求--------------")
	fmt.Println(clp)
	fmt.Println("------------------------------------")

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
	// rpc 调用
	commentService := c.Keys[constant.ClientComment].(comment.DyCommentService)
	resp, err := commentService.CommentList(context.Background(), &req)

	if err != nil {
		HttpResponse(c, resp)
		return
	}

	fmt.Println(resp.CommentList)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.StatusMsg)

	HttpResponse(c, resp)

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

	var uid int64
	if uidStr, ok := c.Get(middlewares.ContextUserID); ok {
		uid = uidStr.(int64)
	}

	req := comment.CommentActionRequest{
		UserId:      uid,
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

func createCommentHandler(c *gin.Context, req *comment.CommentActionRequest) {

	if len(req.CommentText) <= 0 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("评论不能为空"))
		return
	}

	// 取出实例
	commentService := c.Keys[constant.ClientComment].(comment.DyCommentService)
	// rpc 调用
	resp, err := commentService.CreateComment(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		HttpResponse(c, resp)
		return
	}
	HttpResponse(c, resp)
}

func deleteCommentHandler(c *gin.Context, req *comment.CommentActionRequest) {
	// 声明一个response

	// 取出实例
	commentService := c.Keys[constant.ClientComment].(comment.DyCommentService)
	// rpc 调用
	resp, err := commentService.DeleteComment(context.Background(), req)
	if err != nil {
		HttpResponse(c, resp)
		return
	}
	HttpResponse(c, resp)
}
