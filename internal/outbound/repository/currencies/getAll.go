package currencies

import (
	"context"

	model "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) GetAll(ctx context.Context, orm *gorm.DB, req model.Currencies) ([]model.Currencies, error) {
	var currencies []model.Currencies

	if err := orm.WithContext(ctx).Where(req).Find(&currencies).Error; err != nil {
		r.resource.Logger.Errorf("[CurrenciesRepository] GetAll: %v", err)
		return nil, err
	}

	return currencies, nil
}

type IGetAll interface {
	GetAll(ctx context.Context, orm *gorm.DB, req model.Currencies) ([]model.Currencies, error)
}
