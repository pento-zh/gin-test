package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var redisConfig *redis.Options = &redis.Options{
	Addr:         "192.168.1.13:6379", // Redis服务器地址
	Password:     "RT4l8pcR6ed6",      // 密码
	DB:           0,                   // 默认数据库
	PoolSize:     10,                  // 连接池大小
	MinIdleConns: 3,                   // 最小空闲连接
}

// 初始化Redis连接

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		// 添加连接池配置
		redisClient = redis.NewClient(redisConfig)

		// 检查Redis连接
		ctx := context.Background()
		if err := redisClient.Ping(ctx).Err(); err != nil {
			panic(err)
		}
	}
	return redisClient
}
