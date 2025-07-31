package resource

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/pkg/resource/injection"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Register Postgres
	if err := container.Provide(injection.NewPostgres); err != nil {
		return fmt.Errorf("[DI] failed to register PostgresConfig: %w", err)
	}
	// Register Logger
	if err := container.Provide(injection.NewLogger); err != nil {
		return fmt.Errorf("[DI] failed to register Logger: %w", err)
	}
	// Register Server
	if err := container.Provide(injection.NewServer); err != nil {
		return fmt.Errorf("[DI] failed to register Server: %w", err)
	}
	// Register Jwt
	if err := container.Provide(injection.NewJwt); err != nil {
		return fmt.Errorf("[DI] failed to register Jwt: %w", err)
	}

	return nil
}
