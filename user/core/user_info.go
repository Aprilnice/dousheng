package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
	"dousheng/user/service"
)

func (*UserService) UserInfo(ctx context.Context, req *service.DouyinUserRequest, res *service.DouyinUserResponse) error {
	userId := req.GetUserId()
	userinfo, err := mysqldb.GetUserInfoById(userId)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		return err
	}
	tmp := errdeal.NewResponse(errdeal.CodeSuccess).WithData("nil")
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
