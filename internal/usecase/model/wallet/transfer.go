package wallet

type (
	TransferRequest struct {
		UserID           int64   `json:"user_id"`
		FromWalletID     int64   `json:"from_wallet_id"`
		ToWalletID       int64   `json:"to_wallet_id"`
		FromCurrencyCode string  `json:"from_currency_code"`
		ToCurrencyCode   string  `json:"to_currency_code"`
		Amount           float64 `json:"amount"`
	}

	TransferResponse struct {
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
	}
)
