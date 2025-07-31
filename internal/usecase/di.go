package usecase

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/usecase/shared"
	"github.com/asnur/vocagame-be-interview/internal/usecase/transaction"
	"github.com/asnur/vocagame-be-interview/internal/usecase/user"
	"github.com/asnur/vocagame-be-interview/internal/usecase/wallet"
	"go.uber.org/dig"
)

type (
	UseCase struct {
		dig.In

		Shared shared.UseCase

		Transaction transaction.UseCase

		User user.UseCase

		Wallet wallet.UseCase
	}
)

func Register(container *dig.Container) error {
	// Register Shared
	if err := container.Provide(shared.New); err != nil {
		return fmt.Errorf("[DI] failed to register Shared Use Case: %w", err)
	}
	// Register Transaction
	if err := container.Provide(transaction.New); err != nil {
		return fmt.Errorf("[DI] failed to register Transaction Use Case: %w", err)
	}
	// Register User
	if err := container.Provide(user.New); err != nil {
		return fmt.Errorf("[DI] failed to register User Use Case: %w", err)
	}
	// Register Wallet
	if err := container.Provide(wallet.New); err != nil {
		return fmt.Errorf("[DI] failed to register Wallet Use Case: %w", err)
	}

	// You can register other use cases here as needed
	return nil
}
