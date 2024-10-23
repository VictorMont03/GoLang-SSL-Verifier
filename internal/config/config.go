package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	CSVUrl            string        `mapstructure:"CSV_URL"`
	SSLCertTimeout    time.Duration `mapstructure:"SSL_CERT_TIMEOUT"`
	HTTPClientTimeout time.Duration `mapstructure:"HTTP_CLIENT_TIMEOUT"`
	LogLevel          string        `mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetDefault("CSV_URL", "https://downloads.majestic.com/majestic_million.csv")
	viper.SetDefault("SSL_CERT_TIMEOUT", 180*time.Second)
	viper.SetDefault("HTTP_CLIENT_TIMEOUT", 10*time.Second)
	viper.SetDefault("LOG_LEVEL", "INFO")

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error while loading the config %v", err)
		return config, err
	}

	return config, nil
}
