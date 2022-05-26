package core

import (
	"context"
	comment "dousheng/comment/service"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
)

// CommentList 获取评论列表
func (*CommentService) CommentList(ctx context.Context, req *comment.CommentListRequest, resp *comment.CommentListResponse) (err error) {
	lists, err := mysqldb.CommentListByVideoID(req.VideoId)
	if err != nil {
		ResponseErr(err).BindToCommentList(resp)
		return err
	}
	ids, err := mysqldb.CommentUserIDByVideoID(req.VideoId)
	if err != nil {
		ResponseErr(err).BindToCommentList(resp)
		return err
	}
	var users = make(map[int64]*comment.User, len(ids))
	for _, id := range ids {
		u := comment.User{
			Id:   id,
			Name: "Test",
		}
		users[id] = &u
	}
	var commentLists = make([]*comment.Comment, len(lists))
	for i, list := range lists {
		c := comment.Comment{
			User:       users[list.UserID],
			Content:    list.CommentText,
			CreateDate: list.CreatedAt.String(),
		}
		commentLists[i] = &c
	}

	resp.StatusCode = int32(errdeal.CodeSuccess)
	resp.StatusMsg = errdeal.CodeSuccess.Message()
	resp.CommentList = commentLists
	return nil
}
