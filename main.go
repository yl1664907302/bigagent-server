package main

import (
	"bigagent_server/config/global"
	_ "bigagent_server/docs"
	"bigagent_server/inits"
)

func init() {
	inits.Viper()
	inits.Logger()
	inits.MysqlDB()
	inits.RedisDB()
}

// @title BigAgent API
// @version 1.0
// @description This is a BigAgent API server.
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	inits.CronTask()
	inits.RunG()
	r := inits.Router()
	panic(r.Run(global.CONF.System.Addr))
}
