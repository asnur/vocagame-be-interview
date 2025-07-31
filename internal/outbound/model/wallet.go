package model

type Wallets struct {
	BaseModel
	ID            int64  `gorm:"primaryKey;"`
	Name          string `gorm:"not null;column:name;"`
	UserID        int64  `gorm:"not null;column:user_id;"`
	User          Users
	WalletBalance []WalletBalance `gorm:"foreignKey:WalletID;references:ID;"`
}

func (Wallets) TableName() string {
	return "wallets"
}
