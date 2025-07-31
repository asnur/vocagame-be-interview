package outbound

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/outbound/repository"
	"go.uber.org/dig"
)

type (
	Outbound struct {
		dig.In

		Repository repository.Repository
	}
)

func Register(container *dig.Container) error {
	// Register Repository
	if err := repository.Register(container); err != nil {
		return fmt.Errorf("[DI] failed to register repositories: %w", err)
	}

	// You can register other outbound components here as needed
	return nil
}
