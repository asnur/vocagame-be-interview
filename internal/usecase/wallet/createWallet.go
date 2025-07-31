package wallet

import (
	"context"
	"strings"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
)

func (u *usecase) CreateWallet(ctx context.Context, req ucModel.CreateWalletRequest) (ucModel.CreateWalletResponse, error) {
	var (
		orm      = u.resource.Postgres.DB
		response ucModel.CreateWalletResponse
	)

	tx := orm.WithContext(ctx).Begin()
	if tx.Error != nil {
		u.resource.Logger.Errorf("[WalletUseCase] CreateWallet: %v", tx.Error)
		return response, pkgErr.ErrInternalServerError
	}

	// Get All Currencies
	currencies, err := u.Repository.Currencies.GetAll(ctx, tx, obModel.Currencies{})
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] CreateWallet: %v", err)
		tx.Rollback()

		return response, err
	}

	wallet, err := u.Repository.Wallet.Create(ctx, tx, req.ToObModel(currencies))
	if err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] CreateWallet: %v", err)
		tx.Rollback()

		if strings.Contains(err.Error(), pkgErr.ErrDuplicate.Error()) {
			return response, pkgErr.ErrIdentityAlreadyExists
		}

		return response, err
	}

	if err := tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[WalletUseCase] CreateWallet: failed to commit transaction: %v", err)
		return response, err
	}

	// Prepare response
	response = ucModel.CreateWalletResponse{
		ID:   wallet.ID,
		Name: wallet.Name,
	}

	return response, nil
}

type ICreateWallet interface {
	CreateWallet(ctx context.Context, req ucModel.CreateWalletRequest) (ucModel.CreateWalletResponse, error)
}
