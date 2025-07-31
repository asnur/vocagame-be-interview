package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	BalanceRequest struct {
		WalletID int64 `params:"wallet_id" validate:"required"`
	}

	BalanceTotalRequest struct {
		WalletID     int64  `params:"wallet_id" validate:"required"`
		CurrencyCode string `query:"currency_code"`
	}
)

func (b BalanceRequest) ToUcModel(userId int64) ucModel.BalanceRequest {
	return ucModel.BalanceRequest{
		UserID:   userId,
		WalletID: b.WalletID,
	}
}

func (b BalanceTotalRequest) ToUcModel(userId int64) ucModel.BalanceTotalRequest {
	return ucModel.BalanceTotalRequest{
		UserID:       userId,
		WalletID:     b.WalletID,
		CurrencyCode: b.CurrencyCode,
	}
}
