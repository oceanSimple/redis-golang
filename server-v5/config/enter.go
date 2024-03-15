package viper

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server-v5/log"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.SystemLog.Fatal("Failed to read the configuration file.",
			zap.Error(err))
	}
}
