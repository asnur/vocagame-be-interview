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

func (u *usecase) Payment(ctx context.Context, req ucModel.PaymentRequest) (ucModel.PaymentResponse, error) {
	var (
		response = ucModel.PaymentResponse{
			Status: "failed",
		}
		orm = u.resource.Postgres.DB
	)

	if req.Amount <= 0 {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: invalid amount %f", req.Amount)
		return response, pkgErr.ErrInvalidAmount
	}

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	// Validate Ownership Wallet
	wallet, err := u.Repository.Wallet.Get(ctx, tx, obModel.Wallets{
		ID:     req.WalletID,
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to get wallet: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrWalletNotFound
		}
		return response, err
	}

	// Get Currency
	currency, err := u.Repository.Currencies.Get(ctx, tx, obModel.Currencies{
		Code: req.CurrencyCode,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to get currency: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrCurrencyNotFound
		}
		return response, err
	}

	// Get Wallet Balance
	walletBalance, err := u.Repository.WalletBalance.Get(ctx, tx, obModel.WalletBalance{
		WalletID:   wallet.ID,
		CurrencyID: currency.ID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to get wallet balance: %v", err)
		tx.Rollback()

		return response, err
	}

	// Validate Wallet Balance
	if walletBalance.Balance == nil || *walletBalance.Balance < req.Amount {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: insufficient balance in wallet %s", wallet.ID)
		tx.Rollback()
		return response, pkgErr.ErrInsufficientBalance
	}

	// Calculate New Balance
	newBalance := *walletBalance.Balance - req.Amount
	// Update Wallet Balance
	walletBalance.Balance = &newBalance
	if _, err := u.Repository.WalletBalance.Update(ctx, tx, walletBalance); err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to update wallet balance: %v", err)
		tx.Rollback()
		return response, err
	}

	// Create Transaction
	trxId := fmt.Sprintf("%s-%d-%d", strings.ToUpper(pkgConstant.Payment), wallet.ID, time.Now().UnixNano())
	transaction := obModel.Transaction{
		TrxID:       trxId,
		UserID:      req.UserID,
		WalletID:    wallet.ID,
		CurrencyID:  currency.ID,
		Type:        pkgConstant.Payment,
		Amount:      req.Amount,
		Description: req.Description,
	}
	if _, err := u.Repository.Transaction.Create(ctx, tx, transaction); err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to create transaction: %v", err)
		tx.Rollback()
		return response, err
	}

	// Commit Transaction
	if err := tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] Payment: failed to commit transaction: %v", err)
		return response, err
	}

	response = ucModel.PaymentResponse{
		TransactionID: trxId,
		Status:        "success",
	}

	return response, nil
}

type IPayment interface {
	Payment(ctx context.Context, req ucModel.PaymentRequest) (ucModel.PaymentResponse, error)
}
