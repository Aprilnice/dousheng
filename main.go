package main

import (
	"dousheng/comment"
	"dousheng/config"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/snowflaker"
	"dousheng/user"
	"dousheng/video"
	"log"
)

func main() {

	// 初始化配置
	config.Init("./config")

	// 初始化ID生成器
	if err := snowflaker.Init(config.ConfInstance().StartTime, config.ConfInstance().MachineID); err != nil {
		log.Println("ID 生成器初始化失败")
		log.Fatal(err)
	}

	// 初始化数据库
	if err := mysqldb.Init(config.ConfInstance().MySQLConfig); err != nil {
		log.Println("mysql数据库初始化失败")
		log.Fatal(err)
	}

	// 迁移数据库实例
	if err := mysqldb.Migrate(); err != nil {
		log.Println("数据库迁移失败")
		log.Fatal(err)
	}

	comment.Run()
	user.Run()
	video.VideoRun()

}
