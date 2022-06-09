package redisdb

import (
	favorite "dousheng/cmd/favorite/service"
	"dousheng/pkg/rediskey"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const (
	FavoriteValue = 432 // 指定一个赞成票是432分  432*200 = 86400秒 = 一天 也就是说获得200张票可以将视频续一天
)

// CreateFavorite 在redis中新增一条点赞记录
func CreateFavorite(favorite *favorite.FavoriteActionRequest) error {
	pipeline := rdb.TxPipeline()
	uid := favorite.UserId
	vid := favorite.VideoId
	uidStr := strconv.FormatInt(uid, 10)
	vidStr := strconv.FormatInt(vid, 10)
	// 增加用户点赞记录
	userFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, uidStr)

	pipeline.ZAdd(ctx, userFavoriteKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: vid,
	})
	videoFavoriteKey := "VideoFeed"
	// 为该视频的发布时间毫秒数增加 423
	pipeline.ZIncrBy(ctx, videoFavoriteKey, FavoriteValue, vidStr)
	_, err := pipeline.Exec(ctx)
	return err
}

// CancelFavorite 取消点赞
func CancelFavorite(favorite *favorite.FavoriteActionRequest) error {
	pipeline := rdb.TxPipeline()
	uid := favorite.UserId
	vid := favorite.VideoId
	uidStr := strconv.FormatInt(uid, 10)
	vidStr := strconv.FormatInt(vid, 10)
	// 取消用户点赞记录
	userFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, uidStr)
	pipeline.ZRem(ctx, userFavoriteKey, vid)

	videoFavoriteKey := rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, vidStr)
	// 为该视频记录点赞用户
	// 为该视频的发布时间毫秒数减少 423
	pipeline.ZIncrBy(ctx, videoFavoriteKey, -1*FavoriteValue, vidStr)

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
