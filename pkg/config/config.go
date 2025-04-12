package config

import "os"

type Config struct {
	ServerPort string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: os.Getenv("BACKEND_PORT"),
	}
}
