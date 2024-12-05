package global

import (
	"bigagent_server/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	CONF             *config.Server
	MysqlDataConnect *gorm.DB
	RedisDataConnect *redis.Client
)
