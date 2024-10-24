package inits

import (
	"bigagent_server/config/global"
	"bigagent_server/utils/logger"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
	}
}
