package ServerAPI

type Config struct {
	BindAddr string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
	}
}
