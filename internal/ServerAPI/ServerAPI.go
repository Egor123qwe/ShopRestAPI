package ServerAPI

type ServerApi struct {
	config *Config
}

func New(config *Config) *ServerApi {
	return &ServerApi{
		config: config,
	}
}

func (s *ServerApi) Start() error {
	return nil
}
