package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	WithDrawlRequest struct {
		WalletID     int64   `json:"wallet_id" validate:"required"`
		CurrencyCode string  `json:"currency_code" validate:"required"`
		Amount       float64 `json:"amount" validate:"required,gt=0"`
	}
)

func (d WithDrawlRequest) ToUcModel(userId int64) ucModel.WithDrawlRequest {
	return ucModel.WithDrawlRequest{
		UserID:       userId,
		WalletID:     d.WalletID,
		CurrencyCode: d.CurrencyCode,
		Amount:       d.Amount,
	}
}
