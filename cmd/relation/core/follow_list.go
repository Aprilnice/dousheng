package core

import (
	"context"
	"dousheng/cmd/relation/dal/mysqldb"
	"dousheng/cmd/relation/dal/redisdb"
	relation "dousheng/cmd/relation/service"
	"dousheng/pkg/errdeal"
)

func (*RelationService) FollowList(ctx context.Context, req *relation.FollowListRequest,
	resp *relation.FollowerListResponse) error {
	userId := req.GetUserId()
	followIDs, err := redisdb.FollowList(userId)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}

	usersInfo, err := mysqldb.UsersInfo(followIDs)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}

	resp.UserList = usersInfo
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = tmp.StatusCode
	resp.StatusMsg = tmp.StatusMessage
	return nil
}
