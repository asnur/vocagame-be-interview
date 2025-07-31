package wallet

import (
	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
)

type (
	CreateWalletRequest struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
	}

	CreateWalletResponse struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

func (w CreateWalletRequest) ToObModel(currencies []obModel.Currencies) obModel.Wallets {
	initBalance := 0.0 // Initial balance set to 0
	balance := make([]obModel.WalletBalance, len(currencies))
	for i, currency := range currencies {
		balance[i] = obModel.WalletBalance{
			CurrencyID: currency.ID,
			Balance:    &initBalance, // Initial balance is set to 0
		}
	}

	return obModel.Wallets{
		Name:          w.Name,
		UserID:        w.UserID,
		WalletBalance: balance,
	}
}
