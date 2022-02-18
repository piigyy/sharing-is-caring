package config

import (
	"context"
	"flag"
	"fmt"

	"github.com/piigyy/sharing-is-caring/pkg/logger"
	"github.com/spf13/viper"
)

func ReadConfigFromFile(serviceName string, config interface{}) error {
	const caller = "config.ReadConfigFromFile"
	mode := flag.String("mode", "local", "to set environment mode")
	flag.Parse()

	configFileName := fmt.Sprintf("config.%s.%s.yaml", serviceName, *mode)
	logger.Info(context.TODO(), caller, "reading config: %s", configFileName)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.SetConfigFile(configFileName)

	readCfgErr := viper.ReadInConfig()
	if readCfgErr != nil {
		logger.Error(context.TODO(), caller, "viper.ReadInConfig() return an error: %v", readCfgErr)
		return readCfgErr
	}

	unmarshallErr := viper.Unmarshal(&config)

	if unmarshallErr != nil {
		logger.Error(context.TODO(), caller, "viper.Unmarshal return an error: %v", unmarshallErr)
		return unmarshallErr
	}

	logger.Info(context.TODO(), caller, "success reading config %s", configFileName)
	return nil

}
