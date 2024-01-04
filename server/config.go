package server

import (
	"fmt"

	"github.com/thynquest/fizzbuzz/pkg/logging"

	"github.com/spf13/viper"
)

var configTitle = "[fizzbuzz-api-config]"

type Config struct {
	HOST string `mapstructure:HOST`
	PORT string `maprstructure:PORT`
}

func LoadConfig(filePath string) (*Config, error) {
	logging.Info(configTitle, "loading configuration...")
	viper.SetConfigFile(filePath)
	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		logging.Error(configTitle, fmt.Sprintf("error when loading configuration from file: %v", errConfig))
		return &Config{}, errConfig
	}
	var config Config
	errUnMarshal := viper.Unmarshal(&config)
	if errUnMarshal != nil {
		logging.Error(configTitle, fmt.Sprintf("error when unmarshaling configuration: %v", errUnMarshal))
		return &Config{}, errUnMarshal
	}
	return &config, nil
}
