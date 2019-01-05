package config

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
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

func ReadConfig() (*Config, error) {
	var cfg Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return &cfg, nil
}

func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
