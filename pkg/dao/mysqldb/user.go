package mysqldb

import (
	"dousheng/user/service"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//用户名
	Username string `gorm:"not null"`
	//密码
	Password string `gorm:"not null"`
	//用户id
	UserID int64 `gorm:"not null,primary_key"`
}

func migrateUser() error {

	migrator := gormDB.Migrator()
	if migrator.HasTable(&User{}) {
		return nil
	}
	return migrator.CreateTable(&User{})
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
	return gormDB.Create(user).Error
}

func UserLogin(req *service.DouyinUserLoginRequest) bool {
	username := req.GetUsername()
	password := req.GetPassword()
	var user User
	err := gormDB.Where("username = ?", username).Find(&user)
	if err != nil {
		return false
	}
	return user.Password == password
}

func GetUserId(name string) int64 {
	var user User
	gormDB.Where("username = ?", name).Find(&user)
	return user.UserID
}

// UserInfo 用户信息
type UserInfo struct {
}
