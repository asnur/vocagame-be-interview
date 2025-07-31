package main

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/inbound"
	"github.com/asnur/vocagame-be-interview/pkg/di"
	"github.com/asnur/vocagame-be-interview/pkg/resource"
)

func main() {
	// Init Container
	container, err := di.Container()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(resource resource.Resource, inbound inbound.Inbound) error {
		// Initialize Routes
		inbound.Http.Routes(resource.Server.App)

		// Start Server
		if err := resource.Server.Start(); err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
