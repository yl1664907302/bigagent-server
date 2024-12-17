package inits

import (
	"bigagent_server/config/global"
	"github.com/go-redis/redis/v8"
)

func RedisDB() {
	// Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	global.RedisDataConnect = client
}
