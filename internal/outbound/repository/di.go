package repository

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/user"
	"go.uber.org/dig"
)

type (
	Repository struct {
		dig.In

		User user.Repository
	}
)

func Register(container *dig.Container) error {
	// Register User Repository
	if err := container.Provide(user.New); err != nil {
		return fmt.Errorf("[DI] failed to register User Repository: %w", err)
	}

	// You can register other repositories here as needed
	return nil
}
