package inits

import (
	"bigagent_server/config"
	"bigagent_server/logger"
)

func Logger() {
	logger.InitLogger(config.CONF.System.Logfile, "info", "json", true)
}
