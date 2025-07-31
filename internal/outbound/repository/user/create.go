package user

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Create(ctx context.Context, orm *gorm.DB, user model.Users) (model.Users, error) {
	if err := orm.WithContext(ctx).Create(&user).Error; err != nil {
		r.resource.Logger.Errorf("[UserRepository] Create: %v", err)
		return model.Users{}, err
	}

	return user, nil
}

type ICreate interface {
	Create(ctx context.Context, orm *gorm.DB, user model.Users) (model.Users, error)
}
