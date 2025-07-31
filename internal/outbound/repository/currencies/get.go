package currencies

import (
	"context"

	model "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, currency model.Currencies) (model.Currencies, error) {
	var foundCurrency model.Currencies

	if err := orm.WithContext(ctx).Where(currency).First(&foundCurrency).Error; err != nil {
		r.resource.Logger.Errorf("[CurrenciesRepository] Get: %v", err)
		return model.Currencies{}, err
	}

	return foundCurrency, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, currency model.Currencies) (model.Currencies, error)
}
