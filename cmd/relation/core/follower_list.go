package core

import (
	"context"
	"dousheng/cmd/relation/dal/mysqldb"
	"dousheng/cmd/relation/dal/redisdb"
	relation "dousheng/cmd/relation/service"
	"dousheng/pkg/errdeal"
	"strconv"
)

func (*RelationService) FollowerList(ctx context.Context, req *relation.FollowerListRequest,
	resp *relation.FollowerListResponse) error {
	userId := req.GetUserId()
	followIDs, followedIds, err := redisdb.FollowerList(userId)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}
	followed := make(map[string]bool, len(followedIds))
	for _, id := range followedIds {
		followed[id] = true
	}

	usersInfo, err := mysqldb.UsersInfo(followIDs)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}
	for _, user := range usersInfo {
		idStr := strconv.FormatInt(user.Id, 10)
		if _, ok := followed[idStr]; !ok { // 说明粉丝里有没关注的
			user.IsFollow = false
		}
	}

	resp.UserList = usersInfo
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = tmp.StatusCode
	resp.StatusMsg = tmp.StatusMessage
	return nil
}
