package transaction

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, transaction model.Transaction) (model.Transaction, error) {
	var (
		result model.Transaction
		err    error
	)

	if err = orm.WithContext(ctx).Preload(clause.Associations).Where(transaction).First(&result).Error; err != nil {
		return model.Transaction{}, err
	}

	return result, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, transaction model.Transaction) (model.Transaction, error)
}
