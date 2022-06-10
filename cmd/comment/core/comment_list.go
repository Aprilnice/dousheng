package core

import (
	"context"
	commentDB "dousheng/cmd/comment/dal/mysqldb"
	"dousheng/cmd/comment/service"
	user "dousheng/cmd/user/dal/mysqldb"
	"dousheng/pkg/format"
	"fmt"
	"sync"
)

// CommentList 获取评论列表
func (*CommentService) CommentList(ctx context.Context, req *comment.CommentListRequest, resp *comment.CommentListResponse) (err error) {

	fmt.Println("评论列表Service")
	var wg = sync.WaitGroup{}
	wg.Add(2)
	var err1, err2 error
	var preComments []*commentDB.Comment // 存评论额的
	var usersInfo []*user.UserInfo       // 用户信息
	fmt.Println("videoID: ->", req.VideoId)
	// 开狗肉听
	// 查评论列表
	go func() {
		defer wg.Done()
		preComments, err1 = commentDB.VideoCommentList(req.VideoId)
	}()
	// 查评论用户信息
	go func() {
		defer wg.Done()
		usersInfo, err2 = commentDB.CommentUserInfoByVideoID(req.VideoId)
	}()

	wg.Wait()
	if err1 != nil {
		ResponseError(err1).CommentListResponse(resp)
		return err1
	}
	if err2 != nil {
		ResponseError(err1).CommentListResponse(resp)
		return err1
	}

	fmt.Printf("comments: %#v", preComments)

	if len(preComments) == 0 {
		ResponseSuccess().CommentListResponse(resp)
		return nil
	}

	// id -> userInfo
	var commentUsers = make(map[int64]*comment.User, len(usersInfo))
	for _, info := range usersInfo {
		cu := comment.User{
			Id:            info.UserId,
			Name:          info.Name,
			FollowCount:   info.FollowCount,
			FollowerCount: info.FollowerCount,
			IsFollow:      info.IsFollow,
		}
		commentUsers[info.UserId] = &cu
	}
	// 打包评论信息
	var commentsList = make([]*comment.Comment, len(preComments))
	for i, comm := range preComments {
		if _, ok := commentUsers[comm.UserID]; !ok { // 查不到这条用户信息
			continue
		}
		c := comment.Comment{
			Id:         comm.CommentID,
			User:       commentUsers[comm.UserID],
			Content:    comm.CommentText,
			CreateDate: format.GormDateFormat(comm.CreatedAt),
		}
		commentsList[i] = &c
	}

	ResponseSuccess().CommentListResponse(resp)
	resp.CommentList = commentsList // 填充数据
	fmt.Println(resp.CommentList)
	return nil
}
