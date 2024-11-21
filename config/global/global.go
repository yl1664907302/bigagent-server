package global

import (
	"bigagent_server/config"
	"gorm.io/gorm"
)

var (
	CONF             *config.Server
	MysqlDataConnect *gorm.DB
)
