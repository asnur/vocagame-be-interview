package model

type (
	Users struct {
		BaseModel
		ID       int64  `gorm:"column:id;primaryKey" json:"id"`
		Username string `gorm:"column:username;uniqueIndex" json:"username"`
		Email    string `gorm:"column:email;uniqueIndex" json:"email"`
		Password string `gorm:"column:password" json:"-"`
	}
)

func (Users) TableName() string {
	return "users"
}
