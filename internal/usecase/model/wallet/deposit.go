package wallet

// obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"

type (
	DepositRequest struct {
		UserID       int64   `json:"user_id"`
		WalletID     int64   `json:"wallet_id"`
		CurrencyCode string  `json:"currency_code"`
		Amount       float64 `json:"amount"`
	}

	DepositResponse struct {
		TrxID   string  `json:"trx_id"`
		Balance float64 `json:"balance"`
	}
)

// func (DepositRequest) ToObModel() obModel.WalletBalance {}
