package mysqldb

import (
	relation "dousheng/cmd/relation/service"
	userDB "dousheng/cmd/user/dal/mysqldb"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Relation struct {
	gorm.Model
	//用户本身UserId
	UserId int64 `gorm:"not null"`
	//对象用户UserId
	ToUserId int64 `gorm:"not null"`
}

func UsersInfo(userIDs []string) ([]*relation.User, error) {
	var usersInfo []*userDB.UserInfo
	err := gormDB.Model(&userDB.UserInfo{}).Debug().Where("user_id in ?", userIDs).Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "FIELD(user_id,?)",
			Vars:               []interface{}{userIDs},
			WithoutParentheses: true},
	}).Find(&usersInfo).Error

	users := make([]*relation.User, len(usersInfo))
	for i, u := range usersInfo {
		tmp := relation.User{
			Id:            u.UserId,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      true,
		}
		users[i] = &tmp
	}

	return users, err
}

func DoFollow(userID, toUserID int64) error {

	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	relationModel := Relation{
		UserId:   userID,
		ToUserId: toUserID,
	}
	if IsFollow(userID, toUserID) {
		if err := tx.Unscoped().Model(&Relation{}).
			Where("user_id = ? and to_user_id = ?", userID, toUserID).
			Update("deleted_at", nil).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 更新关系表
		if err := tx.Create(&relationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新用户表
	if err := tx.Model(&userDB.UserInfo{}).Where("user_id = ?", userID).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&userDB.UserInfo{}).Where("user_id = ?", toUserID).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func CancelFollow(userID, toUserID int64) error {

	tx := gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新关系表 采用软删除
	if err := tx.Where("user_id = ? and to_user_id = ?", userID, toUserID).
		Delete(&Relation{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新用户表
	if err := tx.Model(&userDB.UserInfo{}).Where("user_id = ?", userID).
		UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&userDB.UserInfo{}).Where("user_id = ?", toUserID).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func IsFollow(userID, toUserID int64) bool {
	var dest Relation
	// 查不到会返回 record not found 错误
	if err := gormDB.Model(&Relation{}).Unscoped().Where("user_id = ? and to_user_id = ?", userID, toUserID).
		First(&dest).Error; err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("dest: ", dest)
	return true
}
