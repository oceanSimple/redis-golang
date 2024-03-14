package viper

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server-v5/log"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./viper")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.SystemLog.Fatal("Failed to read the configuration file.",
			zap.Error(err))
	}
}
