package inits

import (
	"bigagent_server/config/global"
	"bigagent_server/utils/logger"
)

func Logger() {
	logger.InitLogger(global.CONF.System.Logfile, "info", "json", true)
}
