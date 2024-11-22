package main

import (
	"bigagent_server/config/global"
	"bigagent_server/inits"
)

func init() {
	inits.Viper()
	inits.Logger()
	inits.MysqlDB()

}

func main() {
	inits.CronTask()
	inits.RunG()
	r := inits.Router()
	panic(r.Run(global.CONF.System.Addr))
}
