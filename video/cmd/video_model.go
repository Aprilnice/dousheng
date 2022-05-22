package cmd

import (
	"dousheng/video/pb"
	"github.com/jinzhu/gorm"
)

type (
	VideoModel interface {
		// VideoPublish 发布视频
		VideoPublish(video []byte) (resp *pb.DouyinPublishActionResponse, err error)
		// VideoFeed 视频流获取
		VideoFeed(latestTime int64) (resp *pb.DouyinFeedResponse, err error)
		// PlayVideo 视频播放(下载)
		PlayVideo(id int64) (resp *pb.PlayVideoResp, err error)
		// GetCover 封面下载
		GetCover(id int64) (resp *pb.GetCoverResp, err error)
	}

	defaultVideoModel struct {
		// db MySql数据库
		db *gorm.DB
	}
)

// NewVideoModel 初始化VideoModel
func NewVideoModel(db *gorm.DB) *defaultVideoModel {
	return &defaultVideoModel{
		db: db,
	}
}