package mysqldb

import (
	"dousheng/user/service"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//用户名
	Username string `gorm:"not null,primary_key"`
	//密码
	Password string `gorm:"not null"`
	//用户id
	UserID int64 `gorm:"not null"`
}

type UserInfo struct {
	gorm.Model
	//用户UserID
	UserId int64 `gorm:"not null"`
	//用户名
	Name string `gorm:"not null"`
	//关注总数
	FollowCount int64 `gorm:"not null"`
	//粉丝总数
	FollowerCount int64 `gorm:"not null"`
	//是否关注
	IsFollow bool `gorm:"not null"`
}

func migrateUser() error {

	migrator := gormDB.Migrator()
	if migrator.HasTable(&User{}) {
		return nil
	}
	return migrator.CreateTable(&User{})
}
func migrateUserInfo() error {

	migrators := gormDB.Migrator()
	if migrators.HasTable(&UserInfo{}) {
		return nil
	}
	return migrators.CreateTable(&UserInfo{})
}

// BindWithReq 将Req的请求数据绑定到自己的字段里
func (u *User) BindWithReq(req *service.DouyinUserRegisterRequest) error {
	if u != nil {
		u.Username = req.GetUsername()
		u.Password = req.GetPassword()
		return nil
	}

	return errors.New("model.user: nil pointer reference")
}

//注册用户
func UserRegister(user *User) error {
	gormDB.Where("username = ?", user.Username).First(&user)
	if user.ID != 0 {
		return errors.New("same username")
	}
	return gormDB.Create(user).Error

}

func UserLogin(req *service.DouyinUserLoginRequest) bool {
	username := req.GetUsername()
	password := req.GetPassword()
	var user User
	gormDB.Where("username = ?", username).First(&user)
	return user.Password == password
}

func GetUserId(name string) int64 {
	var user User
	gormDB.Where("username = ?", name).Find(&user)
	return user.UserID
}

func SetUserInfo(ui *UserInfo) error {
	return gormDB.Create(ui).Error
}

func GetUserInfoById(userId int64) (*UserInfo, error) {
	var userInfo *UserInfo
	err := gormDB.Where("user_id = ?", userId).Find(&userInfo).Error
	if err != nil {
		return &UserInfo{}, err
	}
	return userInfo, nil
}
