package exchange_rate

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) GetAll(ctx context.Context, orm *gorm.DB, exchangeRate model.ExchangeRate) ([]model.ExchangeRate, error) {
	var (
		result []model.ExchangeRate
		err    error
	)

	if err = orm.WithContext(ctx).Where(exchangeRate).Find(&result).Error; err != nil {
		r.resource.Logger.Errorf("[ExchangeRateRepository] GetAll: failed to get exchange rates: %v", err)
		return nil, err
	}

	return result, nil
}

type IGetAll interface {
	GetAll(ctx context.Context, orm *gorm.DB, exchangeRate model.ExchangeRate) ([]model.ExchangeRate, error)
}
