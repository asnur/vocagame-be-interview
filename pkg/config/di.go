package config

import (
	"fmt"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Register AppConfig
	if err := container.Provide(NewAppConfig); err != nil {
		return fmt.Errorf("[DI] failed to register AppConfig: %w", err)
	}
	// Register PostgresConfig
	if err := container.Provide(NewPostgresConfig); err != nil {
		return fmt.Errorf("[DI] failed to register PostgresConfig: %w", err)
	}
	// Register ServerConfig
	if err := container.Provide(NewServerConfig); err != nil {
		return fmt.Errorf("[DI] failed to register ServerConfig: %w", err)
	}
	// Register TokenConfig
	if err := container.Provide(NewTokenConfig); err != nil {
		return fmt.Errorf("[DI] failed to register TokenConfig: %w", err)
	}

	// You can register other configurations here as needed
	return nil
}
