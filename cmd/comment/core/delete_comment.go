package core

import (
	"context"
	"dousheng/cmd/comment/dal/mysqldb"
	"dousheng/cmd/comment/service"
	"fmt"
)

//DeleteComment 删除评论
func (*CommentService) DeleteComment(ctx context.Context, req *comment.CommentRequest, resp *comment.CommentResponse) (err error) {
	fmt.Println("delete", req)

	if err = mysqldb.DeleteComment(req.GetCommentId()); err != nil {
		ResponseErr(err).BindTo(resp)
		return err
	}
	ResponseSuccess(nil).BindTo(resp)
	return nil
}
