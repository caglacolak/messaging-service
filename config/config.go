// config.go
package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Redis Redis `mapstructure:"redis"`
	Postgres Postgres `mapstructure:"postgres"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}


type Redis struct {
	URI string `mapstructure:"uri"`
}

type Postgres struct {
	Environment struct {
		User     string `mapstructure:"POSTGRES_USER"`
		Password string `mapstructure:"POSTGRES_PASSWORD"`
		DB       string `mapstructure:"POSTGRES_DB"`
	}
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}