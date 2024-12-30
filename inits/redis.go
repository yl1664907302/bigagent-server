package inits

import (
	"bigagent_server/internel/config"
	"github.com/go-redis/redis/v8"
)

func RedisDB() {
	// Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	config.RedisDataConnect = client
}
