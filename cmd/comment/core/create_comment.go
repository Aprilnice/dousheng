package core

import (
	"context"
	"dousheng/cmd/comment/dal/mysqldb"
	"dousheng/cmd/comment/service"
	userDB "dousheng/cmd/user/dal/mysqldb"
	"dousheng/pkg/format"
	"dousheng/pkg/snowflaker"
	"fmt"
	"sync"
	"time"
)

//CreateComment 创建评论
func (*CommentService) CreateComment(ctx context.Context, req *comment.CommentActionRequest, resp *comment.CommentActionResponse) (err error) {
	// 初始化Comment
	commentModel := new(mysqldb.Comment)
	// 生成commentID
	req.CommentId = snowflaker.NextID()

	// 绑定数据
	if err = commentModel.BindWithReq(req); err != nil {
		// 出现错误 创建一个NewResponse 并绑定到CommentResponse中返回
		ResponseError(err).CommentActionResponse(resp)
		return err
	}
	var user *userDB.UserInfo
	var wg sync.WaitGroup
	wg.Add(2)
	// 查询用户信息
	go func() {
		defer wg.Done()
		user, err = mysqldb.CommentUserInfo(req.UserId)
	}()

	// 存储
	go func() {
		defer wg.Done()
		err = mysqldb.CreateComment(commentModel)
	}()
	wg.Wait()

	if err != nil {
		ResponseError(err).CommentActionResponse(resp)
		fmt.Println("err: ", err)
		return err
	}
	fmt.Println("user: ", user, req.UserId)

	commUser := comment.User{
		Id:            user.UserId,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
	commentReturn := comment.Comment{
		User:       &commUser,
		Id:         req.CommentId,
		Content:    req.CommentText,
		CreateDate: format.GormDateFormat(time.Now()),
	}
	// 成功
	ResponseSuccess().CommentActionResponse(resp)
	resp.Comment = &commentReturn
	return nil

}
