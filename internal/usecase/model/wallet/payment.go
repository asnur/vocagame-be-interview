package wallet

type (
	PaymentRequest struct {
		UserID       int64   `json:"user_id"`
		WalletID     int64   `json:"wallet_id"`
		CurrencyCode string  `json:"currency_code"`
		Amount       float64 `json:"amount"`
		Description  string  `json:"description"`
	}

	PaymentResponse struct {
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
	}
)
