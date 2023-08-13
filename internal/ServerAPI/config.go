package ServerAPI

import (
	"ShopRestAPI/internal/Storage"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	serverPort string
	store      *Storage.Config
}

func NewConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main_config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		serverPort: viper.GetString("server_port"),
		store:      Storage.NewConfig(),
	}
}
