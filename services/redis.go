package services

import (
	"baqi/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	// 构建连接参数
	options := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	}

	// 连接 Redis
	client := redis.NewClient(&options)

	// 检查连接是否成功
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	RedisClient = client
	return client, nil
}
