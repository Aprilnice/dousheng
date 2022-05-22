package mysql

// VideoInfo 视频信息表
type VideoInfo struct {
	// Id 视频唯一标识
	Id				int64	`json:"id" form:"id" gorm:"AUTO_INCREMENT;primary_key"`
	// AuthorId 作者id
	AuthorId 		int64	`json:"author_id" form:"author_id"`
	// PlayUrl 视频播放地址
	PlayUrl 		string	`json:"play_url" form:"play_url"`
	// CoverUrl 视频封面地址
	CoverUrl 		string	`json:"cover_url" form:"cover_url"`
	// FavoriteCount 视频的点赞总数
	FavoriteCount 	int64	`json:"favorite_count" form:"favorite_count"`
	// CommentCount 视频的评论总数
	CommentCount 	int64	`json:"comment_count" form:"comment_count"`
}