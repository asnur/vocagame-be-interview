package transaction

import (
	"context"
	"errors"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/transaction"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"gorm.io/gorm"
)

func (u *usecase) GetAll(ctx context.Context, req ucModel.TransactionGetAllRequest) ([]ucModel.TransactionResponse, error) {
	var (
		orm = u.resource.Postgres.DB
	)

	// Call the repository method to get all transactions
	transactions, err := u.Repository.Transaction.GetAll(ctx, orm, obModel.Transaction{
		UserID: req.UserID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[TransactionUseCase] GetAll: failed to get transactions: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkgErr.ErrTransactionNotFound
		}
		return nil, err
	}

	response := make([]ucModel.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		response[i] = ucModel.TransactionResponse{
			TrxId:              transaction.TrxID,
			WalletName:         transaction.Wallet.Name,
			CurrencyCode:       transaction.Currency.Code,
			Amount:             transaction.Amount,
			Type:               transaction.Type,
			RefrenceWalletName: transaction.RefrenceWallet.Name,
		}
	}

	return response, nil
}

type IGetAll interface {
	GetAll(ctx context.Context, req ucModel.TransactionGetAllRequest) ([]ucModel.TransactionResponse, error)
}
