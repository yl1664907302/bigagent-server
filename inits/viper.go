package inits

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/logger"
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
	err = v.Unmarshal(&config.CONF)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
	}
}
