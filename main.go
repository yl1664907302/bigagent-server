package main

import (
	"bigagent_server/config/global"
	"bigagent_server/inits"
)

func init() {
	inits.Viper()
	inits.RunG()
	r := inits.Router()
	panic(r.Run(global.CONF.System.Addr))
}

func main() {}
