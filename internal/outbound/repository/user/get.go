package user

import (
	"context"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, user obModel.Users) (obModel.Users, error) {
	var userData obModel.Users

	if err := orm.WithContext(ctx).Preload(clause.Associations).Where(user).First(&userData).Error; err != nil {
		r.resource.Logger.Errorf("[UserRepository] Get: %v", err)
		return obModel.Users{}, err
	}

	return userData, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, user obModel.Users) (obModel.Users, error)
}
