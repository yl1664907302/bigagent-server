package inits

import (
	"bigagent_server/internel/config"
	"github.com/go-redis/redis/v8"
)

func RedisDB() {
	// Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     config.CONF.System.Redisaddr,     // Redis 地址
		Password: config.CONF.System.Redispassword, // Redis 密码
		DB:       config.CONF.System.Redisdb,       // 选择数据库
	})
	config.RedisDataConnect = client
}
