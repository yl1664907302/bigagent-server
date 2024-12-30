package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	CONF             *Server
	MysqlDataConnect *gorm.DB
	RedisDataConnect *redis.Client
)
