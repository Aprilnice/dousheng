package mysqldb

import (
	"dousheng/comment/service"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	// 评论ID
	CommentID int64 `gorm:"not null"`
	// 用户ID
	UserID int64 `gorm:"not null"`
	// 视频ID
	VideoID int64 `gorm:"not null"`
	// 评论内容
	CommentText string `gorm:"not null"`
}

// MigrateComment 迁移评论表
func migrateComment() error {

	migrator := gormDB.Migrator()
	if migrator.HasTable(&Comment{}) {
		return nil
	}
	return migrator.CreateTable(&Comment{})
}

// BindWithReq 将Req的请求数据绑定到自己的字段里
func (c *Comment) BindWithReq(req *service.CommentRequest) error {
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
	fmt.Println(commentID)
	return gormDB.Unscoped().Where("comment_id = ?", commentID).Delete(&Comment{}).Error
}
