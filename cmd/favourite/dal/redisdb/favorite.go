package redisdb

import (
	favorite "dousheng/cmd/favorite/service"
	"dousheng/pkg/rediskey"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// CreateFavorite 在redis中新增一条点赞记录
func CreateFavorite(favorite *favorite.FavoriteActionRequest) error {
	pipeline := rdb.TxPipeline()
	uid := favorite.UserId
	vid := favorite.VideoId
	uidStr := strconv.FormatInt(uid, 10)

	// 增加用户点赞记录
	userFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, uidStr)

	pipeline.ZAdd(ctx, userFavoriteKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: vid,
	})
	// 点赞数加一
	vidStr := strconv.FormatInt(vid, 10)
	pipeline.HIncrBy(ctx,
		rediskey.NewRedisKey(rediskey.KeyVideoHash, vidStr),
		"FavoriteCount",
		1)
	_, err := pipeline.Exec(ctx)
	return err

}

// CancelFavorite 取消点赞
func CancelFavorite(favorite *favorite.FavoriteActionRequest) error {
	pipeline := rdb.TxPipeline()
	uid := favorite.UserId
	vid := favorite.VideoId
	uidStr := strconv.FormatInt(uid, 10)
	// 取消用户点赞记录
	userFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, uidStr)
	pipeline.ZRem(ctx, userFavoriteKey, vid)

	// 点赞数减一
	vidStr := strconv.FormatInt(vid, 10)
	pipeline.HIncrBy(ctx,
		rediskey.NewRedisKey(rediskey.KeyVideoHash, vidStr),
		"FavoriteCount",
		-1)

	_, err := pipeline.Exec(ctx)

	return err
}

// FavoriteVideosID 查询用户点赞的视频列表
func FavoriteVideosID(userID int64) []string {
	uidStr := strconv.FormatInt(userID, 10)
	userFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, uidStr)
	videosId, _ := rdb.ZRevRangeByScore(ctx, userFavoriteKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	return videosId
}
