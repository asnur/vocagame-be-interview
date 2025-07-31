package model

type (
	ExchangeRate struct {
		BaseModel
		ID             int64   `json:"id" gorm:"primaryKey;autoIncrement"`
		FromCurrencyId int64   `json:"from_currency_id" gorm:"not null"`
		ToCurrencyId   int64   `json:"to_currency_id" gorm:"not null"`
		Rate           float64 `json:"rate" gorm:"not null"`
	}
)

func (ExchangeRate) TableName() string {
	return "exchange_rates"
}
