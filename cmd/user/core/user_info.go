package core

import (
	"context"
	"dousheng/cmd/user/dal/mysqldb"
	"dousheng/cmd/user/service"
	"dousheng/pkg/errdeal"
)

func (*UserService) UserInfo(ctx context.Context, req *service.DouyinUserRequest, res *service.DouyinUserResponse) error {
	userId := req.GetUserId()
	userinfo, err := mysqldb.GetUserInfoById(userId)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		return err
	}
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	res.StatusCode = tmp.StatusCode
	res.StatusMsg = tmp.StatusMessage
	res.User = &service.User{
		Id:            userinfo.UserId,
		Name:          userinfo.Name,
		FollowCount:   userinfo.FollowCount,
		FollowerCount: userinfo.FollowerCount,
		IsFollow:      userinfo.IsFollow,
	}
	return nil
}
