package services

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis(ctx context.Context, addr, password string, db int) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	return err
}

func Redis() *redis.Client {
	return rdb
}
