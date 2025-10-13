package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"server_port"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
	MapsAPIKey string `mapstructure:"maps_api_key"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully")
	return &cfg, nil
}
