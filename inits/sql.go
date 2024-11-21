package inits

import (
	"bigagent_server/config/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1",
		3306,
		"bigagent")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}
	global.MysqlDataConnect = db
}

//func MysqlDB() {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		global.CONF.System.Database.MysqlUser,
//		global.CONF.System.Database.MysqlPassword,
//		global.CONF.System.Database.MysqlHost,
//		global.CONF.System.Database.MysqlPort,
//		global.CONF.System.Database.MysqlDatabasename)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Print(err)
//	}
//	global.MysqlDataConnect = db
//}
