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

func (u *usecase) Deposit(ctx context.Context, req ucModel.DepositRequest) (ucModel.DepositResponse, error) {
	var (
		orm      = u.resource.Postgres.DB
		response ucModel.DepositResponse
		amount   = 0.0
	)

	if req.Amount <= 0 {
		return response, pkgErr.ErrInvalidAmount
	}

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	// Validate Ownership Wallet
	_, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID:     req.WalletID,
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: failed to get wallet: %v", tx.Error)
		tx.Rollback()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}

		return response, err
	}

	// Get Currency
	currency, err := u.Repository.Currencies.Get(ctx, tx, obModel.Currencies{Code: req.CurrencyCode})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: %v", err)
		tx.Rollback()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}

		return response, err
	}

	// Get Last Balance
	lastBalance, err := u.Repository.WalletBalance.Get(ctx, tx, obModel.WalletBalance{
		WalletID:   req.WalletID,
		CurrencyID: currency.ID,
		Locking:    true, // Locking to prevent
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
	}

	// Calculate New Balance
	amount = *lastBalance.Balance + req.Amount

	// Update Wallet Balance
	_, err = u.Repository.WalletBalance.Update(ctx, tx, obModel.WalletBalance{
		WalletID:   req.WalletID,
		CurrencyID: currency.ID,
		Balance:    &amount,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: %v", err)
		tx.Rollback()

		return response, err
	}

	// Create Transaction Record
	trxId := fmt.Sprintf("%s-%d-%d", strings.ToUpper(pkgConstant.Deposit), req.WalletID, time.Now().UnixNano())
	transaction := obModel.Transaction{
		TrxID:      trxId,
		UserID:     req.UserID,
		WalletID:   req.WalletID,
		CurrencyID: currency.ID,
		Type:       pkgConstant.Deposit,
		Amount:     amount,
	}

	_, err = u.Repository.Transaction.Create(ctx, tx, transaction)
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: %v", err)
		tx.Rollback()

		if strings.Contains(err.Error(), pkgErr.ErrDuplicate.Error()) {
			return response, pkgErr.ErrIdentityAlreadyExists
		}

		return response, err
	}

	// Commit the transaction if currency is found
	if err := tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Deposit: failed to commit transaction: %v", err)
		return response, err
	}

	response = ucModel.DepositResponse{
		TrxID:   trxId,
		Balance: amount,
	}

	return response, nil
}

type IDeposit interface {
	Deposit(ctx context.Context, req ucModel.DepositRequest) (ucModel.DepositResponse, error)
}
