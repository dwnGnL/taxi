package apiserver

// Config ...
type Config struct {
	BindAdress string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	//Store      *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAdress: ":8080",
		LogLevel:   "debug",
		//Store:      store.NewConfig(),
	}
}
