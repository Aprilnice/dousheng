package redisdb

import (
	"dousheng/pkg/rediskey"
	"strconv"
)

// CreateComment 增加评论数
func CreateComment(videoID int64) error {
	pipeline := rdb.TxPipeline()

	vidStr := strconv.FormatInt(videoID, 10)
	pipeline.HIncrBy(ctx,
		rediskey.NewRedisKey(rediskey.KeyVideoHash, vidStr),
		"CommentCount",
		1)
	_, err := pipeline.Exec(ctx)
	return err
}

// DeleteComment 删除评论
func DeleteComment(videoID int64) error {
	pipeline := rdb.TxPipeline()
	vidStr := strconv.FormatInt(videoID, 10)
	pipeline.HIncrBy(ctx,
		rediskey.NewRedisKey(rediskey.KeyVideoHash, vidStr),
		"CommentCount",
		-1)
	_, err := pipeline.Exec(ctx)
	return err
}
