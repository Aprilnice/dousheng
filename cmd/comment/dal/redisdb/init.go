package redisdb

import (
	"context"
	"dousheng/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func Init(cfg *config.RedisConfig) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping(ctx).Result()

	return
}

func CLose() {
	_ = rdb.Close()
}
