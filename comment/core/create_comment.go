package core

import (
	"context"
	comment "dousheng/comment/service"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/snowflaker"
)

//CreateComment 创建评论
func (*CommentService) CreateComment(ctx context.Context, req *comment.CommentRequest, resp *comment.CommentResponse) (err error) {
	// 初始化Comment
	commentModel := new(mysqldb.Comment)
	// 生成commentID
	req.CommentId = snowflaker.NextID()

	// 绑定数据
	if err = commentModel.BindWithReq(req); err != nil {
		// 出现错误 创建一个NewResponse 并绑定到CommentResponse中返回
		ResponseErr(err).BindTo(resp)
		return err
	}
	// 存储
	if err = mysqldb.CreateComment(commentModel); err != nil {
		// 出现错误  这里一般都是数据库错误
		ResponseErr(err).BindTo(resp)
		return err
	}
	// 成功
	ResponseSuccess(nil).BindTo(resp)
	return nil

}
