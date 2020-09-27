package apiserver

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
}

func DefaultConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}

func (config *Config) LoadFromTOML(path string) {
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		log.Fatal(err)
	}
}

func FromTOML(path string) *Config {
	config := DefaultConfig()
	config.LoadFromTOML(path)
	return config
}
