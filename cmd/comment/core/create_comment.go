package core

import (
	"context"
	"dousheng/cmd/comment/dal/mysqldb"
	"dousheng/cmd/comment/service"
	"dousheng/pkg/snowflaker"
)

//CreateComment 创建评论
func (*CommentService) CreateComment(ctx context.Context, req *comment.CommentActionRequest, resp *comment.CommentActionResponse) (err error) {
	// 初始化Comment
	commentModel := new(mysqldb.CommentAction)
	// 生成commentID
	req.CommentId = snowflaker.NextID()

	// 绑定数据
	if err = commentModel.BindWithReq(req); err != nil {
		// 出现错误 创建一个NewResponse 并绑定到CommentResponse中返回
		ResponseError(err).CommentActionResponse(resp)
		return err
	}
	// 存储
	if err = mysqldb.CreateComment(commentModel); err != nil {
		// 出现错误  这里一般都是数据库错误
		ResponseError(err).CommentActionResponse(resp)
		return err
	}
	// 成功
	ResponseSuccess().CommentActionResponse(resp)

	return nil

}
