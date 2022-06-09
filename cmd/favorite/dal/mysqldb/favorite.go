package mysqldb

import (
	favorite "dousheng/cmd/favorite/service"
	userDB "dousheng/cmd/user/dal/mysqldb"
	videoDB "dousheng/cmd/video/dal/mysqldb"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Favorite struct {
	gorm.Model

	// 用户ID
	UserID int64 `gorm:"not null,index"`
	// 视频ID
	VideoID int64 `gorm:"not null,index"`
}

// BindWithReq 将Req的请求数据绑定到自己的字段里
func (f *Favorite) BindWithReq(req *favorite.FavoriteActionRequest) {

	f.UserID = req.UserId
	f.VideoID = req.VideoId
}

// CreateFavorite 创建一条评论
func CreateFavorite(favorite *Favorite) error {
	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 如果存在 更新字段即可
	if IsFavorite(favorite.UserID, favorite.VideoID) {

		if err := tx.Unscoped().Model(&Favorite{}).
			Where("user_id = ? and video_id = ?", favorite.UserID, favorite.VideoID).
			Update("deleted_at", nil).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Create(favorite).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新视频表
	if err := tx.Model(&videoDB.VideoInfo{}).Where("id = ?", favorite.VideoID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// CancelFavorite 取消点赞
func CancelFavorite(favorite *Favorite) error {
	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("user_id = ? and video_id = ?", favorite.UserID, favorite.VideoID).
		Delete(&Favorite{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新视频表
	if err := tx.Model(&videoDB.VideoInfo{}).Where("id = ?", favorite.VideoID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func IsFavorite(userID, videoID int64) bool {
	var dest Favorite
	if err := gormDB.Unscoped().Where("user_id = ? and video_id = ?", userID, videoID).
		First(&dest).Error; err != nil {
		fmt.Println(err)
		return false
	}
	if dest.UserID == 0 || dest.VideoID == 0 {
		return false
	}
	return true
}

// QueryVideosInfo 查询视频信息
func QueryVideosInfo(videoIds []string, ch chan<- int) ([]*videoDB.VideoInfo, error) {
	var videosInfo []*videoDB.VideoInfo
	//err := gormDB.Table("t_video_infos").Where("id in ?", videoIds).Find(&videosInfo).Clauses().Error
	err := gormDB.Debug().Where("id in ?", videoIds).Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "FIELD(id,?)",
			Vars:               []interface{}{videoIds},
			WithoutParentheses: true},
	}).Find(&videosInfo).Error

	defer func() {
		ch <- 1
		if err != nil {
			ch <- -1
		}
	}()

	return videosInfo, err
}

func QueryAuthorsInfo(videoIds []string, ch chan<- int) ([]*userDB.UserInfo, error) {
	var userIds []int64
	err := gormDB.Table("t_video_infos").Select([]string{"author_id"}).
		Where("id in ?", videoIds).Scan(&userIds).Error
	if err != nil || userIds == nil {
		ch <- -1
		return nil, err
	}
	var usersInfo []*userDB.UserInfo
	err = gormDB.Table("t_user_infos").Where("user_id in ?", userIds).Find(&usersInfo).Error
	defer func() {
		ch <- 1
		if err != nil || usersInfo == nil {
			ch <- -1
		}
	}()

	return usersInfo, err
}

func FavoriteVideosID(userID int64) ([]int64, error) {
	var videoIds []int64
	err := gormDB.Model(&Favorite{}).Distinct().
		Select("video_id").Where("user_id = ?", userID).
		Order("updated_at desc").Scan(&videoIds).Error
	return videoIds, err
}
