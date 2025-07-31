package transaction

import (
	"context"
	"errors"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/transaction"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"gorm.io/gorm"
)

func (u *usecase) Get(ctx context.Context, req ucModel.TransactionGet) (ucModel.TransactionResponse, error) {
	var (
		response = ucModel.TransactionResponse{}
		orm      = u.resource.Postgres.DB
	)

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[TransactionUseCase] Get: failed to begin transaction: %v", tx.Error)
		return response, tx.Error
	}

	transaction, err := u.Repository.Transaction.Get(ctx, tx, obModel.Transaction{
		UserID: req.UserID,
		TrxID:  req.TrxID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[TransactionUseCase] Get: failed to get transaction: %v", err)
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, pkgErr.ErrTransactionNotFound
		}

		return response, err
	}

	response = ucModel.TransactionResponse{
		TrxId:              transaction.TrxID,
		WalletName:         transaction.Wallet.Name,
		CurrencyCode:       transaction.Currency.Code,
		Amount:             transaction.Amount,
		Type:               transaction.Type,
		RefrenceWalletName: transaction.RefrenceWallet.Name,
	}

	return response, nil
}

type IGet interface {
	Get(ctx context.Context, req ucModel.TransactionGet) (ucModel.TransactionResponse, error)
}
