package server

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort string
	DbURL      string
	DbDriver   string
	SessionKey string
}

func NewConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main_config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		ServerPort: viper.GetString("server_port"),
		DbURL:      viper.GetString("db_url"),
		DbDriver:   viper.GetString("db_driver_name"),
		SessionKey: viper.GetString("session_key"),
	}
}
