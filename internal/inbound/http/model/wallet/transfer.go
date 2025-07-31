package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	TransferRequest struct {
		FromWalletID     int64   `json:"from_wallet_id" validate:"required"`
		ToWalletID       int64   `json:"to_wallet_id" validate:"required"`
		FromCurrencyCode string  `json:"from_currency_code" validate:"required"`
		ToCurrencyCode   string  `json:"to_currency_code" validate:"required"`
		Amount           float64 `json:"amount" validate:"required,gt=0"`
	}
)

func (t TransferRequest) ToUcModel(userId int64) ucModel.TransferRequest {
	return ucModel.TransferRequest{
		UserID:           userId,
		FromWalletID:     t.FromWalletID,
		ToWalletID:       t.ToWalletID,
		FromCurrencyCode: t.FromCurrencyCode,
		ToCurrencyCode:   t.ToCurrencyCode,
		Amount:           t.Amount,
	}
}
