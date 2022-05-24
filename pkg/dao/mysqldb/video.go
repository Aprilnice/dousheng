package mysqldb

// VideoInfo 视频信息表
type VideoInfo struct {
	// Id 视频唯一标识
	Id				int64	`form:"id" gorm:"unsigned;primary_key"`
	// Title 视频标题
	Title 		    string	`form:"title"`
	// AuthorId 作者id
	AuthorId 		int64	`form:"author_id"`
	// PlayUrl 视频播放地址
	PlayUrl 		string	`form:"play_url"`
	// CoverUrl 视频封面地址
	CoverUrl 		string	`form:"cover_url"`
	// FavoriteCount 视频的点赞总数
	FavoriteCount 	int64	`form:"favorite_count"`
	// CommentCount 视频的评论总数
	CommentCount 	int64	`form:"comment_count"`
}

// migrateVideoInfo 迁移视频信息表
func migrateVideoInfo() error {
	migrator := gormDB.Migrator()
	if migrator.HasTable(&VideoInfo{}) {
		return nil
	}
	return migrator.CreateTable(&VideoInfo{})
}

// PublishVideo 发布一个视频
func PublishVideo(video *VideoInfo) error {
	return gormDB.Create(video).Error
}