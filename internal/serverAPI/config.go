package serverAPI

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	serverPort string
	dbURL      string
	dbDriver   string
}

func NewConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main_config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		serverPort: viper.GetString("server_port"),
		dbURL:      viper.GetString("db_url"),
		dbDriver:   viper.GetString("db_driver_name"),
	}
}
