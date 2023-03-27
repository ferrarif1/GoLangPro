package config

import (
	"github.com/spf13/viper"
)

var AppConfig = viper.New()

func InitConfig(cfgFile string) {

	AppConfig.SetConfigFile(cfgFile)

	if err := AppConfig.ReadInConfig(); err != nil {
		panic("Failed to read config file")
	}
}
