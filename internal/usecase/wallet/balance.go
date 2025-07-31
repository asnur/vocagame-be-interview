package wallet

import (
	"context"
	"errors"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"gorm.io/gorm"
)

func (u *usecase) Balance(ctx context.Context, req ucModel.BalanceRequest) (ucModel.BalanceResponse, error) {
	var (
		response = ucModel.BalanceResponse{}
		orm      = u.resource.Postgres.DB
	)

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Balance: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	// Check Wallet Ownership
	wallet, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID:     req.WalletID,
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Balance: failed to get wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}
		return response, err
	}

	// Get Wallet Balance
	balance, err := u.Repository.WalletBalance.GetAll(ctx, tx, obModel.WalletBalance{
		WalletID: wallet.ID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Balance: failed to get wallet balance: %v", err)
		tx.Rollback()

		return response, err
	}

	if err = tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Balance: failed to commit transaction: %v", err)
		return response, err
	}

	response.Name = wallet.Name
	response.Balance = make([]ucModel.BalanceCurrency, len(balance))
	for i, bal := range balance {
		response.Balance[i] = ucModel.BalanceCurrency{
			CurrencyCode: bal.Currency.Code,
			Balance:      *bal.Balance,
		}
	}

	return response, nil
}

func (u *usecase) BalanceTotal(ctx context.Context, req ucModel.BalanceTotalRequest) (ucModel.BalanceTotalResponse, error) {
	var (
		response = ucModel.BalanceTotalResponse{}
		orm      = u.resource.Postgres.DB
	)

	if req.CurrencyCode == "" {
		req.CurrencyCode = u.resource.AppConfig.Currency
	}

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	// Check Wallet Ownership
	wallet, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID:     req.WalletID,
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to get wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}
		return response, err
	}

	// Get Balance All Currencies
	balance, err := u.Repository.WalletBalance.GetAll(ctx, tx, obModel.WalletBalance{
		WalletID: wallet.ID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to get total balance: %v", err)
		tx.Rollback()

		return response, err
	}

	// Get Currency
	currency, err := u.Repository.Currencies.Get(ctx, tx, obModel.Currencies{
		Code: req.CurrencyCode,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to get currency: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	// Get Exchange Rate
	exchangeRates, err := u.Repository.ExchangeRate.GetAll(ctx, tx, obModel.ExchangeRate{
		ToCurrencyId: currency.ID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to get exchange rates: %v", err)
		tx.Rollback()
		return response, err
	}

	if err = tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] BalanceTotal: failed to commit transaction: %v", err)
		return response, err
	}

	total := 0.0
	for _, bal := range balance {
		if bal.Currency.ID == currency.ID {
			// Same currency, no conversion needed
			total += *bal.Balance
		} else {
			// Different currency, need conversion
			for _, rate := range exchangeRates {
				if rate.FromCurrencyId == bal.Currency.ID && rate.ToCurrencyId == currency.ID {
					total += *bal.Balance * rate.Rate
					break
				}
			}
		}
	}

	response.Name = wallet.Name
	response.CurrencyCode = req.CurrencyCode
	response.Total = total

	return response, nil
}

type IBalance interface {
	Balance(ctx context.Context, req ucModel.BalanceRequest) (ucModel.BalanceResponse, error)
	BalanceTotal(ctx context.Context, req ucModel.BalanceTotalRequest) (ucModel.BalanceTotalResponse, error)
}
