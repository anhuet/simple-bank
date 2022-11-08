package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBdriver       string        `mapstructure:"DB_DRIVER"`
	DBSource       string        `mapstructure:"DB_SOURCE"`
	ServerAdress   string        `mapstructure:"SERVER_ADDRESS"`
	TokenSecretKey string        `mapstructure:"TOKEN_SECRET_KEY"`
	TokenDuration  time.Duration `mapstructure:"TOKEN_SECRET_DURATION"`
}

func LoadConfigFile(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
