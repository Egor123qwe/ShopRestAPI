package Storage

import "github.com/spf13/viper"

type Config struct {
	dbURL    string
	dbDriver string
}

func NewConfig() *Config {
	return &Config{
		dbURL:    viper.GetString("db_url"),
		dbDriver: viper.GetString("db_driver_name"),
	}
}
