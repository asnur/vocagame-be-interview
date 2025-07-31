package di

import (
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository"
	"github.com/asnur/vocagame-be-interview/internal/usecase"
	"github.com/asnur/vocagame-be-interview/pkg/config"
	"github.com/asnur/vocagame-be-interview/pkg/resource"
	"go.uber.org/dig"
)

func Container() (*dig.Container, error) {
	container := dig.New()

	// Register Config
	if err := config.Register(container); err != nil {
		return nil, err
	}

	// Register Resource
	if err := resource.Register(container); err != nil {
		return nil, err
	}

	// Register Repositories
	if err := repository.Register(container); err != nil {
		return nil, err
	}

	// Register Use Cases
	if err := usecase.Register(container); err != nil {
		return nil, err
	}

	return container, nil
}
