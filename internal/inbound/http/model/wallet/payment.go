package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	PaymentRequest struct {
		WalletID     int64   `json:"wallet_id" validate:"required"`
		CurrencyCode string  `json:"currency_code" validate:"required"`
		Amount       float64 `json:"amount" validate:"required,gt=0"`
		Description  string  `json:"description" validate:"required"`
	}
)

func (p PaymentRequest) ToUcModel(userId int64) ucModel.PaymentRequest {
	return ucModel.PaymentRequest{
		UserID:       userId,
		WalletID:     p.WalletID,
		CurrencyCode: p.CurrencyCode,
		Amount:       p.Amount,
		Description:  p.Description,
	}
}
