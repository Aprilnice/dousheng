package redisdb

import (
	"context"
	"dousheng/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	// ctx 操作redis所需的上下文
	ctx = context.Background()
	// rdb redis客户端实例
	rdb *redis.Client
)

// InitRedisClient 初始化redis客户端
func InitRedisClient(cfg *config.RedisConfig) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
