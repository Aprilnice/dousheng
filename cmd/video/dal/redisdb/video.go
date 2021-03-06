package redisdb

import (
	userInfo "dousheng/cmd/user/dal/mysqldb"
	"dousheng/cmd/video/dal/mysqldb"
	"dousheng/pkg/rediskey"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// AddVideoInfo redis中使用Hash存储视频信息
func AddVideoInfo(video *mysqldb.VideoInfo) error {
	return rdb.HSet(
		ctx,
		// key值为视频id
		rediskey.NewRedisKey(rediskey.KeyVideoHash, strconv.FormatInt(video.Id, 10)),
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
		rediskey.NewRedisKey(rediskey.KeyFeedZSet),
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
		rediskey.NewRedisKey(rediskey.KeyFeedZSet),
		&redis.ZRangeBy{
			Min:    "-inf",
			Max:    strconv.FormatInt(latestTime, 10),
			Offset: 0,
			Count:  30,
		},
	).Result()
	if err != nil {
		return videos, nil
	}

	// 一次最多返回30个视频
	num := 30
	if len(videoList) < 30 {
		num = len(videoList)
	}

	// 获取具体视频信息
	var tmp mysqldb.VideoInfo
	for i := 0; i < num; i++ {
		// 获取视频信息
		videosInfo, err := rdb.HGetAll(
			ctx,
			rediskey.NewRedisKey(rediskey.KeyVideoHash, videoList[i]),
		).Result()
		if err != nil {
			return videos, err
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

// JudgeUser 判断作者信息是否在redis中
func JudgeUser(userId int64) bool {
	result, _ := rdb.HGetAll(
		ctx,
		rediskey.NewRedisKey(rediskey.KeyUserHash, strconv.FormatInt(userId, 10)),
	).Result()
	if len(result) == 0 {
		return false
	} else {
		return true
	}
}

// AddUserInfo 向redis中添加用户信息
func AddUserInfo(user userInfo.UserInfo) error {
	return rdb.HSet(
		ctx,
		// key值为用户id
		rediskey.NewRedisKey(rediskey.KeyUserHash, strconv.FormatInt(user.UserId, 10)),
		// field 为 Name
		"Name",
		user.Name,
		// field 为 FollowCount
		"FollowCount",
		user.FollowCount,
		// field 为 FollowerCount
		"FollowerCount",
		user.FollowerCount,
	).Err()
}

// GetAuthorInfo redis获取作者信息
func GetAuthorInfo(authorId int64) (user userInfo.UserInfo, err error) {
	info, err := rdb.HGetAll(
		ctx,
		rediskey.NewRedisKey(rediskey.KeyUserHash, strconv.FormatInt(authorId, 10)),
	).Result()
	if len(info) <= 0 {
		return
	}
	user.UserId = authorId
	user.Name = info["Name"]
	user.FollowCount, _ = strconv.ParseInt(info["FollowCount"], 10, 64)
	user.FollowerCount, _ = strconv.ParseInt(info["FollowerCount"], 10, 64)
	return user, nil
}

// GetLike 获取是否点赞
func GetLike(userId int64, videoId int64) bool {
	_, err := rdb.ZRank(
		ctx,
		rediskey.NewRedisKey(rediskey.KeyFavoriteZSet, strconv.FormatInt(userId, 10)),
		strconv.FormatInt(videoId, 10),
	).Result()
	// key不存在err = redis: nil
	if err != nil {
		return false
	} else {
		return true
	}
}

// GetFollow 获取是否关注
func GetFollow(userId int64, authorId int64) bool {
	if _, err := rdb.ZRank( // key不存在err = redis: nil
		ctx,
		rediskey.NewRedisKey(rediskey.KeyUserFollow, strconv.FormatInt(userId, 10)),
		strconv.FormatInt(authorId, 10),
	).Result(); err != nil {
		return false
	}
	return true

}

// AddVideoToUser 视频信息存入用户信息便于获取发布列表
func AddVideoToUser(userId int64, videoId int64) error {
	return rdb.HSet(
		ctx,
		// key值为用户id
		rediskey.NewRedisKey(rediskey.KeyUserHash, strconv.FormatInt(userId, 10)),
		// field 为 video
		"videoId",
		// value为空
		videoId,
	).Err()
}

// GetVideoList 获取用户视频列表
func GetVideoList(userId int64) (videos []mysqldb.VideoInfo, err error) {
	// 获取所有视频id
	videoList, err := rdb.HGetAll(
		ctx,
		rediskey.NewRedisKey(rediskey.KeyUserHash, strconv.FormatInt(userId, 10)),
	).Result()

	// 获取具体视频信息
	var tmp mysqldb.VideoInfo
	for i := 0; i < len(videoList)-3; i++ {
		// 获取视频信息
		videosInfo, err := rdb.HGetAll(
			ctx,
			rediskey.NewRedisKey(rediskey.KeyVideoHash, videoList["videoId"]),
		).Result()
		if err != nil {
			return videos, err
		}

		// 提取信息
		tmp.Id, _ = strconv.ParseInt(videoList["videoId"], 10, 64)
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
