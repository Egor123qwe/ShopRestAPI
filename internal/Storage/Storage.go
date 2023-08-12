package Storage

type Store struct {
	config *Config
}

//func New(config *Config) *Config {/
//	return &Config{
//		config: config,
//	}
//}

func (*Store) Open() error {
	return nil
}

func (*Store) Close() {

}
