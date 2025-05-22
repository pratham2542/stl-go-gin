package config

import (
	"os"
)

type ServerConfig struct {
	port string
}

func NewServerConfig() (*ServerConfig, error) {
	port, err := GetEnvVar(PORT)
	if err != nil {
		os.Exit(1)
	}
	return &ServerConfig{
		port: port,
	}, nil
}

func (s *ServerConfig) GetPort() string {
	return s.port
}
