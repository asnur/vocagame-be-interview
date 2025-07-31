package model

type Currencies struct {
	BaseModel
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Code string `json:"code" gorm:"type:varchar(3);unique;not null"`
	Name string `json:"name" gorm:"type:varchar(100);not null"`
}

func (Currencies) TableName() string {
	return "currencies"
}
