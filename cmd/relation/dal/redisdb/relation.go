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
	_, err := pipeline.Exec(ctx)
	return err
}

// FollowList 关注列表
func FollowList(userID int64) ([]string, error) {
	userIDStr := strconv.FormatInt(userID, 10)
	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, userIDStr)
	follows, err := rdb.ZRevRangeByScore(ctx, followKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	return follows, err
}

// FollowerList 粉丝列表  同时返回已关注的
func FollowerList(userID int64) ([]string, []string, error) {
	userIDStr := strconv.FormatInt(userID, 10)
	followerKey := rediskey.NewRedisKey(rediskey.KeyUserFollower, userIDStr)
	followers, err := rdb.ZRevRangeByScore(ctx, followerKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	followKey := rediskey.NewRedisKey(rediskey.KeyUserFollow, userIDStr)
	// ISFollow
	isFollowKey := rediskey.NewRedisKey(rediskey.KeyIsFollow, userIDStr)
	if rdb.Exists(ctx, isFollowKey).Val() < 1 {
		rdb.ZInterStore(ctx, isFollowKey, &redis.ZStore{
			Keys:      []string{followKey, followerKey},
			Aggregate: "MAX",
		})
		rdb.Expire(ctx, isFollowKey, 180*time.Second) // 设置三分钟过期
	}
	followed, err := rdb.ZRevRangeByScore(ctx, isFollowKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	return followers, followed, err
}
