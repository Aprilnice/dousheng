package redisdb

import (
	"dousheng/pkg/rediskey"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// DoFollow 关注
func DoFollow(userID, toUserID int64) error {
	// 关注列表新增一个用户
	// 本人关注列表key:
	userIDStr := strconv.FormatInt(userID, 10)
	toUserIDStr := strconv.FormatInt(toUserID, 10)
	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, userIDStr)
	pipeline := rdb.Pipeline()
	pipeline.ZAddNX(ctx, followKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: toUserID,
	})
	// 对方粉丝；列表新增一个粉丝
	followerKey := rediskey.NewRedisKey(rediskey.KeyUserFollower, toUserIDStr)
	pipeline.ZAddNX(ctx, followerKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: userID,
	})
	// 修改用户信息
	userKey := rediskey.NewRedisKey(rediskey.KeyUserHash, userIDStr)
	toUserKey := rediskey.NewRedisKey(rediskey.KeyUserHash, toUserIDStr)
	pipeline.HIncrBy(ctx, userKey, "FollowCount", 1)     // 关注加一
	pipeline.HIncrBy(ctx, toUserKey, "FollowerCount", 1) // 对面粉丝加一
	_, err := pipeline.Exec(ctx)
	return err

}

// CancelFollow 取消关注
func CancelFollow(userID, toUserID int64) error {
	// 关注列表删除一个用户
	// 本人关注列表key:
	userIDStr := strconv.FormatInt(userID, 10)
	toUserIDStr := strconv.FormatInt(toUserID, 10)
	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, userIDStr)
	pipeline := rdb.Pipeline()
	pipeline.ZRem(ctx, followKey, toUserID)
	// 对方粉丝；列表删除一个粉丝
	followerKey := rediskey.NewRedisKey(rediskey.KeyUserFollower, toUserIDStr)
	pipeline.ZRem(ctx, followerKey, userID)
	// 修改用户信息
	userKey := rediskey.NewRedisKey(rediskey.KeyUserHash, userIDStr)
	toUserKey := rediskey.NewRedisKey(rediskey.KeyUserHash, toUserIDStr)
	pipeline.HIncrBy(ctx, userKey, "FollowCount", -1)     // 关注减一
	pipeline.HIncrBy(ctx, toUserKey, "FollowerCount", -1) // 对面粉丝减一
	_, err := pipeline.Exec(ctx)
	return err
}

// FollowList 关注列表
func FollowList(userID, selfID int64) ([]string, []string, error) {
	userIDStr := strconv.FormatInt(userID, 10)
	selfIDStr := strconv.FormatInt(selfID, 10)
	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, userIDStr)
	follows, err := rdb.ZRevRangeByScore(ctx, followKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	var followed []string
	if selfID == userID { // 如果查看的是自己的关注列表
		followed = follows
	} else {
		// 自己的关注列表
		selfFollowKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, selfIDStr)
		isFollowKey := rediskey.NewRedisKey(rediskey.KeyIsFollow, userIDStr, selfIDStr)
		followed = interStore(isFollowKey, selfFollowKey, followKey)
	}
	return follows, followed, err
}

// FollowerList 粉丝列表  同时返回已关注的
func FollowerList(userID, selfID int64) ([]string, []string, error) {
	userIDStr := strconv.FormatInt(userID, 10)
	selfIDStr := strconv.FormatInt(selfID, 10)
	// 粉丝列表
	followerKey := rediskey.NewRedisKey(rediskey.KeyUserFollower, userIDStr)
	followers, err := rdb.ZRevRangeByScore(ctx, followerKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	// 自己的关注列表
	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, selfIDStr)
	// ISFollow
	isFollowKey := rediskey.NewRedisKey(rediskey.KeyIsFollow, userIDStr, selfIDStr)
	// 返回交集
	followed := interStore(isFollowKey, followKey, followerKey)

	return followers, followed, err
}

func interStore(key, key1, key2 string) []string {

	if rdb.Exists(ctx, key).Val() < 1 {
		rdb.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{key1, key2},
			Aggregate: "MAX",
		})
		rdb.Expire(ctx, key, 60*time.Second) // 设置1分钟过期
	}
	followed, _ := rdb.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	return followed
}
