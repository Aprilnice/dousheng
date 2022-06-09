package redisdb

import (
	"dousheng/cmd/user/dal/mysqldb"
	"strconv"
)

// AddUserInfo 向redis中添加用户信息
func AddUserInfo(user *mysqldb.UserInfo) error {
	return 	rdb.HSet(
		ctx,
		// key值为用户id
		strconv.FormatInt(user.UserId, 10),
		// field 为 Name
		"Name",
		user.Name,
		// field 为 FollowCount
		"FollowCount",
		user.FollowCount,
		// field 为 FollowerCount
		"FollowerCount",
		user.FollowerCount,
		// field 为 PublishTime
	).Err()
}