package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	ServerAddress       string        `mapstructure:"DB_ADDRESS"`
	DBUserName          string        `mapstructure:"DATABASE_USERNAME"`
	DBPassword          string        `mapstructure:"DATABASE_PASSWORD"`
	DBName              string        `mapstructure:"DATABASE_NAME"`
	DBHost              string        `mapstructure:"DATABASE_HOST"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	return
}
