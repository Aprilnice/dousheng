package mysql

import (
	"dousheng/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Init 初始化MySQL连接
func Init(cfg *config.MySQLConfig, server string) (db *gorm.DB, err error) {
	var tmp *gorm.DB
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	tmp, err = gorm.Open("mysql", dsn)
	if err != nil {
		return tmp,err
	}

	// video服务需要建的表
	if server == "video" {
		tmp.AutoMigrate(&VideoInfo{})
	}

	return tmp,nil
}
