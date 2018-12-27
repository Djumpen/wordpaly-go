package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB
	Version     string
	Port        int
	ReleaseMode bool
}

type DB struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

func ReadConfig() *Config {
	var cfg Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &cfg
}
