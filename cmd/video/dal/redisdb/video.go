package redisdb

import (
	userInfo "dousheng/cmd/user/dal/mysqldb"
	"dousheng/cmd/video/dal/mysqldb"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// AddVideoInfo redis中使用Hash存储视频信息
func AddVideoInfo(video *mysqldb.VideoInfo) error {
	return 	rdb.HSet(
		ctx,
		// key值为视频id
		strconv.FormatInt(video.Id, 10),
		// field 为 Title
		"Title",
		video.Title,
		// field 为 AuthorId
		"AuthorId",
		video.AuthorId,
		// filed 为 PlayUrl
		"PlayUrl",
		video.PlayUrl,
		// field 为 CoverUrl
		"CoverUrl",
		video.CoverUrl,
		// field 为 FavoriteCount
		"FavoriteCount",
		video.FavoriteCount,
		// field 为 CommentCount
		"CommentCount",
		video.CommentCount,
		// field 为 PublishTime
		"PublishTime",
		video.PublishTime,
	).Err()
}

// AddVideoId 在有序集合feed中添加视频id
func AddVideoId(videoId int64, publishTime int64) error {
	return rdb.ZAdd(ctx,
		// 集合为VideoFeed
		"VideoFeed",
		&redis.Z{
			// 分数为发布时间,用来排序
			Score: float64(publishTime),
			// 成员名为视频id,用于查找
			Member: videoId,
		},
	).Err()
}

// GetFeed 从redis获取视频列表
func GetFeed(latestTime int64) (videos []mysqldb.VideoInfo, err error) {
	// 获取视频列表中视频的id
	videoList, err := rdb.ZRevRangeByScore(
		ctx,
		"videoFeed",
		&redis.ZRangeBy{
			Min: "-inf",
			Max: strconv.FormatInt(latestTime, 10),
		},
	).Result()

	// 一次最多返回30个视频
	num := 30
	if len(videoList) < 30 {
		num = len(videoList)
	}

	// 获取具体视频信息
	var tmp mysqldb.VideoInfo
	for i := 0; i < num; i++ {
		// 获取视频信息
		videosInfo, err := rdb.HGetAll(ctx, videoList[i]).Result()
		if err != nil {
			return videos,err
		}

		// 提取信息
		tmp.Id, _ = strconv.ParseInt(videoList[i], 10, 64)
		tmp.AuthorId, _ = strconv.ParseInt(videosInfo["AuthorId"], 10, 64)
		tmp.PlayUrl = videosInfo["PlayUrl"]
		tmp.CoverUrl = videosInfo["CoverUrl"]
		tmp.FavoriteCount, _ = strconv.ParseInt(videosInfo["FavoriteCount"], 10, 64)
		tmp.CommentCount, _ = strconv.ParseInt(videosInfo["CommentCount"], 10, 64)
		tmp.PublishTime, _ = strconv.ParseInt(videosInfo["PublishTime"], 10, 64)

		videos = append(videos, tmp)
	}
	return videos, nil
}

// GetAuthorInfo redis获取作者信息
func GetAuthorInfo(authorId int64) (user userInfo.UserInfo, err error) {
	return user, nil
}