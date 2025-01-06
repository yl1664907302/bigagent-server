package inits

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/logger"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func RedisDB() {
	// Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     config.CONF.System.Redisaddr,     // Redis 地址
		Password: config.CONF.System.Redispassword, // Redis 密码
		DB:       config.CONF.System.Redisdb,       // 选择数据库
	})
	config.RedisDataConnect = client
	err := client.Ping(context.Background()).Err()
	if err != nil {
		logger.DefaultLogger.Errorf("redis连接失败: %v", err)
		panic(err)
	}
}
