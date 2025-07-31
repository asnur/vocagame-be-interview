package wallet

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
	pkgConstant "github.com/asnur/vocagame-be-interview/pkg/constants"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"gorm.io/gorm"
)

func (u *usecase) Transfer(ctx context.Context, req ucModel.TransferRequest) (ucModel.TransferResponse, error) {
	var (
		response = ucModel.TransferResponse{
			Status: "failed",
		}
		orm = u.resource.Postgres.DB
	)

	if req.Amount <= 0 {
		return response, pkgErr.ErrInvalidAmount
	}

	// Validate Transfer in the same currency and wallets
	if (req.FromCurrencyCode == req.ToCurrencyCode) && (req.FromWalletID == req.ToWalletID) {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: cannot transfer to the same wallet and currency")
		return response, pkgErr.ErrInvalidAmount
	}

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	// Validate Ownership Wallets
	fromWallet, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID:     req.FromWalletID,
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get from wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}

		return response, err
	}

	// Validate To Wallet
	toWallet, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID: req.ToWalletID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get to wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}

		return response, err
	}

	// Get Currencies From and To
	fromCurrency, err := u.Repository.Currencies.Get(ctx, tx, obModel.Currencies{
		Code: req.FromCurrencyCode,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get from currency: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	toCurrency, err := u.Repository.Currencies.Get(ctx, tx, obModel.Currencies{
		Code: req.ToCurrencyCode,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get to currency: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	// Get Last Balance From Wallet
	lastBalanceFrom, err := u.Repository.WalletBalance.Get(ctx, tx, obModel.WalletBalance{
		WalletID:   fromWallet.ID,
		CurrencyID: fromCurrency.ID,
		Locking:    true, // Locking to prevent concurrent updates
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get last balance from wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	// Check if sufficient balance
	if lastBalanceFrom.Balance == nil || *lastBalanceFrom.Balance < req.Amount {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: insufficient balance in from wallet")
		tx.Rollback()
		return response, pkgErr.ErrInsufficientBalance
	}

	// Calculate New Balance From Wallet
	newBalanceFrom := *lastBalanceFrom.Balance - req.Amount
	// Update Wallet Balance From
	_, err = u.Repository.WalletBalance.Update(ctx, tx, obModel.WalletBalance{
		WalletID:   fromWallet.ID,
		CurrencyID: fromCurrency.ID,
		Balance:    &newBalanceFrom,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to update from wallet balance: %v", err)
		tx.Rollback()
		return response, err
	}

	// Get Last Balance To Wallet
	lastBalanceTo, err := u.Repository.WalletBalance.Get(ctx, tx, obModel.WalletBalance{
		WalletID:   toWallet.ID,
		CurrencyID: toCurrency.ID,
		Locking:    true, // Locking to prevent concurrent updates
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get last balance to wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	// Get Exchange Rate
	rate := 1.0 // Default rate if same currency
	if fromCurrency.ID != toCurrency.ID {
		exchangeRate, err := u.Repository.ExchangeRate.Get(ctx, tx, obModel.ExchangeRate{
			FromCurrencyId: fromCurrency.ID,
			ToCurrencyId:   toCurrency.ID,
		})
		if err != nil {
			u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to get exchange rate: %v", err)
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response, pkgErr.ErrExchangeRateNotFound
			}
			return response, err
		}

		// Replace rate with the exchange rate from From to To currency
		rate = exchangeRate.Rate
	}

	// Convert Amount to To Currency
	convertedAmount := req.Amount * rate

	// Calculate New Balance To Wallet
	if lastBalanceTo.Balance == nil {
		lastBalanceTo.Balance = new(float64)
	}
	*lastBalanceTo.Balance += convertedAmount

	// Update Wallet Balance To
	_, err = u.Repository.WalletBalance.Update(ctx, tx, obModel.WalletBalance{
		WalletID:   toWallet.ID,
		CurrencyID: toCurrency.ID,
		Balance:    lastBalanceTo.Balance,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to update to wallet balance: %v", err)
		tx.Rollback()
		return response, err
	}

	// Create Transaction Record
	trxIdFrom := fmt.Sprintf("%s-%d-%d", strings.ToUpper(pkgConstant.Transfer), fromWallet.ID, time.Now().UnixNano())
	transaction := obModel.Transaction{
		TrxID:            trxIdFrom,
		Type:             pkgConstant.Transfer,
		WalletID:         fromWallet.ID,
		RefrenceWalletID: &toWallet.ID,
		Amount:           req.Amount,
		CurrencyID:       fromCurrency.ID,
	}
	if _, err = u.Repository.Transaction.Create(ctx, tx, transaction); err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to create transaction record: %v", err)
		tx.Rollback()
		return response, err
	}

	// Create Transaction Record for To Wallet
	trxIdTo := fmt.Sprintf("%s-%d-%d", strings.ToUpper(pkgConstant.Transfer), toWallet.ID, time.Now().UnixNano())
	transaction = obModel.Transaction{
		TrxID:            trxIdTo,
		Type:             pkgConstant.Transfer,
		WalletID:         toWallet.ID,
		RefrenceWalletID: &fromWallet.ID,
		Amount:           convertedAmount,
		CurrencyID:       toCurrency.ID,
	}
	if _, err = u.Repository.Transaction.Create(ctx, tx, transaction); err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to create transaction record for to wallet: %v", err)
		tx.Rollback()
		return response, err
	}

	// Commit Transaction
	if err := tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Transfer: failed to commit transaction: %v", err)
		return response, err
	}

	response = ucModel.TransferResponse{
		TransactionID: trxIdFrom,
		Status:        "success",
	}

	return response, nil
}

type ITransfer interface {
	Transfer(ctx context.Context, req ucModel.TransferRequest) (ucModel.TransferResponse, error)
}
