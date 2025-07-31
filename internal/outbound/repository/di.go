package repository

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/currencies"
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/exchange_rate"
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/transaction"
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/user"
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/wallet"
	"github.com/asnur/vocagame-be-interview/internal/outbound/repository/wallet_balance"
	"go.uber.org/dig"
)

type (
	Repository struct {
		dig.In

		ExchangeRate exchange_rate.Repository

		Transaction transaction.Repository

		User user.Repository

		Wallet wallet.Repository

		WalletBalance wallet_balance.Repository

		Currencies currencies.Repository
	}
)

func Register(container *dig.Container) error {
	// Register User Repository
	if err := container.Provide(user.New); err != nil {
		return fmt.Errorf("[DI] failed to register User Repository: %w", err)
	}

	// Register Wallet Repository
	if err := container.Provide(wallet.New); err != nil {
		return fmt.Errorf("[DI] failed to register Wallet Repository: %w", err)
	}

	// Register Currencies Repository
	if err := container.Provide(currencies.New); err != nil {
		return fmt.Errorf("[DI] failed to register Currencies Repository: %w", err)
	}

	// Register Wallet Balance Repository
	if err := container.Provide(wallet_balance.New); err != nil {
		return fmt.Errorf("[DI] failed to register Wallet Balance Repository: %w", err)
	}

	// Register Transaction Repository
	if err := container.Provide(transaction.New); err != nil {
		return fmt.Errorf("[DI] failed to register Transaction Repository: %w", err)
	}

	// Register Exchange Rate Repository
	if err := container.Provide(exchange_rate.New); err != nil {
		return fmt.Errorf("[DI] failed to register Exchange Rate Repository: %w", err)
	}

	return nil
}
