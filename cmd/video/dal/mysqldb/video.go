package mysqldb

import (
	userInfo "dousheng/cmd/user/dal/mysqldb"
)

// VideoInfo 视频信息表
type VideoInfo struct {
	// Id 视频唯一标识
	Id int64 `json:"id" form:"id" gorm:"unsigned;primary_key"`
	// Title 视频标题
	Title string `json:"title" form:"title"`
	// AuthorId 作者id
	AuthorId int64 `json:"author_id" form:"author_id"`
	// PlayUrl 视频播放地址
	PlayUrl string `json:"play_url" form:"play_url"`
	// CoverUrl 视频封面地址
	CoverUrl string `json:"cover_url" form:"cover_url"`
	// FavoriteCount 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count" form:"favorite_count"`
	// CommentCount 视频的评论总数
	CommentCount int64 `json:"comment_count" form:"comment_count"`
	// PublishTime 上传的时间戳
	PublishTime int64 `json:"publish_time" form:"publish_time"`
}

// PublishVideo 发布一个视频
func PublishVideo(video *VideoInfo) error {
	return gormDB.Create(video).Error
}

// GetVideoFeed 获取feed流
func GetVideoFeed(latestTime int64) (videos []VideoInfo, err error) {
	err = gormDB.Table("t_video_infos").Where("publish_time <= ?",latestTime).Order("publish_time desc").Limit(30).Find(&videos).Error
	return videos, err
}

// GetUserInfo 获取用户信息
func GetUserInfo(userId int64) (user userInfo.UserInfo, err error) {
	err = gormDB.Table("t_user_infos").Where("user_id = ?", userId).Find(&user).Error
	return user, err
}
