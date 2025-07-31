package wallet

type (
	WithDrawlRequest struct {
		UserID       int64   `json:"user_id"`
		WalletID     int64   `json:"wallet_id"`
		CurrencyCode string  `json:"currency_code"`
		Amount       float64 `json:"amount"`
	}

	WithDrawlResponse struct {
		TrxID   string  `json:"trx_id"`
		Balance float64 `json:"balance"`
	}
)
