package exchange_rate

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, exchangeRate model.ExchangeRate) (model.ExchangeRate, error) {
	var (
		result model.ExchangeRate
		err    error
	)

	if err = orm.WithContext(ctx).Where("from_currency_id = ? AND to_currency_id = ?", exchangeRate.FromCurrencyId, exchangeRate.ToCurrencyId).First(&result).Error; err != nil {
		return model.ExchangeRate{}, err
	}

	return result, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, exchangeRate model.ExchangeRate) (model.ExchangeRate, error)
}
