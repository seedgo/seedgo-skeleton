package config

import (
	"github.com/spf13/viper"
	"seedgo-skeleton/common"
)

func Init() {
	loadConfig()
	common.ParseSystemConfig()
	common.InitLogger()
	if common.ServerConfig.Debug {
		common.Logger.Info(viper.AllSettings())
	}

	common.ParseDatabaseConfig()
	common.ParseRedisConf()
}

func loadConfig() {
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
