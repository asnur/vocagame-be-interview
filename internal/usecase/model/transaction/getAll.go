package transaction

type (
	TransactionGetAllRequest struct {
		UserID int64 `json:"user_id"`
	}

	TransactionGet struct {
		UserID int64  `json:"user_id"`
		TrxID  string `json:"trx_id"`
	}

	TransactionResponse struct {
		TrxId              string  `json:"trx_id"`
		WalletName         string  `json:"wallet_name"`
		CurrencyCode       string  `json:"currency_code"`
		Amount             float64 `json:"amount"`
		Type               string  `json:"type"`
		RefrenceWalletName string  `json:"reference_wallet_name,omitempty"`
	}
)
