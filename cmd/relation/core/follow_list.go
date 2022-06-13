package core

import (
	"context"
	"dousheng/cmd/relation/dal/mysqldb"
	"dousheng/cmd/relation/dal/redisdb"
	relation "dousheng/cmd/relation/service"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"strconv"
)

func (*RelationService) FollowList(ctx context.Context, req *relation.FollowListRequest,
	resp *relation.FollowerListResponse) error {
	userId := req.GetUserId()

	// 解析token
	token, _ := doushengjwt.ParseToken(req.Token)
	selfId := token.UserID
	// 返回关注列表 和自己的关注列表
	followIDs, followedIDs, err := redisdb.FollowList(userId, selfId)
	if err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}

	// 把自己关注列表的user设置为true
	followed := make(map[string]bool, len(followedIDs))
	for _, id := range followedIDs {
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
		if _, ok := followed[idStr]; !ok { // 说明有没关注的
			user.IsFollow = false
		}
	}

	resp.UserList = usersInfo
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = tmp.StatusCode
	resp.StatusMsg = tmp.StatusMessage
	return nil
}
