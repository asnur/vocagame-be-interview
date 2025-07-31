package config

import "github.com/asnur/vocagame-be-interview/pkg/utils"

type (
	PostgresConfig struct {
		PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true"`
		PostgresPort     string `envconfig:"POSTGRES_PORT" required:"true"`
		PostgresUsername string `envconfig:"POSTGRES_USERNAME" required:"true"`
		PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
		PostgresDB       string `envconfig:"POSTGRES_DATABASE" required:"true"`
	}
)

func NewPostgresConfig() (PostgresConfig, error) {
	var config PostgresConfig

	if err := utils.NewConfig(&config); err != nil {
		return config, err
	}

	return config, nil
}
