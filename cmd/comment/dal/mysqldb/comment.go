package mysqldb

import (
	"dousheng/cmd/comment/service"
	userDB "dousheng/cmd/user/dal/mysqldb"
	videoDB "dousheng/cmd/video/dal/mysqldb"
	"errors"
	"gorm.io/gorm"
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

//type CommentList struct {
//	CreatedAt time.Time
//	// 评论ID
//	CommentID int64
//	// 用户ID
//	UserID int64
//	// 评论内容
//	CommentText string
//}

// BindWithReq 将Req的请求数据绑定到自己的字段里
func (c *Comment) BindWithReq(req *comment.CommentActionRequest) error {
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
	// 开启事务
	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 更新评论表
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新视频表
	if err := tx.Model(&videoDB.VideoInfo{}).Where("id = ?", comment.VideoID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// DeleteComment 删除评论
func DeleteComment(commentID, videoID int64) error {
	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 删除评论表
	if err := tx.Unscoped().Where("comment_id = ?", commentID).Delete(&Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 视频表评论数减一
	if err := tx.Model(&videoDB.VideoInfo{}).Where("id = ?", videoID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func VideoCommentList(videoID int64) ([]*Comment, error) {
	var commentList []*Comment
	err := gormDB.Select([]string{"created_at", "comment_id", "user_id", "comment_text"}).Order("created_at desc").
		Where("video_id = ?", videoID).Find(&commentList).Error

	return commentList, err
}

func CommentUserInfoByVideoID(videoID int64) ([]*userDB.UserInfo, error) {

	var ids []int64
	if err := gormDB.Table("t_comments").Distinct().
		Select([]string{"user_id"}).Where("video_id = ?", videoID).Scan(&ids).Error; err != nil {
		return nil, err
	}

	var usersInfo []*userDB.UserInfo
	err := gormDB.Table("t_user_infos").Where("user_id in ?", ids).Find(&usersInfo).Error

	return usersInfo, err
}

// CommentUserInfo 评论用户信息
func CommentUserInfo(userId int64) (*userDB.UserInfo, error) {
	var user *userDB.UserInfo
	err := gormDB.Debug().Model(userDB.UserInfo{}).Where("user_id = ?", userId).Find(&user).Error

	return user, err
}
