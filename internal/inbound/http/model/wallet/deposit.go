package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	DepositRequest struct {
		WalletID     int64   `json:"wallet_id" validate:"required"`
		CurrencyCode string  `json:"currency_code" validate:"required"`
		Amount       float64 `json:"amount" validate:"required,gt=0"`
	}
)

func (d DepositRequest) ToUcModel(userId int64) ucModel.DepositRequest {
	return ucModel.DepositRequest{
		UserID:       userId,
		WalletID:     d.WalletID,
		CurrencyCode: d.CurrencyCode,
		Amount:       d.Amount,
	}
}
