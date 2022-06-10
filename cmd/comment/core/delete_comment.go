package core

import (
	"context"
	"dousheng/cmd/comment/dal/mysqldb"
	"dousheng/cmd/comment/dal/redisdb"
	"dousheng/cmd/comment/service"
	"fmt"
)

//DeleteComment 删除评论
func (*CommentService) DeleteComment(ctx context.Context, req *comment.CommentActionRequest, resp *comment.CommentActionResponse) (err error) {
	fmt.Println("delete", req)
	if err = redisdb.DeleteComment(req.VideoId); err != nil {
		ResponseError(err).CommentActionResponse(resp)
		return err
	}
	if err = mysqldb.DeleteComment(req.GetCommentId(), req.GetVideoId()); err != nil {
		ResponseError(err).CommentActionResponse(resp)
		return err
	}

	ResponseSuccess().CommentActionResponse(resp)
	return nil
}
