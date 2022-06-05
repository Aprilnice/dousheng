package mysqldb

import (
	"dousheng/cmd/comment/service"
	"errors"
	"time"
)

type Comment struct {
	ID uint `gorm:"primarykey"`

	CreatedAt time.Time
	// 评论ID
	CommentID int64 `gorm:"not null,index"`
	// 用户ID
	UserID int64 `gorm:"not null"`
	// 视频ID
	VideoID int64 `gorm:"not null,index"`
	// 评论内容
	CommentText string `gorm:"not null"`
}

type CommentList struct {
	CreatedAt time.Time
	// 评论ID
	CommentID int64
	// 用户ID
	UserID int64
	// 评论内容
	CommentText string
}

// BindWithReq 将Req的请求数据绑定到自己的字段里
func (c *Comment) BindWithReq(req *comment.CommentRequest) error {
	if c != nil {
		c.UserID = req.UserId
		c.VideoID = req.VideoId
		c.CommentID = req.CommentId
		c.CommentText = req.CommentText
		return nil
	}
	return errors.New("model.comment: nil pointer reference")
}

// CreateComment 创建一条评论
func CreateComment(comment *Comment) error {
	return gormDB.Create(comment).Error
}

// DeleteComment 删除评论
func DeleteComment(commentID int64) error {
	return gormDB.Unscoped().Where("comment_id = ?", commentID).Delete(&Comment{}).Error
}

// QueryCommentNumsByVideID 查询某条视频的评论数
func QueryCommentNumsByVideID(videoID int64) int64 {
	return gormDB.Where(map[string]interface{}{"VideoID": videoID}).RowsAffected
}

func CommentListByVideoID(videoID int64) ([]*Comment, error) {
	var commentList []*Comment

	err := gormDB.Select([]string{"created_at", "comment_id", "user_id", "comment_text"}).Order("created_at desc").
		Where("video_id = ?", videoID).Find(&commentList).Error
	return commentList, err
}

func CommentUserIDByVideoID(videoID int64) ([]int64, error) {
	var ids []int64
	err := gormDB.Table("t_comments").Distinct().Select([]string{"user_id"}).Where("video_id = ?", videoID).Scan(&ids).Error
	return ids, err
}
