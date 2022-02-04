package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig() (*Config, error) {
	var cfg Config

	mode := flag.String("mode", "local", "to set environment mode")
	flag.Parse()

	configFileName := fmt.Sprintf("config.%s.yaml", *mode)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.SetConfigFile(configFileName)

	readCfgErr := viper.ReadInConfig()
	if readCfgErr != nil {
		return nil, readCfgErr
	}

	unmarshallErr := viper.Unmarshal(&cfg)

	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return &cfg, nil

}
