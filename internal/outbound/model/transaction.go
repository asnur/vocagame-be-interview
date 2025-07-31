package model

type (
	Transaction struct {
		BaseModel
		ID               int64      `gorm:"column:id;primaryKey" json:"id"`
		TrxID            string     `gorm:"column:trx_id;not null;unique" json:"trx_id"`
		UserID           int64      `gorm:"column:user_id;not null" json:"user_id"`
		WalletID         int64      `gorm:"column:wallet_id;not null"`
		Wallet           Wallets    `gorm:"foreignKey:WalletID;references:ID"`
		CurrencyID       int64      `gorm:"column:currency_id;not null"`
		Currency         Currencies `gorm:"foreignKey:CurrencyID;references:ID"`
		Amount           float64    `gorm:"column:amount;not null"`
		Type             string     `gorm:"column:type;not null"`
		Description      string     `gorm:"column:description;not null"`
		RefrenceWalletID *int64     `gorm:"column:refrence_wallet_id;null"`
		RefrenceWallet   Wallets    `gorm:"foreignKey:RefrenceWalletID"`
	}
)

func (Transaction) TableName() string {
	return "transactions"
}
