package mysqldb

import (
	"dousheng/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var gormDB *gorm.DB

// Init 初始化MySQL连接
func Init(cfg *config.MySQLConfig) error {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: cfg.DefaultStringSize,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 加前缀 t_user
			SingularTable: false, // true: 单数 user or not users  false : 表示加s
		},
	})

	if err != nil {
		return err
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 测试连通
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	gormDB = db

	return nil
}

func Migrate() error {
	if err := migrateComment(); err != nil {
		return err
	}
	if err := migrateVideoInfo(); err != nil {
		return err
	}
	if err := migrateUser(); err != nil {
		return err
	}
	return nil
}
