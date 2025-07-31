package model

type (
	WalletBalance struct {
		BaseModel
		ID         int64 `gorm:"column:id;primaryKey" json:"id"`
		WalletID   int64 `gorm:"column:wallet_id;not null"`
		Wallet     Wallets
		CurrencyID int64 `gorm:"column:currency_id;not null"`
		Currency   Currencies
		Balance    *float64 `gorm:"column:balance;not null"`
		Locking    bool     `gorm:"-" json:"locking"`
	}
)

func (WalletBalance) TableName() string {
	return "wallet_balances"
}
