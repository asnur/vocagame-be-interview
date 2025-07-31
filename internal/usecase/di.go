package usecase

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/usecase/shared"
	"github.com/asnur/vocagame-be-interview/internal/usecase/user"
	"go.uber.org/dig"
)

type (
	UseCase struct {
		dig.In

		Shared shared.UseCase

		User user.UseCase
	}
)

func Register(container *dig.Container) error {
	// Register Shared
	if err := container.Provide(shared.New); err != nil {
		return fmt.Errorf("[DI] failed to register Shared Use Case: %w", err)
	}
	// Register User
	if err := container.Provide(user.New); err != nil {
		return fmt.Errorf("[DI] failed to register User Use Case: %w", err)
	}

	// You can register other use cases here as needed
	return nil
}
