package apiserver

// Config of api server, should be read from .toml file
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
}

// NewConfig with default args
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
