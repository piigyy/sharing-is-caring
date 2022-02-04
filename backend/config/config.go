package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadConfig() (*Config, error) {
	var cfg Config

	mode := flag.String("mode", "local", "to set environment mode")
	flag.Parse()

	configFileName := fmt.Sprintf("config.%s.yaml", *mode)
	log.Printf("reading config: %s\n", configFileName)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.SetConfigFile(configFileName)

	readCfgErr := viper.ReadInConfig()
	if readCfgErr != nil {
		log.Printf("error viper.ReadInConfig: %v\n", readCfgErr)
		return nil, readCfgErr
	}

	unmarshallErr := viper.Unmarshal(&cfg)

	if unmarshallErr != nil {
		log.Printf("error viper.Unmarshal: %v\n", readCfgErr)
		return nil, unmarshallErr
	}

	log.Printf("success reading config: %s\n", configFileName)
	return &cfg, nil

}
