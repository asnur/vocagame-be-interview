package wallet

type (
	BalanceRequest struct {
		UserID   int64 `json:"user_id"`
		WalletID int64 `json:"wallet_id"`
	}

	BalanceResponse struct {
		Name    string            `json:"name"`
		Balance []BalanceCurrency `json:"balance"`
	}

	BalanceCurrency struct {
		CurrencyCode string  `json:"currency_code"`
		Balance      float64 `json:"balance"`
	}

	BalanceTotalRequest struct {
		UserID       int64  `json:"user_id"`
		WalletID     int64  `json:"wallet_id"`
		CurrencyCode string `json:"currency_code"`
	}

	BalanceTotalResponse struct {
		Name         string  `json:"name"`
		Total        float64 `json:"total"`
		CurrencyCode string  `json:"currency_code"`
	}
)
