package config

import "github.com/spf13/viper"

func Init() error {
	viper.AddConfigPath("./internal/config")
	viper.AddConfigPath("./internal/s3")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
