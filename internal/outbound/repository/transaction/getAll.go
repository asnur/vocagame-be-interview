package transaction

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) GetAll(ctx context.Context, orm *gorm.DB, transaction model.Transaction) ([]model.Transaction, error) {
	var (
		transactions []model.Transaction
		err          error
	)

	if err = orm.WithContext(ctx).Preload(clause.Associations).Where(transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

type IGetAll interface {
	GetAll(ctx context.Context, orm *gorm.DB, transaction model.Transaction) ([]model.Transaction, error)
}
