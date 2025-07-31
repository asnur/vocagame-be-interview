package config

import "github.com/asnur/vocagame-be-interview/pkg/utils"

type (
	ServerConfig struct {
		Host string `envconfig:"SERVER_HOST" required:"true"`
		Port string `envconfig:"SERVER_PORT" required:"true"`
	}
)

func NewServerConfig() (ServerConfig, error) {
	var config ServerConfig

	if err := utils.NewConfig(&config); err != nil {
		return ServerConfig{}, err
	}

	return config, nil
}
