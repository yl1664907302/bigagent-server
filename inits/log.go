package inits

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/logger"
)

func Logger() {
	logger.InitLogger(config.CONF.System.Logfile, "info", "json", true)
}
