package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Port int
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
		return nil, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
		return nil, err
	}

	return &config, nil
}
