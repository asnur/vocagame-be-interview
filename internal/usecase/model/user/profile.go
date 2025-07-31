package user

import obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"

type (
	ProfileRequest struct {
		UserID int64 `json:"user_id"`
	}

	ProfileResponse struct {
		ID       int64           `json:"id"`
		Username string          `json:"username"`
		Email    string          `json:"email"`
		Wallets  []ProfileWallet `json:"wallets,omitempty"`
	}

	ProfileWallet struct {
		ID       int64                  `json:"id"`
		Name     string                 `json:"name"`
		Balances []ProfileWalletBalance `json:"balances,omitempty"` // Changed to slice to accommodate multiple currencies
	}

	ProfileWalletBalance struct {
		CurrencyCode string  `json:"currency_code"`
		Balance      float64 `json:"balance"`
	}
)

func (u ProfileRequest) ToObUserModel() obModel.Users {
	return obModel.Users{
		ID: u.UserID,
	}
}
